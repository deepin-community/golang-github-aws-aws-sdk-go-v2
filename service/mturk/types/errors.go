// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Your request is invalid.
type RequestError struct {
	Message *string

	ErrorCodeOverride *string

	TurkErrorCode *string

	noSmithyDocumentSerde
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RequestError"
	}
	return *e.ErrorCodeOverride
}
func (e *RequestError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Mechanical Turk is temporarily unable to process your request. Try your
// call again.
type ServiceFault struct {
	Message *string

	ErrorCodeOverride *string

	TurkErrorCode *string

	noSmithyDocumentSerde
}

func (e *ServiceFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceFault) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }