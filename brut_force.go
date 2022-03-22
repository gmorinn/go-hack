package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func checkResponse(res string) bool {
	for _, v := range _response {
		if strings.Contains(res, v) {
			return true
		}
	}
	return false
}

func brutForcePOST() {
	start := time.Now()
	queue := make(chan bool, _concurrencyLevel)
	var index int64

	scanner := getFile(_wordlistFile)
	for scanner.Scan() {
		queue <- true
		go func(value string) {
			defer func() { <-queue }()

			if err := scanner.Err(); err != nil { // check scan file
				log.Fatal(err)
			}

			if index >= _limit && _limit > -1 { // check limit
				fmt.Println("Limit: %d", _limit)
				log.Printf("Took %s", time.Since(start))
				os.Exit(0)
			}

			// set payload and request
			payload := []byte(fmt.Sprintf("{\"%s\":\"%s\"}", _payload, value))
			buffer := bytes.NewBuffer(payload)
			resp, err := http.Post(_url, "application/json", buffer)
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			defer resp.Body.Close()

			//Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("\nTest n°%d\t|\tpayload: %v\t|\tresponse: %v\n", index, string(payload), string(body))
			if checkResponse(string(body)) {
				log.Printf("Took %s", time.Since(start))
				os.Exit(0)
			}
			index++
		}(scanner.Text())
	}
	for i := 0; i < int(_concurrencyLevel); i++ {
		queue <- true
	}
	log.Printf("Took %s", time.Since(start))
}

func brutForceGET() {

}

func brutForcePUT() {

}

func brutForcePATCH() {

}

func brutForceDELETE() {

}
