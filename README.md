# go-tamboon

## Installation

```sh
# Install gopls
go install golang.org/x/tools/gopls@latest

# Install golint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.2.2

# Install redocly:
sudo npm install @redocly/cli -g

# Install goplantuml:
go get github.com/jfeliu007/goplantuml/parser
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest