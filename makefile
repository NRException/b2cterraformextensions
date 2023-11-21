host_name = battlestation
namespace = b2cextensions
type = main
version = 0.0.1
target = linux_amd64

all: clean build install

tftest:
	cd terraform && TF_LOG=trace terraform init -plugin-dir=/home/nrexception/.terraform.d/plugins
install:
	mkdir -p  ~/.terraform.d/plugins/terraform.local/local/bcextensions/0.0.1/linux_amd64
	cp ./bin/terraform-provider-bcextensions ~/.terraform.d/plugins/terraform.local/local/bcextensions/0.0.1/linux_amd64/terraform-provider-bcextensions

build: 
	env GOOS=linux GOARCH=amd64 go build -o "./bin/terraform-provider-bcextensions"
	env GOOS=windows GOARCH=amd64 go build -o "./bin/terraform-provider-bcextensions.exe"

clean:
	rm -f ./bin/*
	rm -f ~/.terraform.d/plugins/terraform.local/local/bcextensions/0.0.1/linux_amd64/terraform-provider-bcextensions

