.PHONY: lint
lint:
	golangci-lint run

.PHONY: rungomock
rungomock:
	sh ./script/gomock.sh

.PHONY: run
run:
	sh ./script/run.sh