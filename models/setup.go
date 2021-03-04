package models

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	var (
		DBConnection = os.Getenv("DB_CONNECTION")
		DBHost       = os.Getenv("DB_HOST")
		DBPort       = os.Getenv("DB_PORT")
		DBUsername   = os.Getenv("DB_USERNAME")
		DBPassword   = os.Getenv("DB_PASSWORD")
	)

	// connect to the database
	fmt.Println("Connecting to MongoDB")
	dbUrl := DBConnection + "://" + DBUsername + ":" + DBPassword + "@" + DBHost + ":" + DBPort
	fmt.Println(dbUrl)
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DB = client.Database(os.Getenv("DB_DATABASE"))
}

func Response(c *gin.Context, status bool, data interface{}) {
	var httpStatus int
	var res interface{}
	if status == true {
		httpStatus = http.StatusOK
		res = gin.H{
			"status": status,
			"data":   data,
		}
	} else {
		httpStatus = http.StatusBadRequest
		res = gin.H{
			"status": status,
			"error":  data,
		}
	}

	c.JSON(httpStatus, res)
	return
}
