// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about all public keys in your account, in the AWS
// region where the API request is processed.
func (c *Client) ListPublicKeys(ctx context.Context, params *ListPublicKeysInput, optFns ...func(*Options)) (*ListPublicKeysOutput, error) {
	if params == nil {
		params = &ListPublicKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublicKeys", params, optFns, c.addOperationListPublicKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublicKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublicKeysInput struct {

	// Maximum number of results to return. Default: 50.
	MaxResults *int32

	// The first public key to retrieve. This is used for pagination; see the nextToken
	// response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPublicKeysOutput struct {

	// List of the matching public keys (summary information only).
	//
	// This member is required.
	PublicKeys []types.PublicKeySummary

	// If there are more public keys than maxResults , use nextToken in the request to
	// get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPublicKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPublicKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPublicKeys{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPublicKeys"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPublicKeys(options.Region), middleware.Before); err != nil {
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

// ListPublicKeysPaginatorOptions is the paginator options for ListPublicKeys
type ListPublicKeysPaginatorOptions struct {
	// Maximum number of results to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPublicKeysPaginator is a paginator for ListPublicKeys
type ListPublicKeysPaginator struct {
	options   ListPublicKeysPaginatorOptions
	client    ListPublicKeysAPIClient
	params    *ListPublicKeysInput
	nextToken *string
	firstPage bool
}

// NewListPublicKeysPaginator returns a new ListPublicKeysPaginator
func NewListPublicKeysPaginator(client ListPublicKeysAPIClient, params *ListPublicKeysInput, optFns ...func(*ListPublicKeysPaginatorOptions)) *ListPublicKeysPaginator {
	if params == nil {
		params = &ListPublicKeysInput{}
	}

	options := ListPublicKeysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPublicKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPublicKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPublicKeys page.
func (p *ListPublicKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPublicKeysOutput, error) {
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
	result, err := p.client.ListPublicKeys(ctx, &params, optFns...)
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

// ListPublicKeysAPIClient is a client that implements the ListPublicKeys
// operation.
type ListPublicKeysAPIClient interface {
	ListPublicKeys(context.Context, *ListPublicKeysInput, ...func(*Options)) (*ListPublicKeysOutput, error)
}

var _ ListPublicKeysAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPublicKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPublicKeys",
	}
}