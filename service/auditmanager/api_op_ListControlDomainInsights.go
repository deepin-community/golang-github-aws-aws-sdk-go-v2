// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the latest analytics data for control domains across all of your active
// assessments.
//
// Audit Manager supports the control domains that are provided by Amazon Web
// Services Control Catalog. For information about how to find a list of available
// control domains, see [ListDomains]ListDomains in the Amazon Web Services Control Catalog API
// Reference.
//
// A control domain is listed only if at least one of the controls within that
// domain collected evidence on the lastUpdated date of controlDomainInsights . If
// this condition isn’t met, no data is listed for that control domain.
//
// [ListDomains]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListDomains.html
func (c *Client) ListControlDomainInsights(ctx context.Context, params *ListControlDomainInsightsInput, optFns ...func(*Options)) (*ListControlDomainInsightsOutput, error) {
	if params == nil {
		params = &ListControlDomainInsightsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListControlDomainInsights", params, optFns, c.addOperationListControlDomainInsightsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListControlDomainInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListControlDomainInsightsInput struct {

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListControlDomainInsightsOutput struct {

	// The control domain analytics data that the ListControlDomainInsights API
	// returned.
	ControlDomainInsights []types.ControlDomainInsights

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListControlDomainInsightsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListControlDomainInsights{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListControlDomainInsights{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListControlDomainInsights"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListControlDomainInsights(options.Region), middleware.Before); err != nil {
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

// ListControlDomainInsightsPaginatorOptions is the paginator options for
// ListControlDomainInsights
type ListControlDomainInsightsPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListControlDomainInsightsPaginator is a paginator for ListControlDomainInsights
type ListControlDomainInsightsPaginator struct {
	options   ListControlDomainInsightsPaginatorOptions
	client    ListControlDomainInsightsAPIClient
	params    *ListControlDomainInsightsInput
	nextToken *string
	firstPage bool
}

// NewListControlDomainInsightsPaginator returns a new
// ListControlDomainInsightsPaginator
func NewListControlDomainInsightsPaginator(client ListControlDomainInsightsAPIClient, params *ListControlDomainInsightsInput, optFns ...func(*ListControlDomainInsightsPaginatorOptions)) *ListControlDomainInsightsPaginator {
	if params == nil {
		params = &ListControlDomainInsightsInput{}
	}

	options := ListControlDomainInsightsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListControlDomainInsightsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListControlDomainInsightsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListControlDomainInsights page.
func (p *ListControlDomainInsightsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListControlDomainInsightsOutput, error) {
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
	result, err := p.client.ListControlDomainInsights(ctx, &params, optFns...)
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

// ListControlDomainInsightsAPIClient is a client that implements the
// ListControlDomainInsights operation.
type ListControlDomainInsightsAPIClient interface {
	ListControlDomainInsights(context.Context, *ListControlDomainInsightsInput, ...func(*Options)) (*ListControlDomainInsightsOutput, error)
}

var _ ListControlDomainInsightsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListControlDomainInsights(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListControlDomainInsights",
	}
}