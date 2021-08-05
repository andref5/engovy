

install:
	curl -L https://getenvoy.io/cli | bash -s -- -b /usr/local/bin

ENVOY_VERSION:=1.19.0
envoy:
	func-e run -c ./deploy/envoy-bootstrap-xds.yaml --drain-time-s 1 -l debug

