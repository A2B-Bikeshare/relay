all:
	go install

build:
	go get -d -v ./... && go build -v ./...

test:
	go test -v

travis: 
	go get -d -v ./... && go build -v ./...	
	go test -v ./...