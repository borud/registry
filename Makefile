all: lint gen

lint:
	@buf lint

gen:
	@buf generate

breaking:
	@buf breaking --against "https://github.com/borud/registry.git#branch=main"

clean:
	@rm -rf pkg

help:
	@echo ""
	@echo "  lint     - lint the proto files"
	@echo "  gen      - perform code generation"
	@echo "  breaking - check if you have made any breaking changes (check against main branch)"
	@echo "  clean    - remove generated files"
	@echo ""
	@echo "  default target is all, which runs lint and gen"
	@echo ""