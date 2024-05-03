package policy

import "testing"

func TestFullEndpointURL(t *testing.T) {
	policy := Policy{
		RoutingPath:         "/v1/Example/People",
		EndpointFQDN:        "https://example.com",
		EndpointReadTimeout: 5,
		EndpointPort:        443,
		EndpointPath:        "/api/people",
	}

	fullEndpoint := policy.FullEndpointURL()

	if fullEndpoint != "https://example.com:443/api/people" {
		t.Errorf("FullEndpointURL is not properly formatted: %s", fullEndpoint)
	}
}
