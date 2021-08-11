package hetzner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirewallClient_Info(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	resp, err := client.Firewall().Info("1.2.3.4")
	require.NoError(t, err)

	expected := &Firewall{
		ServerIP:     "1.2.3.4",
		ServerNumber: 1234,
		Status:       "disabled",
		WhitelistHOS: true,
		Port:         "main",
		Rules: FirewallRules{
			Input: []*FirewallRule{
				{
					IPVersion: "ipv4",
					Name:      "Allow SSH",
					DstIP:     String("0.0.0.0/0"),
					SrcIP:     String("0.0.0.0/0"),
					DstPort:   String("22"),
					SrcPort:   String("0-65535"),
					Protocol:  String("tcp"),
					TcpFlags:  nil,
					Action:    "accept",
				},
				{

					IPVersion: "ipv4",
					Name:      "Allow HTTPS",
					DstIP:     String("0.0.0.0/0"),
					SrcIP:     String("0.0.0.0/0"),
					DstPort:   String("443"),
					SrcPort:   String("0-65535"),
					Protocol:  nil,
					TcpFlags:  nil,
					Action:    "accept",
				},
			},
		},
	}
	require.Equal(t, expected, resp)
}

func TestFirewallClient_Update(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	resp, err := client.Firewall().Update(&FirewallRequest{
		ServerIP:     "1.2.3.4",
		Status:       String("active"),
		WhitelistHOS: Bool(false),
		TemplateID:   nil,
		Rules: FirewallRules{
			Input: []*FirewallRule{
				{
					IPVersion: "ipv4",
					Name:      "rule 1",
					SrcIP:     String("1.1.1.1"),
					DstPort:   String("80"),
					Action:    "accept",
				},
				{
					IPVersion: "ipv4",
					Name:      "Allow MySQL",
					DstPort:   String("3306"),
					Action:    "accept",
				},
			},
		},
	})
	require.NoError(t, err)

	expected := &Firewall{
		ServerIP:     "1.2.3.4",
		ServerNumber: 1234,
		Status:       "in process",
		WhitelistHOS: true,
		Port:         "main",
		Rules: FirewallRules{
			Input: []*FirewallRule{
				{
					IPVersion: "ipv4",
					Name:      "rule 1",
					SrcIP:     String("1.1.1.1"),
					DstPort:   String("80"),
					Action:    "accept",
				},
				{

					IPVersion: "ipv4",
					Name:      "Allow MySQL",
					DstPort:   String("3306"),
					Action:    "accept",
				},
			},
		},
	}
	require.Equal(t, expected, resp)
}

func TestFirewallClient_Delete(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	resp, err := client.Firewall().Delete("1.2.3.4")
	require.NoError(t, err)

	expected := &Firewall{
		ServerIP:     "1.2.3.4",
		ServerNumber: 1234,
		Status:       "in process",
		WhitelistHOS: true,
		Port:         "main",
		Rules: FirewallRules{
			Input: nil,
		},
	}
	require.Equal(t, expected, resp)
}

func TestFirewallClient_EmptyRules(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	expected := &Firewall{
		ServerIP:     "2.3.4.5",
		ServerNumber: 1234,
		WhitelistHOS: true,
		Port:         "main",
		Status:       "disabled",
	}

	resp, err := client.Firewall().Info("2.3.4.5")
	require.NoError(t, err)
	require.Equal(t, expected, resp)

	resp, err = client.Firewall().Update(&FirewallRequest{
		ServerIP: "2.3.4.5",
	})
	require.NoError(t, err)
	require.Equal(t, expected, resp)

	resp, err = client.Firewall().Delete("2.3.4.5")
	require.NoError(t, err)
	require.Equal(t, expected, resp)
}
