.PHONY: test

test:
	@go test ./...

fmt:
	@go fmt ./...

chart-doc-gen:
	@docker run --rm --volume "$(shell pwd)/helm/github-rate-limit-exporter:/helm-docs" -u $(shell id -u) jnorwood/helm-docs:latest