// Code generated by ent, DO NOT EDIT.

package monthlysubscription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldIsActive, v))
}

// Paid applies equality check predicate on the "paid" field. It's identical to PaidEQ.
func Paid(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldPaid, v))
}

// IsTrial applies equality check predicate on the "is_trial" field. It's identical to IsTrialEQ.
func IsTrial(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldIsTrial, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldStartedAt, v))
}

// ExpiredOn applies equality check predicate on the "expired_on" field. It's identical to ExpiredOnEQ.
func ExpiredOn(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldExpiredOn, v))
}

// CancelledAt applies equality check predicate on the "cancelled_at" field. It's identical to CancelledAtEQ.
func CancelledAt(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldCancelledAt, v))
}

// PayingProfileID applies equality check predicate on the "paying_profile_id" field. It's identical to PayingProfileIDEQ.
func PayingProfileID(v int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldPayingProfileID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLTE(FieldUpdatedAt, v))
}

// ProductEQ applies the EQ predicate on the "product" field.
func ProductEQ(v Product) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldProduct, v))
}

// ProductNEQ applies the NEQ predicate on the "product" field.
func ProductNEQ(v Product) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldProduct, v))
}

// ProductIn applies the In predicate on the "product" field.
func ProductIn(vs ...Product) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldProduct, vs...))
}

// ProductNotIn applies the NotIn predicate on the "product" field.
func ProductNotIn(vs ...Product) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldProduct, vs...))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldIsActive, v))
}

// PaidEQ applies the EQ predicate on the "paid" field.
func PaidEQ(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldPaid, v))
}

// PaidNEQ applies the NEQ predicate on the "paid" field.
func PaidNEQ(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldPaid, v))
}

// IsTrialEQ applies the EQ predicate on the "is_trial" field.
func IsTrialEQ(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldIsTrial, v))
}

// IsTrialNEQ applies the NEQ predicate on the "is_trial" field.
func IsTrialNEQ(v bool) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldIsTrial, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLTE(FieldStartedAt, v))
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIsNull(FieldStartedAt))
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotNull(FieldStartedAt))
}

// ExpiredOnEQ applies the EQ predicate on the "expired_on" field.
func ExpiredOnEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldExpiredOn, v))
}

// ExpiredOnNEQ applies the NEQ predicate on the "expired_on" field.
func ExpiredOnNEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldExpiredOn, v))
}

// ExpiredOnIn applies the In predicate on the "expired_on" field.
func ExpiredOnIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldExpiredOn, vs...))
}

// ExpiredOnNotIn applies the NotIn predicate on the "expired_on" field.
func ExpiredOnNotIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldExpiredOn, vs...))
}

// ExpiredOnGT applies the GT predicate on the "expired_on" field.
func ExpiredOnGT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGT(FieldExpiredOn, v))
}

// ExpiredOnGTE applies the GTE predicate on the "expired_on" field.
func ExpiredOnGTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGTE(FieldExpiredOn, v))
}

// ExpiredOnLT applies the LT predicate on the "expired_on" field.
func ExpiredOnLT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLT(FieldExpiredOn, v))
}

// ExpiredOnLTE applies the LTE predicate on the "expired_on" field.
func ExpiredOnLTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLTE(FieldExpiredOn, v))
}

// ExpiredOnIsNil applies the IsNil predicate on the "expired_on" field.
func ExpiredOnIsNil() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIsNull(FieldExpiredOn))
}

// ExpiredOnNotNil applies the NotNil predicate on the "expired_on" field.
func ExpiredOnNotNil() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotNull(FieldExpiredOn))
}

// CancelledAtEQ applies the EQ predicate on the "cancelled_at" field.
func CancelledAtEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldCancelledAt, v))
}

// CancelledAtNEQ applies the NEQ predicate on the "cancelled_at" field.
func CancelledAtNEQ(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldCancelledAt, v))
}

// CancelledAtIn applies the In predicate on the "cancelled_at" field.
func CancelledAtIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldCancelledAt, vs...))
}

// CancelledAtNotIn applies the NotIn predicate on the "cancelled_at" field.
func CancelledAtNotIn(vs ...time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldCancelledAt, vs...))
}

// CancelledAtGT applies the GT predicate on the "cancelled_at" field.
func CancelledAtGT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGT(FieldCancelledAt, v))
}

// CancelledAtGTE applies the GTE predicate on the "cancelled_at" field.
func CancelledAtGTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldGTE(FieldCancelledAt, v))
}

// CancelledAtLT applies the LT predicate on the "cancelled_at" field.
func CancelledAtLT(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLT(FieldCancelledAt, v))
}

// CancelledAtLTE applies the LTE predicate on the "cancelled_at" field.
func CancelledAtLTE(v time.Time) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldLTE(FieldCancelledAt, v))
}

// CancelledAtIsNil applies the IsNil predicate on the "cancelled_at" field.
func CancelledAtIsNil() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIsNull(FieldCancelledAt))
}

// CancelledAtNotNil applies the NotNil predicate on the "cancelled_at" field.
func CancelledAtNotNil() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotNull(FieldCancelledAt))
}

// PayingProfileIDEQ applies the EQ predicate on the "paying_profile_id" field.
func PayingProfileIDEQ(v int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldEQ(FieldPayingProfileID, v))
}

// PayingProfileIDNEQ applies the NEQ predicate on the "paying_profile_id" field.
func PayingProfileIDNEQ(v int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNEQ(FieldPayingProfileID, v))
}

// PayingProfileIDIn applies the In predicate on the "paying_profile_id" field.
func PayingProfileIDIn(vs ...int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldIn(FieldPayingProfileID, vs...))
}

// PayingProfileIDNotIn applies the NotIn predicate on the "paying_profile_id" field.
func PayingProfileIDNotIn(vs ...int) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.FieldNotIn(FieldPayingProfileID, vs...))
}

// HasBenefactors applies the HasEdge predicate on the "benefactors" edge.
func HasBenefactors() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, BenefactorsTable, BenefactorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBenefactorsWith applies the HasEdge predicate on the "benefactors" edge with a given conditions (other predicates).
func HasBenefactorsWith(preds ...predicate.Profile) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(func(s *sql.Selector) {
		step := newBenefactorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayer applies the HasEdge predicate on the "payer" edge.
func HasPayer() predicate.MonthlySubscription {
	return predicate.MonthlySubscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PayerTable, PayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPayerWith applies the HasEdge predicate on the "payer" edge with a given conditions (other predicates).
func HasPayerWith(preds ...predicate.Profile) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(func(s *sql.Selector) {
		step := newPayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MonthlySubscription) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MonthlySubscription) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MonthlySubscription) predicate.MonthlySubscription {
	return predicate.MonthlySubscription(sql.NotPredicates(p))
}