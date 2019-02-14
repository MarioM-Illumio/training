// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction transaction
// swagger:model transaction
type Transaction struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// receiver
	// Required: true
	Receiver *int64 `json:"receiver"`

	// sender
	// Required: true
	Sender *int64 `json:"sender"`

	// id
	// Required: true
	ID *string `json:"id"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Transaction) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Amount *float64 `json:"amount"`

		Receiver *int64 `json:"receiver"`

		Sender *int64 `json:"sender"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Amount = dataAO0.Amount

	m.Receiver = dataAO0.Receiver

	m.Sender = dataAO0.Sender

	// AO1
	var dataAO1 struct {
		ID *string `json:"id"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Transaction) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Amount *float64 `json:"amount"`

		Receiver *int64 `json:"receiver"`

		Sender *int64 `json:"sender"`
	}

	dataAO0.Amount = m.Amount

	dataAO0.Receiver = m.Receiver

	dataAO0.Sender = m.Sender

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		ID *string `json:"id"`
	}

	dataAO1.ID = m.ID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateReceiver(formats strfmt.Registry) error {

	if err := validate.Required("receiver", "body", m.Receiver); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
