package controllers

import (
	"github.com/Nazarii14/golang-auth-jwt/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	helper "golang-jwt-auth/helpers"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var userCollection *mongo.Collection = database.OpenConnection(database.Client, "users")
var validate = validator.New()

func HashPassword(password string) string {
	return password
}

func VerifyPassword(hashedPassword string, password string) bool {
	return hashedPassword == password
}

func SignUp(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}

func RefreshToken(c *gin.Context) {

}

func GetUsers(c *gin.Context) {

}

func GetUser(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.Error})
		}
	}
}
