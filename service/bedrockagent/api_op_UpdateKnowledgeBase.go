// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of a knowledge base with the fields that you specify.
// Because all fields will be overwritten, you must include the same values for
// fields that you want to keep the same.
//
// You can change the following fields:
//
//   - name
//
//   - description
//
//   - roleArn
//
// You can't change the knowledgeBaseConfiguration or storageConfiguration fields,
// so you must specify the same configurations as when you created the knowledge
// base. You can send a [GetKnowledgeBase]request and copy the same configurations.
//
// [GetKnowledgeBase]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetKnowledgeBase.html
func (c *Client) UpdateKnowledgeBase(ctx context.Context, params *UpdateKnowledgeBaseInput, optFns ...func(*Options)) (*UpdateKnowledgeBaseOutput, error) {
	if params == nil {
		params = &UpdateKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKnowledgeBase", params, optFns, c.addOperationUpdateKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKnowledgeBaseInput struct {

	// Specifies the configuration for the embeddings model used for the knowledge
	// base. You must use the same configuration as when the knowledge base was
	// created.
	//
	// This member is required.
	KnowledgeBaseConfiguration *types.KnowledgeBaseConfiguration

	// The unique identifier of the knowledge base to update.
	//
	// This member is required.
	KnowledgeBaseId *string

	// Specifies a new name for the knowledge base.
	//
	// This member is required.
	Name *string

	// Specifies a different Amazon Resource Name (ARN) of the IAM role with
	// permissions to invoke API operations on the knowledge base.
	//
	// This member is required.
	RoleArn *string

	// Specifies the configuration for the vector store used for the knowledge base.
	// You must use the same configuration as when the knowledge base was created.
	//
	// This member is required.
	StorageConfiguration *types.StorageConfiguration

	// Specifies a new description for the knowledge base.
	Description *string

	noSmithyDocumentSerde
}

type UpdateKnowledgeBaseOutput struct {

	// Contains details about the knowledge base.
	//
	// This member is required.
	KnowledgeBase *types.KnowledgeBase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateKnowledgeBase"); err != nil {
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
	if err = addOpUpdateKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKnowledgeBase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateKnowledgeBase",
	}
}