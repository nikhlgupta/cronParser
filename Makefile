.PHONY: setup
setup:
	@mkdir -p ./out

.PHONY: clean
clean:
	rm -rf ./out

.PHONY: build
build:
	@echo Building "./out/${APP}"...
	go build -mod=readonly -o  "./out/${APP}" cmd/cron_parser.go

.PHONY: test
test:
	go test -p 1 -v ./test/... -count=1 -covermode=atomic -coverpkg=./... -coverprofile=coverage-functional.out

.PHONY: run
run:
	make build
	./out/${APP}
