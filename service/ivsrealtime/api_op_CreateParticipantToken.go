// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an additional token for a specified stage. This can be done after stage
// creation or when tokens expire. Tokens always are scoped to the stage for which
// they are created.
//
// Encryption keys are owned by Amazon IVS and never used directly by your
// application.
func (c *Client) CreateParticipantToken(ctx context.Context, params *CreateParticipantTokenInput, optFns ...func(*Options)) (*CreateParticipantTokenOutput, error) {
	if params == nil {
		params = &CreateParticipantTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateParticipantToken", params, optFns, c.addOperationCreateParticipantTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateParticipantTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateParticipantTokenInput struct {

	// ARN of the stage to which this token is scoped.
	//
	// This member is required.
	StageArn *string

	// Application-provided attributes to encode into the token and attach to a stage.
	// Map keys and values can contain UTF-8 encoded text. The maximum length of this
	// field is 1 KB total. This field is exposed to all stage participants and should
	// not be used for personally identifying, confidential, or sensitive information.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the stage. Default:
	// PUBLISH, SUBSCRIBE .
	Capabilities []types.ParticipantTokenCapability

	// Duration (in minutes), after which the token expires. Default: 720 (12 hours).
	Duration *int32

	// Name that can be specified to help identify the token. This can be any UTF-8
	// encoded text. This field is exposed to all stage participants and should not be
	// used for personally identifying, confidential, or sensitive information.
	UserId *string

	noSmithyDocumentSerde
}

type CreateParticipantTokenOutput struct {

	// The participant token that was created.
	ParticipantToken *types.ParticipantToken

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateParticipantTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateParticipantToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateParticipantToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateParticipantToken"); err != nil {
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
	if err = addOpCreateParticipantTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateParticipantToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateParticipantToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateParticipantToken",
	}
}