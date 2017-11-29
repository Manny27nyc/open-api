// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SiteBuildSettings site build settings
// swagger:model siteBuildSettings
type SiteBuildSettings struct {

	// allowed branches
	AllowedBranches []string `json:"allowed_branches"`

	// branch
	Branch string `json:"branch,omitempty"`

	// cmd
	Cmd string `json:"cmd,omitempty"`

	// dir
	Dir string `json:"dir,omitempty"`
}

// Validate validates this site build settings
func (m *SiteBuildSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedBranches(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteBuildSettings) validateAllowedBranches(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedBranches) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteBuildSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteBuildSettings) UnmarshalBinary(b []byte) error {
	var res SiteBuildSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
