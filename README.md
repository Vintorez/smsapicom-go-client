## smsapi.com go(lang) client

Simple smsapi.com client for sending sms.

### usage

```go
go get github.com/vintorez/smsapicom-go-client
```

example with no optional params:

* "username" - user name
* "password"  - password
* "log" - logger interface (optional)
* "insecure" - use certificate (optional)
* "localCertFile" - path to local certificate (optional)
* "httpClient" - custom http client (optional)

```go
package main

import (
	"github.com/vintorez/smsapicom-go-client"
)

service := smsapicom.NewSmsService("username", "password", "log", "insecure", "localCertFile", "httpClient")
report, err := service.Send("from", "to", "text message", nil)
```
