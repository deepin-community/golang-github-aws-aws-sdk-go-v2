// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The UpdatePipelineStatus operation pauses or reactivates a pipeline, so that
// the pipeline stops or restarts the processing of jobs.
//
// Changing the pipeline status is useful if you want to cancel one or more jobs.
// You can't cancel jobs after Elastic Transcoder has started processing them; if
// you pause the pipeline to which you submitted the jobs, you have more time to
// get the job IDs for the jobs that you want to cancel, and to send a CancelJobrequest.
func (c *Client) UpdatePipelineStatus(ctx context.Context, params *UpdatePipelineStatusInput, optFns ...func(*Options)) (*UpdatePipelineStatusOutput, error) {
	if params == nil {
		params = &UpdatePipelineStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePipelineStatus", params, optFns, c.addOperationUpdatePipelineStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePipelineStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdatePipelineStatusRequest structure.
type UpdatePipelineStatusInput struct {

	// The identifier of the pipeline to update.
	//
	// This member is required.
	Id *string

	// The desired status of the pipeline:
	//
	//   - Active : The pipeline is processing jobs.
	//
	//   - Paused : The pipeline is not currently processing jobs.
	//
	// This member is required.
	Status *string

	noSmithyDocumentSerde
}

// When you update status for a pipeline, Elastic Transcoder returns the values
// that you specified in the request.
type UpdatePipelineStatusOutput struct {

	// A section of the response body that provides information about the pipeline.
	Pipeline *types.Pipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePipelineStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePipelineStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePipelineStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePipelineStatus"); err != nil {
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
	if err = addOpUpdatePipelineStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipelineStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePipelineStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePipelineStatus",
	}
}