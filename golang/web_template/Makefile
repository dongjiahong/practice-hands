GOCMD = go
GOBUILD = ${GOCMD} build
GOBIN = web
DOCS = docs

all: serve

serve: docs
	swag init
	${GOBUILD}
	./${GOBIN}

docs:
	exist = $(shell if [ -f $(DOCS) ]; then echo "exist"; else echo "notexist"; fi;)
	ifeq ("$(exist)", "notexist")
	$(shell mkdir $(DOCS))
	endif
