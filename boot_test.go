package hetzner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBootClient_Rescue(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	r, err := client.Boot().Rescue(&RescueRequest{
		ServerIP: "1.2.3.4",
		OS:       "linux",
		Arch:     64,
	})
	require.NoError(t, err)

	expected := &Rescue{
		ServerIP:       "1.2.3.4",
		ServerNumber:   1234,
		OS:             "linux",
		Arch:           64,
		Active:         true,
		Password:       "secret",
		AuthorizedKeys: []string{},
		HostKeys:       []string{},
	}
	require.Equal(t, expected, r)
}
