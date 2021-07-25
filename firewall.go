package hetzner

import (
	"encoding/json"

	"github.com/go-playground/form"
	"github.com/mitchellh/mapstructure"
)

type FirewallClient struct {
	c *Client
}

func (c *Client) Firewall() *FirewallClient {
	return &FirewallClient{c: c}
}

func (f *FirewallClient) Info(ip string) (*Firewall, error) {
	var d map[string]*Firewall
	if err := f.c.get("firewall/"+ip, &d); err != nil {
		return nil, err
	}
	return d["firewall"], nil
}

func (f *FirewallClient) Update(req *FirewallRequest) (*Firewall, error) {
	encoder := form.NewEncoder()
	body, err := encoder.Encode(req)
	if err != nil {
		return nil, err
	}

	content, err := f.c.do("POST", "firewall/"+req.ServerIP, body)
	if err != nil {
		return nil, err
	}

	var d map[string]interface{}
	if err := json.Unmarshal(content, &d); err != nil {
		return nil, err
	}

	d = d["firewall"].(map[string]interface{})

	if _, ok := d["rules"].([]interface{}); ok && len(d) == 0 {
		d["rules"] = map[string]interface{}{}
	}

	var res Firewall
	if err := mapstructure.Decode(d, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (f *FirewallClient) Delete(ip string) (*Firewall, error) {
	var d map[string]*Firewall
	if err := f.c.delete("firewall/"+ip, &d); err != nil {
		return nil, err
	}
	return d["firewall"], nil
}
