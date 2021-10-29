package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/georgechristman/golang-okta-rest-api-viper-config-demo/util"
	"github.com/gin-gonic/gin"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type Payload struct {
	Email  string   `json: "email"`
	Groups []string `json:"groups"`
}

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(config util.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		verifierSetup := jwtverifier.JwtVerifier{
			Issuer: config.OktaIssuer,
			ClaimsToValidate: map[string]string{
				"aud": config.Audience,
				"cid": config.OktaClientId,
			},
		}
		// Verify okta access token
		verifier := verifierSetup.New()
		jwtToken, err := verifier.VerifyAccessToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		payload, err := payloadToStruct(jwtToken.Claims)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func payloadToStruct(claims map[string]interface{}) (*Payload, error) {

	// Convert claims to json string
	jsonStr, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	payload := &Payload{}

	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func HasGroup(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
