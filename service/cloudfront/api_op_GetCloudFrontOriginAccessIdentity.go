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

// Get the information about an origin access identity.
func (c *Client) GetCloudFrontOriginAccessIdentity(ctx context.Context, params *GetCloudFrontOriginAccessIdentityInput, optFns ...func(*Options)) (*GetCloudFrontOriginAccessIdentityOutput, error) {
	if params == nil {
		params = &GetCloudFrontOriginAccessIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCloudFrontOriginAccessIdentity", params, optFns, c.addOperationGetCloudFrontOriginAccessIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCloudFrontOriginAccessIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to get an origin access identity's information.
type GetCloudFrontOriginAccessIdentityInput struct {

	// The identity's ID.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type GetCloudFrontOriginAccessIdentityOutput struct {

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *types.CloudFrontOriginAccessIdentity

	// The current version of the origin access identity's information. For example:
	// E2QWRUHAPOMQZL .
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCloudFrontOriginAccessIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetCloudFrontOriginAccessIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetCloudFrontOriginAccessIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCloudFrontOriginAccessIdentity"); err != nil {
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
	if err = addOpGetCloudFrontOriginAccessIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCloudFrontOriginAccessIdentity",
	}
}