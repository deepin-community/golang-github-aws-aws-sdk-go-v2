// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels an existing Amazon FSx for Lustre data repository task if that task is
// in either the PENDING or EXECUTING state. When you cancel am export task,
// Amazon FSx does the following.
//
//   - Any files that FSx has already exported are not reverted.
//
//   - FSx continues to export any files that are in-flight when the cancel
//     operation is received.
//
//   - FSx does not export any files that have not yet been exported.
//
// For a release task, Amazon FSx will stop releasing files upon cancellation. Any
// files that have already been released will remain in the released state.
func (c *Client) CancelDataRepositoryTask(ctx context.Context, params *CancelDataRepositoryTaskInput, optFns ...func(*Options)) (*CancelDataRepositoryTaskOutput, error) {
	if params == nil {
		params = &CancelDataRepositoryTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelDataRepositoryTask", params, optFns, c.addOperationCancelDataRepositoryTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelDataRepositoryTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Cancels a data repository task.
type CancelDataRepositoryTaskInput struct {

	// Specifies the data repository task to cancel.
	//
	// This member is required.
	TaskId *string

	noSmithyDocumentSerde
}

type CancelDataRepositoryTaskOutput struct {

	// The lifecycle status of the data repository task, as follows:
	//
	//   - PENDING - Amazon FSx has not started the task.
	//
	//   - EXECUTING - Amazon FSx is processing the task.
	//
	//   - FAILED - Amazon FSx was not able to complete the task. For example, there
	//   may be files the task failed to process. The DataRepositoryTaskFailureDetailsproperty provides more
	//   information about task failures.
	//
	//   - SUCCEEDED - FSx completed the task successfully.
	//
	//   - CANCELED - Amazon FSx canceled the task and it did not complete.
	//
	//   - CANCELING - FSx is in process of canceling the task.
	Lifecycle types.DataRepositoryTaskLifecycle

	// The ID of the task being canceled.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelDataRepositoryTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelDataRepositoryTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelDataRepositoryTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelDataRepositoryTask"); err != nil {
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
	if err = addOpCancelDataRepositoryTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelDataRepositoryTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelDataRepositoryTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelDataRepositoryTask",
	}
}