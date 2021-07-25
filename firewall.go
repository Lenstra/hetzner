package hetzner

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
	var d map[string]*Firewall
	if err := f.c.post("firewall/"+req.ServerIP, req, &d); err != nil {
		return nil, err
	}
	return d["firewall"], nil
}

func (f *FirewallClient) Delete(ip string) (*Firewall, error) {
	var d map[string]*Firewall
	if err := f.c.delete("firewall/"+ip, &d); err != nil {
		return nil, err
	}
	return d["firewall"], nil
}
