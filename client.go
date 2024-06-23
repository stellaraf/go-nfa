package nfa

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-nfa/flow"
)

type Client struct {
	http *resty.Client
}

type AuthResponse struct {
	Result struct {
		Code    string `json:"code"`
		Message string `json:"message,omitempty"`
	} `json:"result"`
}

type ProcessingResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

var ErrZeroPercentile = errors.New("percentile value was zero; retry query")

func getAuthCookies(u *url.URL, user, pass string) ([]*http.Cookie, error) {
	req := resty.New().R().
		SetBody(map[string]string{"username": user, "password": pass}).
		SetResult(&AuthResponse{}).
		SetError(&AuthResponse{})

	u.Path += "/login"
	res, err := req.Post(u.String())
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		data, ok := res.Error().(*AuthResponse)
		if !ok {
			return nil, fmt.Errorf("failed to authenticate with %s", req.URL)
		}
		return nil, fmt.Errorf("%s: %s", data.Result.Code, data.Result.Message)
	}
	return res.Cookies(), nil
}

func (client *Client) PercentileQuery(prefixes, exclude []string, percentile uint8) (*flow.ResPercentileQuery, error) {
	q, err := NewPercentileQuery(prefixes, exclude, percentile)
	if err != nil {
		return nil, err
	}
	qp, err := q.MarshalQuery()
	if err != nil {
		return nil, err
	}
	req := client.http.R().
		SetQueryParams(qp).
		SetResult(&flow.ResPercentileQuery{}).
		AddRetryCondition(func(res *resty.Response, _ error) bool {
			return res.StatusCode() == http.StatusAccepted
		})
	res, err := req.Get("/reports/flows")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(string(res.Body()))
	}
	data, ok := res.Result().(*flow.ResPercentileQuery)
	if !ok {
		return nil, fmt.Errorf("failed to parse response")
	}
	if data.PercentileValue == uint64(0) {
		return nil, ErrZeroPercentile
	}
	return data, nil
}

func New(options ...OptionSetter) (*Client, error) {
	r := resty.New()
	opts := &Options{
		Insecure:   false,
		RetryCount: 5,
		RetryTime:  time.Second * 3,
	}
	for _, setter := range options {
		setter(opts)
	}
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	r.SetBaseURL(opts.URL.String())
	cookies, err := getAuthCookies(opts.URL, opts.Username, opts.Password)
	if err != nil {
		return nil, err
	}
	client := &Client{
		http: r,
	}
	r.SetCookies(cookies)
	r.OnBeforeRequest(func(rc *resty.Client, req *resty.Request) error {
		for _, c := range req.Cookies {
			if c.Name == "Expires" {
				expiry, err := time.Parse(time.RFC1123, c.Value)
				if err != nil {
					return fmt.Errorf("failed to parse session expiry")
				}
				if time.Now().After(expiry) {
					cookies, err := getAuthCookies(opts.URL, opts.Username, opts.Password)
					if err != nil {
						return err
					}
					client.http.SetCookies(cookies)
				}
			}
		}
		return nil
	})

	r.SetRetryCount(opts.RetryCount)
	r.SetRetryWaitTime(opts.RetryTime)
	r.SetRetryMaxWaitTime(opts.RetryTime * time.Duration(opts.RetryCount))

	r.AddRetryCondition(func(res *resty.Response, _ error) bool {
		return res.StatusCode() == http.StatusUnauthorized
	})
	r.AddRetryHook(func(res *resty.Response, err error) {
		if res.StatusCode() == http.StatusUnauthorized {
			cookies, err := getAuthCookies(opts.URL, opts.Username, opts.Password)
			if err == nil {
				client.http.SetCookies(cookies)
			}
		}
	})
	return client, nil
}
