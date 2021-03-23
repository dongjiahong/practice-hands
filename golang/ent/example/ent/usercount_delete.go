// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sqlent/ent/predicate"
	"sqlent/ent/usercount"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCountDelete is the builder for deleting a UserCount entity.
type UserCountDelete struct {
	config
	hooks    []Hook
	mutation *UserCountMutation
}

// Where adds a new predicate to the UserCountDelete builder.
func (ucd *UserCountDelete) Where(ps ...predicate.UserCount) *UserCountDelete {
	ucd.mutation.predicates = append(ucd.mutation.predicates, ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserCountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ucd.hooks) == 0 {
		affected, err = ucd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucd.mutation = mutation
			affected, err = ucd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucd.hooks) - 1; i >= 0; i-- {
			mut = ucd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserCountDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usercount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usercount.FieldID,
			},
		},
	}
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
}

// UserCountDeleteOne is the builder for deleting a single UserCount entity.
type UserCountDeleteOne struct {
	ucd *UserCountDelete
}

// Exec executes the deletion query.
func (ucdo *UserCountDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usercount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserCountDeleteOne) ExecX(ctx context.Context) {
	ucdo.ucd.ExecX(ctx)
}