.PHONY: build test clean

build:
	go build -o bin/myvisor main.go

test:
	go test ./...

clean:
	rm -rf bin/

run: build
	bin/myvisor

sos: build
	bin/myvisor sos -H localhost -P 22334 -u msandbox -p msandbox
