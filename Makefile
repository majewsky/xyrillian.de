# NOTE: need to symlink dl/ to https://dl.xyrillian.de/noises/ for build to work

all: noises/index.html thoughts/index.html res/xyrillian.css

noises/index.html: build/noises/* build/noises/shownotes/*
	go run ./build/noises

thoughts/index.html: build/thoughts/* build/thoughts/posts/*
	go run ./build/thoughts

res/xyrillian.css: build/res/*.scss
	sassc -t compressed -I build/res build/res/main.scss res/xyrillian.css
