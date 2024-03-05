// Code generated by ent, DO NOT EDIT.

package tier

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Tier {
	return predicate.Tier(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Tier {
	return predicate.Tier(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Tier {
	return predicate.Tier(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Tier {
	return predicate.Tier(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Tier {
	return predicate.Tier(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Tier {
	return predicate.Tier(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Tier {
	return predicate.Tier(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Tier {
	return predicate.Tier(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Tier {
	return predicate.Tier(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldName, v))
}

// DiskMB applies equality check predicate on the "disk_mb" field. It's identical to DiskMBEQ.
func DiskMB(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldDiskMB, v))
}

// ConcurrentInstances applies equality check predicate on the "concurrent_instances" field. It's identical to ConcurrentInstancesEQ.
func ConcurrentInstances(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldConcurrentInstances, v))
}

// MaxLengthHours applies equality check predicate on the "max_length_hours" field. It's identical to MaxLengthHoursEQ.
func MaxLengthHours(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldMaxLengthHours, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tier {
	return predicate.Tier(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tier {
	return predicate.Tier(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tier {
	return predicate.Tier(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tier {
	return predicate.Tier(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tier {
	return predicate.Tier(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tier {
	return predicate.Tier(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tier {
	return predicate.Tier(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tier {
	return predicate.Tier(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tier {
	return predicate.Tier(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tier {
	return predicate.Tier(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tier {
	return predicate.Tier(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tier {
	return predicate.Tier(sql.FieldContainsFold(FieldName, v))
}

// DiskMBEQ applies the EQ predicate on the "disk_mb" field.
func DiskMBEQ(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldDiskMB, v))
}

// DiskMBNEQ applies the NEQ predicate on the "disk_mb" field.
func DiskMBNEQ(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldNEQ(FieldDiskMB, v))
}

// DiskMBIn applies the In predicate on the "disk_mb" field.
func DiskMBIn(vs ...int64) predicate.Tier {
	return predicate.Tier(sql.FieldIn(FieldDiskMB, vs...))
}

// DiskMBNotIn applies the NotIn predicate on the "disk_mb" field.
func DiskMBNotIn(vs ...int64) predicate.Tier {
	return predicate.Tier(sql.FieldNotIn(FieldDiskMB, vs...))
}

// DiskMBGT applies the GT predicate on the "disk_mb" field.
func DiskMBGT(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldGT(FieldDiskMB, v))
}

// DiskMBGTE applies the GTE predicate on the "disk_mb" field.
func DiskMBGTE(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldGTE(FieldDiskMB, v))
}

// DiskMBLT applies the LT predicate on the "disk_mb" field.
func DiskMBLT(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldLT(FieldDiskMB, v))
}

// DiskMBLTE applies the LTE predicate on the "disk_mb" field.
func DiskMBLTE(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldLTE(FieldDiskMB, v))
}

// ConcurrentInstancesEQ applies the EQ predicate on the "concurrent_instances" field.
func ConcurrentInstancesEQ(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldConcurrentInstances, v))
}

// ConcurrentInstancesNEQ applies the NEQ predicate on the "concurrent_instances" field.
func ConcurrentInstancesNEQ(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldNEQ(FieldConcurrentInstances, v))
}

// ConcurrentInstancesIn applies the In predicate on the "concurrent_instances" field.
func ConcurrentInstancesIn(vs ...int64) predicate.Tier {
	return predicate.Tier(sql.FieldIn(FieldConcurrentInstances, vs...))
}

// ConcurrentInstancesNotIn applies the NotIn predicate on the "concurrent_instances" field.
func ConcurrentInstancesNotIn(vs ...int64) predicate.Tier {
	return predicate.Tier(sql.FieldNotIn(FieldConcurrentInstances, vs...))
}

// ConcurrentInstancesGT applies the GT predicate on the "concurrent_instances" field.
func ConcurrentInstancesGT(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldGT(FieldConcurrentInstances, v))
}

// ConcurrentInstancesGTE applies the GTE predicate on the "concurrent_instances" field.
func ConcurrentInstancesGTE(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldGTE(FieldConcurrentInstances, v))
}

// ConcurrentInstancesLT applies the LT predicate on the "concurrent_instances" field.
func ConcurrentInstancesLT(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldLT(FieldConcurrentInstances, v))
}

// ConcurrentInstancesLTE applies the LTE predicate on the "concurrent_instances" field.
func ConcurrentInstancesLTE(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldLTE(FieldConcurrentInstances, v))
}

// MaxLengthHoursEQ applies the EQ predicate on the "max_length_hours" field.
func MaxLengthHoursEQ(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldEQ(FieldMaxLengthHours, v))
}

// MaxLengthHoursNEQ applies the NEQ predicate on the "max_length_hours" field.
func MaxLengthHoursNEQ(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldNEQ(FieldMaxLengthHours, v))
}

// MaxLengthHoursIn applies the In predicate on the "max_length_hours" field.
func MaxLengthHoursIn(vs ...int64) predicate.Tier {
	return predicate.Tier(sql.FieldIn(FieldMaxLengthHours, vs...))
}

// MaxLengthHoursNotIn applies the NotIn predicate on the "max_length_hours" field.
func MaxLengthHoursNotIn(vs ...int64) predicate.Tier {
	return predicate.Tier(sql.FieldNotIn(FieldMaxLengthHours, vs...))
}

// MaxLengthHoursGT applies the GT predicate on the "max_length_hours" field.
func MaxLengthHoursGT(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldGT(FieldMaxLengthHours, v))
}

// MaxLengthHoursGTE applies the GTE predicate on the "max_length_hours" field.
func MaxLengthHoursGTE(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldGTE(FieldMaxLengthHours, v))
}

// MaxLengthHoursLT applies the LT predicate on the "max_length_hours" field.
func MaxLengthHoursLT(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldLT(FieldMaxLengthHours, v))
}

// MaxLengthHoursLTE applies the LTE predicate on the "max_length_hours" field.
func MaxLengthHoursLTE(v int64) predicate.Tier {
	return predicate.Tier(sql.FieldLTE(FieldMaxLengthHours, v))
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.Tier {
	return predicate.Tier(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeamsTable, TeamsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Team
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.Tier {
	return predicate.Tier(func(s *sql.Selector) {
		step := newTeamsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Team
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tier) predicate.Tier {
	return predicate.Tier(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tier) predicate.Tier {
	return predicate.Tier(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tier) predicate.Tier {
	return predicate.Tier(sql.NotPredicates(p))
}
