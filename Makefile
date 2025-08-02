UTILS := $(shell find . -mindepth 1 -maxdepth 1 -type d ! -name bin ! -name .vscode ! -name .git) 

BIN_DIR := bin

all: $(UTILS)

$(UTILS):
	@util_name=$(notdir $@); \
	echo "Building $$util_name..."; \
	go build -o $(BIN_DIR)/$$util_name ./$$util_name

clean:
	rm -rf $(BIN_DIR)/*

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: all clean $(UTILS)