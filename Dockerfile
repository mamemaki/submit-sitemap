FROM golang:1.23-bookworm

# Install common tools
RUN apt-get update \
    && apt-get install -y git jq unzip tar gzip less

# Install go packages
RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3 \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/mgechev/revive@latest
