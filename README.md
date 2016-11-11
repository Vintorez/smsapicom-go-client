## smsapi.com go(lang) client

Simple smsapi.com client for sending sms.

example with default http-client and no options params:

`service := smsapicom.NewSmsService("user_name", "user_password", nil)`
`report, err := service.Send("from", "to", "text message", nil)`
