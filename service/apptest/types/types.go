// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Defines a batch.
type Batch struct {

	// The job name of the batch.
	//
	// This member is required.
	BatchJobName *string

	// The batch job parameters of the batch.
	BatchJobParameters map[string]string

	// The export data set names of the batch.
	ExportDataSetNames []string

	noSmithyDocumentSerde
}

// Defines a batch step input.
type BatchStepInput struct {

	// The batch job name of the batch step input.
	//
	// This member is required.
	BatchJobName *string

	// The resource of the batch step input.
	//
	// This member is required.
	Resource MainframeResourceSummary

	// The batch job parameters of the batch step input.
	BatchJobParameters map[string]string

	// The export data set names of the batch step input.
	ExportDataSetNames []string

	// The properties of the batch step input.
	Properties *MainframeActionProperties

	noSmithyDocumentSerde
}

// Defines a batch step output.
type BatchStepOutput struct {

	// The data set details of the batch step output.
	DataSetDetails []DataSet

	// The data set export location of the batch step output.
	DataSetExportLocation *string

	// The Database Migration Service (DMS) output location of the batch step output.
	DmsOutputLocation *string

	noSmithyDocumentSerde
}

// Summarizes a batch job.
type BatchSummary struct {

	// The step input of the batch summary.
	//
	// This member is required.
	StepInput *BatchStepInput

	// The step output of the batch summary.
	StepOutput *BatchStepOutput

	noSmithyDocumentSerde
}

// Specifies the CloudFormation template and its parameters.
type CloudFormation struct {

	// The template location of the CloudFormation template.
	//
	// This member is required.
	TemplateLocation *string

	// The CloudFormation properties in the CloudFormation template.
	Parameters map[string]string

	noSmithyDocumentSerde
}

// Specifies the CloudFormation action.
type CloudFormationAction struct {

	// The resource of the CloudFormation action.
	//
	// This member is required.
	Resource *string

	// The action type of the CloudFormation action.
	ActionType CloudFormationActionType

	noSmithyDocumentSerde
}

// Specifies the CloudFormation step summary.
//
// The following types satisfy this interface:
//
//	CloudFormationStepSummaryMemberCreateCloudformation
//	CloudFormationStepSummaryMemberDeleteCloudformation
type CloudFormationStepSummary interface {
	isCloudFormationStepSummary()
}

// Creates the CloudFormation summary of the step.
type CloudFormationStepSummaryMemberCreateCloudformation struct {
	Value CreateCloudFormationSummary

	noSmithyDocumentSerde
}

func (*CloudFormationStepSummaryMemberCreateCloudformation) isCloudFormationStepSummary() {}

// Deletes the CloudFormation summary of the CloudFormation step summary.
type CloudFormationStepSummaryMemberDeleteCloudformation struct {
	Value DeleteCloudFormationSummary

	noSmithyDocumentSerde
}

func (*CloudFormationStepSummaryMemberDeleteCloudformation) isCloudFormationStepSummary() {}

// Compares the action.
type CompareAction struct {

	// The input of the compare action.
	//
	// This member is required.
	Input Input

	// The output of the compare action.
	Output Output

	noSmithyDocumentSerde
}

// Specifies the compare action summary.
type CompareActionSummary struct {

	// The type of the compare action summary.
	//
	// This member is required.
	Type File

	noSmithyDocumentSerde
}

// Compares the database Change Data Capture (CDC) step input.
type CompareDatabaseCDCStepInput struct {

	// The source location of the compare database CDC step input.
	//
	// This member is required.
	SourceLocation *string

	// The source metadata of the compare database CDC step input.
	//
	// This member is required.
	SourceMetadata *SourceDatabaseMetadata

	// The target location of the compare database CDC step input.
	//
	// This member is required.
	TargetLocation *string

	// The target metadata location of the compare database CDC step input.
	//
	// This member is required.
	TargetMetadata *TargetDatabaseMetadata

	// The output location of the compare database CDC step input.
	OutputLocation *string

	noSmithyDocumentSerde
}

// Compares the database CDC step output.
type CompareDatabaseCDCStepOutput struct {

	// The comparison output of the compare database CDC step output.
	//
	// This member is required.
	ComparisonOutputLocation *string

	// The comparison status of the compare database CDC step output.
	//
	// This member is required.
	ComparisonStatus ComparisonStatusEnum

	noSmithyDocumentSerde
}

// Compares the database CDC summary.
type CompareDatabaseCDCSummary struct {

	// The step input of the compare database CDC summary.
	//
	// This member is required.
	StepInput *CompareDatabaseCDCStepInput

	// The step output of the compare database CDC summary.
	StepOutput *CompareDatabaseCDCStepOutput

	noSmithyDocumentSerde
}

// Specifies the compare data sets step input.
type CompareDataSetsStepInput struct {

	// The source data sets of the compare data sets step input location.
	//
	// This member is required.
	SourceDataSets []DataSet

	// The source location of the compare data sets step input location.
	//
	// This member is required.
	SourceLocation *string

	// The target data sets of the compare data sets step input location.
	//
	// This member is required.
	TargetDataSets []DataSet

	// The target location of the compare data sets step input location.
	//
	// This member is required.
	TargetLocation *string

	noSmithyDocumentSerde
}

// Specifies the compare data sets step output.
type CompareDataSetsStepOutput struct {

	// The comparison output location of the compare data sets step output.
	//
	// This member is required.
	ComparisonOutputLocation *string

	// The comparison status of the compare data sets step output.
	//
	// This member is required.
	ComparisonStatus ComparisonStatusEnum

	noSmithyDocumentSerde
}

// Compares data sets summary.
type CompareDataSetsSummary struct {

	// The step input of the compare data sets summary.
	//
	// This member is required.
	StepInput *CompareDataSetsStepInput

	// The step output of the compare data sets summary.
	StepOutput *CompareDataSetsStepOutput

	noSmithyDocumentSerde
}

// Compares the file type.
//
// The following types satisfy this interface:
//
//	CompareFileTypeMemberDatabaseCDC
//	CompareFileTypeMemberDatasets
type CompareFileType interface {
	isCompareFileType()
}

// The database CDC of the compare file type.
type CompareFileTypeMemberDatabaseCDC struct {
	Value CompareDatabaseCDCSummary

	noSmithyDocumentSerde
}

func (*CompareFileTypeMemberDatabaseCDC) isCompareFileType() {}

// The data sets in the compare file type.
type CompareFileTypeMemberDatasets struct {
	Value CompareDataSetsSummary

	noSmithyDocumentSerde
}

func (*CompareFileTypeMemberDatasets) isCompareFileType() {}

// Creates the CloudFormation step input.
type CreateCloudFormationStepInput struct {

	// The template location of the CloudFormation step input.
	//
	// This member is required.
	TemplateLocation *string

	// The CloudFormation properties of the CloudFormation step input.
	Parameters map[string]string

	noSmithyDocumentSerde
}

// Creates a CloudFormation step output.
type CreateCloudFormationStepOutput struct {

	// The stack ID of the CloudFormation step output.
	//
	// This member is required.
	StackId *string

	// The exports of the CloudFormation step output.
	Exports map[string]string

	noSmithyDocumentSerde
}

// Creates a CloudFormation summary.
type CreateCloudFormationSummary struct {

	// The step input of the CloudFormation summary.
	//
	// This member is required.
	StepInput *CreateCloudFormationStepInput

	// The step output of the CloudFormation summary.
	StepOutput *CreateCloudFormationStepOutput

	noSmithyDocumentSerde
}

// Defines the Change Data Capture (CDC) of the database.
type DatabaseCDC struct {

	// The source metadata of the database CDC.
	//
	// This member is required.
	SourceMetadata *SourceDatabaseMetadata

	// The target metadata of the database CDC.
	//
	// This member is required.
	TargetMetadata *TargetDatabaseMetadata

	noSmithyDocumentSerde
}

// Defines a data set.
type DataSet struct {

	// The CCSID of the data set.
	//
	// This member is required.
	Ccsid *string

	// The format of the data set.
	//
	// This member is required.
	Format Format

	// The length of the data set.
	//
	// This member is required.
	Length *int32

	// The name of the data set.
	//
	// This member is required.
	Name *string

	// The type of the data set.
	//
	// This member is required.
	Type DataSetType

	noSmithyDocumentSerde
}

// Deletes the CloudFormation step input.
type DeleteCloudFormationStepInput struct {

	// The stack ID of the deleted CloudFormation step input.
	//
	// This member is required.
	StackId *string

	noSmithyDocumentSerde
}

// Deletes the CloudFormation summary step output.
type DeleteCloudFormationStepOutput struct {
	noSmithyDocumentSerde
}

// Deletes the CloudFormation summary.
type DeleteCloudFormationSummary struct {

	// The step input of the deleted CloudFormation summary.
	//
	// This member is required.
	StepInput *DeleteCloudFormationStepInput

	// The step output of the deleted CloudFormation summary.
	StepOutput *DeleteCloudFormationStepOutput

	noSmithyDocumentSerde
}

// Defines a file.
//
// The following types satisfy this interface:
//
//	FileMemberFileType
type File interface {
	isFile()
}

// The file type of the file.
type FileMemberFileType struct {
	Value CompareFileType

	noSmithyDocumentSerde
}

func (*FileMemberFileType) isFile() {}

// Specifies a file metadata.
//
// The following types satisfy this interface:
//
//	FileMetadataMemberDatabaseCDC
//	FileMetadataMemberDataSets
type FileMetadata interface {
	isFileMetadata()
}

// The database CDC of the file metadata.
type FileMetadataMemberDatabaseCDC struct {
	Value DatabaseCDC

	noSmithyDocumentSerde
}

func (*FileMetadataMemberDatabaseCDC) isFileMetadata() {}

// The data sets of the file metadata.
type FileMetadataMemberDataSets struct {
	Value []DataSet

	noSmithyDocumentSerde
}

func (*FileMetadataMemberDataSets) isFileMetadata() {}

// Specifies the input.
//
// The following types satisfy this interface:
//
//	InputMemberFile
type Input interface {
	isInput()
}

// The file in the input.
type InputMemberFile struct {
	Value InputFile

	noSmithyDocumentSerde
}

func (*InputMemberFile) isInput() {}

// Specifies the input file.
type InputFile struct {

	// The file metadata of the input file.
	//
	// This member is required.
	FileMetadata FileMetadata

	// The source location of the input file.
	//
	// This member is required.
	SourceLocation *string

	// The target location of the input file.
	//
	// This member is required.
	TargetLocation *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed action properties.
type M2ManagedActionProperties struct {

	// Force stops the AWS Mainframe Modernization managed action properties.
	ForceStop *bool

	// The import data set location of the AWS Mainframe Modernization managed action
	// properties.
	ImportDataSetLocation *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed application.
type M2ManagedApplication struct {

	// The application ID of the AWS Mainframe Modernization managed application.
	//
	// This member is required.
	ApplicationId *string

	// The runtime of the AWS Mainframe Modernization managed application.
	//
	// This member is required.
	Runtime M2ManagedRuntime

	// The listener port of the AWS Mainframe Modernization managed application.
	ListenerPort *string

	// The VPC endpoint service name of the AWS Mainframe Modernization managed
	// application.
	VpcEndpointServiceName *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed application action.
type M2ManagedApplicationAction struct {

	// The action type of the AWS Mainframe Modernization managed application action.
	//
	// This member is required.
	ActionType M2ManagedActionType

	// The resource of the AWS Mainframe Modernization managed application action.
	//
	// This member is required.
	Resource *string

	// The properties of the AWS Mainframe Modernization managed application action.
	Properties *M2ManagedActionProperties

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed application step input.
type M2ManagedApplicationStepInput struct {

	// The action type of the AWS Mainframe Modernization managed application step
	// input.
	//
	// This member is required.
	ActionType M2ManagedActionType

	// The application ID of the AWS Mainframe Modernization managed application step
	// input.
	//
	// This member is required.
	ApplicationId *string

	// The runtime of the AWS Mainframe Modernization managed application step input.
	//
	// This member is required.
	Runtime *string

	// The listener port of the AWS Mainframe Modernization managed application step
	// input.
	ListenerPort *int32

	// The properties of the AWS Mainframe Modernization managed application step
	// input.
	Properties *M2ManagedActionProperties

	// The VPC endpoint service name of the AWS Mainframe Modernization managed
	// application step input.
	VpcEndpointServiceName *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed application step output.
type M2ManagedApplicationStepOutput struct {

	// The import data set summary of the AWS Mainframe Modernization managed
	// application step output.
	ImportDataSetSummary map[string]string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed application step summary.
type M2ManagedApplicationStepSummary struct {

	// The step input of the AWS Mainframe Modernization managed application step
	// summary.
	//
	// This member is required.
	StepInput *M2ManagedApplicationStepInput

	// The step output of the AWS Mainframe Modernization managed application step
	// summary.
	StepOutput *M2ManagedApplicationStepOutput

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization managed application summary.
type M2ManagedApplicationSummary struct {

	// The application ID of the AWS Mainframe Modernization managed application
	// summary.
	//
	// This member is required.
	ApplicationId *string

	// The runtime of the AWS Mainframe Modernization managed application summary.
	//
	// This member is required.
	Runtime M2ManagedRuntime

	// The listener port of the AWS Mainframe Modernization managed application
	// summary.
	ListenerPort *int32

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization non-managed application.
type M2NonManagedApplication struct {

	// The listener port of the AWS Mainframe Modernization non-managed application.
	//
	// This member is required.
	ListenerPort *string

	// The runtime of the AWS Mainframe Modernization non-managed application.
	//
	// This member is required.
	Runtime M2NonManagedRuntime

	// The VPC endpoint service name of the AWS Mainframe Modernization non-managed
	// application.
	//
	// This member is required.
	VpcEndpointServiceName *string

	// The web application name of the AWS Mainframe Modernization non-managed
	// application.
	WebAppName *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization non-managed application action.
type M2NonManagedApplicationAction struct {

	// The action type of the AWS Mainframe Modernization non-managed application
	// action.
	//
	// This member is required.
	ActionType M2NonManagedActionType

	// The resource of the AWS Mainframe Modernization non-managed application action.
	//
	// This member is required.
	Resource *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization non-managed application step input.
type M2NonManagedApplicationStepInput struct {

	// The action type of the AWS Mainframe Modernization non-managed application step
	// input.
	//
	// This member is required.
	ActionType M2NonManagedActionType

	// The listener port of the AWS Mainframe Modernization non-managed application
	// step input.
	//
	// This member is required.
	ListenerPort *int32

	// The runtime of the AWS Mainframe Modernization non-managed application step
	// input.
	//
	// This member is required.
	Runtime M2NonManagedRuntime

	// The VPC endpoint service name of the AWS Mainframe Modernization non-managed
	// application step input.
	//
	// This member is required.
	VpcEndpointServiceName *string

	// The web app name of the AWS Mainframe Modernization non-managed application
	// step input.
	WebAppName *string

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization non-managed application step output.
type M2NonManagedApplicationStepOutput struct {
	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization non-managed application step summary.
type M2NonManagedApplicationStepSummary struct {

	// The step input of the AWS Mainframe Modernization non-managed application step
	// summary.
	//
	// This member is required.
	StepInput *M2NonManagedApplicationStepInput

	// The step output of the AWS Mainframe Modernization non-managed application step
	// summary.
	StepOutput *M2NonManagedApplicationStepOutput

	noSmithyDocumentSerde
}

// Specifies the AWS Mainframe Modernization non-managed application summary.
type M2NonManagedApplicationSummary struct {

	// The listener port of the AWS Mainframe Modernization non-managed application
	// summary.
	//
	// This member is required.
	ListenerPort *int32

	// The runtime of the AWS Mainframe Modernization non-managed application summary.
	//
	// This member is required.
	Runtime M2NonManagedRuntime

	// The VPC endpoint service name of the AWS Mainframe Modernization non-managed
	// application summary.
	//
	// This member is required.
	VpcEndpointServiceName *string

	// The web application name of the AWS Mainframe Modernization non-managed
	// application summary.
	WebAppName *string

	noSmithyDocumentSerde
}

// Specifies the mainframe action.
type MainframeAction struct {

	// The action type of the mainframe action.
	//
	// This member is required.
	ActionType MainframeActionType

	// The resource of the mainframe action.
	//
	// This member is required.
	Resource *string

	// The properties of the mainframe action.
	Properties *MainframeActionProperties

	noSmithyDocumentSerde
}

// Specifies the mainframe action properties.
type MainframeActionProperties struct {

	// The DMS task ARN of the mainframe action properties.
	DmsTaskArn *string

	noSmithyDocumentSerde
}

// Specifies the mainframe action summary.
//
// The following types satisfy this interface:
//
//	MainframeActionSummaryMemberBatch
//	MainframeActionSummaryMemberTn3270
type MainframeActionSummary interface {
	isMainframeActionSummary()
}

// The batch of the mainframe action summary.
type MainframeActionSummaryMemberBatch struct {
	Value BatchSummary

	noSmithyDocumentSerde
}

func (*MainframeActionSummaryMemberBatch) isMainframeActionSummary() {}

// The tn3270 port of the mainframe action summary.
type MainframeActionSummaryMemberTn3270 struct {
	Value TN3270Summary

	noSmithyDocumentSerde
}

func (*MainframeActionSummaryMemberTn3270) isMainframeActionSummary() {}

// Specifies the mainframe action type.
//
// The following types satisfy this interface:
//
//	MainframeActionTypeMemberBatch
//	MainframeActionTypeMemberTn3270
type MainframeActionType interface {
	isMainframeActionType()
}

// The batch of the mainframe action type.
type MainframeActionTypeMemberBatch struct {
	Value Batch

	noSmithyDocumentSerde
}

func (*MainframeActionTypeMemberBatch) isMainframeActionType() {}

// The tn3270 port of the mainframe action type.
type MainframeActionTypeMemberTn3270 struct {
	Value TN3270

	noSmithyDocumentSerde
}

func (*MainframeActionTypeMemberTn3270) isMainframeActionType() {}

// Specifies the mainframe resource summary.
//
// The following types satisfy this interface:
//
//	MainframeResourceSummaryMemberM2ManagedApplication
//	MainframeResourceSummaryMemberM2NonManagedApplication
type MainframeResourceSummary interface {
	isMainframeResourceSummary()
}

// The AWS Mainframe Modernization managed application in the mainframe resource
// summary.
type MainframeResourceSummaryMemberM2ManagedApplication struct {
	Value M2ManagedApplicationSummary

	noSmithyDocumentSerde
}

func (*MainframeResourceSummaryMemberM2ManagedApplication) isMainframeResourceSummary() {}

// The AWS Mainframe Modernization non-managed application in the mainframe
// resource summary.
type MainframeResourceSummaryMemberM2NonManagedApplication struct {
	Value M2NonManagedApplicationSummary

	noSmithyDocumentSerde
}

func (*MainframeResourceSummaryMemberM2NonManagedApplication) isMainframeResourceSummary() {}

// Specifies an output.
//
// The following types satisfy this interface:
//
//	OutputMemberFile
type Output interface {
	isOutput()
}

// The file of the output.
type OutputMemberFile struct {
	Value OutputFile

	noSmithyDocumentSerde
}

func (*OutputMemberFile) isOutput() {}

// Specifies an output file.
type OutputFile struct {

	// The file location of the output file.
	FileLocation *string

	noSmithyDocumentSerde
}

// Specifies a resource.
type Resource struct {

	// The name of the resource.
	//
	// This member is required.
	Name *string

	// The type of the resource.
	//
	// This member is required.
	Type ResourceType

	noSmithyDocumentSerde
}

// Specifies a resource action.
//
// The following types satisfy this interface:
//
//	ResourceActionMemberCloudFormationAction
//	ResourceActionMemberM2ManagedApplicationAction
//	ResourceActionMemberM2NonManagedApplicationAction
type ResourceAction interface {
	isResourceAction()
}

// The CloudFormation action of the resource action.
type ResourceActionMemberCloudFormationAction struct {
	Value CloudFormationAction

	noSmithyDocumentSerde
}

func (*ResourceActionMemberCloudFormationAction) isResourceAction() {}

// The AWS Mainframe Modernization managed application action of the resource
// action.
type ResourceActionMemberM2ManagedApplicationAction struct {
	Value M2ManagedApplicationAction

	noSmithyDocumentSerde
}

func (*ResourceActionMemberM2ManagedApplicationAction) isResourceAction() {}

// The AWS Mainframe Modernization non-managed application action of the resource
// action.
type ResourceActionMemberM2NonManagedApplicationAction struct {
	Value M2NonManagedApplicationAction

	noSmithyDocumentSerde
}

func (*ResourceActionMemberM2NonManagedApplicationAction) isResourceAction() {}

// Specifies the resource action summary.
//
// The following types satisfy this interface:
//
//	ResourceActionSummaryMemberCloudFormation
//	ResourceActionSummaryMemberM2ManagedApplication
//	ResourceActionSummaryMemberM2NonManagedApplication
type ResourceActionSummary interface {
	isResourceActionSummary()
}

// The CloudFormation template of the resource action summary.
type ResourceActionSummaryMemberCloudFormation struct {
	Value CloudFormationStepSummary

	noSmithyDocumentSerde
}

func (*ResourceActionSummaryMemberCloudFormation) isResourceActionSummary() {}

// The AWS Mainframe Modernization managed application of the resource action
// summary.
type ResourceActionSummaryMemberM2ManagedApplication struct {
	Value M2ManagedApplicationStepSummary

	noSmithyDocumentSerde
}

func (*ResourceActionSummaryMemberM2ManagedApplication) isResourceActionSummary() {}

// The AWS Mainframe Modernization non-managed application of the resource action
// summary.
type ResourceActionSummaryMemberM2NonManagedApplication struct {
	Value M2NonManagedApplicationStepSummary

	noSmithyDocumentSerde
}

func (*ResourceActionSummaryMemberM2NonManagedApplication) isResourceActionSummary() {}

// Specifies the resource type.
//
// The following types satisfy this interface:
//
//	ResourceTypeMemberCloudFormation
//	ResourceTypeMemberM2ManagedApplication
//	ResourceTypeMemberM2NonManagedApplication
type ResourceType interface {
	isResourceType()
}

// The CloudFormation template of the resource type.
type ResourceTypeMemberCloudFormation struct {
	Value CloudFormation

	noSmithyDocumentSerde
}

func (*ResourceTypeMemberCloudFormation) isResourceType() {}

// The AWS Mainframe Modernization managed application of the resource type.
type ResourceTypeMemberM2ManagedApplication struct {
	Value M2ManagedApplication

	noSmithyDocumentSerde
}

func (*ResourceTypeMemberM2ManagedApplication) isResourceType() {}

// The AWS Mainframe Modernization non-managed application of the resource type.
type ResourceTypeMemberM2NonManagedApplication struct {
	Value M2NonManagedApplication

	noSmithyDocumentSerde
}

func (*ResourceTypeMemberM2NonManagedApplication) isResourceType() {}

// Specifies the script.
type Script struct {

	// The script location of the scripts.
	//
	// This member is required.
	ScriptLocation *string

	// The type of the scripts.
	//
	// This member is required.
	Type ScriptType

	noSmithyDocumentSerde
}

// Specifies the scripts summary.
type ScriptSummary struct {

	// The script location of the script summary.
	//
	// This member is required.
	ScriptLocation *string

	// The type of the script summary.
	//
	// This member is required.
	Type ScriptType

	noSmithyDocumentSerde
}

// Specifies the service settings.
type ServiceSettings struct {

	// The KMS key ID of the service settings.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// Specifies the source database metadata.
type SourceDatabaseMetadata struct {

	// The capture tool of the source database metadata.
	//
	// This member is required.
	CaptureTool CaptureTool

	// The type of the source database metadata.
	//
	// This member is required.
	Type SourceDatabase

	noSmithyDocumentSerde
}

// Defines a step.
type Step struct {

	// The action of the step.
	//
	// This member is required.
	Action StepAction

	// The name of the step.
	//
	// This member is required.
	Name *string

	// The description of the step.
	Description *string

	noSmithyDocumentSerde
}

// Specifies a step action.
//
// The following types satisfy this interface:
//
//	StepActionMemberCompareAction
//	StepActionMemberMainframeAction
//	StepActionMemberResourceAction
type StepAction interface {
	isStepAction()
}

// The compare action of the step action.
type StepActionMemberCompareAction struct {
	Value CompareAction

	noSmithyDocumentSerde
}

func (*StepActionMemberCompareAction) isStepAction() {}

// The mainframe action of the step action.
type StepActionMemberMainframeAction struct {
	Value MainframeAction

	noSmithyDocumentSerde
}

func (*StepActionMemberMainframeAction) isStepAction() {}

// The resource action of the step action.
type StepActionMemberResourceAction struct {
	Value ResourceAction

	noSmithyDocumentSerde
}

func (*StepActionMemberResourceAction) isStepAction() {}

// Defines the step run summary.
//
// The following types satisfy this interface:
//
//	StepRunSummaryMemberCompareAction
//	StepRunSummaryMemberMainframeAction
//	StepRunSummaryMemberResourceAction
type StepRunSummary interface {
	isStepRunSummary()
}

// The compare action of the step run summary.
type StepRunSummaryMemberCompareAction struct {
	Value CompareActionSummary

	noSmithyDocumentSerde
}

func (*StepRunSummaryMemberCompareAction) isStepRunSummary() {}

// The mainframe action of the step run summary.
type StepRunSummaryMemberMainframeAction struct {
	Value MainframeActionSummary

	noSmithyDocumentSerde
}

func (*StepRunSummaryMemberMainframeAction) isStepRunSummary() {}

// The resource action of the step run summary.
type StepRunSummaryMemberResourceAction struct {
	Value ResourceActionSummary

	noSmithyDocumentSerde
}

func (*StepRunSummaryMemberResourceAction) isStepRunSummary() {}

// Specifies a target database metadata.
type TargetDatabaseMetadata struct {

	// The capture tool of the target database metadata.
	//
	// This member is required.
	CaptureTool CaptureTool

	// The type of the target database metadata.
	//
	// This member is required.
	Type TargetDatabase

	noSmithyDocumentSerde
}

// Specifies the latest version of a test case.
type TestCaseLatestVersion struct {

	// The status of the test case latest version.
	//
	// This member is required.
	Status TestCaseLifecycle

	// The version of the test case latest version.
	//
	// This member is required.
	Version *int32

	// The status reason of the test case latest version.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies the test case run summary.
type TestCaseRunSummary struct {

	// The run start time of the test case run summary.
	//
	// This member is required.
	RunStartTime *time.Time

	// The status of the test case run summary.
	//
	// This member is required.
	Status TestCaseRunStatus

	// The test case id of the test case run summary.
	//
	// This member is required.
	TestCaseId *string

	// The test case version of the test case run summary.
	//
	// This member is required.
	TestCaseVersion *int32

	// The test run id of the test case run summary.
	//
	// This member is required.
	TestRunId *string

	// The run end time of the test case run summary.
	RunEndTime *time.Time

	// The status reason of the test case run summary.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies test cases.
//
// The following types satisfy this interface:
//
//	TestCasesMemberSequential
type TestCases interface {
	isTestCases()
}

// The sequential of the test case.
type TestCasesMemberSequential struct {
	Value []string

	noSmithyDocumentSerde
}

func (*TestCasesMemberSequential) isTestCases() {}

// Specifies a test case summary.
type TestCaseSummary struct {

	// The creation time of the test case summary.
	//
	// This member is required.
	CreationTime *time.Time

	// The last update time of the test case summary.
	//
	// This member is required.
	LastUpdateTime *time.Time

	// The latest version of the test case summary.
	//
	// This member is required.
	LatestVersion *int32

	// The name of the test case summary.
	//
	// This member is required.
	Name *string

	// The status of the test case summary.
	//
	// This member is required.
	Status TestCaseLifecycle

	// The test case Amazon Resource Name (ARN) of the test case summary.
	//
	// This member is required.
	TestCaseArn *string

	// The test case ID of the test case summary.
	//
	// This member is required.
	TestCaseId *string

	// The status reason of the test case summary.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies the latest version of the test configuration.
type TestConfigurationLatestVersion struct {

	// The status of the test configuration latest version.
	//
	// This member is required.
	Status TestConfigurationLifecycle

	// The version of the test configuration latest version.
	//
	// This member is required.
	Version *int32

	// The status reason of the test configuration latest version.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies a test configuration summary.
type TestConfigurationSummary struct {

	// The creation time of the test configuration summary.
	//
	// This member is required.
	CreationTime *time.Time

	// The last update time of the test configuration summary.
	//
	// This member is required.
	LastUpdateTime *time.Time

	// The latest version of the test configuration summary.
	//
	// This member is required.
	LatestVersion *int32

	// The name of the test configuration summary.
	//
	// This member is required.
	Name *string

	// The status of the test configuration summary.
	//
	// This member is required.
	Status TestConfigurationLifecycle

	// The test configuration ARN of the test configuration summary.
	//
	// This member is required.
	TestConfigurationArn *string

	// The test configuration ID of the test configuration summary.
	//
	// This member is required.
	TestConfigurationId *string

	// The status reason of the test configuration summary.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies a test run step summary.
type TestRunStepSummary struct {

	// The run start time of the test run step summary.
	//
	// This member is required.
	RunStartTime *time.Time

	// The status of the test run step summary.
	//
	// This member is required.
	Status StepRunStatus

	// The step name of the test run step summary.
	//
	// This member is required.
	StepName *string

	// The test run ID of the test run step summary.
	//
	// This member is required.
	TestRunId *string

	// The after step of the test run step summary.
	AfterStep *bool

	// The before step of the test run step summary.
	BeforeStep *bool

	// The run end time of the test run step summary.
	RunEndTime *time.Time

	// The status reason of the test run step summary.
	StatusReason *string

	// The test case ID of the test run step summary.
	TestCaseId *string

	// The test case version of the test run step summary.
	TestCaseVersion *int32

	// The test suite ID of the test run step summary.
	TestSuiteId *string

	// The test suite version of the test run step summary.
	TestSuiteVersion *int32

	noSmithyDocumentSerde
}

// Specifies a test run summary.
type TestRunSummary struct {

	// The run start time of the test run summary.
	//
	// This member is required.
	RunStartTime *time.Time

	// The status of the test run summary.
	//
	// This member is required.
	Status TestRunStatus

	// The test run ARN of the test run summary.
	//
	// This member is required.
	TestRunArn *string

	// The test run ID of the test run summary.
	//
	// This member is required.
	TestRunId *string

	// The test suite ID of the test run summary.
	//
	// This member is required.
	TestSuiteId *string

	// The test suite version of the test run summary.
	//
	// This member is required.
	TestSuiteVersion *int32

	// The run end time of the test run summary.
	RunEndTime *time.Time

	// The status reason of the test run summary.
	StatusReason *string

	// The test configuration ID of the test run summary.
	TestConfigurationId *string

	// The test configuration version of the test run summary.
	TestConfigurationVersion *int32

	noSmithyDocumentSerde
}

// Specifies the latest version of a test suite.
type TestSuiteLatestVersion struct {

	// The status of the test suite latest version.
	//
	// This member is required.
	Status TestSuiteLifecycle

	// The version of the test suite latest version.
	//
	// This member is required.
	Version *int32

	// The status reason of the test suite latest version.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies the test suite summary.
type TestSuiteSummary struct {

	// The creation time of the test suite summary.
	//
	// This member is required.
	CreationTime *time.Time

	// The last update time of the test suite summary.
	//
	// This member is required.
	LastUpdateTime *time.Time

	// The latest version of the test suite summary.
	//
	// This member is required.
	LatestVersion *int32

	// The name of the test suite summary.
	//
	// This member is required.
	Name *string

	// The status of the test suite summary.
	//
	// This member is required.
	Status TestSuiteLifecycle

	// The test suite Amazon Resource Name (ARN) of the test suite summary.
	//
	// This member is required.
	TestSuiteArn *string

	// The test suite ID of the test suite summary.
	//
	// This member is required.
	TestSuiteId *string

	// The status reason of the test suite summary.
	StatusReason *string

	noSmithyDocumentSerde
}

// Specifies the TN3270 protocol.
type TN3270 struct {

	// The script of the TN3270 protocol.
	//
	// This member is required.
	Script *Script

	// The data set names of the TN3270 protocol.
	ExportDataSetNames []string

	noSmithyDocumentSerde
}

// Specifies a TN3270 step input.
type TN3270StepInput struct {

	// The resource of the TN3270 step input.
	//
	// This member is required.
	Resource MainframeResourceSummary

	// The script of the TN3270 step input.
	//
	// This member is required.
	Script *ScriptSummary

	// The export data set names of the TN3270 step input.
	ExportDataSetNames []string

	// The properties of the TN3270 step input.
	Properties *MainframeActionProperties

	noSmithyDocumentSerde
}

// Specifies a TN3270 step output.
type TN3270StepOutput struct {

	// The script output location of the TN3270 step output.
	//
	// This member is required.
	ScriptOutputLocation *string

	// The data set details of the TN3270 step output.
	DataSetDetails []DataSet

	// The data set export location of the TN3270 step output.
	DataSetExportLocation *string

	// The output location of the TN3270 step output.
	DmsOutputLocation *string

	noSmithyDocumentSerde
}

// Specifies a TN3270 summary.
type TN3270Summary struct {

	// The step input of the TN3270 summary.
	//
	// This member is required.
	StepInput *TN3270StepInput

	// The step output of the TN3270 summary.
	StepOutput *TN3270StepOutput

	noSmithyDocumentSerde
}

// Specifies a validation exception field.
type ValidationExceptionField struct {

	// The message stating reason for why service validation failed.
	//
	// This member is required.
	Message *string

	// The name of the validation exception field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isCloudFormationStepSummary() {}
func (*UnknownUnionMember) isCompareFileType()           {}
func (*UnknownUnionMember) isFile()                      {}
func (*UnknownUnionMember) isFileMetadata()              {}
func (*UnknownUnionMember) isInput()                     {}
func (*UnknownUnionMember) isMainframeActionSummary()    {}
func (*UnknownUnionMember) isMainframeActionType()       {}
func (*UnknownUnionMember) isMainframeResourceSummary()  {}
func (*UnknownUnionMember) isOutput()                    {}
func (*UnknownUnionMember) isResourceAction()            {}
func (*UnknownUnionMember) isResourceActionSummary()     {}
func (*UnknownUnionMember) isResourceType()              {}
func (*UnknownUnionMember) isStepAction()                {}
func (*UnknownUnionMember) isStepRunSummary()            {}
func (*UnknownUnionMember) isTestCases()                 {}