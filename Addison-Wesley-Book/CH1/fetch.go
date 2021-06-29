// makes an http get request to the url passed as command-line argument
// Run in terminal
// go build fetch.go`
// Then
// ./fetch http://gopl.io
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
import "net/http"

func main() {
	url := os.Args[1]

	// make the request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured in http get %s", err.Error())
		os.Exit(1)
	}
	// The Body is of type io.ReadCloser and can be read by  ioutil.ReadAll
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err occured while reading the response body %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
