// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all your sets of featured results for a given index. Features results are
// placed above all other results for certain queries. If there's an exact match of
// a query, then one or more specific documents are featured in the search results.
func (c *Client) ListFeaturedResultsSets(ctx context.Context, params *ListFeaturedResultsSetsInput, optFns ...func(*Options)) (*ListFeaturedResultsSetsOutput, error) {
	if params == nil {
		params = &ListFeaturedResultsSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFeaturedResultsSets", params, optFns, c.addOperationListFeaturedResultsSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFeaturedResultsSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFeaturedResultsSetsInput struct {

	// The identifier of the index used for featuring results.
	//
	// This member is required.
	IndexId *string

	// The maximum number of featured results sets to return.
	MaxResults *int32

	// If the response is truncated, Amazon Kendra returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of featured
	// results sets.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFeaturedResultsSetsOutput struct {

	// An array of summary information for one or more featured results sets.
	FeaturedResultsSetSummaryItems []types.FeaturedResultsSetSummary

	// If the response is truncated, Amazon Kendra returns a pagination token in the
	// response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFeaturedResultsSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFeaturedResultsSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFeaturedResultsSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFeaturedResultsSets"); err != nil {
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
	if err = addOpListFeaturedResultsSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFeaturedResultsSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFeaturedResultsSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFeaturedResultsSets",
	}
}