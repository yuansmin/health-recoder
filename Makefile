run:
	# GIN_MODE=release
	GIN_MODE=debug go run -mod=vendor cmd/main.go
build:
	go build -mod=vendor -o bins/app cmd/main.go
vet:
	go vet cmd/main.go
vendor:
	go mod vendor

test:
	go test -v -count=1 ./tests/...

