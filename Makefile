build_react:
	@cd ./react/my-app/ && npm run build

build: build_react
	@go build -o ./bin/golang-react-app

run: build
	@./bin/golang-react-app
