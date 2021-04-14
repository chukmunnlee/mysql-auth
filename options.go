package main

import (
	"flag"
	"log"
)

const (
	DEFAULT_PORT              = "5000"
	DEFAULT_DSN               = "fred:fred@tcp(localhost:3306)/auth"
	DEFAULT_AUTH_TOKEN_HEADER = "X-Auth-Token"
)

type Options struct {
	Port        string
	CORS        bool
	Logger      bool
	DSN         string
	TokenHeader string
	SignKey     string
}

func ParseOptions() *Options {

	var port string
	var cors bool
	var logger bool
	var dsn string
	var tokenHeader string
	var signKey string

	flag.StringVar(&port, "port", DEFAULT_PORT, "port number")
	flag.StringVar(&dsn, "dsn", DEFAULT_DSN, "connection string for MySQL")
	flag.StringVar(&tokenHeader, "token", DEFAULT_AUTH_TOKEN_HEADER, "token HTTP header")
	flag.StringVar(&signKey, "signKey", "", "token signing key")
	flag.BoolVar(&cors, "cors", true, "enable cors")
	flag.BoolVar(&logger, "log", true, "enable logging")

	flag.Parse()

	return &Options{
		Port:        port,
		CORS:        cors,
		Logger:      logger,
		DSN:         dsn,
		TokenHeader: tokenHeader,
		SignKey:     signKey,
	}
}

// business rules goes here
func (o *Options) Validate() {

	if 0 == len(o.SignKey) {
		log.Fatalln("Set token signing key")
	}

	if DEFAULT_DSN == o.DSN {
		log.Println("WARNING: Using default DSN")
	}
}
