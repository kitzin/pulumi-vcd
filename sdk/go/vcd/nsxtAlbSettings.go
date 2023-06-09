// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxtAlbSettings struct {
	pulumi.CustomResourceState

	// Edge gateway ID
	EdgeGatewayId pulumi.StringOutput `pulumi:"edgeGatewayId"`
	// Defines if ALB is enabled on Edge Gateway
	IsActive pulumi.BoolOutput `pulumi:"isActive"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25)
	ServiceNetworkSpecification pulumi.StringOutput `pulumi:"serviceNetworkSpecification"`
	// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
	SupportedFeatureSet pulumi.StringOutput `pulumi:"supportedFeatureSet"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringOutput `pulumi:"vdc"`
}

// NewNsxtAlbSettings registers a new resource with the given unique name, arguments, and options.
func NewNsxtAlbSettings(ctx *pulumi.Context,
	name string, args *NsxtAlbSettingsArgs, opts ...pulumi.ResourceOption) (*NsxtAlbSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGatewayId'")
	}
	if args.IsActive == nil {
		return nil, errors.New("invalid value for required argument 'IsActive'")
	}
	var resource NsxtAlbSettings
	err := ctx.RegisterResource("vcd:index/nsxtAlbSettings:NsxtAlbSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtAlbSettings gets an existing NsxtAlbSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtAlbSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtAlbSettingsState, opts ...pulumi.ResourceOption) (*NsxtAlbSettings, error) {
	var resource NsxtAlbSettings
	err := ctx.ReadResource("vcd:index/nsxtAlbSettings:NsxtAlbSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtAlbSettings resources.
type nsxtAlbSettingsState struct {
	// Edge gateway ID
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// Defines if ALB is enabled on Edge Gateway
	IsActive *bool `pulumi:"isActive"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25)
	ServiceNetworkSpecification *string `pulumi:"serviceNetworkSpecification"`
	// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
	SupportedFeatureSet *string `pulumi:"supportedFeatureSet"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

type NsxtAlbSettingsState struct {
	// Edge gateway ID
	EdgeGatewayId pulumi.StringPtrInput
	// Defines if ALB is enabled on Edge Gateway
	IsActive pulumi.BoolPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25)
	ServiceNetworkSpecification pulumi.StringPtrInput
	// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
	SupportedFeatureSet pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput
}

func (NsxtAlbSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbSettingsState)(nil)).Elem()
}

type nsxtAlbSettingsArgs struct {
	// Edge gateway ID
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// Defines if ALB is enabled on Edge Gateway
	IsActive bool `pulumi:"isActive"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25)
	ServiceNetworkSpecification *string `pulumi:"serviceNetworkSpecification"`
	// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
	SupportedFeatureSet *string `pulumi:"supportedFeatureSet"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NsxtAlbSettings resource.
type NsxtAlbSettingsArgs struct {
	// Edge gateway ID
	EdgeGatewayId pulumi.StringInput
	// Defines if ALB is enabled on Edge Gateway
	IsActive pulumi.BoolInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25)
	ServiceNetworkSpecification pulumi.StringPtrInput
	// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
	SupportedFeatureSet pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput
}

func (NsxtAlbSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbSettingsArgs)(nil)).Elem()
}

type NsxtAlbSettingsInput interface {
	pulumi.Input

	ToNsxtAlbSettingsOutput() NsxtAlbSettingsOutput
	ToNsxtAlbSettingsOutputWithContext(ctx context.Context) NsxtAlbSettingsOutput
}

func (*NsxtAlbSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbSettings)(nil)).Elem()
}

func (i *NsxtAlbSettings) ToNsxtAlbSettingsOutput() NsxtAlbSettingsOutput {
	return i.ToNsxtAlbSettingsOutputWithContext(context.Background())
}

func (i *NsxtAlbSettings) ToNsxtAlbSettingsOutputWithContext(ctx context.Context) NsxtAlbSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbSettingsOutput)
}

// NsxtAlbSettingsArrayInput is an input type that accepts NsxtAlbSettingsArray and NsxtAlbSettingsArrayOutput values.
// You can construct a concrete instance of `NsxtAlbSettingsArrayInput` via:
//
//	NsxtAlbSettingsArray{ NsxtAlbSettingsArgs{...} }
type NsxtAlbSettingsArrayInput interface {
	pulumi.Input

	ToNsxtAlbSettingsArrayOutput() NsxtAlbSettingsArrayOutput
	ToNsxtAlbSettingsArrayOutputWithContext(context.Context) NsxtAlbSettingsArrayOutput
}

type NsxtAlbSettingsArray []NsxtAlbSettingsInput

func (NsxtAlbSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbSettings)(nil)).Elem()
}

func (i NsxtAlbSettingsArray) ToNsxtAlbSettingsArrayOutput() NsxtAlbSettingsArrayOutput {
	return i.ToNsxtAlbSettingsArrayOutputWithContext(context.Background())
}

func (i NsxtAlbSettingsArray) ToNsxtAlbSettingsArrayOutputWithContext(ctx context.Context) NsxtAlbSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbSettingsArrayOutput)
}

// NsxtAlbSettingsMapInput is an input type that accepts NsxtAlbSettingsMap and NsxtAlbSettingsMapOutput values.
// You can construct a concrete instance of `NsxtAlbSettingsMapInput` via:
//
//	NsxtAlbSettingsMap{ "key": NsxtAlbSettingsArgs{...} }
type NsxtAlbSettingsMapInput interface {
	pulumi.Input

	ToNsxtAlbSettingsMapOutput() NsxtAlbSettingsMapOutput
	ToNsxtAlbSettingsMapOutputWithContext(context.Context) NsxtAlbSettingsMapOutput
}

type NsxtAlbSettingsMap map[string]NsxtAlbSettingsInput

func (NsxtAlbSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbSettings)(nil)).Elem()
}

func (i NsxtAlbSettingsMap) ToNsxtAlbSettingsMapOutput() NsxtAlbSettingsMapOutput {
	return i.ToNsxtAlbSettingsMapOutputWithContext(context.Background())
}

func (i NsxtAlbSettingsMap) ToNsxtAlbSettingsMapOutputWithContext(ctx context.Context) NsxtAlbSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbSettingsMapOutput)
}

type NsxtAlbSettingsOutput struct{ *pulumi.OutputState }

func (NsxtAlbSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbSettings)(nil)).Elem()
}

func (o NsxtAlbSettingsOutput) ToNsxtAlbSettingsOutput() NsxtAlbSettingsOutput {
	return o
}

func (o NsxtAlbSettingsOutput) ToNsxtAlbSettingsOutputWithContext(ctx context.Context) NsxtAlbSettingsOutput {
	return o
}

// Edge gateway ID
func (o NsxtAlbSettingsOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbSettings) pulumi.StringOutput { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// Defines if ALB is enabled on Edge Gateway
func (o NsxtAlbSettingsOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v *NsxtAlbSettings) pulumi.BoolOutput { return v.IsActive }).(pulumi.BoolOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NsxtAlbSettingsOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtAlbSettings) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Optional custom network CIDR definition for ALB Service Engine placement (VCD default is 192.168.255.1/25)
func (o NsxtAlbSettingsOutput) ServiceNetworkSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbSettings) pulumi.StringOutput { return v.ServiceNetworkSpecification }).(pulumi.StringOutput)
}

// Feature set for ALB in this Edge Gateway. One of 'STANDARD', 'PREMIUM'.
func (o NsxtAlbSettingsOutput) SupportedFeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbSettings) pulumi.StringOutput { return v.SupportedFeatureSet }).(pulumi.StringOutput)
}

// The name of VDC to use, optional if defined at provider level
//
// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
func (o NsxtAlbSettingsOutput) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbSettings) pulumi.StringOutput { return v.Vdc }).(pulumi.StringOutput)
}

type NsxtAlbSettingsArrayOutput struct{ *pulumi.OutputState }

func (NsxtAlbSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbSettings)(nil)).Elem()
}

func (o NsxtAlbSettingsArrayOutput) ToNsxtAlbSettingsArrayOutput() NsxtAlbSettingsArrayOutput {
	return o
}

func (o NsxtAlbSettingsArrayOutput) ToNsxtAlbSettingsArrayOutputWithContext(ctx context.Context) NsxtAlbSettingsArrayOutput {
	return o
}

func (o NsxtAlbSettingsArrayOutput) Index(i pulumi.IntInput) NsxtAlbSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtAlbSettings {
		return vs[0].([]*NsxtAlbSettings)[vs[1].(int)]
	}).(NsxtAlbSettingsOutput)
}

type NsxtAlbSettingsMapOutput struct{ *pulumi.OutputState }

func (NsxtAlbSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbSettings)(nil)).Elem()
}

func (o NsxtAlbSettingsMapOutput) ToNsxtAlbSettingsMapOutput() NsxtAlbSettingsMapOutput {
	return o
}

func (o NsxtAlbSettingsMapOutput) ToNsxtAlbSettingsMapOutputWithContext(ctx context.Context) NsxtAlbSettingsMapOutput {
	return o
}

func (o NsxtAlbSettingsMapOutput) MapIndex(k pulumi.StringInput) NsxtAlbSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtAlbSettings {
		return vs[0].(map[string]*NsxtAlbSettings)[vs[1].(string)]
	}).(NsxtAlbSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbSettingsInput)(nil)).Elem(), &NsxtAlbSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbSettingsArrayInput)(nil)).Elem(), NsxtAlbSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbSettingsMapInput)(nil)).Elem(), NsxtAlbSettingsMap{})
	pulumi.RegisterOutputType(NsxtAlbSettingsOutput{})
	pulumi.RegisterOutputType(NsxtAlbSettingsArrayOutput{})
	pulumi.RegisterOutputType(NsxtAlbSettingsMapOutput{})
}
