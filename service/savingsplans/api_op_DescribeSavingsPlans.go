// Code generated by smithy-go-codegen DO NOT EDIT.

package savingsplans

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified Savings Plans.
func (c *Client) DescribeSavingsPlans(ctx context.Context, params *DescribeSavingsPlansInput, optFns ...func(*Options)) (*DescribeSavingsPlansOutput, error) {
	if params == nil {
		params = &DescribeSavingsPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSavingsPlans", params, optFns, c.addOperationDescribeSavingsPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSavingsPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSavingsPlansInput struct {

	// The filters.
	Filters []types.SavingsPlanFilter

	// The maximum number of results to return with a single call. To retrieve
	// additional results, make another call with the returned token value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Names (ARN) of the Savings Plans.
	SavingsPlanArns []string

	// The IDs of the Savings Plans.
	SavingsPlanIds []string

	// The current states of the Savings Plans.
	States []types.SavingsPlanState

	noSmithyDocumentSerde
}

type DescribeSavingsPlansOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the Savings Plans.
	SavingsPlans []types.SavingsPlan

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSavingsPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSavingsPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSavingsPlans{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSavingsPlans"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSavingsPlans(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSavingsPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSavingsPlans",
	}
}