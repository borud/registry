all: gen

gen:
	@buf generate

clean:
	@rm -rf pkg