package nfa

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var ErrCfgNoUsername = errors.New("username required")
var ErrCfgNoPassword = errors.New("username required")
var ErrCfgNoURL = errors.New("URL required")

type Options struct {
	Insecure bool
	URL      *url.URL
	Username string
	Password string
}

func (opts *Options) Validate() error {
	if opts.Username == "" {
		return ErrCfgNoUsername
	}
	if opts.Password == "" {
		return ErrCfgNoPassword
	}
	if opts.URL == nil {
		return ErrCfgNoURL
	}
	return nil
}

type OptionSetter func(*Options)

// Insecure disables SSL certificate validation when communicating with NFA.
func Insecure() OptionSetter {
	return func(opts *Options) {
		opts.Insecure = true
	}
}

// Username sets the authentication username. This option is required.
func Username(u string) OptionSetter {
	return func(opts *Options) {
		opts.Username = u
	}
}

// Password sets the authentication password. This option is required.
func Password(p string) OptionSetter {
	return func(opts *Options) {
		opts.Password = p
	}
}

// URL sets the NFA base URL. This option is required. See ParseURL helper.
func URL(u *url.URL) OptionSetter {
	return func(opts *Options) {
		opts.URL = u
	}
}

// ParseURL parses an NFA URL from a string to a *url.URL, sets the appropriate URI schema, and
// sets the correct path.
func ParseURL(u string) (*url.URL, error) {
	if !strings.HasPrefix(u, "http") {
		u = fmt.Sprintf("https://%s", u)
	}
	pu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	pu.Scheme = "https"
	if pu.Path == "/" || pu.Path == "" {
		pu.Path = "/api"
	}
	return pu, nil
}
