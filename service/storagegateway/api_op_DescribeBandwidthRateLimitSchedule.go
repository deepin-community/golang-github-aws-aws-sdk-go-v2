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

//	Returns information about the bandwidth rate limit schedule of a gateway. By
//
// default, gateways do not have bandwidth rate limit schedules, which means no
// bandwidth rate limiting is in effect. This operation is supported only for
// volume, tape and S3 file gateways. FSx file gateways do not support bandwidth
// rate limits.
//
// This operation returns information about a gateway's bandwidth rate limit
// schedule. A bandwidth rate limit schedule consists of one or more bandwidth rate
// limit intervals. A bandwidth rate limit interval defines a period of time on one
// or more days of the week, during which bandwidth rate limits are specified for
// uploading, downloading, or both.
//
// A bandwidth rate limit interval consists of one or more days of the week, a
// start hour and minute, an ending hour and minute, and bandwidth rate limits for
// uploading and downloading
//
// If no bandwidth rate limit schedule intervals are set for the gateway, this
// operation returns an empty response. To specify which gateway to describe, use
// the Amazon Resource Name (ARN) of the gateway in your request.
func (c *Client) DescribeBandwidthRateLimitSchedule(ctx context.Context, params *DescribeBandwidthRateLimitScheduleInput, optFns ...func(*Options)) (*DescribeBandwidthRateLimitScheduleOutput, error) {
	if params == nil {
		params = &DescribeBandwidthRateLimitScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBandwidthRateLimitSchedule", params, optFns, c.addOperationDescribeBandwidthRateLimitScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBandwidthRateLimitScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBandwidthRateLimitScheduleInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

type DescribeBandwidthRateLimitScheduleOutput struct {

	//  An array that contains the bandwidth rate limit intervals for a tape or volume
	// gateway.
	BandwidthRateLimitIntervals []types.BandwidthRateLimitInterval

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBandwidthRateLimitScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBandwidthRateLimitSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBandwidthRateLimitSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBandwidthRateLimitSchedule"); err != nil {
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
	if err = addOpDescribeBandwidthRateLimitScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBandwidthRateLimitSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBandwidthRateLimitSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBandwidthRateLimitSchedule",
	}
}