package server

import (
	"time"

	_ "github.com/Scrip7/imdb-api/docs" // Docs are generated by Swag CLI, must be imported.
	"github.com/Scrip7/imdb-api/server/handler"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

func GetFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		Prefork:           viper.GetBool("PREFORK"),
		EnablePrintRoutes: viper.GetBool("PRINT_ROUTES"),
		GETOnly:           true,
		ErrorHandler:      errorHandler,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Cache middleware
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") != "" && c.GetReqHeaders()["Refresh"] != ""
		},
		CacheHeader:          "X-Cache",
		Expiration:           viper.GetDuration("CACHE_DURATION") * time.Millisecond,
		StoreResponseHeaders: false,
		CacheControl:         false,
		MaxBytes:             viper.GetUint("CACHE_MAX_BYTES"),
		KeyGenerator: func(c *fiber.Ctx) string {
			queryString := c.Context().QueryArgs().String()
			if queryString == "" {
				return c.Path()
			}
			return c.Path() + "?" + queryString
		},
	}))

	t := app.Group("/title/:id")
	t.Get("/", handler.TitleIndex)
	t.Get("/keywords", handler.TitleKeywords)

	c := app.Group("/chart/")
	c.Get("box-office", handler.ChartBoxOffice)
	c.Get("common", handler.ChartCommon)
	c.Get("moviemeter", handler.ChartMovieMeter)

	app.Use(fallback)

	return app
}
