build:
	GOOS=darwin go build -o ./bin/cloudsteps.Darwin.amd64 ./cmd/cloudsteps
	GOOS=linux GOARCH=amd64 go build -o ./bin/cloudsteps.Linux.amd64 ./cmd/cloudsteps
	GOOS=linux GOARCH=386 go build -o ./bin/cloudsteps.Linux.386 ./cmd/cloudsteps
	GOOS=linux GOARCH=arm go build -o ./bin/cloudsteps.Linux.armhf ./cmd/cloudsteps