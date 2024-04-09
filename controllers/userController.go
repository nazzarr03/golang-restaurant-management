package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nazzarr03/golang-restaurant-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func SignUp(c *gin.Context) {

}
