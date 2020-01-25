

Default authorizaton module used on GoAuth2
==============

[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/default_auth)](https://goreportcard.com/report/github.com/Ulbora/default_auth)


Can interface with any GoAuth2 proxy for any type of authentication service. GoAuth2Users meets all of the requirements below for a GoAuth2 Proxy.

### Any GoAuth2 Proxy service must meet the following requirement:
The GoAuth2 Proxy must implement the code inside the "login" service to interface with the target authentication service.

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


# Usage

```

import (
	px "github.com/Ulbora/GoProxy"
	au "github.com/Ulbora/auth_interface"
	"testing"
)

var authURL = "http://localhost:3001/rs/user/login"

var proxy px.GoProxy

var login au.Login
login.Username = "admin"
login.Password = "admin"
login.ClientID = 10

var da DefaultAuth
da.AuthServerURL = authURL
da.Proxy = proxy.GetNewProxy()
ai := da.GetNew()
val := ai.UserLogin(&login)

```