// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxvIpSet struct {
	pulumi.CustomResourceState

	// IP set description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A set of IP address, CIDR, IP range objects
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Allows visibility in underlying scopes (Default is true)
	IsInheritanceAllowed pulumi.BoolPtrOutput `pulumi:"isInheritanceAllowed"`
	// IP set name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewNsxvIpSet registers a new resource with the given unique name, arguments, and options.
func NewNsxvIpSet(ctx *pulumi.Context,
	name string, args *NsxvIpSetArgs, opts ...pulumi.ResourceOption) (*NsxvIpSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddresses == nil {
		return nil, errors.New("invalid value for required argument 'IpAddresses'")
	}
	var resource NsxvIpSet
	err := ctx.RegisterResource("vcd:index/nsxvIpSet:NsxvIpSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxvIpSet gets an existing NsxvIpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxvIpSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxvIpSetState, opts ...pulumi.ResourceOption) (*NsxvIpSet, error) {
	var resource NsxvIpSet
	err := ctx.ReadResource("vcd:index/nsxvIpSet:NsxvIpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxvIpSet resources.
type nsxvIpSetState struct {
	// IP set description
	Description *string `pulumi:"description"`
	// A set of IP address, CIDR, IP range objects
	IpAddresses []string `pulumi:"ipAddresses"`
	// Allows visibility in underlying scopes (Default is true)
	IsInheritanceAllowed *bool `pulumi:"isInheritanceAllowed"`
	// IP set name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type NsxvIpSetState struct {
	// IP set description
	Description pulumi.StringPtrInput
	// A set of IP address, CIDR, IP range objects
	IpAddresses pulumi.StringArrayInput
	// Allows visibility in underlying scopes (Default is true)
	IsInheritanceAllowed pulumi.BoolPtrInput
	// IP set name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NsxvIpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxvIpSetState)(nil)).Elem()
}

type nsxvIpSetArgs struct {
	// IP set description
	Description *string `pulumi:"description"`
	// A set of IP address, CIDR, IP range objects
	IpAddresses []string `pulumi:"ipAddresses"`
	// Allows visibility in underlying scopes (Default is true)
	IsInheritanceAllowed *bool `pulumi:"isInheritanceAllowed"`
	// IP set name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NsxvIpSet resource.
type NsxvIpSetArgs struct {
	// IP set description
	Description pulumi.StringPtrInput
	// A set of IP address, CIDR, IP range objects
	IpAddresses pulumi.StringArrayInput
	// Allows visibility in underlying scopes (Default is true)
	IsInheritanceAllowed pulumi.BoolPtrInput
	// IP set name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NsxvIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxvIpSetArgs)(nil)).Elem()
}

type NsxvIpSetInput interface {
	pulumi.Input

	ToNsxvIpSetOutput() NsxvIpSetOutput
	ToNsxvIpSetOutputWithContext(ctx context.Context) NsxvIpSetOutput
}

func (*NsxvIpSet) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxvIpSet)(nil)).Elem()
}

func (i *NsxvIpSet) ToNsxvIpSetOutput() NsxvIpSetOutput {
	return i.ToNsxvIpSetOutputWithContext(context.Background())
}

func (i *NsxvIpSet) ToNsxvIpSetOutputWithContext(ctx context.Context) NsxvIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvIpSetOutput)
}

// NsxvIpSetArrayInput is an input type that accepts NsxvIpSetArray and NsxvIpSetArrayOutput values.
// You can construct a concrete instance of `NsxvIpSetArrayInput` via:
//
//	NsxvIpSetArray{ NsxvIpSetArgs{...} }
type NsxvIpSetArrayInput interface {
	pulumi.Input

	ToNsxvIpSetArrayOutput() NsxvIpSetArrayOutput
	ToNsxvIpSetArrayOutputWithContext(context.Context) NsxvIpSetArrayOutput
}

type NsxvIpSetArray []NsxvIpSetInput

func (NsxvIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxvIpSet)(nil)).Elem()
}

func (i NsxvIpSetArray) ToNsxvIpSetArrayOutput() NsxvIpSetArrayOutput {
	return i.ToNsxvIpSetArrayOutputWithContext(context.Background())
}

func (i NsxvIpSetArray) ToNsxvIpSetArrayOutputWithContext(ctx context.Context) NsxvIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvIpSetArrayOutput)
}

// NsxvIpSetMapInput is an input type that accepts NsxvIpSetMap and NsxvIpSetMapOutput values.
// You can construct a concrete instance of `NsxvIpSetMapInput` via:
//
//	NsxvIpSetMap{ "key": NsxvIpSetArgs{...} }
type NsxvIpSetMapInput interface {
	pulumi.Input

	ToNsxvIpSetMapOutput() NsxvIpSetMapOutput
	ToNsxvIpSetMapOutputWithContext(context.Context) NsxvIpSetMapOutput
}

type NsxvIpSetMap map[string]NsxvIpSetInput

func (NsxvIpSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxvIpSet)(nil)).Elem()
}

func (i NsxvIpSetMap) ToNsxvIpSetMapOutput() NsxvIpSetMapOutput {
	return i.ToNsxvIpSetMapOutputWithContext(context.Background())
}

func (i NsxvIpSetMap) ToNsxvIpSetMapOutputWithContext(ctx context.Context) NsxvIpSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvIpSetMapOutput)
}

type NsxvIpSetOutput struct{ *pulumi.OutputState }

func (NsxvIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxvIpSet)(nil)).Elem()
}

func (o NsxvIpSetOutput) ToNsxvIpSetOutput() NsxvIpSetOutput {
	return o
}

func (o NsxvIpSetOutput) ToNsxvIpSetOutputWithContext(ctx context.Context) NsxvIpSetOutput {
	return o
}

// IP set description
func (o NsxvIpSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvIpSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A set of IP address, CIDR, IP range objects
func (o NsxvIpSetOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NsxvIpSet) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Allows visibility in underlying scopes (Default is true)
func (o NsxvIpSetOutput) IsInheritanceAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxvIpSet) pulumi.BoolPtrOutput { return v.IsInheritanceAllowed }).(pulumi.BoolPtrOutput)
}

// IP set name
func (o NsxvIpSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvIpSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NsxvIpSetOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvIpSet) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o NsxvIpSetOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvIpSet) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type NsxvIpSetArrayOutput struct{ *pulumi.OutputState }

func (NsxvIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxvIpSet)(nil)).Elem()
}

func (o NsxvIpSetArrayOutput) ToNsxvIpSetArrayOutput() NsxvIpSetArrayOutput {
	return o
}

func (o NsxvIpSetArrayOutput) ToNsxvIpSetArrayOutputWithContext(ctx context.Context) NsxvIpSetArrayOutput {
	return o
}

func (o NsxvIpSetArrayOutput) Index(i pulumi.IntInput) NsxvIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxvIpSet {
		return vs[0].([]*NsxvIpSet)[vs[1].(int)]
	}).(NsxvIpSetOutput)
}

type NsxvIpSetMapOutput struct{ *pulumi.OutputState }

func (NsxvIpSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxvIpSet)(nil)).Elem()
}

func (o NsxvIpSetMapOutput) ToNsxvIpSetMapOutput() NsxvIpSetMapOutput {
	return o
}

func (o NsxvIpSetMapOutput) ToNsxvIpSetMapOutputWithContext(ctx context.Context) NsxvIpSetMapOutput {
	return o
}

func (o NsxvIpSetMapOutput) MapIndex(k pulumi.StringInput) NsxvIpSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxvIpSet {
		return vs[0].(map[string]*NsxvIpSet)[vs[1].(string)]
	}).(NsxvIpSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvIpSetInput)(nil)).Elem(), &NsxvIpSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvIpSetArrayInput)(nil)).Elem(), NsxvIpSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvIpSetMapInput)(nil)).Elem(), NsxvIpSetMap{})
	pulumi.RegisterOutputType(NsxvIpSetOutput{})
	pulumi.RegisterOutputType(NsxvIpSetArrayOutput{})
	pulumi.RegisterOutputType(NsxvIpSetMapOutput{})
}