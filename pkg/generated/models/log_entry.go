// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogEntry log entry
//
// swagger:model LogEntry
type LogEntry map[string]LogEntryAnon

// Validate validates this log entry
func (m LogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if swag.IsZero(m[k]) { // not required
			continue
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// LogEntryAnon log entry anon
//
// swagger:model LogEntryAnon
type LogEntryAnon struct {

	// extra data
	ExtraData map[string]string `json:"extraData,omitempty"`

	// log index
	// Required: true
	// Minimum: 1
	LogIndex *int64 `json:"logIndex"`

	// signature
	// Required: true
	Signature *LogEntryAnonSignature `json:"signature"`

	// signed content s h a256
	// Required: true
	// Pattern: ^[0-9a-fA-F]{64}$
	SignedContentSHA256 *string `json:"signedContentSHA256"`
}

// Validate validates this log entry anon
func (m *LogEntryAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignedContentSHA256(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnon) validateLogIndex(formats strfmt.Registry) error {

	if err := validate.Required("logIndex", "body", m.LogIndex); err != nil {
		return err
	}

	if err := validate.MinimumInt("logIndex", "body", int64(*m.LogIndex), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *LogEntryAnon) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

func (m *LogEntryAnon) validateSignedContentSHA256(formats strfmt.Registry) error {

	if err := validate.Required("signedContentSHA256", "body", m.SignedContentSHA256); err != nil {
		return err
	}

	if err := validate.Pattern("signedContentSHA256", "body", string(*m.SignedContentSHA256), `^[0-9a-fA-F]{64}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogEntryAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntryAnon) UnmarshalBinary(b []byte) error {
	var res LogEntryAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LogEntryAnonSignature log entry anon signature
//
// swagger:model LogEntryAnonSignature
type LogEntryAnonSignature struct {

	// content
	// Required: true
	// Format: byte
	Content *strfmt.Base64 `json:"content"`

	// format
	// Required: true
	Format SupportedPKIFormats `json:"format"`

	// public key
	// Required: true
	// Format: byte
	PublicKey *strfmt.Base64 `json:"publicKey"`
}

// Validate validates this log entry anon signature
func (m *LogEntryAnonSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnonSignature) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("signature"+"."+"content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *LogEntryAnonSignature) validateFormat(formats strfmt.Registry) error {

	if err := m.Format.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signature" + "." + "format")
		}
		return err
	}

	return nil
}

func (m *LogEntryAnonSignature) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("signature"+"."+"publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogEntryAnonSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntryAnonSignature) UnmarshalBinary(b []byte) error {
	var res LogEntryAnonSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}