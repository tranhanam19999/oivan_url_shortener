package main

import (
	"url-shortener/internal/api/httpurlshortener"
	"url-shortener/internal/repository"
	"url-shortener/internal/service/urlshortener"

	"url-shortener/tools/config"
	"url-shortener/tools/dbutil"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo instance
	cfg := config.Load()
	db, err := dbutil.New(&cfg.DB)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	e := echo.New()

	g := e.Group("/url-shortener")
	// Define routes
	urlshortenerSvc := urlshortener.NewService(
		&repos.UrlShortener,
	)

	httpurlshortener.NewHTTP(urlshortenerSvc, g)

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
