package main

import (
	"os"
	"strings"
)

type Methods string

const (
	GET    Methods = "GET"
	POST           = "POST"
	PATCH          = "PATCH"
	PUT            = "PUT"
	DELETE         = "DELETE"
)

var (
	_concurrencyLevel int64 = 3
	_limit            int64 = -1
	_wordlistFile     string
	_url              string
	_payload          string
	_method           Methods
	_response         []string
)

func manageBrutForce() {
	switch _method {
	case GET:
		brutForceGET()
		break
	case POST:
		brutForcePOST()
		break
	case PUT:
		brutForcePUT()
		break
	case DELETE:
		brutForceDELETE()
		break
	case PATCH:
		brutForceGET()
		break
	}
}

func manageError() {
	if _url == "" {
		exitError("Url is empty")
	}
	if _concurrencyLevel <= 0 {
		exitError("-t should be greater than 0")
	}
	if _wordlistFile == "" {
		exitError("Wordlist is empty")
	}
	if _limit == -1 && len(_response) == 0 {
		exitError("Not limit or response specified")
	}
}

func manageArgs(args []string) {
	for i, v := range args {
		if (v == "-t" || v == "--threads") && (i+1) < len(args) {
			setThread(args[i+1])
		}
		if (v == "-w" || v == "--wordlist") && (i+1) < len(args) {
			setWordlist(args[i+1])
		}
		if (v == "-u" || v == "--url") && (i+1) < len(args) {
			setUrl(args[i+1])
		}
		if (v == "-m" || v == "--method") && (i+1) < len(args) {
			setMethod(Methods(args[i+1]))
		}
		if (v == "-p" || v == "--payload") && (i+1) < len(args) {
			setPayload(args[i+1])
		}
		if (v == "-r" || v == "--response") && (i+1) < len(args) {
			setResponse(strings.Split(args[i+1], ","))
		}
		if (v == "-l" || v == "--limit") && (i+1) < len(args) {
			setLimit(args[i+1])
		}
	}
	manageError()
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		help()
	} else {
		manageArgs(os.Args)
		manageBrutForce()
	}
	os.Exit(0)
}
