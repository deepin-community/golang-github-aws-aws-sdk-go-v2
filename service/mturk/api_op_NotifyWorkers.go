// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The NotifyWorkers operation sends an email to one or more Workers that you
//
// specify with the Worker ID. You can specify up to 100 Worker IDs to send the
// same message with a single call to the NotifyWorkers operation. The
// NotifyWorkers operation will send a notification email to a Worker only if you
// have previously approved or rejected work from the Worker.
func (c *Client) NotifyWorkers(ctx context.Context, params *NotifyWorkersInput, optFns ...func(*Options)) (*NotifyWorkersOutput, error) {
	if params == nil {
		params = &NotifyWorkersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyWorkers", params, optFns, c.addOperationNotifyWorkersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyWorkersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyWorkersInput struct {

	// The text of the email message to send. Can include up to 4,096 characters
	//
	// This member is required.
	MessageText *string

	// The subject line of the email message to send. Can include up to 200 characters.
	//
	// This member is required.
	Subject *string

	// A list of Worker IDs you wish to notify. You can notify upto 100 Workers at a
	// time.
	//
	// This member is required.
	WorkerIds []string

	noSmithyDocumentSerde
}

type NotifyWorkersOutput struct {

	//  When MTurk sends notifications to the list of Workers, it returns back any
	// failures it encounters in this list of NotifyWorkersFailureStatus objects.
	NotifyWorkersFailureStatuses []types.NotifyWorkersFailureStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyWorkersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyWorkers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyWorkers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "NotifyWorkers"); err != nil {
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
	if err = addOpNotifyWorkersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyWorkers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNotifyWorkers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NotifyWorkers",
	}
}