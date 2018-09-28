all: noises/index.html

noises/index.html: build/noises/data.yaml build/noises/*.go build/noises/*.tpl
	go run ./build/noises
