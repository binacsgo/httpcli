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
	bs, err := httpcli.Get(url)
	checkAndPrint(bs, err)

	log.Printf("\n\n ============ Post ============\n")
	bs, err = httpcli.Post(url, nil)
	checkAndPrint(bs, err)

	log.Printf("\n\n ============ GetSkipVerify ============\n")
	bs, err = httpcli.GetSkipVerify(urlskip)
	checkAndPrint(bs, err)

	log.Printf("\n\n ============ PostSkipVerify ============\n")
	bs, err = httpcli.PostSkipVerify(urlskip, nil)
	checkAndPrint(bs, err)

	log.Println("test end!")
}

func checkAndPrint(bs []byte, err error) {
	if err != nil {
		log.Printf("err = %+v\n", err)
		//return
	}
	log.Printf("response:\n%s\n", string(bs[:]))
}
