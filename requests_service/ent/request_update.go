// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-boilerplate/ent/predicate"
	"golang-boilerplate/ent/request"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RequestUpdate is the builder for updating Request entities.
type RequestUpdate struct {
	config
	hooks    []Hook
	mutation *RequestMutation
}

// Where appends a list predicates to the RequestUpdate builder.
func (ru *RequestUpdate) Where(ps ...predicate.Request) *RequestUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetContractID sets the "contract_id" field.
func (ru *RequestUpdate) SetContractID(i int) *RequestUpdate {
	ru.mutation.ResetContractID()
	ru.mutation.SetContractID(i)
	return ru
}

// AddContractID adds i to the "contract_id" field.
func (ru *RequestUpdate) AddContractID(i int) *RequestUpdate {
	ru.mutation.AddContractID(i)
	return ru
}

// SetVendorID sets the "vendor_id" field.
func (ru *RequestUpdate) SetVendorID(i int) *RequestUpdate {
	ru.mutation.ResetVendorID()
	ru.mutation.SetVendorID(i)
	return ru
}

// AddVendorID adds i to the "vendor_id" field.
func (ru *RequestUpdate) AddVendorID(i int) *RequestUpdate {
	ru.mutation.AddVendorID(i)
	return ru
}

// SetContractorID sets the "contractor_id" field.
func (ru *RequestUpdate) SetContractorID(i int) *RequestUpdate {
	ru.mutation.ResetContractorID()
	ru.mutation.SetContractorID(i)
	return ru
}

// AddContractorID adds i to the "contractor_id" field.
func (ru *RequestUpdate) AddContractorID(i int) *RequestUpdate {
	ru.mutation.AddContractorID(i)
	return ru
}

// SetAmount sets the "amount" field.
func (ru *RequestUpdate) SetAmount(i int) *RequestUpdate {
	ru.mutation.ResetAmount()
	ru.mutation.SetAmount(i)
	return ru
}

// AddAmount adds i to the "amount" field.
func (ru *RequestUpdate) AddAmount(i int) *RequestUpdate {
	ru.mutation.AddAmount(i)
	return ru
}

// SetStatus sets the "status" field.
func (ru *RequestUpdate) SetStatus(r request.Status) *RequestUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RequestUpdate) SetNillableStatus(r *request.Status) *RequestUpdate {
	if r != nil {
		ru.SetStatus(*r)
	}
	return ru
}

// SetCollectionDate sets the "collection_date" field.
func (ru *RequestUpdate) SetCollectionDate(t time.Time) *RequestUpdate {
	ru.mutation.SetCollectionDate(t)
	return ru
}

// SetNillableCollectionDate sets the "collection_date" field if the given value is not nil.
func (ru *RequestUpdate) SetNillableCollectionDate(t *time.Time) *RequestUpdate {
	if t != nil {
		ru.SetCollectionDate(*t)
	}
	return ru
}

// ClearCollectionDate clears the value of the "collection_date" field.
func (ru *RequestUpdate) ClearCollectionDate() *RequestUpdate {
	ru.mutation.ClearCollectionDate()
	return ru
}

// Mutation returns the RequestMutation object of the builder.
func (ru *RequestUpdate) Mutation() *RequestMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RequestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RequestUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RequestUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RequestUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RequestUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := request.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RequestUpdate) check() error {
	if v, ok := ru.mutation.Status(); ok {
		if err := request.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Request.status": %w`, err)}
		}
	}
	return nil
}

func (ru *RequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   request.Table,
			Columns: request.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: request.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ContractID(); ok {
		_spec.SetField(request.FieldContractID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedContractID(); ok {
		_spec.AddField(request.FieldContractID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.VendorID(); ok {
		_spec.SetField(request.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedVendorID(); ok {
		_spec.AddField(request.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.ContractorID(); ok {
		_spec.SetField(request.FieldContractorID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedContractorID(); ok {
		_spec.AddField(request.FieldContractorID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Amount(); ok {
		_spec.SetField(request.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedAmount(); ok {
		_spec.AddField(request.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(request.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ru.mutation.CollectionDate(); ok {
		_spec.SetField(request.FieldCollectionDate, field.TypeTime, value)
	}
	if ru.mutation.CollectionDateCleared() {
		_spec.ClearField(request.FieldCollectionDate, field.TypeTime)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(request.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{request.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RequestUpdateOne is the builder for updating a single Request entity.
type RequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RequestMutation
}

// SetContractID sets the "contract_id" field.
func (ruo *RequestUpdateOne) SetContractID(i int) *RequestUpdateOne {
	ruo.mutation.ResetContractID()
	ruo.mutation.SetContractID(i)
	return ruo
}

// AddContractID adds i to the "contract_id" field.
func (ruo *RequestUpdateOne) AddContractID(i int) *RequestUpdateOne {
	ruo.mutation.AddContractID(i)
	return ruo
}

// SetVendorID sets the "vendor_id" field.
func (ruo *RequestUpdateOne) SetVendorID(i int) *RequestUpdateOne {
	ruo.mutation.ResetVendorID()
	ruo.mutation.SetVendorID(i)
	return ruo
}

// AddVendorID adds i to the "vendor_id" field.
func (ruo *RequestUpdateOne) AddVendorID(i int) *RequestUpdateOne {
	ruo.mutation.AddVendorID(i)
	return ruo
}

// SetContractorID sets the "contractor_id" field.
func (ruo *RequestUpdateOne) SetContractorID(i int) *RequestUpdateOne {
	ruo.mutation.ResetContractorID()
	ruo.mutation.SetContractorID(i)
	return ruo
}

// AddContractorID adds i to the "contractor_id" field.
func (ruo *RequestUpdateOne) AddContractorID(i int) *RequestUpdateOne {
	ruo.mutation.AddContractorID(i)
	return ruo
}

// SetAmount sets the "amount" field.
func (ruo *RequestUpdateOne) SetAmount(i int) *RequestUpdateOne {
	ruo.mutation.ResetAmount()
	ruo.mutation.SetAmount(i)
	return ruo
}

// AddAmount adds i to the "amount" field.
func (ruo *RequestUpdateOne) AddAmount(i int) *RequestUpdateOne {
	ruo.mutation.AddAmount(i)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RequestUpdateOne) SetStatus(r request.Status) *RequestUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RequestUpdateOne) SetNillableStatus(r *request.Status) *RequestUpdateOne {
	if r != nil {
		ruo.SetStatus(*r)
	}
	return ruo
}

// SetCollectionDate sets the "collection_date" field.
func (ruo *RequestUpdateOne) SetCollectionDate(t time.Time) *RequestUpdateOne {
	ruo.mutation.SetCollectionDate(t)
	return ruo
}

// SetNillableCollectionDate sets the "collection_date" field if the given value is not nil.
func (ruo *RequestUpdateOne) SetNillableCollectionDate(t *time.Time) *RequestUpdateOne {
	if t != nil {
		ruo.SetCollectionDate(*t)
	}
	return ruo
}

// ClearCollectionDate clears the value of the "collection_date" field.
func (ruo *RequestUpdateOne) ClearCollectionDate() *RequestUpdateOne {
	ruo.mutation.ClearCollectionDate()
	return ruo
}

// Mutation returns the RequestMutation object of the builder.
func (ruo *RequestUpdateOne) Mutation() *RequestMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RequestUpdateOne) Select(field string, fields ...string) *RequestUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Request entity.
func (ruo *RequestUpdateOne) Save(ctx context.Context) (*Request, error) {
	var (
		err  error
		node *Request
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Request)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RequestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RequestUpdateOne) SaveX(ctx context.Context) *Request {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RequestUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RequestUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RequestUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := request.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RequestUpdateOne) check() error {
	if v, ok := ruo.mutation.Status(); ok {
		if err := request.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Request.status": %w`, err)}
		}
	}
	return nil
}

func (ruo *RequestUpdateOne) sqlSave(ctx context.Context) (_node *Request, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   request.Table,
			Columns: request.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: request.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Request.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, request.FieldID)
		for _, f := range fields {
			if !request.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != request.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ContractID(); ok {
		_spec.SetField(request.FieldContractID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedContractID(); ok {
		_spec.AddField(request.FieldContractID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.VendorID(); ok {
		_spec.SetField(request.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedVendorID(); ok {
		_spec.AddField(request.FieldVendorID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.ContractorID(); ok {
		_spec.SetField(request.FieldContractorID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedContractorID(); ok {
		_spec.AddField(request.FieldContractorID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Amount(); ok {
		_spec.SetField(request.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedAmount(); ok {
		_spec.AddField(request.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(request.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ruo.mutation.CollectionDate(); ok {
		_spec.SetField(request.FieldCollectionDate, field.TypeTime, value)
	}
	if ruo.mutation.CollectionDateCleared() {
		_spec.ClearField(request.FieldCollectionDate, field.TypeTime)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(request.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Request{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{request.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}