// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified registration type definitions. You can use
// DescribeRegistrationTypeDefinitions to view the requirements for creating,
// filling out, and submitting each registration type.
func (c *Client) DescribeRegistrationTypeDefinitions(ctx context.Context, params *DescribeRegistrationTypeDefinitionsInput, optFns ...func(*Options)) (*DescribeRegistrationTypeDefinitionsOutput, error) {
	if params == nil {
		params = &DescribeRegistrationTypeDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRegistrationTypeDefinitions", params, optFns, c.addOperationDescribeRegistrationTypeDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRegistrationTypeDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegistrationTypeDefinitionsInput struct {

	// An array of RegistrationFilter objects to filter the results.
	Filters []types.RegistrationTypeFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// The type of registration form. The list of RegistrationTypes can be found using
	// the DescribeRegistrationTypeDefinitionsaction.
	RegistrationTypes []string

	noSmithyDocumentSerde
}

type DescribeRegistrationTypeDefinitionsOutput struct {

	// The type of registration form. The list of RegistrationTypes can be found using
	// the DescribeRegistrationTypeDefinitionsaction.
	//
	// This member is required.
	RegistrationTypeDefinitions []types.RegistrationTypeDefinition

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRegistrationTypeDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRegistrationTypeDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRegistrationTypeDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRegistrationTypeDefinitions"); err != nil {
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
	if err = addOpDescribeRegistrationTypeDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegistrationTypeDefinitions(options.Region), middleware.Before); err != nil {
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

// DescribeRegistrationTypeDefinitionsPaginatorOptions is the paginator options
// for DescribeRegistrationTypeDefinitions
type DescribeRegistrationTypeDefinitionsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRegistrationTypeDefinitionsPaginator is a paginator for
// DescribeRegistrationTypeDefinitions
type DescribeRegistrationTypeDefinitionsPaginator struct {
	options   DescribeRegistrationTypeDefinitionsPaginatorOptions
	client    DescribeRegistrationTypeDefinitionsAPIClient
	params    *DescribeRegistrationTypeDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRegistrationTypeDefinitionsPaginator returns a new
// DescribeRegistrationTypeDefinitionsPaginator
func NewDescribeRegistrationTypeDefinitionsPaginator(client DescribeRegistrationTypeDefinitionsAPIClient, params *DescribeRegistrationTypeDefinitionsInput, optFns ...func(*DescribeRegistrationTypeDefinitionsPaginatorOptions)) *DescribeRegistrationTypeDefinitionsPaginator {
	if params == nil {
		params = &DescribeRegistrationTypeDefinitionsInput{}
	}

	options := DescribeRegistrationTypeDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRegistrationTypeDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRegistrationTypeDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRegistrationTypeDefinitions page.
func (p *DescribeRegistrationTypeDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRegistrationTypeDefinitionsOutput, error) {
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
	result, err := p.client.DescribeRegistrationTypeDefinitions(ctx, &params, optFns...)
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

// DescribeRegistrationTypeDefinitionsAPIClient is a client that implements the
// DescribeRegistrationTypeDefinitions operation.
type DescribeRegistrationTypeDefinitionsAPIClient interface {
	DescribeRegistrationTypeDefinitions(context.Context, *DescribeRegistrationTypeDefinitionsInput, ...func(*Options)) (*DescribeRegistrationTypeDefinitionsOutput, error)
}

var _ DescribeRegistrationTypeDefinitionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeRegistrationTypeDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRegistrationTypeDefinitions",
	}
}