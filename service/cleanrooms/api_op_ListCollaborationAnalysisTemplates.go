// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists analysis templates within a collaboration.
func (c *Client) ListCollaborationAnalysisTemplates(ctx context.Context, params *ListCollaborationAnalysisTemplatesInput, optFns ...func(*Options)) (*ListCollaborationAnalysisTemplatesOutput, error) {
	if params == nil {
		params = &ListCollaborationAnalysisTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollaborationAnalysisTemplates", params, optFns, c.addOperationListCollaborationAnalysisTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollaborationAnalysisTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollaborationAnalysisTemplatesInput struct {

	// A unique identifier for the collaboration that the analysis templates belong
	// to. Currently accepts collaboration ID.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollaborationAnalysisTemplatesOutput struct {

	// The metadata of the analysis template within a collaboration.
	//
	// This member is required.
	CollaborationAnalysisTemplateSummaries []types.CollaborationAnalysisTemplateSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollaborationAnalysisTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCollaborationAnalysisTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCollaborationAnalysisTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCollaborationAnalysisTemplates"); err != nil {
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
	if err = addOpListCollaborationAnalysisTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollaborationAnalysisTemplates(options.Region), middleware.Before); err != nil {
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

// ListCollaborationAnalysisTemplatesPaginatorOptions is the paginator options for
// ListCollaborationAnalysisTemplates
type ListCollaborationAnalysisTemplatesPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollaborationAnalysisTemplatesPaginator is a paginator for
// ListCollaborationAnalysisTemplates
type ListCollaborationAnalysisTemplatesPaginator struct {
	options   ListCollaborationAnalysisTemplatesPaginatorOptions
	client    ListCollaborationAnalysisTemplatesAPIClient
	params    *ListCollaborationAnalysisTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListCollaborationAnalysisTemplatesPaginator returns a new
// ListCollaborationAnalysisTemplatesPaginator
func NewListCollaborationAnalysisTemplatesPaginator(client ListCollaborationAnalysisTemplatesAPIClient, params *ListCollaborationAnalysisTemplatesInput, optFns ...func(*ListCollaborationAnalysisTemplatesPaginatorOptions)) *ListCollaborationAnalysisTemplatesPaginator {
	if params == nil {
		params = &ListCollaborationAnalysisTemplatesInput{}
	}

	options := ListCollaborationAnalysisTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollaborationAnalysisTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollaborationAnalysisTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCollaborationAnalysisTemplates page.
func (p *ListCollaborationAnalysisTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollaborationAnalysisTemplatesOutput, error) {
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
	result, err := p.client.ListCollaborationAnalysisTemplates(ctx, &params, optFns...)
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

// ListCollaborationAnalysisTemplatesAPIClient is a client that implements the
// ListCollaborationAnalysisTemplates operation.
type ListCollaborationAnalysisTemplatesAPIClient interface {
	ListCollaborationAnalysisTemplates(context.Context, *ListCollaborationAnalysisTemplatesInput, ...func(*Options)) (*ListCollaborationAnalysisTemplatesOutput, error)
}

var _ ListCollaborationAnalysisTemplatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCollaborationAnalysisTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCollaborationAnalysisTemplates",
	}
}