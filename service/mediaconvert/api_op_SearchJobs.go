// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a JSON array that includes job details for up to twenty of your most
// recent jobs. Optionally filter results further according to input file, queue,
// or status. To retrieve the twenty next most recent jobs, use the nextToken
// string returned with the array.
func (c *Client) SearchJobs(ctx context.Context, params *SearchJobsInput, optFns ...func(*Options)) (*SearchJobsOutput, error) {
	if params == nil {
		params = &SearchJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchJobs", params, optFns, c.addOperationSearchJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchJobsInput struct {

	// Optional. Provide your input file URL or your partial input file name. The
	// maximum length for an input file is 300 characters.
	InputFile *string

	// Optional. Number of jobs, up to twenty, that will be returned at one time.
	MaxResults *int32

	// Optional. Use this string, provided with the response to a previous request, to
	// request the next batch of jobs.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order

	// Optional. Provide a queue name, or a queue ARN, to return only jobs from that
	// queue.
	Queue *string

	// Optional. A job's status can be SUBMITTED, PROGRESSING, COMPLETE, CANCELED, or
	// ERROR.
	Status types.JobStatus

	noSmithyDocumentSerde
}

type SearchJobsOutput struct {

	// List of jobs.
	Jobs []types.Job

	// Use this string to request the next batch of jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchJobs(options.Region), middleware.Before); err != nil {
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

// SearchJobsPaginatorOptions is the paginator options for SearchJobs
type SearchJobsPaginatorOptions struct {
	// Optional. Number of jobs, up to twenty, that will be returned at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchJobsPaginator is a paginator for SearchJobs
type SearchJobsPaginator struct {
	options   SearchJobsPaginatorOptions
	client    SearchJobsAPIClient
	params    *SearchJobsInput
	nextToken *string
	firstPage bool
}

// NewSearchJobsPaginator returns a new SearchJobsPaginator
func NewSearchJobsPaginator(client SearchJobsAPIClient, params *SearchJobsInput, optFns ...func(*SearchJobsPaginatorOptions)) *SearchJobsPaginator {
	if params == nil {
		params = &SearchJobsInput{}
	}

	options := SearchJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchJobs page.
func (p *SearchJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchJobsOutput, error) {
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
	result, err := p.client.SearchJobs(ctx, &params, optFns...)
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

// SearchJobsAPIClient is a client that implements the SearchJobs operation.
type SearchJobsAPIClient interface {
	SearchJobs(context.Context, *SearchJobsInput, ...func(*Options)) (*SearchJobsOutput, error)
}

var _ SearchJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchJobs",
	}
}