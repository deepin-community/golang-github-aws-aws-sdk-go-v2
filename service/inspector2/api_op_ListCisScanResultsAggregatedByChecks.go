// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists scan results aggregated by checks.
func (c *Client) ListCisScanResultsAggregatedByChecks(ctx context.Context, params *ListCisScanResultsAggregatedByChecksInput, optFns ...func(*Options)) (*ListCisScanResultsAggregatedByChecksOutput, error) {
	if params == nil {
		params = &ListCisScanResultsAggregatedByChecksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCisScanResultsAggregatedByChecks", params, optFns, c.addOperationListCisScanResultsAggregatedByChecksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCisScanResultsAggregatedByChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCisScanResultsAggregatedByChecksInput struct {

	// The scan ARN.
	//
	// This member is required.
	ScanArn *string

	// The filter criteria.
	FilterCriteria *types.CisScanResultsAggregatedByChecksFilterCriteria

	// The maximum number of scan results aggregated by checks to be returned in a
	// single page of results.
	MaxResults *int32

	// The pagination token from a previous request that's used to retrieve the next
	// page of results.
	NextToken *string

	// The sort by order.
	SortBy types.CisScanResultsAggregatedByChecksSortBy

	// The sort order.
	SortOrder types.CisSortOrder

	noSmithyDocumentSerde
}

type ListCisScanResultsAggregatedByChecksOutput struct {

	// The check aggregations.
	CheckAggregations []types.CisCheckAggregation

	// The pagination token from a previous request that's used to retrieve the next
	// page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCisScanResultsAggregatedByChecksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCisScanResultsAggregatedByChecks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCisScanResultsAggregatedByChecks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCisScanResultsAggregatedByChecks"); err != nil {
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
	if err = addOpListCisScanResultsAggregatedByChecksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCisScanResultsAggregatedByChecks(options.Region), middleware.Before); err != nil {
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

// ListCisScanResultsAggregatedByChecksPaginatorOptions is the paginator options
// for ListCisScanResultsAggregatedByChecks
type ListCisScanResultsAggregatedByChecksPaginatorOptions struct {
	// The maximum number of scan results aggregated by checks to be returned in a
	// single page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCisScanResultsAggregatedByChecksPaginator is a paginator for
// ListCisScanResultsAggregatedByChecks
type ListCisScanResultsAggregatedByChecksPaginator struct {
	options   ListCisScanResultsAggregatedByChecksPaginatorOptions
	client    ListCisScanResultsAggregatedByChecksAPIClient
	params    *ListCisScanResultsAggregatedByChecksInput
	nextToken *string
	firstPage bool
}

// NewListCisScanResultsAggregatedByChecksPaginator returns a new
// ListCisScanResultsAggregatedByChecksPaginator
func NewListCisScanResultsAggregatedByChecksPaginator(client ListCisScanResultsAggregatedByChecksAPIClient, params *ListCisScanResultsAggregatedByChecksInput, optFns ...func(*ListCisScanResultsAggregatedByChecksPaginatorOptions)) *ListCisScanResultsAggregatedByChecksPaginator {
	if params == nil {
		params = &ListCisScanResultsAggregatedByChecksInput{}
	}

	options := ListCisScanResultsAggregatedByChecksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCisScanResultsAggregatedByChecksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCisScanResultsAggregatedByChecksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCisScanResultsAggregatedByChecks page.
func (p *ListCisScanResultsAggregatedByChecksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCisScanResultsAggregatedByChecksOutput, error) {
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
	result, err := p.client.ListCisScanResultsAggregatedByChecks(ctx, &params, optFns...)
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

// ListCisScanResultsAggregatedByChecksAPIClient is a client that implements the
// ListCisScanResultsAggregatedByChecks operation.
type ListCisScanResultsAggregatedByChecksAPIClient interface {
	ListCisScanResultsAggregatedByChecks(context.Context, *ListCisScanResultsAggregatedByChecksInput, ...func(*Options)) (*ListCisScanResultsAggregatedByChecksOutput, error)
}

var _ ListCisScanResultsAggregatedByChecksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCisScanResultsAggregatedByChecks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCisScanResultsAggregatedByChecks",
	}
}