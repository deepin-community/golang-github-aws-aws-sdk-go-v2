// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves metadata for a version of a component.
func (c *Client) DescribeComponent(ctx context.Context, params *DescribeComponentInput, optFns ...func(*Options)) (*DescribeComponentOutput, error) {
	if params == nil {
		params = &DescribeComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComponent", params, optFns, c.addOperationDescribeComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeComponentInput struct {

	// The [ARN] of the component version.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type DescribeComponentOutput struct {

	// The [ARN] of the component version.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	Arn *string

	// The name of the component.
	ComponentName *string

	// The version of the component.
	ComponentVersion *string

	// The time at which the component was created, expressed in ISO 8601 format.
	CreationTimestamp *time.Time

	// The description of the component version.
	Description *string

	// The platforms that the component version supports.
	Platforms []types.ComponentPlatform

	// The publisher of the component version.
	Publisher *string

	// The status of the component version in IoT Greengrass V2. This status is
	// different from the status of the component on a core device.
	Status *types.CloudComponentStatus

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see [Tag your resources]in the IoT Greengrass V2 Developer Guide.
	//
	// [Tag your resources]: https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeComponent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeComponent"); err != nil {
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
	if err = addOpDescribeComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComponent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeComponent",
	}
}