BINARY_NAME := "today"
CURRENT_VERSION := `source ./VERSION && echo $VERSION_FULL`

@tidy:
	go mod tidy

@run name=BINARY_NAME *args="":
	go run ./cmd/{{ name }} {{ args }}

@build name=BINARY_NAME *args="":
	go build -ldflags "-X main.version={{ CURRENT_VERSION }}" -o bin/{{ name }} ./cmd/{{ name }}

@clean:
	rm -rf bin/{{ BINARY_NAME }}

@release:
	goreleaser release  --clean --skip-publish

@snapshot:
	goreleaser release  --clean --snapshot
