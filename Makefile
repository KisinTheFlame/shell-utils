cmd = cmd
build = bin
dest = ~/.local/bin
commands = $(notdir $(patsubst %/,%,$(dir $(wildcard $(cmd)/*/main.go))))
targets = $(patsubst %, $(build)/%, $(commands))
dists = $(patsubst %, $(dest)/%, $(commands))

all: build

build: $(targets)

$(build)/%: $(cmd)/%/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./$(cmd)/$*

go.sum: go.mod
	go mod tidy

install: build $(dists)

$(dest)/%: $(build)/%
	cp $< $@

.PHONY: clean
clean:
	rm -rf $(build)
