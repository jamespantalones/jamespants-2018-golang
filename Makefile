build:
	GOOS=linux go build -o jamespants .

dockerize:
	docker build -t jamespants-test .

run:
	docker run -it --publish 8080:8080 jamespants-test

all: build dockerize