package main

// Set thread
func setThread(concurrencyLevel string) {
	_concurrencyLevel = conversionStringToInt64(concurrencyLevel)
}

// Set wordlist
func setWordlist(wordlistFile string) {
	_wordlistFile = wordlistFile
}

// Set url api
func setUrl(url string) {
	_url = url
}

// Set method
func setMethod(method Methods) {
	_method = method
}

// Set payload
func setPayload(payload string) {
	_payload = payload
}

// Set response
func setResponse(response string) {
	if response[:1] == "{" && response[:len(response)-1] == "}" && len(response) > 3 {
		_response = response[1 : len(response)-2]
	} else {
		_response = response
	}
}

// Set response
func setGeneralPayload(general_payload string) {
	if general_payload[:1] == "{" && general_payload[:len(general_payload)-1] == "}" && len(general_payload) > 3 {
		_generalPayload = general_payload[1 : len(general_payload)-2]
	} else {
		exitError("General payload needs to be an object")
	}
}

// Set limit
func setLimit(limit string) {
	_limit = conversionStringToInt64(limit)
}

// Set not
func setNot(not bool) {
	_not = not
}
