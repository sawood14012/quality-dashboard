// Code generated by ent, DO NOT EDIT.

package prowsuites

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLTE(FieldID, id))
}

// JobID applies equality check predicate on the "job_id" field. It's identical to JobIDEQ.
func JobID(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldJobID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldStatus, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldTime, v))
}

// JobIDEQ applies the EQ predicate on the "job_id" field.
func JobIDEQ(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldJobID, v))
}

// JobIDNEQ applies the NEQ predicate on the "job_id" field.
func JobIDNEQ(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNEQ(FieldJobID, v))
}

// JobIDIn applies the In predicate on the "job_id" field.
func JobIDIn(vs ...string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldIn(FieldJobID, vs...))
}

// JobIDNotIn applies the NotIn predicate on the "job_id" field.
func JobIDNotIn(vs ...string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNotIn(FieldJobID, vs...))
}

// JobIDGT applies the GT predicate on the "job_id" field.
func JobIDGT(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGT(FieldJobID, v))
}

// JobIDGTE applies the GTE predicate on the "job_id" field.
func JobIDGTE(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGTE(FieldJobID, v))
}

// JobIDLT applies the LT predicate on the "job_id" field.
func JobIDLT(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLT(FieldJobID, v))
}

// JobIDLTE applies the LTE predicate on the "job_id" field.
func JobIDLTE(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLTE(FieldJobID, v))
}

// JobIDContains applies the Contains predicate on the "job_id" field.
func JobIDContains(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldContains(FieldJobID, v))
}

// JobIDHasPrefix applies the HasPrefix predicate on the "job_id" field.
func JobIDHasPrefix(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldHasPrefix(FieldJobID, v))
}

// JobIDHasSuffix applies the HasSuffix predicate on the "job_id" field.
func JobIDHasSuffix(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldHasSuffix(FieldJobID, v))
}

// JobIDEqualFold applies the EqualFold predicate on the "job_id" field.
func JobIDEqualFold(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEqualFold(FieldJobID, v))
}

// JobIDContainsFold applies the ContainsFold predicate on the "job_id" field.
func JobIDContainsFold(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldContainsFold(FieldJobID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldContainsFold(FieldStatus, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v float64) predicate.ProwSuites {
	return predicate.ProwSuites(sql.FieldLTE(FieldTime, v))
}

// HasProwSuites applies the HasEdge predicate on the "prow_suites" edge.
func HasProwSuites() predicate.ProwSuites {
	return predicate.ProwSuites(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProwSuitesTable, ProwSuitesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProwSuitesWith applies the HasEdge predicate on the "prow_suites" edge with a given conditions (other predicates).
func HasProwSuitesWith(preds ...predicate.Repository) predicate.ProwSuites {
	return predicate.ProwSuites(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProwSuitesInverseTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProwSuitesTable, ProwSuitesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProwSuites) predicate.ProwSuites {
	return predicate.ProwSuites(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProwSuites) predicate.ProwSuites {
	return predicate.ProwSuites(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProwSuites) predicate.ProwSuites {
	return predicate.ProwSuites(func(s *sql.Selector) {
		p(s.Not())
	})
}
