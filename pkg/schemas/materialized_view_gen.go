// Code generated by sdk-to-schema generator; DO NOT EDIT.

package schemas

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ShowMaterializedViewSchema represents output of SHOW query for the single MaterializedView.
var ShowMaterializedViewSchema = map[string]*schema.Schema{
	"created_on": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"reserved": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"database_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"schema_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cluster_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"rows": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"bytes": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"source_database_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"source_schema_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"source_table_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"refreshed_on": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"compacted_on": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"owner": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"invalid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"invalid_reason": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"behind_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"comment": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"text": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"is_secure": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"automatic_clustering": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"owner_role_type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"budget": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var _ = ShowMaterializedViewSchema

func MaterializedViewToSchema(materializedView *sdk.MaterializedView) map[string]any {
	materializedViewSchema := make(map[string]any)
	materializedViewSchema["created_on"] = materializedView.CreatedOn
	materializedViewSchema["name"] = materializedView.Name
	if materializedView.Reserved != nil {
		materializedViewSchema["reserved"] = materializedView.Reserved
	}
	materializedViewSchema["database_name"] = materializedView.DatabaseName
	materializedViewSchema["schema_name"] = materializedView.SchemaName
	materializedViewSchema["cluster_by"] = materializedView.ClusterBy
	materializedViewSchema["rows"] = materializedView.Rows
	materializedViewSchema["bytes"] = materializedView.Bytes
	materializedViewSchema["source_database_name"] = materializedView.SourceDatabaseName
	materializedViewSchema["source_schema_name"] = materializedView.SourceSchemaName
	materializedViewSchema["source_table_name"] = materializedView.SourceTableName
	materializedViewSchema["refreshed_on"] = materializedView.RefreshedOn.String()
	materializedViewSchema["compacted_on"] = materializedView.CompactedOn.String()
	materializedViewSchema["owner"] = materializedView.Owner
	materializedViewSchema["invalid"] = materializedView.Invalid
	materializedViewSchema["invalid_reason"] = materializedView.InvalidReason
	materializedViewSchema["behind_by"] = materializedView.BehindBy
	materializedViewSchema["comment"] = materializedView.Comment
	materializedViewSchema["text"] = materializedView.Text
	materializedViewSchema["is_secure"] = materializedView.IsSecure
	materializedViewSchema["automatic_clustering"] = materializedView.AutomaticClustering
	materializedViewSchema["owner_role_type"] = materializedView.OwnerRoleType
	materializedViewSchema["budget"] = materializedView.Budget
	return materializedViewSchema
}

var _ = MaterializedViewToSchema
