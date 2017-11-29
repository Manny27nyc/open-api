// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SiteSetupAllOf1 site setup all of1
// swagger:model siteSetupAllOf1
type SiteSetupAllOf1 struct {

	// repo
	Repo *RepoSetup `json:"repo,omitempty"`
}

// Validate validates this site setup all of1
func (m *SiteSetupAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteSetupAllOf1) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteSetupAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteSetupAllOf1) UnmarshalBinary(b []byte) error {
	var res SiteSetupAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
