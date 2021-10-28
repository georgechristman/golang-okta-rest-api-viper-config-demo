package api

import (
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
		verifier := verifierSetup.New()
		payload, err := verifier.VerifyAccessToken(accessToken)
		if err != nil {
			fmt.Printf("fail %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fmt.Printf("pass")
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
