// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Batch scheduling policies.
func (c *Client) ListSchedulingPolicies(ctx context.Context, params *ListSchedulingPoliciesInput, optFns ...func(*Options)) (*ListSchedulingPoliciesOutput, error) {
	if params == nil {
		params = &ListSchedulingPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchedulingPolicies", params, optFns, c.addOperationListSchedulingPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchedulingPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ListSchedulingPolicies .
type ListSchedulingPoliciesInput struct {

	// The maximum number of results that's returned by ListSchedulingPolicies in
	// paginated output. When this parameter is used, ListSchedulingPolicies only
	// returns maxResults results in a single page and a nextToken response element.
	// You can see the remaining results of the initial request by sending another
	// ListSchedulingPolicies request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter isn't used, ListSchedulingPolicies
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value that's returned from a previous paginated
	// ListSchedulingPolicies request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. This value is null when
	// there are no more results to return.
	//
	// Treat this token as an opaque identifier that's only used to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSchedulingPoliciesOutput struct {

	// The nextToken value to include in a future ListSchedulingPolicies request. When
	// the results of a ListSchedulingPolicies request exceed maxResults , this value
	// can be used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// A list of scheduling policies that match the request.
	SchedulingPolicies []types.SchedulingPolicyListingDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchedulingPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSchedulingPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSchedulingPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSchedulingPolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchedulingPolicies(options.Region), middleware.Before); err != nil {
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

// ListSchedulingPoliciesPaginatorOptions is the paginator options for
// ListSchedulingPolicies
type ListSchedulingPoliciesPaginatorOptions struct {
	// The maximum number of results that's returned by ListSchedulingPolicies in
	// paginated output. When this parameter is used, ListSchedulingPolicies only
	// returns maxResults results in a single page and a nextToken response element.
	// You can see the remaining results of the initial request by sending another
	// ListSchedulingPolicies request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter isn't used, ListSchedulingPolicies
	// returns up to 100 results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSchedulingPoliciesPaginator is a paginator for ListSchedulingPolicies
type ListSchedulingPoliciesPaginator struct {
	options   ListSchedulingPoliciesPaginatorOptions
	client    ListSchedulingPoliciesAPIClient
	params    *ListSchedulingPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListSchedulingPoliciesPaginator returns a new ListSchedulingPoliciesPaginator
func NewListSchedulingPoliciesPaginator(client ListSchedulingPoliciesAPIClient, params *ListSchedulingPoliciesInput, optFns ...func(*ListSchedulingPoliciesPaginatorOptions)) *ListSchedulingPoliciesPaginator {
	if params == nil {
		params = &ListSchedulingPoliciesInput{}
	}

	options := ListSchedulingPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSchedulingPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSchedulingPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSchedulingPolicies page.
func (p *ListSchedulingPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSchedulingPoliciesOutput, error) {
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
	result, err := p.client.ListSchedulingPolicies(ctx, &params, optFns...)
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

// ListSchedulingPoliciesAPIClient is a client that implements the
// ListSchedulingPolicies operation.
type ListSchedulingPoliciesAPIClient interface {
	ListSchedulingPolicies(context.Context, *ListSchedulingPoliciesInput, ...func(*Options)) (*ListSchedulingPoliciesOutput, error)
}

var _ ListSchedulingPoliciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSchedulingPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSchedulingPolicies",
	}
}