#! /usr/bin/make
#
# Makefile for goa-swagger
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
# - "build" builds the docker container with the latest master from github
# - "run" run service in docker
# - "deploy" deploys the latest build to GKE
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} ./...)
VERSION=v5
IMAGE=gcr.io/goa-swagger/service-node:$(VERSION)
DEPEND=golang.org/x/tools/cmd/cover golang.org/x/tools/cmd/goimports \
	github.com/golang/lint/golint github.com/onsi/gomega \
	github.com/onsi/ginkgo github.com/onsi/ginkgo/ginkgo

.PHONY: build deploy

all: depend lint test build deploy

depend:
	@go get $(DEPEND)

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -v app | grep -v "should not use dot imports" | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi

test:
	@ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race -skipPackage vendor

build:
	@docker build -t $(IMAGE) .

gke:
	@gcloud container clusters create goa-swagger --num-nodes 2 --machine-type n1-standard-1
	@gcloud container clusters get-credentials goa-swagger
	@kubectl run service-node --image=$(IMAGE) --port=8080
	@kubectl expose rc service-node --type="LoadBalancer"
	@kubectl autoscale rc service-node --min=2 --max=5

run:
	docker run --rm --publish 8080:8080 $(IMAGE)

deploy:
	@gcloud docker push $(IMAGE)
	@kubectl rolling-update --update-period=10ms service-node --image=$(IMAGE)
