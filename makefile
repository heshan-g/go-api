build:
	go build -o bin/server

run: build
	./bin/server

nodemon:
	nodemon -e go --signal SIGTERM --exec 'go build -o bin/server main.go && ./bin/server'
