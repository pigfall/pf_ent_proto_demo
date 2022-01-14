// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/pigfall/pf_ent_proto_demo/ent/car"
	"github.com/pigfall/pf_ent_proto_demo/ent/dept"
	"github.com/pigfall/pf_ent_proto_demo/ent/predicate"
	"github.com/pigfall/pf_ent_proto_demo/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCar  = "Car"
	TypeDept = "Dept"
	TypeUser = "User"
)

// CarMutation represents an operation that mutates the Car nodes in the graph.
type CarMutation struct {
	config
	op            Op
	typ           string
	id            *string
	clearedFields map[string]struct{}
	user          *string
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Car, error)
	predicates    []predicate.Car
}

var _ ent.Mutation = (*CarMutation)(nil)

// carOption allows management of the mutation configuration using functional options.
type carOption func(*CarMutation)

// newCarMutation creates new mutation for the Car entity.
func newCarMutation(c config, op Op, opts ...carOption) *CarMutation {
	m := &CarMutation{
		config:        c,
		op:            op,
		typ:           TypeCar,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCarID sets the ID field of the mutation.
func withCarID(id string) carOption {
	return func(m *CarMutation) {
		var (
			err   error
			once  sync.Once
			value *Car
		)
		m.oldValue = func(ctx context.Context) (*Car, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Car.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCar sets the old Car of the mutation.
func withCar(node *Car) carOption {
	return func(m *CarMutation) {
		m.oldValue = func(context.Context) (*Car, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CarMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CarMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Car entities.
func (m *CarMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CarMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *CarMutation) SetUserID(id string) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *CarMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *CarMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *CarMutation) UserID() (id string, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *CarMutation) UserIDs() (ids []string) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *CarMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the CarMutation builder.
func (m *CarMutation) Where(ps ...predicate.Car) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CarMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CarMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Car field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, car.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case car.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, car.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	switch name {
	case car.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	switch name {
	case car.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	switch name {
	case car.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Car edge %s", name)
}

// DeptMutation represents an operation that mutates the Dept nodes in the graph.
type DeptMutation struct {
	config
	op            Op
	typ           string
	id            *string
	clearedFields map[string]struct{}
	user          map[string]struct{}
	removeduser   map[string]struct{}
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Dept, error)
	predicates    []predicate.Dept
}

var _ ent.Mutation = (*DeptMutation)(nil)

// deptOption allows management of the mutation configuration using functional options.
type deptOption func(*DeptMutation)

// newDeptMutation creates new mutation for the Dept entity.
func newDeptMutation(c config, op Op, opts ...deptOption) *DeptMutation {
	m := &DeptMutation{
		config:        c,
		op:            op,
		typ:           TypeDept,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDeptID sets the ID field of the mutation.
func withDeptID(id string) deptOption {
	return func(m *DeptMutation) {
		var (
			err   error
			once  sync.Once
			value *Dept
		)
		m.oldValue = func(ctx context.Context) (*Dept, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Dept.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDept sets the old Dept of the mutation.
func withDept(node *Dept) deptOption {
	return func(m *DeptMutation) {
		m.oldValue = func(context.Context) (*Dept, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DeptMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DeptMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Dept entities.
func (m *DeptMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DeptMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// AddUserIDs adds the "user" edge to the User entity by ids.
func (m *DeptMutation) AddUserIDs(ids ...string) {
	if m.user == nil {
		m.user = make(map[string]struct{})
	}
	for i := range ids {
		m.user[ids[i]] = struct{}{}
	}
}

// ClearUser clears the "user" edge to the User entity.
func (m *DeptMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *DeptMutation) UserCleared() bool {
	return m.cleareduser
}

// RemoveUserIDs removes the "user" edge to the User entity by IDs.
func (m *DeptMutation) RemoveUserIDs(ids ...string) {
	if m.removeduser == nil {
		m.removeduser = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.user, ids[i])
		m.removeduser[ids[i]] = struct{}{}
	}
}

// RemovedUser returns the removed IDs of the "user" edge to the User entity.
func (m *DeptMutation) RemovedUserIDs() (ids []string) {
	for id := range m.removeduser {
		ids = append(ids, id)
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
func (m *DeptMutation) UserIDs() (ids []string) {
	for id := range m.user {
		ids = append(ids, id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *DeptMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
	m.removeduser = nil
}

// Where appends a list predicates to the DeptMutation builder.
func (m *DeptMutation) Where(ps ...predicate.Dept) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DeptMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Dept).
func (m *DeptMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DeptMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DeptMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DeptMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Dept field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DeptMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Dept field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DeptMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DeptMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DeptMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Dept numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DeptMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DeptMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DeptMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Dept nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DeptMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Dept field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DeptMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, dept.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DeptMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case dept.EdgeUser:
		ids := make([]ent.Value, 0, len(m.user))
		for id := range m.user {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DeptMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeduser != nil {
		edges = append(edges, dept.EdgeUser)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DeptMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case dept.EdgeUser:
		ids := make([]ent.Value, 0, len(m.removeduser))
		for id := range m.removeduser {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DeptMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, dept.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DeptMutation) EdgeCleared(name string) bool {
	switch name {
	case dept.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DeptMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Dept unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DeptMutation) ResetEdge(name string) error {
	switch name {
	case dept.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Dept edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *string
	phone         *string
	race          *user.Race
	clearedFields map[string]struct{}
	car           map[string]struct{}
	removedcar    map[string]struct{}
	clearedcar    bool
	dept          map[string]struct{}
	removeddept   map[string]struct{}
	cleareddept   bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id string) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetPhone sets the "phone" field.
func (m *UserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *UserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ClearPhone clears the value of the "phone" field.
func (m *UserMutation) ClearPhone() {
	m.phone = nil
	m.clearedFields[user.FieldPhone] = struct{}{}
}

// PhoneCleared returns if the "phone" field was cleared in this mutation.
func (m *UserMutation) PhoneCleared() bool {
	_, ok := m.clearedFields[user.FieldPhone]
	return ok
}

// ResetPhone resets all changes to the "phone" field.
func (m *UserMutation) ResetPhone() {
	m.phone = nil
	delete(m.clearedFields, user.FieldPhone)
}

// SetRace sets the "race" field.
func (m *UserMutation) SetRace(u user.Race) {
	m.race = &u
}

// Race returns the value of the "race" field in the mutation.
func (m *UserMutation) Race() (r user.Race, exists bool) {
	v := m.race
	if v == nil {
		return
	}
	return *v, true
}

// OldRace returns the old "race" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRace(ctx context.Context) (v user.Race, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRace is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRace requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRace: %w", err)
	}
	return oldValue.Race, nil
}

// ResetRace resets all changes to the "race" field.
func (m *UserMutation) ResetRace() {
	m.race = nil
}

// AddCarIDs adds the "car" edge to the Car entity by ids.
func (m *UserMutation) AddCarIDs(ids ...string) {
	if m.car == nil {
		m.car = make(map[string]struct{})
	}
	for i := range ids {
		m.car[ids[i]] = struct{}{}
	}
}

// ClearCar clears the "car" edge to the Car entity.
func (m *UserMutation) ClearCar() {
	m.clearedcar = true
}

// CarCleared reports if the "car" edge to the Car entity was cleared.
func (m *UserMutation) CarCleared() bool {
	return m.clearedcar
}

// RemoveCarIDs removes the "car" edge to the Car entity by IDs.
func (m *UserMutation) RemoveCarIDs(ids ...string) {
	if m.removedcar == nil {
		m.removedcar = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.car, ids[i])
		m.removedcar[ids[i]] = struct{}{}
	}
}

// RemovedCar returns the removed IDs of the "car" edge to the Car entity.
func (m *UserMutation) RemovedCarIDs() (ids []string) {
	for id := range m.removedcar {
		ids = append(ids, id)
	}
	return
}

// CarIDs returns the "car" edge IDs in the mutation.
func (m *UserMutation) CarIDs() (ids []string) {
	for id := range m.car {
		ids = append(ids, id)
	}
	return
}

// ResetCar resets all changes to the "car" edge.
func (m *UserMutation) ResetCar() {
	m.car = nil
	m.clearedcar = false
	m.removedcar = nil
}

// AddDeptIDs adds the "dept" edge to the Dept entity by ids.
func (m *UserMutation) AddDeptIDs(ids ...string) {
	if m.dept == nil {
		m.dept = make(map[string]struct{})
	}
	for i := range ids {
		m.dept[ids[i]] = struct{}{}
	}
}

// ClearDept clears the "dept" edge to the Dept entity.
func (m *UserMutation) ClearDept() {
	m.cleareddept = true
}

// DeptCleared reports if the "dept" edge to the Dept entity was cleared.
func (m *UserMutation) DeptCleared() bool {
	return m.cleareddept
}

// RemoveDeptIDs removes the "dept" edge to the Dept entity by IDs.
func (m *UserMutation) RemoveDeptIDs(ids ...string) {
	if m.removeddept == nil {
		m.removeddept = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.dept, ids[i])
		m.removeddept[ids[i]] = struct{}{}
	}
}

// RemovedDept returns the removed IDs of the "dept" edge to the Dept entity.
func (m *UserMutation) RemovedDeptIDs() (ids []string) {
	for id := range m.removeddept {
		ids = append(ids, id)
	}
	return
}

// DeptIDs returns the "dept" edge IDs in the mutation.
func (m *UserMutation) DeptIDs() (ids []string) {
	for id := range m.dept {
		ids = append(ids, id)
	}
	return
}

// ResetDept resets all changes to the "dept" edge.
func (m *UserMutation) ResetDept() {
	m.dept = nil
	m.cleareddept = false
	m.removeddept = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.phone != nil {
		fields = append(fields, user.FieldPhone)
	}
	if m.race != nil {
		fields = append(fields, user.FieldRace)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldPhone:
		return m.Phone()
	case user.FieldRace:
		return m.Race()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldPhone:
		return m.OldPhone(ctx)
	case user.FieldRace:
		return m.OldRace(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case user.FieldRace:
		v, ok := value.(user.Race)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRace(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldPhone) {
		fields = append(fields, user.FieldPhone)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldPhone:
		m.ClearPhone()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldPhone:
		m.ResetPhone()
		return nil
	case user.FieldRace:
		m.ResetRace()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.car != nil {
		edges = append(edges, user.EdgeCar)
	}
	if m.dept != nil {
		edges = append(edges, user.EdgeDept)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCar:
		ids := make([]ent.Value, 0, len(m.car))
		for id := range m.car {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeDept:
		ids := make([]ent.Value, 0, len(m.dept))
		for id := range m.dept {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedcar != nil {
		edges = append(edges, user.EdgeCar)
	}
	if m.removeddept != nil {
		edges = append(edges, user.EdgeDept)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCar:
		ids := make([]ent.Value, 0, len(m.removedcar))
		for id := range m.removedcar {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeDept:
		ids := make([]ent.Value, 0, len(m.removeddept))
		for id := range m.removeddept {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedcar {
		edges = append(edges, user.EdgeCar)
	}
	if m.cleareddept {
		edges = append(edges, user.EdgeDept)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeCar:
		return m.clearedcar
	case user.EdgeDept:
		return m.cleareddept
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCar:
		m.ResetCar()
		return nil
	case user.EdgeDept:
		m.ResetDept()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
