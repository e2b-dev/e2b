// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/envalias"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
)

// EnvAliasUpdate is the builder for updating EnvAlias entities.
type EnvAliasUpdate struct {
	config
	hooks     []Hook
	mutation  *EnvAliasMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EnvAliasUpdate builder.
func (eau *EnvAliasUpdate) Where(ps ...predicate.EnvAlias) *EnvAliasUpdate {
	eau.mutation.Where(ps...)
	return eau
}

// SetEnvID sets the "env_id" field.
func (eau *EnvAliasUpdate) SetEnvID(s string) *EnvAliasUpdate {
	eau.mutation.SetEnvID(s)
	return eau
}

// SetNillableEnvID sets the "env_id" field if the given value is not nil.
func (eau *EnvAliasUpdate) SetNillableEnvID(s *string) *EnvAliasUpdate {
	if s != nil {
		eau.SetEnvID(*s)
	}
	return eau
}

// ClearEnvID clears the value of the "env_id" field.
func (eau *EnvAliasUpdate) ClearEnvID() *EnvAliasUpdate {
	eau.mutation.ClearEnvID()
	return eau
}

// SetIsName sets the "is_name" field.
func (eau *EnvAliasUpdate) SetIsName(b bool) *EnvAliasUpdate {
	eau.mutation.SetIsName(b)
	return eau
}

// SetNillableIsName sets the "is_name" field if the given value is not nil.
func (eau *EnvAliasUpdate) SetNillableIsName(b *bool) *EnvAliasUpdate {
	if b != nil {
		eau.SetIsName(*b)
	}
	return eau
}

// SetEnv sets the "env" edge to the Env entity.
func (eau *EnvAliasUpdate) SetEnv(e *Env) *EnvAliasUpdate {
	return eau.SetEnvID(e.ID)
}

// Mutation returns the EnvAliasMutation object of the builder.
func (eau *EnvAliasUpdate) Mutation() *EnvAliasMutation {
	return eau.mutation
}

// ClearEnv clears the "env" edge to the Env entity.
func (eau *EnvAliasUpdate) ClearEnv() *EnvAliasUpdate {
	eau.mutation.ClearEnv()
	return eau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eau *EnvAliasUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eau.sqlSave, eau.mutation, eau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eau *EnvAliasUpdate) SaveX(ctx context.Context) int {
	affected, err := eau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eau *EnvAliasUpdate) Exec(ctx context.Context) error {
	_, err := eau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eau *EnvAliasUpdate) ExecX(ctx context.Context) {
	if err := eau.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eau *EnvAliasUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EnvAliasUpdate {
	eau.modifiers = append(eau.modifiers, modifiers...)
	return eau
}

func (eau *EnvAliasUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(envalias.Table, envalias.Columns, sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeString))
	if ps := eau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eau.mutation.IsName(); ok {
		_spec.SetField(envalias.FieldIsName, field.TypeBool, value)
	}
	if eau.mutation.EnvCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   envalias.EnvTable,
			Columns: []string{envalias.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = eau.schemaConfig.EnvAlias
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.EnvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   envalias.EnvTable,
			Columns: []string{envalias.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = eau.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = eau.schemaConfig.EnvAlias
	ctx = internal.NewSchemaConfigContext(ctx, eau.schemaConfig)
	_spec.AddModifiers(eau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, eau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{envalias.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eau.mutation.done = true
	return n, nil
}

// EnvAliasUpdateOne is the builder for updating a single EnvAlias entity.
type EnvAliasUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EnvAliasMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEnvID sets the "env_id" field.
func (eauo *EnvAliasUpdateOne) SetEnvID(s string) *EnvAliasUpdateOne {
	eauo.mutation.SetEnvID(s)
	return eauo
}

// SetNillableEnvID sets the "env_id" field if the given value is not nil.
func (eauo *EnvAliasUpdateOne) SetNillableEnvID(s *string) *EnvAliasUpdateOne {
	if s != nil {
		eauo.SetEnvID(*s)
	}
	return eauo
}

// ClearEnvID clears the value of the "env_id" field.
func (eauo *EnvAliasUpdateOne) ClearEnvID() *EnvAliasUpdateOne {
	eauo.mutation.ClearEnvID()
	return eauo
}

// SetIsName sets the "is_name" field.
func (eauo *EnvAliasUpdateOne) SetIsName(b bool) *EnvAliasUpdateOne {
	eauo.mutation.SetIsName(b)
	return eauo
}

// SetNillableIsName sets the "is_name" field if the given value is not nil.
func (eauo *EnvAliasUpdateOne) SetNillableIsName(b *bool) *EnvAliasUpdateOne {
	if b != nil {
		eauo.SetIsName(*b)
	}
	return eauo
}

// SetEnv sets the "env" edge to the Env entity.
func (eauo *EnvAliasUpdateOne) SetEnv(e *Env) *EnvAliasUpdateOne {
	return eauo.SetEnvID(e.ID)
}

// Mutation returns the EnvAliasMutation object of the builder.
func (eauo *EnvAliasUpdateOne) Mutation() *EnvAliasMutation {
	return eauo.mutation
}

// ClearEnv clears the "env" edge to the Env entity.
func (eauo *EnvAliasUpdateOne) ClearEnv() *EnvAliasUpdateOne {
	eauo.mutation.ClearEnv()
	return eauo
}

// Where appends a list predicates to the EnvAliasUpdate builder.
func (eauo *EnvAliasUpdateOne) Where(ps ...predicate.EnvAlias) *EnvAliasUpdateOne {
	eauo.mutation.Where(ps...)
	return eauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eauo *EnvAliasUpdateOne) Select(field string, fields ...string) *EnvAliasUpdateOne {
	eauo.fields = append([]string{field}, fields...)
	return eauo
}

// Save executes the query and returns the updated EnvAlias entity.
func (eauo *EnvAliasUpdateOne) Save(ctx context.Context) (*EnvAlias, error) {
	return withHooks(ctx, eauo.sqlSave, eauo.mutation, eauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eauo *EnvAliasUpdateOne) SaveX(ctx context.Context) *EnvAlias {
	node, err := eauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eauo *EnvAliasUpdateOne) Exec(ctx context.Context) error {
	_, err := eauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eauo *EnvAliasUpdateOne) ExecX(ctx context.Context) {
	if err := eauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eauo *EnvAliasUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EnvAliasUpdateOne {
	eauo.modifiers = append(eauo.modifiers, modifiers...)
	return eauo
}

func (eauo *EnvAliasUpdateOne) sqlSave(ctx context.Context) (_node *EnvAlias, err error) {
	_spec := sqlgraph.NewUpdateSpec(envalias.Table, envalias.Columns, sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeString))
	id, ok := eauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "EnvAlias.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, envalias.FieldID)
		for _, f := range fields {
			if !envalias.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != envalias.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eauo.mutation.IsName(); ok {
		_spec.SetField(envalias.FieldIsName, field.TypeBool, value)
	}
	if eauo.mutation.EnvCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   envalias.EnvTable,
			Columns: []string{envalias.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = eauo.schemaConfig.EnvAlias
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.EnvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   envalias.EnvTable,
			Columns: []string{envalias.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = eauo.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = eauo.schemaConfig.EnvAlias
	ctx = internal.NewSchemaConfigContext(ctx, eauo.schemaConfig)
	_spec.AddModifiers(eauo.modifiers...)
	_node = &EnvAlias{config: eauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{envalias.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eauo.mutation.done = true
	return _node, nil
}