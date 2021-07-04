package freemyip

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (*Client, *http.ServeMux) {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client := New("token", true)
	client.baseURL, _ = url.Parse(server.URL)

	return client, mux
}

func testHandler(body string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(rw, fmt.Sprintf("unsupported method: %s", req.Method), http.StatusMethodNotAllowed)
			return
		}

		_, err := fmt.Fprintln(rw, body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TestClient_UpdateDomain(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("OK\nIP 81.220.38.196 didn't change. No need to update record."))

	_, err := client.UpdateDomain(context.Background(), "example", "")
	require.NoError(t, err)
}

func TestClient_UpdateDomain_error(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("ERROR\nToken not provided"))

	_, err := client.UpdateDomain(context.Background(), "example", "")
	require.Error(t, err)
}

func TestClient_DeleteDomain(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("OK\nmessage"))

	_, err := client.DeleteDomain(context.Background(), "example")
	require.NoError(t, err)
}

func TestClient_DeleteDomain_error(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("ERROR\nToken not provided"))

	_, err := client.DeleteDomain(context.Background(), "example")
	require.Error(t, err)
}

func TestClient_EditTXTRecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("OK\nUpdated TXT for domain example.freemyip.com to: value"))

	_, err := client.EditTXTRecord(context.Background(), "example", "value")
	require.NoError(t, err)
}

func TestClient_EditTXTRecord_error(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("ERROR\nToken not provided"))

	_, err := client.EditTXTRecord(context.Background(), "example", "value")
	require.Error(t, err)
}

func TestClient_DeleteTXTRecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("OK\nUpdated TXT for domain example.freemyip.com to: null"))

	_, err := client.DeleteTXTRecord(context.Background(), "example")
	require.NoError(t, err)
}

func TestClient_DeleteTXTRecord_error(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/", testHandler("ERROR\nToken not provided"))

	_, err := client.DeleteTXTRecord(context.Background(), "example")
	require.Error(t, err)
}
