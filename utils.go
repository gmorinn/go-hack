package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// HELP ./gm -h
func help() {
	fmt.Println("GM HACK (v1)")
	fmt.Println("\t./gm {...flags}")

	fmt.Println("\nDESCRIPTION")
	fmt.Println("\tA perfect tool for hacking in Go")
	fmt.Println("\n\t-u, --url\t{URL_API} Url of the API")
	fmt.Println("\t-t, --threads\t{THREADS} Number of concurrent threads (default=3)")
	fmt.Println("\t-w, --wordlist\t{WORDLIST} Path to the wordlist")
	fmt.Println("\t-m, --method\t{METHOD} Method of the request (GET, POST, PUT, DELETE, PATCH) default=GET")
	fmt.Println("\t-p, --payload\t{PAYLOAD} The attribute you want to modify. Example: -p email -w email.txt")
	fmt.Println("\t-l, --limit\t{LIMIT} Stop the program when the number of request equals ${limit}")
	fmt.Println("\t-r, --response\t{RESPONSE} Corresponds to the attributes returned by the API. If one of these attributes is returned, then the program stops. Example: -r \"success,access_token,refresh_token\"")

	fmt.Println("\nNOTICE !")
	fmt.Println("\t- Only one url is taken into account.")
	fmt.Println("\t- Only one attribute is possible in the payload for version 1.")
	fmt.Println("\t- All flags are not mandatory, except -u && -w && (-r || -l)")

}

// Return message error + exit 84
func exitError(msg string) {
	fmt.Println(msg)
	os.Exit(84)
}

// Convert String to Int64
func conversionStringToInt64(value string) int64 {
	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		exitError("Wrong Argument")
	}
	return number
}

// Convert String to Float64
func conversionStringToFloat64(value string) float64 {
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		exitError("Wrong Argument")
	}
	return number
}

// Open file and return content
func getFile(value string) *bufio.Scanner {
	file, err := os.Open(value)
	if err != nil {
		log.Fatal(err)
	}
	res := bufio.NewScanner(file)
	return res
}
