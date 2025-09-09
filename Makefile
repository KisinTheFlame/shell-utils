cmd = cmd
build = bin
dest = /usr/local/bin
commands = now kreme kuso work_work
targets = $(patsubst %, $(build)/%, $(commands))
dists = $(patsubst %, $(dest)/%, $(commands))

all: build

build: $(targets)

$(build)/now: cmd/now/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/now

$(build)/kreme: cmd/kreme/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/kreme

$(build)/kuso: cmd/kuso/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/kuso

$(build)/work_work: cmd/work_work/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/work_work

go.sum: go.mod
	go mod tidy

install: build $(dists)

$(dest)/%: $(build)/%
	sudo cp $< $@

.PHONY: clean
clean:
	rm -rf $(build)
