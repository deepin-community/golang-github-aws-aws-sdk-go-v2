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

// Returns your gateway's maintenance window schedule information, with values for
// monthly or weekly cadence, specific day and time to begin maintenance, and which
// types of updates to apply. Time values returned are for the gateway's time zone.
func (c *Client) DescribeMaintenanceStartTime(ctx context.Context, params *DescribeMaintenanceStartTimeInput, optFns ...func(*Options)) (*DescribeMaintenanceStartTimeOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceStartTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceStartTime", params, optFns, c.addOperationDescribeMaintenanceStartTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
//
// # DescribeMaintenanceStartTimeOutput$SoftwareUpdatePreferences
//
// # DescribeMaintenanceStartTimeOutput$DayOfMonth
//
// # DescribeMaintenanceStartTimeOutput$DayOfWeek
//
// # DescribeMaintenanceStartTimeOutput$HourOfDay
//
// # DescribeMaintenanceStartTimeOutput$MinuteOfHour
//
// DescribeMaintenanceStartTimeOutput$Timezone
type DescribeMaintenanceStartTimeOutput struct {

	// The day of the month component of the maintenance start time represented as an
	// ordinal number from 1 to 28, where 1 represents the first day of the month. It
	// is not possible to set the maintenance schedule to start on days 29 through 31.
	DayOfMonth *int32

	// An ordinal number between 0 and 6 that represents the day of the week, where 0
	// represents Sunday and 6 represents Saturday. The day of week is in the time zone
	// of the gateway.
	DayOfWeek *int32

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// The hour component of the maintenance start time represented as hh, where hh is
	// the hour (0 to 23). The hour of the day is in the time zone of the gateway.
	HourOfDay *int32

	// The minute component of the maintenance start time represented as mm, where mm
	// is the minute (0 to 59). The minute of the hour is in the time zone of the
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

	// A value that indicates the time zone that is set for the gateway. The start
	// time and day of week specified should be in the time zone of the gateway.
	Timezone *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceStartTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceStartTime{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMaintenanceStartTime"); err != nil {
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
	if err = addOpDescribeMaintenanceStartTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceStartTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMaintenanceStartTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMaintenanceStartTime",
	}
}