// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a query if the query is not in a terminated state, such as CANCELLED ,
// FAILED , TIMED_OUT , or FINISHED . You must specify an ARN value for
// EventDataStore . The ID of the query that you want to cancel is also required.
// When you run CancelQuery , the query status might show as CANCELLED even if the
// operation is not yet finished.
func (c *Client) CancelQuery(ctx context.Context, params *CancelQueryInput, optFns ...func(*Options)) (*CancelQueryOutput, error) {
	if params == nil {
		params = &CancelQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelQuery", params, optFns, c.addOperationCancelQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelQueryInput struct {

	// The ID of the query that you want to cancel. The QueryId comes from the
	// response of a StartQuery operation.
	//
	// This member is required.
	QueryId *string

	// The ARN (or the ID suffix of the ARN) of an event data store on which the
	// specified query is running.
	//
	// Deprecated: EventDataStore is no longer required by CancelQueryRequest
	EventDataStore *string

	noSmithyDocumentSerde
}

type CancelQueryOutput struct {

	// The ID of the canceled query.
	//
	// This member is required.
	QueryId *string

	// Shows the status of a query after a CancelQuery request. Typically, the values
	// shown are either RUNNING or CANCELLED .
	//
	// This member is required.
	QueryStatus types.QueryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelQuery"); err != nil {
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
	if err = addOpCancelQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelQuery",
	}
}