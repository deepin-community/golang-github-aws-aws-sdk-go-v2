// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Designates the IAM role and Amazon Simple Notification Service (SNS) topic that
// Firewall Manager uses to record SNS logs.
//
// To perform this action outside of the console, you must first configure the SNS
// topic's access policy to allow the SnsRoleName to publish SNS logs. If the
// SnsRoleName provided is a role other than the AWSServiceRoleForFMS
// service-linked role, this role must have a trust relationship configured to
// allow the Firewall Manager service principal fms.amazonaws.com to assume this
// role. For information about configuring an SNS access policy, see [Service roles for Firewall Manager]in the
// Firewall Manager Developer Guide.
//
// [Service roles for Firewall Manager]: https://docs.aws.amazon.com/waf/latest/developerguide/fms-security_iam_service-with-iam.html#fms-security_iam_service-with-iam-roles-service
func (c *Client) PutNotificationChannel(ctx context.Context, params *PutNotificationChannelInput, optFns ...func(*Options)) (*PutNotificationChannelOutput, error) {
	if params == nil {
		params = &PutNotificationChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutNotificationChannel", params, optFns, c.addOperationPutNotificationChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutNotificationChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutNotificationChannelInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to record
	// Firewall Manager activity.
	//
	// This member is required.
	SnsRoleName *string

	// The Amazon Resource Name (ARN) of the SNS topic that collects notifications
	// from Firewall Manager.
	//
	// This member is required.
	SnsTopicArn *string

	noSmithyDocumentSerde
}

type PutNotificationChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutNotificationChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutNotificationChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutNotificationChannel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutNotificationChannel"); err != nil {
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
	if err = addOpPutNotificationChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutNotificationChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutNotificationChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutNotificationChannel",
	}
}