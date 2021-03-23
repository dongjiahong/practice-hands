// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sqlent/ent/user"
	"sqlent/ent/userbuyrecord"
	"sqlent/ent/usercount"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetPID sets the "p_id" field.
func (uc *UserCreate) SetPID(i int) *UserCreate {
	uc.mutation.SetPID(i)
	return uc
}

// SetNillablePID sets the "p_id" field if the given value is not nil.
func (uc *UserCreate) SetNillablePID(i *int) *UserCreate {
	if i != nil {
		uc.SetPID(*i)
	}
	return uc
}

// SetInvitedCode sets the "invited_code" field.
func (uc *UserCreate) SetInvitedCode(s string) *UserCreate {
	uc.mutation.SetInvitedCode(s)
	return uc
}

// SetCreated sets the "created" field.
func (uc *UserCreate) SetCreated(i int64) *UserCreate {
	uc.mutation.SetCreated(i)
	return uc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreated(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreated(*i)
	}
	return uc
}

// SetUpdated sets the "updated" field.
func (uc *UserCreate) SetUpdated(i int64) *UserCreate {
	uc.mutation.SetUpdated(i)
	return uc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdated(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdated(*i)
	}
	return uc
}

// SetDeleted sets the "deleted" field.
func (uc *UserCreate) SetDeleted(i int64) *UserCreate {
	uc.mutation.SetDeleted(i)
	return uc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeleted(i *int64) *UserCreate {
	if i != nil {
		uc.SetDeleted(*i)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// SetCountID sets the "count" edge to the UserCount entity by ID.
func (uc *UserCreate) SetCountID(id int) *UserCreate {
	uc.mutation.SetCountID(id)
	return uc
}

// SetNillableCountID sets the "count" edge to the UserCount entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCountID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCountID(*id)
	}
	return uc
}

// SetCount sets the "count" edge to the UserCount entity.
func (uc *UserCreate) SetCount(u *UserCount) *UserCreate {
	return uc.SetCountID(u.ID)
}

// AddBuyRecordIDs adds the "buy_record" edge to the UserBuyRecord entity by IDs.
func (uc *UserCreate) AddBuyRecordIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddBuyRecordIDs(ids...)
	return uc
}

// AddBuyRecord adds the "buy_record" edges to the UserBuyRecord entity.
func (uc *UserCreate) AddBuyRecord(u ...*UserBuyRecord) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddBuyRecordIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.PID(); !ok {
		v := user.DefaultPID
		uc.mutation.SetPID(v)
	}
	if _, ok := uc.mutation.Created(); !ok {
		v := user.DefaultCreated
		uc.mutation.SetCreated(v)
	}
	if _, ok := uc.mutation.Updated(); !ok {
		v := user.DefaultUpdated
		uc.mutation.SetUpdated(v)
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		v := user.DefaultDeleted
		uc.mutation.SetDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if v, ok := uc.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if _, ok := uc.mutation.PID(); !ok {
		return &ValidationError{Name: "p_id", err: errors.New("ent: missing required field \"p_id\"")}
	}
	if _, ok := uc.mutation.InvitedCode(); !ok {
		return &ValidationError{Name: "invited_code", err: errors.New("ent: missing required field \"invited_code\"")}
	}
	if v, ok := uc.mutation.InvitedCode(); ok {
		if err := user.InvitedCodeValidator(v); err != nil {
			return &ValidationError{Name: "invited_code", err: fmt.Errorf("ent: validator failed for field \"invited_code\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New("ent: missing required field \"created\"")}
	}
	if _, ok := uc.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New("ent: missing required field \"updated\"")}
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New("ent: missing required field \"deleted\"")}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.PID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldPID,
		})
		_node.PID = value
	}
	if value, ok := uc.mutation.InvitedCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldInvitedCode,
		})
		_node.InvitedCode = value
	}
	if value, ok := uc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := uc.mutation.Updated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldUpdated,
		})
		_node.Updated = value
	}
	if value, ok := uc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldDeleted,
		})
		_node.Deleted = value
	}
	if nodes := uc.mutation.CountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CountTable,
			Columns: []string{user.CountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usercount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BuyRecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BuyRecordTable,
			Columns: []string{user.BuyRecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: userbuyrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
