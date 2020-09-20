.PHONY : build
build :
	go build -o ./bin/gofgen ./cmd/gofgen/*.go

.PHONY : run
run : build
	./bin/gofgen

.PHONY : test
test :
	go test ./...

.PHONY : install
install : 
	cp ./bin/gofgen /usr/local/bin

