package repositories

import (
	"context"
	"time"

	"github.com/khris-xp/shop-ease-api/database"
	"github.com/khris-xp/shop-ease-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.GetCollection(database.DB, "products")
var productTimeout = 10 * time.Second

type ProductRepositoryInterface interface {
	CreateProduct(ctx context.Context, product models.Product) (string, error)
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (models.Product, error)
	UpdateProduct(ctx context.Context, id primitive.ObjectID, product models.Product) (string, error)
	DeleteProduct(ctx context.Context, id primitive.ObjectID) (string, error)
}

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		collection: database.GetCollection(database.DB, "products"),
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product models.Product) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, productTimeout)
	defer cancel()

	_, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	return "Product created successfully", nil
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, productTimeout)
	defer cancel()

	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id primitive.ObjectID) (models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, productTimeout)
	defer cancel()

	var product models.Product
	err := productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id primitive.ObjectID, product models.Product) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, productTimeout)
	defer cancel()

	_, err := productCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
	if err != nil {
		return "", err
	}

	return "Product updated successfully", nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, productTimeout)
	defer cancel()

	_, err := productCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return "", err
	}

	return "Product deleted successfully", nil
}
