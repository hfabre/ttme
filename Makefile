all: run

run:
	go run ./main.go

build-windows:
	CGO_ENABLED=1 GOARCH=amd64 CC=x86_64-w64-mingw32-gcc GOOS=windows go build -o ./build/ttme-windows ./main.go

build-darwin:
	CGO_ENABLED=1 GOARCH=amd64 CC=gcc GOOS=darwin go build -o ./build/ttme-darwin ./main.go

fbuild: build-windows build-darwin
	mkdir -p ./build/assets
	mkdir -p ./packaged_build
	cp ./assets/tilesetpkm.png ./build/assets/tilesetpkm.png
	zip -r ttme ./build
	mv ./ttme.zip ./packaged_build/ttme.zip

fclean:
	rm -rf ./build/*
	rm -rf ./packaged_build/*
