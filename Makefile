build:
	mkdir -p functions
	go get ./...
	go build -o functions ./...
	hugo --gc --minify --source src