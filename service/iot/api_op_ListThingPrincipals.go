// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the principals associated with the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
//
// Requires permission to access the [ListThingPrincipals] action.
//
// [ListThingPrincipals]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListThingPrincipals(ctx context.Context, params *ListThingPrincipalsInput, optFns ...func(*Options)) (*ListThingPrincipalsOutput, error) {
	if params == nil {
		params = &ListThingPrincipalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingPrincipals", params, optFns, c.addOperationListThingPrincipalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThingPrincipal operation.
type ListThingPrincipalsInput struct {

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	// The maximum number of results to return in this operation.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

// The output from the ListThingPrincipals operation.
type ListThingPrincipalsOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The principals associated with the thing.
	Principals []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingPrincipalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListThingPrincipals"); err != nil {
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
	if err = addOpListThingPrincipalsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingPrincipals(options.Region), middleware.Before); err != nil {
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

// ListThingPrincipalsPaginatorOptions is the paginator options for
// ListThingPrincipals
type ListThingPrincipalsPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingPrincipalsPaginator is a paginator for ListThingPrincipals
type ListThingPrincipalsPaginator struct {
	options   ListThingPrincipalsPaginatorOptions
	client    ListThingPrincipalsAPIClient
	params    *ListThingPrincipalsInput
	nextToken *string
	firstPage bool
}

// NewListThingPrincipalsPaginator returns a new ListThingPrincipalsPaginator
func NewListThingPrincipalsPaginator(client ListThingPrincipalsAPIClient, params *ListThingPrincipalsInput, optFns ...func(*ListThingPrincipalsPaginatorOptions)) *ListThingPrincipalsPaginator {
	if params == nil {
		params = &ListThingPrincipalsInput{}
	}

	options := ListThingPrincipalsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingPrincipalsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingPrincipalsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingPrincipals page.
func (p *ListThingPrincipalsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingPrincipalsOutput, error) {
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
	result, err := p.client.ListThingPrincipals(ctx, &params, optFns...)
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

// ListThingPrincipalsAPIClient is a client that implements the
// ListThingPrincipals operation.
type ListThingPrincipalsAPIClient interface {
	ListThingPrincipals(context.Context, *ListThingPrincipalsInput, ...func(*Options)) (*ListThingPrincipalsOutput, error)
}

var _ ListThingPrincipalsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListThingPrincipals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListThingPrincipals",
	}
}