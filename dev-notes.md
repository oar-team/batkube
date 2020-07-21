# Notes about development

All API code is based on go-swagger. There are a few issues to be taken into account.

The command to re-generate the code from the spec should be :
`swagger generate server -f swagger.json --skip-models`

Models are generated but were modified afterwards. If you want to modify them
you need to do it manually or re-do the changes that were done.

## POST requests

The mime type "application/json" has to be added manually to
the swagger spec of endpoints "produces" part, because go-swagger can't seem
to find any json producers when given a "*/*" wildcard as mime type.

## Nullability of time.Time fields

There is an issue with object's Metadata timestamp fields. Since the timestamps
are not considered to be structs, they are of direct type and therefore not
checked against zero values (because time.Time is a struct and does not have a
zero value that json understands), which causes issues because
deletionTimestamp is not supposed to be set when the object is not deleted.

A quick fix to this is to indirect Metadata.DeletionTimestamp so it has a zero
value (nil).

Removing the "format":"date-time" in the spec does not work, as we end up with
the wrong string format for the timestamp.
Some ways to indicate nullability with go-swagger are specified here : https://goswagger.io/use/models/schemas.html#nullability

A more durable fix would be to force go-swagger to generate a struct by
changing the swagger spec.

## Other general notes

- The original `swagger.json` used to generate the API is included in this
    repo. It is based on [Kubernetes 1.18
    spec](https://github.com/kubernetes/kubernetes/blob/release-1.18/api/openapi-spec/swagger.json)
    and slightly modified for compatibility. These modifications consist of
    adding "application/yaml" to the mime types of implemented POST and PATCH
    requests as described above.
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
- Also, batsky-go needs to be fixed. (The bad timer issue mentioned in the README).
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
- Batkube only supports one workload.
