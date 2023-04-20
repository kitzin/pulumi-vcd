// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtFirewall(ctx *pulumi.Context, args *LookupNsxtFirewallArgs, opts ...pulumi.InvokeOption) (*LookupNsxtFirewallResult, error) {
	var rv LookupNsxtFirewallResult
	err := ctx.Invoke("vcd:index/getNsxtFirewall:getNsxtFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtFirewall.
type LookupNsxtFirewallArgs struct {
	EdgeGatewayId string  `pulumi:"edgeGatewayId"`
	Org           *string `pulumi:"org"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtFirewall.
type LookupNsxtFirewallResult struct {
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id    string                `pulumi:"id"`
	Org   *string               `pulumi:"org"`
	Rules []GetNsxtFirewallRule `pulumi:"rules"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtFirewallOutput(ctx *pulumi.Context, args LookupNsxtFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtFirewallResult, error) {
			args := v.(LookupNsxtFirewallArgs)
			r, err := LookupNsxtFirewall(ctx, &args, opts...)
			var s LookupNsxtFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtFirewallResultOutput)
}

// A collection of arguments for invoking getNsxtFirewall.
type LookupNsxtFirewallOutputArgs struct {
	EdgeGatewayId pulumi.StringInput    `pulumi:"edgeGatewayId"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtFirewallArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtFirewall.
type LookupNsxtFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtFirewallResult)(nil)).Elem()
}

func (o LookupNsxtFirewallResultOutput) ToLookupNsxtFirewallResultOutput() LookupNsxtFirewallResultOutput {
	return o
}

func (o LookupNsxtFirewallResultOutput) ToLookupNsxtFirewallResultOutputWithContext(ctx context.Context) LookupNsxtFirewallResultOutput {
	return o
}

func (o LookupNsxtFirewallResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtFirewallResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtFirewallResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtFirewallResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtFirewallResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtFirewallResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtFirewallResultOutput) Rules() GetNsxtFirewallRuleArrayOutput {
	return o.ApplyT(func(v LookupNsxtFirewallResult) []GetNsxtFirewallRule { return v.Rules }).(GetNsxtFirewallRuleArrayOutput)
}

// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
func (o LookupNsxtFirewallResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtFirewallResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtFirewallResultOutput{})
}