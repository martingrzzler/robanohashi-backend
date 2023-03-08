package main

import (
	"context"
	"log"
	"net/http"
	"robanohashi/api/controllers"
	"robanohashi/internal/dto"
	"robanohashi/persist"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "robanohashi/docs"
)

// @title Roba no hashi API
// @description Query Kanji, Vocabulary, and Radicals with Mnemonics
// @version 1.0.0
// @host robanohashi.org
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {

	db, err := persist.Connect()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	authorized.Use(ValidateFirebaseToken(auth, true))

	authorized.POST("/meaning_mnemonic", controllers.CreateMeaningMnemonic)
	authorized.PUT("/meaning_mnemonic", controllers.UpdateMeaningMnemonic)
	authorized.DELETE("/meaning_mnemonic", controllers.DeleteMeaningMnemonic)
	authorized.POST("/meaning_mnemonic/vote", controllers.VoteMeaningMnemonic)
	authorized.POST("/meaning_mnemonic/toggle_favorite", controllers.ToggleFavoriteMeaningMnemonic)

	r.Use(ValidateFirebaseToken(auth, false))

	r.GET("/search", controllers.Search)
	r.GET("/kanji/:id", controllers.GetKanji)
	r.GET("/subject/:id/meaning_mnemonics", controllers.GetMeaningMnemonics)
	r.GET("/radical/:id", controllers.GetRadical)
	r.GET("/vocabulary/:id", controllers.GetVocabulary)
	r.Run(":5000")
}

func ValidateFirebaseToken(auth *auth.Client, abortOnError bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		bearer := c.GetHeader("Authorization")

		if bearer == "" && abortOnError {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthenticated"})
			return
		} else if bearer == "" && !abortOnError {
			c.Next()
			return
		}

		idToken := strings.Split(bearer, " ")[1]

		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthenticated"})
			return
		}

		uid := token.UID

		c.Set("uid", uid)
		c.Next()
	}
}
