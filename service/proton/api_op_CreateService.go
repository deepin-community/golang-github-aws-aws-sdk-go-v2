// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an Proton service. An Proton service is an instantiation of a service
// template and often includes several service instances and pipeline. For more
// information, see [Services]in the Proton User Guide.
//
// [Services]: https://docs.aws.amazon.com/proton/latest/userguide/ag-services.html
func (c *Client) CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) {
	if params == nil {
		params = &CreateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateService", params, optFns, c.addOperationCreateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceInput struct {

	// The service name.
	//
	// This member is required.
	Name *string

	// A link to a spec file that provides inputs as defined in the service template
	// bundle schema file. The spec file is in YAML format. Don’t include pipeline
	// inputs in the spec if your service template doesn’t include a service pipeline.
	// For more information, see [Create a service]in the Proton User Guide.
	//
	// [Create a service]: https://docs.aws.amazon.com/proton/latest/userguide/ag-create-svc.html
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	Spec *string

	// The major version of the service template that was used to create the service.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The name of the service template that's used to create the service.
	//
	// This member is required.
	TemplateName *string

	// The name of the code repository branch that holds the code that's deployed in
	// Proton. Don't include this parameter if your service template doesn't include a
	// service pipeline.
	BranchName *string

	// A description of the Proton service.
	Description *string

	// The Amazon Resource Name (ARN) of the repository connection. For more
	// information, see [Setting up an AWS CodeStar connection]in the Proton User Guide. Don't include this parameter if your
	// service template doesn't include a service pipeline.
	//
	// [Setting up an AWS CodeStar connection]: https://docs.aws.amazon.com/proton/latest/userguide/setting-up-for-service.html#setting-up-vcontrol
	RepositoryConnectionArn *string

	// The ID of the code repository. Don't include this parameter if your service
	// template doesn't include a service pipeline.
	RepositoryId *string

	// An optional list of metadata items that you can associate with the Proton
	// service. A tag is a key-value pair.
	//
	// For more information, see [Proton resources and tagging] in the Proton User Guide.
	//
	// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
	Tags []types.Tag

	// The minor version of the service template that was used to create the service.
	TemplateMinorVersion *string

	noSmithyDocumentSerde
}

type CreateServiceOutput struct {

	// The service detail data that's returned by Proton.
	//
	// This member is required.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateService"); err != nil {
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
	if err = addOpCreateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateService",
	}
}