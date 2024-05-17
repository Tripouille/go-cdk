build:
	@GOOS=linux GOARCH=amd64 go build -o bootstrap
	@zip function.zip bootstrap
	@rm bootstrap

clean:
	@rm function.zip

.PHONY: build clean