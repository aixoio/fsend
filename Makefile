GO = go
PRODARGS = -ldflags "-s -w"

dev:
	$(GO) build

prod:
	$(GO) build $(PRODARGS)

.phony: dev prod

.DEFAULT_GOAL=prod
