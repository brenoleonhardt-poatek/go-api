
.PHONY: all docs

start:
	air

docs:
	swag fmt
	swag init --parseDependency true --parseInternal true
