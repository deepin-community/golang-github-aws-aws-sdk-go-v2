// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This exception is thrown when an internal service error occurs.
type MarketplaceCommerceAnalyticsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MarketplaceCommerceAnalyticsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MarketplaceCommerceAnalyticsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MarketplaceCommerceAnalyticsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MarketplaceCommerceAnalyticsException"
	}
	return *e.ErrorCodeOverride
}
func (e *MarketplaceCommerceAnalyticsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}