.PHONY: help chain-init chain-up chain-down faucet-up faucet-down keyring-up keyring-down mokr-up mokr-down web-init web-up web-down services-up services-down kill-port

BLOCKCHAIN_DIR := blockchain
WEB_DIR := web
PORT_KILL_CMD := lsof -ti

define kill_port
	@$(PORT_KILL_CMD):$(1) | xargs kill
endef

define colorecho
      @tput setaf $2
      @echo $1
      @tput sgr0
endef

help:
	$(call colorecho,"Fusion Makefile - 2024",6)
	@echo "Available commands:"
	@echo
	@echo "  make chain-init         - Initialize & start the Fusion chain"
	@echo "  make chain-up           - Start the Fusion chain"
	@echo "  make chain-down         - Stop the Fusion chain"
	@echo
	@echo "  make faucet-up          - Start the Faucet"
	@echo "  make faucet-down        - Stop the Faucet"
	@echo "  make keyring-up         - Start the Keyring"
	@echo "  make mokr-up            - Start the Mock Keyring"
	@echo
	@echo "  make web-init           - Initialize & start the Web Service"
	@echo "  make web-up             - Start the Web Service"
	@echo "  make web-down           - Stop the Web Service"
	@echo
	@echo "  make services-up        - Start the Whole Stack"
	@echo "  make services-down      - Stop the Whole Stack"

chain-init:
	@cd $(BLOCKCHAIN_DIR) && ./init.sh

chain-up:
	@cd $(BLOCKCHAIN_DIR) && fusiond start

chain-down:
	@$(call kill_port,26656)

faucet-up:
	@cd $(BLOCKCHAIN_DIR) && go run cmd/faucet/faucet.go

faucet-down:
	@$(call kill_port,8000)

keyring-up:
	@cd keyring/cmd/mpc-relayer && go run main.go

mokr-up:
	@cd mokr && go run .

web-init:
	@cd $(WEB_DIR) && pnpm install && pnpm run dev

web-up:
	@cd $(WEB_DIR) && pnpm run dev

web-down:
	@$(call kill_port,5173) # Adjust the port if your Web Service runs on a different one

