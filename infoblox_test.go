package infoblox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type Fixtures []Resources

type Resources struct {
	WapiType    string                   `json:"type"`
	WapiObjects []map[string]interface{} `json:"values"`
}

func init() {
	ts = dummyInfoblox()
	ib = NewClient(ts.URL, "admin", "infoblox", false, false)

	//Populate basic informations

	//fixtures :=
	//	[]Resources{
	//		{
	//			WapiType: "record:a",
	//			WapiObjects: []map[string]interface{}{
	//				map[string]interface{}{
	//					"Name":     "foo",
	//					"Ipv4addr": "10.0.0.1/32",
	//				},
	//			},
	//		},
	//	}
	//fmt.Print(fixtures)
	text := `
[
	{
		"type": "record:a",
		"values": [
			    {
			        "_ref": "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQubG9jYWxka3QudGVzdCxmb28tam51LDEwLjAuMC41:foo1.test.local/default", 
			        "ipv4addr": "10.0.0.1", 
			        "name": "foo1.test.local", 
			        "view": "default"
			    }, 
			    {
			        "_ref": "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQubG9jYWxka3QudGVzdCxmb28tam51MiwxMC4wLjAuNg:foo2.test.local/default", 
			        "ipv4addr": "10.0.0.2", 
			        "name": "foo2.test.local", 
			        "view": "default"
			    }
		]
	}, {
		"type": "record:ptr",
		"values": [
			    {
			        "_ref": "record:ptr/ZG5zLmJpbmRfcHRyJC5fZGVmYXVsdC5hcnBhLmluLWFkZHIuMTI3LjAuMC4xLmxvY2FsaG9zdA:1.0.0.127.in-addr.arpa/default", 
			        "ptrdname": "localhost", 
			        "view": "default"
			    }, 
			    {
			        "_ref": "record:ptr/ZG5zLmJpbmRfcHRyJC5fZGVmYXVsdC5hcnBhLmlwNi4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjAuMC4wLjEuLmxvY2FsaG9zdA:1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa/default", 
			        "ptrdname": "localhost", 
			        "view": "default"
			    }
		]

	}
]
`
	var f Fixtures
	err := json.Unmarshal([]byte(text), &f)
	if err != nil {
		fmt.Print(err)
	}
	for _, thing := range f {
		for _, wapiObjet := range thing.WapiObjects {
			//TODO Can be replace by HTTP POST call or Object creation
			dummySaveObjectByType(thing.WapiType, wapiObjet)
		}
	}
	// TODO Create a dummy DNS zone
	// TODO Create a dummy Shared Group
	// TODO Create a dummy DNS Record Zone
	fmt.Printf("=== End init ===\n")
}

func TestSendRequest(t *testing.T) {
	//Get all pre-load record:a
	resp, err := ib.SendRequest("GET", ts.URL+BASE_PATH+"record:a", "", nil)
	if err != nil {
		t.Errorf("Unexpected error : %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected code : %v, got :%v\n", http.StatusOK, resp.StatusCode)
	}
	//Get content and check there are only 2 elements
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var res []RecordAObject
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		t.Errorf("Unable to convert result to array of RecordAObjects\n")
	}
	if len(res) != len(DummyRecordAObjects) {
		t.Errorf("There is %s objetcs in memory but only %s in response", len(DummyRecordAObjects), len(res))
	}
	//Call unknown wapiObject
	resp, err = ib.SendRequest("GET", ts.URL+BASE_PATH+"record:foo", "", nil)
	if err != nil {
		t.Errorf("Unexpected error : %v\n", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected code : %v, got :%v\n", http.StatusBadRequest, resp.StatusCode)
	}
}
