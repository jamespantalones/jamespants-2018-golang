TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

install:
	go get .

build: install
	go build -ldflags "-X main.version=$(TAG)" -o jamespants .

