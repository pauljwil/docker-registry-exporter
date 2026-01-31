#!/bin/sh

exec ./docker-registry-exporter \
 --registry-address=$REGISTRY_ADDRESS \
 --listen-address=0.0.0.0:9055