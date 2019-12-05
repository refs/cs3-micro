[![Codacy Badge](https://api.codacy.com/project/badge/Grade/43bc431c4afd41bca5adc827513d3dc4)](https://www.codacy.com/manual/refs_2/cs3-micro?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=refs/cs3-micro&amp;utm_campaign=Badge_Grade)

# CS3-Micro

A proof of concept on how micro would play along with CERN's CS3 apis. It leverages teh CS3 api definitions and its implementation (as any of the multiple drivers) and simply implements micro client / server interfaces.

## Usage

## Generating Go gRPC from protobuf definitions
`prototool generate` to generate grpc and micro definitions

Start micro services

```sh
> micro api
```

Start the proof of concept:

`go run main.go`

Send a message to the service:

`micro call go.micro.api.authprovider ProviderAPI.Authenticate '{}'`

You should see the following response:

```json
{
        "user": {
                "username": "einstein",
                "mail": "einstein@example.org",
                "display_name": "Albert Einstein"
        }
}
```
