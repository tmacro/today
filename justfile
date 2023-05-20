BINARY_NAME := "today"
CURRENT_VERSION := `source ./VERSION && echo $VERSION_FULL`

@tidy:
	go mod tidy

@run name=BINARY_NAME *args="":
	#!/bin/bash
	go run cmd/{{ name }}/main.go {{ args }}

@build name=BINARY_NAME *args="":
	go build -ldflags "-X main.version={{ CURRENT_VERSION }}" -o bin/{{ name }} cmd/{{ name }}/main.go

@clean:
	rm -rf bin/{{ BINARY_NAME }}
