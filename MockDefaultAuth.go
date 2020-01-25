package defaultauth

import (
	px "github.com/Ulbora/GoProxy"
	au "github.com/Ulbora/auth_interface"
)

//MockDefaultAuth MockDefaultAuth
type MockDefaultAuth struct {
	Proxy         px.Proxy
	AuthServerURL string
	MockValid     bool
}

//UserLogin UserLogin
func (m *MockDefaultAuth) UserLogin(login *au.Login) bool {
	return m.MockValid
}

//GetNew GetNew
func (m *MockDefaultAuth) GetNew() au.AuthInterface {
	var rtn au.AuthInterface
	rtn = m
	return rtn
}
