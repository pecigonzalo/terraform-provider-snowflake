// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type DatabaseRoleResourceAssert struct {
	*assert.ResourceAssert
}

func DatabaseRoleResource(t *testing.T, name string) *DatabaseRoleResourceAssert {
	t.Helper()

	return &DatabaseRoleResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedDatabaseRoleResource(t *testing.T, id string) *DatabaseRoleResourceAssert {
	t.Helper()

	return &DatabaseRoleResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (d *DatabaseRoleResourceAssert) HasCommentString(expected string) *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueSet("comment", expected))
	return d
}

func (d *DatabaseRoleResourceAssert) HasDatabaseString(expected string) *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueSet("database", expected))
	return d
}

func (d *DatabaseRoleResourceAssert) HasFullyQualifiedNameString(expected string) *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return d
}

func (d *DatabaseRoleResourceAssert) HasNameString(expected string) *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueSet("name", expected))
	return d
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (d *DatabaseRoleResourceAssert) HasNoComment() *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueNotSet("comment"))
	return d
}

func (d *DatabaseRoleResourceAssert) HasNoDatabase() *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueNotSet("database"))
	return d
}

func (d *DatabaseRoleResourceAssert) HasNoFullyQualifiedName() *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return d
}

func (d *DatabaseRoleResourceAssert) HasNoName() *DatabaseRoleResourceAssert {
	d.AddAssertion(assert.ValueNotSet("name"))
	return d
}
