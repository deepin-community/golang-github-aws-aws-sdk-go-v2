// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the generation requests made for a bot locale.
func (c *Client) ListBotResourceGenerations(ctx context.Context, params *ListBotResourceGenerationsInput, optFns ...func(*Options)) (*ListBotResourceGenerationsOutput, error) {
	if params == nil {
		params = &ListBotResourceGenerationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBotResourceGenerations", params, optFns, c.addOperationListBotResourceGenerationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBotResourceGenerationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBotResourceGenerationsInput struct {

	// The unique identifier of the bot whose generation requests you want to view.
	//
	// This member is required.
	BotId *string

	// The version of the bot whose generation requests you want to view.
	//
	// This member is required.
	BotVersion *string

	// The locale of the bot whose generation requests you want to view.
	//
	// This member is required.
	LocaleId *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// If the total number of results is greater than the number specified in the
	// maxResults , the response returns a token in the nextToken field. Use this
	// token when making a request to return the next batch of results.
	NextToken *string

	// An object containing information about the attribute and the method by which to
	// sort the results
	SortBy *types.GenerationSortBy

	noSmithyDocumentSerde
}

type ListBotResourceGenerationsOutput struct {

	// The unique identifier of the bot for which the generation requests were made.
	BotId *string

	// The version of the bot for which the generation requests were made.
	BotVersion *string

	// A list of objects, each containing information about a generation request for
	// the bot locale.
	GenerationSummaries []types.GenerationSummary

	// The locale of the bot for which the generation requests were made.
	LocaleId *string

	// If the total number of results is greater than the number specified in the
	// maxResults , the response returns a token in the nextToken field. Use this
	// token when making a request to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBotResourceGenerationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBotResourceGenerations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBotResourceGenerations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBotResourceGenerations"); err != nil {
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
	if err = addOpListBotResourceGenerationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBotResourceGenerations(options.Region), middleware.Before); err != nil {
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

// ListBotResourceGenerationsPaginatorOptions is the paginator options for
// ListBotResourceGenerations
type ListBotResourceGenerationsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBotResourceGenerationsPaginator is a paginator for
// ListBotResourceGenerations
type ListBotResourceGenerationsPaginator struct {
	options   ListBotResourceGenerationsPaginatorOptions
	client    ListBotResourceGenerationsAPIClient
	params    *ListBotResourceGenerationsInput
	nextToken *string
	firstPage bool
}

// NewListBotResourceGenerationsPaginator returns a new
// ListBotResourceGenerationsPaginator
func NewListBotResourceGenerationsPaginator(client ListBotResourceGenerationsAPIClient, params *ListBotResourceGenerationsInput, optFns ...func(*ListBotResourceGenerationsPaginatorOptions)) *ListBotResourceGenerationsPaginator {
	if params == nil {
		params = &ListBotResourceGenerationsInput{}
	}

	options := ListBotResourceGenerationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBotResourceGenerationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBotResourceGenerationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBotResourceGenerations page.
func (p *ListBotResourceGenerationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBotResourceGenerationsOutput, error) {
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
	result, err := p.client.ListBotResourceGenerations(ctx, &params, optFns...)
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

// ListBotResourceGenerationsAPIClient is a client that implements the
// ListBotResourceGenerations operation.
type ListBotResourceGenerationsAPIClient interface {
	ListBotResourceGenerations(context.Context, *ListBotResourceGenerationsInput, ...func(*Options)) (*ListBotResourceGenerationsOutput, error)
}

var _ ListBotResourceGenerationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBotResourceGenerations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBotResourceGenerations",
	}
}