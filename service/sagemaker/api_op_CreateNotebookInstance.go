// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an SageMaker notebook instance. A notebook instance is a machine
// learning (ML) compute instance running on a Jupyter notebook.
//
// In a CreateNotebookInstance request, specify the type of ML compute instance
// that you want to run. SageMaker launches the instance, installs common libraries
// that you can use to explore datasets for model training, and attaches an ML
// storage volume to the notebook instance.
//
// SageMaker also provides a set of example notebooks. Each notebook demonstrates
// how to use SageMaker with a specific algorithm or with a machine learning
// framework.
//
// After receiving the request, SageMaker does the following:
//
//   - Creates a network interface in the SageMaker VPC.
//
//   - (Option) If you specified SubnetId , SageMaker creates a network interface
//     in your own VPC, which is inferred from the subnet ID that you provide in the
//     input. When creating this network interface, SageMaker attaches the security
//     group that you specified in the request to the network interface that it creates
//     in your VPC.
//
//   - Launches an EC2 instance of the type specified in the request in the
//     SageMaker VPC. If you specified SubnetId of your VPC, SageMaker specifies both
//     network interfaces when launching this instance. This enables inbound traffic
//     from your own VPC to the notebook instance, assuming that the security groups
//     allow it.
//
// After creating the notebook instance, SageMaker returns its Amazon Resource
// Name (ARN). You can't change the name of a notebook instance after you create
// it.
//
// After SageMaker creates the notebook instance, you can connect to the Jupyter
// server and work in Jupyter notebooks. For example, you can write code to explore
// a dataset that you can use for model training, train a model, host models by
// creating SageMaker endpoints, and validate hosted models.
//
// For more information, see [How It Works].
//
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
func (c *Client) CreateNotebookInstance(ctx context.Context, params *CreateNotebookInstanceInput, optFns ...func(*Options)) (*CreateNotebookInstanceOutput, error) {
	if params == nil {
		params = &CreateNotebookInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNotebookInstance", params, optFns, c.addOperationCreateNotebookInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotebookInstanceInput struct {

	// The type of ML compute instance to launch for the notebook instance.
	//
	// This member is required.
	InstanceType types.InstanceType

	// The name of the new notebook instance.
	//
	// This member is required.
	NotebookInstanceName *string

	//  When you send any requests to Amazon Web Services resources from the notebook
	// instance, SageMaker assumes this role to perform tasks on your behalf. You must
	// grant this role necessary permissions so SageMaker can perform these tasks. The
	// policy must allow the SageMaker service principal (sagemaker.amazonaws.com)
	// permissions to assume this role. For more information, see [SageMaker Roles].
	//
	// To be able to pass this role to SageMaker, the caller of this API must have the
	// iam:PassRole permission.
	//
	// [SageMaker Roles]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html
	//
	// This member is required.
	RoleArn *string

	// A list of Elastic Inference (EI) instance types to associate with this notebook
	// instance. Currently, only one instance type can be associated with a notebook
	// instance. For more information, see [Using Elastic Inference in Amazon SageMaker].
	//
	// [Using Elastic Inference in Amazon SageMaker]: https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html
	AcceleratorTypes []types.NotebookInstanceAcceleratorType

	// An array of up to three Git repositories to associate with the notebook
	// instance. These can be either the names of Git repositories stored as resources
	// in your account, or the URL of Git repositories in [Amazon Web Services CodeCommit]or in any other Git
	// repository. These repositories are cloned at the same level as the default
	// repository of your notebook instance. For more information, see [Associating Git Repositories with SageMaker Notebook Instances].
	//
	// [Amazon Web Services CodeCommit]: https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html
	// [Associating Git Repositories with SageMaker Notebook Instances]: https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html
	AdditionalCodeRepositories []string

	// A Git repository to associate with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in [Amazon Web Services CodeCommit]or in any other Git
	// repository. When you open a notebook instance, it opens in the directory that
	// contains this repository. For more information, see [Associating Git Repositories with SageMaker Notebook Instances].
	//
	// [Amazon Web Services CodeCommit]: https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html
	// [Associating Git Repositories with SageMaker Notebook Instances]: https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html
	DefaultCodeRepository *string

	// Sets whether SageMaker provides internet access to the notebook instance. If
	// you set this to Disabled this notebook instance is able to access resources
	// only in your VPC, and is not be able to connect to SageMaker training and
	// endpoint services unless you configure a NAT Gateway in your VPC.
	//
	// For more information, see [Notebook Instances Are Internet-Enabled by Default]. You can set the value of this parameter to Disabled
	// only if you set a value for the SubnetId parameter.
	//
	// [Notebook Instances Are Internet-Enabled by Default]: https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access
	DirectInternetAccess types.DirectInternetAccess

	// Information on the IMDS configuration of the notebook instance
	InstanceMetadataServiceConfiguration *types.InstanceMetadataServiceConfiguration

	// The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service
	// key that SageMaker uses to encrypt data on the storage volume attached to your
	// notebook instance. The KMS key you provide must be enabled. For information, see
	// [Enabling and Disabling Keys]in the Amazon Web Services Key Management Service Developer Guide.
	//
	// [Enabling and Disabling Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html
	KmsKeyId *string

	// The name of a lifecycle configuration to associate with the notebook instance.
	// For information about lifestyle configurations, see [Step 2.1: (Optional) Customize a Notebook Instance].
	//
	// [Step 2.1: (Optional) Customize a Notebook Instance]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
	LifecycleConfigName *string

	// The platform identifier of the notebook instance runtime environment.
	PlatformIdentifier *string

	// Whether root access is enabled or disabled for users of the notebook instance.
	// The default value is Enabled .
	//
	// Lifecycle configurations need root access to be able to set up a notebook
	// instance. Because of this, lifecycle configurations associated with a notebook
	// instance always run with root access even if you disable root access for users.
	RootAccess types.RootAccess

	// The VPC security group IDs, in the form sg-xxxxxxxx. The security groups must
	// be for the same VPC as specified in the subnet.
	SecurityGroupIds []string

	// The ID of the subnet in a VPC to which you would like to have a connectivity
	// from your ML compute instance.
	SubnetId *string

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see [Tagging Amazon Web Services Resources].
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	// The default value is 5 GB.
	VolumeSizeInGB *int32

	noSmithyDocumentSerde
}

type CreateNotebookInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the notebook instance.
	NotebookInstanceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNotebookInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNotebookInstance"); err != nil {
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
	if err = addOpCreateNotebookInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotebookInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNotebookInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNotebookInstance",
	}
}