package handlers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// TestMain changes working dir for correct template path
func TestMain(m *testing.M) {
	os.Chdir("../")
	os.Exit(m.Run())
}

const (
	expectedInBody string = "You our 1 visitor"
	sessionName    string = "appsession"
)

func TestNewHomeHandler(t *testing.T) {
	logger := log.New(ioutil.Discard, "TEST:", log.Ldate|log.Ltime|log.Lshortfile)

	t.Run("NewHomeHandler should return no error if required template is present", func(t *testing.T) {
		if _, err := NewHomeHandler(logger); err != nil {
			t.Errorf("Expected err - \"%v\" to be nil", err)
		}
	})
}
func TestHomeHandlerServeHTTP(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "TEST:", log.Ldate|log.Ltime|log.Lshortfile)
	homeHandler, _ := NewHomeHandler(logger)

	ts := httptest.NewServer(homeHandler)
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("ServeHTTP for HomeHandler normally returns 200 status", func(t *testing.T) {
		if s := resp.StatusCode; s != http.StatusOK {
			t.Errorf("Expected %v Status Code, got - %v", http.StatusOK, s)
		}
	})

	t.Run("ServeHTTP for HomeHandler normally returns body with specific text", func(t *testing.T) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if !strings.Contains(string(body), expectedInBody) {
			t.Errorf("Expected body to contains \"%v\", got - \"%v\"", string(body), expectedInBody)
		}
	})

	t.Run("ServeHTTP for HomeHandler normally sets session cookie", func(t *testing.T) {
		var cookie *http.Cookie
		for _, c := range resp.Cookies() {
			if c.Name == sessionName {
				cookie = c
			}
		}
		if cookie == nil {
			t.Errorf("Failed to find cookie with name \"%v\"", sessionName)
		}
	})

}
