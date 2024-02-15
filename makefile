# ==============================================================================
# Build, add permissions, move to executables

build:
	go build -o termiai src/main.go

add-permissions:
	chmod +x termiai

move:
	sudo mv termiai /usr/local/bin/termiai

deploy: build add-permissions move