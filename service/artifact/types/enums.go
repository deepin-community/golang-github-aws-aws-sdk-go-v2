// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AcceptanceType string

// Enum values for AcceptanceType
const (
	// Do not require explicit click-through acceptance of the Term associated with
	// this Report.
	AcceptanceTypePassthrough AcceptanceType = "PASSTHROUGH"
	// Require explicit click-through acceptance of the Term associated with this
	// Report.
	AcceptanceTypeExplicit AcceptanceType = "EXPLICIT"
)

// Values returns all known values for AcceptanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AcceptanceType) Values() []AcceptanceType {
	return []AcceptanceType{
		"PASSTHROUGH",
		"EXPLICIT",
	}
}

type NotificationSubscriptionStatus string

// Enum values for NotificationSubscriptionStatus
const (
	// The account is subscribed for notification.
	NotificationSubscriptionStatusSubscribed NotificationSubscriptionStatus = "SUBSCRIBED"
	// The account is not subscribed for notification.
	NotificationSubscriptionStatusNotSubscribed NotificationSubscriptionStatus = "NOT_SUBSCRIBED"
)

// Values returns all known values for NotificationSubscriptionStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NotificationSubscriptionStatus) Values() []NotificationSubscriptionStatus {
	return []NotificationSubscriptionStatus{
		"SUBSCRIBED",
		"NOT_SUBSCRIBED",
	}
}

type PublishedState string

// Enum values for PublishedState
const (
	// The resource is published for consumption.
	PublishedStatePublished PublishedState = "PUBLISHED"
	// The resource is not published for consumption.
	PublishedStateUnpublished PublishedState = "UNPUBLISHED"
)

// Values returns all known values for PublishedState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PublishedState) Values() []PublishedState {
	return []PublishedState{
		"PUBLISHED",
		"UNPUBLISHED",
	}
}

type UploadState string

// Enum values for UploadState
const (
	UploadStateProcessing UploadState = "PROCESSING"
	UploadStateComplete   UploadState = "COMPLETE"
	UploadStateFailed     UploadState = "FAILED"
	UploadStateFault      UploadState = "FAULT"
)

// Values returns all known values for UploadState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (UploadState) Values() []UploadState {
	return []UploadState{
		"PROCESSING",
		"COMPLETE",
		"FAILED",
		"FAULT",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonInvalidToken          ValidationExceptionReason = "invalidToken"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"invalidToken",
		"other",
	}
}