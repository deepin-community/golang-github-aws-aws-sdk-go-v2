// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdminStatus string

// Enum values for AdminStatus
const (
	AdminStatusEnabled           AdminStatus = "ENABLED"
	AdminStatusDisableInProgress AdminStatus = "DISABLE_IN_PROGRESS"
)

// Values returns all known values for AdminStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdminStatus) Values() []AdminStatus {
	return []AdminStatus{
		"ENABLED",
		"DISABLE_IN_PROGRESS",
	}
}

type AutoEnableMembers string

// Enum values for AutoEnableMembers
const (
	AutoEnableMembersNew  AutoEnableMembers = "NEW"
	AutoEnableMembersAll  AutoEnableMembers = "ALL"
	AutoEnableMembersNone AutoEnableMembers = "NONE"
)

// Values returns all known values for AutoEnableMembers. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutoEnableMembers) Values() []AutoEnableMembers {
	return []AutoEnableMembers{
		"NEW",
		"ALL",
		"NONE",
	}
}

type CoverageFilterCriterionKey string

// Enum values for CoverageFilterCriterionKey
const (
	CoverageFilterCriterionKeyAccountId      CoverageFilterCriterionKey = "ACCOUNT_ID"
	CoverageFilterCriterionKeyClusterName    CoverageFilterCriterionKey = "CLUSTER_NAME"
	CoverageFilterCriterionKeyResourceType   CoverageFilterCriterionKey = "RESOURCE_TYPE"
	CoverageFilterCriterionKeyCoverageStatus CoverageFilterCriterionKey = "COVERAGE_STATUS"
	CoverageFilterCriterionKeyAddonVersion   CoverageFilterCriterionKey = "ADDON_VERSION"
	CoverageFilterCriterionKeyManagementType CoverageFilterCriterionKey = "MANAGEMENT_TYPE"
	CoverageFilterCriterionKeyEksClusterName CoverageFilterCriterionKey = "EKS_CLUSTER_NAME"
	CoverageFilterCriterionKeyEcsClusterName CoverageFilterCriterionKey = "ECS_CLUSTER_NAME"
	CoverageFilterCriterionKeyAgentVersion   CoverageFilterCriterionKey = "AGENT_VERSION"
	CoverageFilterCriterionKeyInstanceId     CoverageFilterCriterionKey = "INSTANCE_ID"
	CoverageFilterCriterionKeyClusterArn     CoverageFilterCriterionKey = "CLUSTER_ARN"
)

// Values returns all known values for CoverageFilterCriterionKey. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CoverageFilterCriterionKey) Values() []CoverageFilterCriterionKey {
	return []CoverageFilterCriterionKey{
		"ACCOUNT_ID",
		"CLUSTER_NAME",
		"RESOURCE_TYPE",
		"COVERAGE_STATUS",
		"ADDON_VERSION",
		"MANAGEMENT_TYPE",
		"EKS_CLUSTER_NAME",
		"ECS_CLUSTER_NAME",
		"AGENT_VERSION",
		"INSTANCE_ID",
		"CLUSTER_ARN",
	}
}

type CoverageSortKey string

// Enum values for CoverageSortKey
const (
	CoverageSortKeyAccountId      CoverageSortKey = "ACCOUNT_ID"
	CoverageSortKeyClusterName    CoverageSortKey = "CLUSTER_NAME"
	CoverageSortKeyCoverageStatus CoverageSortKey = "COVERAGE_STATUS"
	CoverageSortKeyIssue          CoverageSortKey = "ISSUE"
	CoverageSortKeyAddonVersion   CoverageSortKey = "ADDON_VERSION"
	CoverageSortKeyUpdatedAt      CoverageSortKey = "UPDATED_AT"
	CoverageSortKeyEksClusterName CoverageSortKey = "EKS_CLUSTER_NAME"
	CoverageSortKeyEcsClusterName CoverageSortKey = "ECS_CLUSTER_NAME"
	CoverageSortKeyInstanceId     CoverageSortKey = "INSTANCE_ID"
)

// Values returns all known values for CoverageSortKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CoverageSortKey) Values() []CoverageSortKey {
	return []CoverageSortKey{
		"ACCOUNT_ID",
		"CLUSTER_NAME",
		"COVERAGE_STATUS",
		"ISSUE",
		"ADDON_VERSION",
		"UPDATED_AT",
		"EKS_CLUSTER_NAME",
		"ECS_CLUSTER_NAME",
		"INSTANCE_ID",
	}
}

type CoverageStatisticsType string

// Enum values for CoverageStatisticsType
const (
	CoverageStatisticsTypeCountByResourceType   CoverageStatisticsType = "COUNT_BY_RESOURCE_TYPE"
	CoverageStatisticsTypeCountByCoverageStatus CoverageStatisticsType = "COUNT_BY_COVERAGE_STATUS"
)

// Values returns all known values for CoverageStatisticsType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CoverageStatisticsType) Values() []CoverageStatisticsType {
	return []CoverageStatisticsType{
		"COUNT_BY_RESOURCE_TYPE",
		"COUNT_BY_COVERAGE_STATUS",
	}
}

type CoverageStatus string

// Enum values for CoverageStatus
const (
	CoverageStatusHealthy   CoverageStatus = "HEALTHY"
	CoverageStatusUnhealthy CoverageStatus = "UNHEALTHY"
)

// Values returns all known values for CoverageStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CoverageStatus) Values() []CoverageStatus {
	return []CoverageStatus{
		"HEALTHY",
		"UNHEALTHY",
	}
}

type CriterionKey string

// Enum values for CriterionKey
const (
	CriterionKeyEc2InstanceArn     CriterionKey = "EC2_INSTANCE_ARN"
	CriterionKeyScanId             CriterionKey = "SCAN_ID"
	CriterionKeyAccountId          CriterionKey = "ACCOUNT_ID"
	CriterionKeyGuarddutyFindingId CriterionKey = "GUARDDUTY_FINDING_ID"
	CriterionKeyScanStartTime      CriterionKey = "SCAN_START_TIME"
	CriterionKeyScanStatus         CriterionKey = "SCAN_STATUS"
	CriterionKeyScanType           CriterionKey = "SCAN_TYPE"
)

// Values returns all known values for CriterionKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CriterionKey) Values() []CriterionKey {
	return []CriterionKey{
		"EC2_INSTANCE_ARN",
		"SCAN_ID",
		"ACCOUNT_ID",
		"GUARDDUTY_FINDING_ID",
		"SCAN_START_TIME",
		"SCAN_STATUS",
		"SCAN_TYPE",
	}
}

type DataSource string

// Enum values for DataSource
const (
	DataSourceFlowLogs            DataSource = "FLOW_LOGS"
	DataSourceCloudTrail          DataSource = "CLOUD_TRAIL"
	DataSourceDnsLogs             DataSource = "DNS_LOGS"
	DataSourceS3Logs              DataSource = "S3_LOGS"
	DataSourceKubernetesAuditLogs DataSource = "KUBERNETES_AUDIT_LOGS"
	DataSourceEc2MalwareScan      DataSource = "EC2_MALWARE_SCAN"
)

// Values returns all known values for DataSource. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSource) Values() []DataSource {
	return []DataSource{
		"FLOW_LOGS",
		"CLOUD_TRAIL",
		"DNS_LOGS",
		"S3_LOGS",
		"KUBERNETES_AUDIT_LOGS",
		"EC2_MALWARE_SCAN",
	}
}

type DataSourceStatus string

// Enum values for DataSourceStatus
const (
	DataSourceStatusEnabled  DataSourceStatus = "ENABLED"
	DataSourceStatusDisabled DataSourceStatus = "DISABLED"
)

// Values returns all known values for DataSourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceStatus) Values() []DataSourceStatus {
	return []DataSourceStatus{
		"ENABLED",
		"DISABLED",
	}
}

type DestinationType string

// Enum values for DestinationType
const (
	DestinationTypeS3 DestinationType = "S3"
)

// Values returns all known values for DestinationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DestinationType) Values() []DestinationType {
	return []DestinationType{
		"S3",
	}
}

type DetectorFeature string

// Enum values for DetectorFeature
const (
	DetectorFeatureS3DataEvents         DetectorFeature = "S3_DATA_EVENTS"
	DetectorFeatureEksAuditLogs         DetectorFeature = "EKS_AUDIT_LOGS"
	DetectorFeatureEbsMalwareProtection DetectorFeature = "EBS_MALWARE_PROTECTION"
	DetectorFeatureRdsLoginEvents       DetectorFeature = "RDS_LOGIN_EVENTS"
	DetectorFeatureEksRuntimeMonitoring DetectorFeature = "EKS_RUNTIME_MONITORING"
	DetectorFeatureLambdaNetworkLogs    DetectorFeature = "LAMBDA_NETWORK_LOGS"
	DetectorFeatureRuntimeMonitoring    DetectorFeature = "RUNTIME_MONITORING"
)

// Values returns all known values for DetectorFeature. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DetectorFeature) Values() []DetectorFeature {
	return []DetectorFeature{
		"S3_DATA_EVENTS",
		"EKS_AUDIT_LOGS",
		"EBS_MALWARE_PROTECTION",
		"RDS_LOGIN_EVENTS",
		"EKS_RUNTIME_MONITORING",
		"LAMBDA_NETWORK_LOGS",
		"RUNTIME_MONITORING",
	}
}

type DetectorFeatureResult string

// Enum values for DetectorFeatureResult
const (
	DetectorFeatureResultFlowLogs             DetectorFeatureResult = "FLOW_LOGS"
	DetectorFeatureResultCloudTrail           DetectorFeatureResult = "CLOUD_TRAIL"
	DetectorFeatureResultDnsLogs              DetectorFeatureResult = "DNS_LOGS"
	DetectorFeatureResultS3DataEvents         DetectorFeatureResult = "S3_DATA_EVENTS"
	DetectorFeatureResultEksAuditLogs         DetectorFeatureResult = "EKS_AUDIT_LOGS"
	DetectorFeatureResultEbsMalwareProtection DetectorFeatureResult = "EBS_MALWARE_PROTECTION"
	DetectorFeatureResultRdsLoginEvents       DetectorFeatureResult = "RDS_LOGIN_EVENTS"
	DetectorFeatureResultEksRuntimeMonitoring DetectorFeatureResult = "EKS_RUNTIME_MONITORING"
	DetectorFeatureResultLambdaNetworkLogs    DetectorFeatureResult = "LAMBDA_NETWORK_LOGS"
	DetectorFeatureResultRuntimeMonitoring    DetectorFeatureResult = "RUNTIME_MONITORING"
)

// Values returns all known values for DetectorFeatureResult. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DetectorFeatureResult) Values() []DetectorFeatureResult {
	return []DetectorFeatureResult{
		"FLOW_LOGS",
		"CLOUD_TRAIL",
		"DNS_LOGS",
		"S3_DATA_EVENTS",
		"EKS_AUDIT_LOGS",
		"EBS_MALWARE_PROTECTION",
		"RDS_LOGIN_EVENTS",
		"EKS_RUNTIME_MONITORING",
		"LAMBDA_NETWORK_LOGS",
		"RUNTIME_MONITORING",
	}
}

type DetectorStatus string

// Enum values for DetectorStatus
const (
	DetectorStatusEnabled  DetectorStatus = "ENABLED"
	DetectorStatusDisabled DetectorStatus = "DISABLED"
)

// Values returns all known values for DetectorStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DetectorStatus) Values() []DetectorStatus {
	return []DetectorStatus{
		"ENABLED",
		"DISABLED",
	}
}

type EbsSnapshotPreservation string

// Enum values for EbsSnapshotPreservation
const (
	EbsSnapshotPreservationNoRetention          EbsSnapshotPreservation = "NO_RETENTION"
	EbsSnapshotPreservationRetentionWithFinding EbsSnapshotPreservation = "RETENTION_WITH_FINDING"
)

// Values returns all known values for EbsSnapshotPreservation. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EbsSnapshotPreservation) Values() []EbsSnapshotPreservation {
	return []EbsSnapshotPreservation{
		"NO_RETENTION",
		"RETENTION_WITH_FINDING",
	}
}

type FeatureAdditionalConfiguration string

// Enum values for FeatureAdditionalConfiguration
const (
	FeatureAdditionalConfigurationEksAddonManagement        FeatureAdditionalConfiguration = "EKS_ADDON_MANAGEMENT"
	FeatureAdditionalConfigurationEcsFargateAgentManagement FeatureAdditionalConfiguration = "ECS_FARGATE_AGENT_MANAGEMENT"
	FeatureAdditionalConfigurationEc2AgentManagement        FeatureAdditionalConfiguration = "EC2_AGENT_MANAGEMENT"
)

// Values returns all known values for FeatureAdditionalConfiguration. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FeatureAdditionalConfiguration) Values() []FeatureAdditionalConfiguration {
	return []FeatureAdditionalConfiguration{
		"EKS_ADDON_MANAGEMENT",
		"ECS_FARGATE_AGENT_MANAGEMENT",
		"EC2_AGENT_MANAGEMENT",
	}
}

type FeatureStatus string

// Enum values for FeatureStatus
const (
	FeatureStatusEnabled  FeatureStatus = "ENABLED"
	FeatureStatusDisabled FeatureStatus = "DISABLED"
)

// Values returns all known values for FeatureStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FeatureStatus) Values() []FeatureStatus {
	return []FeatureStatus{
		"ENABLED",
		"DISABLED",
	}
}

type Feedback string

// Enum values for Feedback
const (
	FeedbackUseful    Feedback = "USEFUL"
	FeedbackNotUseful Feedback = "NOT_USEFUL"
)

// Values returns all known values for Feedback. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Feedback) Values() []Feedback {
	return []Feedback{
		"USEFUL",
		"NOT_USEFUL",
	}
}

type FilterAction string

// Enum values for FilterAction
const (
	FilterActionNoop    FilterAction = "NOOP"
	FilterActionArchive FilterAction = "ARCHIVE"
)

// Values returns all known values for FilterAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FilterAction) Values() []FilterAction {
	return []FilterAction{
		"NOOP",
		"ARCHIVE",
	}
}

type FindingPublishingFrequency string

// Enum values for FindingPublishingFrequency
const (
	FindingPublishingFrequencyFifteenMinutes FindingPublishingFrequency = "FIFTEEN_MINUTES"
	FindingPublishingFrequencyOneHour        FindingPublishingFrequency = "ONE_HOUR"
	FindingPublishingFrequencySixHours       FindingPublishingFrequency = "SIX_HOURS"
)

// Values returns all known values for FindingPublishingFrequency. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FindingPublishingFrequency) Values() []FindingPublishingFrequency {
	return []FindingPublishingFrequency{
		"FIFTEEN_MINUTES",
		"ONE_HOUR",
		"SIX_HOURS",
	}
}

type FindingStatisticType string

// Enum values for FindingStatisticType
const (
	FindingStatisticTypeCountBySeverity FindingStatisticType = "COUNT_BY_SEVERITY"
)

// Values returns all known values for FindingStatisticType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FindingStatisticType) Values() []FindingStatisticType {
	return []FindingStatisticType{
		"COUNT_BY_SEVERITY",
	}
}

type FreeTrialFeatureResult string

// Enum values for FreeTrialFeatureResult
const (
	FreeTrialFeatureResultFlowLogs                 FreeTrialFeatureResult = "FLOW_LOGS"
	FreeTrialFeatureResultCloudTrail               FreeTrialFeatureResult = "CLOUD_TRAIL"
	FreeTrialFeatureResultDnsLogs                  FreeTrialFeatureResult = "DNS_LOGS"
	FreeTrialFeatureResultS3DataEvents             FreeTrialFeatureResult = "S3_DATA_EVENTS"
	FreeTrialFeatureResultEksAuditLogs             FreeTrialFeatureResult = "EKS_AUDIT_LOGS"
	FreeTrialFeatureResultEbsMalwareProtection     FreeTrialFeatureResult = "EBS_MALWARE_PROTECTION"
	FreeTrialFeatureResultRdsLoginEvents           FreeTrialFeatureResult = "RDS_LOGIN_EVENTS"
	FreeTrialFeatureResultEksRuntimeMonitoring     FreeTrialFeatureResult = "EKS_RUNTIME_MONITORING"
	FreeTrialFeatureResultLambdaNetworkLogs        FreeTrialFeatureResult = "LAMBDA_NETWORK_LOGS"
	FreeTrialFeatureResultFargateRuntimeMonitoring FreeTrialFeatureResult = "FARGATE_RUNTIME_MONITORING"
	FreeTrialFeatureResultEc2RuntimeMonitoring     FreeTrialFeatureResult = "EC2_RUNTIME_MONITORING"
)

// Values returns all known values for FreeTrialFeatureResult. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FreeTrialFeatureResult) Values() []FreeTrialFeatureResult {
	return []FreeTrialFeatureResult{
		"FLOW_LOGS",
		"CLOUD_TRAIL",
		"DNS_LOGS",
		"S3_DATA_EVENTS",
		"EKS_AUDIT_LOGS",
		"EBS_MALWARE_PROTECTION",
		"RDS_LOGIN_EVENTS",
		"EKS_RUNTIME_MONITORING",
		"LAMBDA_NETWORK_LOGS",
		"FARGATE_RUNTIME_MONITORING",
		"EC2_RUNTIME_MONITORING",
	}
}

type IpSetFormat string

// Enum values for IpSetFormat
const (
	IpSetFormatTxt        IpSetFormat = "TXT"
	IpSetFormatStix       IpSetFormat = "STIX"
	IpSetFormatOtxCsv     IpSetFormat = "OTX_CSV"
	IpSetFormatAlienVault IpSetFormat = "ALIEN_VAULT"
	IpSetFormatProofPoint IpSetFormat = "PROOF_POINT"
	IpSetFormatFireEye    IpSetFormat = "FIRE_EYE"
)

// Values returns all known values for IpSetFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IpSetFormat) Values() []IpSetFormat {
	return []IpSetFormat{
		"TXT",
		"STIX",
		"OTX_CSV",
		"ALIEN_VAULT",
		"PROOF_POINT",
		"FIRE_EYE",
	}
}

type IpSetStatus string

// Enum values for IpSetStatus
const (
	IpSetStatusInactive      IpSetStatus = "INACTIVE"
	IpSetStatusActivating    IpSetStatus = "ACTIVATING"
	IpSetStatusActive        IpSetStatus = "ACTIVE"
	IpSetStatusDeactivating  IpSetStatus = "DEACTIVATING"
	IpSetStatusError         IpSetStatus = "ERROR"
	IpSetStatusDeletePending IpSetStatus = "DELETE_PENDING"
	IpSetStatusDeleted       IpSetStatus = "DELETED"
)

// Values returns all known values for IpSetStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IpSetStatus) Values() []IpSetStatus {
	return []IpSetStatus{
		"INACTIVE",
		"ACTIVATING",
		"ACTIVE",
		"DEACTIVATING",
		"ERROR",
		"DELETE_PENDING",
		"DELETED",
	}
}

type MalwareProtectionPlanStatus string

// Enum values for MalwareProtectionPlanStatus
const (
	MalwareProtectionPlanStatusActive  MalwareProtectionPlanStatus = "ACTIVE"
	MalwareProtectionPlanStatusWarning MalwareProtectionPlanStatus = "WARNING"
	MalwareProtectionPlanStatusError   MalwareProtectionPlanStatus = "ERROR"
)

// Values returns all known values for MalwareProtectionPlanStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MalwareProtectionPlanStatus) Values() []MalwareProtectionPlanStatus {
	return []MalwareProtectionPlanStatus{
		"ACTIVE",
		"WARNING",
		"ERROR",
	}
}

type MalwareProtectionPlanTaggingActionStatus string

// Enum values for MalwareProtectionPlanTaggingActionStatus
const (
	MalwareProtectionPlanTaggingActionStatusEnabled  MalwareProtectionPlanTaggingActionStatus = "ENABLED"
	MalwareProtectionPlanTaggingActionStatusDisabled MalwareProtectionPlanTaggingActionStatus = "DISABLED"
)

// Values returns all known values for MalwareProtectionPlanTaggingActionStatus.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MalwareProtectionPlanTaggingActionStatus) Values() []MalwareProtectionPlanTaggingActionStatus {
	return []MalwareProtectionPlanTaggingActionStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ManagementType string

// Enum values for ManagementType
const (
	ManagementTypeAutoManaged ManagementType = "AUTO_MANAGED"
	ManagementTypeManual      ManagementType = "MANUAL"
	ManagementTypeDisabled    ManagementType = "DISABLED"
)

// Values returns all known values for ManagementType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ManagementType) Values() []ManagementType {
	return []ManagementType{
		"AUTO_MANAGED",
		"MANUAL",
		"DISABLED",
	}
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

// Values returns all known values for OrderBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrderBy) Values() []OrderBy {
	return []OrderBy{
		"ASC",
		"DESC",
	}
}

type OrgFeature string

// Enum values for OrgFeature
const (
	OrgFeatureS3DataEvents         OrgFeature = "S3_DATA_EVENTS"
	OrgFeatureEksAuditLogs         OrgFeature = "EKS_AUDIT_LOGS"
	OrgFeatureEbsMalwareProtection OrgFeature = "EBS_MALWARE_PROTECTION"
	OrgFeatureRdsLoginEvents       OrgFeature = "RDS_LOGIN_EVENTS"
	OrgFeatureEksRuntimeMonitoring OrgFeature = "EKS_RUNTIME_MONITORING"
	OrgFeatureLambdaNetworkLogs    OrgFeature = "LAMBDA_NETWORK_LOGS"
	OrgFeatureRuntimeMonitoring    OrgFeature = "RUNTIME_MONITORING"
)

// Values returns all known values for OrgFeature. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrgFeature) Values() []OrgFeature {
	return []OrgFeature{
		"S3_DATA_EVENTS",
		"EKS_AUDIT_LOGS",
		"EBS_MALWARE_PROTECTION",
		"RDS_LOGIN_EVENTS",
		"EKS_RUNTIME_MONITORING",
		"LAMBDA_NETWORK_LOGS",
		"RUNTIME_MONITORING",
	}
}

type OrgFeatureAdditionalConfiguration string

// Enum values for OrgFeatureAdditionalConfiguration
const (
	OrgFeatureAdditionalConfigurationEksAddonManagement        OrgFeatureAdditionalConfiguration = "EKS_ADDON_MANAGEMENT"
	OrgFeatureAdditionalConfigurationEcsFargateAgentManagement OrgFeatureAdditionalConfiguration = "ECS_FARGATE_AGENT_MANAGEMENT"
	OrgFeatureAdditionalConfigurationEc2AgentManagement        OrgFeatureAdditionalConfiguration = "EC2_AGENT_MANAGEMENT"
)

// Values returns all known values for OrgFeatureAdditionalConfiguration. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrgFeatureAdditionalConfiguration) Values() []OrgFeatureAdditionalConfiguration {
	return []OrgFeatureAdditionalConfiguration{
		"EKS_ADDON_MANAGEMENT",
		"ECS_FARGATE_AGENT_MANAGEMENT",
		"EC2_AGENT_MANAGEMENT",
	}
}

type OrgFeatureStatus string

// Enum values for OrgFeatureStatus
const (
	OrgFeatureStatusNew  OrgFeatureStatus = "NEW"
	OrgFeatureStatusNone OrgFeatureStatus = "NONE"
	OrgFeatureStatusAll  OrgFeatureStatus = "ALL"
)

// Values returns all known values for OrgFeatureStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrgFeatureStatus) Values() []OrgFeatureStatus {
	return []OrgFeatureStatus{
		"NEW",
		"NONE",
		"ALL",
	}
}

type ProfileSubtype string

// Enum values for ProfileSubtype
const (
	ProfileSubtypeFrequent   ProfileSubtype = "FREQUENT"
	ProfileSubtypeInfrequent ProfileSubtype = "INFREQUENT"
	ProfileSubtypeUnseen     ProfileSubtype = "UNSEEN"
	ProfileSubtypeRare       ProfileSubtype = "RARE"
)

// Values returns all known values for ProfileSubtype. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProfileSubtype) Values() []ProfileSubtype {
	return []ProfileSubtype{
		"FREQUENT",
		"INFREQUENT",
		"UNSEEN",
		"RARE",
	}
}

type ProfileType string

// Enum values for ProfileType
const (
	ProfileTypeFrequency ProfileType = "FREQUENCY"
)

// Values returns all known values for ProfileType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProfileType) Values() []ProfileType {
	return []ProfileType{
		"FREQUENCY",
	}
}

type PublishingStatus string

// Enum values for PublishingStatus
const (
	PublishingStatusPendingVerification                   PublishingStatus = "PENDING_VERIFICATION"
	PublishingStatusPublishing                            PublishingStatus = "PUBLISHING"
	PublishingStatusUnableToPublishFixDestinationProperty PublishingStatus = "UNABLE_TO_PUBLISH_FIX_DESTINATION_PROPERTY"
	PublishingStatusStopped                               PublishingStatus = "STOPPED"
)

// Values returns all known values for PublishingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PublishingStatus) Values() []PublishingStatus {
	return []PublishingStatus{
		"PENDING_VERIFICATION",
		"PUBLISHING",
		"UNABLE_TO_PUBLISH_FIX_DESTINATION_PROPERTY",
		"STOPPED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeEks ResourceType = "EKS"
	ResourceTypeEcs ResourceType = "ECS"
	ResourceTypeEc2 ResourceType = "EC2"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"EKS",
		"ECS",
		"EC2",
	}
}

type ScanCriterionKey string

// Enum values for ScanCriterionKey
const (
	ScanCriterionKeyEc2InstanceTag ScanCriterionKey = "EC2_INSTANCE_TAG"
)

// Values returns all known values for ScanCriterionKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScanCriterionKey) Values() []ScanCriterionKey {
	return []ScanCriterionKey{
		"EC2_INSTANCE_TAG",
	}
}

type ScanResult string

// Enum values for ScanResult
const (
	ScanResultClean    ScanResult = "CLEAN"
	ScanResultInfected ScanResult = "INFECTED"
)

// Values returns all known values for ScanResult. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScanResult) Values() []ScanResult {
	return []ScanResult{
		"CLEAN",
		"INFECTED",
	}
}

type ScanStatus string

// Enum values for ScanStatus
const (
	ScanStatusRunning   ScanStatus = "RUNNING"
	ScanStatusCompleted ScanStatus = "COMPLETED"
	ScanStatusFailed    ScanStatus = "FAILED"
	ScanStatusSkipped   ScanStatus = "SKIPPED"
)

// Values returns all known values for ScanStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScanStatus) Values() []ScanStatus {
	return []ScanStatus{
		"RUNNING",
		"COMPLETED",
		"FAILED",
		"SKIPPED",
	}
}

type ScanType string

// Enum values for ScanType
const (
	ScanTypeGuarddutyInitiated ScanType = "GUARDDUTY_INITIATED"
	ScanTypeOnDemand           ScanType = "ON_DEMAND"
)

// Values returns all known values for ScanType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScanType) Values() []ScanType {
	return []ScanType{
		"GUARDDUTY_INITIATED",
		"ON_DEMAND",
	}
}

type ThreatIntelSetFormat string

// Enum values for ThreatIntelSetFormat
const (
	ThreatIntelSetFormatTxt        ThreatIntelSetFormat = "TXT"
	ThreatIntelSetFormatStix       ThreatIntelSetFormat = "STIX"
	ThreatIntelSetFormatOtxCsv     ThreatIntelSetFormat = "OTX_CSV"
	ThreatIntelSetFormatAlienVault ThreatIntelSetFormat = "ALIEN_VAULT"
	ThreatIntelSetFormatProofPoint ThreatIntelSetFormat = "PROOF_POINT"
	ThreatIntelSetFormatFireEye    ThreatIntelSetFormat = "FIRE_EYE"
)

// Values returns all known values for ThreatIntelSetFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThreatIntelSetFormat) Values() []ThreatIntelSetFormat {
	return []ThreatIntelSetFormat{
		"TXT",
		"STIX",
		"OTX_CSV",
		"ALIEN_VAULT",
		"PROOF_POINT",
		"FIRE_EYE",
	}
}

type ThreatIntelSetStatus string

// Enum values for ThreatIntelSetStatus
const (
	ThreatIntelSetStatusInactive      ThreatIntelSetStatus = "INACTIVE"
	ThreatIntelSetStatusActivating    ThreatIntelSetStatus = "ACTIVATING"
	ThreatIntelSetStatusActive        ThreatIntelSetStatus = "ACTIVE"
	ThreatIntelSetStatusDeactivating  ThreatIntelSetStatus = "DEACTIVATING"
	ThreatIntelSetStatusError         ThreatIntelSetStatus = "ERROR"
	ThreatIntelSetStatusDeletePending ThreatIntelSetStatus = "DELETE_PENDING"
	ThreatIntelSetStatusDeleted       ThreatIntelSetStatus = "DELETED"
)

// Values returns all known values for ThreatIntelSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThreatIntelSetStatus) Values() []ThreatIntelSetStatus {
	return []ThreatIntelSetStatus{
		"INACTIVE",
		"ACTIVATING",
		"ACTIVE",
		"DEACTIVATING",
		"ERROR",
		"DELETE_PENDING",
		"DELETED",
	}
}

type UsageFeature string

// Enum values for UsageFeature
const (
	UsageFeatureFlowLogs                    UsageFeature = "FLOW_LOGS"
	UsageFeatureCloudTrail                  UsageFeature = "CLOUD_TRAIL"
	UsageFeatureDnsLogs                     UsageFeature = "DNS_LOGS"
	UsageFeatureS3DataEvents                UsageFeature = "S3_DATA_EVENTS"
	UsageFeatureEksAuditLogs                UsageFeature = "EKS_AUDIT_LOGS"
	UsageFeatureEbsMalwareProtection        UsageFeature = "EBS_MALWARE_PROTECTION"
	UsageFeatureRdsLoginEvents              UsageFeature = "RDS_LOGIN_EVENTS"
	UsageFeatureLambdaNetworkLogs           UsageFeature = "LAMBDA_NETWORK_LOGS"
	UsageFeatureEksRuntimeMonitoring        UsageFeature = "EKS_RUNTIME_MONITORING"
	UsageFeatureFargateRuntimeMonitoring    UsageFeature = "FARGATE_RUNTIME_MONITORING"
	UsageFeatureEc2RuntimeMonitoring        UsageFeature = "EC2_RUNTIME_MONITORING"
	UsageFeatureRdsDbiProtectionProvisioned UsageFeature = "RDS_DBI_PROTECTION_PROVISIONED"
	UsageFeatureRdsDbiProtectionServerless  UsageFeature = "RDS_DBI_PROTECTION_SERVERLESS"
)

// Values returns all known values for UsageFeature. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UsageFeature) Values() []UsageFeature {
	return []UsageFeature{
		"FLOW_LOGS",
		"CLOUD_TRAIL",
		"DNS_LOGS",
		"S3_DATA_EVENTS",
		"EKS_AUDIT_LOGS",
		"EBS_MALWARE_PROTECTION",
		"RDS_LOGIN_EVENTS",
		"LAMBDA_NETWORK_LOGS",
		"EKS_RUNTIME_MONITORING",
		"FARGATE_RUNTIME_MONITORING",
		"EC2_RUNTIME_MONITORING",
		"RDS_DBI_PROTECTION_PROVISIONED",
		"RDS_DBI_PROTECTION_SERVERLESS",
	}
}

type UsageStatisticType string

// Enum values for UsageStatisticType
const (
	UsageStatisticTypeSumByAccount         UsageStatisticType = "SUM_BY_ACCOUNT"
	UsageStatisticTypeSumByDataSource      UsageStatisticType = "SUM_BY_DATA_SOURCE"
	UsageStatisticTypeSumByResource        UsageStatisticType = "SUM_BY_RESOURCE"
	UsageStatisticTypeTopResources         UsageStatisticType = "TOP_RESOURCES"
	UsageStatisticTypeSumByFeatures        UsageStatisticType = "SUM_BY_FEATURES"
	UsageStatisticTypeTopAccountsByFeature UsageStatisticType = "TOP_ACCOUNTS_BY_FEATURE"
)

// Values returns all known values for UsageStatisticType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UsageStatisticType) Values() []UsageStatisticType {
	return []UsageStatisticType{
		"SUM_BY_ACCOUNT",
		"SUM_BY_DATA_SOURCE",
		"SUM_BY_RESOURCE",
		"TOP_RESOURCES",
		"SUM_BY_FEATURES",
		"TOP_ACCOUNTS_BY_FEATURE",
	}
}