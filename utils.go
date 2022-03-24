package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// HELP ./gm -h
func help() {
	fmt.Println(`

	 ██████╗ ███╗   ███╗    ██╗  ██╗ █████╗  ██████╗██╗  ██╗
	██╔════╝ ████╗ ████║    ██║  ██║██╔══██╗██╔════╝██║ ██╔╝
	██║  ███╗██╔████╔██║    ███████║███████║██║     █████╔╝ 
	██║   ██║██║╚██╔╝██║    ██╔══██║██╔══██║██║     ██╔═██╗ 
	╚██████╔╝██║ ╚═╝ ██║    ██║  ██║██║  ██║╚██████╗██║  ██╗
	 ╚═════╝ ╚═╝     ╚═╝    ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
								(v1)	
	`)
	fmt.Println("\t./gm {...flags}")

	flag.Usage()

	fmt.Println("\nNOTICE !")
	fmt.Println("\t- Only one url is taken into account.")
	fmt.Println("\t- Only one attribute is possible in the payload for version 1.")
	fmt.Println("\t- All flags are not mandatory, except -u && -w && (-r || -l)")
	fmt.Println("\t- You can stop the program at any moment with CRTL + C")

}

// Return message error + exit 84
func exitError(msg string) {
	fmt.Println("\n-->\t[ERROR]", msg)
	os.Exit(0)
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
