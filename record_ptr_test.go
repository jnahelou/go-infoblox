package infoblox

var DummyRecordPtrObjects map[string]RecordPtrObject

func buildRecordPtrObject(o map[string]interface{}) RecordPtrObject {
	obj := RecordPtrObject{}
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
	if v, ok := o["ipv6addr"]; ok {
		obj.Ipv6Addr = v.(string)
	}
	if v, ok := o["ptrdname"]; ok {
		obj.PtrDname = v.(string)
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
