// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the description of an endpoint configuration created using the
// CreateEndpointConfig API.
func (c *Client) DescribeEndpointConfig(ctx context.Context, params *DescribeEndpointConfigInput, optFns ...func(*Options)) (*DescribeEndpointConfigOutput, error) {
	if params == nil {
		params = &DescribeEndpointConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointConfig", params, optFns, c.addOperationDescribeEndpointConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointConfigInput struct {

	// The name of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string

	noSmithyDocumentSerde
}

type DescribeEndpointConfigOutput struct {

	// A timestamp that shows when the endpoint configuration was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigArn *string

	// Name of the SageMaker endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string

	// An array of ProductionVariant objects, one for each model that you want to host
	// at this endpoint.
	//
	// This member is required.
	ProductionVariants []types.ProductionVariant

	// Returns the description of an endpoint configuration created using the [CreateEndpointConfig]
	// CreateEndpointConfig API.
	//
	// [CreateEndpointConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html
	AsyncInferenceConfig *types.AsyncInferenceConfig

	// Configuration to control how SageMaker captures inference data.
	DataCaptureConfig *types.DataCaptureConfig

	// Indicates whether all model containers deployed to the endpoint are isolated.
	// If they are, no inbound or outbound network calls can be made to or from the
	// model containers.
	EnableNetworkIsolation *bool

	// The Amazon Resource Name (ARN) of the IAM role that you assigned to the
	// endpoint configuration.
	ExecutionRoleArn *string

	// The configuration parameters for an explainer.
	ExplainerConfig *types.ExplainerConfig

	// Amazon Web Services KMS key ID Amazon SageMaker uses to encrypt data when
	// storing it on the ML storage volume attached to the instance.
	KmsKeyId *string

	// An array of ProductionVariant objects, one for each model that you want to host
	// at this endpoint in shadow mode with production traffic replicated from the
	// model specified on ProductionVariants .
	ShadowProductionVariants []types.ProductionVariant

	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs,
	// hosted models, and compute resources have access to. You can control access to
	// and from your resources by configuring a VPC. For more information, see [Give SageMaker Access to Resources in your Amazon VPC].
	//
	// [Give SageMaker Access to Resources in your Amazon VPC]: https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html
	VpcConfig *types.VpcConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpointConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpointConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEndpointConfig"); err != nil {
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
	if err = addOpDescribeEndpointConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEndpointConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEndpointConfig",
	}
}