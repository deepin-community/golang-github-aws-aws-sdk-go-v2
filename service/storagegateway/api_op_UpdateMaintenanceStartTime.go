// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a gateway's maintenance window schedule, with settings for monthly or
// weekly cadence, specific day and time to begin maintenance, and which types of
// updates to apply. Time configuration uses the gateway's time zone. You can pass
// values for a complete maintenance schedule, or update policy, or both. Previous
// values will persist for whichever setting you choose not to modify. If an
// incomplete or invalid maintenance schedule is passed, the entire request will be
// rejected with an error and no changes will occur.
//
// A complete maintenance schedule must include values for both MinuteOfHour and
// HourOfDay , and either DayOfMonth or DayOfWeek .
//
// We recommend keeping maintenance updates turned on, except in specific use
// cases where the brief disruptions caused by updating the gateway could
// critically impact your deployment.
func (c *Client) UpdateMaintenanceStartTime(ctx context.Context, params *UpdateMaintenanceStartTimeInput, optFns ...func(*Options)) (*UpdateMaintenanceStartTimeOutput, error) {
	if params == nil {
		params = &UpdateMaintenanceStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMaintenanceStartTime", params, optFns, c.addOperationUpdateMaintenanceStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the following fields:
//
// # UpdateMaintenanceStartTimeInput$SoftwareUpdatePreferences
//
// # UpdateMaintenanceStartTimeInput$DayOfMonth
//
// # UpdateMaintenanceStartTimeInput$DayOfWeek
//
// # UpdateMaintenanceStartTimeInput$HourOfDay
//
// UpdateMaintenanceStartTimeInput$MinuteOfHour
type UpdateMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The day of the month component of the maintenance start time represented as an
	// ordinal number from 1 to 28, where 1 represents the first day of the month. It
	// is not possible to set the maintenance schedule to start on days 29 through 31.
	DayOfMonth *int32

	// The day of the week component of the maintenance start time week represented as
	// an ordinal number from 0 to 6, where 0 represents Sunday and 6 represents
	// Saturday.
	DayOfWeek *int32

	// The hour component of the maintenance start time represented as hh, where hh is
	// the hour (00 to 23). The hour of the day is in the time zone of the gateway.
	HourOfDay *int32

	// The minute component of the maintenance start time represented as mm, where mm
	// is the minute (00 to 59). The minute of the hour is in the time zone of the
	// gateway.
	MinuteOfHour *int32

	// A set of variables indicating the software update preferences for the gateway.
	//
	// Includes AutomaticUpdatePolicy field with the following inputs:
	//
	// ALL_VERSIONS - Enables regular gateway maintenance updates.
	//
	// EMERGENCY_VERSIONS_ONLY - Disables regular gateway maintenance updates.
	SoftwareUpdatePreferences *types.SoftwareUpdatePreferences

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// maintenance start time is updated.
type UpdateMaintenanceStartTimeOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMaintenanceStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMaintenanceStartTime"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateMaintenanceStartTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceStartTime(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateMaintenanceStartTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMaintenanceStartTime",
	}
}