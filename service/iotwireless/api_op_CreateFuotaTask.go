// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a FUOTA task.
func (c *Client) CreateFuotaTask(ctx context.Context, params *CreateFuotaTaskInput, optFns ...func(*Options)) (*CreateFuotaTaskOutput, error) {
	if params == nil {
		params = &CreateFuotaTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFuotaTask", params, optFns, c.addOperationCreateFuotaTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFuotaTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFuotaTaskInput struct {

	// The S3 URI points to a firmware update image that is to be used with a FUOTA
	// task.
	//
	// This member is required.
	FirmwareUpdateImage *string

	// The firmware update role that is to be used with a FUOTA task.
	//
	// This member is required.
	FirmwareUpdateRole *string

	// Each resource must have a unique client request token. The client token is used
	// to implement idempotency. It ensures that the request completes no more than one
	// time. If you retry a request with the same token and the same parameters, the
	// request will complete successfully. However, if you try to create a new resource
	// using the same token but different parameters, an HTTP 409 conflict occurs. If
	// you omit this value, AWS SDKs will automatically generate a unique client
	// request. For more information about idempotency, see [Ensuring idempotency in Amazon EC2 API requests].
	//
	// [Ensuring idempotency in Amazon EC2 API requests]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientRequestToken *string

	// The description of the new resource.
	Description *string

	// The interval for sending fragments in milliseconds, rounded to the nearest
	// second.
	//
	// This interval only determines the timing for when the Cloud sends down the
	// fragments to yor device. There can be a delay for when your device will receive
	// these fragments. This delay depends on the device's class and the communication
	// delay with the cloud.
	FragmentIntervalMS *int32

	// The size of each fragment in bytes. This parameter is supported only for FUOTA
	// tasks with multicast groups.
	FragmentSizeBytes *int32

	// The LoRaWAN information used with a FUOTA task.
	LoRaWAN *types.LoRaWANFuotaTask

	// The name of a FUOTA task.
	Name *string

	// The percentage of the added fragments that are redundant. For example, if the
	// size of the firmware image file is 100 bytes and the fragment size is 10 bytes,
	// with RedundancyPercent set to 50(%), the final number of encoded fragments is
	// (100 / 10) + (100 / 10 * 50%) = 15.
	RedundancyPercent *int32

	// The tag to attach to the specified resource. Tags are metadata that you can use
	// to manage a resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFuotaTaskOutput struct {

	// The arn of a FUOTA task.
	Arn *string

	// The ID of a FUOTA task.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFuotaTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFuotaTask"); err != nil {
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
	if err = addIdempotencyToken_opCreateFuotaTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFuotaTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFuotaTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFuotaTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFuotaTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFuotaTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFuotaTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFuotaTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFuotaTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFuotaTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFuotaTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFuotaTask",
	}
}