package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(os.Getenv("MONGODB_URI")).
		SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to ping MongoDB: %v", err)
	}

	fmt.Println("[SUCCESS] connected to MongoDB!")
	return client, nil
}

func SaveDataToMongoDB(data Data) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to insert data into MongoDB: %v", err)
	}

	fmt.Println("[SUCCESS] data saved to MongoDB!")
	return nil
}

func GetData(username string) (Data, error) {
	client, err := GetMongoClient()
	if err != nil {
		return Data{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	filter := bson.M{
		"username": username,
	}
	var cityData Data
	err = collection.FindOne(context.Background(), filter).Decode(&cityData)
	if err != nil {
		return Data{}, fmt.Errorf("[ERROR] failed to fetch city data from MongoDB: %v", err)
	}

	return cityData, nil
}

func FetchDataFromMongoDB(username string) ([]Data, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	filter := bson.M{
		"username": username,
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var citiesData []Data
	err = cursor.All(context.Background(), &citiesData)
	if err != nil {
		return nil, err
	}

	return citiesData, nil
}

func AsyncSaveDataToMongoDB(data []Data, done chan<- bool) {
	client, err := GetMongoClient()
	if err != nil {
		fmt.Printf("[ERROR] failed to connect to MongoDB: %v\n", err)
		done <- true
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	_, err = collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Printf("[ERROR] failed to delete existing data from MongoDB: %v\n", err)
		done <- true
		return
	}

	var documents []interface{}
	for _, cityData := range data {
		documents = append(documents, cityData)
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Printf("[ERROR] failed to insert data into MongoDB: %v\n", err)
	}

	fmt.Println("[SUCCESS] data saved to MongoDB!")
	done <- true
}
