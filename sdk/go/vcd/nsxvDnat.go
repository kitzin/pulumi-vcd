// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxvDnat struct {
	pulumi.CustomResourceState

	// NAT rule description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Edge gateway name in which NAT Rule is located
	EdgeGateway pulumi.StringOutput `pulumi:"edgeGateway"`
	// Whether the rule should be enabled. Default 'true'
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
	// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
	// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType pulumi.StringPtrOutput `pulumi:"icmpType"`
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled pulumi.BoolPtrOutput `pulumi:"loggingEnabled"`
	// Org or external network name
	NetworkName pulumi.StringOutput `pulumi:"networkName"`
	// Network type. One of 'ext', 'org'
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress pulumi.StringOutput `pulumi:"originalAddress"`
	// Original port. This is the destination port for DNAT rules
	OriginalPort pulumi.StringPtrOutput `pulumi:"originalPort"`
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Optional. Allows to set custom rule tag
	RuleTag pulumi.IntOutput `pulumi:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'
	RuleType pulumi.StringOutput `pulumi:"ruleType"`
	// Translated address or address range
	TranslatedAddress pulumi.StringPtrOutput `pulumi:"translatedAddress"`
	// Translated port
	TranslatedPort pulumi.StringPtrOutput `pulumi:"translatedPort"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewNsxvDnat registers a new resource with the given unique name, arguments, and options.
func NewNsxvDnat(ctx *pulumi.Context,
	name string, args *NsxvDnatArgs, opts ...pulumi.ResourceOption) (*NsxvDnat, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGateway == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGateway'")
	}
	if args.NetworkName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkName'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	if args.OriginalAddress == nil {
		return nil, errors.New("invalid value for required argument 'OriginalAddress'")
	}
	var resource NsxvDnat
	err := ctx.RegisterResource("vcd:index/nsxvDnat:NsxvDnat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxvDnat gets an existing NsxvDnat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxvDnat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxvDnatState, opts ...pulumi.ResourceOption) (*NsxvDnat, error) {
	var resource NsxvDnat
	err := ctx.ReadResource("vcd:index/nsxvDnat:NsxvDnat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxvDnat resources.
type nsxvDnatState struct {
	// NAT rule description
	Description *string `pulumi:"description"`
	// Edge gateway name in which NAT Rule is located
	EdgeGateway *string `pulumi:"edgeGateway"`
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `pulumi:"enabled"`
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
	// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
	// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType *string `pulumi:"icmpType"`
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `pulumi:"loggingEnabled"`
	// Org or external network name
	NetworkName *string `pulumi:"networkName"`
	// Network type. One of 'ext', 'org'
	NetworkType *string `pulumi:"networkType"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress *string `pulumi:"originalAddress"`
	// Original port. This is the destination port for DNAT rules
	OriginalPort *string `pulumi:"originalPort"`
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol *string `pulumi:"protocol"`
	// Optional. Allows to set custom rule tag
	RuleTag *int `pulumi:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `pulumi:"ruleType"`
	// Translated address or address range
	TranslatedAddress *string `pulumi:"translatedAddress"`
	// Translated port
	TranslatedPort *string `pulumi:"translatedPort"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type NsxvDnatState struct {
	// NAT rule description
	Description pulumi.StringPtrInput
	// Edge gateway name in which NAT Rule is located
	EdgeGateway pulumi.StringPtrInput
	// Whether the rule should be enabled. Default 'true'
	Enabled pulumi.BoolPtrInput
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
	// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
	// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType pulumi.StringPtrInput
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled pulumi.BoolPtrInput
	// Org or external network name
	NetworkName pulumi.StringPtrInput
	// Network type. One of 'ext', 'org'
	NetworkType pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress pulumi.StringPtrInput
	// Original port. This is the destination port for DNAT rules
	OriginalPort pulumi.StringPtrInput
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol pulumi.StringPtrInput
	// Optional. Allows to set custom rule tag
	RuleTag pulumi.IntPtrInput
	// Read only. Possible values 'user', 'internal_high'
	RuleType pulumi.StringPtrInput
	// Translated address or address range
	TranslatedAddress pulumi.StringPtrInput
	// Translated port
	TranslatedPort pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NsxvDnatState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxvDnatState)(nil)).Elem()
}

type nsxvDnatArgs struct {
	// NAT rule description
	Description *string `pulumi:"description"`
	// Edge gateway name in which NAT Rule is located
	EdgeGateway string `pulumi:"edgeGateway"`
	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `pulumi:"enabled"`
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
	// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
	// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType *string `pulumi:"icmpType"`
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `pulumi:"loggingEnabled"`
	// Org or external network name
	NetworkName string `pulumi:"networkName"`
	// Network type. One of 'ext', 'org'
	NetworkType string `pulumi:"networkType"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress string `pulumi:"originalAddress"`
	// Original port. This is the destination port for DNAT rules
	OriginalPort *string `pulumi:"originalPort"`
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol *string `pulumi:"protocol"`
	// Optional. Allows to set custom rule tag
	RuleTag *int `pulumi:"ruleTag"`
	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `pulumi:"ruleType"`
	// Translated address or address range
	TranslatedAddress *string `pulumi:"translatedAddress"`
	// Translated port
	TranslatedPort *string `pulumi:"translatedPort"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NsxvDnat resource.
type NsxvDnatArgs struct {
	// NAT rule description
	Description pulumi.StringPtrInput
	// Edge gateway name in which NAT Rule is located
	EdgeGateway pulumi.StringInput
	// Whether the rule should be enabled. Default 'true'
	Enabled pulumi.BoolPtrInput
	// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
	// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
	// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
	IcmpType pulumi.StringPtrInput
	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled pulumi.BoolPtrInput
	// Org or external network name
	NetworkName pulumi.StringInput
	// Network type. One of 'ext', 'org'
	NetworkType pulumi.StringInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Original address or address range. This is the the destination address for DNAT rules.
	OriginalAddress pulumi.StringInput
	// Original port. This is the destination port for DNAT rules
	OriginalPort pulumi.StringPtrInput
	// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
	Protocol pulumi.StringPtrInput
	// Optional. Allows to set custom rule tag
	RuleTag pulumi.IntPtrInput
	// Read only. Possible values 'user', 'internal_high'
	RuleType pulumi.StringPtrInput
	// Translated address or address range
	TranslatedAddress pulumi.StringPtrInput
	// Translated port
	TranslatedPort pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NsxvDnatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxvDnatArgs)(nil)).Elem()
}

type NsxvDnatInput interface {
	pulumi.Input

	ToNsxvDnatOutput() NsxvDnatOutput
	ToNsxvDnatOutputWithContext(ctx context.Context) NsxvDnatOutput
}

func (*NsxvDnat) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxvDnat)(nil)).Elem()
}

func (i *NsxvDnat) ToNsxvDnatOutput() NsxvDnatOutput {
	return i.ToNsxvDnatOutputWithContext(context.Background())
}

func (i *NsxvDnat) ToNsxvDnatOutputWithContext(ctx context.Context) NsxvDnatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvDnatOutput)
}

// NsxvDnatArrayInput is an input type that accepts NsxvDnatArray and NsxvDnatArrayOutput values.
// You can construct a concrete instance of `NsxvDnatArrayInput` via:
//
//	NsxvDnatArray{ NsxvDnatArgs{...} }
type NsxvDnatArrayInput interface {
	pulumi.Input

	ToNsxvDnatArrayOutput() NsxvDnatArrayOutput
	ToNsxvDnatArrayOutputWithContext(context.Context) NsxvDnatArrayOutput
}

type NsxvDnatArray []NsxvDnatInput

func (NsxvDnatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxvDnat)(nil)).Elem()
}

func (i NsxvDnatArray) ToNsxvDnatArrayOutput() NsxvDnatArrayOutput {
	return i.ToNsxvDnatArrayOutputWithContext(context.Background())
}

func (i NsxvDnatArray) ToNsxvDnatArrayOutputWithContext(ctx context.Context) NsxvDnatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvDnatArrayOutput)
}

// NsxvDnatMapInput is an input type that accepts NsxvDnatMap and NsxvDnatMapOutput values.
// You can construct a concrete instance of `NsxvDnatMapInput` via:
//
//	NsxvDnatMap{ "key": NsxvDnatArgs{...} }
type NsxvDnatMapInput interface {
	pulumi.Input

	ToNsxvDnatMapOutput() NsxvDnatMapOutput
	ToNsxvDnatMapOutputWithContext(context.Context) NsxvDnatMapOutput
}

type NsxvDnatMap map[string]NsxvDnatInput

func (NsxvDnatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxvDnat)(nil)).Elem()
}

func (i NsxvDnatMap) ToNsxvDnatMapOutput() NsxvDnatMapOutput {
	return i.ToNsxvDnatMapOutputWithContext(context.Background())
}

func (i NsxvDnatMap) ToNsxvDnatMapOutputWithContext(ctx context.Context) NsxvDnatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxvDnatMapOutput)
}

type NsxvDnatOutput struct{ *pulumi.OutputState }

func (NsxvDnatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxvDnat)(nil)).Elem()
}

func (o NsxvDnatOutput) ToNsxvDnatOutput() NsxvDnatOutput {
	return o
}

func (o NsxvDnatOutput) ToNsxvDnatOutputWithContext(ctx context.Context) NsxvDnatOutput {
	return o
}

// NAT rule description
func (o NsxvDnatOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Edge gateway name in which NAT Rule is located
func (o NsxvDnatOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringOutput { return v.EdgeGateway }).(pulumi.StringOutput)
}

// Whether the rule should be enabled. Default 'true'
func (o NsxvDnatOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
func (o NsxvDnatOutput) IcmpType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.IcmpType }).(pulumi.StringPtrOutput)
}

// Whether logging should be enabled for this rule. Default 'false'
func (o NsxvDnatOutput) LoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.BoolPtrOutput { return v.LoggingEnabled }).(pulumi.BoolPtrOutput)
}

// Org or external network name
func (o NsxvDnatOutput) NetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringOutput { return v.NetworkName }).(pulumi.StringOutput)
}

// Network type. One of 'ext', 'org'
func (o NsxvDnatOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NsxvDnatOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Original address or address range. This is the the destination address for DNAT rules.
func (o NsxvDnatOutput) OriginalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringOutput { return v.OriginalAddress }).(pulumi.StringOutput)
}

// Original port. This is the destination port for DNAT rules
func (o NsxvDnatOutput) OriginalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.OriginalPort }).(pulumi.StringPtrOutput)
}

// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
func (o NsxvDnatOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// Optional. Allows to set custom rule tag
func (o NsxvDnatOutput) RuleTag() pulumi.IntOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.IntOutput { return v.RuleTag }).(pulumi.IntOutput)
}

// Read only. Possible values 'user', 'internal_high'
func (o NsxvDnatOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

// Translated address or address range
func (o NsxvDnatOutput) TranslatedAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.TranslatedAddress }).(pulumi.StringPtrOutput)
}

// Translated port
func (o NsxvDnatOutput) TranslatedPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.TranslatedPort }).(pulumi.StringPtrOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o NsxvDnatOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxvDnat) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type NsxvDnatArrayOutput struct{ *pulumi.OutputState }

func (NsxvDnatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxvDnat)(nil)).Elem()
}

func (o NsxvDnatArrayOutput) ToNsxvDnatArrayOutput() NsxvDnatArrayOutput {
	return o
}

func (o NsxvDnatArrayOutput) ToNsxvDnatArrayOutputWithContext(ctx context.Context) NsxvDnatArrayOutput {
	return o
}

func (o NsxvDnatArrayOutput) Index(i pulumi.IntInput) NsxvDnatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxvDnat {
		return vs[0].([]*NsxvDnat)[vs[1].(int)]
	}).(NsxvDnatOutput)
}

type NsxvDnatMapOutput struct{ *pulumi.OutputState }

func (NsxvDnatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxvDnat)(nil)).Elem()
}

func (o NsxvDnatMapOutput) ToNsxvDnatMapOutput() NsxvDnatMapOutput {
	return o
}

func (o NsxvDnatMapOutput) ToNsxvDnatMapOutputWithContext(ctx context.Context) NsxvDnatMapOutput {
	return o
}

func (o NsxvDnatMapOutput) MapIndex(k pulumi.StringInput) NsxvDnatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxvDnat {
		return vs[0].(map[string]*NsxvDnat)[vs[1].(string)]
	}).(NsxvDnatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvDnatInput)(nil)).Elem(), &NsxvDnat{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvDnatArrayInput)(nil)).Elem(), NsxvDnatArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxvDnatMapInput)(nil)).Elem(), NsxvDnatMap{})
	pulumi.RegisterOutputType(NsxvDnatOutput{})
	pulumi.RegisterOutputType(NsxvDnatArrayOutput{})
	pulumi.RegisterOutputType(NsxvDnatMapOutput{})
}
