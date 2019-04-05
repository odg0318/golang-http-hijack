# golang-http-hijack
This repository shows how to use hijack in `net/http`.

## Dumping headers and body
```console
$ curl -v http://127.0.0.1:8000/hijack
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8000 (#0)
> GET /hijack HTTP/1.1
> Host: 127.0.0.1:8000
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=utf-8
< Content-Length: 5
<
* Connection #0 to host 127.0.0.1 left intact
hello
```

## Testing time to be spent
```console
$ curl -LO http://127.0.0.1:8000/hijack
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     5  100     5    0     0      1      0  0:00:05  0:00:04  0:00:01     1
```
