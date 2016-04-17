package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	"github.com/go-openapi/strfmt"
)

// NewDeletePetParams creates a new DeletePetParams object
// with the default values initialized.
func NewDeletePetParams() DeletePetParams {
	var ()
	return DeletePetParams{}
}

// DeletePetParams contains all the bound params for the delete pet operation
// typically these are obtained from a http.Request
//
// swagger:parameters deletePet
type DeletePetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: header
	*/
	APIKey string
	/*Pet id to delete
	  Required: true
	  In: path
	*/
	PetID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeletePetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := o.bindAPIKey(r.Header["api_key"], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rPetID, rhkPetID, _ := route.Params.GetOK("petId")
	if err := o.bindPetID(rPetID, rhkPetID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePetParams) bindAPIKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("api_key", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("api_key", "header", raw); err != nil {
		return err
	}

	o.APIKey = raw

	return nil
}

func (o *DeletePetParams) bindPetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("petId", "path", "int64", raw)
	}
	o.PetID = value

	return nil
}
