package git

import (
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestJenkins(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/job/test-job/build" && r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
			}
		}))
		defer server.Close()

		err := Jenkins("user", "pass", server.URL, "test-job", "secret-token")
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	})

	t.Run("Unauthorized", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer server.Close()

		err := Jenkins("user", "pass", server.URL, "test-job", "secret-token")
		if err == nil || err.Error() != "authentication failed: invalid credentials" {
			t.Errorf("Expected authentication failed error, but got %v", err)
		}
	})

	t.Run("Job Not Found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		err := Jenkins("user", "pass", server.URL, "non-existent-job", "secret-token")
		if err == nil || err.Error() != "job not found: check the job name and URL" {
			t.Errorf("Expected job not found error, but got %v", err)
		}
	})
}

func TestJenkinsLog(t *testing.T){
	t.Run("Success",func (t *testing.T){
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/job/test-job/lastBuild/consoleText" && r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
				_,err := w.Write([]byte("Hello"))
				if err != nil {
					t.Errorf("Error writing response: %v", err)
				}
			}
		}))
		defer server.Close()

		log,err := JenkinsLog("user", "pass", server.URL, "test-job",0)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if log != "Hello" {
			t.Errorf("Body is: %s",log)
		}
	})
	t.Run("Unauthorized",func (t *testing.T){
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer server.Close()

		_,err := JenkinsLog("user", "pass", server.URL, "test-job",0)
		if err == nil || err.Error() != "authentication failed: invalid credentials" {
			t.Errorf("Expected authentication failed error, but got %v", err)
		}
	})
	t.Run("Job_Not_Found",func(t *testing.T){
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		_,err := JenkinsLog("user", "pass", server.URL, "non-existent-job",0)
		if err == nil || err.Error() != "job not found: check the job name and URL" {
			t.Errorf("Expected job not found error, but got %v", err)
		}
	})
}
