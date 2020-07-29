#!/bin/sh

swagger generate server -f swagger.json --flag-strategy=go-flags

find ./models -type f ! -name "io_k8s_apimachinery_pkg_apis_meta_v1_time.go" | xargs sed -i 's/IoK8sApimachineryPkgApisMetaV1Time/\*IoK8sApimachineryPkgApisMetaV1Time/g'
