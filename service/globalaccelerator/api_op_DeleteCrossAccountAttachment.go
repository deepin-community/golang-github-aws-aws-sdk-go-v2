// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a cross-account attachment. When you delete an attachment, Global
// Accelerator revokes the permission to use the resources in the attachment from
// all principals in the list of principals. Global Accelerator revokes the
// permission for specific resources.
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
func (c *Client) DeleteCrossAccountAttachment(ctx context.Context, params *DeleteCrossAccountAttachmentInput, optFns ...func(*Options)) (*DeleteCrossAccountAttachmentOutput, error) {
	if params == nil {
		params = &DeleteCrossAccountAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCrossAccountAttachment", params, optFns, c.addOperationDeleteCrossAccountAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCrossAccountAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCrossAccountAttachmentInput struct {

	// The Amazon Resource Name (ARN) for the cross-account attachment to delete.
	//
	// This member is required.
	AttachmentArn *string

	noSmithyDocumentSerde
}

type DeleteCrossAccountAttachmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCrossAccountAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCrossAccountAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCrossAccountAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteCrossAccountAttachment"); err != nil {
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
	if err = addOpDeleteCrossAccountAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCrossAccountAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCrossAccountAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteCrossAccountAttachment",
	}
}