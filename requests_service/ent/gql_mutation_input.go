// Code generated by ent, DO NOT EDIT.

package ent

import (
	"golang-boilerplate/ent/request"
	"time"
)

// CreateRequestInput represents a mutation input for creating requests.
type CreateRequestInput struct {
	ContractID     int
	VendorID       int
	ContractorID   int
	Amount         int
	Status         *request.Status
	CollectionDate *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

// Mutate applies the CreateRequestInput on the RequestMutation builder.
func (i *CreateRequestInput) Mutate(m *RequestMutation) {
	m.SetContractID(i.ContractID)
	m.SetVendorID(i.VendorID)
	m.SetContractorID(i.ContractorID)
	m.SetAmount(i.Amount)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.CollectionDate; v != nil {
		m.SetCollectionDate(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateRequestInput on the RequestCreate builder.
func (c *RequestCreate) SetInput(i CreateRequestInput) *RequestCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRequestInput represents a mutation input for updating requests.
type UpdateRequestInput struct {
	ContractID          *int
	VendorID            *int
	ContractorID        *int
	Amount              *int
	Status              *request.Status
	ClearCollectionDate bool
	CollectionDate      *time.Time
}

// Mutate applies the UpdateRequestInput on the RequestMutation builder.
func (i *UpdateRequestInput) Mutate(m *RequestMutation) {
	if v := i.ContractID; v != nil {
		m.SetContractID(*v)
	}
	if v := i.VendorID; v != nil {
		m.SetVendorID(*v)
	}
	if v := i.ContractorID; v != nil {
		m.SetContractorID(*v)
	}
	if v := i.Amount; v != nil {
		m.SetAmount(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearCollectionDate {
		m.ClearCollectionDate()
	}
	if v := i.CollectionDate; v != nil {
		m.SetCollectionDate(*v)
	}
}

// SetInput applies the change-set in the UpdateRequestInput on the RequestUpdate builder.
func (c *RequestUpdate) SetInput(i UpdateRequestInput) *RequestUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRequestInput on the RequestUpdateOne builder.
func (c *RequestUpdateOne) SetInput(i UpdateRequestInput) *RequestUpdateOne {
	i.Mutate(c.Mutation())
	return c
}