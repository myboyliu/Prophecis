// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserGroupInfo user group info
// swagger:model UserGroupInfo
type UserGroupInfo struct {

	// Whether the data is valid
	EnableFlag int8 `json:"enableFlag,omitempty"`

	// the gid of user
	Gid int64 `json:"gid,omitempty"`

	// group
	Group *Group `json:"group,omitempty"`

	// true or false
	GUIDCheck int64 `json:"guidCheck,omitempty"`

	// the id of user
	ID int64 `json:"id,omitempty"`

	// the username
	Name string `json:"name,omitempty"`

	// the uid remarks user
	Remarks string `json:"remarks,omitempty"`

	// the client_token of user
	Token string `json:"token,omitempty"`

	// the type of user, user or system
	Type string `json:"type,omitempty"`

	// the uid of user
	UID int64 `json:"uid,omitempty"`
}

// Validate validates this user group info
func (m *UserGroupInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGroupInfo) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserGroupInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroupInfo) UnmarshalBinary(b []byte) error {
	var res UserGroupInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
