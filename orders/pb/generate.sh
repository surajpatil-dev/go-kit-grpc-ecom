#!/usr/bin/env sh
protoc ./order.proto --go_out=plugin=grpc:.