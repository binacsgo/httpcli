package httpcli

// PostSkipVerify send POST request
func PostSkipVerify(url string, request, response interface{}) error {
	return call(gHTTPClientSkipVerify, url, "POST", request, response)
}

// GetiSkipVerify send GET request
func GetSkipVerify(url string, request, response interface{}) error {
	return call(gHTTPClientSkipVerify, url, "GET", request, response)
}

// PutSkipVerify send PUT request
func PutSkipVerify(url string, request, response interface{}) error {
	return call(gHTTPClientSkipVerify, url, "PUT", request, response)
}

// DeleteSkipVerify send DELETE request
func DeleteSkipVerify(url string, request, response interface{}) error {
	return call(gHTTPClientSkipVerify, url, "DELETE", request, response)
}
