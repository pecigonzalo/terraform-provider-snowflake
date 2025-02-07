// Code generated by assertions generator; DO NOT EDIT.

package resourceparametersassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type UserResourceParametersAssert struct {
	*assert.ResourceAssert
}

func UserResourceParameters(t *testing.T, name string) *UserResourceParametersAssert {
	t.Helper()

	resourceParameterAssert := UserResourceParametersAssert{
		ResourceAssert: assert.NewResourceAssert(name, "parameters"),
	}
	resourceParameterAssert.AddAssertion(assert.ValueSet("parameters.#", "1"))
	return &resourceParameterAssert
}

func ImportedUserResourceParameters(t *testing.T, id string) *UserResourceParametersAssert {
	t.Helper()

	resourceParameterAssert := UserResourceParametersAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported parameters"),
	}
	resourceParameterAssert.AddAssertion(assert.ValueSet("parameters.#", "1"))
	return &resourceParameterAssert
}

////////////////////////////
// Parameter value checks //
////////////////////////////

func (u *UserResourceParametersAssert) HasEnableUnredactedQuerySyntaxError(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterEnableUnredactedQuerySyntaxError, expected))
	return u
}

func (u *UserResourceParametersAssert) HasNetworkPolicy(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterNetworkPolicy, expected))
	return u
}

func (u *UserResourceParametersAssert) HasPreventUnloadToInternalStages(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterPreventUnloadToInternalStages, expected))
	return u
}

func (u *UserResourceParametersAssert) HasAbortDetachedQuery(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterAbortDetachedQuery, expected))
	return u
}

func (u *UserResourceParametersAssert) HasAutocommit(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterAutocommit, expected))
	return u
}

func (u *UserResourceParametersAssert) HasBinaryInputFormat(expected sdk.BinaryInputFormat) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterBinaryInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasBinaryOutputFormat(expected sdk.BinaryOutputFormat) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterBinaryOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientMemoryLimit(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterClientMemoryLimit, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientMetadataRequestUseConnectionCtx(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterClientMetadataRequestUseConnectionCtx, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientPrefetchThreads(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterClientPrefetchThreads, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientResultChunkSize(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterClientResultChunkSize, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientResultColumnCaseInsensitive(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterClientResultColumnCaseInsensitive, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientSessionKeepAlive(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterClientSessionKeepAlive, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientSessionKeepAliveHeartbeatFrequency(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterClientSessionKeepAliveHeartbeatFrequency, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientTimestampTypeMapping(expected sdk.ClientTimestampTypeMapping) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterClientTimestampTypeMapping, expected))
	return u
}

func (u *UserResourceParametersAssert) HasDateInputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterDateInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasDateOutputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterDateOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasEnableUnloadPhysicalTypeOptimization(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterEnableUnloadPhysicalTypeOptimization, expected))
	return u
}

func (u *UserResourceParametersAssert) HasErrorOnNondeterministicMerge(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterErrorOnNondeterministicMerge, expected))
	return u
}

func (u *UserResourceParametersAssert) HasErrorOnNondeterministicUpdate(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterErrorOnNondeterministicUpdate, expected))
	return u
}

func (u *UserResourceParametersAssert) HasGeographyOutputFormat(expected sdk.GeographyOutputFormat) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterGeographyOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasGeometryOutputFormat(expected sdk.GeometryOutputFormat) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterGeometryOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJdbcTreatDecimalAsInt(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterJdbcTreatDecimalAsInt, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJdbcTreatTimestampNtzAsUtc(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterJdbcTreatTimestampNtzAsUtc, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJdbcUseSessionTimezone(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterJdbcUseSessionTimezone, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJsonIndent(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterJsonIndent, expected))
	return u
}

func (u *UserResourceParametersAssert) HasLockTimeout(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterLockTimeout, expected))
	return u
}

func (u *UserResourceParametersAssert) HasLogLevel(expected sdk.LogLevel) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterLogLevel, expected))
	return u
}

func (u *UserResourceParametersAssert) HasMultiStatementCount(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterMultiStatementCount, expected))
	return u
}

func (u *UserResourceParametersAssert) HasNoorderSequenceAsDefault(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterNoorderSequenceAsDefault, expected))
	return u
}

func (u *UserResourceParametersAssert) HasOdbcTreatDecimalAsInt(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterOdbcTreatDecimalAsInt, expected))
	return u
}

func (u *UserResourceParametersAssert) HasQueryTag(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterQueryTag, expected))
	return u
}

func (u *UserResourceParametersAssert) HasQuotedIdentifiersIgnoreCase(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterQuotedIdentifiersIgnoreCase, expected))
	return u
}

func (u *UserResourceParametersAssert) HasRowsPerResultset(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterRowsPerResultset, expected))
	return u
}

func (u *UserResourceParametersAssert) HasS3StageVpceDnsName(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterS3StageVpceDnsName, expected))
	return u
}

func (u *UserResourceParametersAssert) HasSearchPath(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterSearchPath, expected))
	return u
}

func (u *UserResourceParametersAssert) HasSimulatedDataSharingConsumer(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterSimulatedDataSharingConsumer, expected))
	return u
}

func (u *UserResourceParametersAssert) HasStatementQueuedTimeoutInSeconds(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterStatementQueuedTimeoutInSeconds, expected))
	return u
}

func (u *UserResourceParametersAssert) HasStatementTimeoutInSeconds(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterStatementTimeoutInSeconds, expected))
	return u
}

func (u *UserResourceParametersAssert) HasStrictJsonOutput(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterStrictJsonOutput, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampDayIsAlways24h(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterTimestampDayIsAlways24h, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampInputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimestampInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampLtzOutputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimestampLtzOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampNtzOutputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimestampNtzOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampOutputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimestampOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampTypeMapping(expected sdk.TimestampTypeMapping) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterTimestampTypeMapping, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampTzOutputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimestampTzOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimezone(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimezone, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimeInputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimeInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimeOutputFormat(expected string) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterValueSet(sdk.UserParameterTimeOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTraceLevel(expected sdk.TraceLevel) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterTraceLevel, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTransactionAbortOnError(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterTransactionAbortOnError, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTransactionDefaultIsolationLevel(expected sdk.TransactionDefaultIsolationLevel) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterTransactionDefaultIsolationLevel, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTwoDigitCenturyStart(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterTwoDigitCenturyStart, expected))
	return u
}

func (u *UserResourceParametersAssert) HasUnsupportedDdlAction(expected sdk.UnsupportedDDLAction) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterStringUnderlyingValueSet(sdk.UserParameterUnsupportedDdlAction, expected))
	return u
}

func (u *UserResourceParametersAssert) HasUseCachedResult(expected bool) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterBoolValueSet(sdk.UserParameterUseCachedResult, expected))
	return u
}

func (u *UserResourceParametersAssert) HasWeekOfYearPolicy(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterWeekOfYearPolicy, expected))
	return u
}

func (u *UserResourceParametersAssert) HasWeekStart(expected int) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterIntValueSet(sdk.UserParameterWeekStart, expected))
	return u
}

////////////////////////////
// Parameter level checks //
////////////////////////////

func (u *UserResourceParametersAssert) HasEnableUnredactedQuerySyntaxErrorLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterEnableUnredactedQuerySyntaxError, expected))
	return u
}

func (u *UserResourceParametersAssert) HasNetworkPolicyLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterNetworkPolicy, expected))
	return u
}

func (u *UserResourceParametersAssert) HasPreventUnloadToInternalStagesLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterPreventUnloadToInternalStages, expected))
	return u
}

func (u *UserResourceParametersAssert) HasAbortDetachedQueryLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterAbortDetachedQuery, expected))
	return u
}

func (u *UserResourceParametersAssert) HasAutocommitLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterAutocommit, expected))
	return u
}

func (u *UserResourceParametersAssert) HasBinaryInputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterBinaryInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasBinaryOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterBinaryOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientMemoryLimitLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientMemoryLimit, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientMetadataRequestUseConnectionCtxLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientMetadataRequestUseConnectionCtx, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientPrefetchThreadsLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientPrefetchThreads, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientResultChunkSizeLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientResultChunkSize, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientResultColumnCaseInsensitiveLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientResultColumnCaseInsensitive, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientSessionKeepAliveLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientSessionKeepAlive, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientSessionKeepAliveHeartbeatFrequencyLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientSessionKeepAliveHeartbeatFrequency, expected))
	return u
}

func (u *UserResourceParametersAssert) HasClientTimestampTypeMappingLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterClientTimestampTypeMapping, expected))
	return u
}

func (u *UserResourceParametersAssert) HasDateInputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterDateInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasDateOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterDateOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasEnableUnloadPhysicalTypeOptimizationLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterEnableUnloadPhysicalTypeOptimization, expected))
	return u
}

func (u *UserResourceParametersAssert) HasErrorOnNondeterministicMergeLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterErrorOnNondeterministicMerge, expected))
	return u
}

func (u *UserResourceParametersAssert) HasErrorOnNondeterministicUpdateLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterErrorOnNondeterministicUpdate, expected))
	return u
}

func (u *UserResourceParametersAssert) HasGeographyOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterGeographyOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasGeometryOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterGeometryOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJdbcTreatDecimalAsIntLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterJdbcTreatDecimalAsInt, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJdbcTreatTimestampNtzAsUtcLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterJdbcTreatTimestampNtzAsUtc, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJdbcUseSessionTimezoneLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterJdbcUseSessionTimezone, expected))
	return u
}

func (u *UserResourceParametersAssert) HasJsonIndentLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterJsonIndent, expected))
	return u
}

func (u *UserResourceParametersAssert) HasLockTimeoutLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterLockTimeout, expected))
	return u
}

func (u *UserResourceParametersAssert) HasLogLevelLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterLogLevel, expected))
	return u
}

func (u *UserResourceParametersAssert) HasMultiStatementCountLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterMultiStatementCount, expected))
	return u
}

func (u *UserResourceParametersAssert) HasNoorderSequenceAsDefaultLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterNoorderSequenceAsDefault, expected))
	return u
}

func (u *UserResourceParametersAssert) HasOdbcTreatDecimalAsIntLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterOdbcTreatDecimalAsInt, expected))
	return u
}

func (u *UserResourceParametersAssert) HasQueryTagLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterQueryTag, expected))
	return u
}

func (u *UserResourceParametersAssert) HasQuotedIdentifiersIgnoreCaseLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterQuotedIdentifiersIgnoreCase, expected))
	return u
}

func (u *UserResourceParametersAssert) HasRowsPerResultsetLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterRowsPerResultset, expected))
	return u
}

func (u *UserResourceParametersAssert) HasS3StageVpceDnsNameLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterS3StageVpceDnsName, expected))
	return u
}

func (u *UserResourceParametersAssert) HasSearchPathLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterSearchPath, expected))
	return u
}

func (u *UserResourceParametersAssert) HasSimulatedDataSharingConsumerLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterSimulatedDataSharingConsumer, expected))
	return u
}

func (u *UserResourceParametersAssert) HasStatementQueuedTimeoutInSecondsLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterStatementQueuedTimeoutInSeconds, expected))
	return u
}

func (u *UserResourceParametersAssert) HasStatementTimeoutInSecondsLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterStatementTimeoutInSeconds, expected))
	return u
}

func (u *UserResourceParametersAssert) HasStrictJsonOutputLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterStrictJsonOutput, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampDayIsAlways24hLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampDayIsAlways24h, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampInputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampLtzOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampLtzOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampNtzOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampNtzOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampTypeMappingLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampTypeMapping, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimestampTzOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimestampTzOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimezoneLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimezone, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimeInputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimeInputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTimeOutputFormatLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTimeOutputFormat, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTraceLevelLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTraceLevel, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTransactionAbortOnErrorLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTransactionAbortOnError, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTransactionDefaultIsolationLevelLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTransactionDefaultIsolationLevel, expected))
	return u
}

func (u *UserResourceParametersAssert) HasTwoDigitCenturyStartLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterTwoDigitCenturyStart, expected))
	return u
}

func (u *UserResourceParametersAssert) HasUnsupportedDdlActionLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterUnsupportedDdlAction, expected))
	return u
}

func (u *UserResourceParametersAssert) HasUseCachedResultLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterUseCachedResult, expected))
	return u
}

func (u *UserResourceParametersAssert) HasWeekOfYearPolicyLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterWeekOfYearPolicy, expected))
	return u
}

func (u *UserResourceParametersAssert) HasWeekStartLevel(expected sdk.ParameterType) *UserResourceParametersAssert {
	u.AddAssertion(assert.ResourceParameterLevelSet(sdk.UserParameterWeekStart, expected))
	return u
}
