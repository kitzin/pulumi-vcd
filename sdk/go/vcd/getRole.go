// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRole(ctx *pulumi.Context, args *LookupRoleArgs, opts ...pulumi.InvokeOption) (*LookupRoleResult, error) {
	var rv LookupRoleResult
	err := ctx.Invoke("vcd:index/getRole:getRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type LookupRoleArgs struct {
	Name string  `pulumi:"name"`
	Org  *string `pulumi:"org"`
}

// A collection of values returned by getRole.
type LookupRoleResult struct {
	BundleKey   string `pulumi:"bundleKey"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id       string   `pulumi:"id"`
	Name     string   `pulumi:"name"`
	Org      *string  `pulumi:"org"`
	ReadOnly bool     `pulumi:"readOnly"`
	Rights   []string `pulumi:"rights"`
}

func LookupRoleOutput(ctx *pulumi.Context, args LookupRoleOutputArgs, opts ...pulumi.InvokeOption) LookupRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleResult, error) {
			args := v.(LookupRoleArgs)
			r, err := LookupRole(ctx, &args, opts...)
			var s LookupRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleResultOutput)
}

// A collection of arguments for invoking getRole.
type LookupRoleOutputArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Org  pulumi.StringPtrInput `pulumi:"org"`
}

func (LookupRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleArgs)(nil)).Elem()
}

// A collection of values returned by getRole.
type LookupRoleResultOutput struct{ *pulumi.OutputState }

func (LookupRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleResult)(nil)).Elem()
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutput() LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutputWithContext(ctx context.Context) LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) BundleKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.BundleKey }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupRoleResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRoleResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func (o LookupRoleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleResultOutput{})
}
