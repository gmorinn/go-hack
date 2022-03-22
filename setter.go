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
func setResponse(response []string) {
	_response = response
}

// Set limit
func setLimit(limit string) {
	_limit = conversionStringToInt64(limit)
}
