// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of messages associated with an Amazon Q Business web experience.
func (c *Client) ListMessages(ctx context.Context, params *ListMessagesInput, optFns ...func(*Options)) (*ListMessagesOutput, error) {
	if params == nil {
		params = &ListMessagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMessages", params, optFns, c.addOperationListMessagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMessagesInput struct {

	// The identifier for the Amazon Q Business application.
	//
	// This member is required.
	ApplicationId *string

	// The identifier of the Amazon Q Business web experience conversation.
	//
	// This member is required.
	ConversationId *string

	// The maximum number of messages to return.
	MaxResults *int32

	// If the number of retrievers returned exceeds maxResults , Amazon Q Business
	// returns a next token as a pagination token to retrieve the next set of messages.
	NextToken *string

	// The identifier of the user involved in the Amazon Q Business web experience
	// conversation.
	UserId *string

	noSmithyDocumentSerde
}

type ListMessagesOutput struct {

	// An array of information on one or more messages.
	Messages []types.Message

	// If the response is truncated, Amazon Q Business returns this token, which you
	// can use in a later request to list the next set of messages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMessagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMessages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMessages{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMessages"); err != nil {
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
	if err = addOpListMessagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMessages(options.Region), middleware.Before); err != nil {
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

// ListMessagesPaginatorOptions is the paginator options for ListMessages
type ListMessagesPaginatorOptions struct {
	// The maximum number of messages to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMessagesPaginator is a paginator for ListMessages
type ListMessagesPaginator struct {
	options   ListMessagesPaginatorOptions
	client    ListMessagesAPIClient
	params    *ListMessagesInput
	nextToken *string
	firstPage bool
}

// NewListMessagesPaginator returns a new ListMessagesPaginator
func NewListMessagesPaginator(client ListMessagesAPIClient, params *ListMessagesInput, optFns ...func(*ListMessagesPaginatorOptions)) *ListMessagesPaginator {
	if params == nil {
		params = &ListMessagesInput{}
	}

	options := ListMessagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMessagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMessagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMessages page.
func (p *ListMessagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMessagesOutput, error) {
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
	result, err := p.client.ListMessages(ctx, &params, optFns...)
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

// ListMessagesAPIClient is a client that implements the ListMessages operation.
type ListMessagesAPIClient interface {
	ListMessages(context.Context, *ListMessagesInput, ...func(*Options)) (*ListMessagesOutput, error)
}

var _ ListMessagesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListMessages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMessages",
	}
}