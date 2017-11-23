# 	README

##	开发类似于cloudgo的web服务程序

基于net/http库，实现命令行选择在特定端口启动服务器。

因为只涉及到查询字符串，个人感觉net/http库足够了。

关键代码：

service/server.go

```go
package service

import (
	"fmt"
	"net/http"
)

type MyServer struct {
	mux *http.ServeMux
}

// set the running address
func (sr *MyServer) Run(addr string) {
	http.ListenAndServe(addr, sr)
}

func (sr *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if sr.mux == nil {
		sr.mux = http.DefaultServeMux
	}
	sr.mux.ServeHTTP(w, r)
	return
}

// NewServer configures and returns a Server.
func NewServer() *MyServer {
	sr := &MyServer{}
	sr.mux = http.NewServeMux()
	sr.mux.HandleFunc("/", sayhelloName)
	return sr
}

// define a handle function
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"]
	if len(name) == 0 {
		fmt.Fprintf(w, "Hello\n")
	} else {
		fmt.Fprintf(w, "Hello "+name[0]+"\n")
	}

}

```

main.go

```go
package main

import (
	"os"

	"github.com/dzc15331066/golang-learning/web/server1/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
```



```shell
~$ go run main.go -p8088
listening at port :8088
```

curl 测试

```shell
~$ curl -v localhost:8000/?name=haha
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> GET /?name=haha HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 15 Nov 2017 06:47:22 GMT
< Content-Length: 11
< Content-Type: text/plain; charset=utf-8
< 
Hello haha
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact

```

ab压力测试

```shell
 ~$ ab -n 1000 -c 100 http://localhost:8000/?name=haha
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /?name=haha
Document Length:        11 bytes

Concurrency Level:      100
Time taken for tests:   0.114 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      128000 bytes
HTML transferred:       11000 bytes
Requests per second:    8800.18 [#/sec] (mean)
Time per request:       11.363 [ms] (mean)
Time per request:       0.114 [ms] (mean, across all concurrent requests)
Transfer rate:          1100.02 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    4   1.2      4       8
Processing:     2    7   4.0      6      21
Waiting:        1    5   3.6      4      19
Total:          6   11   3.6     10      23

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     10
  75%     12
  80%     13
  90%     17
  95%     19
  98%     21
  99%     22
 100%     23 (longest request)

```

一些参数的解释

Server Software
表示被测试的Web服务器软件名称，这里是Apache/2.2.19,它来自于http响应数据的头信息，所以如果是我们自己编写的Web服务器软或者修改开源Web服务器软件的源代码，便可以随意改写这里的名称。
Server Hostname

表示请求的URL中的主机部分名称

Server Port
表示被测试的Web服务器软件的监听端口

Document Path
表示请求的URL中根绝对路径

Document Length

表示http响应数据的正文长度。

Concurrency Level
表示并发用户数，通过-c设置的参数。

Time taken for tests
表示所有这些请求被处理完成花费的总时间。

Complete requests

表示总请求数，这是通过-n设置的相应参数。

Failed requests
表示失败的请求数，这里的失败是指请求的连接服务器、发送数据、接收数据等环节发生异常，以及无响应后超时的情况。对于超时时间的设置可以用ab的-t参数。
而如果接受到的http响应数据的头信息中含有2xx以外的状态码，则会在测试结果显示另一个名为“Non-2xx responses”的统计项，用于统计这部分请求数，这些请求并不算是失败的请求。

Total transferred
表示所有请求的响应数据长度总和，包括每个http响应数据的头信息和正文数据的长度。注意这里不包括http请求数据的长度，所以Total transferred代表了从Web服务器流向用户PC的应用层数据总长度。通过使用ab的-v参数即可查看详细的http头信息。

HTML transferred
表示所有请求的响应数据中正文数据的总和，也就是减去了Total transferred中http响应数据中头信息的长度。

**Requests per second**
这便是吞吐率，它等于：
Complete requests / Time taken for tests

**Time per request**
这便是用户平均请求等待时间，它等于:
Time taken for tests / (Complete requests /Concurrency Level)

**Time per request?(across all concurrent requests)**
这是服务器平均请求处理时间，等于：
Time taken for tests / Complete requests

Transfer rate
表示这些请求在单位时间内从服务器获取的数据长度，它等于：
Total transferred / Time taken for tests
这个统计项可以很好的说明服务器在处理能力达到限制时，其出口带宽的需求量。
利用前面介绍的有关带宽的知识，不难计算出结果。

**Percentage of the requests served within a certain time(ms)**
这部分数据用于描述每个请求处理时间的分布情况。