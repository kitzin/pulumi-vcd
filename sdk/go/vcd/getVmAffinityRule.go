// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVmAffinityRule(ctx *pulumi.Context, args *LookupVmAffinityRuleArgs, opts ...pulumi.InvokeOption) (*LookupVmAffinityRuleResult, error) {
	var rv LookupVmAffinityRuleResult
	err := ctx.Invoke("vcd:index/getVmAffinityRule:getVmAffinityRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmAffinityRule.
type LookupVmAffinityRuleArgs struct {
	Name   *string `pulumi:"name"`
	Org    *string `pulumi:"org"`
	RuleId *string `pulumi:"ruleId"`
	Vdc    *string `pulumi:"vdc"`
}

// A collection of values returned by getVmAffinityRule.
type LookupVmAffinityRuleResult struct {
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id       string   `pulumi:"id"`
	Name     *string  `pulumi:"name"`
	Org      *string  `pulumi:"org"`
	Polarity string   `pulumi:"polarity"`
	Required bool     `pulumi:"required"`
	RuleId   *string  `pulumi:"ruleId"`
	Vdc      *string  `pulumi:"vdc"`
	VmIds    []string `pulumi:"vmIds"`
}

func LookupVmAffinityRuleOutput(ctx *pulumi.Context, args LookupVmAffinityRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVmAffinityRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVmAffinityRuleResult, error) {
			args := v.(LookupVmAffinityRuleArgs)
			r, err := LookupVmAffinityRule(ctx, &args, opts...)
			var s LookupVmAffinityRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVmAffinityRuleResultOutput)
}

// A collection of arguments for invoking getVmAffinityRule.
type LookupVmAffinityRuleOutputArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Org    pulumi.StringPtrInput `pulumi:"org"`
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
	Vdc    pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupVmAffinityRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmAffinityRuleArgs)(nil)).Elem()
}

// A collection of values returned by getVmAffinityRule.
type LookupVmAffinityRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVmAffinityRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmAffinityRuleResult)(nil)).Elem()
}

func (o LookupVmAffinityRuleResultOutput) ToLookupVmAffinityRuleResultOutput() LookupVmAffinityRuleResultOutput {
	return o
}

func (o LookupVmAffinityRuleResultOutput) ToLookupVmAffinityRuleResultOutputWithContext(ctx context.Context) LookupVmAffinityRuleResultOutput {
	return o
}

func (o LookupVmAffinityRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVmAffinityRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVmAffinityRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVmAffinityRuleResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupVmAffinityRuleResultOutput) Polarity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) string { return v.Polarity }).(pulumi.StringOutput)
}

func (o LookupVmAffinityRuleResultOutput) Required() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) bool { return v.Required }).(pulumi.BoolOutput)
}

func (o LookupVmAffinityRuleResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func (o LookupVmAffinityRuleResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func (o LookupVmAffinityRuleResultOutput) VmIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) []string { return v.VmIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmAffinityRuleResultOutput{})
}
