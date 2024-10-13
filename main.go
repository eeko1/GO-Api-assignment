package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Author      string             `json:"author" bson:"author"`
	Content     string             `json:"content" bson:"content"`
	PublishDate time.Time          `json:"publish_date" bson:"publish_date"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("your_database_name").Collection("articles")

	app := fiber.New()

	app.Get("/api/articles", getArticles)
	app.Post("/api/articles", postArticle)
	app.Put("/api/articles/:id", putArticle)
	app.Delete("/api/articles/:id", deleteArticle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getArticles(c *fiber.Ctx) error {
	var articles []Article

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var article Article
		if err := cursor.Decode(&article); err != nil {
			return err
		}
		articles = append(articles, article)
	}

	return c.JSON(articles)
}

func postArticle(c *fiber.Ctx) error {
	article := new(Article)

	if err := c.BodyParser(article); err != nil {
		return err
	}

	if article.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Title cannot be empty"})
	}

	article.PublishDate = time.Now()

	insertResult, err := collection.InsertOne(context.Background(), article)
	if err != nil {
		return err
	}

	article.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(fiber.Map{"message": "Article added successfully"})
}

func putArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid article ID"})
	}

	var articleUpdate Article
	if err := c.BodyParser(&articleUpdate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"title":        articleUpdate.Title,
			"author":       articleUpdate.Author,
			"content":      articleUpdate.Content,
			"publish_date": articleUpdate.PublishDate,
		},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update article"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
	}

	return c.JSON(fiber.Map{
		"message":        "Article updated successfully",
		"modified_count": result.ModifiedCount,
	})
}

func deleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"message": "Article deleted successfully"})
}
