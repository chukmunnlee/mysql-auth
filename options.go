package main

import "flag"

type Options struct {
	Port   string
	CORS   bool
	Logger bool
}

func ParseOptions() *Options {

	var port string
	var cors bool
	var logger bool

	flag.StringVar(&port, "port", "5000", "port number")
	flag.BoolVar(&cors, "cors", true, "enable cors")
	flag.BoolVar(&logger, "log", true, "enable logging")
	flag.Parse()

	return &Options{
		Port:   port,
		CORS:   cors,
		Logger: logger,
	}
}
