package hetzner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBootClient_Reset(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	r, err := client.Reset().Reset(&ResetRequest{
		ServerIP: "1.2.3.4",
		Type:     "hw",
	})
	require.NoError(t, err)

	expected := &Reset{
		ServerIP: "1.2.3.4",
		Type:     "hw",
	}
	require.Equal(t, expected, r)
}
