#!/bin/sh

BIN=$1
shift
$BIN --master http://localhost:8001 --kube-api-content-type=application/json --leader-elect=false --scheduler-name=random-scheduler $@
