.PHONY: build-win
build-win:
	go build -o ./cmd/citatnik/citatnik.exe ./cmd/citatnik/main.go

.PHONY: run-win
run-win: build-win
	./cmd/citatnik/citatnik.exe

.PHONY: build-linux
build-linux:
	go build -o ./cmd/citatnik/citatnik ./cmd/citatnik/main.go

.PHONY: run-linux
run-linux: build-linux
	chmod +x ./cmd/citatnik/citatnik
	./cmd/citatnik/citatnik