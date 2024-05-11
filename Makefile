build:
	@go build -o ./bin/golang-react-app

run: build
	@./bin/golang-react-app
