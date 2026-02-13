cmd = cmd
build = bin
dest = ~/.local/bin
commands = health_server kreme kuso now work_work
targets = $(patsubst %, $(build)/%, $(commands))
dists = $(patsubst %, $(dest)/%, $(commands))

all: build

build: $(targets)

$(build)/health_server: cmd/health_server/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/health_server

$(build)/kreme: cmd/kreme/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/kreme

$(build)/kuso: cmd/kuso/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/kuso

$(build)/now: cmd/now/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/now

$(build)/work_work: cmd/work_work/main.go go.mod go.sum
	@mkdir -p $(build)
	go build -o $@ ./cmd/work_work

go.sum: go.mod
	go mod tidy

install: build $(dists)

$(dest)/%: $(build)/%
	cp $< $@

.PHONY: clean
clean:
	rm -rf $(build)
