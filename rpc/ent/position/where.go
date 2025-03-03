// Code generated by ent, DO NOT EDIT.

package position

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/iot-synergy/oms-core/rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldSort, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCode, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldStatus))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v uint32) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...uint32) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...uint32) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v uint32) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v uint32) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v uint32) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v uint32) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldSort, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldCode, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldRemark, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Position) predicate.Position {
	return predicate.Position(sql.NotPredicates(p))
}
