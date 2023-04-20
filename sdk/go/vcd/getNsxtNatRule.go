// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtNatRule(ctx *pulumi.Context, args *LookupNsxtNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupNsxtNatRuleResult, error) {
	var rv LookupNsxtNatRuleResult
	err := ctx.Invoke("vcd:index/getNsxtNatRule:getNsxtNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtNatRule.
type LookupNsxtNatRuleArgs struct {
	EdgeGatewayId string  `pulumi:"edgeGatewayId"`
	Name          string  `pulumi:"name"`
	Org           *string `pulumi:"org"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtNatRule.
type LookupNsxtNatRuleResult struct {
	AppPortProfileId string `pulumi:"appPortProfileId"`
	Description      string `pulumi:"description"`
	DnatExternalPort string `pulumi:"dnatExternalPort"`
	EdgeGatewayId    string `pulumi:"edgeGatewayId"`
	Enabled          bool   `pulumi:"enabled"`
	ExternalAddress  string `pulumi:"externalAddress"`
	FirewallMatch    string `pulumi:"firewallMatch"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string  `pulumi:"id"`
	InternalAddress        string  `pulumi:"internalAddress"`
	Logging                bool    `pulumi:"logging"`
	Name                   string  `pulumi:"name"`
	Org                    *string `pulumi:"org"`
	Priority               int     `pulumi:"priority"`
	RuleType               string  `pulumi:"ruleType"`
	SnatDestinationAddress string  `pulumi:"snatDestinationAddress"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtNatRuleOutput(ctx *pulumi.Context, args LookupNsxtNatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtNatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtNatRuleResult, error) {
			args := v.(LookupNsxtNatRuleArgs)
			r, err := LookupNsxtNatRule(ctx, &args, opts...)
			var s LookupNsxtNatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtNatRuleResultOutput)
}

// A collection of arguments for invoking getNsxtNatRule.
type LookupNsxtNatRuleOutputArgs struct {
	EdgeGatewayId pulumi.StringInput    `pulumi:"edgeGatewayId"`
	Name          pulumi.StringInput    `pulumi:"name"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtNatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtNatRuleArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtNatRule.
type LookupNsxtNatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtNatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtNatRuleResult)(nil)).Elem()
}

func (o LookupNsxtNatRuleResultOutput) ToLookupNsxtNatRuleResultOutput() LookupNsxtNatRuleResultOutput {
	return o
}

func (o LookupNsxtNatRuleResultOutput) ToLookupNsxtNatRuleResultOutputWithContext(ctx context.Context) LookupNsxtNatRuleResultOutput {
	return o
}

func (o LookupNsxtNatRuleResultOutput) AppPortProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.AppPortProfileId }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) DnatExternalPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.DnatExternalPort }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupNsxtNatRuleResultOutput) ExternalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.ExternalAddress }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) FirewallMatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.FirewallMatch }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtNatRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) InternalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.InternalAddress }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) Logging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) bool { return v.Logging }).(pulumi.BoolOutput)
}

func (o LookupNsxtNatRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtNatRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

func (o LookupNsxtNatRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o LookupNsxtNatRuleResultOutput) SnatDestinationAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) string { return v.SnatDestinationAddress }).(pulumi.StringOutput)
}

// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
func (o LookupNsxtNatRuleResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNatRuleResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtNatRuleResultOutput{})
}
