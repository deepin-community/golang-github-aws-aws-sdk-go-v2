// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about a specific identity associated with your Amazon
// Pinpoint account, including the identity's verification status, its DKIM
// authentication status, and its custom Mail-From settings.
func (c *Client) GetEmailIdentity(ctx context.Context, params *GetEmailIdentityInput, optFns ...func(*Options)) (*GetEmailIdentityOutput, error) {
	if params == nil {
		params = &GetEmailIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEmailIdentity", params, optFns, c.addOperationGetEmailIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEmailIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to return details about an email identity.
type GetEmailIdentityInput struct {

	// The email identity that you want to retrieve details for.
	//
	// This member is required.
	EmailIdentity *string

	noSmithyDocumentSerde
}

// Details about an email identity.
type GetEmailIdentityOutput struct {

	// An object that contains information about the DKIM attributes for the identity.
	// This object includes the tokens that you use to create the CNAME records that
	// are required to complete the DKIM verification process.
	DkimAttributes *types.DkimAttributes

	// The feedback forwarding configuration for the identity.
	//
	// If the value is true , Amazon Pinpoint sends you email notifications when bounce
	// or complaint events occur. Amazon Pinpoint sends this notification to the
	// address that you specified in the Return-Path header of the original email.
	//
	// When you set this value to false , Amazon Pinpoint sends notifications through
	// other mechanisms, such as by notifying an Amazon SNS topic or another event
	// destination. You're required to have a method of tracking bounces and
	// complaints. If you haven't set up another mechanism for receiving bounce or
	// complaint notifications, Amazon Pinpoint sends an email notification when these
	// events occur (even if this setting is disabled).
	FeedbackForwardingStatus bool

	// The email identity type.
	IdentityType types.IdentityType

	// An object that contains information about the Mail-From attributes for the
	// email identity.
	MailFromAttributes *types.MailFromAttributes

	// An array of objects that define the tags (keys and values) that are associated
	// with the email identity.
	Tags []types.Tag

	// Specifies whether or not the identity is verified. In Amazon Pinpoint, you can
	// only send email from verified email addresses or domains. For more information
	// about verifying identities, see the [Amazon Pinpoint User Guide].
	//
	// [Amazon Pinpoint User Guide]: https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html
	VerifiedForSendingStatus bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEmailIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEmailIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEmailIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEmailIdentity"); err != nil {
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
	if err = addOpGetEmailIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEmailIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEmailIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEmailIdentity",
	}
}