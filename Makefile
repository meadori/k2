.PHONY: all
all: bin/k2

src := $(shell find . -type f -name '*.go')

bin/%: $(src)
	@mkdir -p $(dir $@)
	go build -o $@ ./cmd/$(@F)

.PHONY: fmt
fmt:
	go fmt $(shell dirname $(src)) 

.PHONY: clean
clean:
	rm -rf bin
