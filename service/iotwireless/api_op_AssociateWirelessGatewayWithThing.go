// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a wireless gateway with a thing.
func (c *Client) AssociateWirelessGatewayWithThing(ctx context.Context, params *AssociateWirelessGatewayWithThingInput, optFns ...func(*Options)) (*AssociateWirelessGatewayWithThingOutput, error) {
	if params == nil {
		params = &AssociateWirelessGatewayWithThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWirelessGatewayWithThing", params, optFns, c.addOperationAssociateWirelessGatewayWithThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWirelessGatewayWithThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWirelessGatewayWithThingInput struct {

	// The ID of the resource to update.
	//
	// This member is required.
	Id *string

	// The ARN of the thing to associate with the wireless gateway.
	//
	// This member is required.
	ThingArn *string

	noSmithyDocumentSerde
}

type AssociateWirelessGatewayWithThingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateWirelessGatewayWithThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateWirelessGatewayWithThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateWirelessGatewayWithThing{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateWirelessGatewayWithThing"); err != nil {
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
	if err = addOpAssociateWirelessGatewayWithThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWirelessGatewayWithThing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateWirelessGatewayWithThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateWirelessGatewayWithThing",
	}
}