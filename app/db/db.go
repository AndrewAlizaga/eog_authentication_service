package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//singleton poool conection
var PoolClient *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

//close mongo conection and cancel the context
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	///cancelfun to cancel to context
	defer cancel()

	defer func() {

		//client disconnect method has deadline
		//a mongodb conection.
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

//Ping to db
func Ping(client *mongo.Client, ctx context.Context) error {

	//mongo client has ping to ping mongodb, deadline of ping was previously defined by ctx
	// ping method returns error if any occurred
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("ping successfully")
	return nil
}

func Insert(col string, doc interface{}) (*mongo.InsertOneResult, error) {
	//

	//mongodb client from singleton

	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	fmt.Printf("DB METHOD INSERTION")
	collection := client.Database("local").Collection(col)

	fmt.Println("COLLECTION INSERTION")
	result, err := collection.InsertOne(context.TODO(), doc)
	return result, err

}

//TODO:
func InsertMany(col string, docs []interface{}) (*mongo.InsertManyResult, error) {
	//

	//mongodb client from singleton

	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	fmt.Printf("DB METHOD INSERTION")
	collection := client.Database("local").Collection(col)

	fmt.Println("COLLECTION INSERTION")
	result, err := collection.InsertMany(context.TODO(), docs)

	if err != nil {
		return nil, err
	}

	return result, err

}

//TODO:
func Update(col string, filter interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	//mongodb client from singleton

	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	fmt.Printf("DB METHOD INSERTION")
	collection := client.Database("local").Collection(col)

	fmt.Println("COLLECTION INSERTION")
	result, err := collection.UpdateOne(context.TODO(), filter, doc)
	return result, err

}

//TODO
func Delete(col string, filter interface{}, doc interface{}) (*mongo.DeleteResult, error) {
	//mongodb client from singleton

	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	fmt.Printf("DB METHOD INSERTION")
	collection := client.Database("local").Collection(col)

	fmt.Println("COLLECTION INSERTION")
	result, err := collection.DeleteOne(context.TODO(), filter)
	return result, err

}

//TODO
func FindOne(col string, filter interface{}) (*mongo.SingleResult, error) {
	//mongodb client from singleton

	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	fmt.Printf("DB METHOD INSERTION")
	collection := client.Database("local").Collection(col)

	fmt.Println("COLLECTION INSERTION")
	result := collection.FindOne(context.TODO(), filter)

	return result, nil

}

//TODO
func FindMany(col string, doc interface{}) (*mongo.Cursor, error) {
	//mongodb client from singleton

	client, err := GetMongoClient()

	if err != nil {
		return nil, err
	}

	fmt.Printf("DB METHOD INSERTION")
	collection := client.Database("local").Collection(col)

	fmt.Println("COLLECTION INSERTION")
	result, err := collection.Find(context.TODO(), doc)
	return result, err

}

func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		//clientOptions := options.Client().ApplyURI("mongodb:27017")

		connectionString := os.Getenv("MONGO_CONNECTION_STRING")

		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPIOptions)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
			fmt.Println("ERROR: ", clientInstanceError)
			return
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
			fmt.Println("ERROR: ", clientInstanceError)
			return
		}
		PoolClient = client
	})
	return PoolClient, clientInstanceError
}
