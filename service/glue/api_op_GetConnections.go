// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of connection definitions from the Data Catalog.
func (c *Client) GetConnections(ctx context.Context, params *GetConnectionsInput, optFns ...func(*Options)) (*GetConnectionsOutput, error) {
	if params == nil {
		params = &GetConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnections", params, optFns, c.addOperationGetConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectionsInput struct {

	// The ID of the Data Catalog in which the connections reside. If none is
	// provided, the Amazon Web Services account ID is used by default.
	CatalogId *string

	// A filter that controls which connections are returned.
	Filter *types.GetConnectionsFilter

	// Allows you to retrieve the connection metadata without returning the password.
	// For instance, the Glue console uses this flag to retrieve the connection, and
	// does not display the password. Set this parameter when the caller might not have
	// permission to use the KMS key to decrypt the password, but it does have
	// permission to access the rest of the connection properties.
	HidePassword bool

	// The maximum number of connections to return in one response.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConnectionsOutput struct {

	// A list of requested connection definitions.
	ConnectionList []types.Connection

	// A continuation token, if the list of connections returned does not include the
	// last of the filtered connections.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConnections"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnections(options.Region), middleware.Before); err != nil {
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

// GetConnectionsPaginatorOptions is the paginator options for GetConnections
type GetConnectionsPaginatorOptions struct {
	// The maximum number of connections to return in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetConnectionsPaginator is a paginator for GetConnections
type GetConnectionsPaginator struct {
	options   GetConnectionsPaginatorOptions
	client    GetConnectionsAPIClient
	params    *GetConnectionsInput
	nextToken *string
	firstPage bool
}

// NewGetConnectionsPaginator returns a new GetConnectionsPaginator
func NewGetConnectionsPaginator(client GetConnectionsAPIClient, params *GetConnectionsInput, optFns ...func(*GetConnectionsPaginatorOptions)) *GetConnectionsPaginator {
	if params == nil {
		params = &GetConnectionsInput{}
	}

	options := GetConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetConnections page.
func (p *GetConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetConnectionsOutput, error) {
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
	result, err := p.client.GetConnections(ctx, &params, optFns...)
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

// GetConnectionsAPIClient is a client that implements the GetConnections
// operation.
type GetConnectionsAPIClient interface {
	GetConnections(context.Context, *GetConnectionsInput, ...func(*Options)) (*GetConnectionsOutput, error)
}

var _ GetConnectionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConnections",
	}
}