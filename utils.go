package main

import "net/http"

func getHeader(name string, req *http.Request) string {
	return req.Header.Get(name)
}

func generateToken(username string, signKey string) {

}
