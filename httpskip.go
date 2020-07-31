package httpcli

// PostSkipVerify send POST request
func PostSkipVerify(url string, request interface{}) ([]byte, error) {
	return post(gHTTPClientSkipVerify, url, request)
}

// GetSkipVerify send GET request
func GetSkipVerify(addr string) ([]byte, error) {
	return get(gHTTPClientSkipVerify, addr)
}
