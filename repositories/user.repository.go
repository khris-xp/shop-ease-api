package repositories

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/khris-xp/shop-ease-api/configs"
	"github.com/khris-xp/shop-ease-api/database"
	"github.com/khris-xp/shop-ease-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = []byte(configs.EnvSecretKey())
)

var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")
var userTimeout = 10 * time.Second

type UserRepositoryInterface interface {
	RegisterUser(ctx context.Context, user models.User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error)
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: database.GetCollection(database.DB, "users"),
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	existingUser, err := userCollection.Find(ctx, bson.M{"email": user.Email})

	if existingUser.Next(ctx) {
		return "", err
	} else if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)
	user.Role = "customer"
	user.Cart = []models.Cart{}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *UserRepository) LoginUser(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *UserRepository) GetUserProfile(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) RefreshToken(ctx context.Context, token string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, userTimeout)
	defer cancel()

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", err
	}

	user, err := r.GetUserProfile(ctx, email)
	if err != nil {
		return "", err
	}

	newToken := jwt.New(jwt.SigningMethodHS256)
	newClaims := newToken.Claims.(jwt.MapClaims)
	newClaims["email"] = user.Email
	newClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := newToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
