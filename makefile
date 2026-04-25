.PHONY: auth order payment all

## VARIABLES
SERVICES_PATH=.

target: auth order payment internal

auth:
	cd $(SERVICES_PATH)/auth-service && air -c .air.toml

order:
	cd $(SERVICES_PATH)/order-service && air -c .air.toml

payment:
	cd $(SERVICES_PATH)/payment-service && air -c .air.toml

internal:
	cd $(SERVICES_PATH)/internal && air -c .air.toml