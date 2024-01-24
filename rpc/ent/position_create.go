// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/oms-core/rpc/ent/position"
	"github.com/iot-synergy/oms-core/rpc/ent/user"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PositionCreate) SetCreatedAt(t time.Time) *PositionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreatedAt(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PositionCreate) SetUpdatedAt(t time.Time) *PositionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PositionCreate) SetNillableUpdatedAt(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PositionCreate) SetStatus(u uint8) *PositionCreate {
	pc.mutation.SetStatus(u)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PositionCreate) SetNillableStatus(u *uint8) *PositionCreate {
	if u != nil {
		pc.SetStatus(*u)
	}
	return pc
}

// SetSort sets the "sort" field.
func (pc *PositionCreate) SetSort(u uint32) *PositionCreate {
	pc.mutation.SetSort(u)
	return pc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pc *PositionCreate) SetNillableSort(u *uint32) *PositionCreate {
	if u != nil {
		pc.SetSort(*u)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PositionCreate) SetName(s string) *PositionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetCode sets the "code" field.
func (pc *PositionCreate) SetCode(s string) *PositionCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetRemark sets the "remark" field.
func (pc *PositionCreate) SetRemark(s string) *PositionCreate {
	pc.mutation.SetRemark(s)
	return pc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pc *PositionCreate) SetNillableRemark(s *string) *PositionCreate {
	if s != nil {
		pc.SetRemark(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PositionCreate) SetID(u uint64) *PositionCreate {
	pc.mutation.SetID(u)
	return pc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pc *PositionCreate) AddUserIDs(ids ...uuid.UUID) *PositionCreate {
	pc.mutation.AddUserIDs(ids...)
	return pc
}

// AddUsers adds the "users" edges to the User entity.
func (pc *PositionCreate) AddUsers(u ...*User) *PositionCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PositionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PositionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PositionCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := position.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := position.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := position.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Sort(); !ok {
		v := position.DefaultSort
		pc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PositionCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Position.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Position.updated_at"`)}
	}
	if _, ok := pc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Position.sort"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Position.name"`)}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Position.code"`)}
	}
	return nil
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		_node = &Position{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(position.Table, sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint64))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(position.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(position.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.Sort(); ok {
		_spec.SetField(position.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if nodes := pc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: position.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PositionCreateBulk is the builder for creating many Position entities in bulk.
type PositionCreateBulk struct {
	config
	err      error
	builders []*PositionCreate
}

// Save creates the Position entities in the database.
func (pcb *PositionCreateBulk) Save(ctx context.Context) ([]*Position, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Position, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PositionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PositionCreateBulk) SaveX(ctx context.Context) []*Position {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PositionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PositionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
