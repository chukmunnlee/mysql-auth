package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	opts := ParseOptions()

	server := echo.New()

	if opts.CORS {
		log.Printf("Enabling CORS")
		server.Use(middleware.CORS())
	}

	if opts.Logger {
		log.Printf("Enabling Logging")
		server.Use(middleware.Logger())
	}

	api := server.Group("/api")

	api.GET("/", func(ctx echo.Context) error {
		msg := fmt.Sprintf("<h1>The current time is %s</h1>", time.Now().Format(time.RFC850))
		return ctx.HTML(http.StatusOK, msg)
	})

	api.GET("/healthz", func(ctx echo.Context) error {
		resp := OKResponse{
			OK:      true,
			Message: "",
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	log.Println(fmt.Sprintf("Starting server on port %s", opts.Port))

	if err := server.Start(fmt.Sprintf(":%s", opts.Port)); nil != err {
		log.Panicf("Cannot start server: %s\n", err.Error())
	}
}
