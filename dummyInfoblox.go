package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

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

//Messages from my free trial appliance
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

func BadReferenceResponse(ref string) string {
	return `{ "Error": "AdmConDataNotFoundError: Reference ` + ref + ` not found", 
  "code": "Client.Ibap.Data.NotFound", 
  "text": "Reference ` + ref + ` not found"
}`
}

// Basic credential validation
func checkCred(r *http.Request) bool {
	user, pass, ok := r.BasicAuth()
	return ok && user == "admin" && pass == "infoblox"
}

// Store objects in-memory
func dummySaveObjectByType(t string, o map[string]interface{}) error {
	switch t {
	case "record:a":
		recordA := buildRecordAObject(o)
		if v, ok := o["_ref"]; ok {
			DummyRecordAObjects[v.(string)] = recordA
		}
	case "record:aaaa":
		fmt.Printf("record:aaaa")
	case "record:cname":
		fmt.Printf("record:cname")
	case "record:host":
		fmt.Printf("record:host")
	case "record:ptr":
		recordPtr := buildRecordPtrObject(o)
		if v, ok := o["_ref"]; ok {
			DummyRecordPtrObjects[v.(string)] = recordPtr
		}
	case "scheduledtask":
		fmt.Printf("scheduledtask")
	case "ipv4address":
		fmt.Printf("ipv4address")
	case "network":
		fmt.Printf("network")
	case "sharedrecord:a":
		fmt.Printf("sharedrecord:a")
	}
	return nil
}

// Get Handler
func dummyInfobloxGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !checkCred(r) {
		http.Error(w, UnAuthorizedResponse, http.StatusUnauthorized)
		return
	}
	if _, ok := supportedWapiObjects[ps.ByName("wapiObject")]; !ok {
		http.Error(w, BadRequestResponse(ps.ByName("wapiObject")), http.StatusBadRequest)
		return
	}
	wapiObject := ps.ByName("wapiObject")
	switch wapiObject {
	case "record:a":
		query := ps.ByName("query")
		if query != "" { //return object
			//TODO : need to implement filters..
			ref := strings.TrimPrefix(r.URL.Path, BASE_PATH)
			if val, ok := DummyRecordAObjects[ref]; ok {
				resp, err := json.Marshal(val)
				if err != nil {
					http.Error(w, "Dummy Internal Error", http.StatusInternalServerError)
				}
				fmt.Fprintf(w, string(resp))
			} else {
				//Key not found, return infoblox error message
				fmt.Fprintf(w, BadReferenceResponse(ref))
			}
		} else { //return all objects
			var recordAObjectsList []RecordAObject
			for _, v := range DummyRecordAObjects {
				recordAObjectsList = append(recordAObjectsList, v)
			}
			resp, err := json.Marshal(recordAObjectsList)
			if err != nil {
				http.Error(w, "Dummy Internal Error", http.StatusInternalServerError)
			}
			fmt.Fprintf(w, string(resp))
		}
		return
	case "record:aaaa":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "record:cname":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "record:host":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "record:ptr":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "scheduledtask":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "ipv4address":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "network":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	case "sharedrecord:a":
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}
	return
}

// Post Handler
func dummyInfobloxPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !checkCred(r) {
		http.Error(w, UnAuthorizedResponse, http.StatusUnauthorized)
		return
	}
	if _, ok := supportedWapiObjects[ps.ByName("wapiObject")]; !ok {
		http.Error(w, BadRequestResponse(ps.ByName("wapiObject")), http.StatusBadRequest)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
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

	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	return
}

func dummyInfoblox() *httptest.Server {
	//Init storage memory
	DummyRecordAObjects = make(map[string]RecordAObject)
	DummyRecordPtrObjects = make(map[string]RecordPtrObject)

	//Start http server
	router := httprouter.New()
	router.GET("/wapi/v2.2/:wapiObject", dummyInfobloxGet)
	router.GET("/wapi/v2.2/:wapiObject/*query", dummyInfobloxGet)
	router.POST("/wapi/v2.2/:wapiObject", dummyInfobloxPost)
	router.DELETE("/wapi/v2.2/:wapiObject", dummyInfobloxDelete)
	return httptest.NewTLSServer(router)
}
