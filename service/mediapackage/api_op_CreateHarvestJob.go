// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new HarvestJob record.
func (c *Client) CreateHarvestJob(ctx context.Context, params *CreateHarvestJobInput, optFns ...func(*Options)) (*CreateHarvestJobOutput, error) {
	if params == nil {
		params = &CreateHarvestJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHarvestJob", params, optFns, c.addOperationCreateHarvestJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHarvestJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Configuration parameters used to create a new HarvestJob.
type CreateHarvestJobInput struct {

	// The end of the time-window which will be harvested
	//
	// This member is required.
	EndTime *string

	// The ID of the HarvestJob. The ID must be unique within the region and it cannot
	// be changed after the HarvestJob is submitted
	//
	// This member is required.
	Id *string

	// The ID of the OriginEndpoint that the HarvestJob will harvest from. This cannot
	// be changed after the HarvestJob is submitted.
	//
	// This member is required.
	OriginEndpointId *string

	// Configuration parameters for where in an S3 bucket to place the harvested
	// content
	//
	// This member is required.
	S3Destination *types.S3Destination

	// The start of the time-window which will be harvested
	//
	// This member is required.
	StartTime *string

	noSmithyDocumentSerde
}

type CreateHarvestJobOutput struct {

	// The Amazon Resource Name (ARN) assigned to the HarvestJob.
	Arn *string

	// The ID of the Channel that the HarvestJob will harvest from.
	ChannelId *string

	// The date and time the HarvestJob was submitted.
	CreatedAt *string

	// The end of the time-window which will be harvested.
	EndTime *string

	// The ID of the HarvestJob. The ID must be unique within the region and it cannot
	// be changed after the HarvestJob is submitted.
	Id *string

	// The ID of the OriginEndpoint that the HarvestJob will harvest from. This cannot
	// be changed after the HarvestJob is submitted.
	OriginEndpointId *string

	// Configuration parameters for where in an S3 bucket to place the harvested
	// content
	S3Destination *types.S3Destination

	// The start of the time-window which will be harvested.
	StartTime *string

	// The current status of the HarvestJob. Consider setting up a CloudWatch Event to
	// listen for HarvestJobs as they succeed or fail. In the event of failure, the
	// CloudWatch Event will include an explanation of why the HarvestJob failed.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHarvestJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateHarvestJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateHarvestJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHarvestJob"); err != nil {
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
	if err = addOpCreateHarvestJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHarvestJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHarvestJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHarvestJob",
	}
}