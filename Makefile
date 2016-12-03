hello:
	mkdir -p dst/basic/hello
	python src/basic/00hello.py | gofmt > dst/basic/hello/main.go
	go run dst/basic/hello/main.go

setup:
	pip install -r requirements.txt

