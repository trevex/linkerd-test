#!/usr/bin/env bash
docker build -t trevex/testapp:$1 .
docker push trevex/testapp:$1
