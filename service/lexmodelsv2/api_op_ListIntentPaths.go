// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves summary statistics for a path of intents that users take over
// sessions with your bot. The following fields are required:
//
//   - startDateTime and endDateTime – Define a time range for which you want to
//     retrieve results.
//
//   - intentPath – Define an order of intents for which you want to retrieve
//     metrics. Separate intents in the path with a forward slash. For example,
//     populate the intentPath field with /BookCar/BookHotel to see details about how
//     many times users invoked the BookCar and BookHotel intents in that order.
//
// Use the optional filters field to filter the results.
func (c *Client) ListIntentPaths(ctx context.Context, params *ListIntentPathsInput, optFns ...func(*Options)) (*ListIntentPathsOutput, error) {
	if params == nil {
		params = &ListIntentPathsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIntentPaths", params, optFns, c.addOperationListIntentPathsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIntentPathsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIntentPathsInput struct {

	// The identifier for the bot for which you want to retrieve intent path metrics.
	//
	// This member is required.
	BotId *string

	// The date and time that marks the end of the range of time for which you want to
	// see intent path metrics.
	//
	// This member is required.
	EndDateTime *time.Time

	// The intent path for which you want to retrieve metrics. Use a forward slash to
	// separate intents in the path. For example:
	//
	//   - /BookCar
	//
	//   - /BookCar/BookHotel
	//
	//   - /BookHotel/BookCar
	//
	// This member is required.
	IntentPath *string

	// The date and time that marks the beginning of the range of time for which you
	// want to see intent path metrics.
	//
	// This member is required.
	StartDateTime *time.Time

	// A list of objects, each describes a condition by which you want to filter the
	// results.
	Filters []types.AnalyticsPathFilter

	noSmithyDocumentSerde
}

type ListIntentPathsOutput struct {

	// A list of objects, each of which contains information about a node in the
	// intent path for which you requested metrics.
	NodeSummaries []types.AnalyticsIntentNodeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIntentPathsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIntentPaths{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIntentPaths{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIntentPaths"); err != nil {
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
	if err = addOpListIntentPathsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIntentPaths(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListIntentPaths(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIntentPaths",
	}
}