// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List repository sync definitions with detail data.
func (c *Client) ListRepositorySyncDefinitions(ctx context.Context, params *ListRepositorySyncDefinitionsInput, optFns ...func(*Options)) (*ListRepositorySyncDefinitionsOutput, error) {
	if params == nil {
		params = &ListRepositorySyncDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositorySyncDefinitions", params, optFns, c.addOperationListRepositorySyncDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositorySyncDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositorySyncDefinitionsInput struct {

	// The repository name.
	//
	// This member is required.
	RepositoryName *string

	// The repository provider.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The sync type. The only supported value is TEMPLATE_SYNC .
	//
	// This member is required.
	SyncType types.SyncType

	// A token that indicates the location of the next repository sync definition in
	// the array of repository sync definitions, after the list of repository sync
	// definitions previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRepositorySyncDefinitionsOutput struct {

	// An array of repository sync definitions.
	//
	// This member is required.
	SyncDefinitions []types.RepositorySyncDefinition

	// A token that indicates the location of the next repository sync definition in
	// the array of repository sync definitions, after the current requested list of
	// repository sync definitions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRepositorySyncDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRepositorySyncDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRepositorySyncDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRepositorySyncDefinitions"); err != nil {
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
	if err = addOpListRepositorySyncDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositorySyncDefinitions(options.Region), middleware.Before); err != nil {
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

// ListRepositorySyncDefinitionsPaginatorOptions is the paginator options for
// ListRepositorySyncDefinitions
type ListRepositorySyncDefinitionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRepositorySyncDefinitionsPaginator is a paginator for
// ListRepositorySyncDefinitions
type ListRepositorySyncDefinitionsPaginator struct {
	options   ListRepositorySyncDefinitionsPaginatorOptions
	client    ListRepositorySyncDefinitionsAPIClient
	params    *ListRepositorySyncDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListRepositorySyncDefinitionsPaginator returns a new
// ListRepositorySyncDefinitionsPaginator
func NewListRepositorySyncDefinitionsPaginator(client ListRepositorySyncDefinitionsAPIClient, params *ListRepositorySyncDefinitionsInput, optFns ...func(*ListRepositorySyncDefinitionsPaginatorOptions)) *ListRepositorySyncDefinitionsPaginator {
	if params == nil {
		params = &ListRepositorySyncDefinitionsInput{}
	}

	options := ListRepositorySyncDefinitionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRepositorySyncDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRepositorySyncDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRepositorySyncDefinitions page.
func (p *ListRepositorySyncDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRepositorySyncDefinitionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListRepositorySyncDefinitions(ctx, &params, optFns...)
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

// ListRepositorySyncDefinitionsAPIClient is a client that implements the
// ListRepositorySyncDefinitions operation.
type ListRepositorySyncDefinitionsAPIClient interface {
	ListRepositorySyncDefinitions(context.Context, *ListRepositorySyncDefinitionsInput, ...func(*Options)) (*ListRepositorySyncDefinitionsOutput, error)
}

var _ ListRepositorySyncDefinitionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRepositorySyncDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRepositorySyncDefinitions",
	}
}