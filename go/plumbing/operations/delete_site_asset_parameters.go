// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSiteAssetParams creates a new DeleteSiteAssetParams object
// with the default values initialized.
func NewDeleteSiteAssetParams() *DeleteSiteAssetParams {
	var ()
	return &DeleteSiteAssetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSiteAssetParamsWithTimeout creates a new DeleteSiteAssetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSiteAssetParamsWithTimeout(timeout time.Duration) *DeleteSiteAssetParams {
	var ()
	return &DeleteSiteAssetParams{

		timeout: timeout,
	}
}

// NewDeleteSiteAssetParamsWithContext creates a new DeleteSiteAssetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSiteAssetParamsWithContext(ctx context.Context) *DeleteSiteAssetParams {
	var ()
	return &DeleteSiteAssetParams{

		Context: ctx,
	}
}

// NewDeleteSiteAssetParamsWithHTTPClient creates a new DeleteSiteAssetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSiteAssetParamsWithHTTPClient(client *http.Client) *DeleteSiteAssetParams {
	var ()
	return &DeleteSiteAssetParams{
		HTTPClient: client,
	}
}

/*DeleteSiteAssetParams contains all the parameters to send to the API endpoint
for the delete site asset operation typically these are written to a http.Request
*/
type DeleteSiteAssetParams struct {

	/*AssetID*/
	AssetID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete site asset params
func (o *DeleteSiteAssetParams) WithTimeout(timeout time.Duration) *DeleteSiteAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete site asset params
func (o *DeleteSiteAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete site asset params
func (o *DeleteSiteAssetParams) WithContext(ctx context.Context) *DeleteSiteAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete site asset params
func (o *DeleteSiteAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete site asset params
func (o *DeleteSiteAssetParams) WithHTTPClient(client *http.Client) *DeleteSiteAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete site asset params
func (o *DeleteSiteAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the delete site asset params
func (o *DeleteSiteAssetParams) WithAssetID(assetID string) *DeleteSiteAssetParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the delete site asset params
func (o *DeleteSiteAssetParams) SetAssetID(assetID string) {
	o.AssetID = assetID
}

// WithSiteID adds the siteID to the delete site asset params
func (o *DeleteSiteAssetParams) WithSiteID(siteID string) *DeleteSiteAssetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete site asset params
func (o *DeleteSiteAssetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSiteAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param asset_id
	if err := r.SetPathParam("asset_id", o.AssetID); err != nil {
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
