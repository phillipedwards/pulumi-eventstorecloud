// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventstorecloud

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AWSCloudWatchMetricsIntegration struct {
	pulumi.CustomResourceState

	// AWS IAM access key
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// Clusters to be used with this integration
	ClusterIds pulumi.StringArrayOutput `pulumi:"clusterIds"`
	// Human readable description of the integration
	Description pulumi.StringOutput `pulumi:"description"`
	// Name of the CloudWatch namespace
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// ID of the project to which the integration applies
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// AWS region for group
	Region pulumi.StringOutput `pulumi:"region"`
	// AWS IAM secret access key
	SecretAccessKey pulumi.StringPtrOutput `pulumi:"secretAccessKey"`
}

// NewAWSCloudWatchMetricsIntegration registers a new resource with the given unique name, arguments, and options.
func NewAWSCloudWatchMetricsIntegration(ctx *pulumi.Context,
	name string, args *AWSCloudWatchMetricsIntegrationArgs, opts ...pulumi.ResourceOption) (*AWSCloudWatchMetricsIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterIds == nil {
		return nil, errors.New("invalid value for required argument 'ClusterIds'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AWSCloudWatchMetricsIntegration
	err := ctx.RegisterResource("eventstorecloud:index/aWSCloudWatchMetricsIntegration:AWSCloudWatchMetricsIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAWSCloudWatchMetricsIntegration gets an existing AWSCloudWatchMetricsIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAWSCloudWatchMetricsIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AWSCloudWatchMetricsIntegrationState, opts ...pulumi.ResourceOption) (*AWSCloudWatchMetricsIntegration, error) {
	var resource AWSCloudWatchMetricsIntegration
	err := ctx.ReadResource("eventstorecloud:index/aWSCloudWatchMetricsIntegration:AWSCloudWatchMetricsIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AWSCloudWatchMetricsIntegration resources.
type awscloudWatchMetricsIntegrationState struct {
	// AWS IAM access key
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Clusters to be used with this integration
	ClusterIds []string `pulumi:"clusterIds"`
	// Human readable description of the integration
	Description *string `pulumi:"description"`
	// Name of the CloudWatch namespace
	Namespace *string `pulumi:"namespace"`
	// ID of the project to which the integration applies
	ProjectId *string `pulumi:"projectId"`
	// AWS region for group
	Region *string `pulumi:"region"`
	// AWS IAM secret access key
	SecretAccessKey *string `pulumi:"secretAccessKey"`
}

type AWSCloudWatchMetricsIntegrationState struct {
	// AWS IAM access key
	AccessKeyId pulumi.StringPtrInput
	// Clusters to be used with this integration
	ClusterIds pulumi.StringArrayInput
	// Human readable description of the integration
	Description pulumi.StringPtrInput
	// Name of the CloudWatch namespace
	Namespace pulumi.StringPtrInput
	// ID of the project to which the integration applies
	ProjectId pulumi.StringPtrInput
	// AWS region for group
	Region pulumi.StringPtrInput
	// AWS IAM secret access key
	SecretAccessKey pulumi.StringPtrInput
}

func (AWSCloudWatchMetricsIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*awscloudWatchMetricsIntegrationState)(nil)).Elem()
}

type awscloudWatchMetricsIntegrationArgs struct {
	// AWS IAM access key
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Clusters to be used with this integration
	ClusterIds []string `pulumi:"clusterIds"`
	// Human readable description of the integration
	Description string `pulumi:"description"`
	// Name of the CloudWatch namespace
	Namespace string `pulumi:"namespace"`
	// ID of the project to which the integration applies
	ProjectId string `pulumi:"projectId"`
	// AWS region for group
	Region string `pulumi:"region"`
	// AWS IAM secret access key
	SecretAccessKey *string `pulumi:"secretAccessKey"`
}

// The set of arguments for constructing a AWSCloudWatchMetricsIntegration resource.
type AWSCloudWatchMetricsIntegrationArgs struct {
	// AWS IAM access key
	AccessKeyId pulumi.StringPtrInput
	// Clusters to be used with this integration
	ClusterIds pulumi.StringArrayInput
	// Human readable description of the integration
	Description pulumi.StringInput
	// Name of the CloudWatch namespace
	Namespace pulumi.StringInput
	// ID of the project to which the integration applies
	ProjectId pulumi.StringInput
	// AWS region for group
	Region pulumi.StringInput
	// AWS IAM secret access key
	SecretAccessKey pulumi.StringPtrInput
}

func (AWSCloudWatchMetricsIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awscloudWatchMetricsIntegrationArgs)(nil)).Elem()
}

type AWSCloudWatchMetricsIntegrationInput interface {
	pulumi.Input

	ToAWSCloudWatchMetricsIntegrationOutput() AWSCloudWatchMetricsIntegrationOutput
	ToAWSCloudWatchMetricsIntegrationOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationOutput
}

func (*AWSCloudWatchMetricsIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**AWSCloudWatchMetricsIntegration)(nil)).Elem()
}

func (i *AWSCloudWatchMetricsIntegration) ToAWSCloudWatchMetricsIntegrationOutput() AWSCloudWatchMetricsIntegrationOutput {
	return i.ToAWSCloudWatchMetricsIntegrationOutputWithContext(context.Background())
}

func (i *AWSCloudWatchMetricsIntegration) ToAWSCloudWatchMetricsIntegrationOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AWSCloudWatchMetricsIntegrationOutput)
}

// AWSCloudWatchMetricsIntegrationArrayInput is an input type that accepts AWSCloudWatchMetricsIntegrationArray and AWSCloudWatchMetricsIntegrationArrayOutput values.
// You can construct a concrete instance of `AWSCloudWatchMetricsIntegrationArrayInput` via:
//
//          AWSCloudWatchMetricsIntegrationArray{ AWSCloudWatchMetricsIntegrationArgs{...} }
type AWSCloudWatchMetricsIntegrationArrayInput interface {
	pulumi.Input

	ToAWSCloudWatchMetricsIntegrationArrayOutput() AWSCloudWatchMetricsIntegrationArrayOutput
	ToAWSCloudWatchMetricsIntegrationArrayOutputWithContext(context.Context) AWSCloudWatchMetricsIntegrationArrayOutput
}

type AWSCloudWatchMetricsIntegrationArray []AWSCloudWatchMetricsIntegrationInput

func (AWSCloudWatchMetricsIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AWSCloudWatchMetricsIntegration)(nil)).Elem()
}

func (i AWSCloudWatchMetricsIntegrationArray) ToAWSCloudWatchMetricsIntegrationArrayOutput() AWSCloudWatchMetricsIntegrationArrayOutput {
	return i.ToAWSCloudWatchMetricsIntegrationArrayOutputWithContext(context.Background())
}

func (i AWSCloudWatchMetricsIntegrationArray) ToAWSCloudWatchMetricsIntegrationArrayOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AWSCloudWatchMetricsIntegrationArrayOutput)
}

// AWSCloudWatchMetricsIntegrationMapInput is an input type that accepts AWSCloudWatchMetricsIntegrationMap and AWSCloudWatchMetricsIntegrationMapOutput values.
// You can construct a concrete instance of `AWSCloudWatchMetricsIntegrationMapInput` via:
//
//          AWSCloudWatchMetricsIntegrationMap{ "key": AWSCloudWatchMetricsIntegrationArgs{...} }
type AWSCloudWatchMetricsIntegrationMapInput interface {
	pulumi.Input

	ToAWSCloudWatchMetricsIntegrationMapOutput() AWSCloudWatchMetricsIntegrationMapOutput
	ToAWSCloudWatchMetricsIntegrationMapOutputWithContext(context.Context) AWSCloudWatchMetricsIntegrationMapOutput
}

type AWSCloudWatchMetricsIntegrationMap map[string]AWSCloudWatchMetricsIntegrationInput

func (AWSCloudWatchMetricsIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AWSCloudWatchMetricsIntegration)(nil)).Elem()
}

func (i AWSCloudWatchMetricsIntegrationMap) ToAWSCloudWatchMetricsIntegrationMapOutput() AWSCloudWatchMetricsIntegrationMapOutput {
	return i.ToAWSCloudWatchMetricsIntegrationMapOutputWithContext(context.Background())
}

func (i AWSCloudWatchMetricsIntegrationMap) ToAWSCloudWatchMetricsIntegrationMapOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AWSCloudWatchMetricsIntegrationMapOutput)
}

type AWSCloudWatchMetricsIntegrationOutput struct{ *pulumi.OutputState }

func (AWSCloudWatchMetricsIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AWSCloudWatchMetricsIntegration)(nil)).Elem()
}

func (o AWSCloudWatchMetricsIntegrationOutput) ToAWSCloudWatchMetricsIntegrationOutput() AWSCloudWatchMetricsIntegrationOutput {
	return o
}

func (o AWSCloudWatchMetricsIntegrationOutput) ToAWSCloudWatchMetricsIntegrationOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationOutput {
	return o
}

type AWSCloudWatchMetricsIntegrationArrayOutput struct{ *pulumi.OutputState }

func (AWSCloudWatchMetricsIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AWSCloudWatchMetricsIntegration)(nil)).Elem()
}

func (o AWSCloudWatchMetricsIntegrationArrayOutput) ToAWSCloudWatchMetricsIntegrationArrayOutput() AWSCloudWatchMetricsIntegrationArrayOutput {
	return o
}

func (o AWSCloudWatchMetricsIntegrationArrayOutput) ToAWSCloudWatchMetricsIntegrationArrayOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationArrayOutput {
	return o
}

func (o AWSCloudWatchMetricsIntegrationArrayOutput) Index(i pulumi.IntInput) AWSCloudWatchMetricsIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AWSCloudWatchMetricsIntegration {
		return vs[0].([]*AWSCloudWatchMetricsIntegration)[vs[1].(int)]
	}).(AWSCloudWatchMetricsIntegrationOutput)
}

type AWSCloudWatchMetricsIntegrationMapOutput struct{ *pulumi.OutputState }

func (AWSCloudWatchMetricsIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AWSCloudWatchMetricsIntegration)(nil)).Elem()
}

func (o AWSCloudWatchMetricsIntegrationMapOutput) ToAWSCloudWatchMetricsIntegrationMapOutput() AWSCloudWatchMetricsIntegrationMapOutput {
	return o
}

func (o AWSCloudWatchMetricsIntegrationMapOutput) ToAWSCloudWatchMetricsIntegrationMapOutputWithContext(ctx context.Context) AWSCloudWatchMetricsIntegrationMapOutput {
	return o
}

func (o AWSCloudWatchMetricsIntegrationMapOutput) MapIndex(k pulumi.StringInput) AWSCloudWatchMetricsIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AWSCloudWatchMetricsIntegration {
		return vs[0].(map[string]*AWSCloudWatchMetricsIntegration)[vs[1].(string)]
	}).(AWSCloudWatchMetricsIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AWSCloudWatchMetricsIntegrationInput)(nil)).Elem(), &AWSCloudWatchMetricsIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*AWSCloudWatchMetricsIntegrationArrayInput)(nil)).Elem(), AWSCloudWatchMetricsIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AWSCloudWatchMetricsIntegrationMapInput)(nil)).Elem(), AWSCloudWatchMetricsIntegrationMap{})
	pulumi.RegisterOutputType(AWSCloudWatchMetricsIntegrationOutput{})
	pulumi.RegisterOutputType(AWSCloudWatchMetricsIntegrationArrayOutput{})
	pulumi.RegisterOutputType(AWSCloudWatchMetricsIntegrationMapOutput{})
}
