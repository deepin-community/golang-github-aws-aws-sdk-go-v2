// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Lists all channels that a particular AppInstanceUser is a part of. Only an
//
// AppInstanceAdmin can call the API with a user ARN that is not their own.
//
// The x-amz-chime-bearer request header is mandatory. Use the AppInstanceUserArn
// of the user that makes the API call as the value in the header.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [ListChannelMembershipsForAppInstanceUser], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by ListChannelMembershipsForAppInstanceUser in the Amazon
// Chime SDK Messaging Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [ListChannelMembershipsForAppInstanceUser]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_messaging-chime_ListChannelMembershipsForAppInstanceUser.html
func (c *Client) ListChannelMembershipsForAppInstanceUser(ctx context.Context, params *ListChannelMembershipsForAppInstanceUserInput, optFns ...func(*Options)) (*ListChannelMembershipsForAppInstanceUserOutput, error) {
	if params == nil {
		params = &ListChannelMembershipsForAppInstanceUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannelMembershipsForAppInstanceUser", params, optFns, c.addOperationListChannelMembershipsForAppInstanceUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelMembershipsForAppInstanceUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelMembershipsForAppInstanceUserInput struct {

	// The ARN of the AppInstanceUser s
	AppInstanceUserArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	// The maximum number of users that you want returned.
	MaxResults *int32

	// The token returned from previous API requests until the number of channel
	// memberships is reached.
	NextToken *string

	noSmithyDocumentSerde
}

type ListChannelMembershipsForAppInstanceUserOutput struct {

	// The information for the requested channel memberships.
	ChannelMemberships []types.ChannelMembershipForAppInstanceUserSummary

	// The token passed by previous API calls until all requested users are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelMembershipsForAppInstanceUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannelMembershipsForAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannelMembershipsForAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListChannelMembershipsForAppInstanceUser"); err != nil {
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
	if err = addEndpointPrefix_opListChannelMembershipsForAppInstanceUserMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannelMembershipsForAppInstanceUser(options.Region), middleware.Before); err != nil {
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

// ListChannelMembershipsForAppInstanceUserPaginatorOptions is the paginator
// options for ListChannelMembershipsForAppInstanceUser
type ListChannelMembershipsForAppInstanceUserPaginatorOptions struct {
	// The maximum number of users that you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelMembershipsForAppInstanceUserPaginator is a paginator for
// ListChannelMembershipsForAppInstanceUser
type ListChannelMembershipsForAppInstanceUserPaginator struct {
	options   ListChannelMembershipsForAppInstanceUserPaginatorOptions
	client    ListChannelMembershipsForAppInstanceUserAPIClient
	params    *ListChannelMembershipsForAppInstanceUserInput
	nextToken *string
	firstPage bool
}

// NewListChannelMembershipsForAppInstanceUserPaginator returns a new
// ListChannelMembershipsForAppInstanceUserPaginator
func NewListChannelMembershipsForAppInstanceUserPaginator(client ListChannelMembershipsForAppInstanceUserAPIClient, params *ListChannelMembershipsForAppInstanceUserInput, optFns ...func(*ListChannelMembershipsForAppInstanceUserPaginatorOptions)) *ListChannelMembershipsForAppInstanceUserPaginator {
	if params == nil {
		params = &ListChannelMembershipsForAppInstanceUserInput{}
	}

	options := ListChannelMembershipsForAppInstanceUserPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelMembershipsForAppInstanceUserPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelMembershipsForAppInstanceUserPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannelMembershipsForAppInstanceUser page.
func (p *ListChannelMembershipsForAppInstanceUserPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelMembershipsForAppInstanceUserOutput, error) {
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
	result, err := p.client.ListChannelMembershipsForAppInstanceUser(ctx, &params, optFns...)
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

type endpointPrefix_opListChannelMembershipsForAppInstanceUserMiddleware struct {
}

func (*endpointPrefix_opListChannelMembershipsForAppInstanceUserMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListChannelMembershipsForAppInstanceUserMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListChannelMembershipsForAppInstanceUserMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListChannelMembershipsForAppInstanceUserMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListChannelMembershipsForAppInstanceUserAPIClient is a client that implements
// the ListChannelMembershipsForAppInstanceUser operation.
type ListChannelMembershipsForAppInstanceUserAPIClient interface {
	ListChannelMembershipsForAppInstanceUser(context.Context, *ListChannelMembershipsForAppInstanceUserInput, ...func(*Options)) (*ListChannelMembershipsForAppInstanceUserOutput, error)
}

var _ ListChannelMembershipsForAppInstanceUserAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListChannelMembershipsForAppInstanceUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListChannelMembershipsForAppInstanceUser",
	}
}