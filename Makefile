all:
	go install

test:
	go test -v

travis: 
	go get -d -v ./... && go build -v ./...	
	go test -v ./...