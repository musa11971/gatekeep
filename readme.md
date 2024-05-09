# Gatekeep
Gatekeep is a proof of concept API gateway written in Go. By acting as a single entry point for all client requests, Gatekeep aims to provide a centralized solution for controlling access, enforcing security policies, and so forth.

## Features
- Policy based routing to API endpoints.
- Per-policy configuration of:
    - Routing path
    - Endpoint FQDN
    - Endpoint read timeout
    - Endpoint port
    - Endpoint path

## To do
- Improve tests.
- The URL path after the RoutingPath should be forwarded to the API endpoint.
- The request (POST-)data should be forwarded to the API endpoint.
- Policies should be managed externally, not hard-coded.
- Logequests.
- Authentication.

## Getting started
Running locally:
```sh
git clone https://github.com/musa11971/gatekeep && cd gatekeep
go run .
```

To run tests locally use the following command:
```sh
go test -v ./...
```

## Contributing
Contributions are welcome! If you have ideas for new features, improvements, or bug fixes, please submit a pull request or open an issue on the GitHub repository.