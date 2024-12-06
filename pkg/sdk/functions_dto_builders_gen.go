// Code generated by dto builder generator; DO NOT EDIT.

package sdk

// imports edited manually
import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/datatypes"
)

func NewCreateForJavaFunctionRequest(
	name SchemaObjectIdentifier,
	Returns FunctionReturnsRequest,
	Handler string,
) *CreateForJavaFunctionRequest {
	s := CreateForJavaFunctionRequest{}
	s.name = name
	s.Returns = Returns
	s.Handler = Handler
	return &s
}

func (s *CreateForJavaFunctionRequest) WithOrReplace(OrReplace bool) *CreateForJavaFunctionRequest {
	s.OrReplace = &OrReplace
	return s
}

func (s *CreateForJavaFunctionRequest) WithTemporary(Temporary bool) *CreateForJavaFunctionRequest {
	s.Temporary = &Temporary
	return s
}

func (s *CreateForJavaFunctionRequest) WithSecure(Secure bool) *CreateForJavaFunctionRequest {
	s.Secure = &Secure
	return s
}

func (s *CreateForJavaFunctionRequest) WithIfNotExists(IfNotExists bool) *CreateForJavaFunctionRequest {
	s.IfNotExists = &IfNotExists
	return s
}

func (s *CreateForJavaFunctionRequest) WithArguments(Arguments []FunctionArgumentRequest) *CreateForJavaFunctionRequest {
	s.Arguments = Arguments
	return s
}

func (s *CreateForJavaFunctionRequest) WithCopyGrants(CopyGrants bool) *CreateForJavaFunctionRequest {
	s.CopyGrants = &CopyGrants
	return s
}

func (s *CreateForJavaFunctionRequest) WithReturnNullValues(ReturnNullValues ReturnNullValues) *CreateForJavaFunctionRequest {
	s.ReturnNullValues = &ReturnNullValues
	return s
}

func (s *CreateForJavaFunctionRequest) WithNullInputBehavior(NullInputBehavior NullInputBehavior) *CreateForJavaFunctionRequest {
	s.NullInputBehavior = &NullInputBehavior
	return s
}

func (s *CreateForJavaFunctionRequest) WithReturnResultsBehavior(ReturnResultsBehavior ReturnResultsBehavior) *CreateForJavaFunctionRequest {
	s.ReturnResultsBehavior = &ReturnResultsBehavior
	return s
}

func (s *CreateForJavaFunctionRequest) WithRuntimeVersion(RuntimeVersion string) *CreateForJavaFunctionRequest {
	s.RuntimeVersion = &RuntimeVersion
	return s
}

func (s *CreateForJavaFunctionRequest) WithComment(Comment string) *CreateForJavaFunctionRequest {
	s.Comment = &Comment
	return s
}

func (s *CreateForJavaFunctionRequest) WithImports(Imports []FunctionImportRequest) *CreateForJavaFunctionRequest {
	s.Imports = Imports
	return s
}

func (s *CreateForJavaFunctionRequest) WithPackages(Packages []FunctionPackageRequest) *CreateForJavaFunctionRequest {
	s.Packages = Packages
	return s
}

func (s *CreateForJavaFunctionRequest) WithExternalAccessIntegrations(ExternalAccessIntegrations []AccountObjectIdentifier) *CreateForJavaFunctionRequest {
	s.ExternalAccessIntegrations = ExternalAccessIntegrations
	return s
}

func (s *CreateForJavaFunctionRequest) WithSecrets(Secrets []SecretReference) *CreateForJavaFunctionRequest {
	s.Secrets = Secrets
	return s
}

func (s *CreateForJavaFunctionRequest) WithTargetPath(TargetPath string) *CreateForJavaFunctionRequest {
	s.TargetPath = &TargetPath
	return s
}

func (s *CreateForJavaFunctionRequest) WithFunctionDefinition(FunctionDefinition string) *CreateForJavaFunctionRequest {
	s.FunctionDefinition = &FunctionDefinition
	return s
}

func NewFunctionArgumentRequest(
	ArgName string,
	ArgDataType datatypes.DataType,
) *FunctionArgumentRequest {
	s := FunctionArgumentRequest{}
	s.ArgName = ArgName
	s.ArgDataType = ArgDataType
	return &s
}

func (s *FunctionArgumentRequest) WithArgDataTypeOld(ArgDataTypeOld DataType) *FunctionArgumentRequest {
	s.ArgDataTypeOld = ArgDataTypeOld
	return s
}

func (s *FunctionArgumentRequest) WithDefaultValue(DefaultValue string) *FunctionArgumentRequest {
	s.DefaultValue = &DefaultValue
	return s
}

func NewFunctionReturnsRequest() *FunctionReturnsRequest {
	return &FunctionReturnsRequest{}
}

func (s *FunctionReturnsRequest) WithResultDataType(ResultDataType FunctionReturnsResultDataTypeRequest) *FunctionReturnsRequest {
	s.ResultDataType = &ResultDataType
	return s
}

func (s *FunctionReturnsRequest) WithTable(Table FunctionReturnsTableRequest) *FunctionReturnsRequest {
	s.Table = &Table
	return s
}

func NewFunctionReturnsResultDataTypeRequest(
	ResultDataType datatypes.DataType,
) *FunctionReturnsResultDataTypeRequest {
	s := FunctionReturnsResultDataTypeRequest{}
	s.ResultDataType = ResultDataType
	return &s
}

func (s *FunctionReturnsResultDataTypeRequest) WithResultDataTypeOld(ResultDataTypeOld DataType) *FunctionReturnsResultDataTypeRequest {
	s.ResultDataTypeOld = ResultDataTypeOld
	return s
}

func NewFunctionReturnsTableRequest() *FunctionReturnsTableRequest {
	return &FunctionReturnsTableRequest{}
}

func (s *FunctionReturnsTableRequest) WithColumns(Columns []FunctionColumnRequest) *FunctionReturnsTableRequest {
	s.Columns = Columns
	return s
}

func NewFunctionColumnRequest(
	ColumnName string,
	ColumnDataType datatypes.DataType,
) *FunctionColumnRequest {
	s := FunctionColumnRequest{}
	s.ColumnName = ColumnName
	s.ColumnDataType = ColumnDataType
	return &s
}

func (s *FunctionColumnRequest) WithColumnDataTypeOld(ColumnDataTypeOld DataType) *FunctionColumnRequest {
	s.ColumnDataTypeOld = ColumnDataTypeOld
	return s
}

func NewFunctionImportRequest() *FunctionImportRequest {
	return &FunctionImportRequest{}
}

func (s *FunctionImportRequest) WithImport(Import string) *FunctionImportRequest {
	s.Import = Import
	return s
}

func NewFunctionPackageRequest() *FunctionPackageRequest {
	return &FunctionPackageRequest{}
}

func (s *FunctionPackageRequest) WithPackage(Package string) *FunctionPackageRequest {
	s.Package = Package
	return s
}

func NewCreateForJavascriptFunctionRequest(
	name SchemaObjectIdentifier,
	Returns FunctionReturnsRequest,
	FunctionDefinition string,
) *CreateForJavascriptFunctionRequest {
	s := CreateForJavascriptFunctionRequest{}
	s.name = name
	s.Returns = Returns
	s.FunctionDefinition = FunctionDefinition
	return &s
}

func (s *CreateForJavascriptFunctionRequest) WithOrReplace(OrReplace bool) *CreateForJavascriptFunctionRequest {
	s.OrReplace = &OrReplace
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithTemporary(Temporary bool) *CreateForJavascriptFunctionRequest {
	s.Temporary = &Temporary
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithSecure(Secure bool) *CreateForJavascriptFunctionRequest {
	s.Secure = &Secure
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithArguments(Arguments []FunctionArgumentRequest) *CreateForJavascriptFunctionRequest {
	s.Arguments = Arguments
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithCopyGrants(CopyGrants bool) *CreateForJavascriptFunctionRequest {
	s.CopyGrants = &CopyGrants
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithReturnNullValues(ReturnNullValues ReturnNullValues) *CreateForJavascriptFunctionRequest {
	s.ReturnNullValues = &ReturnNullValues
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithNullInputBehavior(NullInputBehavior NullInputBehavior) *CreateForJavascriptFunctionRequest {
	s.NullInputBehavior = &NullInputBehavior
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithReturnResultsBehavior(ReturnResultsBehavior ReturnResultsBehavior) *CreateForJavascriptFunctionRequest {
	s.ReturnResultsBehavior = &ReturnResultsBehavior
	return s
}

func (s *CreateForJavascriptFunctionRequest) WithComment(Comment string) *CreateForJavascriptFunctionRequest {
	s.Comment = &Comment
	return s
}

func NewCreateForPythonFunctionRequest(
	name SchemaObjectIdentifier,
	Returns FunctionReturnsRequest,
	RuntimeVersion string,
	Handler string,
) *CreateForPythonFunctionRequest {
	s := CreateForPythonFunctionRequest{}
	s.name = name
	s.Returns = Returns
	s.RuntimeVersion = RuntimeVersion
	s.Handler = Handler
	return &s
}

func (s *CreateForPythonFunctionRequest) WithOrReplace(OrReplace bool) *CreateForPythonFunctionRequest {
	s.OrReplace = &OrReplace
	return s
}

func (s *CreateForPythonFunctionRequest) WithTemporary(Temporary bool) *CreateForPythonFunctionRequest {
	s.Temporary = &Temporary
	return s
}

func (s *CreateForPythonFunctionRequest) WithSecure(Secure bool) *CreateForPythonFunctionRequest {
	s.Secure = &Secure
	return s
}

func (s *CreateForPythonFunctionRequest) WithIfNotExists(IfNotExists bool) *CreateForPythonFunctionRequest {
	s.IfNotExists = &IfNotExists
	return s
}

func (s *CreateForPythonFunctionRequest) WithArguments(Arguments []FunctionArgumentRequest) *CreateForPythonFunctionRequest {
	s.Arguments = Arguments
	return s
}

func (s *CreateForPythonFunctionRequest) WithCopyGrants(CopyGrants bool) *CreateForPythonFunctionRequest {
	s.CopyGrants = &CopyGrants
	return s
}

func (s *CreateForPythonFunctionRequest) WithReturnNullValues(ReturnNullValues ReturnNullValues) *CreateForPythonFunctionRequest {
	s.ReturnNullValues = &ReturnNullValues
	return s
}

func (s *CreateForPythonFunctionRequest) WithNullInputBehavior(NullInputBehavior NullInputBehavior) *CreateForPythonFunctionRequest {
	s.NullInputBehavior = &NullInputBehavior
	return s
}

func (s *CreateForPythonFunctionRequest) WithReturnResultsBehavior(ReturnResultsBehavior ReturnResultsBehavior) *CreateForPythonFunctionRequest {
	s.ReturnResultsBehavior = &ReturnResultsBehavior
	return s
}

func (s *CreateForPythonFunctionRequest) WithComment(Comment string) *CreateForPythonFunctionRequest {
	s.Comment = &Comment
	return s
}

func (s *CreateForPythonFunctionRequest) WithImports(Imports []FunctionImportRequest) *CreateForPythonFunctionRequest {
	s.Imports = Imports
	return s
}

func (s *CreateForPythonFunctionRequest) WithPackages(Packages []FunctionPackageRequest) *CreateForPythonFunctionRequest {
	s.Packages = Packages
	return s
}

func (s *CreateForPythonFunctionRequest) WithExternalAccessIntegrations(ExternalAccessIntegrations []AccountObjectIdentifier) *CreateForPythonFunctionRequest {
	s.ExternalAccessIntegrations = ExternalAccessIntegrations
	return s
}

func (s *CreateForPythonFunctionRequest) WithSecrets(Secrets []SecretReference) *CreateForPythonFunctionRequest {
	s.Secrets = Secrets
	return s
}

func (s *CreateForPythonFunctionRequest) WithFunctionDefinition(FunctionDefinition string) *CreateForPythonFunctionRequest {
	s.FunctionDefinition = &FunctionDefinition
	return s
}

func NewCreateForScalaFunctionRequest(
	name SchemaObjectIdentifier,
	ResultDataType datatypes.DataType,
	Handler string,
) *CreateForScalaFunctionRequest {
	s := CreateForScalaFunctionRequest{}
	s.name = name
	s.ResultDataType = ResultDataType
	s.Handler = Handler
	return &s
}

func (s *CreateForScalaFunctionRequest) WithOrReplace(OrReplace bool) *CreateForScalaFunctionRequest {
	s.OrReplace = &OrReplace
	return s
}

func (s *CreateForScalaFunctionRequest) WithTemporary(Temporary bool) *CreateForScalaFunctionRequest {
	s.Temporary = &Temporary
	return s
}

func (s *CreateForScalaFunctionRequest) WithSecure(Secure bool) *CreateForScalaFunctionRequest {
	s.Secure = &Secure
	return s
}

func (s *CreateForScalaFunctionRequest) WithIfNotExists(IfNotExists bool) *CreateForScalaFunctionRequest {
	s.IfNotExists = &IfNotExists
	return s
}

func (s *CreateForScalaFunctionRequest) WithArguments(Arguments []FunctionArgumentRequest) *CreateForScalaFunctionRequest {
	s.Arguments = Arguments
	return s
}

func (s *CreateForScalaFunctionRequest) WithCopyGrants(CopyGrants bool) *CreateForScalaFunctionRequest {
	s.CopyGrants = &CopyGrants
	return s
}

func (s *CreateForScalaFunctionRequest) WithResultDataTypeOld(ResultDataTypeOld DataType) *CreateForScalaFunctionRequest {
	s.ResultDataTypeOld = ResultDataTypeOld
	return s
}

func (s *CreateForScalaFunctionRequest) WithReturnNullValues(ReturnNullValues ReturnNullValues) *CreateForScalaFunctionRequest {
	s.ReturnNullValues = &ReturnNullValues
	return s
}

func (s *CreateForScalaFunctionRequest) WithNullInputBehavior(NullInputBehavior NullInputBehavior) *CreateForScalaFunctionRequest {
	s.NullInputBehavior = &NullInputBehavior
	return s
}

func (s *CreateForScalaFunctionRequest) WithReturnResultsBehavior(ReturnResultsBehavior ReturnResultsBehavior) *CreateForScalaFunctionRequest {
	s.ReturnResultsBehavior = &ReturnResultsBehavior
	return s
}

func (s *CreateForScalaFunctionRequest) WithRuntimeVersion(RuntimeVersion string) *CreateForScalaFunctionRequest {
	s.RuntimeVersion = &RuntimeVersion
	return s
}

func (s *CreateForScalaFunctionRequest) WithComment(Comment string) *CreateForScalaFunctionRequest {
	s.Comment = &Comment
	return s
}

func (s *CreateForScalaFunctionRequest) WithImports(Imports []FunctionImportRequest) *CreateForScalaFunctionRequest {
	s.Imports = Imports
	return s
}

func (s *CreateForScalaFunctionRequest) WithPackages(Packages []FunctionPackageRequest) *CreateForScalaFunctionRequest {
	s.Packages = Packages
	return s
}

func (s *CreateForScalaFunctionRequest) WithTargetPath(TargetPath string) *CreateForScalaFunctionRequest {
	s.TargetPath = &TargetPath
	return s
}

func (s *CreateForScalaFunctionRequest) WithFunctionDefinition(FunctionDefinition string) *CreateForScalaFunctionRequest {
	s.FunctionDefinition = &FunctionDefinition
	return s
}

func NewCreateForSQLFunctionRequest(
	name SchemaObjectIdentifier,
	Returns FunctionReturnsRequest,
	FunctionDefinition string,
) *CreateForSQLFunctionRequest {
	s := CreateForSQLFunctionRequest{}
	s.name = name
	s.Returns = Returns
	s.FunctionDefinition = FunctionDefinition
	return &s
}

func (s *CreateForSQLFunctionRequest) WithOrReplace(OrReplace bool) *CreateForSQLFunctionRequest {
	s.OrReplace = &OrReplace
	return s
}

func (s *CreateForSQLFunctionRequest) WithTemporary(Temporary bool) *CreateForSQLFunctionRequest {
	s.Temporary = &Temporary
	return s
}

func (s *CreateForSQLFunctionRequest) WithSecure(Secure bool) *CreateForSQLFunctionRequest {
	s.Secure = &Secure
	return s
}

func (s *CreateForSQLFunctionRequest) WithArguments(Arguments []FunctionArgumentRequest) *CreateForSQLFunctionRequest {
	s.Arguments = Arguments
	return s
}

func (s *CreateForSQLFunctionRequest) WithCopyGrants(CopyGrants bool) *CreateForSQLFunctionRequest {
	s.CopyGrants = &CopyGrants
	return s
}

func (s *CreateForSQLFunctionRequest) WithReturnNullValues(ReturnNullValues ReturnNullValues) *CreateForSQLFunctionRequest {
	s.ReturnNullValues = &ReturnNullValues
	return s
}

func (s *CreateForSQLFunctionRequest) WithReturnResultsBehavior(ReturnResultsBehavior ReturnResultsBehavior) *CreateForSQLFunctionRequest {
	s.ReturnResultsBehavior = &ReturnResultsBehavior
	return s
}

func (s *CreateForSQLFunctionRequest) WithMemoizable(Memoizable bool) *CreateForSQLFunctionRequest {
	s.Memoizable = &Memoizable
	return s
}

func (s *CreateForSQLFunctionRequest) WithComment(Comment string) *CreateForSQLFunctionRequest {
	s.Comment = &Comment
	return s
}

func NewAlterFunctionRequest(
	name SchemaObjectIdentifierWithArguments,
) *AlterFunctionRequest {
	s := AlterFunctionRequest{}
	s.name = name
	return &s
}

func (s *AlterFunctionRequest) WithIfExists(IfExists bool) *AlterFunctionRequest {
	s.IfExists = &IfExists
	return s
}

func (s *AlterFunctionRequest) WithRenameTo(RenameTo SchemaObjectIdentifier) *AlterFunctionRequest {
	s.RenameTo = &RenameTo
	return s
}

func (s *AlterFunctionRequest) WithSetComment(SetComment string) *AlterFunctionRequest {
	s.SetComment = &SetComment
	return s
}

func (s *AlterFunctionRequest) WithSetLogLevel(SetLogLevel string) *AlterFunctionRequest {
	s.SetLogLevel = &SetLogLevel
	return s
}

func (s *AlterFunctionRequest) WithSetTraceLevel(SetTraceLevel string) *AlterFunctionRequest {
	s.SetTraceLevel = &SetTraceLevel
	return s
}

func (s *AlterFunctionRequest) WithSetSecure(SetSecure bool) *AlterFunctionRequest {
	s.SetSecure = &SetSecure
	return s
}

func (s *AlterFunctionRequest) WithUnsetSecure(UnsetSecure bool) *AlterFunctionRequest {
	s.UnsetSecure = &UnsetSecure
	return s
}

func (s *AlterFunctionRequest) WithUnsetLogLevel(UnsetLogLevel bool) *AlterFunctionRequest {
	s.UnsetLogLevel = &UnsetLogLevel
	return s
}

func (s *AlterFunctionRequest) WithUnsetTraceLevel(UnsetTraceLevel bool) *AlterFunctionRequest {
	s.UnsetTraceLevel = &UnsetTraceLevel
	return s
}

func (s *AlterFunctionRequest) WithUnsetComment(UnsetComment bool) *AlterFunctionRequest {
	s.UnsetComment = &UnsetComment
	return s
}

func (s *AlterFunctionRequest) WithSetTags(SetTags []TagAssociation) *AlterFunctionRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterFunctionRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterFunctionRequest {
	s.UnsetTags = UnsetTags
	return s
}

func NewDropFunctionRequest(
	name SchemaObjectIdentifierWithArguments,
) *DropFunctionRequest {
	s := DropFunctionRequest{}
	s.name = name
	return &s
}

func (s *DropFunctionRequest) WithIfExists(IfExists bool) *DropFunctionRequest {
	s.IfExists = &IfExists
	return s
}

func NewShowFunctionRequest() *ShowFunctionRequest {
	return &ShowFunctionRequest{}
}

func (s *ShowFunctionRequest) WithLike(Like Like) *ShowFunctionRequest {
	s.Like = &Like
	return s
}

func (s *ShowFunctionRequest) WithIn(In In) *ShowFunctionRequest {
	s.In = &In
	return s
}

func NewDescribeFunctionRequest(
	name SchemaObjectIdentifierWithArguments,
) *DescribeFunctionRequest {
	s := DescribeFunctionRequest{}
	s.name = name
	return &s
}
