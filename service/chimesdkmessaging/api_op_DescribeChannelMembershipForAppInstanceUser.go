// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns the details of a channel based on the membership of the specified
//
// AppInstanceUser or AppInstanceBot .
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func (c *Client) DescribeChannelMembershipForAppInstanceUser(ctx context.Context, params *DescribeChannelMembershipForAppInstanceUserInput, optFns ...func(*Options)) (*DescribeChannelMembershipForAppInstanceUserOutput, error) {
	if params == nil {
		params = &DescribeChannelMembershipForAppInstanceUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChannelMembershipForAppInstanceUser", params, optFns, c.addOperationDescribeChannelMembershipForAppInstanceUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChannelMembershipForAppInstanceUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChannelMembershipForAppInstanceUserInput struct {

	// The ARN of the user or bot in a channel.
	//
	// This member is required.
	AppInstanceUserArn *string

	// The ARN of the channel to which the user belongs.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	noSmithyDocumentSerde
}

type DescribeChannelMembershipForAppInstanceUserOutput struct {

	// The channel to which a user belongs.
	ChannelMembership *types.ChannelMembershipForAppInstanceUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChannelMembershipForAppInstanceUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannelMembershipForAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannelMembershipForAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeChannelMembershipForAppInstanceUser"); err != nil {
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
	if err = addOpDescribeChannelMembershipForAppInstanceUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannelMembershipForAppInstanceUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeChannelMembershipForAppInstanceUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeChannelMembershipForAppInstanceUser",
	}
}