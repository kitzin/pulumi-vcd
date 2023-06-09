// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxtAlbCloud struct {
	pulumi.CustomResourceState

	// NSX-T ALB Controller ID
	ControllerId pulumi.StringOutput `pulumi:"controllerId"`
	// NSX-T ALB Cloud description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// NSX-T ALB Cloud detailed health message
	HealthMessage pulumi.StringOutput `pulumi:"healthMessage"`
	// NSX-T ALB Cloud health status
	HealthStatus pulumi.StringOutput `pulumi:"healthStatus"`
	// NSX-T ALB Importable Cloud ID
	ImportableCloudId pulumi.StringOutput `pulumi:"importableCloudId"`
	// NSX-T ALB Cloud name
	Name pulumi.StringOutput `pulumi:"name"`
	// Network pool ID for NSX-T ALB Importable Cloud
	NetworkPoolId pulumi.StringOutput `pulumi:"networkPoolId"`
	// Network pool name of NSX-T ALB Cloud
	NetworkPoolName pulumi.StringOutput `pulumi:"networkPoolName"`
}

// NewNsxtAlbCloud registers a new resource with the given unique name, arguments, and options.
func NewNsxtAlbCloud(ctx *pulumi.Context,
	name string, args *NsxtAlbCloudArgs, opts ...pulumi.ResourceOption) (*NsxtAlbCloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControllerId == nil {
		return nil, errors.New("invalid value for required argument 'ControllerId'")
	}
	if args.ImportableCloudId == nil {
		return nil, errors.New("invalid value for required argument 'ImportableCloudId'")
	}
	if args.NetworkPoolId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkPoolId'")
	}
	var resource NsxtAlbCloud
	err := ctx.RegisterResource("vcd:index/nsxtAlbCloud:NsxtAlbCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtAlbCloud gets an existing NsxtAlbCloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtAlbCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtAlbCloudState, opts ...pulumi.ResourceOption) (*NsxtAlbCloud, error) {
	var resource NsxtAlbCloud
	err := ctx.ReadResource("vcd:index/nsxtAlbCloud:NsxtAlbCloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtAlbCloud resources.
type nsxtAlbCloudState struct {
	// NSX-T ALB Controller ID
	ControllerId *string `pulumi:"controllerId"`
	// NSX-T ALB Cloud description
	Description *string `pulumi:"description"`
	// NSX-T ALB Cloud detailed health message
	HealthMessage *string `pulumi:"healthMessage"`
	// NSX-T ALB Cloud health status
	HealthStatus *string `pulumi:"healthStatus"`
	// NSX-T ALB Importable Cloud ID
	ImportableCloudId *string `pulumi:"importableCloudId"`
	// NSX-T ALB Cloud name
	Name *string `pulumi:"name"`
	// Network pool ID for NSX-T ALB Importable Cloud
	NetworkPoolId *string `pulumi:"networkPoolId"`
	// Network pool name of NSX-T ALB Cloud
	NetworkPoolName *string `pulumi:"networkPoolName"`
}

type NsxtAlbCloudState struct {
	// NSX-T ALB Controller ID
	ControllerId pulumi.StringPtrInput
	// NSX-T ALB Cloud description
	Description pulumi.StringPtrInput
	// NSX-T ALB Cloud detailed health message
	HealthMessage pulumi.StringPtrInput
	// NSX-T ALB Cloud health status
	HealthStatus pulumi.StringPtrInput
	// NSX-T ALB Importable Cloud ID
	ImportableCloudId pulumi.StringPtrInput
	// NSX-T ALB Cloud name
	Name pulumi.StringPtrInput
	// Network pool ID for NSX-T ALB Importable Cloud
	NetworkPoolId pulumi.StringPtrInput
	// Network pool name of NSX-T ALB Cloud
	NetworkPoolName pulumi.StringPtrInput
}

func (NsxtAlbCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbCloudState)(nil)).Elem()
}

type nsxtAlbCloudArgs struct {
	// NSX-T ALB Controller ID
	ControllerId string `pulumi:"controllerId"`
	// NSX-T ALB Cloud description
	Description *string `pulumi:"description"`
	// NSX-T ALB Importable Cloud ID
	ImportableCloudId string `pulumi:"importableCloudId"`
	// NSX-T ALB Cloud name
	Name *string `pulumi:"name"`
	// Network pool ID for NSX-T ALB Importable Cloud
	NetworkPoolId string `pulumi:"networkPoolId"`
}

// The set of arguments for constructing a NsxtAlbCloud resource.
type NsxtAlbCloudArgs struct {
	// NSX-T ALB Controller ID
	ControllerId pulumi.StringInput
	// NSX-T ALB Cloud description
	Description pulumi.StringPtrInput
	// NSX-T ALB Importable Cloud ID
	ImportableCloudId pulumi.StringInput
	// NSX-T ALB Cloud name
	Name pulumi.StringPtrInput
	// Network pool ID for NSX-T ALB Importable Cloud
	NetworkPoolId pulumi.StringInput
}

func (NsxtAlbCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbCloudArgs)(nil)).Elem()
}

type NsxtAlbCloudInput interface {
	pulumi.Input

	ToNsxtAlbCloudOutput() NsxtAlbCloudOutput
	ToNsxtAlbCloudOutputWithContext(ctx context.Context) NsxtAlbCloudOutput
}

func (*NsxtAlbCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbCloud)(nil)).Elem()
}

func (i *NsxtAlbCloud) ToNsxtAlbCloudOutput() NsxtAlbCloudOutput {
	return i.ToNsxtAlbCloudOutputWithContext(context.Background())
}

func (i *NsxtAlbCloud) ToNsxtAlbCloudOutputWithContext(ctx context.Context) NsxtAlbCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbCloudOutput)
}

// NsxtAlbCloudArrayInput is an input type that accepts NsxtAlbCloudArray and NsxtAlbCloudArrayOutput values.
// You can construct a concrete instance of `NsxtAlbCloudArrayInput` via:
//
//	NsxtAlbCloudArray{ NsxtAlbCloudArgs{...} }
type NsxtAlbCloudArrayInput interface {
	pulumi.Input

	ToNsxtAlbCloudArrayOutput() NsxtAlbCloudArrayOutput
	ToNsxtAlbCloudArrayOutputWithContext(context.Context) NsxtAlbCloudArrayOutput
}

type NsxtAlbCloudArray []NsxtAlbCloudInput

func (NsxtAlbCloudArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbCloud)(nil)).Elem()
}

func (i NsxtAlbCloudArray) ToNsxtAlbCloudArrayOutput() NsxtAlbCloudArrayOutput {
	return i.ToNsxtAlbCloudArrayOutputWithContext(context.Background())
}

func (i NsxtAlbCloudArray) ToNsxtAlbCloudArrayOutputWithContext(ctx context.Context) NsxtAlbCloudArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbCloudArrayOutput)
}

// NsxtAlbCloudMapInput is an input type that accepts NsxtAlbCloudMap and NsxtAlbCloudMapOutput values.
// You can construct a concrete instance of `NsxtAlbCloudMapInput` via:
//
//	NsxtAlbCloudMap{ "key": NsxtAlbCloudArgs{...} }
type NsxtAlbCloudMapInput interface {
	pulumi.Input

	ToNsxtAlbCloudMapOutput() NsxtAlbCloudMapOutput
	ToNsxtAlbCloudMapOutputWithContext(context.Context) NsxtAlbCloudMapOutput
}

type NsxtAlbCloudMap map[string]NsxtAlbCloudInput

func (NsxtAlbCloudMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbCloud)(nil)).Elem()
}

func (i NsxtAlbCloudMap) ToNsxtAlbCloudMapOutput() NsxtAlbCloudMapOutput {
	return i.ToNsxtAlbCloudMapOutputWithContext(context.Background())
}

func (i NsxtAlbCloudMap) ToNsxtAlbCloudMapOutputWithContext(ctx context.Context) NsxtAlbCloudMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbCloudMapOutput)
}

type NsxtAlbCloudOutput struct{ *pulumi.OutputState }

func (NsxtAlbCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbCloud)(nil)).Elem()
}

func (o NsxtAlbCloudOutput) ToNsxtAlbCloudOutput() NsxtAlbCloudOutput {
	return o
}

func (o NsxtAlbCloudOutput) ToNsxtAlbCloudOutputWithContext(ctx context.Context) NsxtAlbCloudOutput {
	return o
}

// NSX-T ALB Controller ID
func (o NsxtAlbCloudOutput) ControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.ControllerId }).(pulumi.StringOutput)
}

// NSX-T ALB Cloud description
func (o NsxtAlbCloudOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// NSX-T ALB Cloud detailed health message
func (o NsxtAlbCloudOutput) HealthMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.HealthMessage }).(pulumi.StringOutput)
}

// NSX-T ALB Cloud health status
func (o NsxtAlbCloudOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

// NSX-T ALB Importable Cloud ID
func (o NsxtAlbCloudOutput) ImportableCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.ImportableCloudId }).(pulumi.StringOutput)
}

// NSX-T ALB Cloud name
func (o NsxtAlbCloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network pool ID for NSX-T ALB Importable Cloud
func (o NsxtAlbCloudOutput) NetworkPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.NetworkPoolId }).(pulumi.StringOutput)
}

// Network pool name of NSX-T ALB Cloud
func (o NsxtAlbCloudOutput) NetworkPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbCloud) pulumi.StringOutput { return v.NetworkPoolName }).(pulumi.StringOutput)
}

type NsxtAlbCloudArrayOutput struct{ *pulumi.OutputState }

func (NsxtAlbCloudArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbCloud)(nil)).Elem()
}

func (o NsxtAlbCloudArrayOutput) ToNsxtAlbCloudArrayOutput() NsxtAlbCloudArrayOutput {
	return o
}

func (o NsxtAlbCloudArrayOutput) ToNsxtAlbCloudArrayOutputWithContext(ctx context.Context) NsxtAlbCloudArrayOutput {
	return o
}

func (o NsxtAlbCloudArrayOutput) Index(i pulumi.IntInput) NsxtAlbCloudOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtAlbCloud {
		return vs[0].([]*NsxtAlbCloud)[vs[1].(int)]
	}).(NsxtAlbCloudOutput)
}

type NsxtAlbCloudMapOutput struct{ *pulumi.OutputState }

func (NsxtAlbCloudMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbCloud)(nil)).Elem()
}

func (o NsxtAlbCloudMapOutput) ToNsxtAlbCloudMapOutput() NsxtAlbCloudMapOutput {
	return o
}

func (o NsxtAlbCloudMapOutput) ToNsxtAlbCloudMapOutputWithContext(ctx context.Context) NsxtAlbCloudMapOutput {
	return o
}

func (o NsxtAlbCloudMapOutput) MapIndex(k pulumi.StringInput) NsxtAlbCloudOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtAlbCloud {
		return vs[0].(map[string]*NsxtAlbCloud)[vs[1].(string)]
	}).(NsxtAlbCloudOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbCloudInput)(nil)).Elem(), &NsxtAlbCloud{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbCloudArrayInput)(nil)).Elem(), NsxtAlbCloudArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbCloudMapInput)(nil)).Elem(), NsxtAlbCloudMap{})
	pulumi.RegisterOutputType(NsxtAlbCloudOutput{})
	pulumi.RegisterOutputType(NsxtAlbCloudArrayOutput{})
	pulumi.RegisterOutputType(NsxtAlbCloudMapOutput{})
}
