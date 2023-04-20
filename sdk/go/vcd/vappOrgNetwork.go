// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VappOrgNetwork struct {
	pulumi.CustomResourceState

	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are
	// accessed in this vApp
	IsFenced pulumi.BoolPtrOutput `pulumi:"isFenced"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Organization network name to which vApp network is connected to
	OrgNetworkName pulumi.StringOutput `pulumi:"orgNetworkName"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled pulumi.BoolPtrOutput `pulumi:"retainIpMacEnabled"`
	// vApp network name
	VappName pulumi.StringOutput `pulumi:"vappName"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewVappOrgNetwork registers a new resource with the given unique name, arguments, and options.
func NewVappOrgNetwork(ctx *pulumi.Context,
	name string, args *VappOrgNetworkArgs, opts ...pulumi.ResourceOption) (*VappOrgNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'OrgNetworkName'")
	}
	if args.VappName == nil {
		return nil, errors.New("invalid value for required argument 'VappName'")
	}
	var resource VappOrgNetwork
	err := ctx.RegisterResource("vcd:index/vappOrgNetwork:VappOrgNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVappOrgNetwork gets an existing VappOrgNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappOrgNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VappOrgNetworkState, opts ...pulumi.ResourceOption) (*VappOrgNetwork, error) {
	var resource VappOrgNetwork
	err := ctx.ReadResource("vcd:index/vappOrgNetwork:VappOrgNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VappOrgNetwork resources.
type vappOrgNetworkState struct {
	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are
	// accessed in this vApp
	IsFenced *bool `pulumi:"isFenced"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Organization network name to which vApp network is connected to
	OrgNetworkName *string `pulumi:"orgNetworkName"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled *bool `pulumi:"retainIpMacEnabled"`
	// vApp network name
	VappName *string `pulumi:"vappName"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type VappOrgNetworkState struct {
	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are
	// accessed in this vApp
	IsFenced pulumi.BoolPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Organization network name to which vApp network is connected to
	OrgNetworkName pulumi.StringPtrInput
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled pulumi.BoolPtrInput
	// vApp network name
	VappName pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (VappOrgNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vappOrgNetworkState)(nil)).Elem()
}

type vappOrgNetworkArgs struct {
	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are
	// accessed in this vApp
	IsFenced *bool `pulumi:"isFenced"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Organization network name to which vApp network is connected to
	OrgNetworkName string `pulumi:"orgNetworkName"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled *bool `pulumi:"retainIpMacEnabled"`
	// vApp network name
	VappName string `pulumi:"vappName"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a VappOrgNetwork resource.
type VappOrgNetworkArgs struct {
	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are
	// accessed in this vApp
	IsFenced pulumi.BoolPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Organization network name to which vApp network is connected to
	OrgNetworkName pulumi.StringInput
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled pulumi.BoolPtrInput
	// vApp network name
	VappName pulumi.StringInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (VappOrgNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vappOrgNetworkArgs)(nil)).Elem()
}

type VappOrgNetworkInput interface {
	pulumi.Input

	ToVappOrgNetworkOutput() VappOrgNetworkOutput
	ToVappOrgNetworkOutputWithContext(ctx context.Context) VappOrgNetworkOutput
}

func (*VappOrgNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VappOrgNetwork)(nil)).Elem()
}

func (i *VappOrgNetwork) ToVappOrgNetworkOutput() VappOrgNetworkOutput {
	return i.ToVappOrgNetworkOutputWithContext(context.Background())
}

func (i *VappOrgNetwork) ToVappOrgNetworkOutputWithContext(ctx context.Context) VappOrgNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappOrgNetworkOutput)
}

// VappOrgNetworkArrayInput is an input type that accepts VappOrgNetworkArray and VappOrgNetworkArrayOutput values.
// You can construct a concrete instance of `VappOrgNetworkArrayInput` via:
//
//	VappOrgNetworkArray{ VappOrgNetworkArgs{...} }
type VappOrgNetworkArrayInput interface {
	pulumi.Input

	ToVappOrgNetworkArrayOutput() VappOrgNetworkArrayOutput
	ToVappOrgNetworkArrayOutputWithContext(context.Context) VappOrgNetworkArrayOutput
}

type VappOrgNetworkArray []VappOrgNetworkInput

func (VappOrgNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappOrgNetwork)(nil)).Elem()
}

func (i VappOrgNetworkArray) ToVappOrgNetworkArrayOutput() VappOrgNetworkArrayOutput {
	return i.ToVappOrgNetworkArrayOutputWithContext(context.Background())
}

func (i VappOrgNetworkArray) ToVappOrgNetworkArrayOutputWithContext(ctx context.Context) VappOrgNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappOrgNetworkArrayOutput)
}

// VappOrgNetworkMapInput is an input type that accepts VappOrgNetworkMap and VappOrgNetworkMapOutput values.
// You can construct a concrete instance of `VappOrgNetworkMapInput` via:
//
//	VappOrgNetworkMap{ "key": VappOrgNetworkArgs{...} }
type VappOrgNetworkMapInput interface {
	pulumi.Input

	ToVappOrgNetworkMapOutput() VappOrgNetworkMapOutput
	ToVappOrgNetworkMapOutputWithContext(context.Context) VappOrgNetworkMapOutput
}

type VappOrgNetworkMap map[string]VappOrgNetworkInput

func (VappOrgNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappOrgNetwork)(nil)).Elem()
}

func (i VappOrgNetworkMap) ToVappOrgNetworkMapOutput() VappOrgNetworkMapOutput {
	return i.ToVappOrgNetworkMapOutputWithContext(context.Background())
}

func (i VappOrgNetworkMap) ToVappOrgNetworkMapOutputWithContext(ctx context.Context) VappOrgNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappOrgNetworkMapOutput)
}

type VappOrgNetworkOutput struct{ *pulumi.OutputState }

func (VappOrgNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VappOrgNetwork)(nil)).Elem()
}

func (o VappOrgNetworkOutput) ToVappOrgNetworkOutput() VappOrgNetworkOutput {
	return o
}

func (o VappOrgNetworkOutput) ToVappOrgNetworkOutputWithContext(ctx context.Context) VappOrgNetworkOutput {
	return o
}

// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are
// accessed in this vApp
func (o VappOrgNetworkOutput) IsFenced() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappOrgNetwork) pulumi.BoolPtrOutput { return v.IsFenced }).(pulumi.BoolPtrOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o VappOrgNetworkOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappOrgNetwork) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Organization network name to which vApp network is connected to
func (o VappOrgNetworkOutput) OrgNetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v *VappOrgNetwork) pulumi.StringOutput { return v.OrgNetworkName }).(pulumi.StringOutput)
}

// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
func (o VappOrgNetworkOutput) RetainIpMacEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappOrgNetwork) pulumi.BoolPtrOutput { return v.RetainIpMacEnabled }).(pulumi.BoolPtrOutput)
}

// vApp network name
func (o VappOrgNetworkOutput) VappName() pulumi.StringOutput {
	return o.ApplyT(func(v *VappOrgNetwork) pulumi.StringOutput { return v.VappName }).(pulumi.StringOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o VappOrgNetworkOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappOrgNetwork) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type VappOrgNetworkArrayOutput struct{ *pulumi.OutputState }

func (VappOrgNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappOrgNetwork)(nil)).Elem()
}

func (o VappOrgNetworkArrayOutput) ToVappOrgNetworkArrayOutput() VappOrgNetworkArrayOutput {
	return o
}

func (o VappOrgNetworkArrayOutput) ToVappOrgNetworkArrayOutputWithContext(ctx context.Context) VappOrgNetworkArrayOutput {
	return o
}

func (o VappOrgNetworkArrayOutput) Index(i pulumi.IntInput) VappOrgNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VappOrgNetwork {
		return vs[0].([]*VappOrgNetwork)[vs[1].(int)]
	}).(VappOrgNetworkOutput)
}

type VappOrgNetworkMapOutput struct{ *pulumi.OutputState }

func (VappOrgNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappOrgNetwork)(nil)).Elem()
}

func (o VappOrgNetworkMapOutput) ToVappOrgNetworkMapOutput() VappOrgNetworkMapOutput {
	return o
}

func (o VappOrgNetworkMapOutput) ToVappOrgNetworkMapOutputWithContext(ctx context.Context) VappOrgNetworkMapOutput {
	return o
}

func (o VappOrgNetworkMapOutput) MapIndex(k pulumi.StringInput) VappOrgNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VappOrgNetwork {
		return vs[0].(map[string]*VappOrgNetwork)[vs[1].(string)]
	}).(VappOrgNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VappOrgNetworkInput)(nil)).Elem(), &VappOrgNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappOrgNetworkArrayInput)(nil)).Elem(), VappOrgNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappOrgNetworkMapInput)(nil)).Elem(), VappOrgNetworkMap{})
	pulumi.RegisterOutputType(VappOrgNetworkOutput{})
	pulumi.RegisterOutputType(VappOrgNetworkArrayOutput{})
	pulumi.RegisterOutputType(VappOrgNetworkMapOutput{})
}