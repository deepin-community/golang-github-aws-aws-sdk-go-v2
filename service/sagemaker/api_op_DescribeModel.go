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

// Describes a model that you created using the CreateModel API.
func (c *Client) DescribeModel(ctx context.Context, params *DescribeModelInput, optFns ...func(*Options)) (*DescribeModelOutput, error) {
	if params == nil {
		params = &DescribeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModel", params, optFns, c.addOperationDescribeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelInput struct {

	// The name of the model.
	//
	// This member is required.
	ModelName *string

	noSmithyDocumentSerde
}

type DescribeModelOutput struct {

	// A timestamp that shows when the model was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model.
	//
	// This member is required.
	ModelArn *string

	// Name of the SageMaker model.
	//
	// This member is required.
	ModelName *string

	// The containers in the inference pipeline.
	Containers []types.ContainerDefinition

	// A set of recommended deployment configurations for the model.
	DeploymentRecommendation *types.DeploymentRecommendation

	// If True , no inbound or outbound network calls can be made to or from the model
	// container.
	EnableNetworkIsolation *bool

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the model.
	ExecutionRoleArn *string

	// Specifies details of how containers in a multi-container endpoint are called.
	InferenceExecutionConfig *types.InferenceExecutionConfig

	// The location of the primary inference code, associated artifacts, and custom
	// environment map that the inference code uses when it is deployed in production.
	PrimaryContainer *types.ContainerDefinition

	// A [VpcConfig] object that specifies the VPC that this model has access to. For more
	// information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud]
	//
	// [VpcConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html
	// [Protect Endpoints by Using an Amazon Virtual Private Cloud]: https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html
	VpcConfig *types.VpcConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeModel"); err != nil {
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
	if err = addOpDescribeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeModel",
	}
}