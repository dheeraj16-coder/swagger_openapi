#!/bin/sh

# Start ADOT collector in background
aws-otel-collector --config /etc/adot/config.yaml &

# Start Go server in foreground
exec server
