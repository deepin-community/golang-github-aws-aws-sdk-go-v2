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

// Get the configuration information about an origin access identity.
func (c *Client) GetCloudFrontOriginAccessIdentityConfig(ctx context.Context, params *GetCloudFrontOriginAccessIdentityConfigInput, optFns ...func(*Options)) (*GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	if params == nil {
		params = &GetCloudFrontOriginAccessIdentityConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCloudFrontOriginAccessIdentityConfig", params, optFns, c.addOperationGetCloudFrontOriginAccessIdentityConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCloudFrontOriginAccessIdentityConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The origin access identity's configuration information. For more information,
// see [CloudFrontOriginAccessIdentityConfig].
//
// [CloudFrontOriginAccessIdentityConfig]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CloudFrontOriginAccessIdentityConfig.html
type GetCloudFrontOriginAccessIdentityConfigInput struct {

	// The identity's ID.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type GetCloudFrontOriginAccessIdentityConfigOutput struct {

	// The origin access identity's configuration information.
	CloudFrontOriginAccessIdentityConfig *types.CloudFrontOriginAccessIdentityConfig

	// The current version of the configuration. For example: E2QWRUHAPOMQZL .
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCloudFrontOriginAccessIdentityConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetCloudFrontOriginAccessIdentityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetCloudFrontOriginAccessIdentityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCloudFrontOriginAccessIdentityConfig"); err != nil {
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
	if err = addOpGetCloudFrontOriginAccessIdentityConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentityConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentityConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCloudFrontOriginAccessIdentityConfig",
	}
}