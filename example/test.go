package main

import (
	"log"

	"github.com/binacsgo/httpcli"
)

var (
	url     = "http://baidu.com"
	urlskip = "https://baidu.com"
)

func main() {
	log.Println("test start!")

	log.Printf("\n\n ============ GET ============\n")
	_ = httpcli.Get(url, nil, nil)

	log.Printf("\n\n ============ Post ============\n")
	_ = httpcli.PostRaw(url, nil, nil)

	log.Printf("\n\n ============ GetSkipVerify ============\n")
	_ = httpcli.GetSkipVerify(urlskip, nil, nil)

	log.Printf("\n\n ============ PostSkipVerify ============\n")
	_ = httpcli.PostSkipVerify(urlskip, nil, nil)

	log.Println("test end!")
}
