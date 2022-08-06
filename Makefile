.PHONY: *
init config page api wire test build all container help:
	find app -mindepth 1 -maxdepth 1 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) $@'

.DEFAULT_GOAL := help