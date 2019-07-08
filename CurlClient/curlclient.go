package main

import (
	"net/http"
)

var (
	httpMethod            string
	body                  string
	followRedirects       bool
	httpHeaders           http.Header
	saveOutput            bool
	outputFile            string
	redirectFollowedCount int
)

const (
	defaultUrlScheme = "http"
	maxRedirects     = 10
)

func init() {
	//flag.Var(&httpHeaders,"H","Setup http headers")
}

func main() {

}
