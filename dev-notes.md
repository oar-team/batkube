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

#### Nullability of time.Time fields
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
