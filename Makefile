GOFILES= $$(go list -f '{{join .GoFiles " "}}')

run:
	go run $(GOFILES)

build:
	go build -o $(GOPATH)/bin/graphql $(GOFILES)

dgraph:
	docker run --rm -it -p 8000:8000 -p 8080:8080 -p 9080:9080 dgraph/standalone:master
