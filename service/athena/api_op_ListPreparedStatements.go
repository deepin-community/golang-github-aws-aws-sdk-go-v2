// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the prepared statements in the specified workgroup.
func (c *Client) ListPreparedStatements(ctx context.Context, params *ListPreparedStatementsInput, optFns ...func(*Options)) (*ListPreparedStatementsOutput, error) {
	if params == nil {
		params = &ListPreparedStatementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPreparedStatements", params, optFns, c.addOperationListPreparedStatementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPreparedStatementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPreparedStatementsInput struct {

	// The workgroup to list the prepared statements for.
	//
	// This member is required.
	WorkGroup *string

	// The maximum number of results to return in this request.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPreparedStatementsOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The list of prepared statements for the workgroup.
	PreparedStatements []types.PreparedStatementSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPreparedStatementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPreparedStatements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPreparedStatements{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPreparedStatements"); err != nil {
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
	if err = addOpListPreparedStatementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPreparedStatements(options.Region), middleware.Before); err != nil {
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

// ListPreparedStatementsPaginatorOptions is the paginator options for
// ListPreparedStatements
type ListPreparedStatementsPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPreparedStatementsPaginator is a paginator for ListPreparedStatements
type ListPreparedStatementsPaginator struct {
	options   ListPreparedStatementsPaginatorOptions
	client    ListPreparedStatementsAPIClient
	params    *ListPreparedStatementsInput
	nextToken *string
	firstPage bool
}

// NewListPreparedStatementsPaginator returns a new ListPreparedStatementsPaginator
func NewListPreparedStatementsPaginator(client ListPreparedStatementsAPIClient, params *ListPreparedStatementsInput, optFns ...func(*ListPreparedStatementsPaginatorOptions)) *ListPreparedStatementsPaginator {
	if params == nil {
		params = &ListPreparedStatementsInput{}
	}

	options := ListPreparedStatementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPreparedStatementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPreparedStatementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPreparedStatements page.
func (p *ListPreparedStatementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPreparedStatementsOutput, error) {
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
	result, err := p.client.ListPreparedStatements(ctx, &params, optFns...)
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

// ListPreparedStatementsAPIClient is a client that implements the
// ListPreparedStatements operation.
type ListPreparedStatementsAPIClient interface {
	ListPreparedStatements(context.Context, *ListPreparedStatementsInput, ...func(*Options)) (*ListPreparedStatementsOutput, error)
}

var _ ListPreparedStatementsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPreparedStatements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPreparedStatements",
	}
}