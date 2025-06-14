package uri

// Improved URI library for Go
//
// Author: torstein at skybert dot  net

import (
	"net/url"
)

var portByScheme map[string]string

func init() {
	portByScheme = map[string]string{
		"http": "80",
		"https": "443",
		"ftp": "21",
		"scp": "22",
		"ldap": "389",
	}
}

type URI struct {
	queryParam map[string]string
	url        *url.URL
}

func New() *URI {
	return &URI{
		queryParam: make(map[string]string)}
}

func FromString(uri string) (*URI, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	ur := New()
	ur.url = u
	return ur, nil
}

type URIBuilder struct {
}

func (b *URI) SetHost(host string) error {
	b.url.Host = host
	return nil
}

func (b *URI) QueryParam(key string) string {
	return b.url.Query().Get(key)
}

func (b *URI) AddQueryParam(key, value string) error {
	// TODO Don't really need this, do we?
	b.queryParam[key] = value

	q := b.url.Query()
	q.Set(key, value)
	b.url.RawQuery = q.Encode()

	return nil
}

func (b *URI) RemoveQueryParam(key string) error {
	q := b.url.Query()
	q.Del(key)
	b.url.RawQuery = q.Encode()

	return nil
}


func (b *URI) Port() string {
	if b.url.Port() != "" {
		return b.url.Port()
	}

	return portByScheme[b.url.Scheme]
}
