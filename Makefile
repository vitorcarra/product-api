check_install:
	which swagger || GO111MODULE=off go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models