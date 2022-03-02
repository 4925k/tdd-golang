package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
	t        *testing.T
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	content := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("request cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		content <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-content:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, fmt.Errorf("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("normal execution", func(t *testing.T) {
		content := "hello world"
		store := &StubStore{content, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != content {
			t.Errorf("got %v, want %v", response.Body.String(), content)
		}
	})

	t.Run("cancel work midway", func(t *testing.T) {
		content := "hello world"
		store := &StubStore{content, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("was not expecting a value")
		}

	})
}
