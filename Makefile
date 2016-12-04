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

jsontogo:
	mkdir -p dst/jsontogo/
	python src/jsontogo/01*.py json/github.json | gofmt > dst/jsontogo/github.go
	python src/jsontogo/02*.py json/github.json | gofmt > dst/jsontogo/github2.go
	python src/jsontogo/03*.py json/github.json | gofmt > dst/jsontogo/github3.go
	python src/jsontogo/04*.py json/github.json | gofmt > dst/jsontogo/github4.go
	python src/jsontogo/05*.py json/github.json | gofmt > dst/jsontogo/github5.go
	python src/jsontogo/06*.py json/github.json | gofmt > dst/jsontogo/github6.go
	python src/jsontogo/07*.py json/github.json | gofmt > dst/jsontogo/github7.go
	python src/jsontogo/08*.py json/github.json | gofmt > dst/jsontogo/github8.go
	python src/jsontogo/09*.py json/github.json | gofmt > dst/jsontogo/github9.go

article:
	mkdir -p dst/article
	python src/jsontogo/04*.py json/article.json --name Article | gofmt > dst/article/article4.go
	python src/jsontogo/05*.py json/article.json --name Article | gofmt > dst/article/article5.go
	python src/jsontogo/07*.py json/article.json --name Article | gofmt > dst/article/article7.go

tree:
	mkdir -p dst/tree
	python src/jsontogo/01*.py json/tree.json --name Tree | gofmt > dst/tree/tree.go
	python src/jsontogo/09*.py json/tree.json --name Tree | gofmt > dst/tree/tree9.go

misc:
	rm -f misc/*.output*
	for i in misc/*.py; do python $$i > $$i.output.go; done

setup:
	pip install -r requirements.txt

.PHONY: misc
