// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPortgroup(ctx *pulumi.Context, args *GetPortgroupArgs, opts ...pulumi.InvokeOption) (*GetPortgroupResult, error) {
	var rv GetPortgroupResult
	err := ctx.Invoke("vcd:index/getPortgroup:getPortgroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPortgroup.
type GetPortgroupArgs struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

// A collection of values returned by getPortgroup.
type GetPortgroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

func GetPortgroupOutput(ctx *pulumi.Context, args GetPortgroupOutputArgs, opts ...pulumi.InvokeOption) GetPortgroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPortgroupResult, error) {
			args := v.(GetPortgroupArgs)
			r, err := GetPortgroup(ctx, &args, opts...)
			var s GetPortgroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPortgroupResultOutput)
}

// A collection of arguments for invoking getPortgroup.
type GetPortgroupOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetPortgroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPortgroupArgs)(nil)).Elem()
}

// A collection of values returned by getPortgroup.
type GetPortgroupResultOutput struct{ *pulumi.OutputState }

func (GetPortgroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPortgroupResult)(nil)).Elem()
}

func (o GetPortgroupResultOutput) ToGetPortgroupResultOutput() GetPortgroupResultOutput {
	return o
}

func (o GetPortgroupResultOutput) ToGetPortgroupResultOutputWithContext(ctx context.Context) GetPortgroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPortgroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortgroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPortgroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortgroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetPortgroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortgroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPortgroupResultOutput{})
}