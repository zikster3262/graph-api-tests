init:
	go run github.com/99designs/gqlgen init
server:
	go run ./server.go
generate:
	go get github.com/99designs/gqlgen@v0.17.9; go run github.com/99designs/gqlgen generate
