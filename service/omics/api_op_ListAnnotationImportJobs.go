// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of annotation import jobs.
func (c *Client) ListAnnotationImportJobs(ctx context.Context, params *ListAnnotationImportJobsInput, optFns ...func(*Options)) (*ListAnnotationImportJobsOutput, error) {
	if params == nil {
		params = &ListAnnotationImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnnotationImportJobs", params, optFns, c.addOperationListAnnotationImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnnotationImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnnotationImportJobsInput struct {

	// A filter to apply to the list.
	Filter *types.ListAnnotationImportJobsFilter

	// IDs of annotation import jobs to retrieve.
	Ids []string

	// The maximum number of jobs to return in one page of results.
	MaxResults *int32

	// Specifies the pagination token from a previous request to retrieve the next
	// page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAnnotationImportJobsOutput struct {

	// A list of jobs.
	AnnotationImportJobs []types.AnnotationImportJobItem

	// Specifies the pagination token from a previous request to retrieve the next
	// page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnnotationImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnnotationImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnnotationImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAnnotationImportJobs"); err != nil {
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
	if err = addEndpointPrefix_opListAnnotationImportJobsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnnotationImportJobs(options.Region), middleware.Before); err != nil {
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

// ListAnnotationImportJobsPaginatorOptions is the paginator options for
// ListAnnotationImportJobs
type ListAnnotationImportJobsPaginatorOptions struct {
	// The maximum number of jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnnotationImportJobsPaginator is a paginator for ListAnnotationImportJobs
type ListAnnotationImportJobsPaginator struct {
	options   ListAnnotationImportJobsPaginatorOptions
	client    ListAnnotationImportJobsAPIClient
	params    *ListAnnotationImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListAnnotationImportJobsPaginator returns a new
// ListAnnotationImportJobsPaginator
func NewListAnnotationImportJobsPaginator(client ListAnnotationImportJobsAPIClient, params *ListAnnotationImportJobsInput, optFns ...func(*ListAnnotationImportJobsPaginatorOptions)) *ListAnnotationImportJobsPaginator {
	if params == nil {
		params = &ListAnnotationImportJobsInput{}
	}

	options := ListAnnotationImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnnotationImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnnotationImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnnotationImportJobs page.
func (p *ListAnnotationImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnnotationImportJobsOutput, error) {
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
	result, err := p.client.ListAnnotationImportJobs(ctx, &params, optFns...)
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

type endpointPrefix_opListAnnotationImportJobsMiddleware struct {
}

func (*endpointPrefix_opListAnnotationImportJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAnnotationImportJobsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListAnnotationImportJobsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListAnnotationImportJobsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListAnnotationImportJobsAPIClient is a client that implements the
// ListAnnotationImportJobs operation.
type ListAnnotationImportJobsAPIClient interface {
	ListAnnotationImportJobs(context.Context, *ListAnnotationImportJobsInput, ...func(*Options)) (*ListAnnotationImportJobsOutput, error)
}

var _ ListAnnotationImportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAnnotationImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAnnotationImportJobs",
	}
}