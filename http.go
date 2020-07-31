package httpcli

// Post send POST request
func Post(url string, request interface{}) ([]byte, error) {
	return post(gHTTPClient, url, request)
}

// Get send GET request
func Get(addr string) ([]byte, error) {
	return get(gHTTPClient, addr)
}
