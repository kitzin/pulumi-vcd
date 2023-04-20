// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxvSnat(ctx *pulumi.Context, args *LookupNsxvSnatArgs, opts ...pulumi.InvokeOption) (*LookupNsxvSnatResult, error) {
	var rv LookupNsxvSnatResult
	err := ctx.Invoke("vcd:index/getNsxvSnat:getNsxvSnat", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxvSnat.
type LookupNsxvSnatArgs struct {
	EdgeGateway string  `pulumi:"edgeGateway"`
	Org         *string `pulumi:"org"`
	RuleId      string  `pulumi:"ruleId"`
	Vdc         *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxvSnat.
type LookupNsxvSnatResult struct {
	Description string `pulumi:"description"`
	EdgeGateway string `pulumi:"edgeGateway"`
	Enabled     bool   `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                string  `pulumi:"id"`
	LoggingEnabled    bool    `pulumi:"loggingEnabled"`
	NetworkName       string  `pulumi:"networkName"`
	NetworkType       string  `pulumi:"networkType"`
	Org               *string `pulumi:"org"`
	OriginalAddress   string  `pulumi:"originalAddress"`
	RuleId            string  `pulumi:"ruleId"`
	RuleTag           int     `pulumi:"ruleTag"`
	RuleType          string  `pulumi:"ruleType"`
	TranslatedAddress string  `pulumi:"translatedAddress"`
	Vdc               *string `pulumi:"vdc"`
}

func LookupNsxvSnatOutput(ctx *pulumi.Context, args LookupNsxvSnatOutputArgs, opts ...pulumi.InvokeOption) LookupNsxvSnatResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxvSnatResult, error) {
			args := v.(LookupNsxvSnatArgs)
			r, err := LookupNsxvSnat(ctx, &args, opts...)
			var s LookupNsxvSnatResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxvSnatResultOutput)
}

// A collection of arguments for invoking getNsxvSnat.
type LookupNsxvSnatOutputArgs struct {
	EdgeGateway pulumi.StringInput    `pulumi:"edgeGateway"`
	Org         pulumi.StringPtrInput `pulumi:"org"`
	RuleId      pulumi.StringInput    `pulumi:"ruleId"`
	Vdc         pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxvSnatOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvSnatArgs)(nil)).Elem()
}

// A collection of values returned by getNsxvSnat.
type LookupNsxvSnatResultOutput struct{ *pulumi.OutputState }

func (LookupNsxvSnatResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvSnatResult)(nil)).Elem()
}

func (o LookupNsxvSnatResultOutput) ToLookupNsxvSnatResultOutput() LookupNsxvSnatResultOutput {
	return o
}

func (o LookupNsxvSnatResultOutput) ToLookupNsxvSnatResultOutputWithContext(ctx context.Context) LookupNsxvSnatResultOutput {
	return o
}

func (o LookupNsxvSnatResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxvSnatResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) LoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) bool { return v.LoggingEnabled }).(pulumi.BoolOutput)
}

func (o LookupNsxvSnatResultOutput) NetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.NetworkName }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.NetworkType }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxvSnatResultOutput) OriginalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.OriginalAddress }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.RuleId }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) RuleTag() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) int { return v.RuleTag }).(pulumi.IntOutput)
}

func (o LookupNsxvSnatResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) TranslatedAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) string { return v.TranslatedAddress }).(pulumi.StringOutput)
}

func (o LookupNsxvSnatResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvSnatResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxvSnatResultOutput{})
}