build:
	go build -o build/autodnd .

install:
	make -s build
	sudo cp build/autodnd /usr/bin/autodnd
