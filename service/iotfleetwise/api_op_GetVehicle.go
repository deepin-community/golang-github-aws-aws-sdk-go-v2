// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a vehicle.
func (c *Client) GetVehicle(ctx context.Context, params *GetVehicleInput, optFns ...func(*Options)) (*GetVehicleOutput, error) {
	if params == nil {
		params = &GetVehicleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVehicle", params, optFns, c.addOperationGetVehicleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVehicleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVehicleInput struct {

	//  The ID of the vehicle to retrieve information about.
	//
	// This member is required.
	VehicleName *string

	noSmithyDocumentSerde
}

type GetVehicleOutput struct {

	//  The Amazon Resource Name (ARN) of the vehicle to retrieve information about.
	Arn *string

	// Static information about a vehicle in a key-value pair. For example:
	//
	// "engineType" : "1.3 L R2"
	Attributes map[string]string

	//  The time the vehicle was created in seconds since epoch (January 1, 1970 at
	// midnight UTC time).
	CreationTime *time.Time

	//  The ARN of a decoder manifest associated with the vehicle.
	DecoderManifestArn *string

	//  The time the vehicle was last updated in seconds since epoch (January 1, 1970
	// at midnight UTC time).
	LastModificationTime *time.Time

	//  The ARN of a vehicle model (model manifest) associated with the vehicle.
	ModelManifestArn *string

	// The ID of the vehicle.
	VehicleName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVehicleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetVehicle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetVehicle{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetVehicle"); err != nil {
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
	if err = addOpGetVehicleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVehicle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVehicle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetVehicle",
	}
}