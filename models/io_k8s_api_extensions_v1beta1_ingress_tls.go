// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIExtensionsV1beta1IngressTLS IngressTLS describes the transport layer security associated with an Ingress.
//
// swagger:model io.k8s.api.extensions.v1beta1.IngressTLS
type IoK8sAPIExtensionsV1beta1IngressTLS struct {

	// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Hosts []string `json:"hosts"`

	// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	SecretName string `json:"secretName,omitempty"`
}

// Validate validates this io k8s api extensions v1beta1 ingress TLS
func (m *IoK8sAPIExtensionsV1beta1IngressTLS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIExtensionsV1beta1IngressTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIExtensionsV1beta1IngressTLS) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIExtensionsV1beta1IngressTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}