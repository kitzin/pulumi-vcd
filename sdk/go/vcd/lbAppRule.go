// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LbAppRule struct {
	pulumi.CustomResourceState

	// Edge gateway name in which the LB Application Rule is located
	EdgeGateway pulumi.StringOutput `pulumi:"edgeGateway"`
	// Unique LB Application Rule name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// The script for the application rule. Note - you may find HEREDOC useful to pass multiline strings
	Script pulumi.StringOutput `pulumi:"script"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewLbAppRule registers a new resource with the given unique name, arguments, and options.
func NewLbAppRule(ctx *pulumi.Context,
	name string, args *LbAppRuleArgs, opts ...pulumi.ResourceOption) (*LbAppRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGateway == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGateway'")
	}
	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	var resource LbAppRule
	err := ctx.RegisterResource("vcd:index/lbAppRule:LbAppRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbAppRule gets an existing LbAppRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbAppRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbAppRuleState, opts ...pulumi.ResourceOption) (*LbAppRule, error) {
	var resource LbAppRule
	err := ctx.ReadResource("vcd:index/lbAppRule:LbAppRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbAppRule resources.
type lbAppRuleState struct {
	// Edge gateway name in which the LB Application Rule is located
	EdgeGateway *string `pulumi:"edgeGateway"`
	// Unique LB Application Rule name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// The script for the application rule. Note - you may find HEREDOC useful to pass multiline strings
	Script *string `pulumi:"script"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type LbAppRuleState struct {
	// Edge gateway name in which the LB Application Rule is located
	EdgeGateway pulumi.StringPtrInput
	// Unique LB Application Rule name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// The script for the application rule. Note - you may find HEREDOC useful to pass multiline strings
	Script pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (LbAppRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbAppRuleState)(nil)).Elem()
}

type lbAppRuleArgs struct {
	// Edge gateway name in which the LB Application Rule is located
	EdgeGateway string `pulumi:"edgeGateway"`
	// Unique LB Application Rule name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// The script for the application rule. Note - you may find HEREDOC useful to pass multiline strings
	Script string `pulumi:"script"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a LbAppRule resource.
type LbAppRuleArgs struct {
	// Edge gateway name in which the LB Application Rule is located
	EdgeGateway pulumi.StringInput
	// Unique LB Application Rule name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// The script for the application rule. Note - you may find HEREDOC useful to pass multiline strings
	Script pulumi.StringInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (LbAppRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbAppRuleArgs)(nil)).Elem()
}

type LbAppRuleInput interface {
	pulumi.Input

	ToLbAppRuleOutput() LbAppRuleOutput
	ToLbAppRuleOutputWithContext(ctx context.Context) LbAppRuleOutput
}

func (*LbAppRule) ElementType() reflect.Type {
	return reflect.TypeOf((**LbAppRule)(nil)).Elem()
}

func (i *LbAppRule) ToLbAppRuleOutput() LbAppRuleOutput {
	return i.ToLbAppRuleOutputWithContext(context.Background())
}

func (i *LbAppRule) ToLbAppRuleOutputWithContext(ctx context.Context) LbAppRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbAppRuleOutput)
}

// LbAppRuleArrayInput is an input type that accepts LbAppRuleArray and LbAppRuleArrayOutput values.
// You can construct a concrete instance of `LbAppRuleArrayInput` via:
//
//	LbAppRuleArray{ LbAppRuleArgs{...} }
type LbAppRuleArrayInput interface {
	pulumi.Input

	ToLbAppRuleArrayOutput() LbAppRuleArrayOutput
	ToLbAppRuleArrayOutputWithContext(context.Context) LbAppRuleArrayOutput
}

type LbAppRuleArray []LbAppRuleInput

func (LbAppRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbAppRule)(nil)).Elem()
}

func (i LbAppRuleArray) ToLbAppRuleArrayOutput() LbAppRuleArrayOutput {
	return i.ToLbAppRuleArrayOutputWithContext(context.Background())
}

func (i LbAppRuleArray) ToLbAppRuleArrayOutputWithContext(ctx context.Context) LbAppRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbAppRuleArrayOutput)
}

// LbAppRuleMapInput is an input type that accepts LbAppRuleMap and LbAppRuleMapOutput values.
// You can construct a concrete instance of `LbAppRuleMapInput` via:
//
//	LbAppRuleMap{ "key": LbAppRuleArgs{...} }
type LbAppRuleMapInput interface {
	pulumi.Input

	ToLbAppRuleMapOutput() LbAppRuleMapOutput
	ToLbAppRuleMapOutputWithContext(context.Context) LbAppRuleMapOutput
}

type LbAppRuleMap map[string]LbAppRuleInput

func (LbAppRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbAppRule)(nil)).Elem()
}

func (i LbAppRuleMap) ToLbAppRuleMapOutput() LbAppRuleMapOutput {
	return i.ToLbAppRuleMapOutputWithContext(context.Background())
}

func (i LbAppRuleMap) ToLbAppRuleMapOutputWithContext(ctx context.Context) LbAppRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbAppRuleMapOutput)
}

type LbAppRuleOutput struct{ *pulumi.OutputState }

func (LbAppRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbAppRule)(nil)).Elem()
}

func (o LbAppRuleOutput) ToLbAppRuleOutput() LbAppRuleOutput {
	return o
}

func (o LbAppRuleOutput) ToLbAppRuleOutputWithContext(ctx context.Context) LbAppRuleOutput {
	return o
}

// Edge gateway name in which the LB Application Rule is located
func (o LbAppRuleOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAppRule) pulumi.StringOutput { return v.EdgeGateway }).(pulumi.StringOutput)
}

// Unique LB Application Rule name
func (o LbAppRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAppRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o LbAppRuleOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbAppRule) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// The script for the application rule. Note - you may find HEREDOC useful to pass multiline strings
func (o LbAppRuleOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAppRule) pulumi.StringOutput { return v.Script }).(pulumi.StringOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o LbAppRuleOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbAppRule) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type LbAppRuleArrayOutput struct{ *pulumi.OutputState }

func (LbAppRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbAppRule)(nil)).Elem()
}

func (o LbAppRuleArrayOutput) ToLbAppRuleArrayOutput() LbAppRuleArrayOutput {
	return o
}

func (o LbAppRuleArrayOutput) ToLbAppRuleArrayOutputWithContext(ctx context.Context) LbAppRuleArrayOutput {
	return o
}

func (o LbAppRuleArrayOutput) Index(i pulumi.IntInput) LbAppRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbAppRule {
		return vs[0].([]*LbAppRule)[vs[1].(int)]
	}).(LbAppRuleOutput)
}

type LbAppRuleMapOutput struct{ *pulumi.OutputState }

func (LbAppRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbAppRule)(nil)).Elem()
}

func (o LbAppRuleMapOutput) ToLbAppRuleMapOutput() LbAppRuleMapOutput {
	return o
}

func (o LbAppRuleMapOutput) ToLbAppRuleMapOutputWithContext(ctx context.Context) LbAppRuleMapOutput {
	return o
}

func (o LbAppRuleMapOutput) MapIndex(k pulumi.StringInput) LbAppRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbAppRule {
		return vs[0].(map[string]*LbAppRule)[vs[1].(string)]
	}).(LbAppRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbAppRuleInput)(nil)).Elem(), &LbAppRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbAppRuleArrayInput)(nil)).Elem(), LbAppRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbAppRuleMapInput)(nil)).Elem(), LbAppRuleMap{})
	pulumi.RegisterOutputType(LbAppRuleOutput{})
	pulumi.RegisterOutputType(LbAppRuleArrayOutput{})
	pulumi.RegisterOutputType(LbAppRuleMapOutput{})
}