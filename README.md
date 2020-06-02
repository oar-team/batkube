# Batkube

Kubernetes cluster simulator for batch processing simulations.

This acts as an interface between the infrastructure simulator [Batsim](https://github.com/oar-team/batsim) and any Kubernetes scheduler.

## Usage
- Clone any Kubernetes scheduler and adapt it using the [batsky-go installer](https://github.com/oar-team/batsky-go-installer)
- Launch the scheduler
- Launch Batkube : `go run cmd/kubernetes-server/main.go --tls-certificate certs/server.crt --tls-key certs/server.key --tls-ca certs/ca.crt --tls-port 6443`
- `mkdir /tmp/expe-out`
- Wait for the Batkube API to launch (it should say `Serving kubernetes at https://127.0.0.1:6443`)
- Launch Batsim : `batsim -p examples/platforms/platform_graphene_16nodes.xml -w examples/workloads/test_3_delay10_job.json -e "/tmp/expe-out/out"`

Notes :
- The original `swagger.json` used to generate the API is included in this repo. It is based on [Kubernetes 1.18 spec](https://github.com/kubernetes/kubernetes/blob/release-1.18/api/openapi-spec/swagger.json) and slightly modified for compatibility.
- `kubeconfig.yaml` provides a base for out-of-cluster configuration for your scheduler.
