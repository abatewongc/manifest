```
go build -buildmode=exe -ldflags="-s -w -H windowsgui -linkmode external -extldflags -static" .
```