# Batkube

Batkube is an interface between the
[Batsim](https://github.com/oar-team/batsim) infrastructure simulator and any
Kubernetes scheduler written in Go, in order to run basic Kubernets cluster simulations.

## Usage

- Clone any Kubernetes scheduler and adapt it using the [batsky-go
    installer](https://github.com/oar-team/batsky-go-installer)
- Launch the scheduler, give it `kubeconfig.yaml` as input for it to find the
    simulated cluster.
- Launch Batkube `go run cmd/kubernetes-server/main.go --scheme=http --port 8001 [OPTIONS]`
- Create a directory for Batsim output (ex : `mkdir /tmp/expe-out`)
- Wait for the Batkube API to launch (it should say `Serving kubernetes at http://127.0.0.1:8001`)
- Launch Batsim with your desired platform and workload (ex : `batsim -p
    examples/platforms/platform_graphene_16nodes.xml -w
    examples/workloads/spaced_150_delay200.json -e "/tmp/expe-out/out"`)

Notes :

- Internally, resources are not cleaned up periodically because there are not
    implemented in a thread safe way. This mays cause memory issues with long
    running simulations.
- [batsky-go](https://github.com/oar-team/batsky-go) suffers from an unknown
    technical issue which leads it to panic with error `bad timer`. In that
    case, just re-launch the simulation. It does not happen very often.

Please read the [dev notes](dev-notes.md) if you intend to take over this
project for further development.
