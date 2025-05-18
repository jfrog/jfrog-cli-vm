.PHONY: build install uninstall clean bootstrap

JFVM_BIN := jfvm
SHIM_BIN := jf
BIN_DIR := $(HOME)/.jfvm/bin

build:
	@echo "🔧 Building jfvm CLI..."
	go build -o $(JFVM_BIN) .
	@echo "🔧 Building jf shim..."
	cd shim && go build -o $(SHIM_BIN) .

install: build
	@echo "📂 Creating bin directory: $(BIN_DIR)"
	mkdir -p $(BIN_DIR)
	@echo "📥 Installing binaries to $(BIN_DIR)"
	cp $(JFVM_BIN) $(BIN_DIR)/
	cp shim/$(SHIM_BIN) $(BIN_DIR)/
	@echo "✅ Binaries installed."

bootstrap: install
	@echo "🔁 Checking shell config for PATH..."
	@grep -q '.jfvm/bin' ~/.bashrc 2>/dev/null || echo 'export PATH="$$HOME/.jfvm/bin:$$PATH"' >> ~/.bashrc
	@grep -q '.jfvm/bin' ~/.zshrc 2>/dev/null || echo 'export PATH="$$HOME/.jfvm/bin:$$PATH"' >> ~/.zshrc
	@grep -q '.jfvm/bin' ~/.profile 2>/dev/null || echo 'export PATH="$$HOME/.jfvm/bin:$$PATH"' >> ~/.profile
	@echo "✅ PATH updated in shell config. Run 'source ~/.bashrc' or 'source ~/.zshrc' to apply."

uninstall:
	@echo "🗑️ Removing installed binaries..."
	rm -f $(BIN_DIR)/$(JFVM_BIN) $(BIN_DIR)/$(SHIM_BIN)
	@echo "✅ Uninstalled."

clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f $(JFVM_BIN)
	cd shim && rm -f $(SHIM_BIN)