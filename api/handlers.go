//create a new book

package api
import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var DB *gorm.DB // pointer to the instance of GORM 

func Initialze(){
	// fetching the env file from the root of the project 
   err:=godotenv.Load()
    if err!=nil{
     log.Fatal("Failed to load the Database:" , err)
    }
    
    
// fetching the connection string from the env file 
connectStr := os.Getenv("DB_u")
 DB , err = gorm.Open(postgres.Open(connectStr) , &gorm.Config{})
 if err != nil {
 	log.Fatal("Failed to connect to the Database")
 }
 
 //migrate the schema 
if err := DB.AutoMigrate(&Book{}); err != nil {
log.Fatal("Failed to migrate schema:", err)

}
}
func CreateBook(a *gin.Context){
	var book Book
   if err:=a.ShouldBindJSON(&book); err != nil{
	 ResponseJSON(a, http.StatusBadRequest, "Invalid input", nil)
	 return
	} 
	DB.Create(&book)
	ResponseJSON(a, http.StatusCreated, "Book created successfully", book)
}
func GetBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)
	ResponseJSON(c, http.StatusOK, "Books retrieved successfully", books)
}
func GetBook(c *gin.Context) {
	var book Book
	if err := DB.First(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Book retrieved successfully", book)
}
func UpdateBook(c *gin.Context) {
	var book Book
	if err := DB.First(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// bind the request body
	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	DB.Save(&book)
	ResponseJSON(c, http.StatusOK, "Book updated successfully", book)
}
func DeleteBook(c *gin.Context) {
	var book Book
	if err := DB.Delete(&book, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Book deleted successfully", nil)
}
func GenerateJWT(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	if loginRequest.Username != "admin" || loginRequest.Password != "password" {
		ResponseJSON(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})
	// Sign the token
	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Token generated successfully", gin.H{"token": tokenString})
}
