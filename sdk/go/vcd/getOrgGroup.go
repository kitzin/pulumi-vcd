// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrgGroup(ctx *pulumi.Context, args *LookupOrgGroupArgs, opts ...pulumi.InvokeOption) (*LookupOrgGroupResult, error) {
	var rv LookupOrgGroupResult
	err := ctx.Invoke("vcd:index/getOrgGroup:getOrgGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgGroup.
type LookupOrgGroupArgs struct {
	Name string  `pulumi:"name"`
	Org  *string `pulumi:"org"`
}

// A collection of values returned by getOrgGroup.
type LookupOrgGroupResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id           string   `pulumi:"id"`
	Name         string   `pulumi:"name"`
	Org          *string  `pulumi:"org"`
	ProviderType string   `pulumi:"providerType"`
	Role         string   `pulumi:"role"`
	UserNames    []string `pulumi:"userNames"`
}

func LookupOrgGroupOutput(ctx *pulumi.Context, args LookupOrgGroupOutputArgs, opts ...pulumi.InvokeOption) LookupOrgGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgGroupResult, error) {
			args := v.(LookupOrgGroupArgs)
			r, err := LookupOrgGroup(ctx, &args, opts...)
			var s LookupOrgGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgGroupResultOutput)
}

// A collection of arguments for invoking getOrgGroup.
type LookupOrgGroupOutputArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Org  pulumi.StringPtrInput `pulumi:"org"`
}

func (LookupOrgGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgGroupArgs)(nil)).Elem()
}

// A collection of values returned by getOrgGroup.
type LookupOrgGroupResultOutput struct{ *pulumi.OutputState }

func (LookupOrgGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgGroupResult)(nil)).Elem()
}

func (o LookupOrgGroupResultOutput) ToLookupOrgGroupResultOutput() LookupOrgGroupResultOutput {
	return o
}

func (o LookupOrgGroupResultOutput) ToLookupOrgGroupResultOutputWithContext(ctx context.Context) LookupOrgGroupResultOutput {
	return o
}

func (o LookupOrgGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrgGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrgGroupResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupOrgGroupResultOutput) ProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) string { return v.ProviderType }).(pulumi.StringOutput)
}

func (o LookupOrgGroupResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupOrgGroupResultOutput) UserNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgGroupResult) []string { return v.UserNames }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgGroupResultOutput{})
}