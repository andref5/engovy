# Envoy control-plane example
This is a simple example of Envoy control plane connected via gRPC using Go

## How it works
```
                        | /config_dump (1900)
                        | 
                        |
  http req (8000)    +---------------+   proxy to 5000   +-----------------+
---------------------|     envoy     |-------------------| upstream-tester |
  (/ping | /pong)    +---------------+  (/ping | /pong)  +-----------------+
                           |  |
                           |  |
                           |  | gRPC (18000)
                           |  |
                           |  |
                     +---------------+
                     |    golang     |
                     | control-plane |
                     +---------------+
                        |
                        |
                        | /changeRoute (8001)
```

```
├── main.go (Startup control-plane(18000) and a http server(8001) to test change envoy configs)
├── pkg
│   ├── control-plane (Envoy control plane impl. Adapted from https://github.com/envoyproxy/go-control-plane/tree/main/internal/)
│   │   ├── logger.go (simple logger impl to set go-control-plane)
│   │   ├── resource.go (funcs to create envoy API structs https://www.envoyproxy.io/docs/envoy/latest/api/api)
│   │   └── server.go (control plane gRPC server. Envoy will establish gRPC stream to sync configs)
│   └── upstream-tester (Simple Go http server to test proxy to upstream cluster)
│       └── main.go (port 5000 - paths "/ping" "/pong")
└── deploy
    └── envoy-bootstrap-xds.yaml (Base config for Envoy xDS https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration)
```

## Running

#### Install getenvoy.io
```
sudo make install
```

#### Runs an instance of Envoy
```
make envoy
```

#### New terminal to startup golang gRPC control-plane and HTTP server
```
go run main.go
```

#### New terminal to startup upstream tester
```
cd pkg/upstream-tester
go run main.go
```

## Testing

### /*\ Only a simple path change implemented
```
curl -sS http://localhost:19000/config_dump | jq '.configs[1].dynamic_active_clusters'
curl -sS http://localhost:19000/config_dump | jq '.configs[4].dynamic_route_configs'

curl -v localhost:8000/ping

curl -v localhost:8001/changeRoute -d path=/pong

curl -sS http://localhost:19000/config_dump | jq '.configs[4].dynamic_route_configs'

curl -v localhost:8000/pong
```

## References

- https://www.envoyproxy.io/docs/envoy/latest/
- https://github.com/envoyproxy/go-control-plane
- https://github.com/envoyproxy/management-plane-api
- https://github.com/solo-io/gloo
- https://dropbox.tech/infrastructure/how-we-migrated-dropbox-from-nginx-to-envoy
- https://eng.uber.com/gatewayuberapi/
- https://eng.uber.com/architecture-api-gateway