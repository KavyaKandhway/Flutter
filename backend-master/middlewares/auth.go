package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"furrble.com/backend/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const valName = "FIREBASE_ID_TOKEN"

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// FirebaseAuthMiddleware is middleware for Firebase Authentication
type AuthMiddleware struct {
	cli          *auth.Client
	unAuthorized func(c *gin.Context)
}

// New is constructor of the middleware
func New(credFileName string, unAuthorized func(c *gin.Context)) (*AuthMiddleware, error) {
	opt := option.WithCredentialsFile(credFileName)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return &AuthMiddleware{
		cli:          auth,
		unAuthorized: unAuthorized,
	}, nil
}

/*
UserMiddlewares function to add auth
*/
func (am *AuthMiddleware) TokenValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		log.Print(clientToken)
		if clientToken == "" {
			log.Print("error1")
			logger.Logger.Errorf("Authorization token was not provided")
			respondWithError(c, 401, "Unauthorised Access")
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		log.Print(extractedToken)
		// Verify if the format of the token is correct
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
			log.Print(clientToken)
		} else {
			logger.Logger.Errorf("Incorrect Format of Authn Token")
			respondWithError(c, 401, "Incorrect Format of Authorization Token")
			return
		}

		// foundInBlacklist := IsBlacklisted(extractedToken[1])
		// TODO: Remove this and implement a isblacklisted function to secure the endpoint
		foundInBlacklist := false

		if foundInBlacklist == true {
			logger.Logger.Infof("Found in Blacklist")
			respondWithError(c, 401, "Invalid Token - Unauthorised access")
			return
		}

		idToken, err := am.cli.VerifyIDToken(context.Background(), clientToken)
		log.Print(err)
		if err != nil {
			if am.unAuthorized != nil {
				am.unAuthorized(c)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": http.StatusText(http.StatusUnauthorized),
				})
			}
			return
		}

		c.Set(valName, idToken)
		c.Next()
	}
}

// ExtractClaims extracts claims
func ExtractClaims(c *gin.Context) *auth.Token {
	idToken, ok := c.Get(valName)
	if !ok {
		return new(auth.Token)
	}
	return idToken.(*auth.Token)
}
