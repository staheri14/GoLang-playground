package main

import (
	"fmt"
	"main/helper"
	"net/http"
)

func main() {
	resp, err := http.Get("https://web3.bamboohr.com/jobs/")
	helper.CheckError(err)

	fmt.Println(resp.Status)

}
