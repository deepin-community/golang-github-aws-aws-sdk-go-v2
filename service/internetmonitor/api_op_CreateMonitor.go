// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a monitor in Amazon CloudWatch Internet Monitor. A monitor is built
// based on information from the application resources that you add: VPCs, Network
// Load Balancers (NLBs), Amazon CloudFront distributions, and Amazon WorkSpaces
// directories. Internet Monitor then publishes internet measurements from Amazon
// Web Services that are specific to the city-networks. That is, the locations and
// ASNs (typically internet service providers or ISPs), where clients access your
// application. For more information, see [Using Amazon CloudWatch Internet Monitor]in the Amazon CloudWatch User Guide.
//
// When you create a monitor, you choose the percentage of traffic that you want
// to monitor. You can also set a maximum limit for the number of city-networks
// where client traffic is monitored, that caps the total traffic that Internet
// Monitor monitors. A city-network maximum is the limit of city-networks, but you
// only pay for the number of city-networks that are actually monitored. You can
// update your monitor at any time to change the percentage of traffic to monitor
// or the city-networks maximum. For more information, see [Choosing a city-network maximum value]in the Amazon
// CloudWatch User Guide.
//
// [Using Amazon CloudWatch Internet Monitor]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html
// [Choosing a city-network maximum value]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html
func (c *Client) CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) {
	if params == nil {
		params = &CreateMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMonitor", params, optFns, c.addOperationCreateMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMonitorInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// A unique, case-sensitive string of up to 64 ASCII characters that you specify
	// to make an idempotent API request. Don't reuse the same client token for other
	// API requests.
	ClientToken *string

	// Defines the threshold percentages and other configuration information for when
	// Amazon CloudWatch Internet Monitor creates a health event. Internet Monitor
	// creates a health event when an internet issue that affects your application end
	// users has a health score percentage that is at or below a specific threshold,
	// and, sometimes, when other criteria are met.
	//
	// If you don't set a health event threshold, the default value is 95%.
	//
	// For more information, see [Change health event thresholds] in the Internet Monitor section of the CloudWatch
	// User Guide.
	//
	// [Change health event thresholds]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview
	HealthEventsConfig *types.HealthEventsConfig

	// Publish internet measurements for Internet Monitor to an Amazon S3 bucket in
	// addition to CloudWatch Logs.
	InternetMeasurementsLogDelivery *types.InternetMeasurementsLogDelivery

	// The maximum number of city-networks to monitor for your resources. A
	// city-network is the location (city) where clients access your application
	// resources from and the ASN or network provider, such as an internet service
	// provider (ISP), that clients access the resources through. Setting this limit
	// can help control billing costs.
	//
	// To learn more, see [Choosing a city-network maximum value] in the Amazon CloudWatch Internet Monitor section of the
	// CloudWatch User Guide.
	//
	// [Choosing a city-network maximum value]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html
	MaxCityNetworksToMonitor *int32

	// The resources to include in a monitor, which you provide as a set of Amazon
	// Resource Names (ARNs). Resources can be VPCs, NLBs, Amazon CloudFront
	// distributions, or Amazon WorkSpaces directories.
	//
	// You can add a combination of VPCs and CloudFront distributions, or you can add
	// WorkSpaces directories, or you can add NLBs. You can't add NLBs or WorkSpaces
	// directories together with any other resources.
	//
	// If you add only Amazon VPC resources, at least one VPC must have an Internet
	// Gateway attached to it, to make sure that it has internet connectivity.
	Resources []string

	// The tags for a monitor. You can add a maximum of 50 tags in Internet Monitor.
	Tags map[string]string

	// The percentage of the internet-facing traffic for your application that you
	// want to monitor with this monitor. If you set a city-networks maximum, that
	// limit overrides the traffic percentage that you set.
	//
	// To learn more, see [Choosing an application traffic percentage to monitor] in the Amazon CloudWatch Internet Monitor section of the
	// CloudWatch User Guide.
	//
	// [Choosing an application traffic percentage to monitor]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMTrafficPercentage.html
	TrafficPercentageToMonitor *int32

	noSmithyDocumentSerde
}

type CreateMonitorOutput struct {

	// The Amazon Resource Name (ARN) of the monitor.
	//
	// This member is required.
	Arn *string

	// The status of a monitor.
	//
	// This member is required.
	Status types.MonitorConfigState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMonitor"); err != nil {
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
	if err = addIdempotencyToken_opCreateMonitorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMonitor(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMonitor struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMonitor) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMonitorInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMonitorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMonitor{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMonitor",
	}
}