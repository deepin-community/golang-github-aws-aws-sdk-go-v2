// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the status of a classification job.
func (c *Client) UpdateClassificationJob(ctx context.Context, params *UpdateClassificationJobInput, optFns ...func(*Options)) (*UpdateClassificationJobOutput, error) {
	if params == nil {
		params = &UpdateClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClassificationJob", params, optFns, c.addOperationUpdateClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClassificationJobInput struct {

	// The unique identifier for the classification job.
	//
	// This member is required.
	JobId *string

	// The new status for the job. Valid values are:
	//
	//   - CANCELLED - Stops the job permanently and cancels it. This value is valid
	//   only if the job's current status is IDLE, PAUSED, RUNNING, or USER_PAUSED.
	//
	// If you specify this value and the job's current status is RUNNING, Amazon Macie
	//   immediately begins to stop all processing tasks for the job. You can't resume or
	//   restart a job after you cancel it.
	//
	//   - RUNNING - Resumes the job. This value is valid only if the job's current
	//   status is USER_PAUSED.
	//
	// If you paused the job while it was actively running and you specify this value
	//   less than 30 days after you paused the job, Macie immediately resumes processing
	//   from the point where you paused the job. Otherwise, Macie resumes the job
	//   according to the schedule and other settings for the job.
	//
	//   - USER_PAUSED - Pauses the job temporarily. This value is valid only if the
	//   job's current status is IDLE, PAUSED, or RUNNING. If you specify this value and
	//   the job's current status is RUNNING, Macie immediately begins to pause all
	//   processing tasks for the job.
	//
	// If you pause a one-time job and you don't resume it within 30 days, the job
	//   expires and Macie cancels the job. If you pause a recurring job when its status
	//   is RUNNING and you don't resume it within 30 days, the job run expires and Macie
	//   cancels the run. To check the expiration date, refer to the
	//   UserPausedDetails.jobExpiresAt property.
	//
	// This member is required.
	JobStatus types.JobStatus

	noSmithyDocumentSerde
}

type UpdateClassificationJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateClassificationJob"); err != nil {
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
	if err = addOpUpdateClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClassificationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateClassificationJob",
	}
}