// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a multicast group.
func (c *Client) GetMulticastGroup(ctx context.Context, params *GetMulticastGroupInput, optFns ...func(*Options)) (*GetMulticastGroupOutput, error) {
	if params == nil {
		params = &GetMulticastGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMulticastGroup", params, optFns, c.addOperationGetMulticastGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMulticastGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMulticastGroupInput struct {

	// The ID of the multicast group.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetMulticastGroupOutput struct {

	// The arn of the multicast group.
	Arn *string

	// Created at timestamp for the resource.
	CreatedAt *time.Time

	// The description of the new resource.
	Description *string

	// The ID of the multicast group.
	Id *string

	// The LoRaWAN information that is to be returned from getting multicast group
	// information.
	LoRaWAN *types.LoRaWANMulticastGet

	// The name of the multicast group.
	Name *string

	// The status of the multicast group.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMulticastGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMulticastGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMulticastGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMulticastGroup"); err != nil {
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
	if err = addOpGetMulticastGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMulticastGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMulticastGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMulticastGroup",
	}
}