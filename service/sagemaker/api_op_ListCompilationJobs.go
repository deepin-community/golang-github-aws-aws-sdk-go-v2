// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists model compilation jobs that satisfy various filters.
//
// To create a model compilation job, use [CreateCompilationJob]. To get information about a particular
// model compilation job you have created, use [DescribeCompilationJob].
//
// [DescribeCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeCompilationJob.html
// [CreateCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateCompilationJob.html
func (c *Client) ListCompilationJobs(ctx context.Context, params *ListCompilationJobsInput, optFns ...func(*Options)) (*ListCompilationJobsOutput, error) {
	if params == nil {
		params = &ListCompilationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCompilationJobs", params, optFns, c.addOperationListCompilationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCompilationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCompilationJobsInput struct {

	// A filter that returns the model compilation jobs that were created after a
	// specified time.
	CreationTimeAfter *time.Time

	// A filter that returns the model compilation jobs that were created before a
	// specified time.
	CreationTimeBefore *time.Time

	// A filter that returns the model compilation jobs that were modified after a
	// specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns the model compilation jobs that were modified before a
	// specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of model compilation jobs to return in the response.
	MaxResults *int32

	// A filter that returns the model compilation jobs whose name contains a
	// specified string.
	NameContains *string

	// If the result of the previous ListCompilationJobs request was truncated, the
	// response includes a NextToken . To retrieve the next set of model compilation
	// jobs, use the token in the next request.
	NextToken *string

	// The field by which to sort results. The default is CreationTime .
	SortBy types.ListCompilationJobsSortBy

	// The sort order for results. The default is Ascending .
	SortOrder types.SortOrder

	// A filter that retrieves model compilation jobs with a specific
	// CompilationJobStatus status.
	StatusEquals types.CompilationJobStatus

	noSmithyDocumentSerde
}

type ListCompilationJobsOutput struct {

	// An array of [CompilationJobSummary] objects, each describing a model compilation job.
	//
	// [CompilationJobSummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CompilationJobSummary.html
	//
	// This member is required.
	CompilationJobSummaries []types.CompilationJobSummary

	// If the response is truncated, Amazon SageMaker returns this NextToken . To
	// retrieve the next set of model compilation jobs, use this token in the next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCompilationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCompilationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCompilationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCompilationJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCompilationJobs(options.Region), middleware.Before); err != nil {
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

// ListCompilationJobsPaginatorOptions is the paginator options for
// ListCompilationJobs
type ListCompilationJobsPaginatorOptions struct {
	// The maximum number of model compilation jobs to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCompilationJobsPaginator is a paginator for ListCompilationJobs
type ListCompilationJobsPaginator struct {
	options   ListCompilationJobsPaginatorOptions
	client    ListCompilationJobsAPIClient
	params    *ListCompilationJobsInput
	nextToken *string
	firstPage bool
}

// NewListCompilationJobsPaginator returns a new ListCompilationJobsPaginator
func NewListCompilationJobsPaginator(client ListCompilationJobsAPIClient, params *ListCompilationJobsInput, optFns ...func(*ListCompilationJobsPaginatorOptions)) *ListCompilationJobsPaginator {
	if params == nil {
		params = &ListCompilationJobsInput{}
	}

	options := ListCompilationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCompilationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCompilationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCompilationJobs page.
func (p *ListCompilationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCompilationJobsOutput, error) {
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
	result, err := p.client.ListCompilationJobs(ctx, &params, optFns...)
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

// ListCompilationJobsAPIClient is a client that implements the
// ListCompilationJobs operation.
type ListCompilationJobsAPIClient interface {
	ListCompilationJobs(context.Context, *ListCompilationJobsInput, ...func(*Options)) (*ListCompilationJobsOutput, error)
}

var _ ListCompilationJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCompilationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCompilationJobs",
	}
}