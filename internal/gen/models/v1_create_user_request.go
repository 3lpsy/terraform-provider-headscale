// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateUserRequest v1 create user request
//
// swagger:model v1CreateUserRequest
type V1CreateUserRequest struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// picture Url
	PictureURL string `json:"pictureUrl,omitempty"`
}

// Validate validates this v1 create user request
func (m *V1CreateUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 create user request based on context it is used
func (m *V1CreateUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateUserRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
