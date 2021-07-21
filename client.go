package hetzner

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	User     string
	Password string
}

type Client struct {
	user, password string
	client         *http.Client
}

func NewClient(config *Config) *Client {
	c := &Client{
		client: cleanhttp.DefaultClient(),
	}

	if config != nil {
		c.user = config.User
		c.password = config.Password
	}

	return c
}

func (c *Client) do(method, url string, body url.Values) ([]byte, error) {
	url = "https://robot-ws.your-server.de/" + url

	var reader io.Reader
	if body != nil {
		reader = strings.NewReader(body.Encode())
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.SetBasicAuth(c.user, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", resp.Status, content)
	}
	return content, nil
}

func (c *Client) list(url string) ([]interface{}, error) {
	content, err := c.do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var d []interface{}
	if err := json.Unmarshal(content, &d); err != nil {
		return nil, err
	}

	return d, nil
}

func (c *Client) get(url string) (map[string]interface{}, error) {
	content, err := c.do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var d map[string]interface{}
	if err := json.Unmarshal(content, &d); err != nil {
		return nil, err
	}

	return d, nil
}

func (c *Client) post(url string, body url.Values) (map[string]interface{}, error) {
	content, err := c.do("POST", url, body)
	if err != nil {
		return nil, err
	}

	var d map[string]interface{}
	if err := json.Unmarshal(content, &d); err != nil {
		return nil, err
	}

	return d, nil
}

func (c *Client) Reset(IP, ty string) (*Reset, error) {
	data := url.Values{}
	data.Set("type", ty)
	d, err := c.post("reset/"+IP, data)
	if err != nil {
		return nil, err
	}

	var r Reset
	if err := mapstructure.Decode(d["reset"], &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) Rescue(IP, os string, arch int, authorizedKeys []string) (*Rescue, error) {
	data := url.Values{}
	data.Set("os", os)
	data.Set("arch", strconv.Itoa(arch))
	for _, key := range authorizedKeys {
		data.Add("authorized_key", key)
	}
	d, err := c.post(fmt.Sprintf("boot/%s/rescue", IP), data)
	if err != nil {
		return nil, err
	}

	var r Rescue
	if err := mapstructure.Decode(d["rescue"], &r); err != nil {
		return nil, err
	}
	return &r, nil
}
