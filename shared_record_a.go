package infoblox

import "fmt"

// https://102.168.2.200/wapidoc/objects/record.a.html
func (c *Client) SharedRecordA() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "sharedrecord:a",
	}
}

type SharedRecordAObject struct {
	Object
	Comment              string            `json:"comment,omitempty"`
	Ipv4Addr             string            `json:"ipv4addr,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Ttl                  int               `json:"ttl,omitempty"`
	View                 string            `json:"view,omitempty"`
	SharedRecordGroup    string            `json:"shared_record_group,omitempty"`
	Disable              bool              `json:"disable,omitempty"`
	Extattrs             map[string]string `json:"extattrs,omitempty"`
	ExtensibleAttributes map[string]string `json:"extensible_attributes,omitempty"`
}

func (c *Client) SharedRecordAObject(ref string) *SharedRecordAObject {
	a := SharedRecordAObject{}
	a.Object = Object{
		Ref: ref,
		r:   c.SharedRecordA(),
	}
	return &a
}

func (c *Client) GetSharedRecordA(ref string, opts *Options) (*SharedRecordAObject, error) {
	resp, err := c.SharedRecordAObject(ref).get(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not get created A shared record: %s", err)
	}
	var out SharedRecordAObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) FindSharedRecordA(name string) ([]SharedRecordAObject, error) {
	field := "name"
	conditions := []Condition{Condition{Field: &field, Value: name}}
	resp, err := c.SharedRecordA().find(conditions, nil)
	if err != nil {
		return nil, err
	}

	var out []SharedRecordAObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
