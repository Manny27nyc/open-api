// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ServiceInstance service instance
// swagger:model serviceInstance
type ServiceInstance struct {

	// auth url
	AuthURL string `json:"auth_url,omitempty"`

	// config
	Config interface{} `json:"config,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// env
	Env interface{} `json:"env,omitempty"`

	// external attributes
	ExternalAttributes interface{} `json:"external_attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// service name
	ServiceName string `json:"service_name,omitempty"`

	// service path
	ServicePath string `json:"service_path,omitempty"`

	// service slug
	ServiceSlug string `json:"service_slug,omitempty"`

	// snippets
	Snippets []interface{} `json:"snippets"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this service instance
func (m *ServiceInstance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstance) UnmarshalBinary(b []byte) error {
	var res ServiceInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
