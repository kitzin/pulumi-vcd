// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxtDistributedFirewall struct {
	pulumi.CustomResourceState

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Ordered list of firewall rules
	Rules NsxtDistributedFirewallRuleArrayOutput `pulumi:"rules"`
	// ID of VDC Group for Distributed Firewall
	VdcGroupId pulumi.StringOutput `pulumi:"vdcGroupId"`
}

// NewNsxtDistributedFirewall registers a new resource with the given unique name, arguments, and options.
func NewNsxtDistributedFirewall(ctx *pulumi.Context,
	name string, args *NsxtDistributedFirewallArgs, opts ...pulumi.ResourceOption) (*NsxtDistributedFirewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.VdcGroupId == nil {
		return nil, errors.New("invalid value for required argument 'VdcGroupId'")
	}
	var resource NsxtDistributedFirewall
	err := ctx.RegisterResource("vcd:index/nsxtDistributedFirewall:NsxtDistributedFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtDistributedFirewall gets an existing NsxtDistributedFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtDistributedFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtDistributedFirewallState, opts ...pulumi.ResourceOption) (*NsxtDistributedFirewall, error) {
	var resource NsxtDistributedFirewall
	err := ctx.ReadResource("vcd:index/nsxtDistributedFirewall:NsxtDistributedFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtDistributedFirewall resources.
type nsxtDistributedFirewallState struct {
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Ordered list of firewall rules
	Rules []NsxtDistributedFirewallRule `pulumi:"rules"`
	// ID of VDC Group for Distributed Firewall
	VdcGroupId *string `pulumi:"vdcGroupId"`
}

type NsxtDistributedFirewallState struct {
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Ordered list of firewall rules
	Rules NsxtDistributedFirewallRuleArrayInput
	// ID of VDC Group for Distributed Firewall
	VdcGroupId pulumi.StringPtrInput
}

func (NsxtDistributedFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtDistributedFirewallState)(nil)).Elem()
}

type nsxtDistributedFirewallArgs struct {
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Ordered list of firewall rules
	Rules []NsxtDistributedFirewallRule `pulumi:"rules"`
	// ID of VDC Group for Distributed Firewall
	VdcGroupId string `pulumi:"vdcGroupId"`
}

// The set of arguments for constructing a NsxtDistributedFirewall resource.
type NsxtDistributedFirewallArgs struct {
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Ordered list of firewall rules
	Rules NsxtDistributedFirewallRuleArrayInput
	// ID of VDC Group for Distributed Firewall
	VdcGroupId pulumi.StringInput
}

func (NsxtDistributedFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtDistributedFirewallArgs)(nil)).Elem()
}

type NsxtDistributedFirewallInput interface {
	pulumi.Input

	ToNsxtDistributedFirewallOutput() NsxtDistributedFirewallOutput
	ToNsxtDistributedFirewallOutputWithContext(ctx context.Context) NsxtDistributedFirewallOutput
}

func (*NsxtDistributedFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtDistributedFirewall)(nil)).Elem()
}

func (i *NsxtDistributedFirewall) ToNsxtDistributedFirewallOutput() NsxtDistributedFirewallOutput {
	return i.ToNsxtDistributedFirewallOutputWithContext(context.Background())
}

func (i *NsxtDistributedFirewall) ToNsxtDistributedFirewallOutputWithContext(ctx context.Context) NsxtDistributedFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtDistributedFirewallOutput)
}

// NsxtDistributedFirewallArrayInput is an input type that accepts NsxtDistributedFirewallArray and NsxtDistributedFirewallArrayOutput values.
// You can construct a concrete instance of `NsxtDistributedFirewallArrayInput` via:
//
//	NsxtDistributedFirewallArray{ NsxtDistributedFirewallArgs{...} }
type NsxtDistributedFirewallArrayInput interface {
	pulumi.Input

	ToNsxtDistributedFirewallArrayOutput() NsxtDistributedFirewallArrayOutput
	ToNsxtDistributedFirewallArrayOutputWithContext(context.Context) NsxtDistributedFirewallArrayOutput
}

type NsxtDistributedFirewallArray []NsxtDistributedFirewallInput

func (NsxtDistributedFirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtDistributedFirewall)(nil)).Elem()
}

func (i NsxtDistributedFirewallArray) ToNsxtDistributedFirewallArrayOutput() NsxtDistributedFirewallArrayOutput {
	return i.ToNsxtDistributedFirewallArrayOutputWithContext(context.Background())
}

func (i NsxtDistributedFirewallArray) ToNsxtDistributedFirewallArrayOutputWithContext(ctx context.Context) NsxtDistributedFirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtDistributedFirewallArrayOutput)
}

// NsxtDistributedFirewallMapInput is an input type that accepts NsxtDistributedFirewallMap and NsxtDistributedFirewallMapOutput values.
// You can construct a concrete instance of `NsxtDistributedFirewallMapInput` via:
//
//	NsxtDistributedFirewallMap{ "key": NsxtDistributedFirewallArgs{...} }
type NsxtDistributedFirewallMapInput interface {
	pulumi.Input

	ToNsxtDistributedFirewallMapOutput() NsxtDistributedFirewallMapOutput
	ToNsxtDistributedFirewallMapOutputWithContext(context.Context) NsxtDistributedFirewallMapOutput
}

type NsxtDistributedFirewallMap map[string]NsxtDistributedFirewallInput

func (NsxtDistributedFirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtDistributedFirewall)(nil)).Elem()
}

func (i NsxtDistributedFirewallMap) ToNsxtDistributedFirewallMapOutput() NsxtDistributedFirewallMapOutput {
	return i.ToNsxtDistributedFirewallMapOutputWithContext(context.Background())
}

func (i NsxtDistributedFirewallMap) ToNsxtDistributedFirewallMapOutputWithContext(ctx context.Context) NsxtDistributedFirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtDistributedFirewallMapOutput)
}

type NsxtDistributedFirewallOutput struct{ *pulumi.OutputState }

func (NsxtDistributedFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtDistributedFirewall)(nil)).Elem()
}

func (o NsxtDistributedFirewallOutput) ToNsxtDistributedFirewallOutput() NsxtDistributedFirewallOutput {
	return o
}

func (o NsxtDistributedFirewallOutput) ToNsxtDistributedFirewallOutputWithContext(ctx context.Context) NsxtDistributedFirewallOutput {
	return o
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NsxtDistributedFirewallOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtDistributedFirewall) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Ordered list of firewall rules
func (o NsxtDistributedFirewallOutput) Rules() NsxtDistributedFirewallRuleArrayOutput {
	return o.ApplyT(func(v *NsxtDistributedFirewall) NsxtDistributedFirewallRuleArrayOutput { return v.Rules }).(NsxtDistributedFirewallRuleArrayOutput)
}

// ID of VDC Group for Distributed Firewall
func (o NsxtDistributedFirewallOutput) VdcGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtDistributedFirewall) pulumi.StringOutput { return v.VdcGroupId }).(pulumi.StringOutput)
}

type NsxtDistributedFirewallArrayOutput struct{ *pulumi.OutputState }

func (NsxtDistributedFirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtDistributedFirewall)(nil)).Elem()
}

func (o NsxtDistributedFirewallArrayOutput) ToNsxtDistributedFirewallArrayOutput() NsxtDistributedFirewallArrayOutput {
	return o
}

func (o NsxtDistributedFirewallArrayOutput) ToNsxtDistributedFirewallArrayOutputWithContext(ctx context.Context) NsxtDistributedFirewallArrayOutput {
	return o
}

func (o NsxtDistributedFirewallArrayOutput) Index(i pulumi.IntInput) NsxtDistributedFirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtDistributedFirewall {
		return vs[0].([]*NsxtDistributedFirewall)[vs[1].(int)]
	}).(NsxtDistributedFirewallOutput)
}

type NsxtDistributedFirewallMapOutput struct{ *pulumi.OutputState }

func (NsxtDistributedFirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtDistributedFirewall)(nil)).Elem()
}

func (o NsxtDistributedFirewallMapOutput) ToNsxtDistributedFirewallMapOutput() NsxtDistributedFirewallMapOutput {
	return o
}

func (o NsxtDistributedFirewallMapOutput) ToNsxtDistributedFirewallMapOutputWithContext(ctx context.Context) NsxtDistributedFirewallMapOutput {
	return o
}

func (o NsxtDistributedFirewallMapOutput) MapIndex(k pulumi.StringInput) NsxtDistributedFirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtDistributedFirewall {
		return vs[0].(map[string]*NsxtDistributedFirewall)[vs[1].(string)]
	}).(NsxtDistributedFirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtDistributedFirewallInput)(nil)).Elem(), &NsxtDistributedFirewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtDistributedFirewallArrayInput)(nil)).Elem(), NsxtDistributedFirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtDistributedFirewallMapInput)(nil)).Elem(), NsxtDistributedFirewallMap{})
	pulumi.RegisterOutputType(NsxtDistributedFirewallOutput{})
	pulumi.RegisterOutputType(NsxtDistributedFirewallArrayOutput{})
	pulumi.RegisterOutputType(NsxtDistributedFirewallMapOutput{})
}