// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the Amazon EMR block public access configuration for your Amazon Web
// Services account in the current Region. For more information see [Configure Block Public Access for Amazon EMR]in the Amazon
// EMR Management Guide.
//
// [Configure Block Public Access for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/configure-block-public-access.html
func (c *Client) GetBlockPublicAccessConfiguration(ctx context.Context, params *GetBlockPublicAccessConfigurationInput, optFns ...func(*Options)) (*GetBlockPublicAccessConfigurationOutput, error) {
	if params == nil {
		params = &GetBlockPublicAccessConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBlockPublicAccessConfiguration", params, optFns, c.addOperationGetBlockPublicAccessConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBlockPublicAccessConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBlockPublicAccessConfigurationInput struct {
	noSmithyDocumentSerde
}

type GetBlockPublicAccessConfigurationOutput struct {

	// A configuration for Amazon EMR block public access. The configuration applies
	// to all clusters created in your account for the current Region. The
	// configuration specifies whether block public access is enabled. If block public
	// access is enabled, security groups associated with the cluster cannot have rules
	// that allow inbound traffic from 0.0.0.0/0 or ::/0 on a port, unless the port is
	// specified as an exception using PermittedPublicSecurityGroupRuleRanges in the
	// BlockPublicAccessConfiguration . By default, Port 22 (SSH) is an exception, and
	// public access is allowed on this port. You can change this by updating the block
	// public access configuration to remove the exception.
	//
	// For accounts that created clusters in a Region before November 25, 2019, block
	// public access is disabled by default in that Region. To use this feature, you
	// must manually enable and configure it. For accounts that did not create an
	// Amazon EMR cluster in a Region before this date, block public access is enabled
	// by default in that Region.
	//
	// This member is required.
	BlockPublicAccessConfiguration *types.BlockPublicAccessConfiguration

	// Properties that describe the Amazon Web Services principal that created the
	// BlockPublicAccessConfiguration using the PutBlockPublicAccessConfiguration
	// action as well as the date and time that the configuration was created. Each
	// time a configuration for block public access is updated, Amazon EMR updates this
	// metadata.
	//
	// This member is required.
	BlockPublicAccessConfigurationMetadata *types.BlockPublicAccessConfigurationMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBlockPublicAccessConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBlockPublicAccessConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBlockPublicAccessConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBlockPublicAccessConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlockPublicAccessConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBlockPublicAccessConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBlockPublicAccessConfiguration",
	}
}