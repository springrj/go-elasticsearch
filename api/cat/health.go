// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"
)

// Health - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-health.html for more info.
//
// options: optional parameters. Supports the following functional options: WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithTs, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Health(options ...*Option) (*HealthResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Health"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &HealthResponse{resp}, err
}

// HealthResponse is the response for Health
type HealthResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}