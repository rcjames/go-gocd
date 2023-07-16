#!/bin/bash

set -e

function wait-for-gocd-start() {
  # Add an initial delay for the server to start responding
  sleep 5

  while true; do
    response=$(curl -s http://localhost:8153/go/api/version | grep "GoCD server is starting" || true)
    if [[ "0" == $(echo -n "$response" | wc -c) ]]; then
      echo "GoCD server has started"
      break
    fi
    echo "$(date +'%F %T') Waiting for GoCD server to start. Response: $response"
    sleep 1
  done
}

function escape() {
  echo "$(date +'%F %T') Shutting down GoCD server"
  docker stop gocd
}

DOCKER_REGISTRY_PLUGIN_VERSION=1.3.1-485
DOCKER_REGISTRY_PLUGIN_FILE=docker-registry-artifact-plugin-${DOCKER_REGISTRY_PLUGIN_VERSION}.jar

if [[ ! -f "${PWD}/test-plugins/${DOCKER_REGISTRY_PLUGIN_FILE}" ]]; then
  echo "$(date +'%F %T') Downloading plugins"
  wget -O ${PWD}/test-plugins/${DOCKER_REGISTRY_PLUGIN_FILE} https://github.com/gocd/docker-registry-artifact-plugin/releases/download/v${DOCKER_REGISTRY_PLUGIN_VERSION}/${DOCKER_REGISTRY_PLUGIN_FILE}
fi

echo "$(date +'%F %T') Starting GoCD server"
bash ./start-gocd-server.sh
trap escape EXIT

echo "$(date +'%F %T') Waiting for GoCD server to start"
wait-for-gocd-start

echo $(date +'%F %T') Running integration tests
go test -tags=integration