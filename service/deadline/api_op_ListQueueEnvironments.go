// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists queue environments.
func (c *Client) ListQueueEnvironments(ctx context.Context, params *ListQueueEnvironmentsInput, optFns ...func(*Options)) (*ListQueueEnvironmentsOutput, error) {
	if params == nil {
		params = &ListQueueEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueueEnvironments", params, optFns, c.addOperationListQueueEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueueEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueueEnvironmentsInput struct {

	// The farm ID for the queue environment list.
	//
	// This member is required.
	FarmId *string

	// The queue ID for the queue environment list.
	//
	// This member is required.
	QueueId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token for the next set of results, or null to start from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListQueueEnvironmentsOutput struct {

	// The environments to include in the queue environments list.
	//
	// This member is required.
	Environments []types.QueueEnvironmentSummary

	// If Deadline Cloud returns nextToken , then there are more results available. The
	// value of nextToken is a unique pagination token for each page. To retrieve the
	// next page, call the operation again using the returned token. Keep all other
	// arguments unchanged. If no results remain, then nextToken is set to null . Each
	// pagination token expires after 24 hours. If you provide a token that isn't
	// valid, then you receive an HTTP 400 ValidationException error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueueEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueueEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueueEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListQueueEnvironments"); err != nil {
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
	if err = addEndpointPrefix_opListQueueEnvironmentsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListQueueEnvironmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueueEnvironments(options.Region), middleware.Before); err != nil {
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

// ListQueueEnvironmentsPaginatorOptions is the paginator options for
// ListQueueEnvironments
type ListQueueEnvironmentsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueueEnvironmentsPaginator is a paginator for ListQueueEnvironments
type ListQueueEnvironmentsPaginator struct {
	options   ListQueueEnvironmentsPaginatorOptions
	client    ListQueueEnvironmentsAPIClient
	params    *ListQueueEnvironmentsInput
	nextToken *string
	firstPage bool
}

// NewListQueueEnvironmentsPaginator returns a new ListQueueEnvironmentsPaginator
func NewListQueueEnvironmentsPaginator(client ListQueueEnvironmentsAPIClient, params *ListQueueEnvironmentsInput, optFns ...func(*ListQueueEnvironmentsPaginatorOptions)) *ListQueueEnvironmentsPaginator {
	if params == nil {
		params = &ListQueueEnvironmentsInput{}
	}

	options := ListQueueEnvironmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueueEnvironmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueueEnvironmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueueEnvironments page.
func (p *ListQueueEnvironmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueueEnvironmentsOutput, error) {
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
	result, err := p.client.ListQueueEnvironments(ctx, &params, optFns...)
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

type endpointPrefix_opListQueueEnvironmentsMiddleware struct {
}

func (*endpointPrefix_opListQueueEnvironmentsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListQueueEnvironmentsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListQueueEnvironmentsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListQueueEnvironmentsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListQueueEnvironmentsAPIClient is a client that implements the
// ListQueueEnvironments operation.
type ListQueueEnvironmentsAPIClient interface {
	ListQueueEnvironments(context.Context, *ListQueueEnvironmentsInput, ...func(*Options)) (*ListQueueEnvironmentsOutput, error)
}

var _ ListQueueEnvironmentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListQueueEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListQueueEnvironments",
	}
}