package uri

import "testing"

func assertEquals(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected: %v Actual: %v", expected, actual)
	}
}

func TestCreateBaseURI(t *testing.T) {
	input := "https://example.com"
	uri, err := FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}

	if uri == nil {
		t.Error("Shouldn't be nil")
	}

	assertEquals("example.com", uri.url.Host, t)
	assertEquals("https", uri.url.Scheme, t)
}

func TestOverrideHost(t *testing.T) {
	input := "https://example.com"
	uri, err := FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}

	newHost := "foo.bar"
	uri.SetHost(newHost)
	assertEquals(newHost, uri.url.Host, t)
	assertEquals("443", uri.Port(), t)
}

func TestDefaultPorts(t *testing.T) {
	input := "https://example.com"
	uri, err := FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}
	assertEquals("443", uri.Port(), t)

	input = "http://example.com"
	uri, err = FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}
	assertEquals("80", uri.Port(), t)
}

func TestPorts(t *testing.T) {
	input := "http://example.com:8000"
	uri, err := FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}
	assertEquals("8000", uri.Port(), t)
}

func TestAddQueryParameter(t *testing.T) {
	input := "https://example.com"
	uri, err := FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}

	uri.AddQueryParam("key", "v")
	assertEquals("v", uri.QueryParam("key"), t)
	uri.AddQueryParam("key", "v2")

	assertEquals("v2", uri.QueryParam("key"), t)
}

func TestRemoveQueryParameter(t *testing.T) {
	input := "https://example.com"
	uri, err := FromString(input)
	if err != nil {
		t.Error("Shouldn't return nil")
	}

	uri.AddQueryParam("key", "v")
	assertEquals("v", uri.QueryParam("key"), t)

	uri.RemoveQueryParam("key")
	assertEquals("", uri.QueryParam("key"), t)
}
