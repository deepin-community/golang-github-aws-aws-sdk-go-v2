// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an event type. An event is a business activity that is
// evaluated for fraud risk. With Amazon Fraud Detector, you generate fraud
// predictions for events. An event type defines the structure for an event sent to
// Amazon Fraud Detector. This includes the variables sent as part of the event,
// the entity performing the event (such as a customer), and the labels that
// classify the event. Example event types include online payment transactions,
// account registrations, and authentications.
func (c *Client) PutEventType(ctx context.Context, params *PutEventTypeInput, optFns ...func(*Options)) (*PutEventTypeOutput, error) {
	if params == nil {
		params = &PutEventTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEventType", params, optFns, c.addOperationPutEventTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventTypeInput struct {

	// The entity type for the event type. Example entity types: customer, merchant,
	// account.
	//
	// This member is required.
	EntityTypes []string

	// The event type variables.
	//
	// This member is required.
	EventVariables []string

	// The name.
	//
	// This member is required.
	Name *string

	// The description of the event type.
	Description *string

	// Specifies if ingestion is enabled or disabled.
	EventIngestion types.EventIngestion

	// Enables or disables event orchestration. If enabled, you can send event
	// predictions to select AWS services for downstream processing of the events.
	EventOrchestration *types.EventOrchestration

	// The event type labels.
	Labels []string

	// A collection of key and value pairs.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutEventTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEventTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEventType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEventType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutEventType"); err != nil {
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
	if err = addOpPutEventTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEventType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEventType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutEventType",
	}
}