package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var persons = []person{
	{Email: "george@testemail.com", FirstName: "George", LastName: "Christman"},
	{Email: "john@testemail.com", FirstName: "John", LastName: "Doe"},
}

func (server *Server) listPersons(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, persons)
}
