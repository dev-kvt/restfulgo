package api
import "github.com/gin-gonic/gin"

type Book struct{
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type JSONResponse struct{
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJSON(a *gin.Context , status int , message string , data any){
	response := JSONResponse{
		Status: status,
		Message: message,
		Data: data,
	}
	a.JSON(status , response)
}