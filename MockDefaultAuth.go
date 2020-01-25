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
