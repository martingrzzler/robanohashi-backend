package main

import (
	"context"
	"log"
	"net/http"
	"robanohashi/api/controllers"
	"robanohashi/persist"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	db, err := persist.Connect()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	authorized := r.Group("")
	authorized.Use(ValidateFirebaseToken(auth))

	authorized.POST("/meaning_mnemonic", controllers.CreateMeaningMnemonic)
	authorized.POST("/meaning_mnemonic/vote", controllers.VoteMeaningMnemonic)

	r.GET("/search", controllers.Search)
	r.GET("/kanji/:id", controllers.GetKanji)
	r.GET("/subject/:id/meaning_mnemonics", controllers.GetMeaningMnemonics)
	r.GET("/radical/:id", controllers.GetRadical)
	r.GET("/vocabulary/:id", controllers.GetVocabulary)
	r.Run(":5000")
}

func ValidateFirebaseToken(auth *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		bearer := c.GetHeader("Authorization")

		if bearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		idToken := strings.Split(bearer, " ")[1]

		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		uid := token.UID

		c.Set("uid", uid)
	}
}
