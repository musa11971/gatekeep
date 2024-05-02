package policy

import "fmt"

var Policies = [...]Policy{
	{
		RoutingPath:  "/v1/StarWars/People",
		EndpointFQDN: "https://swapi.dev",
		EndpointPort: 443,
		EndpointPath: "/api/people",
	},
	{
		RoutingPath:  "/v1/StarWars/Planets",
		EndpointFQDN: "https://swapi.dev",
		EndpointPort: 443,
		EndpointPath: "/api/planets",
	},
}

type Policy struct {
	RoutingPath  string
	EndpointFQDN string
	EndpointPort int
	EndpointPath string
}

func (p Policy) FullEndpointURL() string {
	return fmt.Sprintf("%s:%d%s", p.EndpointFQDN, p.EndpointPort, p.EndpointPath)
}
