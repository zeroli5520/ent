// Code generated (@generated) by entc, DO NOT EDIT.

package card

import (
	"time"

	"github.com/facebookincubator/ent/examples/o2o2types/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// Expired applies equality check predicate on the "expired" field. It's identical to ExpiredEQ.
func Expired(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpired), v))
		},
	)
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
	)
}

// ExpiredEQ applies the EQ predicate on the "expired" field.
func ExpiredEQ(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpired), v))
		},
	)
}

// ExpiredNEQ applies the NEQ predicate on the "expired" field.
func ExpiredNEQ(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldExpired), v))
		},
	)
}

// ExpiredGT applies the GT predicate on the "expired" field.
func ExpiredGT(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldExpired), v))
		},
	)
}

// ExpiredGTE applies the GTE predicate on the "expired" field.
func ExpiredGTE(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldExpired), v))
		},
	)
}

// ExpiredLT applies the LT predicate on the "expired" field.
func ExpiredLT(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldExpired), v))
		},
	)
}

// ExpiredLTE applies the LTE predicate on the "expired" field.
func ExpiredLTE(v time.Time) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldExpired), v))
		},
	)
}

// ExpiredIn applies the In predicate on the "expired" field.
func ExpiredIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldExpired), v...))
		},
	)
}

// ExpiredNotIn applies the NotIn predicate on the "expired" field.
func ExpiredNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldExpired), v...))
		},
	)
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
	)
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNumber), v))
		},
	)
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNumber), v))
		},
	)
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNumber), v))
		},
	)
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNumber), v))
		},
	)
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNumber), v))
		},
	)
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNumber), v...))
		},
	)
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNumber), v...))
		},
	)
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldNumber), v))
		},
	)
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldNumber), v))
		},
	)
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldNumber), v))
		},
	)
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldNumber), v))
		},
	)
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldNumber), v))
		},
	)
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(OwnerColumn)))
		},
	)
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(OwnerInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(OwnerColumn), t2))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}