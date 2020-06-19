#!/bin/sh

$1 --master http://localhost:8001 --kube-api-content-type=application/json --leader-elect=false --scheduler-name=random-scheduler
