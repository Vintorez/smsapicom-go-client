## smsapi.com go(lang) client

Simple smsapi.com client for sending sms.

### usage

```go
go get github.com/vintorez/smsapicom-go-client
```

example with default http-client and no options params:

```go
package main

import (
	"github.com/vintorez/smsapicom-go-client"
)

service := smsapicom.NewSmsService("user_name", "user_password", nil)
report, err := service.Send("from", "to", "text message", nil)
```
