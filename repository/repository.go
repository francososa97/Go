package repository

import (
	"context"
	"errors"
	"log"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID    string  `json:"id" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

type ProductRepository interface {
	GetAll(sortByPriceAsc bool) ([]Product, error)
	GetByID(id string) (*Product, error)
	Create(product *Product) error
	Update(id string, product *Product) error
	Delete(id string) error
}

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(connectionString, dbName, collectionName string) (*MongoRepository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &MongoRepository{Collection: collection}, nil
}

func (r *MongoRepository) GetAll(sortByPriceAsc bool) ([]Product, error) {
	ctx := context.Background()
	filter := bson.D{}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	if sortByPriceAsc {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	} else {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price
		})
	}

	return products, nil
}

func (r *MongoRepository) GetByID(id string) (*Product, error) {
	ctx := context.Background()
	filter := bson.D{{"_id", id}}
	var product Product
	err := r.Collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *MongoRepository) Create(product *Product) error {
	ctx := context.Background()
	_, err := r.Collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Update(id string, product *Product) error {
	ctx := context.Background()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", product}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(id string) error {
	ctx := context.Background()
	filter := bson.D{{"_id", id}}
	_, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
