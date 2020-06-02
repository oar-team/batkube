// Code generated by go-swagger; DO NOT EDIT.

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// ReplaceCertificatesV1beta1CertificateSigningRequestURL generates an URL for the replace certificates v1beta1 certificate signing request operation
type ReplaceCertificatesV1beta1CertificateSigningRequestURL struct {
	Name string

	DryRun       *string
	FieldManager *string
	Pretty       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) WithBasePath(bp string) *ReplaceCertificatesV1beta1CertificateSigningRequestURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on ReplaceCertificatesV1beta1CertificateSigningRequestURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var dryRunQ string
	if o.DryRun != nil {
		dryRunQ = *o.DryRun
	}
	if dryRunQ != "" {
		qs.Set("dryRun", dryRunQ)
	}

	var fieldManagerQ string
	if o.FieldManager != nil {
		fieldManagerQ = *o.FieldManager
	}
	if fieldManagerQ != "" {
		qs.Set("fieldManager", fieldManagerQ)
	}

	var prettyQ string
	if o.Pretty != nil {
		prettyQ = *o.Pretty
	}
	if prettyQ != "" {
		qs.Set("pretty", prettyQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ReplaceCertificatesV1beta1CertificateSigningRequestURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ReplaceCertificatesV1beta1CertificateSigningRequestURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
