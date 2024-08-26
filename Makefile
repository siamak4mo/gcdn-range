GO=go
CC=$(GO) build
CFLAGS=-compiler gccgo
BIN=cdn-range.out
OUT=-o $(BIN)

# test
TEST_1=./providers
TEST_2=./cmd

# main program
MAIN=./main.go
DEPS=$(shell find . -name '*.go' | grep -v '_test.go')


# with gccgo
all: $(BIN)

$(BIN): $(MAIN) $(DEPS)
	$(CC) $(CFLAGS) $(OUT) $(MAIN)


# without gccgo
no_gcc: $(MAIN) $(DEPS)
	$(CC) $(OUT) $(MAIN)

no_gcc_run:
	$(GO) run $(MAIN)


# test
test:
	make -C $(TEST_1) test
	make -C $(TEST_2) test
