// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the name or associated model for a Provisioned Throughput. For more
// information, see [Provisioned Throughput]in the Amazon Bedrock User Guide.
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
func (c *Client) UpdateProvisionedModelThroughput(ctx context.Context, params *UpdateProvisionedModelThroughputInput, optFns ...func(*Options)) (*UpdateProvisionedModelThroughputOutput, error) {
	if params == nil {
		params = &UpdateProvisionedModelThroughputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProvisionedModelThroughput", params, optFns, c.addOperationUpdateProvisionedModelThroughputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProvisionedModelThroughputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisionedModelThroughputInput struct {

	// The Amazon Resource Name (ARN) or name of the Provisioned Throughput to update.
	//
	// This member is required.
	ProvisionedModelId *string

	// The Amazon Resource Name (ARN) of the new model to associate with this
	// Provisioned Throughput. You can't specify this field if this Provisioned
	// Throughput is associated with a base model.
	//
	// If this Provisioned Throughput is associated with a custom model, you can
	// specify one of the following options:
	//
	//   - The base model from which the custom model was customized.
	//
	//   - Another custom model that was customized from the same base model as the
	//   custom model.
	DesiredModelId *string

	// The new name for this Provisioned Throughput.
	DesiredProvisionedModelName *string

	noSmithyDocumentSerde
}

type UpdateProvisionedModelThroughputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProvisionedModelThroughputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProvisionedModelThroughput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProvisionedModelThroughput{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProvisionedModelThroughput"); err != nil {
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
	if err = addOpUpdateProvisionedModelThroughputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisionedModelThroughput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProvisionedModelThroughput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProvisionedModelThroughput",
	}
}