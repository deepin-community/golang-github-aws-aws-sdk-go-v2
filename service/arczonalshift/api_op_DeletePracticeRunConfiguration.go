// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the practice run configuration for a resource. Before you can delete a
// practice run configuration for a resource., you must disable zonal autoshift for
// the resource. Practice runs must be configured for zonal autoshift to be
// enabled.
func (c *Client) DeletePracticeRunConfiguration(ctx context.Context, params *DeletePracticeRunConfigurationInput, optFns ...func(*Options)) (*DeletePracticeRunConfigurationOutput, error) {
	if params == nil {
		params = &DeletePracticeRunConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePracticeRunConfiguration", params, optFns, c.addOperationDeletePracticeRunConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePracticeRunConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePracticeRunConfigurationInput struct {

	// The identifier for the resource that you want to delete the practice run
	// configuration for. The identifier is the Amazon Resource Name (ARN) for the
	// resource.
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type DeletePracticeRunConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the resource that you deleted the practice
	// run for.
	//
	// This member is required.
	Arn *string

	// The name of the resource that you deleted the practice run for.
	//
	// This member is required.
	Name *string

	// The status of zonal autoshift for the resource.
	//
	// This member is required.
	ZonalAutoshiftStatus types.ZonalAutoshiftStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePracticeRunConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeletePracticeRunConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePracticeRunConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePracticeRunConfiguration"); err != nil {
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
	if err = addOpDeletePracticeRunConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePracticeRunConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePracticeRunConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePracticeRunConfiguration",
	}
}