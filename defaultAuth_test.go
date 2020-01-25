package defaultauth

import (
	px "github.com/Ulbora/GoProxy"
	au "github.com/Ulbora/auth_interface"
	"testing"
)

func TestDefaultAuth_UserLogin(t *testing.T) {
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
	if !val {
		t.Fail()
	}
}

func TestDefaultAuth_UserLoginBadReq(t *testing.T) {
	var authURL = "://localhost:3001/rs/user/login"

	var proxy px.GoProxy

	var login au.Login
	login.Username = "admin"
	login.Password = "admin"
	login.ClientID = 10

	var da DefaultAuth
	da.AuthServerURL = authURL
	da.Proxy = proxy.GetNewProxy()
	val := da.UserLogin(&login)
	if val {
		t.Fail()
	}
}
