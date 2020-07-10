# Batkube

Batkube is an interface between the
[Batsim](https://github.com/oar-team/batsim) infrastructure simulator and any
Kubernetes scheduler written in Go, in order to run basic Kubernets cluster simulations.

## Usage

- Clone any Kubernetes scheduler and adapt it using the [batsky-go
    installer](https://github.com/oar-team/batsky-go-installer)
- Launch the scheduler, specify where the master is (at `http://localhost:8001`)
- Launch Batkube `go run cmd/kubernetes-server/main.go --scheme=http --port 8001`
- Create a directory for Batsim output 'ex : `mkdir /tmp/expe-out`
- Wait for the Batkube API to launch (it should say `Serving kubernetes at http://127.0.0.1:8001`)
- Launch Batsim with your desired platform and workload (ex : `batsim -p
    examples/platforms/platform_graphene_16nodes.xml -w
    examples/workloads/spaced_150_delay200.json -e "/tmp/expe-out/out"`)

Notes :

- Internally, resources are not cleaned up perdiocically because there are not
    implemented in a thread safe way. This causes memory issues with long
    running simulations
- [batsky-go](https://github.com/oar-team/batsky-go) suffers from an unknown
    technical issue which leads it to panic with error `bad timer`. In that
    case, just re-launch the simulation. It does not happen very often.
- If you desire to change the parameters of the simulation, change them
    directly in the code, on the first lines of `pkg/broker/broker.go`, and
    then re-compile the code (or run your simulation with `go run` like in the
    aforementioned example)

Technial notes for further development :

- The original `swagger.json` used to generate the API is included in this
    repo. It is based on [Kubernetes 1.18
    spec](https://github.com/kubernetes/kubernetes/blob/release-1.18/api/openapi-spec/swagger.json)
    and slightly modified for compatibility. These modifications consist of
    adding "application/yaml" to the mime types of implemented POST and PATCH
    requests.
- Use `regenerate-api.sh` when you want to re-generate the code after changes
    to the swagger spec. This script takes care of an issue that exists with
    time variables and their zero value by indirecting them. (Otherwise these
    are direct values and their zero value ends up being a time with value
    zero, where it needs to be null.). A better solution would be to edit the
    swagger to indicate nullability of those fields, however that is not
    possible with OpenAPI 2.0 and the go-swagger have [no intent of upgrading
    to OpenAPI 3.0 for
    now](https://github.com/go-swagger/go-swagger/issues/1122). The go-swagger
    documentation offers clues on how to do that for now but I have not been
    able to make it work.
- Enable debugging logs for the API by setting the `debug` variable to "true"
    in `restapi/configure_kubernetes.go`. Setting the environment variable
    `DEBUG` to a non zero value also works (warning : very verbose)
- Kubernetes resources are stored in memory and are not thread safe causing
    memory issues as explained earlier. This needs to be fixed.
- Also, batsky-go needs to be fixed. (The bad timer issue).
- This code could greatly benefit from [Go
    generics](https://blog.golang.org/generics-next-step) as it uses a lot of
    reflection to go through resources and filter them. Consider either
    re-implementing the structs to use interfaces and interface methods, or use
    generics to have a cleaner code.
- This code has no implementation of a proper deep copy function. It uses json
    coding and decoding as a quick and dirty way of copying objects to create
    watch events.
- For a better user experience, flags need to be implemented so we don't have
    to go through the code and re-compile everytime the user wants to change a
    sumulation parameter.
- This code uses a simple hack to update resourceVersions : everytime a request
    is made with a specified resourceVersion, every object's resourceVersion is
    updated to this value. It works very well and doesn't need to be changed,
    for now. Maybe consider propely handling resourceVersions in the future.
