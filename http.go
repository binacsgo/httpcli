package httpcli

// Post send POST request
func Post(url string, request, response interface{}) error {
	return call(gHTTPClient, url, "POST", request, response)
}

// Get send GET request
func Get(url string, request, response interface{}) error {
	return call(gHTTPClient, url, "GET", request, response)
}

// Put send PUT request
func Put(url string, request, response interface{}) error {
	return call(gHTTPClient, url, "PUT", request, response)
}

// Delete send DELETE request
func Delete(url string, request, response interface{}) error {
	return call(gHTTPClient, url, "DELETE", request, response)
}

