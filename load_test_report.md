# Отчет о нагрузочном тестировании. Lab_05.
#### ab -n 25000 -c 10 http://localhost/api/v1/competitions
```
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        nginx/1.27.3
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/competitions
Document Length:        524 bytes

Concurrency Level:      10
Time taken for tests:   7.545 seconds
Complete requests:      25000
Failed requests:        0
Total transferred:      16850000 bytes
HTML transferred:       13100000 bytes
Requests per second:    3313.48 [#/sec] (mean)
Time per request:       3.018 [ms] (mean)
Time per request:       0.302 [ms] (mean, across all concurrent requests)
Transfer rate:          2180.94 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:     0    3   3.7      2      36
Waiting:        0    3   3.7      2      36
Total:          0    3   3.7      2      36

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      3
  80%      3
  90%      9
  95%     13
  98%     15
  99%     17
 100%     36 (longest request)
```
#### ab -n 25000 -c 10 http://localhost/api/v1/no_balance/competitions
```
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        nginx/1.27.3
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/no_balance/competitions
Document Length:        524 bytes

Concurrency Level:      10
Time taken for tests:   13.377 seconds
Complete requests:      25000
Failed requests:        0
Total transferred:      16850000 bytes
HTML transferred:       13100000 bytes
Requests per second:    1868.92 [#/sec] (mean)
Time per request:       5.351 [ms] (mean)
Time per request:       0.535 [ms] (mean, across all concurrent requests)
Transfer rate:          1230.13 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:     0    5   5.7      3      75
Waiting:        0    5   5.7      3      75
Total:          0    5   5.7      3      75

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      9
  80%     10
  90%     12
  95%     14
  98%     19
  99%     27
 100%     75 (longest request)
```
