#!/bin/bash

docker run -d --rm -p 8153:8153 --name gocd -v ${PWD}/test-plugins:/tmp/godata/plugins/external gocd/gocd-server:v23.1.0
docker exec gocd mkdir -p /godata/plugins
docker exec gocd cp -r /tmp/godata/plugins/external/ /godata/plugins/