package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
	Author  string             `json:"author"`
}

// DATABASE INSTANCE
var collection *mongo.Collection

func MessageCollection(c *mongo.Database) {
	collection = c.Collection("messages")
}

func GetAllMessages(c *gin.Context) {
	messages := []Message{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all messages, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var message Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Messages",
		"data":    messages,
	})
	return
}

func CreateMessage(c *gin.Context) {
	var message Message
	c.BindJSON(&message)
	title := message.Title
	content := message.Content
	author := message.Author
	id := primitive.NewObjectID()

	newMessage := Message{
		ID:      id,
		Title:   title,
		Content: content,
		Author:  author,
	}
	_, err := collection.InsertOne(context.TODO(), newMessage)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
	return
}

func GetSingleMessage(c *gin.Context) {
	messageId := c.Param("messageId")
	objectId, err := primitive.ObjectIDFromHex(messageId)

	message := Message{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&message)
	if err != nil {
		log.Printf("Error while getting a single Message, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Message not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Message",
		"data":    message,
	})
	return
}

func DeleteMessage(c *gin.Context) {
	messageId := c.Param("messageId")
	objectId, err := primitive.ObjectIDFromHex(messageId)

	log.Printf(messageId)
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		log.Printf("Error while deleting a single message, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Message deleted successfully",
	})
	return
}

func EditMessage(c *gin.Context) {
	messageId := c.Param("messageId")
	objectId, err := primitive.ObjectIDFromHex(messageId)
	var message Message
	if err := c.ShouldBind(&message); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	newData := bson.M{
		"$set": bson.M{
			"title":   message.Title,
			"content": message.Content,
			"author":  message.Author,
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, newData)
	log.Println(err)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Message Edited Successfully",
	})
	return
}
