// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the installed vCenter clients.
func (c *Client) DescribeVcenterClients(ctx context.Context, params *DescribeVcenterClientsInput, optFns ...func(*Options)) (*DescribeVcenterClientsOutput, error) {
	if params == nil {
		params = &DescribeVcenterClientsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVcenterClients", params, optFns, c.addOperationDescribeVcenterClientsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVcenterClientsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVcenterClientsInput struct {

	// Maximum results to be returned in DescribeVcenterClients.
	MaxResults *int32

	// Next pagination token to be provided for DescribeVcenterClients.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeVcenterClientsOutput struct {

	// List of items returned by DescribeVcenterClients.
	Items []types.VcenterClient

	// Next pagination token returned from DescribeVcenterClients.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVcenterClientsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVcenterClients{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVcenterClients{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVcenterClients"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVcenterClients(options.Region), middleware.Before); err != nil {
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

// DescribeVcenterClientsPaginatorOptions is the paginator options for
// DescribeVcenterClients
type DescribeVcenterClientsPaginatorOptions struct {
	// Maximum results to be returned in DescribeVcenterClients.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVcenterClientsPaginator is a paginator for DescribeVcenterClients
type DescribeVcenterClientsPaginator struct {
	options   DescribeVcenterClientsPaginatorOptions
	client    DescribeVcenterClientsAPIClient
	params    *DescribeVcenterClientsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVcenterClientsPaginator returns a new DescribeVcenterClientsPaginator
func NewDescribeVcenterClientsPaginator(client DescribeVcenterClientsAPIClient, params *DescribeVcenterClientsInput, optFns ...func(*DescribeVcenterClientsPaginatorOptions)) *DescribeVcenterClientsPaginator {
	if params == nil {
		params = &DescribeVcenterClientsInput{}
	}

	options := DescribeVcenterClientsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVcenterClientsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVcenterClientsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVcenterClients page.
func (p *DescribeVcenterClientsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVcenterClientsOutput, error) {
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
	result, err := p.client.DescribeVcenterClients(ctx, &params, optFns...)
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

// DescribeVcenterClientsAPIClient is a client that implements the
// DescribeVcenterClients operation.
type DescribeVcenterClientsAPIClient interface {
	DescribeVcenterClients(context.Context, *DescribeVcenterClientsInput, ...func(*Options)) (*DescribeVcenterClientsOutput, error)
}

var _ DescribeVcenterClientsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeVcenterClients(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVcenterClients",
	}
}