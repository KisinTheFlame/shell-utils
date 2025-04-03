cmd = src/bin
build = target/release
dest = /usr/local/bin
commands = $(patsubst $(cmd)/%.rs, $(cmd)/% ,$(wildcard $(cmd)/*))
targets = $(patsubst $(cmd)/%, $(build)/%, $(commands))
dists = $(patsubst $(cmd)/%, $(dest)/%, $(commands))

all: build

build:
	cargo +nightly fmt
	cargo build --release

install: build $(dists)

$(dest)/%: $(build)/%
	sudo cp $< $@

.PHONY: clean
clean:
	cargo clean
