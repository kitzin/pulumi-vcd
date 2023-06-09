// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxvFirewallRule struct {
	pulumi.CustomResourceState

	// This firewall rule will be inserted above the referred one
	AboveRuleId pulumi.StringPtrOutput `pulumi:"aboveRuleId"`
	// 'accept' or 'deny'. Default 'accept'
	Action      pulumi.StringPtrOutput            `pulumi:"action"`
	Destination NsxvFirewallRuleDestinationOutput `pulumi:"destination"`
	// Edge gateway name in which Firewall Rule is located
	EdgeGateway pulumi.StringOutput `pulumi:"edgeGateway"`
	// Whether the rule should be enabled. Default 'true'
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled pulumi.BoolPtrOutput `pulumi:"loggingEnabled"`
	// Firewall rule name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Optional. Allows to set custom rule tag
	RuleTag pulumi.IntOutput `pulumi:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'
	RuleType pulumi.StringOutput                `pulumi:"ruleType"`
	Services NsxvFirewallRuleServiceArrayOutput `pulumi:"services"`
	Source   NsxvFirewallRuleSourceOutput       `pulumi:"source"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewNsxvFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewNsxvFirewallRule(ctx *pulumi.Context,
	name string, args *NsxvFirewallRuleArgs, opts ...pulumi.ResourceOption) (*NsxvFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.EdgeGateway == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGateway'")
	}
	if args.Services == nil {
		return nil, errors.New("invalid value for required argument 'Services'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource NsxvFirewallRule
	err := ctx.RegisterResource("vcd:index/nsxvFirewallRule:NsxvFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxvFirewallRule gets an existing NsxvFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxvFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxvFirewallRuleState, opts ...pulumi.ResourceOption) (*NsxvFirewallRule, error) {
	var resource NsxvFirewallRule
	err := ctx.ReadResource("vcd:index/nsxvFirewallRule:NsxvFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxvFirewallRule resources.
type nsxvFirewallRuleState struct {
	// This firewall rule will be inserted above the referred one
	AboveRuleId *string `pulumi:"aboveRuleId"`
	// 'accept' or 'deny'. Default 'accept'
	Action      *string                      `pulumi:"action"`
	Destination *NsxvFirewallRuleDestination `pulumi:"destination"`
	// Edge gateway name in which Firewall Rule is located
	EdgeGateway *string `pulumi:"edgeGateway"`
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `pulumi:"enabled"`
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `pulumi:"loggingEnabled"`
	// Firewall rule name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Optional. Allows to set custom rule tag
	RuleTag *int `pulumi:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string                   `pulumi:"ruleType"`
	Services []NsxvFirewallRuleService `pulumi:"services"`
	Source   *NsxvFirewallRuleSource   `pulumi:"source"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type NsxvFirewallRuleState struct {
	// This firewall rule will be inserted above the referred one
	AboveRuleId pulumi.StringPtrInput
	// 'accept' or 'deny'. Default 'accept'
	Action      pulumi.StringPtrInput
	Destination NsxvFirewallRuleDestinationPtrInput
	// Edge gateway name in which Firewall Rule is located
	EdgeGateway pulumi.StringPtrInput
	// Whether the rule should be enabled. Default 'true'
	Enabled pulumi.BoolPtrInput
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled pulumi.BoolPtrInput
	// Firewall rule name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Optional. Allows to set custom rule tag
	RuleTag pulumi.IntPtrInput
	// Read only. Possible values 'user', 'internal_high'
	RuleType pulumi.StringPtrInput
	Services NsxvFirewallRuleServiceArrayInput
	Source   NsxvFirewallRuleSourcePtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NsxvFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxvFirewallRuleState)(nil)).Elem()
}

type nsxvFirewallRuleArgs struct {
	// This firewall rule will be inserted above the referred one
	AboveRuleId *string `pulumi:"aboveRuleId"`
	// 'accept' or 'deny'. Default 'accept'
	Action      *string                     `pulumi:"action"`
	Destination NsxvFirewallRuleDestination `pulumi:"destination"`
	// Edge gateway name in which Firewall Rule is located
	EdgeGateway string `pulumi:"edgeGateway"`
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `pulumi:"enabled"`
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `pulumi:"loggingEnabled"`
	// Firewall rule name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Optional. Allows to set custom rule tag
	RuleTag *int `pulumi:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string                   `pulumi:"ruleType"`
	Services []NsxvFirewallRuleService `pulumi:"services"`
	Source   NsxvFirewallRuleSource    `pulumi:"source"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NsxvFirewallRule resource.
type NsxvFirewallRuleArgs struct {
	// This firewall rule will be inserted above the referred one
	AboveRuleId pulumi.StringPtrInput
	// 'accept' or 'deny'. Default 'accept'
	Action      pulumi.StringPtrInput
	Destination NsxvFirewallRuleDestinationInput
	// Edge gateway name in which Firewall Rule is located
	EdgeGateway pulumi.StringInput
	// Whether the rule should be enabled. Default 'true'
	Enabled pulumi.BoolPtrInput
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled pulumi.BoolPtrInput
	// Firewall rule name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Optional. Allows to set custom rule tag
	RuleTag pulumi.IntPtrInput
	// Read only. Possible values 'user', 'internal_high'
	RuleType pulumi.StringPtrInput
	Services NsxvFirewallRuleServiceArrayInput
	Source   NsxvFirewallRuleSourceInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NsxvFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxvFirewallRuleArgs)(nil)).Elem()
}

type NsxvFirewallRuleInput interface {
	pulumi.Input

	ToNsxvFirewallRuleOutput() NsxvFirewallRuleOutput
	ToNsxvFirewallRuleOutputWithContext(ctx context.Context) NsxvFirewallRuleOutput
}

func (*NsxvFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxvFirewallRule)(nil)).Elem()
}

func (i *NsxvFirewallRule) ToNsxvFirewallRuleOutput() NsxvFirewallRuleOutput {
	return i.ToNsxvFirewallRuleOutputWithContext(context.Background())
}

func (i *NsxvFirewallRule) ToNsxvFirewallRuleOutputWithContext(ctx context.Context) NsxvFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvFirewallRuleOutput)
}

// NsxvFirewallRuleArrayInput is an input type that accepts NsxvFirewallRuleArray and NsxvFirewallRuleArrayOutput values.
// You can construct a concrete instance of `NsxvFirewallRuleArrayInput` via:
//
//	NsxvFirewallRuleArray{ NsxvFirewallRuleArgs{...} }
type NsxvFirewallRuleArrayInput interface {
	pulumi.Input

	ToNsxvFirewallRuleArrayOutput() NsxvFirewallRuleArrayOutput
	ToNsxvFirewallRuleArrayOutputWithContext(context.Context) NsxvFirewallRuleArrayOutput
}

type NsxvFirewallRuleArray []NsxvFirewallRuleInput

func (NsxvFirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxvFirewallRule)(nil)).Elem()
}

func (i NsxvFirewallRuleArray) ToNsxvFirewallRuleArrayOutput() NsxvFirewallRuleArrayOutput {
	return i.ToNsxvFirewallRuleArrayOutputWithContext(context.Background())
}

func (i NsxvFirewallRuleArray) ToNsxvFirewallRuleArrayOutputWithContext(ctx context.Context) NsxvFirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvFirewallRuleArrayOutput)
}

// NsxvFirewallRuleMapInput is an input type that accepts NsxvFirewallRuleMap and NsxvFirewallRuleMapOutput values.
// You can construct a concrete instance of `NsxvFirewallRuleMapInput` via:
//
//	NsxvFirewallRuleMap{ "key": NsxvFirewallRuleArgs{...} }
type NsxvFirewallRuleMapInput interface {
	pulumi.Input

	ToNsxvFirewallRuleMapOutput() NsxvFirewallRuleMapOutput
	ToNsxvFirewallRuleMapOutputWithContext(context.Context) NsxvFirewallRuleMapOutput
}

type NsxvFirewallRuleMap map[string]NsxvFirewallRuleInput

func (NsxvFirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxvFirewallRule)(nil)).Elem()
}

func (i NsxvFirewallRuleMap) ToNsxvFirewallRuleMapOutput() NsxvFirewallRuleMapOutput {
	return i.ToNsxvFirewallRuleMapOutputWithContext(context.Background())
}

func (i NsxvFirewallRuleMap) ToNsxvFirewallRuleMapOutputWithContext(ctx context.Context) NsxvFirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvFirewallRuleMapOutput)
}

type NsxvFirewallRuleOutput struct{ *pulumi.OutputState }

func (NsxvFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxvFirewallRule)(nil)).Elem()
}

func (o NsxvFirewallRuleOutput) ToNsxvFirewallRuleOutput() NsxvFirewallRuleOutput {
	return o
}

func (o NsxvFirewallRuleOutput) ToNsxvFirewallRuleOutputWithContext(ctx context.Context) NsxvFirewallRuleOutput {
	return o
}

// This firewall rule will be inserted above the referred one
func (o NsxvFirewallRuleOutput) AboveRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringPtrOutput { return v.AboveRuleId }).(pulumi.StringPtrOutput)
}

// 'accept' or 'deny'. Default 'accept'
func (o NsxvFirewallRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NsxvFirewallRuleOutput) Destination() NsxvFirewallRuleDestinationOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) NsxvFirewallRuleDestinationOutput { return v.Destination }).(NsxvFirewallRuleDestinationOutput)
}

// Edge gateway name in which Firewall Rule is located
func (o NsxvFirewallRuleOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringOutput { return v.EdgeGateway }).(pulumi.StringOutput)
}

// Whether the rule should be enabled. Default 'true'
func (o NsxvFirewallRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Whether logging should be enabled for this rule. Default 'false'
func (o NsxvFirewallRuleOutput) LoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.BoolPtrOutput { return v.LoggingEnabled }).(pulumi.BoolPtrOutput)
}

// Firewall rule name
func (o NsxvFirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NsxvFirewallRuleOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Optional. Allows to set custom rule tag
func (o NsxvFirewallRuleOutput) RuleTag() pulumi.IntOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.IntOutput { return v.RuleTag }).(pulumi.IntOutput)
}

// Read only. Possible values 'user', 'internal_high'
func (o NsxvFirewallRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

func (o NsxvFirewallRuleOutput) Services() NsxvFirewallRuleServiceArrayOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) NsxvFirewallRuleServiceArrayOutput { return v.Services }).(NsxvFirewallRuleServiceArrayOutput)
}

func (o NsxvFirewallRuleOutput) Source() NsxvFirewallRuleSourceOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) NsxvFirewallRuleSourceOutput { return v.Source }).(NsxvFirewallRuleSourceOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o NsxvFirewallRuleOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvFirewallRule) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type NsxvFirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (NsxvFirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxvFirewallRule)(nil)).Elem()
}

func (o NsxvFirewallRuleArrayOutput) ToNsxvFirewallRuleArrayOutput() NsxvFirewallRuleArrayOutput {
	return o
}

func (o NsxvFirewallRuleArrayOutput) ToNsxvFirewallRuleArrayOutputWithContext(ctx context.Context) NsxvFirewallRuleArrayOutput {
	return o
}

func (o NsxvFirewallRuleArrayOutput) Index(i pulumi.IntInput) NsxvFirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxvFirewallRule {
		return vs[0].([]*NsxvFirewallRule)[vs[1].(int)]
	}).(NsxvFirewallRuleOutput)
}

type NsxvFirewallRuleMapOutput struct{ *pulumi.OutputState }

func (NsxvFirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxvFirewallRule)(nil)).Elem()
}

func (o NsxvFirewallRuleMapOutput) ToNsxvFirewallRuleMapOutput() NsxvFirewallRuleMapOutput {
	return o
}

func (o NsxvFirewallRuleMapOutput) ToNsxvFirewallRuleMapOutputWithContext(ctx context.Context) NsxvFirewallRuleMapOutput {
	return o
}

func (o NsxvFirewallRuleMapOutput) MapIndex(k pulumi.StringInput) NsxvFirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxvFirewallRule {
		return vs[0].(map[string]*NsxvFirewallRule)[vs[1].(string)]
	}).(NsxvFirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvFirewallRuleInput)(nil)).Elem(), &NsxvFirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvFirewallRuleArrayInput)(nil)).Elem(), NsxvFirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvFirewallRuleMapInput)(nil)).Elem(), NsxvFirewallRuleMap{})
	pulumi.RegisterOutputType(NsxvFirewallRuleOutput{})
	pulumi.RegisterOutputType(NsxvFirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(NsxvFirewallRuleMapOutput{})
}
