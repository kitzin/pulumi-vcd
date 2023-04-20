// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VmAffinityRule struct {
	pulumi.CustomResourceState

	// True if this affinity rule is enabled
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// VM affinity rule name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// One of 'Affinity', 'Anti-Affinity'
	Polarity pulumi.StringOutput `pulumi:"polarity"`
	// True if this affinity rule is required. When a rule is mandatory, a host failover will not power on the VM if doing so
	// would violate the rule
	Required pulumi.BoolPtrOutput `pulumi:"required"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
	// Set of VM IDs assigned to this rule
	VmIds pulumi.StringArrayOutput `pulumi:"vmIds"`
}

// NewVmAffinityRule registers a new resource with the given unique name, arguments, and options.
func NewVmAffinityRule(ctx *pulumi.Context,
	name string, args *VmAffinityRuleArgs, opts ...pulumi.ResourceOption) (*VmAffinityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Polarity == nil {
		return nil, errors.New("invalid value for required argument 'Polarity'")
	}
	if args.VmIds == nil {
		return nil, errors.New("invalid value for required argument 'VmIds'")
	}
	var resource VmAffinityRule
	err := ctx.RegisterResource("vcd:index/vmAffinityRule:VmAffinityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmAffinityRule gets an existing VmAffinityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmAffinityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmAffinityRuleState, opts ...pulumi.ResourceOption) (*VmAffinityRule, error) {
	var resource VmAffinityRule
	err := ctx.ReadResource("vcd:index/vmAffinityRule:VmAffinityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmAffinityRule resources.
type vmAffinityRuleState struct {
	// True if this affinity rule is enabled
	Enabled *bool `pulumi:"enabled"`
	// VM affinity rule name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// One of 'Affinity', 'Anti-Affinity'
	Polarity *string `pulumi:"polarity"`
	// True if this affinity rule is required. When a rule is mandatory, a host failover will not power on the VM if doing so
	// would violate the rule
	Required *bool `pulumi:"required"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
	// Set of VM IDs assigned to this rule
	VmIds []string `pulumi:"vmIds"`
}

type VmAffinityRuleState struct {
	// True if this affinity rule is enabled
	Enabled pulumi.BoolPtrInput
	// VM affinity rule name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// One of 'Affinity', 'Anti-Affinity'
	Polarity pulumi.StringPtrInput
	// True if this affinity rule is required. When a rule is mandatory, a host failover will not power on the VM if doing so
	// would violate the rule
	Required pulumi.BoolPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
	// Set of VM IDs assigned to this rule
	VmIds pulumi.StringArrayInput
}

func (VmAffinityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmAffinityRuleState)(nil)).Elem()
}

type vmAffinityRuleArgs struct {
	// True if this affinity rule is enabled
	Enabled *bool `pulumi:"enabled"`
	// VM affinity rule name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// One of 'Affinity', 'Anti-Affinity'
	Polarity string `pulumi:"polarity"`
	// True if this affinity rule is required. When a rule is mandatory, a host failover will not power on the VM if doing so
	// would violate the rule
	Required *bool `pulumi:"required"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
	// Set of VM IDs assigned to this rule
	VmIds []string `pulumi:"vmIds"`
}

// The set of arguments for constructing a VmAffinityRule resource.
type VmAffinityRuleArgs struct {
	// True if this affinity rule is enabled
	Enabled pulumi.BoolPtrInput
	// VM affinity rule name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// One of 'Affinity', 'Anti-Affinity'
	Polarity pulumi.StringInput
	// True if this affinity rule is required. When a rule is mandatory, a host failover will not power on the VM if doing so
	// would violate the rule
	Required pulumi.BoolPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
	// Set of VM IDs assigned to this rule
	VmIds pulumi.StringArrayInput
}

func (VmAffinityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmAffinityRuleArgs)(nil)).Elem()
}

type VmAffinityRuleInput interface {
	pulumi.Input

	ToVmAffinityRuleOutput() VmAffinityRuleOutput
	ToVmAffinityRuleOutputWithContext(ctx context.Context) VmAffinityRuleOutput
}

func (*VmAffinityRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VmAffinityRule)(nil)).Elem()
}

func (i *VmAffinityRule) ToVmAffinityRuleOutput() VmAffinityRuleOutput {
	return i.ToVmAffinityRuleOutputWithContext(context.Background())
}

func (i *VmAffinityRule) ToVmAffinityRuleOutputWithContext(ctx context.Context) VmAffinityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmAffinityRuleOutput)
}

// VmAffinityRuleArrayInput is an input type that accepts VmAffinityRuleArray and VmAffinityRuleArrayOutput values.
// You can construct a concrete instance of `VmAffinityRuleArrayInput` via:
//
//	VmAffinityRuleArray{ VmAffinityRuleArgs{...} }
type VmAffinityRuleArrayInput interface {
	pulumi.Input

	ToVmAffinityRuleArrayOutput() VmAffinityRuleArrayOutput
	ToVmAffinityRuleArrayOutputWithContext(context.Context) VmAffinityRuleArrayOutput
}

type VmAffinityRuleArray []VmAffinityRuleInput

func (VmAffinityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmAffinityRule)(nil)).Elem()
}

func (i VmAffinityRuleArray) ToVmAffinityRuleArrayOutput() VmAffinityRuleArrayOutput {
	return i.ToVmAffinityRuleArrayOutputWithContext(context.Background())
}

func (i VmAffinityRuleArray) ToVmAffinityRuleArrayOutputWithContext(ctx context.Context) VmAffinityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmAffinityRuleArrayOutput)
}

// VmAffinityRuleMapInput is an input type that accepts VmAffinityRuleMap and VmAffinityRuleMapOutput values.
// You can construct a concrete instance of `VmAffinityRuleMapInput` via:
//
//	VmAffinityRuleMap{ "key": VmAffinityRuleArgs{...} }
type VmAffinityRuleMapInput interface {
	pulumi.Input

	ToVmAffinityRuleMapOutput() VmAffinityRuleMapOutput
	ToVmAffinityRuleMapOutputWithContext(context.Context) VmAffinityRuleMapOutput
}

type VmAffinityRuleMap map[string]VmAffinityRuleInput

func (VmAffinityRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmAffinityRule)(nil)).Elem()
}

func (i VmAffinityRuleMap) ToVmAffinityRuleMapOutput() VmAffinityRuleMapOutput {
	return i.ToVmAffinityRuleMapOutputWithContext(context.Background())
}

func (i VmAffinityRuleMap) ToVmAffinityRuleMapOutputWithContext(ctx context.Context) VmAffinityRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmAffinityRuleMapOutput)
}

type VmAffinityRuleOutput struct{ *pulumi.OutputState }

func (VmAffinityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmAffinityRule)(nil)).Elem()
}

func (o VmAffinityRuleOutput) ToVmAffinityRuleOutput() VmAffinityRuleOutput {
	return o
}

func (o VmAffinityRuleOutput) ToVmAffinityRuleOutputWithContext(ctx context.Context) VmAffinityRuleOutput {
	return o
}

// True if this affinity rule is enabled
func (o VmAffinityRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// VM affinity rule name
func (o VmAffinityRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o VmAffinityRuleOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// One of 'Affinity', 'Anti-Affinity'
func (o VmAffinityRuleOutput) Polarity() pulumi.StringOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.StringOutput { return v.Polarity }).(pulumi.StringOutput)
}

// True if this affinity rule is required. When a rule is mandatory, a host failover will not power on the VM if doing so
// would violate the rule
func (o VmAffinityRuleOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.BoolPtrOutput { return v.Required }).(pulumi.BoolPtrOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o VmAffinityRuleOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

// Set of VM IDs assigned to this rule
func (o VmAffinityRuleOutput) VmIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VmAffinityRule) pulumi.StringArrayOutput { return v.VmIds }).(pulumi.StringArrayOutput)
}

type VmAffinityRuleArrayOutput struct{ *pulumi.OutputState }

func (VmAffinityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmAffinityRule)(nil)).Elem()
}

func (o VmAffinityRuleArrayOutput) ToVmAffinityRuleArrayOutput() VmAffinityRuleArrayOutput {
	return o
}

func (o VmAffinityRuleArrayOutput) ToVmAffinityRuleArrayOutputWithContext(ctx context.Context) VmAffinityRuleArrayOutput {
	return o
}

func (o VmAffinityRuleArrayOutput) Index(i pulumi.IntInput) VmAffinityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VmAffinityRule {
		return vs[0].([]*VmAffinityRule)[vs[1].(int)]
	}).(VmAffinityRuleOutput)
}

type VmAffinityRuleMapOutput struct{ *pulumi.OutputState }

func (VmAffinityRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmAffinityRule)(nil)).Elem()
}

func (o VmAffinityRuleMapOutput) ToVmAffinityRuleMapOutput() VmAffinityRuleMapOutput {
	return o
}

func (o VmAffinityRuleMapOutput) ToVmAffinityRuleMapOutputWithContext(ctx context.Context) VmAffinityRuleMapOutput {
	return o
}

func (o VmAffinityRuleMapOutput) MapIndex(k pulumi.StringInput) VmAffinityRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VmAffinityRule {
		return vs[0].(map[string]*VmAffinityRule)[vs[1].(string)]
	}).(VmAffinityRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmAffinityRuleInput)(nil)).Elem(), &VmAffinityRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmAffinityRuleArrayInput)(nil)).Elem(), VmAffinityRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmAffinityRuleMapInput)(nil)).Elem(), VmAffinityRuleMap{})
	pulumi.RegisterOutputType(VmAffinityRuleOutput{})
	pulumi.RegisterOutputType(VmAffinityRuleArrayOutput{})
	pulumi.RegisterOutputType(VmAffinityRuleMapOutput{})
}