// Code generated by assertions generator; DO NOT EDIT.

package objectassert

import (
	"fmt"
	"testing"
	"time"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type DatabaseAssert struct {
	*assert.SnowflakeObjectAssert[sdk.Database, sdk.AccountObjectIdentifier]
}

func Database(t *testing.T, id sdk.AccountObjectIdentifier) *DatabaseAssert {
	t.Helper()
	return &DatabaseAssert{
		assert.NewSnowflakeObjectAssertWithProvider(sdk.ObjectTypeDatabase, id, acc.TestClient().Database.Show),
	}
}

func DatabaseFromObject(t *testing.T, database *sdk.Database) *DatabaseAssert {
	t.Helper()
	return &DatabaseAssert{
		assert.NewSnowflakeObjectAssertWithObject(sdk.ObjectTypeDatabase, database.ID(), database),
	}
}

func (d *DatabaseAssert) HasCreatedOn(expected time.Time) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.CreatedOn != expected {
			return fmt.Errorf("expected created on: %v; got: %v", expected, o.CreatedOn)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasName(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Name != expected {
			return fmt.Errorf("expected name: %v; got: %v", expected, o.Name)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasIsDefault(expected bool) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.IsDefault != expected {
			return fmt.Errorf("expected is default: %v; got: %v", expected, o.IsDefault)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasIsCurrent(expected bool) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.IsCurrent != expected {
			return fmt.Errorf("expected is current: %v; got: %v", expected, o.IsCurrent)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasOrigin(expected sdk.ExternalObjectIdentifier) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Origin == nil {
			return fmt.Errorf("expected origin to have value; got: nil")
		}
		if *o.Origin != expected {
			return fmt.Errorf("expected origin: %v; got: %v", expected, *o.Origin)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasOwner(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Owner != expected {
			return fmt.Errorf("expected owner: %v; got: %v", expected, o.Owner)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasComment(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Comment != expected {
			return fmt.Errorf("expected comment: %v; got: %v", expected, o.Comment)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasOptions(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Options != expected {
			return fmt.Errorf("expected options: %v; got: %v", expected, o.Options)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasRetentionTime(expected int) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.RetentionTime != expected {
			return fmt.Errorf("expected retention time: %v; got: %v", expected, o.RetentionTime)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasResourceGroup(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.ResourceGroup != expected {
			return fmt.Errorf("expected resource group: %v; got: %v", expected, o.ResourceGroup)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasDroppedOn(expected time.Time) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.DroppedOn != expected {
			return fmt.Errorf("expected dropped on: %v; got: %v", expected, o.DroppedOn)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasTransient(expected bool) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Transient != expected {
			return fmt.Errorf("expected transient: %v; got: %v", expected, o.Transient)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasKind(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.Kind != expected {
			return fmt.Errorf("expected kind: %v; got: %v", expected, o.Kind)
		}
		return nil
	})
	return d
}

func (d *DatabaseAssert) HasOwnerRoleType(expected string) *DatabaseAssert {
	d.AddAssertion(func(t *testing.T, o *sdk.Database) error {
		t.Helper()
		if o.OwnerRoleType != expected {
			return fmt.Errorf("expected owner role type: %v; got: %v", expected, o.OwnerRoleType)
		}
		return nil
	})
	return d
}
