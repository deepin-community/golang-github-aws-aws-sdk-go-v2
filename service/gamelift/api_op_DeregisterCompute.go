// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Removes a compute resource from an Amazon GameLift Anywhere fleet or container
// fleet. Deregistered computes can no longer host game sessions through Amazon
// GameLift.
//
// For an Anywhere fleet or a container fleet that's running the Amazon GameLift
// Agent, the Agent handles all compute registry tasks for you. For an Anywhere
// fleet that doesn't use the Agent, call this operation to deregister fleet
// computes.
//
// To deregister a compute, call this operation from the compute that's being
// deregistered and specify the compute name and the fleet ID.
func (c *Client) DeregisterCompute(ctx context.Context, params *DeregisterComputeInput, optFns ...func(*Options)) (*DeregisterComputeOutput, error) {
	if params == nil {
		params = &DeregisterComputeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterCompute", params, optFns, c.addOperationDeregisterComputeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterComputeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterComputeInput struct {

	// The unique identifier of the compute resource to deregister. For an Anywhere
	// fleet compute, use the registered compute name. For a container fleet, use the
	// compute name (for example,
	// a123b456c789012d3e4567f8a901b23c/1a234b56-7cd8-9e0f-a1b2-c34d567ef8a9 ) or the
	// compute ARN.
	//
	// This member is required.
	ComputeName *string

	// A unique identifier for the fleet the compute resource is currently registered
	// to.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type DeregisterComputeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterComputeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterCompute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterCompute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterCompute"); err != nil {
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
	if err = addOpDeregisterComputeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterCompute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterCompute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterCompute",
	}
}