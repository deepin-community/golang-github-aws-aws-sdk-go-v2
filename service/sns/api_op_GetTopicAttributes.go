// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all of the properties of a topic. Topic properties returned might
// differ based on the authorization of the user.
func (c *Client) GetTopicAttributes(ctx context.Context, params *GetTopicAttributesInput, optFns ...func(*Options)) (*GetTopicAttributesOutput, error) {
	if params == nil {
		params = &GetTopicAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTopicAttributes", params, optFns, c.addOperationGetTopicAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTopicAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetTopicAttributes action.
type GetTopicAttributesInput struct {

	// The ARN of the topic whose properties you want to get.
	//
	// This member is required.
	TopicArn *string

	noSmithyDocumentSerde
}

// Response for GetTopicAttributes action.
type GetTopicAttributesOutput struct {

	// A map of the topic's attributes. Attributes in this map include the following:
	//
	//   - DeliveryPolicy – The JSON serialization of the topic's delivery policy.
	//
	//   - DisplayName – The human-readable name used in the From field for
	//   notifications to email and email-json endpoints.
	//
	//   - EffectiveDeliveryPolicy – The JSON serialization of the effective delivery
	//   policy, taking system defaults into account.
	//
	//   - Owner – The Amazon Web Services account ID of the topic's owner.
	//
	//   - Policy – The JSON serialization of the topic's access control policy.
	//
	//   - SignatureVersion – The signature version corresponds to the hashing
	//   algorithm used while creating the signature of the notifications, subscription
	//   confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	//
	//   - By default, SignatureVersion is set to 1. The signature is a Base64-encoded
	//   SHA1withRSA signature.
	//
	//   - When you set SignatureVersion to 2. Amazon SNS uses a Base64-encoded
	//   SHA256withRSA signature.
	//
	// If the API response does not include the SignatureVersion attribute, it means
	//   that the SignatureVersion for the topic has value 1.
	//
	//   - SubscriptionsConfirmed – The number of confirmed subscriptions for the topic.
	//
	//   - SubscriptionsDeleted – The number of deleted subscriptions for the topic.
	//
	//   - SubscriptionsPending – The number of subscriptions pending confirmation for
	//   the topic.
	//
	//   - TopicArn – The topic's ARN.
	//
	//   - TracingConfig – Tracing mode of an Amazon SNS topic. By default
	//   TracingConfig is set to PassThrough , and the topic passes through the tracing
	//   header it receives from an Amazon SNS publisher to its subscriptions. If set to
	//   Active , Amazon SNS will vend X-Ray segment data to topic owner account if the
	//   sampled flag in the tracing header is true. This is only supported on standard
	//   topics.
	//
	// The following attribute applies only to [server-side-encryption]:
	//
	//   - KmsMasterKeyId - The ID of an Amazon Web Services managed customer master
	//   key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms]. For
	//   more examples, see [KeyId]in the Key Management Service API Reference.
	//
	// The following attributes apply only to [FIFO topics]:
	//
	//   - FifoTopic – When this is set to true , a FIFO topic is created.
	//
	//   - ContentBasedDeduplication – Enables content-based deduplication for FIFO
	//   topics.
	//
	//   - By default, ContentBasedDeduplication is set to false . If you create a FIFO
	//   topic and this attribute is false , you must specify a value for the
	//   MessageDeduplicationId parameter for the [Publish]action.
	//
	//   - When you set ContentBasedDeduplication to true , Amazon SNS uses a SHA-256
	//   hash to generate the MessageDeduplicationId using the body of the message (but
	//   not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the
	//   MessageDeduplicationId parameter for the Publish action.
	//
	// [Key Terms]: https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms
	// [KeyId]: https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters
	// [server-side-encryption]: https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html
	// [Publish]: https://docs.aws.amazon.com/sns/latest/api/API_Publish.html
	// [FIFO topics]: https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html
	Attributes map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTopicAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTopicAttributes"); err != nil {
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
	if err = addOpGetTopicAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTopicAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTopicAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTopicAttributes",
	}
}