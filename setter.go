package main

// Set method
func setMethod(method Methods) {
	_method = method
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
	if general_payload[:1] == "{" && general_payload[len(general_payload)-1:len(general_payload)] == "}" && len(general_payload) > 3 {
		_generalPayload = general_payload[1 : len(general_payload)-1]
	} else {
		exitError("General payload needs to be an object")
	}
}
