// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the projected metrics of Amazon RDS recommendations.
func (c *Client) GetRDSDatabaseRecommendationProjectedMetrics(ctx context.Context, params *GetRDSDatabaseRecommendationProjectedMetricsInput, optFns ...func(*Options)) (*GetRDSDatabaseRecommendationProjectedMetricsOutput, error) {
	if params == nil {
		params = &GetRDSDatabaseRecommendationProjectedMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRDSDatabaseRecommendationProjectedMetrics", params, optFns, c.addOperationGetRDSDatabaseRecommendationProjectedMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRDSDatabaseRecommendationProjectedMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRDSDatabaseRecommendationProjectedMetricsInput struct {

	//  The timestamp of the last projected metrics data point to return.
	//
	// This member is required.
	EndTime *time.Time

	//  The granularity, in seconds, of the projected metrics data points.
	//
	// This member is required.
	Period int32

	//  The ARN that identifies the Amazon RDS.
	//
	// The following is the format of the ARN:
	//
	//     arn:aws:rds:{region}:{accountId}:db:{resourceName}
	//
	// This member is required.
	ResourceArn *string

	//  The timestamp of the first projected metrics data point to return.
	//
	// This member is required.
	StartTime *time.Time

	//  The statistic of the projected metrics.
	//
	// This member is required.
	Stat types.MetricStatistic

	// Describes the recommendation preferences to return in the response of a GetAutoScalingGroupRecommendations, GetEC2InstanceRecommendations, GetEC2RecommendationProjectedMetrics, GetRDSDatabaseRecommendations,
	// and GetRDSDatabaseRecommendationProjectedMetricsrequest.
	RecommendationPreferences *types.RecommendationPreferences

	noSmithyDocumentSerde
}

type GetRDSDatabaseRecommendationProjectedMetricsOutput struct {

	//  An array of objects that describes the projected metrics.
	RecommendedOptionProjectedMetrics []types.RDSDatabaseRecommendedOptionProjectedMetric

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRDSDatabaseRecommendationProjectedMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRDSDatabaseRecommendationProjectedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRDSDatabaseRecommendationProjectedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRDSDatabaseRecommendationProjectedMetrics"); err != nil {
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
	if err = addOpGetRDSDatabaseRecommendationProjectedMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRDSDatabaseRecommendationProjectedMetrics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRDSDatabaseRecommendationProjectedMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRDSDatabaseRecommendationProjectedMetrics",
	}
}