package main

import (
	"database/sql"
	"fmt"
	"github.com/georgechristman/golang-okta-viper-config-demo/util"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var todos []string
var toValidate = map[string]string{
	"aud": "api://hriportal",
	"cid": os.Getenv("HOME"),
}

func Lists(c *gin.Context) {
	if verify(c) {
		c.JSON(http.StatusOK, gin.H{"list": todos})
	}
}

func ListItem(c *gin.Context) {

	if verify(c) {
		errormessage := "Index out of range"
		indexstring := c.Param("index")
		if index, err := strconv.Atoi(indexstring); err == nil && index < len(todos) {
			c.JSON(http.StatusOK, gin.H{"item": todos[index]})
		} else {
			if err != nil {
				errormessage = "Number expected: " + indexstring
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
		}
	}
}

func AddListItem(c *gin.Context) {
	if verify(c) {
		item := c.PostForm("item")
		todos = append(todos, item)
		c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
	}
}

func verify(c *gin.Context) bool {
	status := true
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		verifierSetup := jwtverifier.JwtVerifier{
			Issuer: "https://apps-test.healthresearch.org/oauth2/ausxjstvaaJC9KDTR0h7",
			ClaimsToValidate: toValidate,
		}
		verifier := verifierSetup.New()
		_, err := verifier.VerifyAccessToken(token)
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			print(err.Error())
			status = false
		}
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
		status = false
	}
	return status
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
