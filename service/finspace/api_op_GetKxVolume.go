// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the information about the volume.
func (c *Client) GetKxVolume(ctx context.Context, params *GetKxVolumeInput, optFns ...func(*Options)) (*GetKxVolumeOutput, error) {
	if params == nil {
		params = &GetKxVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKxVolume", params, optFns, c.addOperationGetKxVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKxVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKxVolumeInput struct {

	// A unique identifier for the kdb environment, whose clusters can attach to the
	// volume.
	//
	// This member is required.
	EnvironmentId *string

	// A unique identifier for the volume.
	//
	// This member is required.
	VolumeName *string

	noSmithyDocumentSerde
}

type GetKxVolumeOutput struct {

	//  A list of cluster identifiers that a volume is attached to.
	AttachedClusters []types.KxAttachedCluster

	// The identifier of the availability zones.
	AvailabilityZoneIds []string

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	AzMode types.KxAzMode

	//  The timestamp at which the volume was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	//  A description of the volume.
	Description *string

	// A unique identifier for the kdb environment, whose clusters can attach to the
	// volume.
	EnvironmentId *string

	// The last time that the volume was updated in FinSpace. The value is determined
	// as epoch time in milliseconds. For example, the value for Monday, November 1,
	// 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	//  Specifies the configuration for the Network attached storage (NAS_1) file
	// system volume.
	Nas1Configuration *types.KxNAS1Configuration

	// The status of volume creation.
	//
	//   - CREATING – The volume creation is in progress.
	//
	//   - CREATE_FAILED – The volume creation has failed.
	//
	//   - ACTIVE – The volume is active.
	//
	//   - UPDATING – The volume is in the process of being updated.
	//
	//   - UPDATE_FAILED – The update action failed.
	//
	//   - UPDATED – The volume is successfully updated.
	//
	//   - DELETING – The volume is in the process of being deleted.
	//
	//   - DELETE_FAILED – The system failed to delete the volume.
	//
	//   - DELETED – The volume is successfully deleted.
	Status types.KxVolumeStatus

	// The error message when a failed state occurs.
	StatusReason *string

	//  The ARN identifier of the volume.
	VolumeArn *string

	//  A unique identifier for the volume.
	VolumeName *string

	//  The type of file system volume. Currently, FinSpace only supports NAS_1 volume
	// type.
	VolumeType types.KxVolumeType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetKxVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetKxVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetKxVolume{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetKxVolume"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetKxVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKxVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetKxVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetKxVolume",
	}
}