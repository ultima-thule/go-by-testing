package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if want != got {
			t.Errorf("got %q want %q", want, got)
		}
	})

	t.Run("returns an error if server does not respond within 10 sec", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Millisecond)

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
