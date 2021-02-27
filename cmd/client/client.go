package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/tcnksm/go-httpstat"
)

func doOk() {
	fmt.Println("--- Running Ok ---")

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Println("GET /", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Body read", err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(body))
	fmt.Println("")
}

func doTimeoutFromServer() {
	fmt.Println("--- Running Timeout from Server ---")

	resp, err := http.Get("http://localhost:8080/timeout")
	if err != nil {
		log.Println("GET /timeout", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(body))
	fmt.Println("")
}

func doTimeoutFromClient() {
	fmt.Println("\n--- Running Timeout from Client ---")

	client := &http.Client{
		Timeout: 1 * time.Minute,
	}
	resp, err := client.Get("http://localhost:8080/slow")
	if err != nil {
		log.Println("GET /slow", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(body))
	fmt.Println("")
}

func doWithTraceOnClient() {
	fmt.Println("\n--- Running Client With Trace ---")

	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := &http.Client{
		Timeout: 1 * time.Minute,
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	result.End(time.Now())

	fmt.Printf("%+v\n", result)
	fmt.Println(resp.Status)
	fmt.Println(string(body))
	fmt.Println("")
}

func main() {
	doOk()
	doTimeoutFromServer()
	doTimeoutFromClient()
	doWithTraceOnClient()
}
