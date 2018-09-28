# NOTE: need to symlink dl/ to https://dl.xyrillian.de/noises/ for build to work

all: noises/index.html

noises/index.html: build/noises/data.yaml build/noises/*.go build/noises/*.tpl
	go run ./build/noises
