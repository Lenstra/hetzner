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
	content, err := f.c.do("GET", "firewall/"+ip, nil)
	if err != nil {
		return nil, err
	}
	return decodeFirewall(content)
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

	return decodeFirewall(content)
}

func (f *FirewallClient) Delete(ip string) (*Firewall, error) {
	content, err := f.c.do("DELETE", "firewall/"+ip, nil)
	if err != nil {
		return nil, err
	}
	return decodeFirewall(content)
}

func decodeFirewall(content []byte) (*Firewall, error) {
	var d map[string]interface{}
	if err := json.Unmarshal(content, &d); err != nil {
		return nil, err
	}

	d = d["firewall"].(map[string]interface{})

	if rules, ok := d["rules"].([]interface{}); ok && len(rules) == 0 {
		d["rules"] = map[string]interface{}{}
	}

	var res Firewall
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &res,
		TagName:  "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(d); err != nil {
		return nil, err
	}

	return &res, nil
}
