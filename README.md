# httpcli
Simple http client written in Go.



## 0. Quick Start

```go
bs, err := httpcli.Get(url)
if err != nil {
	log.Printf("err = %+v\n", err)
	// return
}
log.Printf("response:\n%s\n", string(bs[:]))
```



## 1. Usage

```go
bs, err := httpcli.Get(url)
...

bs, err := httpcli.Post(url, request.interface{})
...

bs, err := httpcli.GetSkipVerify(url)
...

bs, err := httpcli.PostSkipVerify(url, request.interface{})
...
```



## 2. About

自己写接口之后需要测试，这个时候的客户端不需要太多的定制，只是和服务端通过 `json` 交互而已。

因此独立出一个 `httpcli` ，放置简单封装的客户端 `HTTP Client` ，减少不同项目测试时要写的重复代码。

