// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a listener for a custom routing accelerator.
func (c *Client) UpdateCustomRoutingListener(ctx context.Context, params *UpdateCustomRoutingListenerInput, optFns ...func(*Options)) (*UpdateCustomRoutingListenerOutput, error) {
	if params == nil {
		params = &UpdateCustomRoutingListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCustomRoutingListener", params, optFns, c.addOperationUpdateCustomRoutingListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCustomRoutingListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCustomRoutingListenerInput struct {

	// The Amazon Resource Name (ARN) of the listener to update.
	//
	// This member is required.
	ListenerArn *string

	// The updated port range to support for connections from clients to your
	// accelerator. If you remove ports that are currently being used by a subnet
	// endpoint, the call fails.
	//
	// Separately, you set port ranges for endpoints. For more information, see [About endpoints for custom routing accelerators].
	//
	// [About endpoints for custom routing accelerators]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-custom-routing-endpoints.html
	//
	// This member is required.
	PortRanges []types.PortRange

	noSmithyDocumentSerde
}

type UpdateCustomRoutingListenerOutput struct {

	// Information for the updated listener for a custom routing accelerator.
	Listener *types.CustomRoutingListener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCustomRoutingListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCustomRoutingListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCustomRoutingListener{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCustomRoutingListener"); err != nil {
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
	if err = addOpUpdateCustomRoutingListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCustomRoutingListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCustomRoutingListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCustomRoutingListener",
	}
}