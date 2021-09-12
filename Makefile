.PHONY: all
all:
	@echo "Try 'make readme'"

.PHONY: readme
readme:
	godoc -templates=./readme-template/ "github.com/skratchdot/open-golang/open" > ./README.md
.PHONY: test
test:
	go test github.com/zwjzxh520/open-golang/open