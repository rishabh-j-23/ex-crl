CMD=init

run:
	go run ./cmd/main.go ${CMD}
build:
	@go build -o ex-crl ./cmd/main.go

install:
	@echo "Downloading dependencies"
	@go mod tidy
	@echo "[INFO] Building ex-crl..."
	@make build
	@echo "[INFO] Setting executable permissions..."
	@chmod +x ./ex-crl
	@echo "[INFO] Moving binary to ~/.local/bin..."
	@mkdir -p ~/.local/bin
	@cp ./ex-crl ~/.local/bin/
	@sudo mv ./ex-crl /usr/bin/

	@echo "[INFO] Creating dir to ~/ex-crl..."
	@mkdir -p ~/ex-crl

	@echo "[INFO] Installation completed successfully."

uninstall:
	@rm -f ~/.local/bin/ex-crl
	@rm -rf ~/.config/ex-crl
	@echo "ex-crl uninstalled."

test:
	go test -v ./...
