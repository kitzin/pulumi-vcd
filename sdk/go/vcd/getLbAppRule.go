// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLbAppRule(ctx *pulumi.Context, args *LookupLbAppRuleArgs, opts ...pulumi.InvokeOption) (*LookupLbAppRuleResult, error) {
	var rv LookupLbAppRuleResult
	err := ctx.Invoke("vcd:index/getLbAppRule:getLbAppRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbAppRule.
type LookupLbAppRuleArgs struct {
	EdgeGateway string  `pulumi:"edgeGateway"`
	Name        string  `pulumi:"name"`
	Org         *string `pulumi:"org"`
	Vdc         *string `pulumi:"vdc"`
}

// A collection of values returned by getLbAppRule.
type LookupLbAppRuleResult struct {
	EdgeGateway string `pulumi:"edgeGateway"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Org    *string `pulumi:"org"`
	Script string  `pulumi:"script"`
	Vdc    *string `pulumi:"vdc"`
}

func LookupLbAppRuleOutput(ctx *pulumi.Context, args LookupLbAppRuleOutputArgs, opts ...pulumi.InvokeOption) LookupLbAppRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbAppRuleResult, error) {
			args := v.(LookupLbAppRuleArgs)
			r, err := LookupLbAppRule(ctx, &args, opts...)
			var s LookupLbAppRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLbAppRuleResultOutput)
}

// A collection of arguments for invoking getLbAppRule.
type LookupLbAppRuleOutputArgs struct {
	EdgeGateway pulumi.StringInput    `pulumi:"edgeGateway"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Org         pulumi.StringPtrInput `pulumi:"org"`
	Vdc         pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupLbAppRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbAppRuleArgs)(nil)).Elem()
}

// A collection of values returned by getLbAppRule.
type LookupLbAppRuleResultOutput struct{ *pulumi.OutputState }

func (LookupLbAppRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbAppRuleResult)(nil)).Elem()
}

func (o LookupLbAppRuleResultOutput) ToLookupLbAppRuleResultOutput() LookupLbAppRuleResultOutput {
	return o
}

func (o LookupLbAppRuleResultOutput) ToLookupLbAppRuleResultOutputWithContext(ctx context.Context) LookupLbAppRuleResultOutput {
	return o
}

func (o LookupLbAppRuleResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppRuleResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbAppRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbAppRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLbAppRuleResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbAppRuleResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupLbAppRuleResultOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppRuleResult) string { return v.Script }).(pulumi.StringOutput)
}

func (o LookupLbAppRuleResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbAppRuleResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbAppRuleResultOutput{})
}