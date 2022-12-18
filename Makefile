COVERPROFILE=.cover.out
COVERDIR=.cover

dep:
	@go get ./...
	
run: 
	@go run main.go

test: 
	@go test -coverprofile=$(COVERPROFILE) ./...

local-cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERPROFILE) -o $(COVERDIR)/index.html