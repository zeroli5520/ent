// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/car"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetID sets the id field.
func (pc *PetCreate) SetID(s string) *PetCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetOwnerID sets the owner edge to User by id.
func (pc *PetCreate) SetOwnerID(id int) *PetCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(id *int) *PetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the owner edge to User.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// AddCarIDs adds the cars edge to Car by ids.
func (pc *PetCreate) AddCarIDs(ids ...int) *PetCreate {
	pc.mutation.AddCarIDs(ids...)
	return pc
}

// AddCars adds the cars edges to Car.
func (pc *PetCreate) AddCars(c ...*Car) *PetCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCarIDs(ids...)
}

// AddFriendIDs adds the friends edge to Pet by ids.
func (pc *PetCreate) AddFriendIDs(ids ...string) *PetCreate {
	pc.mutation.AddFriendIDs(ids...)
	return pc
}

// AddFriends adds the friends edges to Pet.
func (pc *PetCreate) AddFriends(p ...*Pet) *PetCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddFriendIDs(ids...)
}

// SetBestFriendID sets the best_friend edge to Pet by id.
func (pc *PetCreate) SetBestFriendID(id string) *PetCreate {
	pc.mutation.SetBestFriendID(id)
	return pc
}

// SetNillableBestFriendID sets the best_friend edge to Pet by id if the given value is not nil.
func (pc *PetCreate) SetNillableBestFriendID(id *string) *PetCreate {
	if id != nil {
		pc = pc.SetBestFriendID(*id)
	}
	return pc
}

// SetBestFriend sets the best_friend edge to Pet.
func (pc *PetCreate) SetBestFriend(p *Pet) *PetCreate {
	return pc.SetBestFriendID(p.ID)
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	var (
		err  error
		node *Pet
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	var (
		pe    = &Pet{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: pet.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		pe.ID = id
		_spec.ID.Value = id
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.CarsTable,
			Columns: []string{pet.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   pet.FriendsTable,
			Columns: pet.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   pet.BestFriendTable,
			Columns: []string{pet.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pe, nil
}