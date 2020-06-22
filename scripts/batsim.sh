#!/bin/sh

mkdir /expe-out

batsim -p examples/platforms/platform_graphene_16nodes.xml -w examples/workloads/test_30_delay10_job.json -e "expe-out/out" --enable-compute-sharing
