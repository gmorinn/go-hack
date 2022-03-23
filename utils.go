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
	fmt.Println("\n\t-u, --url\t\t\t{URL_API} Url of the API")
	fmt.Println("\t-t, --threads\t\t\t{THREADS} Number of concurrent threads (default=3)")
	fmt.Println("\t-w, --wordlist\t\t\t{WORDLIST} Path to the wordlist")
	fmt.Println("\t-m, --method\t\t\t{METHOD} Method of the request. Default=POST")
	fmt.Println("\t-p, --payload\t\t\t{PAYLOAD} The attribute you want to modify. Example: -p email -w email.txt")
	fmt.Println("\t-l, --limit\t\t\t{LIMIT} Stop the program when the number of request equals ${limit}")
	fmt.Println("\t-gp, --general-payload\t\t{GENERAL_PAYLOAD} The payload that does not change with each request. Exemple: -gp \"{\"email\":\"guillaume@test.com\"}\"")
	fmt.Println("\t-r, --response\t\t\t{RESPONSE} Corresponds to the response or message returned by the API. If the response of the request contains the response specified then the program stops. Example: -r \"{\"success\":\"true\"}\"")
	fmt.Println("\t-n, --not\t\t\t{NOT} Stop the program when the request send a response different of the -r specified. Exemple: -r \"INVALD API KEY\" -n => The program stop when the response is different that \"INVALID API KEY\"")

	fmt.Println("\nNOTICE !")
	fmt.Println("\t- Only one url is taken into account.")
	fmt.Println("\t- Only one attribute is possible in the payload for version 1.")
	fmt.Println("\t- All flags are not mandatory, except -u && -w && (-r || -l)")
	fmt.Println("\t- You can stop the program at any moment with CRTL + C")

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
