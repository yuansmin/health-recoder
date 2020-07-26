run:
	# GIN_MODE=release
	GIN_MODE=debug go run -mod=vendor cmd/main.go
build:
	go build -mod=vendor -o bins/app cmd/main.go

.PHONY: vet
vet:
	go vet cmd/main.go

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: test
test:
	go test -v -count=1 ./tests/...

