// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new origin access control in CloudFront. After you create an origin
// access control, you can add it to an origin in a CloudFront distribution so that
// CloudFront sends authenticated (signed) requests to the origin.
//
// This makes it possible to block public access to the origin, allowing viewers
// (users) to access the origin's content only through CloudFront.
//
// For more information about using a CloudFront origin access control, see [Restricting access to an Amazon Web Services origin] in
// the Amazon CloudFront Developer Guide.
//
// [Restricting access to an Amazon Web Services origin]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-origin.html
func (c *Client) CreateOriginAccessControl(ctx context.Context, params *CreateOriginAccessControlInput, optFns ...func(*Options)) (*CreateOriginAccessControlOutput, error) {
	if params == nil {
		params = &CreateOriginAccessControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOriginAccessControl", params, optFns, c.addOperationCreateOriginAccessControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOriginAccessControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOriginAccessControlInput struct {

	// Contains the origin access control.
	//
	// This member is required.
	OriginAccessControlConfig *types.OriginAccessControlConfig

	noSmithyDocumentSerde
}

type CreateOriginAccessControlOutput struct {

	// The version identifier for the current version of the origin access control.
	ETag *string

	// The URL of the origin access control.
	Location *string

	// Contains an origin access control.
	OriginAccessControl *types.OriginAccessControl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOriginAccessControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateOriginAccessControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateOriginAccessControl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateOriginAccessControl"); err != nil {
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
	if err = addOpCreateOriginAccessControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOriginAccessControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOriginAccessControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateOriginAccessControl",
	}
}