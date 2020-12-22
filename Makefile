DEPS:=filcrypto.h filcrypto.pc libfilcrypto.a

all: $(DEPS)
.PHONY: all

# Create a file so that parallel make doesn't call `./install-filcrypto` for
# each of the deps
$(DEPS): .install-filcrypto  ;

.install-filcrypto:
	./install-filcrypto
	@touch $@

clean:
	rm -rf $(DEPS) .install-filcrypto
.PHONY: clean

go-lint: $(DEPS)
	golangci-lint run -v --concurrency 2 --new-from-rev origin/master
.PHONY: go-lint

shellcheck:
	shellcheck install-filcrypto

lint: shellcheck go-lint
