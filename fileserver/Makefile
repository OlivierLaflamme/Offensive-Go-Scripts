TARGET=./build
ARCHS=amd64 386
LDFLAGS="-s -w"

.DEFAULT_GOAL := all

all: clean windows linux darwin

deps:
	$(GOGET) github.com/gin-gonic/gin
	
windows:
	@mkdir -p ${TARGET} ; \
	for GOARCH in ${ARCHS}; do \
		echo "Building for windows $${GOARCH} ..." ; \
		GOOS=windows GOARCH=$${GOARCH} go build -ldflags=${LDFLAGS} -trimpath -o ${TARGET}/filseserver-windows-$${GOARCH}.exe ; \
	done;

linux:
	@mkdir -p ${TARGET} ; \
	for GOARCH in ${ARCHS}; do \
		echo "Building for linux $${GOARCH} ..." ; \
		GOOS=linux GOARCH=$${GOARCH} go build -ldflags=${LDFLAGS} -trimpath -o ${TARGET}/filseserver-linux-$${GOARCH} ; \
	done;

darwin:
	@mkdir -p ${TARGET} ; \
	for GOARCH in ${ARCHS}; do \
		echo "Building for darwin $${GOARCH} ..." ; \
		GOOS=darwin GOARCH=$${GOARCH} go build -ldflags=${LDFLAGS} -trimpath -o ${TARGET}/filseserver-darwin-$${GOARCH} ; \
	done;

clean:
	@rm -rf ${TARGET}/*
