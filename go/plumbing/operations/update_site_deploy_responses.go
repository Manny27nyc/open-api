// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// UpdateSiteDeployReader is a Reader for the UpdateSiteDeploy structure.
type UpdateSiteDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSiteDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSiteDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSiteDeployDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSiteDeployOK creates a UpdateSiteDeployOK with default headers values
func NewUpdateSiteDeployOK() *UpdateSiteDeployOK {
	return &UpdateSiteDeployOK{}
}

/*UpdateSiteDeployOK handles this case with default header values.

OK
*/
type UpdateSiteDeployOK struct {
	Payload *models.Deploy
}

func (o *UpdateSiteDeployOK) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/deploys/{deploy_id}][%d] updateSiteDeployOK  %+v", 200, o.Payload)
}

func (o *UpdateSiteDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSiteDeployDefault creates a UpdateSiteDeployDefault with default headers values
func NewUpdateSiteDeployDefault(code int) *UpdateSiteDeployDefault {
	return &UpdateSiteDeployDefault{
		_statusCode: code,
	}
}

/*UpdateSiteDeployDefault handles this case with default header values.

error
*/
type UpdateSiteDeployDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update site deploy default response
func (o *UpdateSiteDeployDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSiteDeployDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/deploys/{deploy_id}][%d] updateSiteDeploy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSiteDeployDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
