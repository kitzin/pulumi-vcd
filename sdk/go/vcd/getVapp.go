// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVapp(ctx *pulumi.Context, args *LookupVappArgs, opts ...pulumi.InvokeOption) (*LookupVappResult, error) {
	var rv LookupVappResult
	err := ctx.Invoke("vcd:index/getVapp:getVapp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVapp.
type LookupVappArgs struct {
	Name string  `pulumi:"name"`
	Org  *string `pulumi:"org"`
	Vdc  *string `pulumi:"vdc"`
}

// A collection of values returned by getVapp.
type LookupVappResult struct {
	Description     string                 `pulumi:"description"`
	GuestProperties map[string]interface{} `pulumi:"guestProperties"`
	Href            string                 `pulumi:"href"`
	// The provider-assigned unique ID for this managed resource.
	Id     string         `pulumi:"id"`
	Leases []GetVappLease `pulumi:"leases"`
	// Deprecated: Use metadata_entry instead
	Metadata        map[string]interface{} `pulumi:"metadata"`
	MetadataEntries []GetVappMetadataEntry `pulumi:"metadataEntries"`
	Name            string                 `pulumi:"name"`
	Org             *string                `pulumi:"org"`
	Status          int                    `pulumi:"status"`
	StatusText      string                 `pulumi:"statusText"`
	Vdc             *string                `pulumi:"vdc"`
}

func LookupVappOutput(ctx *pulumi.Context, args LookupVappOutputArgs, opts ...pulumi.InvokeOption) LookupVappResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVappResult, error) {
			args := v.(LookupVappArgs)
			r, err := LookupVapp(ctx, &args, opts...)
			var s LookupVappResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVappResultOutput)
}

// A collection of arguments for invoking getVapp.
type LookupVappOutputArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Org  pulumi.StringPtrInput `pulumi:"org"`
	Vdc  pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupVappOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVappArgs)(nil)).Elem()
}

// A collection of values returned by getVapp.
type LookupVappResultOutput struct{ *pulumi.OutputState }

func (LookupVappResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVappResult)(nil)).Elem()
}

func (o LookupVappResultOutput) ToLookupVappResultOutput() LookupVappResultOutput {
	return o
}

func (o LookupVappResultOutput) ToLookupVappResultOutputWithContext(ctx context.Context) LookupVappResultOutput {
	return o
}

func (o LookupVappResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupVappResultOutput) GuestProperties() pulumi.MapOutput {
	return o.ApplyT(func(v LookupVappResult) map[string]interface{} { return v.GuestProperties }).(pulumi.MapOutput)
}

func (o LookupVappResultOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappResult) string { return v.Href }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVappResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVappResultOutput) Leases() GetVappLeaseArrayOutput {
	return o.ApplyT(func(v LookupVappResult) []GetVappLease { return v.Leases }).(GetVappLeaseArrayOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupVappResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupVappResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupVappResultOutput) MetadataEntries() GetVappMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupVappResult) []GetVappMetadataEntry { return v.MetadataEntries }).(GetVappMetadataEntryArrayOutput)
}

func (o LookupVappResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVappResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVappResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupVappResultOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappResult) int { return v.Status }).(pulumi.IntOutput)
}

func (o LookupVappResultOutput) StatusText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappResult) string { return v.StatusText }).(pulumi.StringOutput)
}

func (o LookupVappResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVappResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVappResultOutput{})
}