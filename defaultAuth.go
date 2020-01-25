package defaultauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	px "github.com/Ulbora/GoProxy"
	au "github.com/Ulbora/auth_interface"
)

//DefaultAuth DefaultAuth
type DefaultAuth struct {
	Proxy         px.Proxy
	AuthServerURL string
}

//LoginRes LoginRes
type LoginRes struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

//UserLogin UserLogin
func (m *DefaultAuth) UserLogin(login *au.Login) bool {
	var rtn bool
	aJSON, _ := json.Marshal(login)
	req, rErr := http.NewRequest("POST", m.AuthServerURL, bytes.NewBuffer(aJSON))
	if rErr != nil {
		log.Println("Request error: ", rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		var lres LoginRes
		fmt.Println("m.Proxy: ", m.Proxy)
		suc, code := m.Proxy.Do(req, &lres)
		// fmt.Println("suc: ", suc)
		// fmt.Println("code: ", code)
		// fmt.Println("lres: ", lres)
		// fmt.Println("lres code: ", lres.Code)
		// fmt.Println("clientid: ", strconv.FormatInt(login.ClientID, 10))
		if suc && code == http.StatusOK && strconv.FormatInt(login.ClientID, 10) == lres.Code {
			rtn = lres.Valid
		}
	}
	return rtn
}

//GetNew GetNew
func (m *DefaultAuth) GetNew() au.AuthInterface {
	var rtn au.AuthInterface
	rtn = m
	return rtn
}

//GO111MODULE=on go mod init github.com/Ulbora/default_auth
