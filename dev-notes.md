# Notes about development

All API code is based on go-swagger. There are a few issues to be taken into account.

The command to re-generate the code from the spec should be :
`swagger generate server -f swagger.json --skip-models`

Models are generated but were modified afterwards. If you want to modify them
you need to do it manually or re-do the changes that were done.

#### POST requests
the mime type "application/json" has to be added manually to
the swagger spec of endpoints "produces" part, because go-swagger can't seem
to find any json producers when given a "*/*" wildcard as mime type.

#### Generated model
go-swagger mixes up types between direct and indirect objects. For example,
in the Pod struct, the Metadata field is indirect. However in the Metadata
struct, deletionTimestamp is direct enventhough they are defined the same way
which causes issues at serialisation when deletionTimestamp isn't supposed to
have a value and should be ignored. The scheduler ends up confused and fails to
schedule. 

The reason for this is that go-swagger treats custom objects defined in the
swagger.json (Metadata, Spec, Status...) as indirect types, and standard types
(string, int...) as direct and eventhough IoK8sApimachineryPkgApisMetaV1Time
(the type defined for the time) is custom, its underlying type is a string.

Therefore to avoid that, some fields ahve been modified manually. Models
shouldn't be re-generated.
