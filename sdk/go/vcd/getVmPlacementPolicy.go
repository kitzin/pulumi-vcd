// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVmPlacementPolicy(ctx *pulumi.Context, args *LookupVmPlacementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupVmPlacementPolicyResult, error) {
	var rv LookupVmPlacementPolicyResult
	err := ctx.Invoke("vcd:index/getVmPlacementPolicy:getVmPlacementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmPlacementPolicy.
type LookupVmPlacementPolicyArgs struct {
	Name          string  `pulumi:"name"`
	ProviderVdcId *string `pulumi:"providerVdcId"`
	VdcId         *string `pulumi:"vdcId"`
}

// A collection of values returned by getVmPlacementPolicy.
type LookupVmPlacementPolicyResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                string   `pulumi:"id"`
	LogicalVmGroupIds []string `pulumi:"logicalVmGroupIds"`
	Name              string   `pulumi:"name"`
	ProviderVdcId     *string  `pulumi:"providerVdcId"`
	VdcId             *string  `pulumi:"vdcId"`
	VmGroupIds        []string `pulumi:"vmGroupIds"`
}

func LookupVmPlacementPolicyOutput(ctx *pulumi.Context, args LookupVmPlacementPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupVmPlacementPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVmPlacementPolicyResult, error) {
			args := v.(LookupVmPlacementPolicyArgs)
			r, err := LookupVmPlacementPolicy(ctx, &args, opts...)
			var s LookupVmPlacementPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVmPlacementPolicyResultOutput)
}

// A collection of arguments for invoking getVmPlacementPolicy.
type LookupVmPlacementPolicyOutputArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	ProviderVdcId pulumi.StringPtrInput `pulumi:"providerVdcId"`
	VdcId         pulumi.StringPtrInput `pulumi:"vdcId"`
}

func (LookupVmPlacementPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmPlacementPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getVmPlacementPolicy.
type LookupVmPlacementPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupVmPlacementPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmPlacementPolicyResult)(nil)).Elem()
}

func (o LookupVmPlacementPolicyResultOutput) ToLookupVmPlacementPolicyResultOutput() LookupVmPlacementPolicyResultOutput {
	return o
}

func (o LookupVmPlacementPolicyResultOutput) ToLookupVmPlacementPolicyResultOutputWithContext(ctx context.Context) LookupVmPlacementPolicyResultOutput {
	return o
}

func (o LookupVmPlacementPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVmPlacementPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVmPlacementPolicyResultOutput) LogicalVmGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) []string { return v.LogicalVmGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupVmPlacementPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVmPlacementPolicyResultOutput) ProviderVdcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) *string { return v.ProviderVdcId }).(pulumi.StringPtrOutput)
}

func (o LookupVmPlacementPolicyResultOutput) VdcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) *string { return v.VdcId }).(pulumi.StringPtrOutput)
}

func (o LookupVmPlacementPolicyResultOutput) VmGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVmPlacementPolicyResult) []string { return v.VmGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmPlacementPolicyResultOutput{})
}
