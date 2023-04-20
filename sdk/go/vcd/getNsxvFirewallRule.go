// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxvFirewallRule(ctx *pulumi.Context, args *LookupNsxvFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupNsxvFirewallRuleResult, error) {
	var rv LookupNsxvFirewallRuleResult
	err := ctx.Invoke("vcd:index/getNsxvFirewallRule:getNsxvFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxvFirewallRule.
type LookupNsxvFirewallRuleArgs struct {
	EdgeGateway string  `pulumi:"edgeGateway"`
	Org         *string `pulumi:"org"`
	RuleId      string  `pulumi:"ruleId"`
	Vdc         *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxvFirewallRule.
type LookupNsxvFirewallRuleResult struct {
	Action       string                           `pulumi:"action"`
	Destinations []GetNsxvFirewallRuleDestination `pulumi:"destinations"`
	EdgeGateway  string                           `pulumi:"edgeGateway"`
	Enabled      bool                             `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id             string                       `pulumi:"id"`
	LoggingEnabled bool                         `pulumi:"loggingEnabled"`
	Name           string                       `pulumi:"name"`
	Org            *string                      `pulumi:"org"`
	RuleId         string                       `pulumi:"ruleId"`
	RuleTag        int                          `pulumi:"ruleTag"`
	RuleType       string                       `pulumi:"ruleType"`
	Services       []GetNsxvFirewallRuleService `pulumi:"services"`
	Sources        []GetNsxvFirewallRuleSource  `pulumi:"sources"`
	Vdc            *string                      `pulumi:"vdc"`
}

func LookupNsxvFirewallRuleOutput(ctx *pulumi.Context, args LookupNsxvFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNsxvFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxvFirewallRuleResult, error) {
			args := v.(LookupNsxvFirewallRuleArgs)
			r, err := LookupNsxvFirewallRule(ctx, &args, opts...)
			var s LookupNsxvFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxvFirewallRuleResultOutput)
}

// A collection of arguments for invoking getNsxvFirewallRule.
type LookupNsxvFirewallRuleOutputArgs struct {
	EdgeGateway pulumi.StringInput    `pulumi:"edgeGateway"`
	Org         pulumi.StringPtrInput `pulumi:"org"`
	RuleId      pulumi.StringInput    `pulumi:"ruleId"`
	Vdc         pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxvFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvFirewallRuleArgs)(nil)).Elem()
}

// A collection of values returned by getNsxvFirewallRule.
type LookupNsxvFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNsxvFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvFirewallRuleResult)(nil)).Elem()
}

func (o LookupNsxvFirewallRuleResultOutput) ToLookupNsxvFirewallRuleResultOutput() LookupNsxvFirewallRuleResultOutput {
	return o
}

func (o LookupNsxvFirewallRuleResultOutput) ToLookupNsxvFirewallRuleResultOutputWithContext(ctx context.Context) LookupNsxvFirewallRuleResultOutput {
	return o
}

func (o LookupNsxvFirewallRuleResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) string { return v.Action }).(pulumi.StringOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Destinations() GetNsxvFirewallRuleDestinationArrayOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) []GetNsxvFirewallRuleDestination { return v.Destinations }).(GetNsxvFirewallRuleDestinationArrayOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxvFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) LoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) bool { return v.LoggingEnabled }).(pulumi.BoolOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) string { return v.RuleId }).(pulumi.StringOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) RuleTag() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) int { return v.RuleTag }).(pulumi.IntOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Services() GetNsxvFirewallRuleServiceArrayOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) []GetNsxvFirewallRuleService { return v.Services }).(GetNsxvFirewallRuleServiceArrayOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Sources() GetNsxvFirewallRuleSourceArrayOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) []GetNsxvFirewallRuleSource { return v.Sources }).(GetNsxvFirewallRuleSourceArrayOutput)
}

func (o LookupNsxvFirewallRuleResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvFirewallRuleResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxvFirewallRuleResultOutput{})
}