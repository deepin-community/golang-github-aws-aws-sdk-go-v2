// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
//
// Updates the outbound caller ID name, number, and outbound whisper flow for a
// specified queue.
//
//   - If the phone number is claimed to a traffic distribution group that was
//     created in the same Region as the Amazon Connect instance where you are calling
//     this API, then you can use a full phone number ARN or a UUID for
//     OutboundCallerIdNumberId . However, if the phone number is claimed to a
//     traffic distribution group that is in one Region, and you are calling this API
//     from an instance in another Amazon Web Services Region that is associated with
//     the traffic distribution group, you must provide a full phone number ARN. If a
//     UUID is provided in this scenario, you will receive a
//     ResourceNotFoundException .
//
//   - Only use the phone number ARN format that doesn't contain instance in the
//     path, for example, arn:aws:connect:us-east-1:1234567890:phone-number/uuid .
//     This is the same ARN format that is returned when you call the [ListPhoneNumbersV2]API.
//
//   - If you plan to use IAM policies to allow/deny access to this API for phone
//     number resources claimed to a traffic distribution group, see [Allow or Deny queue API actions for phone numbers in a replica Region].
//
// [ListPhoneNumbersV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html
// [Allow or Deny queue API actions for phone numbers in a replica Region]: https://docs.aws.amazon.com/connect/latest/adminguide/security_iam_resource-level-policy-examples.html#allow-deny-queue-actions-replica-region
func (c *Client) UpdateQueueOutboundCallerConfig(ctx context.Context, params *UpdateQueueOutboundCallerConfigInput, optFns ...func(*Options)) (*UpdateQueueOutboundCallerConfigOutput, error) {
	if params == nil {
		params = &UpdateQueueOutboundCallerConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQueueOutboundCallerConfig", params, optFns, c.addOperationUpdateQueueOutboundCallerConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQueueOutboundCallerConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQueueOutboundCallerConfigInput struct {

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The outbound caller ID name, number, and outbound whisper flow.
	//
	// This member is required.
	OutboundCallerConfig *types.OutboundCallerConfig

	// The identifier for the queue.
	//
	// This member is required.
	QueueId *string

	noSmithyDocumentSerde
}

type UpdateQueueOutboundCallerConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateQueueOutboundCallerConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateQueueOutboundCallerConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateQueueOutboundCallerConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateQueueOutboundCallerConfig"); err != nil {
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
	if err = addOpUpdateQueueOutboundCallerConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQueueOutboundCallerConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQueueOutboundCallerConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateQueueOutboundCallerConfig",
	}
}