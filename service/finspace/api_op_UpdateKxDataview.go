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

//	Updates the specified dataview. The dataviews get automatically updated when
//
// any new changesets are ingested. Each update of the dataview creates a new
// version, including changeset details and cache configurations
func (c *Client) UpdateKxDataview(ctx context.Context, params *UpdateKxDataviewInput, optFns ...func(*Options)) (*UpdateKxDataviewOutput, error) {
	if params == nil {
		params = &UpdateKxDataviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKxDataview", params, optFns, c.addOperationUpdateKxDataviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKxDataviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKxDataviewInput struct {

	// A token that ensures idempotency. This token expires in 10 minutes.
	//
	// This member is required.
	ClientToken *string

	//  The name of the database.
	//
	// This member is required.
	DatabaseName *string

	// The name of the dataview that you want to update.
	//
	// This member is required.
	DataviewName *string

	// A unique identifier for the kdb environment, where you want to update the
	// dataview.
	//
	// This member is required.
	EnvironmentId *string

	// A unique identifier for the changeset.
	ChangesetId *string

	//  The description for a dataview.
	Description *string

	//  The configuration that contains the database path of the data that you want to
	// place on each selected volume. Each segment must have a unique database path for
	// each volume. If you do not explicitly specify any database path for a volume,
	// they are accessible from the cluster through the default S3/object store
	// segment.
	SegmentConfigurations []types.KxDataviewSegmentConfiguration

	noSmithyDocumentSerde
}

type UpdateKxDataviewOutput struct {

	//  The current active changeset versions of the database on the given dataview.
	ActiveVersions []types.KxDataviewActiveVersion

	// The option to specify whether you want to apply all the future additions and
	// corrections automatically to the dataview when new changesets are ingested. The
	// default value is false.
	AutoUpdate bool

	//  The identifier of the availability zones.
	AvailabilityZoneId *string

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	AzMode types.KxAzMode

	// A unique identifier for the changeset.
	ChangesetId *string

	//  The timestamp at which the dataview was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	//  The name of the database.
	DatabaseName *string

	//  The name of the database under which the dataview was created.
	DataviewName *string

	// A description of the dataview.
	Description *string

	// A unique identifier for the kdb environment, where you want to update the
	// dataview.
	EnvironmentId *string

	//  The last time that the dataview was updated in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	// Returns True if the dataview is created as writeable and False otherwise.
	ReadWrite bool

	//  The configuration that contains the database path of the data that you want to
	// place on each selected volume. Each segment must have a unique database path for
	// each volume. If you do not explicitly specify any database path for a volume,
	// they are accessible from the cluster through the default S3/object store
	// segment.
	SegmentConfigurations []types.KxDataviewSegmentConfiguration

	//  The status of dataview creation.
	//
	//   - CREATING – The dataview creation is in progress.
	//
	//   - UPDATING – The dataview is in the process of being updated.
	//
	//   - ACTIVE – The dataview is active.
	Status types.KxDataviewStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKxDataviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateKxDataview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateKxDataview{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateKxDataview"); err != nil {
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
	if err = addIdempotencyToken_opUpdateKxDataviewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateKxDataviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKxDataview(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateKxDataview struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateKxDataview) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateKxDataview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateKxDataviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateKxDataviewInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateKxDataviewMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateKxDataview{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateKxDataview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateKxDataview",
	}
}