package main

import (
	"context"
	"student/internal/composites"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("config initializing")
	cfg := config.GetLogger()
	logger.Info("mongodb composite initializing")
	router := httprouter.New()

	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthSource)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}
	logger.Info("author composite initializing")
	AuthorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	AuthorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}
	bookComposite.Handler.Register(router)

	// start application

}
