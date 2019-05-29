build:
	go build main.go
	docker build -t go_plus_python .
run:
	docker run -d -it go_plus_python
