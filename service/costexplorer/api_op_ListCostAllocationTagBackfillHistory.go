// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of your historical cost allocation tag backfill requests.
func (c *Client) ListCostAllocationTagBackfillHistory(ctx context.Context, params *ListCostAllocationTagBackfillHistoryInput, optFns ...func(*Options)) (*ListCostAllocationTagBackfillHistoryOutput, error) {
	if params == nil {
		params = &ListCostAllocationTagBackfillHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCostAllocationTagBackfillHistory", params, optFns, c.addOperationListCostAllocationTagBackfillHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCostAllocationTagBackfillHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCostAllocationTagBackfillHistoryInput struct {

	//  The maximum number of objects that are returned for this request.
	MaxResults *int32

	//  The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCostAllocationTagBackfillHistoryOutput struct {

	//  The list of historical cost allocation tag backfill requests.
	BackfillRequests []types.CostAllocationTagBackfillRequest

	//  The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCostAllocationTagBackfillHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCostAllocationTagBackfillHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCostAllocationTagBackfillHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCostAllocationTagBackfillHistory"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCostAllocationTagBackfillHistory(options.Region), middleware.Before); err != nil {
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

// ListCostAllocationTagBackfillHistoryPaginatorOptions is the paginator options
// for ListCostAllocationTagBackfillHistory
type ListCostAllocationTagBackfillHistoryPaginatorOptions struct {
	//  The maximum number of objects that are returned for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCostAllocationTagBackfillHistoryPaginator is a paginator for
// ListCostAllocationTagBackfillHistory
type ListCostAllocationTagBackfillHistoryPaginator struct {
	options   ListCostAllocationTagBackfillHistoryPaginatorOptions
	client    ListCostAllocationTagBackfillHistoryAPIClient
	params    *ListCostAllocationTagBackfillHistoryInput
	nextToken *string
	firstPage bool
}

// NewListCostAllocationTagBackfillHistoryPaginator returns a new
// ListCostAllocationTagBackfillHistoryPaginator
func NewListCostAllocationTagBackfillHistoryPaginator(client ListCostAllocationTagBackfillHistoryAPIClient, params *ListCostAllocationTagBackfillHistoryInput, optFns ...func(*ListCostAllocationTagBackfillHistoryPaginatorOptions)) *ListCostAllocationTagBackfillHistoryPaginator {
	if params == nil {
		params = &ListCostAllocationTagBackfillHistoryInput{}
	}

	options := ListCostAllocationTagBackfillHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCostAllocationTagBackfillHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCostAllocationTagBackfillHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCostAllocationTagBackfillHistory page.
func (p *ListCostAllocationTagBackfillHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCostAllocationTagBackfillHistoryOutput, error) {
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
	result, err := p.client.ListCostAllocationTagBackfillHistory(ctx, &params, optFns...)
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

// ListCostAllocationTagBackfillHistoryAPIClient is a client that implements the
// ListCostAllocationTagBackfillHistory operation.
type ListCostAllocationTagBackfillHistoryAPIClient interface {
	ListCostAllocationTagBackfillHistory(context.Context, *ListCostAllocationTagBackfillHistoryInput, ...func(*Options)) (*ListCostAllocationTagBackfillHistoryOutput, error)
}

var _ ListCostAllocationTagBackfillHistoryAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCostAllocationTagBackfillHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCostAllocationTagBackfillHistory",
	}
}