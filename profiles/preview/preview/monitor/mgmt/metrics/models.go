//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package metrics

import original "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-05-01-preview/metrics"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationType = original.AggregationType

const (
	AggregationTypeAverage AggregationType = original.AggregationTypeAverage
	AggregationTypeCount   AggregationType = original.AggregationTypeCount
	AggregationTypeMaximum AggregationType = original.AggregationTypeMaximum
	AggregationTypeMinimum AggregationType = original.AggregationTypeMinimum
	AggregationTypeNone    AggregationType = original.AggregationTypeNone
	AggregationTypeTotal   AggregationType = original.AggregationTypeTotal
)

type AggregationTypeEnum = original.AggregationTypeEnum

const (
	AggregationTypeEnumAverage AggregationTypeEnum = original.AggregationTypeEnumAverage
	AggregationTypeEnumCount   AggregationTypeEnum = original.AggregationTypeEnumCount
	AggregationTypeEnumMaximum AggregationTypeEnum = original.AggregationTypeEnumMaximum
	AggregationTypeEnumMinimum AggregationTypeEnum = original.AggregationTypeEnumMinimum
	AggregationTypeEnumTotal   AggregationTypeEnum = original.AggregationTypeEnumTotal
)

type BaselineSensitivity = original.BaselineSensitivity

const (
	BaselineSensitivityHigh   BaselineSensitivity = original.BaselineSensitivityHigh
	BaselineSensitivityLow    BaselineSensitivity = original.BaselineSensitivityLow
	BaselineSensitivityMedium BaselineSensitivity = original.BaselineSensitivityMedium
)

type CriterionType = original.CriterionType

const (
	CriterionTypeDynamicThresholdCriterion CriterionType = original.CriterionTypeDynamicThresholdCriterion
	CriterionTypeMultiMetricCriteria       CriterionType = original.CriterionTypeMultiMetricCriteria
	CriterionTypeStaticThresholdCriterion  CriterionType = original.CriterionTypeStaticThresholdCriterion
)

type DynamicThresholdOperator = original.DynamicThresholdOperator

const (
	DynamicThresholdOperatorGreaterOrLessThan DynamicThresholdOperator = original.DynamicThresholdOperatorGreaterOrLessThan
	DynamicThresholdOperatorGreaterThan       DynamicThresholdOperator = original.DynamicThresholdOperatorGreaterThan
	DynamicThresholdOperatorLessThan          DynamicThresholdOperator = original.DynamicThresholdOperatorLessThan
)

type DynamicThresholdSensitivity = original.DynamicThresholdSensitivity

const (
	DynamicThresholdSensitivityHigh   DynamicThresholdSensitivity = original.DynamicThresholdSensitivityHigh
	DynamicThresholdSensitivityLow    DynamicThresholdSensitivity = original.DynamicThresholdSensitivityLow
	DynamicThresholdSensitivityMedium DynamicThresholdSensitivity = original.DynamicThresholdSensitivityMedium
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeCount   MetricAggregationType = original.MetricAggregationTypeCount
	MetricAggregationTypeMaximum MetricAggregationType = original.MetricAggregationTypeMaximum
	MetricAggregationTypeMinimum MetricAggregationType = original.MetricAggregationTypeMinimum
	MetricAggregationTypeNone    MetricAggregationType = original.MetricAggregationTypeNone
	MetricAggregationTypeTotal   MetricAggregationType = original.MetricAggregationTypeTotal
)

type MetricClass = original.MetricClass

const (
	MetricClassAvailability MetricClass = original.MetricClassAvailability
	MetricClassErrors       MetricClass = original.MetricClassErrors
	MetricClassLatency      MetricClass = original.MetricClassLatency
	MetricClassSaturation   MetricClass = original.MetricClassSaturation
	MetricClassTransactions MetricClass = original.MetricClassTransactions
)

type MetricResultType = original.MetricResultType

const (
	MetricResultTypeData     MetricResultType = original.MetricResultTypeData
	MetricResultTypeMetadata MetricResultType = original.MetricResultTypeMetadata
)

type MetricUnit = original.MetricUnit

const (
	MetricUnitBitsPerSecond  MetricUnit = original.MetricUnitBitsPerSecond
	MetricUnitBytes          MetricUnit = original.MetricUnitBytes
	MetricUnitByteSeconds    MetricUnit = original.MetricUnitByteSeconds
	MetricUnitBytesPerSecond MetricUnit = original.MetricUnitBytesPerSecond
	MetricUnitCores          MetricUnit = original.MetricUnitCores
	MetricUnitCount          MetricUnit = original.MetricUnitCount
	MetricUnitCountPerSecond MetricUnit = original.MetricUnitCountPerSecond
	MetricUnitMilliCores     MetricUnit = original.MetricUnitMilliCores
	MetricUnitMilliSeconds   MetricUnit = original.MetricUnitMilliSeconds
	MetricUnitNanoCores      MetricUnit = original.MetricUnitNanoCores
	MetricUnitPercent        MetricUnit = original.MetricUnitPercent
	MetricUnitSeconds        MetricUnit = original.MetricUnitSeconds
	MetricUnitUnspecified    MetricUnit = original.MetricUnitUnspecified
)

type NamespaceClassification = original.NamespaceClassification

const (
	NamespaceClassificationCustom   NamespaceClassification = original.NamespaceClassificationCustom
	NamespaceClassificationPlatform NamespaceClassification = original.NamespaceClassificationPlatform
	NamespaceClassificationQos      NamespaceClassification = original.NamespaceClassificationQos
)

type OdataType = original.OdataType

const (
	OdataTypeMetricAlertCriteria                                         OdataType = original.OdataTypeMetricAlertCriteria
	OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria OdataType = original.OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria
	OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria   OdataType = original.OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria
	OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria    OdataType = original.OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria
)

type Operator = original.Operator

const (
	OperatorEquals             Operator = original.OperatorEquals
	OperatorGreaterThan        Operator = original.OperatorGreaterThan
	OperatorGreaterThanOrEqual Operator = original.OperatorGreaterThanOrEqual
	OperatorLessThan           Operator = original.OperatorLessThan
	OperatorLessThanOrEqual    Operator = original.OperatorLessThanOrEqual
)

type ResultType = original.ResultType

const (
	ResultTypeData     ResultType = original.ResultTypeData
	ResultTypeMetadata ResultType = original.ResultTypeMetadata
)

type Unit = original.Unit

const (
	UnitBitsPerSecond  Unit = original.UnitBitsPerSecond
	UnitBytes          Unit = original.UnitBytes
	UnitByteSeconds    Unit = original.UnitByteSeconds
	UnitBytesPerSecond Unit = original.UnitBytesPerSecond
	UnitCores          Unit = original.UnitCores
	UnitCount          Unit = original.UnitCount
	UnitCountPerSecond Unit = original.UnitCountPerSecond
	UnitMilliCores     Unit = original.UnitMilliCores
	UnitMilliSeconds   Unit = original.UnitMilliSeconds
	UnitNanoCores      Unit = original.UnitNanoCores
	UnitPercent        Unit = original.UnitPercent
	UnitSeconds        Unit = original.UnitSeconds
	UnitUnspecified    Unit = original.UnitUnspecified
)

type AlertAction = original.AlertAction
type AlertCriteria = original.AlertCriteria
type AlertMultipleResourceMultipleMetricCriteria = original.AlertMultipleResourceMultipleMetricCriteria
type AlertProperties = original.AlertProperties
type AlertPropertiesPatch = original.AlertPropertiesPatch
type AlertResource = original.AlertResource
type AlertResourceCollection = original.AlertResourceCollection
type AlertResourcePatch = original.AlertResourcePatch
type AlertSingleResourceMultipleMetricCriteria = original.AlertSingleResourceMultipleMetricCriteria
type AlertStatus = original.AlertStatus
type AlertStatusCollection = original.AlertStatusCollection
type AlertStatusProperties = original.AlertStatusProperties
type AlertsClient = original.AlertsClient
type AlertsStatusClient = original.AlertsStatusClient
type Availability = original.Availability
type BaseClient = original.BaseClient
type BaselineMetadata = original.BaselineMetadata
type BaselinesClient = original.BaselinesClient
type BaselinesProperties = original.BaselinesProperties
type BaselinesResponse = original.BaselinesResponse
type BasicAlertCriteria = original.BasicAlertCriteria
type BasicMultiMetricCriteria = original.BasicMultiMetricCriteria
type Client = original.Client
type Criteria = original.Criteria
type Definition = original.Definition
type DefinitionCollection = original.DefinitionCollection
type DefinitionsClient = original.DefinitionsClient
type Dimension = original.Dimension
type DimensionProperties = original.DimensionProperties
type DynamicMetricCriteria = original.DynamicMetricCriteria
type DynamicThresholdFailingPeriods = original.DynamicThresholdFailingPeriods
type ErrorContract = original.ErrorContract
type ErrorResponse = original.ErrorResponse
type LocalizableString = original.LocalizableString
type LogSpecification = original.LogSpecification
type MetadataValue = original.MetadataValue
type Metric = original.Metric
type MultiMetricCriteria = original.MultiMetricCriteria
type Namespace = original.Namespace
type NamespaceCollection = original.NamespaceCollection
type NamespaceName = original.NamespaceName
type NamespacesClient = original.NamespacesClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type Response = original.Response
type ServiceSpecification = original.ServiceSpecification
type SingleBaseline = original.SingleBaseline
type SingleDimension = original.SingleDimension
type SingleMetricBaseline = original.SingleMetricBaseline
type Specification = original.Specification
type SubscriptionScopeMetric = original.SubscriptionScopeMetric
type SubscriptionScopeMetricDefinition = original.SubscriptionScopeMetricDefinition
type SubscriptionScopeMetricDefinitionCollection = original.SubscriptionScopeMetricDefinitionCollection
type SubscriptionScopeMetricResponse = original.SubscriptionScopeMetricResponse
type SubscriptionScopeMetricsRequestBodyParameters = original.SubscriptionScopeMetricsRequestBodyParameters
type TimeSeriesBaseline = original.TimeSeriesBaseline
type TimeSeriesElement = original.TimeSeriesElement
type Value = original.Value
type WebtestLocationAvailabilityCriteria = original.WebtestLocationAvailabilityCriteria

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertsStatusClient(subscriptionID string) AlertsStatusClient {
	return original.NewAlertsStatusClient(subscriptionID)
}
func NewAlertsStatusClientWithBaseURI(baseURI string, subscriptionID string) AlertsStatusClient {
	return original.NewAlertsStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewBaselinesClient(subscriptionID string) BaselinesClient {
	return original.NewBaselinesClient(subscriptionID)
}
func NewBaselinesClientWithBaseURI(baseURI string, subscriptionID string) BaselinesClient {
	return original.NewBaselinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClient(subscriptionID)
}
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAggregationTypeEnumValues() []AggregationTypeEnum {
	return original.PossibleAggregationTypeEnumValues()
}
func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}
func PossibleBaselineSensitivityValues() []BaselineSensitivity {
	return original.PossibleBaselineSensitivityValues()
}
func PossibleCriterionTypeValues() []CriterionType {
	return original.PossibleCriterionTypeValues()
}
func PossibleDynamicThresholdOperatorValues() []DynamicThresholdOperator {
	return original.PossibleDynamicThresholdOperatorValues()
}
func PossibleDynamicThresholdSensitivityValues() []DynamicThresholdSensitivity {
	return original.PossibleDynamicThresholdSensitivityValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleMetricClassValues() []MetricClass {
	return original.PossibleMetricClassValues()
}
func PossibleMetricResultTypeValues() []MetricResultType {
	return original.PossibleMetricResultTypeValues()
}
func PossibleMetricUnitValues() []MetricUnit {
	return original.PossibleMetricUnitValues()
}
func PossibleNamespaceClassificationValues() []NamespaceClassification {
	return original.PossibleNamespaceClassificationValues()
}
func PossibleOdataTypeValues() []OdataType {
	return original.PossibleOdataTypeValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}
func PossibleUnitValues() []Unit {
	return original.PossibleUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}