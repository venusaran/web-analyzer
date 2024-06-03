package scraper

import (
	"testing"

	"golang.org/x/net/html"
)

func TestIsLoginForm(t *testing.T) {
	tests := []struct {
		name     string
		token    html.Token
		expected bool
	}{
		{
			name: "Login form identified by id attribute",
			token: html.Token{
				Attr: []html.Attribute{
					{Key: "id", Val: "loginForm"},
				},
			},
			expected: true,
		},
		{
			name: "Login form identified by id attribute with mixed case",
			token: html.Token{
				Attr: []html.Attribute{
					{Key: "id", Val: "LoginForm"},
				},
			},
			expected: true,
		},
		{
			name: "Not a login form with different id",
			token: html.Token{
				Attr: []html.Attribute{
					{Key: "id", Val: "registerForm"},
				},
			},
			expected: false,
		},
		{
			name: "Not a login form without id",
			token: html.Token{
				Attr: []html.Attribute{
					{Key: "class", Val: "form"},
				},
			},
			expected: false,
		},
		{
			name:     "Not a login form without any attributes",
			token:    html.Token{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isLoginForm(tt.token)
			if result != tt.expected {
				t.Errorf("isLoginForm() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetHTMLVersion(t *testing.T) {
	tests := []struct {
		name     string
		doctype  string
		expected string
	}{
		{
			name:     "HTML 4.01",
			doctype:  "html PUBLIC \"-//W3C//DTD HTML 4.01//EN\"",
			expected: "HTML 4.01",
		},
		{
			name:     "XHTML 1.2",
			doctype:  "html PUBLIC \"-//W3C//DTD XHTML 1.2//EN\"",
			expected: "XHTML 1.2",
		},
		{
			name:     "HTML5",
			doctype:  "html",
			expected: "HTML5",
		},
		{
			name:     "Unknown HTML version",
			doctype:  "unknown",
			expected: "Unknown HTML version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := html.Token{
				Data: tt.doctype,
			}
			version, err := getHTMLVersion(token)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if version != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, version)
			}
		})
	}
}

func TestGetHref(t *testing.T) {
	tests := []struct {
		input         html.Token
		expectedFound bool
		expectedHref  string
	}{
		// Test case where href attribute is found
		{html.Token{Type: html.StartTagToken, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "https://example.com"}}}, true, "https://example.com"},
		// Test case where href attribute is not found
		{html.Token{Type: html.StartTagToken, Data: "a", Attr: []html.Attribute{{Key: "class", Val: "link"}}}, false, ""},
		// Test case where input token is not a start tag
		{html.Token{Type: html.TextToken, Data: "Hello World"}, false, ""},
		// Test case where input token has no attributes
		{html.Token{Type: html.StartTagToken, Data: "a"}, false, ""},
	}

	for _, test := range tests {
		found, href := getHref(test.input)
		if found != test.expectedFound || href != test.expectedHref {
			t.Errorf("getHref(%v) = (%t, %s), expected (%t, %s)", test.input, found, href, test.expectedFound, test.expectedHref)
		}
	}
}
