

Default authorizaton module used on GoAuth2
==============

[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/default_auth)](https://goreportcard.com/report/github.com/Ulbora/default_auth)


Proxy for any type of authentication service. GoAuth2Users meets all of the requirements below.

### Any Proxy service must meet the following requirement:

### Validate User
```
Method: POST

URL: http://proxyURL/rs/user/login
```


Request headers Example:
```
Content-Type = application/json
```
Request Body Example:
```
{
   "username":"admin",
   "password":"admin",
   "clientId":10   
}
```
Response:
```
{
    "valid": true,
    "code": "10"
}
```