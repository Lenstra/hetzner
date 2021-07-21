package hetzner

import (
	"net/url"

	"github.com/mitchellh/mapstructure"
)

func (c *Client) Servers() *ServersClient {
	return &ServersClient{c: c}
}

func (s *ServersClient) List() ([]*ServerSummary, error) {
	d, err := s.c.list("server")
	if err != nil {
		return nil, err
	}

	var res []*ServerSummary
	for _, elem := range d {
		var s ServerSummary
		err = mapstructure.Decode(elem.(map[string]interface{})["server"], &s)
		if err != nil {
			return nil, err
		}
		res = append(res, &s)
	}

	return res, nil
}

func (s *ServersClient) Info(ip string) (*Server, error) {
	d, err := s.c.get("server/" + ip)
	if err != nil {
		return nil, err
	}

	var server Server
	if err := mapstructure.Decode(d["server"], &server); err != nil {
		return nil, err
	}
	return &server, nil
}

func (s *ServersClient) Update(ip, name string) (*Server, error) {
	body := url.Values{}
	body.Add("server_name", name)
	d, err := s.c.post("server/"+ip, body)
	if err != nil {
		return nil, err
	}

	var server Server
	if err := mapstructure.Decode(d["server"], &server); err != nil {
		return nil, err
	}
	return &server, nil
}
