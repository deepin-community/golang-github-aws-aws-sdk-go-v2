// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound
// and outbound:
//
//   - An inbound Resolver endpoint forwards DNS queries to the DNS service for a
//     VPC from your network.
//
//   - An outbound Resolver endpoint forwards DNS queries from the DNS service for
//     a VPC to your network.
func (c *Client) CreateResolverEndpoint(ctx context.Context, params *CreateResolverEndpointInput, optFns ...func(*Options)) (*CreateResolverEndpointOutput, error) {
	if params == nil {
		params = &CreateResolverEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResolverEndpoint", params, optFns, c.addOperationCreateResolverEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverEndpointInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of running the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// Specify the applicable value:
	//
	//   - INBOUND : Resolver forwards DNS queries to the DNS service for a VPC from
	//   your network
	//
	//   - OUTBOUND : Resolver forwards DNS queries from the DNS service for a VPC to
	//   your network
	//
	// This member is required.
	Direction types.ResolverEndpointDirection

	// The subnets and IP addresses in your VPC that DNS queries originate from (for
	// outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	// The subnet ID uniquely identifies a VPC.
	//
	// Even though the minimum is 1, Route 53 requires that you create at least two.
	//
	// This member is required.
	IpAddresses []types.IpAddressRequest

	// The ID of one or more security groups that you want to use to control access to
	// this VPC. The security group that you specify must include one or more inbound
	// rules (for inbound Resolver endpoints) or outbound rules (for outbound Resolver
	// endpoints). Inbound and outbound rules must allow TCP and UDP access. For
	// inbound access, open port 53. For outbound access, open the port that you're
	// using for DNS queries on your network.
	//
	// Some security group rules will cause your connection to be tracked. For
	// outbound resolver endpoint, it can potentially impact the maximum queries per
	// second from outbound endpoint to your target name server. For inbound resolver
	// endpoint, it can bring down the overall maximum queries per second per IP
	// address to as low as 1500. To avoid connection tracking caused by security
	// group, see [Untracked connections].
	//
	// [Untracked connections]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-connection-tracking.html#untracked-connectionsl
	//
	// This member is required.
	SecurityGroupIds []string

	// A friendly name that lets you easily find a configuration in the Resolver
	// dashboard in the Route 53 console.
	Name *string

	// The Amazon Resource Name (ARN) of the Outpost. If you specify this, you must
	// also specify a value for the PreferredInstanceType .
	OutpostArn *string

	// The instance type. If you specify this, you must also specify a value for the
	// OutpostArn .
	PreferredInstanceType *string

	//  The protocols you want to use for the endpoint. DoH-FIPS is applicable for
	// inbound endpoints only.
	//
	// For an inbound endpoint you can apply the protocols as follows:
	//
	//   - Do53 and DoH in combination.
	//
	//   - Do53 and DoH-FIPS in combination.
	//
	//   - Do53 alone.
	//
	//   - DoH alone.
	//
	//   - DoH-FIPS alone.
	//
	//   - None, which is treated as Do53.
	//
	// For an outbound endpoint you can apply the protocols as follows:
	//
	//   - Do53 and DoH in combination.
	//
	//   - Do53 alone.
	//
	//   - DoH alone.
	//
	//   - None, which is treated as Do53.
	Protocols []types.Protocol

	//  For the endpoint type you can choose either IPv4, IPv6, or dual-stack. A
	// dual-stack endpoint means that it will resolve via both IPv4 and IPv6. This
	// endpoint type is applied to all IP addresses.
	ResolverEndpointType types.ResolverEndpointType

	// A list of the tag keys and values that you want to associate with the endpoint.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateResolverEndpointOutput struct {

	// Information about the CreateResolverEndpoint request, including the status of
	// the request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResolverEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResolverEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResolverEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateResolverEndpoint"); err != nil {
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
	if err = addOpCreateResolverEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolverEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResolverEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateResolverEndpoint",
	}
}