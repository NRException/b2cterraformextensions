all: clean build

build: 
	env GOOS=linux GOARCH=amd64 go build -o "bin/terraform-provider-b2cextensions"
	env GOOS=windows GOARCH=amd64 go build -o "bin/terraform-provider-b2cextensions.exe"

clean:
	rm ./bin/*
