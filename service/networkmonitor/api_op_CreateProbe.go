// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a probe within a monitor. Once you create a probe, and it begins
// monitoring your network traffic, you'll incur billing charges for that probe.
// This action requires the monitorName parameter. Run ListMonitors to get a list
// of monitor names. Note the name of the monitorName you want to create the probe
// for.
func (c *Client) CreateProbe(ctx context.Context, params *CreateProbeInput, optFns ...func(*Options)) (*CreateProbeOutput, error) {
	if params == nil {
		params = &CreateProbeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProbe", params, optFns, c.addOperationCreateProbeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProbeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProbeInput struct {

	// The name of the monitor to associated with the probe.
	//
	// This member is required.
	MonitorName *string

	// Describes the details of an individual probe for a monitor.
	//
	// This member is required.
	Probe *types.ProbeInput

	// Unique, case-sensitive identifier to ensure the idempotency of the request.
	// Only returned if a client token was provided in the request.
	ClientToken *string

	// The list of key-value pairs created and assigned to the probe.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateProbeOutput struct {

	// The destination IP address for the monitor. This must be either an IPv4 or IPv6
	// address.
	//
	// This member is required.
	Destination *string

	// The protocol used for the network traffic between the source and destination .
	// This must be either TCP or ICMP .
	//
	// This member is required.
	Protocol types.Protocol

	// The ARN of the probe.
	//
	// This member is required.
	SourceArn *string

	// Indicates whether the IP address is IPV4 or IPV6 .
	AddressFamily types.AddressFamily

	// The time and date that the probe was created.
	CreatedAt *time.Time

	// The port associated with the destination . This is required only if the protocol
	// is TCP and must be a number between 1 and 65536 .
	DestinationPort *int32

	// The time and date when the probe was last modified.
	ModifiedAt *time.Time

	// The size of the packets sent between the source and destination. This must be a
	// number between 56 and 8500 .
	PacketSize *int32

	// The ARN of the probe.
	ProbeArn *string

	// The ID of the probe for which details are returned.
	ProbeId *string

	// The state of the probe.
	State types.ProbeState

	// The list of key-value pairs assigned to the probe.
	Tags map[string]string

	// The ID of the source VPC or subnet.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProbeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProbe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProbe{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProbe"); err != nil {
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
	if err = addIdempotencyToken_opCreateProbeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProbeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProbe(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateProbe struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProbe) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProbe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProbeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProbeInput ")
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
func addIdempotencyToken_opCreateProbeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateProbe{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProbe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProbe",
	}
}