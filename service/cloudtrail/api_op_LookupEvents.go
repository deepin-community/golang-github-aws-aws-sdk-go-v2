// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Looks up [management events] or [CloudTrail Insights events] that are captured by CloudTrail. You can look up events that
// occurred in a Region within the last 90 days.
//
// LookupEvents returns recent Insights events for trails that enable Insights. To
// view Insights events for an event data store, you can run queries on your
// Insights event data store, and you can also view the Lake dashboard for
// Insights.
//
// Lookup supports the following attributes for management events:
//
//   - Amazon Web Services access key
//
//   - Event ID
//
//   - Event name
//
//   - Event source
//
//   - Read only
//
//   - Resource name
//
//   - Resource type
//
//   - User name
//
// Lookup supports the following attributes for Insights events:
//
//   - Event ID
//
//   - Event name
//
//   - Event source
//
// All attributes are optional. The default number of results returned is 50, with
// a maximum of 50 possible. The response includes a token that you can use to get
// the next page of results.
//
// The rate of lookup requests is limited to two per second, per account, per
// Region. If this limit is exceeded, a throttling error occurs.
//
// [CloudTrail Insights events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-insights-events
// [management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-management-events
func (c *Client) LookupEvents(ctx context.Context, params *LookupEventsInput, optFns ...func(*Options)) (*LookupEventsOutput, error) {
	if params == nil {
		params = &LookupEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LookupEvents", params, optFns, c.addOperationLookupEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LookupEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains a request for LookupEvents.
type LookupEventsInput struct {

	// Specifies that only events that occur before or at the specified time are
	// returned. If the specified end time is before the specified start time, an error
	// is returned.
	EndTime *time.Time

	// Specifies the event category. If you do not specify an event category, events
	// of the category are not returned in the response. For example, if you do not
	// specify insight as the value of EventCategory , no Insights events are returned.
	EventCategory types.EventCategory

	// Contains a list of lookup attributes. Currently the list can contain only one
	// item.
	LookupAttributes []types.LookupAttribute

	// The number of events to return. Possible values are 1 through 50. The default
	// is 50.
	MaxResults *int32

	// The token to use to get the next page of results after a previous API call.
	// This token must be passed in with the same parameters that were specified in the
	// original call. For example, if the original call specified an AttributeKey of
	// 'Username' with a value of 'root', the call with NextToken should include those
	// same parameters.
	NextToken *string

	// Specifies that only events that occur after or at the specified time are
	// returned. If the specified start time is after the specified end time, an error
	// is returned.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Contains a response to a LookupEvents action.
type LookupEventsOutput struct {

	// A list of events returned based on the lookup attributes specified and the
	// CloudTrail event. The events list is sorted by time. The most recent event is
	// listed first.
	Events []types.Event

	// The token to use to get the next page of results after a previous API call. If
	// the token does not appear, there are no more results to return. The token must
	// be passed in with the same parameters as the previous call. For example, if the
	// original call specified an AttributeKey of 'Username' with a value of 'root',
	// the call with NextToken should include those same parameters.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationLookupEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpLookupEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpLookupEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "LookupEvents"); err != nil {
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
	if err = addOpLookupEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLookupEvents(options.Region), middleware.Before); err != nil {
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

// LookupEventsPaginatorOptions is the paginator options for LookupEvents
type LookupEventsPaginatorOptions struct {
	// The number of events to return. Possible values are 1 through 50. The default
	// is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// LookupEventsPaginator is a paginator for LookupEvents
type LookupEventsPaginator struct {
	options   LookupEventsPaginatorOptions
	client    LookupEventsAPIClient
	params    *LookupEventsInput
	nextToken *string
	firstPage bool
}

// NewLookupEventsPaginator returns a new LookupEventsPaginator
func NewLookupEventsPaginator(client LookupEventsAPIClient, params *LookupEventsInput, optFns ...func(*LookupEventsPaginatorOptions)) *LookupEventsPaginator {
	if params == nil {
		params = &LookupEventsInput{}
	}

	options := LookupEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &LookupEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *LookupEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next LookupEvents page.
func (p *LookupEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*LookupEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.LookupEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// LookupEventsAPIClient is a client that implements the LookupEvents operation.
type LookupEventsAPIClient interface {
	LookupEvents(context.Context, *LookupEventsInput, ...func(*Options)) (*LookupEventsOutput, error)
}

var _ LookupEventsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opLookupEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "LookupEvents",
	}
}