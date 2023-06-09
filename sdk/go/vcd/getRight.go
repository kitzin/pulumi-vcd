// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRight(ctx *pulumi.Context, args *GetRightArgs, opts ...pulumi.InvokeOption) (*GetRightResult, error) {
	var rv GetRightResult
	err := ctx.Invoke("vcd:index/getRight:getRight", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRight.
type GetRightArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getRight.
type GetRightResult struct {
	BundleKey   string `pulumi:"bundleKey"`
	CategoryId  string `pulumi:"categoryId"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id            string                 `pulumi:"id"`
	ImpliedRights []GetRightImpliedRight `pulumi:"impliedRights"`
	Name          string                 `pulumi:"name"`
	RightType     string                 `pulumi:"rightType"`
}

func GetRightOutput(ctx *pulumi.Context, args GetRightOutputArgs, opts ...pulumi.InvokeOption) GetRightResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRightResult, error) {
			args := v.(GetRightArgs)
			r, err := GetRight(ctx, &args, opts...)
			var s GetRightResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRightResultOutput)
}

// A collection of arguments for invoking getRight.
type GetRightOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetRightOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRightArgs)(nil)).Elem()
}

// A collection of values returned by getRight.
type GetRightResultOutput struct{ *pulumi.OutputState }

func (GetRightResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRightResult)(nil)).Elem()
}

func (o GetRightResultOutput) ToGetRightResultOutput() GetRightResultOutput {
	return o
}

func (o GetRightResultOutput) ToGetRightResultOutputWithContext(ctx context.Context) GetRightResultOutput {
	return o
}

func (o GetRightResultOutput) BundleKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetRightResult) string { return v.BundleKey }).(pulumi.StringOutput)
}

func (o GetRightResultOutput) CategoryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRightResult) string { return v.CategoryId }).(pulumi.StringOutput)
}

func (o GetRightResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetRightResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRightResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRightResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRightResultOutput) ImpliedRights() GetRightImpliedRightArrayOutput {
	return o.ApplyT(func(v GetRightResult) []GetRightImpliedRight { return v.ImpliedRights }).(GetRightImpliedRightArrayOutput)
}

func (o GetRightResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRightResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetRightResultOutput) RightType() pulumi.StringOutput {
	return o.ApplyT(func(v GetRightResult) string { return v.RightType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRightResultOutput{})
}
