package policy

import (
	"errors"
	"fmt"
	"strings"
)

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

func FindWithRoutingPath(path string) (Policy, error) {
	for _, p := range Policies {
		if strings.HasPrefix(path, p.RoutingPath) {
			return p, nil
		}
	}

	return Policy{}, errors.New("Policy not found for the given RoutingPath")
}
