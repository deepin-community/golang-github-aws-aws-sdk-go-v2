// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new opt-out list.
//
// If the opt-out list name already exists, an error is returned.
//
// An opt-out list is a list of phone numbers that are opted out, meaning you
// can't send SMS or voice messages to them. If end user replies with the keyword
// "STOP," an entry for the phone number is added to the opt-out list. In addition
// to STOP, your recipients can use any supported opt-out keyword, such as CANCEL
// or OPTOUT. For a list of supported opt-out keywords, see [SMS opt out]in the Amazon Pinpoint
// User Guide.
//
// [SMS opt out]: https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-manage.html#channels-sms-manage-optout
func (c *Client) CreateOptOutList(ctx context.Context, params *CreateOptOutListInput, optFns ...func(*Options)) (*CreateOptOutListOutput, error) {
	if params == nil {
		params = &CreateOptOutListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOptOutList", params, optFns, c.addOperationCreateOptOutListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOptOutListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOptOutListInput struct {

	// The name of the new OptOutList.
	//
	// This member is required.
	OptOutListName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	// An array of tags (key and value pairs) to associate with the new OptOutList.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateOptOutListOutput struct {

	// The time when the pool was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) for the OptOutList.
	OptOutListArn *string

	// The name of the new OptOutList.
	OptOutListName *string

	// An array of tags (key and value pairs) associated with the new OptOutList.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOptOutListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateOptOutList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateOptOutList{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateOptOutList"); err != nil {
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
	if err = addIdempotencyToken_opCreateOptOutListMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateOptOutListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOptOutList(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateOptOutList struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateOptOutList) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateOptOutList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateOptOutListInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateOptOutListInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateOptOutListMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateOptOutList{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateOptOutList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateOptOutList",
	}
}