package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte(os.Getenv("SECRET"))

func JWTMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		tokenstr:=c.GetHeader("Authorization")
		 if tokenstr == ""{
				ResponseJSON(c, http.StatusUnauthorized ,"Authorization Token required" , nil)
				c.Abort()
				return
			}
			//parse and validate the token 
			_, err := jwt.Parse(tokenstr ,func(token *jwt.Token) (interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return secretkey , nil
			})
			if err!=nil{
				ResponseJSON(c, http.StatusInternalServerError, "Could not generate token", nil)
				return
			}
			ResponseJSON(c, http.StatusOK, "Token generated successfully", gin.H{"token": tokenstr})

	} 
}

	
