// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a proxy session on the specified Amazon Chime Voice Connector for the
// specified participant phone numbers.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [CreateProxySession], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by CreateProxySession in the Amazon Chime SDK Voice
// Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [CreateProxySession]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_CreateProxySession.html
func (c *Client) CreateProxySession(ctx context.Context, params *CreateProxySessionInput, optFns ...func(*Options)) (*CreateProxySessionOutput, error) {
	if params == nil {
		params = &CreateProxySessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProxySession", params, optFns, c.addOperationCreateProxySessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProxySessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProxySessionInput struct {

	// The proxy session capabilities.
	//
	// This member is required.
	Capabilities []types.Capability

	// The participant phone numbers.
	//
	// This member is required.
	ParticipantPhoneNumbers []string

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The number of minutes allowed for the proxy session.
	ExpiryMinutes *int32

	// The preference for matching the country or area code of the proxy phone number
	// with that of the first participant.
	GeoMatchLevel types.GeoMatchLevel

	// The country and area code for the proxy phone number.
	GeoMatchParams *types.GeoMatchParams

	// The name of the proxy session.
	Name *string

	// The preference for proxy phone number reuse, or stickiness, between the same
	// participants across sessions.
	NumberSelectionBehavior types.NumberSelectionBehavior

	noSmithyDocumentSerde
}

type CreateProxySessionOutput struct {

	// The proxy session details.
	ProxySession *types.ProxySession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProxySessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProxySession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProxySession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProxySession"); err != nil {
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
	if err = addOpCreateProxySessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProxySession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProxySession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProxySession",
	}
}