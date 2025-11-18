package main

import (
	"fmt"
	"url-shortener/docs"
	"url-shortener/internal/api/httpredirect"
	"url-shortener/internal/api/httpurlshortener"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	"url-shortener/internal/service/redirect"
	"url-shortener/internal/service/urlshortener"

	"url-shortener/tools/config"
	"url-shortener/tools/dbutil"

	"github.com/labstack/echo/v4"

	_ "url-shortener/docs"

	"github.com/labstack/echo/v4/middleware"
)

// @title URL Shortener API
// @version 1.0
// @description API for encoding and decoding short URLs.
// @host localhost:8081
// @BasePath /
func main() {
	// Initialize Echo instance
	cfg := config.Load()
	db, err := dbutil.New(&cfg.DB)
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Host = cfg.App.BaseURL // OR read from env directly
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "Swagger for APIs"
	docs.SwaggerInfo.Description = "Swagger API Documentation for the url shortener."
	docs.SwaggerInfo.Version = "1.0"

	// TODO: Migrate to migrations with gormigrate
	db.AutoMigrate(&model.URLShortener{})

	// Init repos
	repos := repository.NewRepository(db)
	e := echo.New()

	// Logger
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Declaring echo groups
	rg := e.Group("/r")
	g := e.Group("/url-shortener")

	// Define services
	urlshortenerSvc := urlshortener.NewService(
		repos.UrlShortener,
		cfg.App.SBaseURL,
	)
	redirectSvc := redirect.NewService()

	httpurlshortener.NewHTTP(urlshortenerSvc, g)
	httpredirect.NewHTTP(redirectSvc, urlshortenerSvc, rg)

	// Start server on port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.App.Port)))
}
