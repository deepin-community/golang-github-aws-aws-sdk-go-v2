// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the action executions that have occurred in a pipeline.
func (c *Client) ListActionExecutions(ctx context.Context, params *ListActionExecutionsInput, optFns ...func(*Options)) (*ListActionExecutionsOutput, error) {
	if params == nil {
		params = &ListActionExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActionExecutions", params, optFns, c.addOperationListActionExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActionExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActionExecutionsInput struct {

	//  The name of the pipeline for which you want to list action execution history.
	//
	// This member is required.
	PipelineName *string

	// Input information used to filter action execution history.
	Filter *types.ActionExecutionFilter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. Action
	// execution history is retained for up to 12 months, based on action execution
	// start times. Default value is 100.
	MaxResults *int32

	// The token that was returned from the previous ListActionExecutions call, which
	// can be used to return the next set of action executions in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListActionExecutionsOutput struct {

	// The details for a list of recent executions, such as action execution ID.
	ActionExecutionDetails []types.ActionExecutionDetail

	// If the amount of returned information is significantly large, an identifier is
	// also returned and can be used in a subsequent ListActionExecutions call to
	// return the next set of action executions in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListActionExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListActionExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListActionExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListActionExecutions"); err != nil {
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
	if err = addOpListActionExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActionExecutions(options.Region), middleware.Before); err != nil {
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

// ListActionExecutionsPaginatorOptions is the paginator options for
// ListActionExecutions
type ListActionExecutionsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. Action
	// execution history is retained for up to 12 months, based on action execution
	// start times. Default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListActionExecutionsPaginator is a paginator for ListActionExecutions
type ListActionExecutionsPaginator struct {
	options   ListActionExecutionsPaginatorOptions
	client    ListActionExecutionsAPIClient
	params    *ListActionExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListActionExecutionsPaginator returns a new ListActionExecutionsPaginator
func NewListActionExecutionsPaginator(client ListActionExecutionsAPIClient, params *ListActionExecutionsInput, optFns ...func(*ListActionExecutionsPaginatorOptions)) *ListActionExecutionsPaginator {
	if params == nil {
		params = &ListActionExecutionsInput{}
	}

	options := ListActionExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListActionExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListActionExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListActionExecutions page.
func (p *ListActionExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListActionExecutionsOutput, error) {
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
	result, err := p.client.ListActionExecutions(ctx, &params, optFns...)
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

// ListActionExecutionsAPIClient is a client that implements the
// ListActionExecutions operation.
type ListActionExecutionsAPIClient interface {
	ListActionExecutions(context.Context, *ListActionExecutionsInput, ...func(*Options)) (*ListActionExecutionsOutput, error)
}

var _ ListActionExecutionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListActionExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListActionExecutions",
	}
}