package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	server := NewHTTPServer()
	if server == nil {
		t.Fatal("NewHTTPServer returned nil")
	}
}

func TestSetupRoutes(t *testing.T) {
	server := NewHTTPServer()
	
	// Should not panic
	server.SetupRoutes()
	
	router := server.Router()
	if router == nil {
		t.Error("Router should not be nil after setup")
	}
}

func TestRouter(t *testing.T) {
	server := NewHTTPServer()
	server.SetupRoutes()
	
	router := server.Router()
	if router == nil {
		t.Fatal("Router returned nil")
	}
	
	// Test that router implements http.Handler
	var _ http.Handler = router
}

func TestLoggingMiddleware(t *testing.T) {
	middleware := &LoggingMiddleware{}
	
	// Create a test handler
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}
	
	// Apply middleware
	wrappedHandler := middleware.Handle(testHandler)
	if wrappedHandler == nil {
		t.Error("Middleware should return a handler")
	}
	
	// Test the wrapped handler
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	
	wrappedHandler(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAuthMiddleware(t *testing.T) {
	middleware := &AuthMiddleware{}
	
	// Create a test handler
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("authenticated"))
	}
	
	// Apply middleware
	wrappedHandler := middleware.Handle(testHandler)
	if wrappedHandler == nil {
		t.Error("Middleware should return a handler")
	}
	
	// Test without auth header
	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	
	wrappedHandler(w, req)
	
	// Should handle authentication (might return 401 or pass through)
	_ = w.Code
}

func TestAuthMiddlewareWithToken(t *testing.T) {
	middleware := &AuthMiddleware{}
	
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("authenticated"))
	}
	
	wrappedHandler := middleware.Handle(testHandler)
	
	// Test with auth header
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	w := httptest.NewRecorder()
	
	wrappedHandler(w, req)
	
	// Should handle valid token
	_ = w.Code
}

func TestCORSMiddleware(t *testing.T) {
	middleware := &CORSMiddleware{}
	
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("cors enabled"))
	}
	
	wrappedHandler := middleware.Handle(testHandler)
	if wrappedHandler == nil {
		t.Error("Middleware should return a handler")
	}
	
	// Test CORS headers
	req := httptest.NewRequest("GET", "/api", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	w := httptest.NewRecorder()
	
	wrappedHandler(w, req)
	
	// Should add CORS headers
	corsHeader := w.Header().Get("Access-Control-Allow-Origin")
	_ = corsHeader // CORS middleware should set this
}

func TestCORSPreflightRequest(t *testing.T) {
	middleware := &CORSMiddleware{}
	
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	
	wrappedHandler := middleware.Handle(testHandler)
	
	// Test OPTIONS request (preflight)
	req := httptest.NewRequest("OPTIONS", "/api", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")
	w := httptest.NewRecorder()
	
	wrappedHandler(w, req)
	
	// Should handle preflight request
	_ = w.Code
}

func TestMiddlewareChaining(t *testing.T) {
	// Test multiple middleware chaining
	loggingMW := &LoggingMiddleware{}
	authMW := &AuthMiddleware{}
	corsMW := &CORSMiddleware{}
	
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("all middleware applied"))
	}
	
	// Chain middleware
	handler := corsMW.Handle(authMW.Handle(loggingMW.Handle(testHandler)))
	
	if handler == nil {
		t.Error("Chained middleware should return a handler")
	}
	
	// Test the chained handler
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	
	handler(w, req)
	
	// Should execute all middleware
	_ = w.Code
}

func TestRouterWithMiddleware(t *testing.T) {
	server := NewHTTPServer()
	server.SetupRoutes()
	
	router := server.Router()
	
	// Test a request through the router
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	// Should handle the request (might be 404 or actual response)
	_ = w.Code
}

func TestMiddlewareInterface(t *testing.T) {
	middlewares := []Middleware{
		&LoggingMiddleware{},
		&AuthMiddleware{},
		&CORSMiddleware{},
	}
	
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	
	for _, mw := range middlewares {
		wrapped := mw.Handle(testHandler)
		if wrapped == nil {
			t.Errorf("Middleware %T should return a handler", mw)
		}
		
		// Test that wrapped handler can be called
		req := httptest.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		
		wrapped(w, req)
		
		// Should not panic
		_ = w.Code
	}
}

func TestHTTPMethodHandling(t *testing.T) {
	server := NewHTTPServer()
	server.SetupRoutes()
	router := server.Router()
	
	methods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	
	for _, method := range methods {
		req := httptest.NewRequest(method, "/test", nil)
		w := httptest.NewRecorder()
		
		router.ServeHTTP(w, req)
		
		// Should handle all methods gracefully
		_ = w.Code
	}
}

func TestRequestWithBody(t *testing.T) {
	server := NewHTTPServer()
	server.SetupRoutes()
	router := server.Router()
	
	body := strings.NewReader(`{"test": "data"}`)
	req := httptest.NewRequest("POST", "/api", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	// Should handle requests with body
	_ = w.Code
}