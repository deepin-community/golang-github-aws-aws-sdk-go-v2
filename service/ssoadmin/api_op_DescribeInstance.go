// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the details of an instance of IAM Identity Center. The status can be
// one of the following:
//
//   - CREATE_IN_PROGRESS - The instance is in the process of being created. When
//     the instance is ready for use, DescribeInstance returns the status of ACTIVE .
//     While the instance is in the CREATE_IN_PROGRESS state, you can call only
//     DescribeInstance and DeleteInstance operations.
//
//   - DELETE_IN_PROGRESS - The instance is being deleted. Returns
//     AccessDeniedException after the delete operation completes.
//
//   - ACTIVE - The instance is active.
func (c *Client) DescribeInstance(ctx context.Context, params *DescribeInstanceInput, optFns ...func(*Options)) (*DescribeInstanceOutput, error) {
	if params == nil {
		params = &DescribeInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstance", params, optFns, c.addOperationDescribeInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceInput struct {

	// The ARN of the instance of IAM Identity Center under which the operation will
	// run.
	//
	// This member is required.
	InstanceArn *string

	noSmithyDocumentSerde
}

type DescribeInstanceOutput struct {

	// The date the instance was created.
	CreatedDate *time.Time

	// The identifier of the identity store that is connected to the instance of IAM
	// Identity Center.
	IdentityStoreId *string

	// The ARN of the instance of IAM Identity Center under which the operation will
	// run. For more information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service Namespacesin the Amazon Web Services General
	// Reference.
	InstanceArn *string

	// Specifies the instance name.
	Name *string

	// The identifier of the Amazon Web Services account for which the instance was
	// created.
	OwnerAccountId *string

	// The status of the instance.
	Status types.InstanceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstance"); err != nil {
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
	if err = addOpDescribeInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstance",
	}
}