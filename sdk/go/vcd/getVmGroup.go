// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVmGroup(ctx *pulumi.Context, args *GetVmGroupArgs, opts ...pulumi.InvokeOption) (*GetVmGroupResult, error) {
	var rv GetVmGroupResult
	err := ctx.Invoke("vcd:index/getVmGroup:getVmGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmGroup.
type GetVmGroupArgs struct {
	Name          string `pulumi:"name"`
	ProviderVdcId string `pulumi:"providerVdcId"`
}

// A collection of values returned by getVmGroup.
type GetVmGroupResult struct {
	ClusterMoref string `pulumi:"clusterMoref"`
	ClusterName  string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	NamedVmGroupId string `pulumi:"namedVmGroupId"`
	ProviderVdcId  string `pulumi:"providerVdcId"`
	VcenterId      string `pulumi:"vcenterId"`
}

func GetVmGroupOutput(ctx *pulumi.Context, args GetVmGroupOutputArgs, opts ...pulumi.InvokeOption) GetVmGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVmGroupResult, error) {
			args := v.(GetVmGroupArgs)
			r, err := GetVmGroup(ctx, &args, opts...)
			var s GetVmGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVmGroupResultOutput)
}

// A collection of arguments for invoking getVmGroup.
type GetVmGroupOutputArgs struct {
	Name          pulumi.StringInput `pulumi:"name"`
	ProviderVdcId pulumi.StringInput `pulumi:"providerVdcId"`
}

func (GetVmGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVmGroupArgs)(nil)).Elem()
}

// A collection of values returned by getVmGroup.
type GetVmGroupResultOutput struct{ *pulumi.OutputState }

func (GetVmGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVmGroupResult)(nil)).Elem()
}

func (o GetVmGroupResultOutput) ToGetVmGroupResultOutput() GetVmGroupResultOutput {
	return o
}

func (o GetVmGroupResultOutput) ToGetVmGroupResultOutputWithContext(ctx context.Context) GetVmGroupResultOutput {
	return o
}

func (o GetVmGroupResultOutput) ClusterMoref() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.ClusterMoref }).(pulumi.StringOutput)
}

func (o GetVmGroupResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVmGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVmGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetVmGroupResultOutput) NamedVmGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.NamedVmGroupId }).(pulumi.StringOutput)
}

func (o GetVmGroupResultOutput) ProviderVdcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.ProviderVdcId }).(pulumi.StringOutput)
}

func (o GetVmGroupResultOutput) VcenterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmGroupResult) string { return v.VcenterId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVmGroupResultOutput{})
}
