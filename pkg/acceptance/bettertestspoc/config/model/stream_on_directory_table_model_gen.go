// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type StreamOnDirectoryTableModel struct {
	Comment            tfconfig.Variable `json:"comment,omitempty"`
	CopyGrants         tfconfig.Variable `json:"copy_grants,omitempty"`
	Database           tfconfig.Variable `json:"database,omitempty"`
	FullyQualifiedName tfconfig.Variable `json:"fully_qualified_name,omitempty"`
	Name               tfconfig.Variable `json:"name,omitempty"`
	Schema             tfconfig.Variable `json:"schema,omitempty"`
	Stage              tfconfig.Variable `json:"stage,omitempty"`
	Stale              tfconfig.Variable `json:"stale,omitempty"`
	StreamType         tfconfig.Variable `json:"stream_type,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func StreamOnDirectoryTable(
	resourceName string,
	database string,
	name string,
	schema string,
	stage string,
) *StreamOnDirectoryTableModel {
	s := &StreamOnDirectoryTableModel{ResourceModelMeta: config.Meta(resourceName, resources.StreamOnDirectoryTable)}
	s.WithDatabase(database)
	s.WithName(name)
	s.WithSchema(schema)
	s.WithStage(stage)
	return s
}

func StreamOnDirectoryTableWithDefaultMeta(
	database string,
	name string,
	schema string,
	stage string,
) *StreamOnDirectoryTableModel {
	s := &StreamOnDirectoryTableModel{ResourceModelMeta: config.DefaultMeta(resources.StreamOnDirectoryTable)}
	s.WithDatabase(database)
	s.WithName(name)
	s.WithSchema(schema)
	s.WithStage(stage)
	return s
}

///////////////////////////////////////////////////////
// set proper json marshalling and handle depends on //
///////////////////////////////////////////////////////

func (s *StreamOnDirectoryTableModel) MarshalJSON() ([]byte, error) {
	type Alias StreamOnDirectoryTableModel
	return json.Marshal(&struct {
		*Alias
		DependsOn []string `json:"depends_on,omitempty"`
	}{
		Alias:     (*Alias)(s),
		DependsOn: s.DependsOn(),
	})
}

func (s *StreamOnDirectoryTableModel) WithDependsOn(values ...string) *StreamOnDirectoryTableModel {
	s.SetDependsOn(values...)
	return s
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (s *StreamOnDirectoryTableModel) WithComment(comment string) *StreamOnDirectoryTableModel {
	s.Comment = tfconfig.StringVariable(comment)
	return s
}

func (s *StreamOnDirectoryTableModel) WithCopyGrants(copyGrants bool) *StreamOnDirectoryTableModel {
	s.CopyGrants = tfconfig.BoolVariable(copyGrants)
	return s
}

func (s *StreamOnDirectoryTableModel) WithDatabase(database string) *StreamOnDirectoryTableModel {
	s.Database = tfconfig.StringVariable(database)
	return s
}

func (s *StreamOnDirectoryTableModel) WithFullyQualifiedName(fullyQualifiedName string) *StreamOnDirectoryTableModel {
	s.FullyQualifiedName = tfconfig.StringVariable(fullyQualifiedName)
	return s
}

func (s *StreamOnDirectoryTableModel) WithName(name string) *StreamOnDirectoryTableModel {
	s.Name = tfconfig.StringVariable(name)
	return s
}

func (s *StreamOnDirectoryTableModel) WithSchema(schema string) *StreamOnDirectoryTableModel {
	s.Schema = tfconfig.StringVariable(schema)
	return s
}

func (s *StreamOnDirectoryTableModel) WithStage(stage string) *StreamOnDirectoryTableModel {
	s.Stage = tfconfig.StringVariable(stage)
	return s
}

func (s *StreamOnDirectoryTableModel) WithStale(stale bool) *StreamOnDirectoryTableModel {
	s.Stale = tfconfig.BoolVariable(stale)
	return s
}

func (s *StreamOnDirectoryTableModel) WithStreamType(streamType string) *StreamOnDirectoryTableModel {
	s.StreamType = tfconfig.StringVariable(streamType)
	return s
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (s *StreamOnDirectoryTableModel) WithCommentValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.Comment = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithCopyGrantsValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.CopyGrants = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithDatabaseValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.Database = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithFullyQualifiedNameValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.FullyQualifiedName = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithNameValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.Name = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithSchemaValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.Schema = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithStageValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.Stage = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithStaleValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.Stale = value
	return s
}

func (s *StreamOnDirectoryTableModel) WithStreamTypeValue(value tfconfig.Variable) *StreamOnDirectoryTableModel {
	s.StreamType = value
	return s
}
