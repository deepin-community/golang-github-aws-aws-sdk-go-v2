// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a deployment. Deployments define the components that run on Greengrass
// core devices.
func (c *Client) GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) {
	if params == nil {
		params = &GetDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeployment", params, optFns, c.addOperationGetDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentInput struct {

	// The ID of the deployment.
	//
	// This member is required.
	DeploymentId *string

	noSmithyDocumentSerde
}

type GetDeploymentOutput struct {

	// The components to deploy. This is a dictionary, where each key is the name of a
	// component, and each key's value is the version and configuration to deploy for
	// that component.
	Components map[string]types.ComponentDeploymentSpecification

	// The time at which the deployment was created, expressed in ISO 8601 format.
	CreationTimestamp *time.Time

	// The ID of the deployment.
	DeploymentId *string

	// The name of the deployment.
	DeploymentName *string

	// The deployment policies for the deployment. These policies define how the
	// deployment updates components and handles failure.
	DeploymentPolicies *types.DeploymentPolicies

	// The status of the deployment.
	DeploymentStatus types.DeploymentStatus

	// The [ARN] of the IoT job that applies the deployment to target devices.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	IotJobArn *string

	// The job configuration for the deployment configuration. The job configuration
	// specifies the rollout, timeout, and stop configurations for the deployment
	// configuration.
	IotJobConfiguration *types.DeploymentIoTJobConfiguration

	// The ID of the IoT job that applies the deployment to target devices.
	IotJobId *string

	// Whether or not the deployment is the latest revision for its target.
	IsLatestForTarget bool

	// The parent deployment's target [ARN] within a subdeployment.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	ParentTargetArn *string

	// The revision number of the deployment.
	RevisionId *string

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see [Tag your resources]in the IoT Greengrass V2 Developer Guide.
	//
	// [Tag your resources]: https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html
	Tags map[string]string

	// The [ARN] of the target IoT thing or thing group.
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	TargetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDeployment"); err != nil {
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
	if err = addOpGetDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDeployment",
	}
}