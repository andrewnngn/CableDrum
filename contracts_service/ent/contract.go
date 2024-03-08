// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golang-boilerplate/ent/contract"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Contract is the model entity for the Contract schema.
type Contract struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VendorID holds the value of the "vendor_id" field.
	VendorID int `json:"vendor_id,omitempty"`
	// Status holds the value of the "status" field.
	Status contract.Status `json:"status,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate *time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate *time.Time `json:"end_date,omitempty"`
	// TotalAmount holds the value of the "total_amount" field.
	TotalAmount int `json:"total_amount,omitempty"`
	// RemainingAmount holds the value of the "remaining_amount" field.
	RemainingAmount int `json:"remaining_amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contract) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contract.FieldID, contract.FieldVendorID, contract.FieldTotalAmount, contract.FieldRemainingAmount:
			values[i] = new(sql.NullInt64)
		case contract.FieldStatus:
			values[i] = new(sql.NullString)
		case contract.FieldStartDate, contract.FieldEndDate, contract.FieldCreatedAt, contract.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contract", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contract fields.
func (c *Contract) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case contract.FieldVendorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vendor_id", values[i])
			} else if value.Valid {
				c.VendorID = int(value.Int64)
			}
		case contract.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = contract.Status(value.String)
			}
		case contract.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				c.StartDate = new(time.Time)
				*c.StartDate = value.Time
			}
		case contract.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				c.EndDate = new(time.Time)
				*c.EndDate = value.Time
			}
		case contract.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				c.TotalAmount = int(value.Int64)
			}
		case contract.FieldRemainingAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remaining_amount", values[i])
			} else if value.Valid {
				c.RemainingAmount = int(value.Int64)
			}
		case contract.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case contract.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Contract.
// Note that you need to call Contract.Unwrap() before calling this method if this Contract
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contract) Update() *ContractUpdateOne {
	return (&ContractClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contract) Unwrap() *Contract {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contract is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contract) String() string {
	var builder strings.Builder
	builder.WriteString("Contract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("vendor_id=")
	builder.WriteString(fmt.Sprintf("%v", c.VendorID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	if v := c.StartDate; v != nil {
		builder.WriteString("start_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.EndDate; v != nil {
		builder.WriteString("end_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(fmt.Sprintf("%v", c.TotalAmount))
	builder.WriteString(", ")
	builder.WriteString("remaining_amount=")
	builder.WriteString(fmt.Sprintf("%v", c.RemainingAmount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Contracts is a parsable slice of Contract.
type Contracts []*Contract

func (c Contracts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}