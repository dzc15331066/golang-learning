# README

我使用xorm来改写entities里面的内容，代码数量果然少了

使用curl测试数据服务。

```sh
$curl -d "username=ooo&departname=1" http://localhost:8080/service/userinfo
{
  "UID": 24,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-11-30T11:10:13.797633824+08:00"
}

```

```sh
$curl http://localhost:8080/service/userinfo?userid=24
{
  "UID": 24,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-11-30T08:00:00+08:00"
}

```

可见能够正常插入记录和查询记录了

ab压力测试

改写前

```sh
$ab -n 10000 -c 100 http://localhost:8080/service/userinfo?userid=24
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=24
Document Length:        96 bytes

Concurrency Level:      100
Time taken for tests:   8.167 seconds
Complete requests:      10000
Failed requests:        0
Non-2xx responses:      10000
Total transferred:      2280000 bytes
HTML transferred:       960000 bytes
Requests per second:    1224.51 [#/sec] (mean)
Time per request:       81.665 [ms] (mean)
Time per request:       0.817 [ms] (mean, across all concurrent requests)
Transfer rate:          272.65 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       5
Processing:     1   81  59.7     67     505
Waiting:        1   81  59.7     67     505
Total:          1   81  59.8     68     505

Percentage of the requests served within a certain time (ms)
  50%     68
  66%     92
  75%    111
  80%    124
  90%    163
  95%    198
  98%    240
  99%    274
 100%    505 (longest request)

```

改写后

```sh
ab -n 10000 -c 100 http://localhost:8080/service/userinfo?userid=24
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=24
Document Length:        101 bytes

Concurrency Level:      100
Time taken for tests:   7.891 seconds
Complete requests:      10000
Failed requests:        26
   (Connect: 0, Receive: 0, Length: 26, Exceptions: 0)
Non-2xx responses:      10000
Total transferred:      2339246 bytes
HTML transferred:       1009272 bytes
Requests per second:    1267.28 [#/sec] (mean)
Time per request:       78.909 [ms] (mean)
Time per request:       0.789 [ms] (mean, across all concurrent requests)
Transfer rate:          289.50 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0       6
Processing:     1   78  59.8     64     380
Waiting:        1   78  59.8     64     379
Total:          1   79  59.9     65     381

Percentage of the requests served within a certain time (ms)
  50%     65
  66%     92
  75%    112
  80%    126
  90%    164
  95%    196
  98%    233
  99%    259
 100%    381 (longest request)

```

发现区别不大，估计瓶颈在IO，反射的影响不大。