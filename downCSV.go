package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "https://docs.google.com/spreadsheets/d/1jfB4muWkzKTR0daklmf8D5F0Uf_IYAgcx_-Ij9McClQ/export?format=csv"
	Download(url, "sum.csv", 100)
}

func Download(url string, filename string, timeout int64) {
	client := http.Client{
		Timeout: time.Duration(timeout * int64(time.Second)),
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error url")
	}

	if resp.StatusCode != 200 {

		fmt.Println("error code error")
	}
	if resp.Header["Content-Type"][0] != "text/csv" {
		fmt.Println("error not csv")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error cant read body")
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		fmt.Println("error write file error")
	}
}
