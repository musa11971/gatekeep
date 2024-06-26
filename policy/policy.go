package policy

import "fmt"

var Policies = [...]Policy{
	{
		RoutingPath:         "/v1/StarWars/People",
		EndpointFQDN:        "https://swapi.dev",
		EndpointReadTimeout: 5,
		EndpointPort:        443,
		EndpointPath:        "/api/people",
	},
	{
		RoutingPath:         "/v1/StarWars/Planets",
		EndpointReadTimeout: 5,
		EndpointFQDN:        "https://swapi.dev",
		EndpointPort:        443,
		EndpointPath:        "/api/planets",
	},
	{
		RoutingPath:         "/Google",
		EndpointReadTimeout: 5,
		EndpointFQDN:        "https://google.com",
		EndpointPort:        443,
		EndpointPath:        "/",
	},
}

type Policy struct {
	RoutingPath         string
	EndpointReadTimeout int
	EndpointFQDN        string
	EndpointPort        int
	EndpointPath        string
}

func (p Policy) FullEndpointURL() string {
	return fmt.Sprintf("%s:%d%s", p.EndpointFQDN, p.EndpointPort, p.EndpointPath)
}
