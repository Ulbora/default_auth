package defaultauth

import (
	"testing"

	px "github.com/Ulbora/GoProxy"
	au "github.com/Ulbora/auth_interface"
)

func TestMockDefaultAuth_UserLogin(t *testing.T) {
	var authURL = "http://localhost:3001/rs/user/login"

	var proxy px.GoProxy

	var login au.Login
	login.Username = "admin"
	login.Password = "admin"
	login.ClientID = 10

	var da MockDefaultAuth
	da.MockValid = true
	da.AuthServerURL = authURL
	da.Proxy = proxy.GetNewProxy()
	var ai au.AuthInterface
	ai = &da
	val := ai.UserLogin(&login)
	if !val {
		t.Fail()
	}
}
