// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"time"
)

// Get the latest thumbnail data for the input device.
func (c *Client) DescribeInputDeviceThumbnail(ctx context.Context, params *DescribeInputDeviceThumbnailInput, optFns ...func(*Options)) (*DescribeInputDeviceThumbnailOutput, error) {
	if params == nil {
		params = &DescribeInputDeviceThumbnailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInputDeviceThumbnail", params, optFns, c.addOperationDescribeInputDeviceThumbnailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInputDeviceThumbnailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeInputDeviceThumbnailRequest
type DescribeInputDeviceThumbnailInput struct {

	// The HTTP Accept header. Indicates the requested type for the thumbnail.
	//
	// This member is required.
	Accept types.AcceptHeader

	// The unique ID of this input device. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DescribeInputDeviceThumbnailResponse
type DescribeInputDeviceThumbnailOutput struct {

	// The binary data for the thumbnail that the Link device has most recently sent
	// to MediaLive.
	Body io.ReadCloser

	// The length of the content.
	ContentLength *int64

	// Specifies the media type of the thumbnail.
	ContentType types.ContentType

	// The unique, cacheable version of this thumbnail.
	ETag *string

	// The date and time the thumbnail was last updated at the device.
	LastModified *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInputDeviceThumbnailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInputDeviceThumbnail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInputDeviceThumbnail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInputDeviceThumbnail"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeInputDeviceThumbnailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInputDeviceThumbnail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInputDeviceThumbnail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInputDeviceThumbnail",
	}
}