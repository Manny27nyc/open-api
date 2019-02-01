// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteServiceInstanceParams creates a new DeleteServiceInstanceParams object
// with the default values initialized.
func NewDeleteServiceInstanceParams() *DeleteServiceInstanceParams {
	var ()
	return &DeleteServiceInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceInstanceParamsWithTimeout creates a new DeleteServiceInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServiceInstanceParamsWithTimeout(timeout time.Duration) *DeleteServiceInstanceParams {
	var ()
	return &DeleteServiceInstanceParams{

		timeout: timeout,
	}
}

// NewDeleteServiceInstanceParamsWithContext creates a new DeleteServiceInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServiceInstanceParamsWithContext(ctx context.Context) *DeleteServiceInstanceParams {
	var ()
	return &DeleteServiceInstanceParams{

		Context: ctx,
	}
}

// NewDeleteServiceInstanceParamsWithHTTPClient creates a new DeleteServiceInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServiceInstanceParamsWithHTTPClient(client *http.Client) *DeleteServiceInstanceParams {
	var ()
	return &DeleteServiceInstanceParams{
		HTTPClient: client,
	}
}

/*DeleteServiceInstanceParams contains all the parameters to send to the API endpoint
for the delete service instance operation typically these are written to a http.Request
*/
type DeleteServiceInstanceParams struct {

	/*Addon*/
	Addon string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete service instance params
func (o *DeleteServiceInstanceParams) WithTimeout(timeout time.Duration) *DeleteServiceInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service instance params
func (o *DeleteServiceInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service instance params
func (o *DeleteServiceInstanceParams) WithContext(ctx context.Context) *DeleteServiceInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service instance params
func (o *DeleteServiceInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service instance params
func (o *DeleteServiceInstanceParams) WithHTTPClient(client *http.Client) *DeleteServiceInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service instance params
func (o *DeleteServiceInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddon adds the addon to the delete service instance params
func (o *DeleteServiceInstanceParams) WithAddon(addon string) *DeleteServiceInstanceParams {
	o.SetAddon(addon)
	return o
}

// SetAddon adds the addon to the delete service instance params
func (o *DeleteServiceInstanceParams) SetAddon(addon string) {
	o.Addon = addon
}

// WithSiteID adds the siteID to the delete service instance params
func (o *DeleteServiceInstanceParams) WithSiteID(siteID string) *DeleteServiceInstanceParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete service instance params
func (o *DeleteServiceInstanceParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addon
	if err := r.SetPathParam("addon", o.Addon); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
