.PHONY: clean vet test test-cover lint

all: clean internal/sqlite/sqlite3.c .build/reqlite.so .build/reqlite

# pass these flags to linker to suppress missing symbol errors in intermediate artifacts
export CGO_LDFLAGS = -Wl,--unresolved-symbols=ignore-in-object-files
ifeq ($(shell uname -s),Darwin)
	export CGO_LDFLAGS = -Wl,-undefined,dynamic_lookup
endif

test:
	@CGO_LDFLAGS="${CGO_LDFLAGS}" go test -v -tags="libsqlite3,sqlite_json1" ./...

test-cover:
	@CGO_LDFLAGS="${CGO_LDFLAGS}" go test -v -tags="libsqlite3,sqlite_json1" -cover -covermode=count -coverprofile=coverage.out ./...

vet:
	@CGO_LDFLAGS="${CGO_LDFLAGS}" go vet -v -tags="libsqlite3,sqlite_json1" ./...

lint:
	@CGO_LDFLAGS="${CGO_LDFLAGS}" golangci-lint run --build-tags libsqlite3,sqlite_json1

.build/reqlite.so: $(shell find . -type f -name '*.go' -o -name '*.c')
	$(call log, $(CYAN), "building $@")
	@CGO_CFLAGS="-DUSE_LIBSQLITE3" CPATH="${PWD}/internal/sqlite" \
		go build -buildmode=c-shared -o $@ -tags="shared" shared.go
	$(call log, $(GREEN), "built $@")

.build/reqlite: $(shell find . -type f -name '*.go' -o -name '*.c')
	$(call log, $(CYAN), "building $@")
	@CGO_LDFLAGS="${CGO_LDFLAGS}" CGO_CFLAGS="-DUSE_LIBSQLITE3" CPATH="${PWD}/internal/sqlite" \
		go build -o $@ -tags="sqlite_json1,static,!shared" main.go
	$(call log, $(GREEN), "built $@")

# target to download latest sqlite3 amalgamation code
internal/sqlite/sqlite3.c:
	$(call log, $(CYAN), "downloading sqlite3 amalgamation source v3.35.0")
	$(eval SQLITE_DOWNLOAD_DIR = $(shell mktemp -d))
	@curl -sSLo $(SQLITE_DOWNLOAD_DIR)/sqlite3.zip https://www.sqlite.org/2021/sqlite-amalgamation-3350000.zip
	$(call log, $(GREEN), "downloaded sqlite3 amalgamation source v3.35.0")
	$(call log, $(CYAN), "unzipping to $(SQLITE_DOWNLOAD_DIR)")
	@(cd $(SQLITE_DOWNLOAD_DIR) && unzip sqlite3.zip > /dev/null)
	@-rm $(SQLITE_DOWNLOAD_DIR)/sqlite-amalgamation-3350000/shell.c
	$(call log, $(CYAN), "moving to internal/sqlite")
	@mv $(SQLITE_DOWNLOAD_DIR)/sqlite-amalgamation-3350000/* internal/sqlite

clean:
	$(call log, $(YELLOW), "nuking .build/")
	@-rm -rf .build/

# ========================================
# some utility methods

# ASCII color codes that can be used with functions that output to stdout
RED		:= 1;31
GREEN	:= 1;32
ORANGE	:= 1;33
YELLOW	:= 1;33
BLUE	:= 1;34
PURPLE	:= 1;35
CYAN	:= 1;36

# log:
#	print out $2 to stdout using $1 as ASCII color codes
define log
	@printf "\033[$(strip $1)m-- %s\033[0m\n" $2
endef
