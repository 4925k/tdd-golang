package selectTesting

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return fast server", func(t *testing.T) {
		slowServer := delayedServer(20 * time.Millisecond)
		fastServer := delayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("return error for taking more than 10s", func(t *testing.T) {
		serverA := delayedServer(1 * time.Second)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 10*time.Millisecond)
		if err == nil {
			t.Error("wanted an error but got none")
		}
	})

}

func delayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		fmt.Println(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
