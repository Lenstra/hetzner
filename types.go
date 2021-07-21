package hetzner

type ServersClient struct {
	c *Client
}

type ServerSummary struct {
	ServerIP         string   `mapstructure:"server_ip"`
	ServerNumber     int      `mapstructure:"server_number"`
	ServerName       string   `mapstructure:"server_name"`
	Product          string   `mapstructure:"product"`
	DC               string   `mapstructure:"dc"`
	Traffic          string   `mapstructure:"traffic"`
	Status           string   `mapstructure:"status"`
	Cancelled        bool     `mapstructure:"cancelled"`
	PaidUntil        string   `mapstructure:"paid_until"`
	IP               []string `mapstructure:"ip"`
	Subnet           []Subnet `mapstructure:"subnet"`
	LinkedStoragebox *int     `mapstructure:"linked_storagebox"`
}

type Server struct {
	ServerIP         string   `mapstructure:"server_ip"`
	ServerNumber     int      `mapstructure:"server_number"`
	ServerName       string   `mapstructure:"server_name"`
	Product          string   `mapstructure:"product"`
	DC               string   `mapstructure:"dc"`
	Traffic          string   `mapstructure:"traffic"`
	Status           string   `mapstructure:"status"`
	Cancelled        bool     `mapstructure:"cancelled"`
	PaidUntil        string   `mapstructure:"paid_until"`
	IP               []string `mapstructure:"ip"`
	Subnet           []Subnet `mapstructure:"subnet"`
	Reset            bool     `mapstructure:"reset"`
	Rescue           bool     `mapstructure:"rescue"`
	Vnc              bool     `mapstructure:"vnc"`
	Windows          bool     `mapstructure:"windows"`
	Plesk            bool     `mapstructure:"plesk"`
	CPanel           bool     `mapstructure:"cpanel"`
	WOL              bool     `mapstructure:"wol"`
	HotSwap          bool     `mapstructure:"hot_swap"`
	LinkedStoragebox *int     `mapstructure:"linked_storagebox"`
}

type Subnet struct {
	IP   string `mapstructure:"ip"`
	Mask string `mapstructure:"mask"`
}

type Rescue struct {
	ServerIP       string   `mapstructure:"server_ip"`
	ServerNumber   int      `mapstructure:"server_number"`
	OS             string   `mapstructure:"os"`
	Arch           int      `mapstructure:"arch"`
	Active         bool     `mapstructure:"active"`
	Password       string   `mapstructure:"password"`
	AuthorizedKeys []string `mapstructure:"authorized_key"`
	HostKeys       []string `mapstructure:"host_key"`
}

type Reset struct {
	ServerIP string `mapstructure:"server_ip"`
	Type     string `mapstructure:"type"`
}
