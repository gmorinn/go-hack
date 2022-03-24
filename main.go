package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"
)

func init() {
	flag.StringVar(&_url, "u", "", "Url of the API")
	flag.Int64Var(&_concurrencyLevel, "t", 3, "Number of concurrent threads (default=3)")
	flag.StringVar(&_wordlistFile, "w", "", "Path to the wordlist")
	flag.StringVar(&_methodStr, "m", "POST", "Method of the request. Default=POST")
	flag.StringVar(&_payload, "p", "", "The attribute you want to modify. Example: -p email -w email.txt")
	flag.Int64Var(&_limit, "l", 0, "Stop the program when the number of request equals ${limit}")
	flag.StringVar(&_generalPayload, "gp", "", "The payload that does not change with each request. Exemple: -gp \"{\\\"email\\\":\\\"guillaume@test.com\\\"}\"")
	flag.StringVar(&_response, "r", "", "Corresponds to the response or message returned by the API. If the response of the request contains the response specified then the program stops. Example: -r \"{\\\"success\\\":true}\"")
	flag.BoolVar(&_not, "n", false, "Stop the program when the request send a response different of the -r specified. Exemple: -r \"INVALD API KEY\" -n => The program stop when the response is different that \"INVALID API KEY\"")
	flag.Parse()
}

type Methods string

const (
	POST Methods = "POST"
	// GET         = "GET"
	// PATCH          = "PATCH"
	// PUT            = "PUT"
	// DELETE         = "DELETE"
)

var (
	_concurrencyLevel int64 = 3
	_limit            int64 = -1
	_wordlistFile     string
	_generalPayload   string
	_url              string
	_payload          string
	_not              bool = false
	_methodStr        string
	_method           Methods = "POST"
	_response         string
)

func manageBrutForce() {
	switch _method {
	case POST:
		brutForcePOST()
		break
	// case GET:
	// 	brutForceGET()
	// 	break
	// case PUT:
	// 	brutForcePUT()
	// 	break
	// case DELETE:
	// 	brutForceDELETE()
	// 	break
	// case PATCH:
	// 	brutForcePATCH()
	// 	break
	default:
		break
	}
}

func manageError() {
	if _url == "" {
		help()
		exitError("Url is empty")
	}
	if _concurrencyLevel <= 0 {
		help()
		exitError("-t should be greater than 0")
	}
	if _wordlistFile == "" {
		help()
		exitError("Wordlist is empty")
	}
	if _limit == -1 && len(_response) == 0 {
		help()
		exitError("Not limit or response specified")
	}
	ping, _ := exec.Command("ping", _url, "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(ping), "Destination Host Unreachable") {
		exitError("Wrong URL")
	}
}

func manageArgs(args []string) {
	setMethod(Methods(_methodStr))
	if _payload != "" {
		setGeneralPayload(_payload)
	}

	if _response != "" {
		setResponse(_response)
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
