// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the route calculator resource details.
func (c *Client) DescribeRouteCalculator(ctx context.Context, params *DescribeRouteCalculatorInput, optFns ...func(*Options)) (*DescribeRouteCalculatorOutput, error) {
	if params == nil {
		params = &DescribeRouteCalculatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRouteCalculator", params, optFns, c.addOperationDescribeRouteCalculatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRouteCalculatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRouteCalculatorInput struct {

	// The name of the route calculator resource.
	//
	// This member is required.
	CalculatorName *string

	noSmithyDocumentSerde
}

type DescribeRouteCalculatorOutput struct {

	// The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN
	// when you specify a resource across Amazon Web Services.
	//
	//   - Format example:
	//   arn:aws:geo:region:account-id:route-calculator/ExampleCalculator
	//
	// This member is required.
	CalculatorArn *string

	// The name of the route calculator resource being described.
	//
	// This member is required.
	CalculatorName *string

	// The timestamp when the route calculator resource was created in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ .
	//
	//   - For example, 2020–07-2T12:15:20.000Z+01:00
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	CreateTime *time.Time

	// The data provider of traffic and road network data. Indicates one of the
	// available providers:
	//
	//   - Esri
	//
	//   - Grab
	//
	//   - Here
	//
	// For more information about data providers, see [Amazon Location Service data providers].
	//
	// [Amazon Location Service data providers]: https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html
	//
	// This member is required.
	DataSource *string

	// The optional description of the route calculator resource.
	//
	// This member is required.
	Description *string

	// The timestamp when the route calculator resource was last updated in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ .
	//
	//   - For example, 2020–07-2T12:15:20.000Z+01:00
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	UpdateTime *time.Time

	// Always returns RequestBasedUsage .
	//
	// Deprecated: Deprecated. Always returns RequestBasedUsage.
	PricingPlan types.PricingPlan

	// Tags associated with route calculator resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRouteCalculatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRouteCalculator"); err != nil {
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
	if err = addEndpointPrefix_opDescribeRouteCalculatorMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeRouteCalculatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRouteCalculator(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeRouteCalculatorMiddleware struct {
}

func (*endpointPrefix_opDescribeRouteCalculatorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeRouteCalculatorMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.routes." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDescribeRouteCalculatorMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeRouteCalculatorMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRouteCalculator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRouteCalculator",
	}
}