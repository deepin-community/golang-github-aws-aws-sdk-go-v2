// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches the provisioned phone numbers in an organization.
func (c *Client) SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) {
	if params == nil {
		params = &SearchAvailablePhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAvailablePhoneNumbers", params, optFns, c.addOperationSearchAvailablePhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAvailablePhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAvailablePhoneNumbersInput struct {

	// Confines a search to just the phone numbers associated with the specified area
	// code.
	AreaCode *string

	// Confines a search to just the phone numbers associated with the specified city.
	City *string

	// Confines a search to just the phone numbers associated with the specified
	// country.
	Country *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token used to return the next page of results.
	NextToken *string

	// Confines a search to just the phone numbers associated with the specified phone
	// number type, either local or toll-free.
	PhoneNumberType types.PhoneNumberType

	// Confines a search to just the phone numbers associated with the specified state.
	State *string

	// Confines a search to just the phone numbers associated with the specified
	// toll-free prefix.
	TollFreePrefix *string

	noSmithyDocumentSerde
}

type SearchAvailablePhoneNumbersOutput struct {

	// Confines a search to just the phone numbers in the E.164 format.
	E164PhoneNumbers []string

	// The token used to return the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchAvailablePhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchAvailablePhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAvailablePhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchAvailablePhoneNumbers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAvailablePhoneNumbers(options.Region), middleware.Before); err != nil {
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

// SearchAvailablePhoneNumbersPaginatorOptions is the paginator options for
// SearchAvailablePhoneNumbers
type SearchAvailablePhoneNumbersPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchAvailablePhoneNumbersPaginator is a paginator for
// SearchAvailablePhoneNumbers
type SearchAvailablePhoneNumbersPaginator struct {
	options   SearchAvailablePhoneNumbersPaginatorOptions
	client    SearchAvailablePhoneNumbersAPIClient
	params    *SearchAvailablePhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewSearchAvailablePhoneNumbersPaginator returns a new
// SearchAvailablePhoneNumbersPaginator
func NewSearchAvailablePhoneNumbersPaginator(client SearchAvailablePhoneNumbersAPIClient, params *SearchAvailablePhoneNumbersInput, optFns ...func(*SearchAvailablePhoneNumbersPaginatorOptions)) *SearchAvailablePhoneNumbersPaginator {
	if params == nil {
		params = &SearchAvailablePhoneNumbersInput{}
	}

	options := SearchAvailablePhoneNumbersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchAvailablePhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchAvailablePhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchAvailablePhoneNumbers page.
func (p *SearchAvailablePhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) {
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
	result, err := p.client.SearchAvailablePhoneNumbers(ctx, &params, optFns...)
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

// SearchAvailablePhoneNumbersAPIClient is a client that implements the
// SearchAvailablePhoneNumbers operation.
type SearchAvailablePhoneNumbersAPIClient interface {
	SearchAvailablePhoneNumbers(context.Context, *SearchAvailablePhoneNumbersInput, ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error)
}

var _ SearchAvailablePhoneNumbersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchAvailablePhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchAvailablePhoneNumbers",
	}
}