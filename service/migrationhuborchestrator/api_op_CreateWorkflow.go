// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a workflow to orchestrate your migrations.
func (c *Client) CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) {
	if params == nil {
		params = &CreateWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkflow", params, optFns, c.addOperationCreateWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkflowInput struct {

	// The input parameters required to create a migration workflow.
	//
	// This member is required.
	InputParameters map[string]types.StepInput

	// The name of the migration workflow.
	//
	// This member is required.
	Name *string

	// The ID of the template.
	//
	// This member is required.
	TemplateId *string

	// The configuration ID of the application configured in Application Discovery
	// Service.
	ApplicationConfigurationId *string

	// The description of the migration workflow.
	Description *string

	// The servers on which a step will be run.
	StepTargets []string

	// The tags to add on a migration workflow.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateWorkflowOutput struct {

	// The configuration ID of the application configured in Application Discovery
	// Service.
	AdsApplicationConfigurationId *string

	// The Amazon Resource Name (ARN) of the migration workflow.
	Arn *string

	// The time at which the migration workflow was created.
	CreationTime *time.Time

	// The description of the migration workflow.
	Description *string

	// The ID of the migration workflow.
	Id *string

	// The name of the migration workflow.
	Name *string

	// The status of the migration workflow.
	Status types.MigrationWorkflowStatusEnum

	// The servers on which a step will be run.
	StepTargets []string

	// The tags to add on a migration workflow.
	Tags map[string]string

	// The ID of the template.
	TemplateId *string

	// The inputs for creating a migration workflow.
	WorkflowInputs map[string]types.StepInput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkflow"); err != nil {
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
	if err = addOpCreateWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkflow",
	}
}