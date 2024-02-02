#!/bin/sh
protoc --plugin=$GOPATH/bin/protoc-gen-go-plugin --go-plugin_out=. --go-plugin_opt=paths=source_relative common/plugin.proto
