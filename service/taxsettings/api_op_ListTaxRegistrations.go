// Code generated by smithy-go-codegen DO NOT EDIT.

package taxsettings

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/taxsettings/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the tax registration of accounts listed in a consolidated billing
// family. This can be used to retrieve up to 100 accounts' tax registrations in
// one call (default 50).
func (c *Client) ListTaxRegistrations(ctx context.Context, params *ListTaxRegistrationsInput, optFns ...func(*Options)) (*ListTaxRegistrationsOutput, error) {
	if params == nil {
		params = &ListTaxRegistrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTaxRegistrations", params, optFns, c.addOperationListTaxRegistrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTaxRegistrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTaxRegistrationsInput struct {

	// Number of accountDetails results you want in one response.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTaxRegistrationsOutput struct {

	// The list of account details. This contains account Ids and TRN Information for
	// each of the linked accounts.
	//
	// This member is required.
	AccountDetails []types.AccountDetails

	//  The token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTaxRegistrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTaxRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTaxRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTaxRegistrations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTaxRegistrations(options.Region), middleware.Before); err != nil {
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

// ListTaxRegistrationsPaginatorOptions is the paginator options for
// ListTaxRegistrations
type ListTaxRegistrationsPaginatorOptions struct {
	// Number of accountDetails results you want in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTaxRegistrationsPaginator is a paginator for ListTaxRegistrations
type ListTaxRegistrationsPaginator struct {
	options   ListTaxRegistrationsPaginatorOptions
	client    ListTaxRegistrationsAPIClient
	params    *ListTaxRegistrationsInput
	nextToken *string
	firstPage bool
}

// NewListTaxRegistrationsPaginator returns a new ListTaxRegistrationsPaginator
func NewListTaxRegistrationsPaginator(client ListTaxRegistrationsAPIClient, params *ListTaxRegistrationsInput, optFns ...func(*ListTaxRegistrationsPaginatorOptions)) *ListTaxRegistrationsPaginator {
	if params == nil {
		params = &ListTaxRegistrationsInput{}
	}

	options := ListTaxRegistrationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTaxRegistrationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTaxRegistrationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTaxRegistrations page.
func (p *ListTaxRegistrationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTaxRegistrationsOutput, error) {
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
	result, err := p.client.ListTaxRegistrations(ctx, &params, optFns...)
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

// ListTaxRegistrationsAPIClient is a client that implements the
// ListTaxRegistrations operation.
type ListTaxRegistrationsAPIClient interface {
	ListTaxRegistrations(context.Context, *ListTaxRegistrationsInput, ...func(*Options)) (*ListTaxRegistrationsOutput, error)
}

var _ ListTaxRegistrationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTaxRegistrations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTaxRegistrations",
	}
}