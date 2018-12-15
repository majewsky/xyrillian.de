all: noises/index.html res/xyrillian.css

noises/index.html: build/noises/*
	go run ./build/noises

res/xyrillian.css: build/res/*.scss
	sassc -t compressed -I build/res build/res/main.scss res/xyrillian.css
