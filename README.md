# http-debug

Ultra small webserver that prints and logs your requests.

```bash
$ docker run -d -p 8080:8080 nousefreak/http-debug
```

```bash
$ curl 127.0.0.1:8080/sdf -H "Host:somehost.com" -H "X-Custom-Header: bla"
GET /sdf HTTP/1.1
Host: somehost.com
Accept: */*
User-Agent: curl/7.54.0
X-Custom-Header: bla
```
