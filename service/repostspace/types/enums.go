// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ConfigurationStatus string

// Enum values for ConfigurationStatus
const (
	ConfigurationStatusConfigured   ConfigurationStatus = "CONFIGURED"
	ConfigurationStatusUnconfigured ConfigurationStatus = "UNCONFIGURED"
)

// Values returns all known values for ConfigurationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationStatus) Values() []ConfigurationStatus {
	return []ConfigurationStatus{
		"CONFIGURED",
		"UNCONFIGURED",
	}
}

type TierLevel string

// Enum values for TierLevel
const (
	TierLevelBasic    TierLevel = "BASIC"
	TierLevelStandard TierLevel = "STANDARD"
)

// Values returns all known values for TierLevel. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TierLevel) Values() []TierLevel {
	return []TierLevel{
		"BASIC",
		"STANDARD",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
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
		"other",
	}
}

type VanityDomainStatus string

// Enum values for VanityDomainStatus
const (
	VanityDomainStatusPending    VanityDomainStatus = "PENDING"
	VanityDomainStatusApproved   VanityDomainStatus = "APPROVED"
	VanityDomainStatusUnapproved VanityDomainStatus = "UNAPPROVED"
)

// Values returns all known values for VanityDomainStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VanityDomainStatus) Values() []VanityDomainStatus {
	return []VanityDomainStatus{
		"PENDING",
		"APPROVED",
		"UNAPPROVED",
	}
}