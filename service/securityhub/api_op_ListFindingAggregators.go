// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If finding aggregation is enabled, then ListFindingAggregators returns the ARN
// of the finding aggregator. You can run this operation from any Region.
func (c *Client) ListFindingAggregators(ctx context.Context, params *ListFindingAggregatorsInput, optFns ...func(*Options)) (*ListFindingAggregatorsOutput, error) {
	if params == nil {
		params = &ListFindingAggregatorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindingAggregators", params, optFns, c.addOperationListFindingAggregatorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingAggregatorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingAggregatorsInput struct {

	// The maximum number of results to return. This operation currently only returns
	// a single result.
	MaxResults *int32

	// The token returned with the previous set of results. Identifies the next set of
	// results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFindingAggregatorsOutput struct {

	// The list of finding aggregators. This operation currently only returns a single
	// result.
	FindingAggregators []types.FindingAggregator

	// If there are more results, this is the token to provide in the next call to
	// ListFindingAggregators .
	//
	// This operation currently only returns a single result.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFindingAggregatorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindingAggregators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindingAggregators{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFindingAggregators"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindingAggregators(options.Region), middleware.Before); err != nil {
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

// ListFindingAggregatorsPaginatorOptions is the paginator options for
// ListFindingAggregators
type ListFindingAggregatorsPaginatorOptions struct {
	// The maximum number of results to return. This operation currently only returns
	// a single result.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingAggregatorsPaginator is a paginator for ListFindingAggregators
type ListFindingAggregatorsPaginator struct {
	options   ListFindingAggregatorsPaginatorOptions
	client    ListFindingAggregatorsAPIClient
	params    *ListFindingAggregatorsInput
	nextToken *string
	firstPage bool
}

// NewListFindingAggregatorsPaginator returns a new ListFindingAggregatorsPaginator
func NewListFindingAggregatorsPaginator(client ListFindingAggregatorsAPIClient, params *ListFindingAggregatorsInput, optFns ...func(*ListFindingAggregatorsPaginatorOptions)) *ListFindingAggregatorsPaginator {
	if params == nil {
		params = &ListFindingAggregatorsInput{}
	}

	options := ListFindingAggregatorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFindingAggregatorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingAggregatorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFindingAggregators page.
func (p *ListFindingAggregatorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingAggregatorsOutput, error) {
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
	result, err := p.client.ListFindingAggregators(ctx, &params, optFns...)
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

// ListFindingAggregatorsAPIClient is a client that implements the
// ListFindingAggregators operation.
type ListFindingAggregatorsAPIClient interface {
	ListFindingAggregators(context.Context, *ListFindingAggregatorsInput, ...func(*Options)) (*ListFindingAggregatorsOutput, error)
}

var _ ListFindingAggregatorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListFindingAggregators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFindingAggregators",
	}
}