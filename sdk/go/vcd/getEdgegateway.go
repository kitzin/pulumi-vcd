// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEdgegateway(ctx *pulumi.Context, args *GetEdgegatewayArgs, opts ...pulumi.InvokeOption) (*GetEdgegatewayResult, error) {
	var rv GetEdgegatewayResult
	err := ctx.Invoke("vcd:index/getEdgegateway:getEdgegateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEdgegateway.
type GetEdgegatewayArgs struct {
	Filter *GetEdgegatewayFilter `pulumi:"filter"`
	Name   *string               `pulumi:"name"`
	Org    *string               `pulumi:"org"`
	Vdc    *string               `pulumi:"vdc"`
}

// A collection of values returned by getEdgegateway.
type GetEdgegatewayResult struct {
	Configuration               string                          `pulumi:"configuration"`
	DefaultExternalNetworkIp    string                          `pulumi:"defaultExternalNetworkIp"`
	Description                 string                          `pulumi:"description"`
	DistributedRouting          bool                            `pulumi:"distributedRouting"`
	ExternalNetworkIps          []string                        `pulumi:"externalNetworkIps"`
	ExternalNetworks            []GetEdgegatewayExternalNetwork `pulumi:"externalNetworks"`
	Filter                      *GetEdgegatewayFilter           `pulumi:"filter"`
	FipsModeEnabled             bool                            `pulumi:"fipsModeEnabled"`
	FwDefaultRuleAction         string                          `pulumi:"fwDefaultRuleAction"`
	FwDefaultRuleLoggingEnabled bool                            `pulumi:"fwDefaultRuleLoggingEnabled"`
	FwEnabled                   bool                            `pulumi:"fwEnabled"`
	HaEnabled                   bool                            `pulumi:"haEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string  `pulumi:"id"`
	LbAccelerationEnabled      bool    `pulumi:"lbAccelerationEnabled"`
	LbEnabled                  bool    `pulumi:"lbEnabled"`
	LbLoggingEnabled           bool    `pulumi:"lbLoggingEnabled"`
	LbLoglevel                 string  `pulumi:"lbLoglevel"`
	Name                       *string `pulumi:"name"`
	Org                        *string `pulumi:"org"`
	UseDefaultRouteForDnsRelay bool    `pulumi:"useDefaultRouteForDnsRelay"`
	Vdc                        *string `pulumi:"vdc"`
}

func GetEdgegatewayOutput(ctx *pulumi.Context, args GetEdgegatewayOutputArgs, opts ...pulumi.InvokeOption) GetEdgegatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEdgegatewayResult, error) {
			args := v.(GetEdgegatewayArgs)
			r, err := GetEdgegateway(ctx, &args, opts...)
			var s GetEdgegatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEdgegatewayResultOutput)
}

// A collection of arguments for invoking getEdgegateway.
type GetEdgegatewayOutputArgs struct {
	Filter GetEdgegatewayFilterPtrInput `pulumi:"filter"`
	Name   pulumi.StringPtrInput        `pulumi:"name"`
	Org    pulumi.StringPtrInput        `pulumi:"org"`
	Vdc    pulumi.StringPtrInput        `pulumi:"vdc"`
}

func (GetEdgegatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgegatewayArgs)(nil)).Elem()
}

// A collection of values returned by getEdgegateway.
type GetEdgegatewayResultOutput struct{ *pulumi.OutputState }

func (GetEdgegatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgegatewayResult)(nil)).Elem()
}

func (o GetEdgegatewayResultOutput) ToGetEdgegatewayResultOutput() GetEdgegatewayResultOutput {
	return o
}

func (o GetEdgegatewayResultOutput) ToGetEdgegatewayResultOutputWithContext(ctx context.Context) GetEdgegatewayResultOutput {
	return o
}

func (o GetEdgegatewayResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) string { return v.Configuration }).(pulumi.StringOutput)
}

func (o GetEdgegatewayResultOutput) DefaultExternalNetworkIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) string { return v.DefaultExternalNetworkIp }).(pulumi.StringOutput)
}

func (o GetEdgegatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetEdgegatewayResultOutput) DistributedRouting() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.DistributedRouting }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) ExternalNetworkIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) []string { return v.ExternalNetworkIps }).(pulumi.StringArrayOutput)
}

func (o GetEdgegatewayResultOutput) ExternalNetworks() GetEdgegatewayExternalNetworkArrayOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) []GetEdgegatewayExternalNetwork { return v.ExternalNetworks }).(GetEdgegatewayExternalNetworkArrayOutput)
}

func (o GetEdgegatewayResultOutput) Filter() GetEdgegatewayFilterPtrOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) *GetEdgegatewayFilter { return v.Filter }).(GetEdgegatewayFilterPtrOutput)
}

func (o GetEdgegatewayResultOutput) FipsModeEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.FipsModeEnabled }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) FwDefaultRuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) string { return v.FwDefaultRuleAction }).(pulumi.StringOutput)
}

func (o GetEdgegatewayResultOutput) FwDefaultRuleLoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.FwDefaultRuleLoggingEnabled }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) FwEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.FwEnabled }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) HaEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.HaEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEdgegatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEdgegatewayResultOutput) LbAccelerationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.LbAccelerationEnabled }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) LbEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.LbEnabled }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) LbLoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.LbLoggingEnabled }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) LbLoglevel() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) string { return v.LbLoglevel }).(pulumi.StringOutput)
}

func (o GetEdgegatewayResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetEdgegatewayResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o GetEdgegatewayResultOutput) UseDefaultRouteForDnsRelay() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) bool { return v.UseDefaultRouteForDnsRelay }).(pulumi.BoolOutput)
}

func (o GetEdgegatewayResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEdgegatewayResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEdgegatewayResultOutput{})
}
