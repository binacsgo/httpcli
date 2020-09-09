# httpcli
Simple http client written in Go.


> 为 JSON 交互定制

## 0. Quick Start

```go
var resp resp_struct{}
err := httpcli.Get(url, &resp)
if err != nil {
	log.Printf("err = %+v\n", err)
	// return
}
log.Printf("response:\n%+v\n", resp)
```



## 1. Usage

```go
err := httpcli.Get(url, request, &response)
...

err := httpcli.Post(url, request, &response)
...

err := httpcli.Put(url, request, &response)
...

err := httpcli.Delete(url, request, &response)
...

err := httpcli.GetSkipVerify(url, request, &response)
...

err := httpcli.PostSkipVerify(url, request, &response)
...

err := httpcli.PutSkipVerify(url, request, &response)
...

err := httpcli.DeleteSkipVerify(url, request, &response)
...
```



## 2. About

自己写接口之后需要测试，这个时候的客户端不需要太多的定制，只是和服务端通过 `json` 交互而已。

因此独立出一个 `httpcli` ，放置简单封装的客户端 `HTTP Client` ，减少不同项目测试时要写的重复代码。

