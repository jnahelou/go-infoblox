package infoblox

import (
	"testing"
)

var DummyRecordAObjects map[string]RecordAObject

func buildRecordAObject(o map[string]interface{}) RecordAObject {
	obj := RecordAObject{}
	if v, ok := o["_ref"]; ok {
		obj.Object = Object{
			Ref: v.(string),
		}
	}
	if v, ok := o["name"]; ok {
		obj.Name = v.(string)
	}
	if v, ok := o["ipv4addr"]; ok {
		obj.Ipv4Addr = v.(string)
	}
	if v, ok := o["comment"]; ok {
		obj.Comment = v.(string)
	}
	if v, ok := o["ttl"]; ok {
		obj.Ttl = v.(int)
	}
	if v, ok := o["view"]; ok {
		obj.View = v.(string)
	}

	return obj
}

func TestGetRecordA(t *testing.T) {
	//Refer to basic fixture
	ref := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQubG9jYWxka3QudGVzdCxmb28tam51LDEwLjAuMC41:foo1.test.local/default"
	obj, err := ib.GetRecordA(ref, nil)
	if err != nil {
		t.Errorf("Unknown error : %v", err)
	}
	if obj.Object.Ref == "" {
		t.Errorf("Unable to find record:a : %v\n", ref)
	}

	//Query non-existing object
	ref = "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQubG9jYWxka3QudGVzdCxmb28tam51LDEwLjAuMC41:not.present.dns/default"
	obj, err = ib.GetRecordA(ref, nil)
	if err != nil {
		t.Errorf("Unknown error : %v", err)
	}
	if obj.Object.Ref != "" {
		t.Errorf("Find non-existing object  : %v\n", obj)
	}

}
