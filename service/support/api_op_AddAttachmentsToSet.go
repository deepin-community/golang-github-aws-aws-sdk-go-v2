// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more attachments to an attachment set.
//
// An attachment set is a temporary container for attachments that you add to a
// case or case communication. The set is available for 1 hour after it's created.
// The expiryTime returned in the response is when the set expires.
//
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func (c *Client) AddAttachmentsToSet(ctx context.Context, params *AddAttachmentsToSetInput, optFns ...func(*Options)) (*AddAttachmentsToSetOutput, error) {
	if params == nil {
		params = &AddAttachmentsToSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddAttachmentsToSet", params, optFns, c.addOperationAddAttachmentsToSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddAttachmentsToSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddAttachmentsToSetInput struct {

	// One or more attachments to add to the set. You can add up to three attachments
	// per set. The size limit is 5 MB per attachment.
	//
	// In the Attachment object, use the data parameter to specify the contents of the
	// attachment file. In the previous request syntax, the value for data appear as
	// blob , which is represented as a base64-encoded string. The value for fileName
	// is the name of the attachment, such as troubleshoot-screenshot.png .
	//
	// This member is required.
	Attachments []types.Attachment

	// The ID of the attachment set. If an attachmentSetId is not specified, a new
	// attachment set is created, and the ID of the set is returned in the response. If
	// an attachmentSetId is specified, the attachments are added to the specified
	// set, if it exists.
	AttachmentSetId *string

	noSmithyDocumentSerde
}

// The ID and expiry time of the attachment set returned by the AddAttachmentsToSet operation.
type AddAttachmentsToSetOutput struct {

	// The ID of the attachment set. If an attachmentSetId was not specified, a new
	// attachment set is created, and the ID of the set is returned in the response. If
	// an attachmentSetId was specified, the attachments are added to the specified
	// set, if it exists.
	AttachmentSetId *string

	// The time and date when the attachment set expires.
	ExpiryTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddAttachmentsToSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddAttachmentsToSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddAttachmentsToSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddAttachmentsToSet"); err != nil {
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
	if err = addOpAddAttachmentsToSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddAttachmentsToSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddAttachmentsToSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddAttachmentsToSet",
	}
}