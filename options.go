package main

import "flag"

type Options struct {
	Port   string
	CORS   bool
	Logger bool
	DSN    string
}

func ParseOptions() *Options {

	var port string
	var cors bool
	var logger bool
	var dsn string

	flag.StringVar(&port, "port", "5000", "port number")
	flag.BoolVar(&cors, "cors", true, "enable cors")
	flag.BoolVar(&logger, "log", true, "enable logging")
	flag.StringVar(&dsn, "dsn", "fred:fred@tcp(localhost:3306)/auth", "connection string for MySQL")

	flag.Parse()

	return &Options{
		Port:   port,
		CORS:   cors,
		Logger: logger,
		DSN:    dsn,
	}
}
