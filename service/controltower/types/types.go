// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/service/controltower/document"
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An object of shape BaselineOperation , returning details about the specified
// Baseline operation ID.
type BaselineOperation struct {

	// The end time of the operation (if applicable), in ISO 8601 format.
	EndTime *time.Time

	// The identifier of the specified operation.
	OperationIdentifier *string

	// An enumerated type ( enum ) with possible values of ENABLE_BASELINE ,
	// DISABLE_BASELINE , UPDATE_ENABLED_BASELINE , or RESET_ENABLED_BASELINE .
	OperationType BaselineOperationType

	// The start time of the operation, in ISO 8601 format.
	StartTime *time.Time

	// An enumerated type ( enum ) with possible values of SUCCEEDED , FAILED , or
	// IN_PROGRESS .
	Status BaselineOperationStatus

	// A status message that gives more information about the operation's status, if
	// applicable.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Returns a summary of information about a Baseline object.
type BaselineSummary struct {

	// The full ARN of a Baseline.
	//
	// This member is required.
	Arn *string

	// The human-readable name of a Baseline.
	//
	// This member is required.
	Name *string

	// A summary description of a Baseline.
	Description *string

	noSmithyDocumentSerde
}

// An operation performed by the control.
type ControlOperation struct {

	// The controlIdentifier of the control for the operation.
	ControlIdentifier *string

	// The controlIdentifier of the enabled control.
	EnabledControlIdentifier *string

	// The time that the operation finished.
	EndTime *time.Time

	// The identifier of the specified operation.
	OperationIdentifier *string

	// One of ENABLE_CONTROL or DISABLE_CONTROL .
	OperationType ControlOperationType

	// The time that the operation began.
	StartTime *time.Time

	// One of IN_PROGRESS , SUCEEDED , or FAILED .
	Status ControlOperationStatus

	// If the operation result is FAILED , this string contains a message explaining
	// why the operation failed.
	StatusMessage *string

	// The target upon which the control operation is working.
	TargetIdentifier *string

	noSmithyDocumentSerde
}

// A filter object that lets you call ListControlOperations with a specific filter.
type ControlOperationFilter struct {

	// The set of controlIdentifier returned by the filter.
	ControlIdentifiers []string

	// The set of ControlOperation objects returned by the filter.
	ControlOperationTypes []ControlOperationType

	// The set controlIdentifier of enabled controls selected by the filter.
	EnabledControlIdentifiers []string

	// Lists the status of control operations.
	Statuses []ControlOperationStatus

	// The set of targetIdentifier objects returned by the filter.
	TargetIdentifiers []string

	noSmithyDocumentSerde
}

// A summary of information about the specified control operation.
type ControlOperationSummary struct {

	// The controlIdentifier of a control.
	ControlIdentifier *string

	// The controlIdentifier of an enabled control.
	EnabledControlIdentifier *string

	// The time at which the control operation was completed.
	EndTime *time.Time

	// The unique identifier of a control operation.
	OperationIdentifier *string

	// The type of operation.
	OperationType ControlOperationType

	// The time at which a control operation began.
	StartTime *time.Time

	// The status of the specified control operation.
	Status ControlOperationStatus

	// A speficic message displayed as part of the control status.
	StatusMessage *string

	// The unique identifier of the target of a control operation.
	TargetIdentifier *string

	noSmithyDocumentSerde
}

// The drift summary of the enabled control.
//
// Amazon Web Services Control Tower expects the enabled control configuration to
// include all supported and governed Regions. If the enabled control differs from
// the expected configuration, it is defined to be in a state of drift. You can
// repair this drift by resetting the enabled control.
type DriftStatusSummary struct {

	//  The drift status of the enabled control.
	//
	// Valid values:
	//
	//   - DRIFTED : The enabledControl deployed in this configuration doesn’t match
	//   the configuration that Amazon Web Services Control Tower expected.
	//
	//   - IN_SYNC : The enabledControl deployed in this configuration matches the
	//   configuration that Amazon Web Services Control Tower expected.
	//
	//   - NOT_CHECKING : Amazon Web Services Control Tower does not check drift for
	//   this enabled control. Drift is not supported for the control type.
	//
	//   - UNKNOWN : Amazon Web Services Control Tower is not able to check the drift
	//   status for the enabled control.
	DriftStatus DriftStatus

	noSmithyDocumentSerde
}

// Details of the EnabledBaseline resource.
type EnabledBaselineDetails struct {

	// The ARN of the EnabledBaseline resource.
	//
	// This member is required.
	Arn *string

	// The specific Baseline enabled as part of the EnabledBaseline resource.
	//
	// This member is required.
	BaselineIdentifier *string

	// The deployment summary of an EnabledControl or EnabledBaseline resource.
	//
	// This member is required.
	StatusSummary *EnablementStatusSummary

	// The target on which to enable the Baseline .
	//
	// This member is required.
	TargetIdentifier *string

	// The enabled version of the Baseline .
	BaselineVersion *string

	// Shows the parameters that are applied when enabling this Baseline .
	Parameters []EnabledBaselineParameterSummary

	noSmithyDocumentSerde
}

// A filter applied on the ListEnabledBaseline operation. Allowed filters are
// baselineIdentifiers and targetIdentifiers . The filter can be applied for
// either, or both.
type EnabledBaselineFilter struct {

	// Identifiers for the Baseline objects returned as part of the filter operation.
	BaselineIdentifiers []string

	// Identifiers for the targets of the Baseline filter operation.
	TargetIdentifiers []string

	noSmithyDocumentSerde
}

// A key-value parameter to an EnabledBaseline resource.
type EnabledBaselineParameter struct {

	// A string denoting the parameter key.
	//
	// This member is required.
	Key *string

	// A low-level Document object of any type (for example, a Java Object).
	//
	// This member is required.
	Value document.Interface

	noSmithyDocumentSerde
}

// Summary of an applied parameter to an EnabledBaseline resource.
type EnabledBaselineParameterSummary struct {

	// A string denoting the parameter key.
	//
	// This member is required.
	Key *string

	// A low-level document object of any type (for example, a Java Object).
	//
	// This member is required.
	Value document.Interface

	noSmithyDocumentSerde
}

// Returns a summary of information about an EnabledBaseline object.
type EnabledBaselineSummary struct {

	// The ARN of the EnabledBaseline resource
	//
	// This member is required.
	Arn *string

	// The specific baseline that is enabled as part of the EnabledBaseline resource.
	//
	// This member is required.
	BaselineIdentifier *string

	// The deployment summary of an EnabledControl or EnabledBaseline resource.
	//
	// This member is required.
	StatusSummary *EnablementStatusSummary

	// The target upon which the baseline is enabled.
	//
	// This member is required.
	TargetIdentifier *string

	// The enabled version of the baseline.
	BaselineVersion *string

	noSmithyDocumentSerde
}

// Information about the enabled control.
type EnabledControlDetails struct {

	// The ARN of the enabled control.
	Arn *string

	// The control identifier of the enabled control. For information on how to find
	// the controlIdentifier , see [the overview page].
	//
	// [the overview page]: https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html
	ControlIdentifier *string

	// The drift status of the enabled control.
	DriftStatusSummary *DriftStatusSummary

	// Array of EnabledControlParameter objects.
	Parameters []EnabledControlParameterSummary

	// The deployment summary of the enabled control.
	StatusSummary *EnablementStatusSummary

	// The ARN of the organizational unit. For information on how to find the
	// targetIdentifier , see [the overview page].
	//
	// [the overview page]: https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html
	TargetIdentifier *string

	// Target Amazon Web Services Regions for the enabled control.
	TargetRegions []Region

	noSmithyDocumentSerde
}

// A structure that returns a set of control identifiers, the control status for
// each control in the set, and the drift status for each control in the set.
type EnabledControlFilter struct {

	// The set of controlIdentifier returned by the filter.
	ControlIdentifiers []string

	// A list of DriftStatus items.
	DriftStatuses []DriftStatus

	// A list of EnablementStatus items.
	Statuses []EnablementStatus

	noSmithyDocumentSerde
}

// A key/value pair, where Key is of type String and Value is of type Document .
type EnabledControlParameter struct {

	// The key of a key/value pair.
	//
	// This member is required.
	Key *string

	// The value of a key/value pair.
	//
	// This member is required.
	Value document.Interface

	noSmithyDocumentSerde
}

// Returns a summary of information about the parameters of an enabled control.
type EnabledControlParameterSummary struct {

	// The key of a key/value pair.
	//
	// This member is required.
	Key *string

	// The value of a key/value pair.
	//
	// This member is required.
	Value document.Interface

	noSmithyDocumentSerde
}

// Returns a summary of information about an enabled control.
type EnabledControlSummary struct {

	// The ARN of the enabled control.
	Arn *string

	// The controlIdentifier of the enabled control.
	ControlIdentifier *string

	// The drift status of the enabled control.
	DriftStatusSummary *DriftStatusSummary

	// A short description of the status of the enabled control.
	StatusSummary *EnablementStatusSummary

	// The ARN of the organizational unit.
	TargetIdentifier *string

	noSmithyDocumentSerde
}

// The deployment summary of an EnabledControl or EnabledBaseline resource.
type EnablementStatusSummary struct {

	// The last operation identifier for the enabled resource.
	LastOperationIdentifier *string

	//  The deployment status of the enabled resource.
	//
	// Valid values:
	//
	//   - SUCCEEDED : The EnabledControl or EnabledBaseline configuration was deployed
	//   successfully.
	//
	//   - UNDER_CHANGE : The EnabledControl or EnabledBaseline configuration is
	//   changing.
	//
	//   - FAILED : The EnabledControl or EnabledBaseline configuration failed to
	//   deploy.
	Status EnablementStatus

	noSmithyDocumentSerde
}

// Information about the landing zone.
type LandingZoneDetail struct {

	// The landing zone manifest JSON text file that specifies the landing zone
	// configurations.
	//
	// This member is required.
	Manifest document.Interface

	// The landing zone's current deployed version.
	//
	// This member is required.
	Version *string

	// The ARN of the landing zone.
	Arn *string

	// The drift status of the landing zone.
	DriftStatus *LandingZoneDriftStatusSummary

	// The latest available version of the landing zone.
	LatestAvailableVersion *string

	// The landing zone deployment status. One of ACTIVE , PROCESSING , FAILED .
	Status LandingZoneStatus

	noSmithyDocumentSerde
}

// The drift status summary of the landing zone.
//
// If the landing zone differs from the expected configuration, it is defined to
// be in a state of drift. You can repair this drift by resetting the landing zone.
type LandingZoneDriftStatusSummary struct {

	// The drift status of the landing zone.
	//
	// Valid values:
	//
	//   - DRIFTED : The landing zone deployed in this configuration does not match the
	//   configuration that Amazon Web Services Control Tower expected.
	//
	//   - IN_SYNC : The landing zone deployed in this configuration matches the
	//   configuration that Amazon Web Services Control Tower expected.
	Status LandingZoneDriftStatus

	noSmithyDocumentSerde
}

// Information about a landing zone operation.
type LandingZoneOperationDetail struct {

	// The landing zone operation end time.
	EndTime *time.Time

	// The operationIdentifier of the landing zone operation.
	OperationIdentifier *string

	// The landing zone operation type.
	//
	// Valid values:
	//
	//   - DELETE : The DeleteLandingZone operation.
	//
	//   - CREATE : The CreateLandingZone operation.
	//
	//   - UPDATE : The UpdateLandingZone operation.
	//
	//   - RESET : The ResetLandingZone operation.
	OperationType LandingZoneOperationType

	// The landing zone operation start time.
	StartTime *time.Time

	// Valid values:
	//
	//   - SUCCEEDED : The landing zone operation succeeded.
	//
	//   - IN_PROGRESS : The landing zone operation is in progress.
	//
	//   - FAILED : The landing zone operation failed.
	Status LandingZoneOperationStatus

	// If the operation result is FAILED, this string contains a message explaining
	// why the operation failed.
	StatusMessage *string

	noSmithyDocumentSerde
}

// A filter object that lets you call ListLandingZoneOperations with a specific
// filter.
type LandingZoneOperationFilter struct {

	// The statuses of the set of landing zone operations selected by the filter.
	Statuses []LandingZoneOperationStatus

	// The set of landing zone operation types selected by the filter.
	Types []LandingZoneOperationType

	noSmithyDocumentSerde
}

// Returns a summary of information about a landing zone operation.
type LandingZoneOperationSummary struct {

	// The operationIdentifier of the landing zone operation.
	OperationIdentifier *string

	// The type of the landing zone operation.
	OperationType LandingZoneOperationType

	// The status of the landing zone operation.
	Status LandingZoneOperationStatus

	noSmithyDocumentSerde
}

// Returns a summary of information about a landing zone.
type LandingZoneSummary struct {

	// The ARN of the landing zone.
	Arn *string

	noSmithyDocumentSerde
}

// An Amazon Web Services Region in which Amazon Web Services Control Tower
// expects to find the control deployed.
//
// The expected Regions are based on the Regions that are governed by the landing
// zone. In certain cases, a control is not actually enabled in the Region as
// expected, such as during drift, or [mixed governance].
//
// [mixed governance]: https://docs.aws.amazon.com/controltower/latest/userguide/region-how.html#mixed-governance
type Region struct {

	// The Amazon Web Services Region name.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde