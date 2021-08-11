package hetzner

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func testServer(t *testing.T) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		require.NoError(t, err)

		method := strings.ToLower(r.Method)

		parts := strings.Split(r.URL.Path, "/")
		last := parts[len(parts)-1]
		parts[len(parts)-1] = fmt.Sprintf("%s-req-%s.txt", method, last)

		path := fmt.Sprintf("fixtures%s", strings.Join(parts, "/"))
		expected, err := os.ReadFile(path)
		if err != nil {
			var pathErr *os.PathError
			require.ErrorAs(t, err, &pathErr)
		}
		expected = bytes.TrimSpace(expected)
		require.Equal(t, string(expected), string(body), "Request body does not match %q", path)

		parts = strings.Split(r.URL.Path, "/")
		last = parts[len(parts)-1]
		parts[len(parts)-1] = fmt.Sprintf("%s-%s.json", method, last)
		path = "fixtures" + strings.Join(parts, "/")

		log.Printf("Serving %s as response", path)
		http.ServeFile(w, r, path)
	}))

	client := NewClient(&Config{
		Address: server.URL,
	})

	return client, server
}

func TestStorageBoxClient_List(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	l, err := client.StorageBox().List()
	require.NoError(t, err)

	expected := []*StorageBoxSummary{
		{
			ID:           123,
			Login:        "u1234",
			Name:         "",
			Product:      "BX10 - inclusive",
			Cancelled:    false,
			Locked:       false,
			Location:     "FSN1",
			LinkedServer: 4567,
			PaidUntil:    "2021-12-31",
		},
	}
	require.Equal(t, expected, l)
}

func TestStorageBoxClient_Info(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	b, err := client.StorageBox().Info(1234)
	require.NoError(t, err)

	expected := &StorageBox{
		ID:                   1234,
		Login:                "u12345",
		Name:                 "",
		Product:              "BX10 - inclusive",
		Cancelled:            false,
		Locked:               false,
		Location:             "FSN1",
		LinkedServer:         5678,
		PaidUntil:            "2021-12-31",
		DiskQuota:            102400,
		DiskUsage:            0,
		DiskUsageData:        0,
		DiskUsageSnapshots:   0,
		Webdav:               false,
		Samba:                false,
		SSH:                  false,
		ExternalReachability: false,
		ZFS:                  false,
		Server:               "u12345.your-storagebox.de",
		HostSystem:           "FSN1-BX123",
	}
	require.Equal(t, expected, b)
}

func TestStorageBoxClient_Update(t *testing.T) {
	client, s := testServer(t)
	defer s.Close()

	b, err := client.StorageBox().Update(&StorageBoxRequest{
		ID:             1234,
		StorageBoxName: String("hello"),
	})
	require.NoError(t, err)

	expected := &StorageBoxSummary{
		ID:           1234,
		Login:        "u12345",
		Name:         "hello",
		Product:      "BX10 - inclusive",
		Cancelled:    false,
		Locked:       false,
		Location:     "FSN1",
		LinkedServer: 5678,
		PaidUntil:    "2021-12-31",
	}
	require.Equal(t, expected, b)
}
