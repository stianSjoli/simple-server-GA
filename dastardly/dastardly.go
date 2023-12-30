package main

import (
    "fmt"
	"flag"
	"github.com/google/go-github/v57/github"
)

func main() {
    // Get a greeting message and print it.
	bearerToken := flag.String("token", "t", "github bearer token")
	flag.Parse()
	
    //client := github.NewClient(nil).WithAuthToken("... your access token ...")
    fmt.Println(*bearerToken)
}
