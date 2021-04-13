package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	CONTENT_TYPE     = "Content-Type"
	FORM_URL_ENCODED = "application/x-www-form-urlencoded"
	JSON             = "application/json"
)

func main() {

	startedOn := time.Now()

	opts := ParseOptions()

	// Connect to database
	authDB := AuthDatabase(opts.DSN)
	if err := authDB.Open(); nil != err {
		log.Panicf("Error. Cannot connect to database. \t%s\n", err.Error())
	}
	defer authDB.Close()

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

	api.POST("/authz", func(ctx echo.Context) error {
		var authzResp PostAuthzResponse

		req := ctx.Request()
		contentType := req.Header.Get(CONTENT_TYPE)

		data := User{}
		if strings.HasPrefix(contentType, FORM_URL_ENCODED) {
			data.Username = req.FormValue("username")
			data.Password = req.FormValue("password")
		} else if strings.HasPrefix(contentType, JSON) {
			if err := ctx.Bind(&data); nil != err {
				authzResp = PostAuthzResponse{
					Message: "Content type not supported",
				}
				return ctx.JSON(http.StatusUnsupportedMediaType, authzResp)
			}
		} else {
			return ctx.JSON(http.StatusUnsupportedMediaType, authzResp)
		}

		v, err := authDB.Validate(data.Username, data.Password)
		if nil != err {
			authzResp = PostAuthzResponse{
				Message: err.Error(),
			}
			return ctx.JSON(http.StatusInternalServerError, authzResp)
		}

		if !v {
			authzResp = PostAuthzResponse{
				Message: "Failed authentication",
			}
			return ctx.JSON(http.StatusUnauthorized, authzResp)
		}

		authzResp = PostAuthzResponse{
			Message: fmt.Sprintf("Authenticated %s", data.Username),
		}

		return ctx.JSON(http.StatusOK, authzResp)
	})

	api.GET("/healthz", func(ctx echo.Context) error {
		var status OKResponse
		if err := authDB.Ping(); nil != err {
			status = OKResponse{
				Status:  false,
				Message: err.Error(),
			}
			return ctx.JSON(http.StatusInternalServerError, status)
		}
		status = OKResponse{
			Status:  true,
			Message: fmt.Sprintf("Uptime: %.2f hrs", time.Now().Sub(startedOn).Hours()),
		}
		return ctx.JSON(http.StatusOK, status)
	})

	log.Println(fmt.Sprintf("Starting server on port %s", opts.Port))

	if err := server.Start(fmt.Sprintf(":%s", opts.Port)); nil != err {
		log.Panicf("Cannot start server: %s\n", err.Error())
	}
}
