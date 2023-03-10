package api

import (
	"context"
	"log"
	"robanohashi/api/controllers"
	"robanohashi/api/controllers/middleware"
	"robanohashi/internal/config"
	"robanohashi/persist"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func Create(cfg config.Config) *gin.Engine {

	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Use(cors.Default())

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
	authorized.Use(middleware.ValidateFirebaseToken(auth, true))

	authorized.POST("/meaning_mnemonic", controllers.CreateMeaningMnemonic)
	authorized.PUT("/meaning_mnemonic", controllers.UpdateMeaningMnemonic)
	authorized.DELETE("/meaning_mnemonic", controllers.DeleteMeaningMnemonic)
	authorized.POST("/meaning_mnemonic/vote", controllers.VoteMeaningMnemonic)
	authorized.POST("/meaning_mnemonic/toggle_favorite", controllers.ToggleFavoriteMeaningMnemonic)

	r.Use(middleware.ValidateFirebaseToken(auth, false))

	r.GET("/search", controllers.Search)
	r.GET("/kanji/:id", controllers.GetKanji)
	r.GET("/subject/:id/meaning_mnemonics", controllers.GetMeaningMnemonics)
	r.GET("/radical/:id", controllers.GetRadical)
	r.GET("/vocabulary/:id", controllers.GetVocabulary)

	return r
}
