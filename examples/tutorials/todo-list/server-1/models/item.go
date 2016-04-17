package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Item item

swagger:model item
*/
type Item struct {

	/* completed
	 */
	Completed bool `json:"completed,omitempty"`

	/* description

	Required: true
	Min Length: 1
	*/
	Description *string `json:"description"`

	/* id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}
