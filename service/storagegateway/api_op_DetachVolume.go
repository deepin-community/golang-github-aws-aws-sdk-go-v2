// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnects a volume from an iSCSI connection and then detaches the volume from
// the specified gateway. Detaching and attaching a volume enables you to recover
// your data from one gateway to a different gateway without creating a snapshot.
// It also makes it easier to move your volumes from an on-premises gateway to a
// gateway hosted on an Amazon EC2 instance. This operation is only supported in
// the volume gateway type.
func (c *Client) DetachVolume(ctx context.Context, params *DetachVolumeInput, optFns ...func(*Options)) (*DetachVolumeOutput, error) {
	if params == nil {
		params = &DetachVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachVolume", params, optFns, c.addOperationDetachVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// AttachVolumeInput
type DetachVolumeInput struct {

	// The Amazon Resource Name (ARN) of the volume to detach from the gateway.
	//
	// This member is required.
	VolumeARN *string

	// Set to true to forcibly remove the iSCSI connection of the target volume and
	// detach the volume. The default is false . If this value is set to false , you
	// must manually disconnect the iSCSI connection from the target volume.
	//
	// Valid Values: true | false
	ForceDetach *bool

	noSmithyDocumentSerde
}

// AttachVolumeOutput
type DetachVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the volume that was detached.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachVolume{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetachVolume"); err != nil {
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
	if err = addOpDetachVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetachVolume",
	}
}