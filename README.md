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
