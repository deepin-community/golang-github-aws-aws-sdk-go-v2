// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the details for an existing device. To remove information for any of
// the parameters, specify an empty string.
func (c *Client) UpdateDevice(ctx context.Context, params *UpdateDeviceInput, optFns ...func(*Options)) (*UpdateDeviceOutput, error) {
	if params == nil {
		params = &UpdateDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDevice", params, optFns, c.addOperationUpdateDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDeviceInput struct {

	// The ID of the device.
	//
	// This member is required.
	DeviceId *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Web Services location of the device, if applicable. For an
	// on-premises device, you can omit this parameter.
	AWSLocation *types.AWSLocation

	// A description of the device.
	//
	// Constraints: Maximum length of 256 characters.
	Description *string

	// Describes a location.
	Location *types.Location

	// The model of the device.
	//
	// Constraints: Maximum length of 128 characters.
	Model *string

	// The serial number of the device.
	//
	// Constraints: Maximum length of 128 characters.
	SerialNumber *string

	// The ID of the site.
	SiteId *string

	// The type of the device.
	Type *string

	// The vendor of the device.
	//
	// Constraints: Maximum length of 128 characters.
	Vendor *string

	noSmithyDocumentSerde
}

type UpdateDeviceOutput struct {

	// Information about the device.
	Device *types.Device

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDevice"); err != nil {
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
	if err = addOpUpdateDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDevice",
	}
}