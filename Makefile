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

# swagger
swagger_setup:
	pip install -r requirements.txt
	go get -v github.com/go-swagger/go-swagger/cmd/swagger
	go get -v github.com/go-openapi/validate
	go get -v github.com/go-openapi/swag
	go get -v github.com/podhmo/go-structjson/cmd/go-structjson
	go get -v github.com/podhmo/go-structjson/cmd/go-funcjson
	go get -v golang.org/x/tools/cmd/goimports
	go get -v github.com/davecgh/go-spew/spew
	npm install

swagger_serve:
	./node_modules/.bin/redocup	yaml/github-swagger.yaml

swagger_fetch:
	mkdir -p yaml
	wget https://api.apis.guru/v2/specs/github.com/v3/swagger.yaml -O yaml/github-swagger.yaml
	gsed -i "s/type: *'null'/type: object/g; s/'+1':/'plus1':/g; s/'-1':/'minus1':/" yaml/github-swagger.yaml

swagger_src:
	mkdir -p src/github
	python src/jsontogo/jsontogo.py json/github-get-authenticated-user.json --package github --name AuthenticatedUser | gofmt > src/github/authenticated_user.go

swagger_gen:
	rm -rf dst/swagger/gen
	mkdir -p dst/swagger/gen
	swagger generate model -f yaml/github-swagger.yaml --target dst/swagger/gen --model-package def
	goimports -w dst/swagger/gen/def/*.go

swagger_extract:
	mkdir -p dst/swagger/convert
	mkdir -p json/extracted
	go-structjson -target src/github > json/extracted/src.json
	go-structjson -target dst/swagger/gen/def > json/extracted/dst.json
	go-funcjson -target dst/swagger/convert > json/extracted/convert.json || echo '{}' > json/extracted/convert.json

swagger_convert:
	mkdir -p dst/swagger/convert
	python src/convert.py --logger=DEBUG --src json/extracted/src.json --dst json/extracted/dst.json --override json/extracted/convert.json > dst/swagger/convert/autogen.go
	gofmt -w dst/swagger/convert/autogen.go

swagger_run1:
	cat json/github-get-authenticated-user.json | go run dst/swagger/main/spew/*.go

swagger_run2:
	cat json/github-get-authenticated-user.json | go run dst/swagger/main/json/*.go

.PHONY: misc
