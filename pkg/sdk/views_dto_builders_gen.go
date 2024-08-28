// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewCreateViewRequest(
	name SchemaObjectIdentifier,
	sql string,
) *CreateViewRequest {
	s := CreateViewRequest{}
	s.name = name
	s.sql = sql
	return &s
}

func (s *CreateViewRequest) WithOrReplace(OrReplace bool) *CreateViewRequest {
	s.OrReplace = &OrReplace
	return s
}

func (s *CreateViewRequest) WithSecure(Secure bool) *CreateViewRequest {
	s.Secure = &Secure
	return s
}

func (s *CreateViewRequest) WithTemporary(Temporary bool) *CreateViewRequest {
	s.Temporary = &Temporary
	return s
}

func (s *CreateViewRequest) WithRecursive(Recursive bool) *CreateViewRequest {
	s.Recursive = &Recursive
	return s
}

func (s *CreateViewRequest) WithIfNotExists(IfNotExists bool) *CreateViewRequest {
	s.IfNotExists = &IfNotExists
	return s
}

func (s *CreateViewRequest) WithColumns(Columns []ViewColumnRequest) *CreateViewRequest {
	s.Columns = Columns
	return s
}

func (s *CreateViewRequest) WithCopyGrants(CopyGrants bool) *CreateViewRequest {
	s.CopyGrants = &CopyGrants
	return s
}

func (s *CreateViewRequest) WithComment(Comment string) *CreateViewRequest {
	s.Comment = &Comment
	return s
}

func (s *CreateViewRequest) WithRowAccessPolicy(RowAccessPolicy ViewRowAccessPolicyRequest) *CreateViewRequest {
	s.RowAccessPolicy = &RowAccessPolicy
	return s
}

func (s *CreateViewRequest) WithAggregationPolicy(AggregationPolicy ViewAggregationPolicyRequest) *CreateViewRequest {
	s.AggregationPolicy = &AggregationPolicy
	return s
}

func (s *CreateViewRequest) WithTag(Tag []TagAssociation) *CreateViewRequest {
	s.Tag = Tag
	return s
}

func NewViewColumnRequest(
	Name string,
) *ViewColumnRequest {
	s := ViewColumnRequest{}
	s.Name = Name
	return &s
}

func (s *ViewColumnRequest) WithProjectionPolicy(ProjectionPolicy ViewColumnProjectionPolicyRequest) *ViewColumnRequest {
	s.ProjectionPolicy = &ProjectionPolicy
	return s
}

func (s *ViewColumnRequest) WithMaskingPolicy(MaskingPolicy ViewColumnMaskingPolicyRequest) *ViewColumnRequest {
	s.MaskingPolicy = &MaskingPolicy
	return s
}

func (s *ViewColumnRequest) WithComment(Comment string) *ViewColumnRequest {
	s.Comment = &Comment
	return s
}

func (s *ViewColumnRequest) WithTag(Tag []TagAssociation) *ViewColumnRequest {
	s.Tag = Tag
	return s
}

func NewViewColumnProjectionPolicyRequest(
	ProjectionPolicy SchemaObjectIdentifier,
) *ViewColumnProjectionPolicyRequest {
	s := ViewColumnProjectionPolicyRequest{}
	s.ProjectionPolicy = ProjectionPolicy
	return &s
}

func NewViewColumnMaskingPolicyRequest(
	MaskingPolicy SchemaObjectIdentifier,
) *ViewColumnMaskingPolicyRequest {
	s := ViewColumnMaskingPolicyRequest{}
	s.MaskingPolicy = MaskingPolicy
	return &s
}

func (s *ViewColumnMaskingPolicyRequest) WithUsing(Using []Column) *ViewColumnMaskingPolicyRequest {
	s.Using = Using
	return s
}

func NewViewRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
	On []Column,
) *ViewRowAccessPolicyRequest {
	s := ViewRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	s.On = On
	return &s
}

func NewViewAggregationPolicyRequest(
	AggregationPolicy SchemaObjectIdentifier,
) *ViewAggregationPolicyRequest {
	s := ViewAggregationPolicyRequest{}
	s.AggregationPolicy = AggregationPolicy
	return &s
}

func (s *ViewAggregationPolicyRequest) WithEntityKey(EntityKey []Column) *ViewAggregationPolicyRequest {
	s.EntityKey = EntityKey
	return s
}

func NewAlterViewRequest(
	name SchemaObjectIdentifier,
) *AlterViewRequest {
	s := AlterViewRequest{}
	s.name = name
	return &s
}

func (s *AlterViewRequest) WithIfExists(IfExists bool) *AlterViewRequest {
	s.IfExists = &IfExists
	return s
}

func (s *AlterViewRequest) WithRenameTo(RenameTo SchemaObjectIdentifier) *AlterViewRequest {
	s.RenameTo = &RenameTo
	return s
}

func (s *AlterViewRequest) WithSetComment(SetComment string) *AlterViewRequest {
	s.SetComment = &SetComment
	return s
}

func (s *AlterViewRequest) WithUnsetComment(UnsetComment bool) *AlterViewRequest {
	s.UnsetComment = &UnsetComment
	return s
}

func (s *AlterViewRequest) WithSetSecure(SetSecure bool) *AlterViewRequest {
	s.SetSecure = &SetSecure
	return s
}

func (s *AlterViewRequest) WithSetChangeTracking(SetChangeTracking bool) *AlterViewRequest {
	s.SetChangeTracking = &SetChangeTracking
	return s
}

func (s *AlterViewRequest) WithUnsetSecure(UnsetSecure bool) *AlterViewRequest {
	s.UnsetSecure = &UnsetSecure
	return s
}

func (s *AlterViewRequest) WithSetTags(SetTags []TagAssociation) *AlterViewRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterViewRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterViewRequest {
	s.UnsetTags = UnsetTags
	return s
}

func (s *AlterViewRequest) WithAddDataMetricFunction(AddDataMetricFunction ViewAddDataMetricFunctionRequest) *AlterViewRequest {
	s.AddDataMetricFunction = &AddDataMetricFunction
	return s
}

func (s *AlterViewRequest) WithDropDataMetricFunction(DropDataMetricFunction ViewDropDataMetricFunctionRequest) *AlterViewRequest {
	s.DropDataMetricFunction = &DropDataMetricFunction
	return s
}

func (s *AlterViewRequest) WithSetDataMetricSchedule(SetDataMetricSchedule ViewSetDataMetricScheduleRequest) *AlterViewRequest {
	s.SetDataMetricSchedule = &SetDataMetricSchedule
	return s
}

func (s *AlterViewRequest) WithUnsetDataMetricSchedule(UnsetDataMetricSchedule ViewUnsetDataMetricScheduleRequest) *AlterViewRequest {
	s.UnsetDataMetricSchedule = &UnsetDataMetricSchedule
	return s
}

func (s *AlterViewRequest) WithAddRowAccessPolicy(AddRowAccessPolicy ViewAddRowAccessPolicyRequest) *AlterViewRequest {
	s.AddRowAccessPolicy = &AddRowAccessPolicy
	return s
}

func (s *AlterViewRequest) WithDropRowAccessPolicy(DropRowAccessPolicy ViewDropRowAccessPolicyRequest) *AlterViewRequest {
	s.DropRowAccessPolicy = &DropRowAccessPolicy
	return s
}

func (s *AlterViewRequest) WithDropAndAddRowAccessPolicy(DropAndAddRowAccessPolicy ViewDropAndAddRowAccessPolicyRequest) *AlterViewRequest {
	s.DropAndAddRowAccessPolicy = &DropAndAddRowAccessPolicy
	return s
}

func (s *AlterViewRequest) WithDropAllRowAccessPolicies(DropAllRowAccessPolicies bool) *AlterViewRequest {
	s.DropAllRowAccessPolicies = &DropAllRowAccessPolicies
	return s
}

func (s *AlterViewRequest) WithSetAggregationPolicy(SetAggregationPolicy ViewSetAggregationPolicyRequest) *AlterViewRequest {
	s.SetAggregationPolicy = &SetAggregationPolicy
	return s
}

func (s *AlterViewRequest) WithUnsetAggregationPolicy(UnsetAggregationPolicy ViewUnsetAggregationPolicyRequest) *AlterViewRequest {
	s.UnsetAggregationPolicy = &UnsetAggregationPolicy
	return s
}

func (s *AlterViewRequest) WithSetMaskingPolicyOnColumn(SetMaskingPolicyOnColumn ViewSetColumnMaskingPolicyRequest) *AlterViewRequest {
	s.SetMaskingPolicyOnColumn = &SetMaskingPolicyOnColumn
	return s
}

func (s *AlterViewRequest) WithUnsetMaskingPolicyOnColumn(UnsetMaskingPolicyOnColumn ViewUnsetColumnMaskingPolicyRequest) *AlterViewRequest {
	s.UnsetMaskingPolicyOnColumn = &UnsetMaskingPolicyOnColumn
	return s
}

func (s *AlterViewRequest) WithSetProjectionPolicyOnColumn(SetProjectionPolicyOnColumn ViewSetProjectionPolicyRequest) *AlterViewRequest {
	s.SetProjectionPolicyOnColumn = &SetProjectionPolicyOnColumn
	return s
}

func (s *AlterViewRequest) WithUnsetProjectionPolicyOnColumn(UnsetProjectionPolicyOnColumn ViewUnsetProjectionPolicyRequest) *AlterViewRequest {
	s.UnsetProjectionPolicyOnColumn = &UnsetProjectionPolicyOnColumn
	return s
}

func (s *AlterViewRequest) WithSetTagsOnColumn(SetTagsOnColumn ViewSetColumnTagsRequest) *AlterViewRequest {
	s.SetTagsOnColumn = &SetTagsOnColumn
	return s
}

func (s *AlterViewRequest) WithUnsetTagsOnColumn(UnsetTagsOnColumn ViewUnsetColumnTagsRequest) *AlterViewRequest {
	s.UnsetTagsOnColumn = &UnsetTagsOnColumn
	return s
}

func NewViewAddDataMetricFunctionRequest(
	DataMetricFunction []ViewDataMetricFunction,
) *ViewAddDataMetricFunctionRequest {
	s := ViewAddDataMetricFunctionRequest{}
	s.DataMetricFunction = DataMetricFunction
	return &s
}

func NewViewDropDataMetricFunctionRequest(
	DataMetricFunction []ViewDataMetricFunction,
) *ViewDropDataMetricFunctionRequest {
	s := ViewDropDataMetricFunctionRequest{}
	s.DataMetricFunction = DataMetricFunction
	return &s
}

func NewViewSetDataMetricScheduleRequest() *ViewSetDataMetricScheduleRequest {
	return &ViewSetDataMetricScheduleRequest{}
}

func (s *ViewSetDataMetricScheduleRequest) WithMinutes(Minutes ViewMinuteRequest) *ViewSetDataMetricScheduleRequest {
	s.Minutes = &Minutes
	return s
}

func (s *ViewSetDataMetricScheduleRequest) WithUsingCron(UsingCron ViewUsingCronRequest) *ViewSetDataMetricScheduleRequest {
	s.UsingCron = &UsingCron
	return s
}

func (s *ViewSetDataMetricScheduleRequest) WithTriggerOnChanges(TriggerOnChanges bool) *ViewSetDataMetricScheduleRequest {
	s.TriggerOnChanges = &TriggerOnChanges
	return s
}

func NewViewMinuteRequest(
	Minutes int,
) *ViewMinuteRequest {
	s := ViewMinuteRequest{}
	s.Minutes = Minutes
	return &s
}

func NewViewUsingCronRequest(
	Cron string,
) *ViewUsingCronRequest {
	s := ViewUsingCronRequest{}
	s.Cron = Cron
	return &s
}

func NewViewUnsetDataMetricScheduleRequest() *ViewUnsetDataMetricScheduleRequest {
	return &ViewUnsetDataMetricScheduleRequest{}
}

func NewViewAddRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
	On []Column,
) *ViewAddRowAccessPolicyRequest {
	s := ViewAddRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	s.On = On
	return &s
}

func NewViewDropRowAccessPolicyRequest(
	RowAccessPolicy SchemaObjectIdentifier,
) *ViewDropRowAccessPolicyRequest {
	s := ViewDropRowAccessPolicyRequest{}
	s.RowAccessPolicy = RowAccessPolicy
	return &s
}

func NewViewDropAndAddRowAccessPolicyRequest(
	Drop ViewDropRowAccessPolicyRequest,
	Add ViewAddRowAccessPolicyRequest,
) *ViewDropAndAddRowAccessPolicyRequest {
	s := ViewDropAndAddRowAccessPolicyRequest{}
	s.Drop = Drop
	s.Add = Add
	return &s
}

func NewViewSetAggregationPolicyRequest(
	AggregationPolicy SchemaObjectIdentifier,
) *ViewSetAggregationPolicyRequest {
	s := ViewSetAggregationPolicyRequest{}
	s.AggregationPolicy = AggregationPolicy
	return &s
}

func (s *ViewSetAggregationPolicyRequest) WithEntityKey(EntityKey []Column) *ViewSetAggregationPolicyRequest {
	s.EntityKey = EntityKey
	return s
}

func (s *ViewSetAggregationPolicyRequest) WithForce(Force bool) *ViewSetAggregationPolicyRequest {
	s.Force = &Force
	return s
}

func NewViewUnsetAggregationPolicyRequest() *ViewUnsetAggregationPolicyRequest {
	return &ViewUnsetAggregationPolicyRequest{}
}

func NewViewSetColumnMaskingPolicyRequest(
	Name string,
	MaskingPolicy SchemaObjectIdentifier,
) *ViewSetColumnMaskingPolicyRequest {
	s := ViewSetColumnMaskingPolicyRequest{}
	s.Name = Name
	s.MaskingPolicy = MaskingPolicy
	return &s
}

func (s *ViewSetColumnMaskingPolicyRequest) WithUsing(Using []Column) *ViewSetColumnMaskingPolicyRequest {
	s.Using = Using
	return s
}

func (s *ViewSetColumnMaskingPolicyRequest) WithForce(Force bool) *ViewSetColumnMaskingPolicyRequest {
	s.Force = &Force
	return s
}

func NewViewUnsetColumnMaskingPolicyRequest(
	Name string,
) *ViewUnsetColumnMaskingPolicyRequest {
	s := ViewUnsetColumnMaskingPolicyRequest{}
	s.Name = Name
	return &s
}

func NewViewSetProjectionPolicyRequest(
	Name string,
	ProjectionPolicy SchemaObjectIdentifier,
) *ViewSetProjectionPolicyRequest {
	s := ViewSetProjectionPolicyRequest{}
	s.Name = Name
	s.ProjectionPolicy = ProjectionPolicy
	return &s
}

func (s *ViewSetProjectionPolicyRequest) WithForce(Force bool) *ViewSetProjectionPolicyRequest {
	s.Force = &Force
	return s
}

func NewViewUnsetProjectionPolicyRequest(
	Name string,
) *ViewUnsetProjectionPolicyRequest {
	s := ViewUnsetProjectionPolicyRequest{}
	s.Name = Name
	return &s
}

func NewViewSetColumnTagsRequest(
	Name string,
	SetTags []TagAssociation,
) *ViewSetColumnTagsRequest {
	s := ViewSetColumnTagsRequest{}
	s.Name = Name
	s.SetTags = SetTags
	return &s
}

func NewViewUnsetColumnTagsRequest(
	Name string,
	UnsetTags []ObjectIdentifier,
) *ViewUnsetColumnTagsRequest {
	s := ViewUnsetColumnTagsRequest{}
	s.Name = Name
	s.UnsetTags = UnsetTags
	return &s
}

func NewDropViewRequest(
	name SchemaObjectIdentifier,
) *DropViewRequest {
	s := DropViewRequest{}
	s.name = name
	return &s
}

func (s *DropViewRequest) WithIfExists(IfExists bool) *DropViewRequest {
	s.IfExists = &IfExists
	return s
}

func NewShowViewRequest() *ShowViewRequest {
	return &ShowViewRequest{}
}

func (s *ShowViewRequest) WithTerse(Terse bool) *ShowViewRequest {
	s.Terse = &Terse
	return s
}

func (s *ShowViewRequest) WithLike(Like Like) *ShowViewRequest {
	s.Like = &Like
	return s
}

func (s *ShowViewRequest) WithIn(In ExtendedIn) *ShowViewRequest {
	s.In = &In
	return s
}

func (s *ShowViewRequest) WithStartsWith(StartsWith string) *ShowViewRequest {
	s.StartsWith = &StartsWith
	return s
}

func (s *ShowViewRequest) WithLimit(Limit LimitFrom) *ShowViewRequest {
	s.Limit = &Limit
	return s
}

func NewDescribeViewRequest(
	name SchemaObjectIdentifier,
) *DescribeViewRequest {
	s := DescribeViewRequest{}
	s.name = name
	return &s
}
