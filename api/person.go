package api

import (
	"errors"
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*Payload)

	hasAccess := authPayload.hasGroup("PORTAL_DASHBOARD_MERLIN_RPCI")

	if !hasAccess {
		err := errors.New("access denied")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, persons)
}

func (server *Server) getPerson(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*Payload)

	hasAccess := authPayload.hasGroup("PORTAL_DASHBOARD_MERLIN_RPC")

	if !hasAccess {
		err := errors.New("access denied")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, persons)
}

func (server *Server) updatePerson(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*Payload)

	hasAccess := authPayload.hasGroups([]string{"PORTAL_DASHBOARD_MERLIN_RPC", "PORTAL_DASHBOARD_MERLIN_RPCI"})

	if !hasAccess {
		err := errors.New("access denied")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, persons)
}
