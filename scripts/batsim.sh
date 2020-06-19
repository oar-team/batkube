#!/bin/sh

mkdir -p /tmp/expe-out

batsim -p examples/platforms/platform_graphene_16nodes.xml -w examples/workloads/test_3_delay10_job.json -e "/tmp/expe-out/out"
