package infoblox

import (
	"net/http"
	"testing"
)

func init() {
	ts = dummyInfoblox()
	ib = NewClient(ts.URL, "admin", "infoblox", false, false)
}

func TestSendRequest(t *testing.T) {
	resp, err := ib.SendRequest("GET", ts.URL+"/wapi/v2.2/record:a", "", nil)
	if err != nil {
		t.Errorf("Unexpected error : %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected code : %v, got :%v\n", http.StatusOK, resp.StatusCode)
	}

	resp, err = ib.SendRequest("GET", ts.URL+"/wapi/v2.2/record:foo", "", nil)
	if err != nil {
		t.Errorf("Unexpected error : %v\n", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected code : %v, got :%v\n", http.StatusBadRequest, resp.StatusCode)
	}
}
