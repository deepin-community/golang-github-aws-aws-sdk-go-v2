// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AsyncOperationName string

// Enum values for AsyncOperationName
const (
	AsyncOperationNameCreateMultiRegionAccessPoint    AsyncOperationName = "CreateMultiRegionAccessPoint"
	AsyncOperationNameDeleteMultiRegionAccessPoint    AsyncOperationName = "DeleteMultiRegionAccessPoint"
	AsyncOperationNamePutMultiRegionAccessPointPolicy AsyncOperationName = "PutMultiRegionAccessPointPolicy"
)

// Values returns all known values for AsyncOperationName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AsyncOperationName) Values() []AsyncOperationName {
	return []AsyncOperationName{
		"CreateMultiRegionAccessPoint",
		"DeleteMultiRegionAccessPoint",
		"PutMultiRegionAccessPointPolicy",
	}
}

type BucketCannedACL string

// Enum values for BucketCannedACL
const (
	BucketCannedACLPrivate           BucketCannedACL = "private"
	BucketCannedACLPublicRead        BucketCannedACL = "public-read"
	BucketCannedACLPublicReadWrite   BucketCannedACL = "public-read-write"
	BucketCannedACLAuthenticatedRead BucketCannedACL = "authenticated-read"
)

// Values returns all known values for BucketCannedACL. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BucketCannedACL) Values() []BucketCannedACL {
	return []BucketCannedACL{
		"private",
		"public-read",
		"public-read-write",
		"authenticated-read",
	}
}

type BucketLocationConstraint string

// Enum values for BucketLocationConstraint
const (
	BucketLocationConstraintEu           BucketLocationConstraint = "EU"
	BucketLocationConstraintEuWest1      BucketLocationConstraint = "eu-west-1"
	BucketLocationConstraintUsWest1      BucketLocationConstraint = "us-west-1"
	BucketLocationConstraintUsWest2      BucketLocationConstraint = "us-west-2"
	BucketLocationConstraintApSouth1     BucketLocationConstraint = "ap-south-1"
	BucketLocationConstraintApSoutheast1 BucketLocationConstraint = "ap-southeast-1"
	BucketLocationConstraintApSoutheast2 BucketLocationConstraint = "ap-southeast-2"
	BucketLocationConstraintApNortheast1 BucketLocationConstraint = "ap-northeast-1"
	BucketLocationConstraintSaEast1      BucketLocationConstraint = "sa-east-1"
	BucketLocationConstraintCnNorth1     BucketLocationConstraint = "cn-north-1"
	BucketLocationConstraintEuCentral1   BucketLocationConstraint = "eu-central-1"
)

// Values returns all known values for BucketLocationConstraint. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BucketLocationConstraint) Values() []BucketLocationConstraint {
	return []BucketLocationConstraint{
		"EU",
		"eu-west-1",
		"us-west-1",
		"us-west-2",
		"ap-south-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-northeast-1",
		"sa-east-1",
		"cn-north-1",
		"eu-central-1",
	}
}

type BucketVersioningStatus string

// Enum values for BucketVersioningStatus
const (
	BucketVersioningStatusEnabled   BucketVersioningStatus = "Enabled"
	BucketVersioningStatusSuspended BucketVersioningStatus = "Suspended"
)

// Values returns all known values for BucketVersioningStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BucketVersioningStatus) Values() []BucketVersioningStatus {
	return []BucketVersioningStatus{
		"Enabled",
		"Suspended",
	}
}

type DeleteMarkerReplicationStatus string

// Enum values for DeleteMarkerReplicationStatus
const (
	DeleteMarkerReplicationStatusEnabled  DeleteMarkerReplicationStatus = "Enabled"
	DeleteMarkerReplicationStatusDisabled DeleteMarkerReplicationStatus = "Disabled"
)

// Values returns all known values for DeleteMarkerReplicationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeleteMarkerReplicationStatus) Values() []DeleteMarkerReplicationStatus {
	return []DeleteMarkerReplicationStatus{
		"Enabled",
		"Disabled",
	}
}

type ExistingObjectReplicationStatus string

// Enum values for ExistingObjectReplicationStatus
const (
	ExistingObjectReplicationStatusEnabled  ExistingObjectReplicationStatus = "Enabled"
	ExistingObjectReplicationStatusDisabled ExistingObjectReplicationStatus = "Disabled"
)

// Values returns all known values for ExistingObjectReplicationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExistingObjectReplicationStatus) Values() []ExistingObjectReplicationStatus {
	return []ExistingObjectReplicationStatus{
		"Enabled",
		"Disabled",
	}
}

type ExpirationStatus string

// Enum values for ExpirationStatus
const (
	ExpirationStatusEnabled  ExpirationStatus = "Enabled"
	ExpirationStatusDisabled ExpirationStatus = "Disabled"
)

// Values returns all known values for ExpirationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExpirationStatus) Values() []ExpirationStatus {
	return []ExpirationStatus{
		"Enabled",
		"Disabled",
	}
}

type Format string

// Enum values for Format
const (
	FormatCsv     Format = "CSV"
	FormatParquet Format = "Parquet"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"CSV",
		"Parquet",
	}
}

type GeneratedManifestFormat string

// Enum values for GeneratedManifestFormat
const (
	GeneratedManifestFormatS3InventoryReportCsv20211130 GeneratedManifestFormat = "S3InventoryReport_CSV_20211130"
)

// Values returns all known values for GeneratedManifestFormat. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GeneratedManifestFormat) Values() []GeneratedManifestFormat {
	return []GeneratedManifestFormat{
		"S3InventoryReport_CSV_20211130",
	}
}

type GranteeType string

// Enum values for GranteeType
const (
	GranteeTypeDirectoryUser  GranteeType = "DIRECTORY_USER"
	GranteeTypeDirectoryGroup GranteeType = "DIRECTORY_GROUP"
	GranteeTypeIam            GranteeType = "IAM"
)

// Values returns all known values for GranteeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GranteeType) Values() []GranteeType {
	return []GranteeType{
		"DIRECTORY_USER",
		"DIRECTORY_GROUP",
		"IAM",
	}
}

type JobManifestFieldName string

// Enum values for JobManifestFieldName
const (
	JobManifestFieldNameIgnore    JobManifestFieldName = "Ignore"
	JobManifestFieldNameBucket    JobManifestFieldName = "Bucket"
	JobManifestFieldNameKey       JobManifestFieldName = "Key"
	JobManifestFieldNameVersionId JobManifestFieldName = "VersionId"
)

// Values returns all known values for JobManifestFieldName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobManifestFieldName) Values() []JobManifestFieldName {
	return []JobManifestFieldName{
		"Ignore",
		"Bucket",
		"Key",
		"VersionId",
	}
}

type JobManifestFormat string

// Enum values for JobManifestFormat
const (
	JobManifestFormatS3BatchOperationsCsv20180820 JobManifestFormat = "S3BatchOperations_CSV_20180820"
	JobManifestFormatS3InventoryReportCsv20161130 JobManifestFormat = "S3InventoryReport_CSV_20161130"
)

// Values returns all known values for JobManifestFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobManifestFormat) Values() []JobManifestFormat {
	return []JobManifestFormat{
		"S3BatchOperations_CSV_20180820",
		"S3InventoryReport_CSV_20161130",
	}
}

type JobReportFormat string

// Enum values for JobReportFormat
const (
	JobReportFormatReportCsv20180820 JobReportFormat = "Report_CSV_20180820"
)

// Values returns all known values for JobReportFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobReportFormat) Values() []JobReportFormat {
	return []JobReportFormat{
		"Report_CSV_20180820",
	}
}

type JobReportScope string

// Enum values for JobReportScope
const (
	JobReportScopeAllTasks        JobReportScope = "AllTasks"
	JobReportScopeFailedTasksOnly JobReportScope = "FailedTasksOnly"
)

// Values returns all known values for JobReportScope. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobReportScope) Values() []JobReportScope {
	return []JobReportScope{
		"AllTasks",
		"FailedTasksOnly",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusActive     JobStatus = "Active"
	JobStatusCancelled  JobStatus = "Cancelled"
	JobStatusCancelling JobStatus = "Cancelling"
	JobStatusComplete   JobStatus = "Complete"
	JobStatusCompleting JobStatus = "Completing"
	JobStatusFailed     JobStatus = "Failed"
	JobStatusFailing    JobStatus = "Failing"
	JobStatusNew        JobStatus = "New"
	JobStatusPaused     JobStatus = "Paused"
	JobStatusPausing    JobStatus = "Pausing"
	JobStatusPreparing  JobStatus = "Preparing"
	JobStatusReady      JobStatus = "Ready"
	JobStatusSuspended  JobStatus = "Suspended"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"Active",
		"Cancelled",
		"Cancelling",
		"Complete",
		"Completing",
		"Failed",
		"Failing",
		"New",
		"Paused",
		"Pausing",
		"Preparing",
		"Ready",
		"Suspended",
	}
}

type MetricsStatus string

// Enum values for MetricsStatus
const (
	MetricsStatusEnabled  MetricsStatus = "Enabled"
	MetricsStatusDisabled MetricsStatus = "Disabled"
)

// Values returns all known values for MetricsStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MetricsStatus) Values() []MetricsStatus {
	return []MetricsStatus{
		"Enabled",
		"Disabled",
	}
}

type MFADelete string

// Enum values for MFADelete
const (
	MFADeleteEnabled  MFADelete = "Enabled"
	MFADeleteDisabled MFADelete = "Disabled"
)

// Values returns all known values for MFADelete. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MFADelete) Values() []MFADelete {
	return []MFADelete{
		"Enabled",
		"Disabled",
	}
}

type MFADeleteStatus string

// Enum values for MFADeleteStatus
const (
	MFADeleteStatusEnabled  MFADeleteStatus = "Enabled"
	MFADeleteStatusDisabled MFADeleteStatus = "Disabled"
)

// Values returns all known values for MFADeleteStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MFADeleteStatus) Values() []MFADeleteStatus {
	return []MFADeleteStatus{
		"Enabled",
		"Disabled",
	}
}

type MultiRegionAccessPointStatus string

// Enum values for MultiRegionAccessPointStatus
const (
	MultiRegionAccessPointStatusReady                     MultiRegionAccessPointStatus = "READY"
	MultiRegionAccessPointStatusInconsistentAcrossRegions MultiRegionAccessPointStatus = "INCONSISTENT_ACROSS_REGIONS"
	MultiRegionAccessPointStatusCreating                  MultiRegionAccessPointStatus = "CREATING"
	MultiRegionAccessPointStatusPartiallyCreated          MultiRegionAccessPointStatus = "PARTIALLY_CREATED"
	MultiRegionAccessPointStatusPartiallyDeleted          MultiRegionAccessPointStatus = "PARTIALLY_DELETED"
	MultiRegionAccessPointStatusDeleting                  MultiRegionAccessPointStatus = "DELETING"
)

// Values returns all known values for MultiRegionAccessPointStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MultiRegionAccessPointStatus) Values() []MultiRegionAccessPointStatus {
	return []MultiRegionAccessPointStatus{
		"READY",
		"INCONSISTENT_ACROSS_REGIONS",
		"CREATING",
		"PARTIALLY_CREATED",
		"PARTIALLY_DELETED",
		"DELETING",
	}
}

type NetworkOrigin string

// Enum values for NetworkOrigin
const (
	NetworkOriginInternet NetworkOrigin = "Internet"
	NetworkOriginVpc      NetworkOrigin = "VPC"
)

// Values returns all known values for NetworkOrigin. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NetworkOrigin) Values() []NetworkOrigin {
	return []NetworkOrigin{
		"Internet",
		"VPC",
	}
}

type ObjectLambdaAccessPointAliasStatus string

// Enum values for ObjectLambdaAccessPointAliasStatus
const (
	ObjectLambdaAccessPointAliasStatusProvisioning ObjectLambdaAccessPointAliasStatus = "PROVISIONING"
	ObjectLambdaAccessPointAliasStatusReady        ObjectLambdaAccessPointAliasStatus = "READY"
)

// Values returns all known values for ObjectLambdaAccessPointAliasStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectLambdaAccessPointAliasStatus) Values() []ObjectLambdaAccessPointAliasStatus {
	return []ObjectLambdaAccessPointAliasStatus{
		"PROVISIONING",
		"READY",
	}
}

type ObjectLambdaAllowedFeature string

// Enum values for ObjectLambdaAllowedFeature
const (
	ObjectLambdaAllowedFeatureGetObjectRange       ObjectLambdaAllowedFeature = "GetObject-Range"
	ObjectLambdaAllowedFeatureGetObjectPartNumber  ObjectLambdaAllowedFeature = "GetObject-PartNumber"
	ObjectLambdaAllowedFeatureHeadObjectRange      ObjectLambdaAllowedFeature = "HeadObject-Range"
	ObjectLambdaAllowedFeatureHeadObjectPartNumber ObjectLambdaAllowedFeature = "HeadObject-PartNumber"
)

// Values returns all known values for ObjectLambdaAllowedFeature. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectLambdaAllowedFeature) Values() []ObjectLambdaAllowedFeature {
	return []ObjectLambdaAllowedFeature{
		"GetObject-Range",
		"GetObject-PartNumber",
		"HeadObject-Range",
		"HeadObject-PartNumber",
	}
}

type ObjectLambdaTransformationConfigurationAction string

// Enum values for ObjectLambdaTransformationConfigurationAction
const (
	ObjectLambdaTransformationConfigurationActionGetObject     ObjectLambdaTransformationConfigurationAction = "GetObject"
	ObjectLambdaTransformationConfigurationActionHeadObject    ObjectLambdaTransformationConfigurationAction = "HeadObject"
	ObjectLambdaTransformationConfigurationActionListObjects   ObjectLambdaTransformationConfigurationAction = "ListObjects"
	ObjectLambdaTransformationConfigurationActionListObjectsV2 ObjectLambdaTransformationConfigurationAction = "ListObjectsV2"
)

// Values returns all known values for
// ObjectLambdaTransformationConfigurationAction. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectLambdaTransformationConfigurationAction) Values() []ObjectLambdaTransformationConfigurationAction {
	return []ObjectLambdaTransformationConfigurationAction{
		"GetObject",
		"HeadObject",
		"ListObjects",
		"ListObjectsV2",
	}
}

type OperationName string

// Enum values for OperationName
const (
	OperationNameLambdaInvoke            OperationName = "LambdaInvoke"
	OperationNameS3PutObjectCopy         OperationName = "S3PutObjectCopy"
	OperationNameS3PutObjectAcl          OperationName = "S3PutObjectAcl"
	OperationNameS3PutObjectTagging      OperationName = "S3PutObjectTagging"
	OperationNameS3DeleteObjectTagging   OperationName = "S3DeleteObjectTagging"
	OperationNameS3InitiateRestoreObject OperationName = "S3InitiateRestoreObject"
	OperationNameS3PutObjectLegalHold    OperationName = "S3PutObjectLegalHold"
	OperationNameS3PutObjectRetention    OperationName = "S3PutObjectRetention"
	OperationNameS3ReplicateObject       OperationName = "S3ReplicateObject"
)

// Values returns all known values for OperationName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OperationName) Values() []OperationName {
	return []OperationName{
		"LambdaInvoke",
		"S3PutObjectCopy",
		"S3PutObjectAcl",
		"S3PutObjectTagging",
		"S3DeleteObjectTagging",
		"S3InitiateRestoreObject",
		"S3PutObjectLegalHold",
		"S3PutObjectRetention",
		"S3ReplicateObject",
	}
}

type OutputSchemaVersion string

// Enum values for OutputSchemaVersion
const (
	OutputSchemaVersionV1 OutputSchemaVersion = "V_1"
)

// Values returns all known values for OutputSchemaVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OutputSchemaVersion) Values() []OutputSchemaVersion {
	return []OutputSchemaVersion{
		"V_1",
	}
}

type OwnerOverride string

// Enum values for OwnerOverride
const (
	OwnerOverrideDestination OwnerOverride = "Destination"
)

// Values returns all known values for OwnerOverride. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OwnerOverride) Values() []OwnerOverride {
	return []OwnerOverride{
		"Destination",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionRead      Permission = "READ"
	PermissionWrite     Permission = "WRITE"
	PermissionReadwrite Permission = "READWRITE"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"READ",
		"WRITE",
		"READWRITE",
	}
}

type Privilege string

// Enum values for Privilege
const (
	PrivilegeMinimal Privilege = "Minimal"
	PrivilegeDefault Privilege = "Default"
)

// Values returns all known values for Privilege. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Privilege) Values() []Privilege {
	return []Privilege{
		"Minimal",
		"Default",
	}
}

type ReplicaModificationsStatus string

// Enum values for ReplicaModificationsStatus
const (
	ReplicaModificationsStatusEnabled  ReplicaModificationsStatus = "Enabled"
	ReplicaModificationsStatusDisabled ReplicaModificationsStatus = "Disabled"
)

// Values returns all known values for ReplicaModificationsStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicaModificationsStatus) Values() []ReplicaModificationsStatus {
	return []ReplicaModificationsStatus{
		"Enabled",
		"Disabled",
	}
}

type ReplicationRuleStatus string

// Enum values for ReplicationRuleStatus
const (
	ReplicationRuleStatusEnabled  ReplicationRuleStatus = "Enabled"
	ReplicationRuleStatusDisabled ReplicationRuleStatus = "Disabled"
)

// Values returns all known values for ReplicationRuleStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationRuleStatus) Values() []ReplicationRuleStatus {
	return []ReplicationRuleStatus{
		"Enabled",
		"Disabled",
	}
}

type ReplicationStatus string

// Enum values for ReplicationStatus
const (
	ReplicationStatusCompleted ReplicationStatus = "COMPLETED"
	ReplicationStatusFailed    ReplicationStatus = "FAILED"
	ReplicationStatusReplica   ReplicationStatus = "REPLICA"
	ReplicationStatusNone      ReplicationStatus = "NONE"
)

// Values returns all known values for ReplicationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationStatus) Values() []ReplicationStatus {
	return []ReplicationStatus{
		"COMPLETED",
		"FAILED",
		"REPLICA",
		"NONE",
	}
}

type ReplicationStorageClass string

// Enum values for ReplicationStorageClass
const (
	ReplicationStorageClassStandard           ReplicationStorageClass = "STANDARD"
	ReplicationStorageClassReducedRedundancy  ReplicationStorageClass = "REDUCED_REDUNDANCY"
	ReplicationStorageClassStandardIa         ReplicationStorageClass = "STANDARD_IA"
	ReplicationStorageClassOnezoneIa          ReplicationStorageClass = "ONEZONE_IA"
	ReplicationStorageClassIntelligentTiering ReplicationStorageClass = "INTELLIGENT_TIERING"
	ReplicationStorageClassGlacier            ReplicationStorageClass = "GLACIER"
	ReplicationStorageClassDeepArchive        ReplicationStorageClass = "DEEP_ARCHIVE"
	ReplicationStorageClassOutposts           ReplicationStorageClass = "OUTPOSTS"
	ReplicationStorageClassGlacierIr          ReplicationStorageClass = "GLACIER_IR"
)

// Values returns all known values for ReplicationStorageClass. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationStorageClass) Values() []ReplicationStorageClass {
	return []ReplicationStorageClass{
		"STANDARD",
		"REDUCED_REDUNDANCY",
		"STANDARD_IA",
		"ONEZONE_IA",
		"INTELLIGENT_TIERING",
		"GLACIER",
		"DEEP_ARCHIVE",
		"OUTPOSTS",
		"GLACIER_IR",
	}
}

type ReplicationTimeStatus string

// Enum values for ReplicationTimeStatus
const (
	ReplicationTimeStatusEnabled  ReplicationTimeStatus = "Enabled"
	ReplicationTimeStatusDisabled ReplicationTimeStatus = "Disabled"
)

// Values returns all known values for ReplicationTimeStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationTimeStatus) Values() []ReplicationTimeStatus {
	return []ReplicationTimeStatus{
		"Enabled",
		"Disabled",
	}
}

type RequestedJobStatus string

// Enum values for RequestedJobStatus
const (
	RequestedJobStatusCancelled RequestedJobStatus = "Cancelled"
	RequestedJobStatusReady     RequestedJobStatus = "Ready"
)

// Values returns all known values for RequestedJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RequestedJobStatus) Values() []RequestedJobStatus {
	return []RequestedJobStatus{
		"Cancelled",
		"Ready",
	}
}

type S3CannedAccessControlList string

// Enum values for S3CannedAccessControlList
const (
	S3CannedAccessControlListPrivate                S3CannedAccessControlList = "private"
	S3CannedAccessControlListPublicRead             S3CannedAccessControlList = "public-read"
	S3CannedAccessControlListPublicReadWrite        S3CannedAccessControlList = "public-read-write"
	S3CannedAccessControlListAwsExecRead            S3CannedAccessControlList = "aws-exec-read"
	S3CannedAccessControlListAuthenticatedRead      S3CannedAccessControlList = "authenticated-read"
	S3CannedAccessControlListBucketOwnerRead        S3CannedAccessControlList = "bucket-owner-read"
	S3CannedAccessControlListBucketOwnerFullControl S3CannedAccessControlList = "bucket-owner-full-control"
)

// Values returns all known values for S3CannedAccessControlList. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3CannedAccessControlList) Values() []S3CannedAccessControlList {
	return []S3CannedAccessControlList{
		"private",
		"public-read",
		"public-read-write",
		"aws-exec-read",
		"authenticated-read",
		"bucket-owner-read",
		"bucket-owner-full-control",
	}
}

type S3ChecksumAlgorithm string

// Enum values for S3ChecksumAlgorithm
const (
	S3ChecksumAlgorithmCrc32  S3ChecksumAlgorithm = "CRC32"
	S3ChecksumAlgorithmCrc32c S3ChecksumAlgorithm = "CRC32C"
	S3ChecksumAlgorithmSha1   S3ChecksumAlgorithm = "SHA1"
	S3ChecksumAlgorithmSha256 S3ChecksumAlgorithm = "SHA256"
)

// Values returns all known values for S3ChecksumAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3ChecksumAlgorithm) Values() []S3ChecksumAlgorithm {
	return []S3ChecksumAlgorithm{
		"CRC32",
		"CRC32C",
		"SHA1",
		"SHA256",
	}
}

type S3GlacierJobTier string

// Enum values for S3GlacierJobTier
const (
	S3GlacierJobTierBulk     S3GlacierJobTier = "BULK"
	S3GlacierJobTierStandard S3GlacierJobTier = "STANDARD"
)

// Values returns all known values for S3GlacierJobTier. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3GlacierJobTier) Values() []S3GlacierJobTier {
	return []S3GlacierJobTier{
		"BULK",
		"STANDARD",
	}
}

type S3GranteeTypeIdentifier string

// Enum values for S3GranteeTypeIdentifier
const (
	S3GranteeTypeIdentifierCanonical    S3GranteeTypeIdentifier = "id"
	S3GranteeTypeIdentifierEmailAddress S3GranteeTypeIdentifier = "emailAddress"
	S3GranteeTypeIdentifierGroup        S3GranteeTypeIdentifier = "uri"
)

// Values returns all known values for S3GranteeTypeIdentifier. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3GranteeTypeIdentifier) Values() []S3GranteeTypeIdentifier {
	return []S3GranteeTypeIdentifier{
		"id",
		"emailAddress",
		"uri",
	}
}

type S3MetadataDirective string

// Enum values for S3MetadataDirective
const (
	S3MetadataDirectiveCopy    S3MetadataDirective = "COPY"
	S3MetadataDirectiveReplace S3MetadataDirective = "REPLACE"
)

// Values returns all known values for S3MetadataDirective. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3MetadataDirective) Values() []S3MetadataDirective {
	return []S3MetadataDirective{
		"COPY",
		"REPLACE",
	}
}

type S3ObjectLockLegalHoldStatus string

// Enum values for S3ObjectLockLegalHoldStatus
const (
	S3ObjectLockLegalHoldStatusOff S3ObjectLockLegalHoldStatus = "OFF"
	S3ObjectLockLegalHoldStatusOn  S3ObjectLockLegalHoldStatus = "ON"
)

// Values returns all known values for S3ObjectLockLegalHoldStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3ObjectLockLegalHoldStatus) Values() []S3ObjectLockLegalHoldStatus {
	return []S3ObjectLockLegalHoldStatus{
		"OFF",
		"ON",
	}
}

type S3ObjectLockMode string

// Enum values for S3ObjectLockMode
const (
	S3ObjectLockModeCompliance S3ObjectLockMode = "COMPLIANCE"
	S3ObjectLockModeGovernance S3ObjectLockMode = "GOVERNANCE"
)

// Values returns all known values for S3ObjectLockMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3ObjectLockMode) Values() []S3ObjectLockMode {
	return []S3ObjectLockMode{
		"COMPLIANCE",
		"GOVERNANCE",
	}
}

type S3ObjectLockRetentionMode string

// Enum values for S3ObjectLockRetentionMode
const (
	S3ObjectLockRetentionModeCompliance S3ObjectLockRetentionMode = "COMPLIANCE"
	S3ObjectLockRetentionModeGovernance S3ObjectLockRetentionMode = "GOVERNANCE"
)

// Values returns all known values for S3ObjectLockRetentionMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3ObjectLockRetentionMode) Values() []S3ObjectLockRetentionMode {
	return []S3ObjectLockRetentionMode{
		"COMPLIANCE",
		"GOVERNANCE",
	}
}

type S3Permission string

// Enum values for S3Permission
const (
	S3PermissionFullControl S3Permission = "FULL_CONTROL"
	S3PermissionRead        S3Permission = "READ"
	S3PermissionWrite       S3Permission = "WRITE"
	S3PermissionReadAcp     S3Permission = "READ_ACP"
	S3PermissionWriteAcp    S3Permission = "WRITE_ACP"
)

// Values returns all known values for S3Permission. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3Permission) Values() []S3Permission {
	return []S3Permission{
		"FULL_CONTROL",
		"READ",
		"WRITE",
		"READ_ACP",
		"WRITE_ACP",
	}
}

type S3PrefixType string

// Enum values for S3PrefixType
const (
	S3PrefixTypeObject S3PrefixType = "Object"
)

// Values returns all known values for S3PrefixType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3PrefixType) Values() []S3PrefixType {
	return []S3PrefixType{
		"Object",
	}
}

type S3SSEAlgorithm string

// Enum values for S3SSEAlgorithm
const (
	S3SSEAlgorithmAes256 S3SSEAlgorithm = "AES256"
	S3SSEAlgorithmKms    S3SSEAlgorithm = "KMS"
)

// Values returns all known values for S3SSEAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3SSEAlgorithm) Values() []S3SSEAlgorithm {
	return []S3SSEAlgorithm{
		"AES256",
		"KMS",
	}
}

type S3StorageClass string

// Enum values for S3StorageClass
const (
	S3StorageClassStandard           S3StorageClass = "STANDARD"
	S3StorageClassStandardIa         S3StorageClass = "STANDARD_IA"
	S3StorageClassOnezoneIa          S3StorageClass = "ONEZONE_IA"
	S3StorageClassGlacier            S3StorageClass = "GLACIER"
	S3StorageClassIntelligentTiering S3StorageClass = "INTELLIGENT_TIERING"
	S3StorageClassDeepArchive        S3StorageClass = "DEEP_ARCHIVE"
	S3StorageClassGlacierIr          S3StorageClass = "GLACIER_IR"
)

// Values returns all known values for S3StorageClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3StorageClass) Values() []S3StorageClass {
	return []S3StorageClass{
		"STANDARD",
		"STANDARD_IA",
		"ONEZONE_IA",
		"GLACIER",
		"INTELLIGENT_TIERING",
		"DEEP_ARCHIVE",
		"GLACIER_IR",
	}
}

type SseKmsEncryptedObjectsStatus string

// Enum values for SseKmsEncryptedObjectsStatus
const (
	SseKmsEncryptedObjectsStatusEnabled  SseKmsEncryptedObjectsStatus = "Enabled"
	SseKmsEncryptedObjectsStatusDisabled SseKmsEncryptedObjectsStatus = "Disabled"
)

// Values returns all known values for SseKmsEncryptedObjectsStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SseKmsEncryptedObjectsStatus) Values() []SseKmsEncryptedObjectsStatus {
	return []SseKmsEncryptedObjectsStatus{
		"Enabled",
		"Disabled",
	}
}

type TransitionStorageClass string

// Enum values for TransitionStorageClass
const (
	TransitionStorageClassGlacier            TransitionStorageClass = "GLACIER"
	TransitionStorageClassStandardIa         TransitionStorageClass = "STANDARD_IA"
	TransitionStorageClassOnezoneIa          TransitionStorageClass = "ONEZONE_IA"
	TransitionStorageClassIntelligentTiering TransitionStorageClass = "INTELLIGENT_TIERING"
	TransitionStorageClassDeepArchive        TransitionStorageClass = "DEEP_ARCHIVE"
)

// Values returns all known values for TransitionStorageClass. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TransitionStorageClass) Values() []TransitionStorageClass {
	return []TransitionStorageClass{
		"GLACIER",
		"STANDARD_IA",
		"ONEZONE_IA",
		"INTELLIGENT_TIERING",
		"DEEP_ARCHIVE",
	}
}