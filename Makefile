hello:
	mkdir -p dst/basic/hello
	python src/basic/00hello.py | gofmt > dst/basic/hello/main.go
	go run dst/basic/hello/main.go

cross:
	mkdir -p dst/basic/cross
	python src/basic/01cross*.py | gofmt > dst/basic/cross/cross2.go
	python src/basic/02cross*.py | gofmt > dst/basic/cross/cross5.go

fizzbuzz:
	mkdir -p dst/basic/fizzbuzz
	python src/basic/03*.py | gofmt > dst/basic/fizzbuzz/main.go
	go run dst/basic/fizzbuzz/main.go

setup:
	pip install -r requirements.txt
