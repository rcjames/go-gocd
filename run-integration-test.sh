#!/bin/bash

set -e

if [[ -z "$(ls ${PWD}/test-plugins | grep docker-registry)" ]]; then
  # TODO - Download plugin
  echo "Please download docker registry plugin to test-plugins dir"
  exit 1
fi

docker run -d --rm -p 8153:8153 --name gocd -v ${PWD}/test-plugins:/tmp/godata/plugins/external gocd/gocd-server:v23.1.0
docker exec gocd mkdir -p /godata/plugins
docker exec gocd cp -r /tmp/godata/plugins/external/ /godata/plugins/

# TODO - Health check
sleep 30

go test -tags=integration