package hetzner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServersClient_List(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	l, err := client.Servers().List()
	require.NoError(t, err)

	expected := []*ServerSummary{
		{
			ServerIP:     "1.2.3.4",
			ServerNumber: 12345,
			ServerName:   "",
			Product:      "AX41-NVMe",
			DC:           "FSN1-DC1",
			Traffic:      "unlimited",
			Status:       "ready",
			Cancelled:    false,
			PaidUntil:    "2021-12-31",
			IP: []string{
				"1.2.3.4",
			},
			Subnet: []Subnet{
				{
					IP:   "::1",
					Mask: "128",
				},
			},
			LinkedStoragebox: Int(1234),
		},
	}
	require.Equal(t, expected, l)
}

func TestServersClient_Info(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	resp, err := client.Servers().Info("1.2.3.4")
	require.NoError(t, err)

	expected := &Server{
		ServerIP:     "1.2.3.4",
		ServerNumber: 12345,
		ServerName:   "",
		Product:      "AX41-NVMe",
		DC:           "FSN1-DC1",
		Traffic:      "unlimited",
		Status:       "ready",
		Cancelled:    false,
		PaidUntil:    "2021-12-31",
		IP:           []string{"1.2.3.4"},
		Subnet: []Subnet{
			{
				IP:   "::1",
				Mask: "128",
			},
		},
		Reset:            true,
		Rescue:           true,
		Vnc:              true,
		Windows:          false,
		Plesk:            false,
		CPanel:           false,
		WOL:              true,
		HotSwap:          false,
		LinkedStoragebox: Int(1234),
	}
	require.Equal(t, expected, resp)
}

func TestServersClient_Update(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	resp, err := client.Servers().Update(&ServerRequest{
		ServerIP:   "1.2.3.4",
		ServerName: "hello",
	})
	require.NoError(t, err)

	expected := &Server{
		ServerIP:     "1.2.3.4",
		ServerNumber: 12345,
		ServerName:   "hello",
		Product:      "AX41-NVMe",
		DC:           "FSN1-DC1",
		Traffic:      "unlimited",
		Status:       "ready",
		Cancelled:    false,
		PaidUntil:    "2021-12-31",
		IP:           []string{"1.2.3.4"},
		Subnet: []Subnet{
			{
				IP:   "::1",
				Mask: "128",
			},
		},
		Reset:            true,
		Rescue:           true,
		Vnc:              true,
		Windows:          false,
		Plesk:            false,
		CPanel:           false,
		WOL:              true,
		HotSwap:          false,
		LinkedStoragebox: Int(1234),
	}
	require.Equal(t, expected, resp)
}
