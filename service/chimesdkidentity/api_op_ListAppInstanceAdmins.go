// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the administrators in the AppInstance .
func (c *Client) ListAppInstanceAdmins(ctx context.Context, params *ListAppInstanceAdminsInput, optFns ...func(*Options)) (*ListAppInstanceAdminsOutput, error) {
	if params == nil {
		params = &ListAppInstanceAdminsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppInstanceAdmins", params, optFns, c.addOperationListAppInstanceAdminsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppInstanceAdminsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppInstanceAdminsInput struct {

	// The ARN of the AppInstance .
	//
	// This member is required.
	AppInstanceArn *string

	// The maximum number of administrators that you want to return.
	MaxResults *int32

	// The token returned from previous API requests until the number of
	// administrators is reached.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppInstanceAdminsOutput struct {

	// The information for each administrator.
	AppInstanceAdmins []types.AppInstanceAdminSummary

	// The ARN of the AppInstance .
	AppInstanceArn *string

	// The token returned from previous API requests until the number of
	// administrators is reached.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppInstanceAdminsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppInstanceAdmins{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppInstanceAdmins{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppInstanceAdmins"); err != nil {
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
	if err = addOpListAppInstanceAdminsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppInstanceAdmins(options.Region), middleware.Before); err != nil {
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

// ListAppInstanceAdminsPaginatorOptions is the paginator options for
// ListAppInstanceAdmins
type ListAppInstanceAdminsPaginatorOptions struct {
	// The maximum number of administrators that you want to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppInstanceAdminsPaginator is a paginator for ListAppInstanceAdmins
type ListAppInstanceAdminsPaginator struct {
	options   ListAppInstanceAdminsPaginatorOptions
	client    ListAppInstanceAdminsAPIClient
	params    *ListAppInstanceAdminsInput
	nextToken *string
	firstPage bool
}

// NewListAppInstanceAdminsPaginator returns a new ListAppInstanceAdminsPaginator
func NewListAppInstanceAdminsPaginator(client ListAppInstanceAdminsAPIClient, params *ListAppInstanceAdminsInput, optFns ...func(*ListAppInstanceAdminsPaginatorOptions)) *ListAppInstanceAdminsPaginator {
	if params == nil {
		params = &ListAppInstanceAdminsInput{}
	}

	options := ListAppInstanceAdminsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppInstanceAdminsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppInstanceAdminsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppInstanceAdmins page.
func (p *ListAppInstanceAdminsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppInstanceAdminsOutput, error) {
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
	result, err := p.client.ListAppInstanceAdmins(ctx, &params, optFns...)
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

// ListAppInstanceAdminsAPIClient is a client that implements the
// ListAppInstanceAdmins operation.
type ListAppInstanceAdminsAPIClient interface {
	ListAppInstanceAdmins(context.Context, *ListAppInstanceAdminsInput, ...func(*Options)) (*ListAppInstanceAdminsOutput, error)
}

var _ ListAppInstanceAdminsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAppInstanceAdmins(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppInstanceAdmins",
	}
}