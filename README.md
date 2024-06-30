# Go TCP primal number testing server

https://protohackers.com/problem/1

### run the tcp server
```
go build
./prime_time
```

### telnet to the server
```
telnet localhost 8080
```

## protocol

### input
```
{ "method": "isPrime", "number": 123}
```

### output
```
{ "method": "isPrime", "prime": false}
```