// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the cross-account attachments that have been created in Global Accelerator.
func (c *Client) ListCrossAccountAttachments(ctx context.Context, params *ListCrossAccountAttachmentsInput, optFns ...func(*Options)) (*ListCrossAccountAttachmentsOutput, error) {
	if params == nil {
		params = &ListCrossAccountAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCrossAccountAttachments", params, optFns, c.addOperationListCrossAccountAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCrossAccountAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCrossAccountAttachmentsInput struct {

	// The number of cross-account attachment objects that you want to return with
	// this call. The default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCrossAccountAttachmentsOutput struct {

	// Information about the cross-account attachments.
	CrossAccountAttachments []types.Attachment

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCrossAccountAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCrossAccountAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCrossAccountAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCrossAccountAttachments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCrossAccountAttachments(options.Region), middleware.Before); err != nil {
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

// ListCrossAccountAttachmentsPaginatorOptions is the paginator options for
// ListCrossAccountAttachments
type ListCrossAccountAttachmentsPaginatorOptions struct {
	// The number of cross-account attachment objects that you want to return with
	// this call. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCrossAccountAttachmentsPaginator is a paginator for
// ListCrossAccountAttachments
type ListCrossAccountAttachmentsPaginator struct {
	options   ListCrossAccountAttachmentsPaginatorOptions
	client    ListCrossAccountAttachmentsAPIClient
	params    *ListCrossAccountAttachmentsInput
	nextToken *string
	firstPage bool
}

// NewListCrossAccountAttachmentsPaginator returns a new
// ListCrossAccountAttachmentsPaginator
func NewListCrossAccountAttachmentsPaginator(client ListCrossAccountAttachmentsAPIClient, params *ListCrossAccountAttachmentsInput, optFns ...func(*ListCrossAccountAttachmentsPaginatorOptions)) *ListCrossAccountAttachmentsPaginator {
	if params == nil {
		params = &ListCrossAccountAttachmentsInput{}
	}

	options := ListCrossAccountAttachmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCrossAccountAttachmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCrossAccountAttachmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCrossAccountAttachments page.
func (p *ListCrossAccountAttachmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCrossAccountAttachmentsOutput, error) {
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
	result, err := p.client.ListCrossAccountAttachments(ctx, &params, optFns...)
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

// ListCrossAccountAttachmentsAPIClient is a client that implements the
// ListCrossAccountAttachments operation.
type ListCrossAccountAttachmentsAPIClient interface {
	ListCrossAccountAttachments(context.Context, *ListCrossAccountAttachmentsInput, ...func(*Options)) (*ListCrossAccountAttachmentsOutput, error)
}

var _ ListCrossAccountAttachmentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCrossAccountAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCrossAccountAttachments",
	}
}