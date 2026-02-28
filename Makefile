# Makefile pour Port-Scanner

# Variables
BINARY_NAME=port-scanner
BINARY_WINDOWS=$(BINARY_NAME).exe
BINARY_LINUX=$(BINARY_NAME)
BINARY_MAC=$(BINARY_NAME)

# Commandes
.PHONY: all build clean test run help

all: clean build

# Compiler pour Windows
build:
	go build -o $(BINARY_WINDOWS) main.go

# Compiler pour plusieurs plateformes
build-all:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_WINDOWS) main.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-mac main.go

# Nettoyer les binaires
clean:
	go clean
	del /Q $(BINARY_WINDOWS) 2>nul || echo Cleaned

# Exécuter les tests
test:
	go test -v ./...

# Lancer le programme
run:
	go run main.go

# Lancer avec des arguments par défaut
run-quick:
	go run main.go --host localhost --quick

# Vérifier le code
lint:
	go fmt ./...
	go vet ./...

# Afficher l'aide
help:
	@echo "Commandes disponibles:"
	@echo "  make build       - Compiler le programme"
	@echo "  make build-all   - Compiler pour toutes les plateformes"
	@echo "  make clean       - Nettoyer les fichiers compiles"
	@echo "  make test        - Lancer les tests"
	@echo "  make run         - Executer le programme"
	@echo "  make run-quick   - Executer en mode quick"
	@echo "  make lint        - Verifier le code"
	@echo "  make help        - Afficher cette aide"
