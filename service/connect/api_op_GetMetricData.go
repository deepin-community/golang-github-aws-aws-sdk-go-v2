// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets historical metric data from the specified Amazon Connect instance.
//
// For a description of each historical metric, see [Historical Metrics Definitions] in the Amazon Connect
// Administrator Guide.
//
// We recommend using the [GetMetricDataV2] API. It provides more flexibility, features, and the
// ability to query longer time ranges than GetMetricData . Use it to retrieve
// historical agent and contact metrics for the last 3 months, at varying
// intervals. You can also use it to build custom dashboards to measure historical
// queue and agent performance. For example, you can track the number of incoming
// contacts for the last 7 days, with data split by day, to see how contact volume
// changed per day of the week.
//
// [GetMetricDataV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_GetMetricDataV2.html
// [Historical Metrics Definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics-definitions.html
func (c *Client) GetMetricData(ctx context.Context, params *GetMetricDataInput, optFns ...func(*Options)) (*GetMetricDataOutput, error) {
	if params == nil {
		params = &GetMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricData", params, optFns, c.addOperationGetMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricDataInput struct {

	// The timestamp, in UNIX Epoch time format, at which to end the reporting
	// interval for the retrieval of historical metrics data. The time must be
	// specified using an interval of 5 minutes, such as 11:00, 11:05, 11:10, and must
	// be later than the start time timestamp.
	//
	// The time range between the start and end time must be less than 24 hours.
	//
	// This member is required.
	EndTime *time.Time

	// The queues, up to 100, or channels, to use to filter the metrics returned.
	// Metric data is retrieved only for the resources associated with the queues or
	// channels included in the filter. You can include both queue IDs and queue ARNs
	// in the same request. VOICE, CHAT, and TASK channels are supported.
	//
	// RoutingStepExpression is not a valid filter for GetMetricData and we recommend
	// switching to GetMetricDataV2 for more up-to-date features.
	//
	// To filter by Queues , enter the queue ID/ARN, not the name of the queue.
	//
	// This member is required.
	Filters *types.Filters

	// The metrics to retrieve. Specify the name, unit, and statistic for each metric.
	// The following historical metrics are available. For a description of each
	// metric, see [Historical Metrics Definitions]in the Amazon Connect Administrator Guide.
	//
	// This API does not support a contacts incoming metric (there's no
	// CONTACTS_INCOMING metric missing from the documented list).
	//
	// ABANDON_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// AFTER_CONTACT_WORK_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// API_CONTACTS_HANDLED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CALLBACK_CONTACTS_HANDLED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_ABANDONED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_AGENT_HUNG_UP_FIRST Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_CONSULTED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_HANDLED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_HANDLED_INCOMING Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_HANDLED_OUTBOUND Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_HOLD_ABANDONS Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_MISSED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_QUEUED Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_TRANSFERRED_IN Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_TRANSFERRED_IN_FROM_QUEUE Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_TRANSFERRED_OUT Unit: COUNT
	//
	// Statistic: SUM
	//
	// CONTACTS_TRANSFERRED_OUT_FROM_QUEUE Unit: COUNT
	//
	// Statistic: SUM
	//
	// HANDLE_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// HOLD_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// INTERACTION_AND_HOLD_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// INTERACTION_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// OCCUPANCY Unit: PERCENT
	//
	// Statistic: AVG
	//
	// QUEUE_ANSWER_TIME Unit: SECONDS
	//
	// Statistic: AVG
	//
	// QUEUED_TIME Unit: SECONDS
	//
	// Statistic: MAX
	//
	// SERVICE_LEVEL You can include up to 20 SERVICE_LEVEL metrics in a request.
	//
	// Unit: PERCENT
	//
	// Statistic: AVG
	//
	// Threshold: For ThresholdValue , enter any whole number from 1 to 604800
	// (inclusive), in seconds. For Comparison , you must enter LT (for "Less than").
	//
	// [Historical Metrics Definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics-definitions.html
	//
	// This member is required.
	HistoricalMetrics []types.HistoricalMetric

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The timestamp, in UNIX Epoch time format, at which to start the reporting
	// interval for the retrieval of historical metrics data. The time must be
	// specified using a multiple of 5 minutes, such as 10:05, 10:10, 10:15.
	//
	// The start time cannot be earlier than 24 hours before the time of the request.
	// Historical metrics are available only for 24 hours.
	//
	// This member is required.
	StartTime *time.Time

	// The grouping applied to the metrics returned. For example, when results are
	// grouped by queue, the metrics returned are grouped by queue. The values returned
	// apply to the metrics for each queue rather than aggregated for all queues.
	//
	// If no grouping is specified, a summary of metrics for all queues is returned.
	//
	// RoutingStepExpression is not a valid filter for GetMetricData and we recommend
	// switching to GetMetricDataV2 for more up-to-date features.
	Groupings []types.Grouping

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetMetricDataOutput struct {

	// Information about the historical metrics.
	//
	// If no grouping is specified, a summary of metric data is returned.
	MetricResults []types.HistoricalMetricResult

	// If there are additional results, this is the token for the next set of results.
	//
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMetricData"); err != nil {
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
	if err = addOpGetMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricData(options.Region), middleware.Before); err != nil {
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

// GetMetricDataPaginatorOptions is the paginator options for GetMetricData
type GetMetricDataPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMetricDataPaginator is a paginator for GetMetricData
type GetMetricDataPaginator struct {
	options   GetMetricDataPaginatorOptions
	client    GetMetricDataAPIClient
	params    *GetMetricDataInput
	nextToken *string
	firstPage bool
}

// NewGetMetricDataPaginator returns a new GetMetricDataPaginator
func NewGetMetricDataPaginator(client GetMetricDataAPIClient, params *GetMetricDataInput, optFns ...func(*GetMetricDataPaginatorOptions)) *GetMetricDataPaginator {
	if params == nil {
		params = &GetMetricDataInput{}
	}

	options := GetMetricDataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMetricDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMetricDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetMetricData page.
func (p *GetMetricDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMetricDataOutput, error) {
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
	result, err := p.client.GetMetricData(ctx, &params, optFns...)
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

// GetMetricDataAPIClient is a client that implements the GetMetricData operation.
type GetMetricDataAPIClient interface {
	GetMetricData(context.Context, *GetMetricDataInput, ...func(*Options)) (*GetMetricDataOutput, error)
}

var _ GetMetricDataAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMetricData",
	}
}