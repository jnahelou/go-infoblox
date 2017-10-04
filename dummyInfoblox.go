package infoblox

import (
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
)

var supportedWapiObjects = map[string]bool{
	"record:a":       true,
	"record:aaaa":    true,
	"record:cname":   true,
	"record:host":    true,
	"record:ptr":     true,
	"scheduledtask":  true,
	"ipv4address":    true,
	"network":        true,
	"sharedrecord:a": true,
}

var ib *Client
var ts *httptest.Server

var UnAuthorizedResponse = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<title>401 Authorization Required</title>
</head><body>
<h1>Authorization Required</h1>
<p>This server could not verify that you
are authorized to access the document
requested.  Either you supplied the wrong
credentials (e.g., bad password), or your
browser doesn't understand how to supply
the credentials required.</p>
</body></html>`

func BadRequestResponse(wapiObject string) string {
	return `{ "Error": "AdmConProtoError: Unknown object type (` + wapiObject + `)",
  "code": "Client.Ibap.Proto",
  "text": "Unknown object type (` + wapiObject + `)"
}`
}

func checkCred(r *http.Request) bool {
	user, pass, ok := r.BasicAuth()
	return ok && user == "admin" && pass == "infoblox"
}

func dummyInfobloxGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !checkCred(r) {
		http.Error(w, UnAuthorizedResponse, http.StatusUnauthorized)
		return
	}
	if _, ok := supportedWapiObjects[ps.ByName("wapiObject")]; !ok {
		http.Error(w, BadRequestResponse(ps.ByName("wapiObject")), http.StatusBadRequest)
		return
	}
	return
}

func dummyInfobloxPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !checkCred(r) {
		http.Error(w, UnAuthorizedResponse, http.StatusUnauthorized)
		return
	}
	if _, ok := supportedWapiObjects[ps.ByName("wapiObject")]; !ok {
		http.Error(w, BadRequestResponse(ps.ByName("wapiObject")), http.StatusBadRequest)
		return
	}
	return
}
func dummyInfobloxDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !checkCred(r) {
		http.Error(w, UnAuthorizedResponse, http.StatusUnauthorized)
		return
	}
	if _, ok := supportedWapiObjects[ps.ByName("wapiObject")]; !ok {
		http.Error(w, BadRequestResponse(ps.ByName("wapiObject")), http.StatusBadRequest)
		return
	}
	return
}

func dummyInfoblox() *httptest.Server {
	router := httprouter.New()
	router.GET("/wapi/v2.2/:wapiObject", dummyInfobloxGet)
	router.POST("/wapi/v2.2/:wapiObject", dummyInfobloxPost)
	router.DELETE("/wapi/v2.2/:wapiObject", dummyInfobloxDelete)
	return httptest.NewTLSServer(router)

}
