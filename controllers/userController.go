package controllers

import (
	"context"
	"fmt"
	"go-gin/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetUsers(c *gin.Context) {
	cur, err := models.DB.Collection("users").Find(context.Background(), bson.M{})
	// .Collection("users").Find(context.Background(), bson.M{})
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}
	defer cur.Close(context.Background())

	var users []*models.User
	for cur.Next(context.Background()) {
		user := &models.User{} // satuan
		err := cur.Decode(user)
		if err != nil {
			models.Response(c, false, fmt.Sprintf("%v", err.Error()))
			return
		}
		users = append(users, user)
	}

	if users == nil {
		users = []*models.User{}
	}

	models.Response(c, true, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// Validasi
	type reqUser struct {
		Name string `json:"name" binding:"required"`
	}
	var userInput reqUser

	if err := c.ShouldBindJSON(&userInput); err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	// Create User
	user := models.User{ID: primitive.NewObjectID(), Name: userInput.Name}
	_, err := models.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   &user,
	})
}

func (uc *UserController) DetailUser(c *gin.Context) {
	// Validasi
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	var user *models.User
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   user,
		"id":     id,
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	// Validasi
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	var json models.User
	c.Bind(&json)

	_, err = models.DB.Collection("users").UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{
		"$set": bson.M{"name": json.Name, "isActive": json.IsActive},
	})
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	// Validasi
	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	_, err = models.DB.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		models.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
