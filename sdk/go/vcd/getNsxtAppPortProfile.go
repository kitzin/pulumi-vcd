// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtAppPortProfile(ctx *pulumi.Context, args *LookupNsxtAppPortProfileArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAppPortProfileResult, error) {
	var rv LookupNsxtAppPortProfileResult
	err := ctx.Invoke("vcd:index/getNsxtAppPortProfile:getNsxtAppPortProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAppPortProfile.
type LookupNsxtAppPortProfileArgs struct {
	ContextId *string `pulumi:"contextId"`
	Name      string  `pulumi:"name"`
	// Deprecated: Deprecated in favor of 'context_id'
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	Org           *string `pulumi:"org"`
	Scope         string  `pulumi:"scope"`
	// Deprecated: Deprecated in favor of 'context_id'
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtAppPortProfile.
type LookupNsxtAppPortProfileResult struct {
	AppPorts    []GetNsxtAppPortProfileAppPort `pulumi:"appPorts"`
	ContextId   string                         `pulumi:"contextId"`
	Description string                         `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Deprecated: Deprecated in favor of 'context_id'
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	Org           *string `pulumi:"org"`
	Scope         string  `pulumi:"scope"`
	// Deprecated: Deprecated in favor of 'context_id'
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtAppPortProfileOutput(ctx *pulumi.Context, args LookupNsxtAppPortProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAppPortProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtAppPortProfileResult, error) {
			args := v.(LookupNsxtAppPortProfileArgs)
			r, err := LookupNsxtAppPortProfile(ctx, &args, opts...)
			var s LookupNsxtAppPortProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtAppPortProfileResultOutput)
}

// A collection of arguments for invoking getNsxtAppPortProfile.
type LookupNsxtAppPortProfileOutputArgs struct {
	ContextId pulumi.StringPtrInput `pulumi:"contextId"`
	Name      pulumi.StringInput    `pulumi:"name"`
	// Deprecated: Deprecated in favor of 'context_id'
	NsxtManagerId pulumi.StringPtrInput `pulumi:"nsxtManagerId"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
	Scope         pulumi.StringInput    `pulumi:"scope"`
	// Deprecated: Deprecated in favor of 'context_id'
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtAppPortProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAppPortProfileArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAppPortProfile.
type LookupNsxtAppPortProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAppPortProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAppPortProfileResult)(nil)).Elem()
}

func (o LookupNsxtAppPortProfileResultOutput) ToLookupNsxtAppPortProfileResultOutput() LookupNsxtAppPortProfileResultOutput {
	return o
}

func (o LookupNsxtAppPortProfileResultOutput) ToLookupNsxtAppPortProfileResultOutputWithContext(ctx context.Context) LookupNsxtAppPortProfileResultOutput {
	return o
}

func (o LookupNsxtAppPortProfileResultOutput) AppPorts() GetNsxtAppPortProfileAppPortArrayOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) []GetNsxtAppPortProfileAppPort { return v.AppPorts }).(GetNsxtAppPortProfileAppPortArrayOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) ContextId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.ContextId }).(pulumi.StringOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAppPortProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favor of 'context_id'
func (o LookupNsxtAppPortProfileResultOutput) NsxtManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) *string { return v.NsxtManagerId }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favor of 'context_id'
func (o LookupNsxtAppPortProfileResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAppPortProfileResultOutput{})
}
