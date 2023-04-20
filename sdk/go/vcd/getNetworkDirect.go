// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkDirect(ctx *pulumi.Context, args *LookupNetworkDirectArgs, opts ...pulumi.InvokeOption) (*LookupNetworkDirectResult, error) {
	var rv LookupNetworkDirectResult
	err := ctx.Invoke("vcd:index/getNetworkDirect:getNetworkDirect", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkDirect.
type LookupNetworkDirectArgs struct {
	Filter *GetNetworkDirectFilter `pulumi:"filter"`
	Name   *string                 `pulumi:"name"`
	Org    *string                 `pulumi:"org"`
	Vdc    *string                 `pulumi:"vdc"`
}

// A collection of values returned by getNetworkDirect.
type LookupNetworkDirectResult struct {
	Description              string                  `pulumi:"description"`
	ExternalNetwork          string                  `pulumi:"externalNetwork"`
	ExternalNetworkDns1      string                  `pulumi:"externalNetworkDns1"`
	ExternalNetworkDns2      string                  `pulumi:"externalNetworkDns2"`
	ExternalNetworkDnsSuffix string                  `pulumi:"externalNetworkDnsSuffix"`
	ExternalNetworkGateway   string                  `pulumi:"externalNetworkGateway"`
	ExternalNetworkNetmask   string                  `pulumi:"externalNetworkNetmask"`
	Filter                   *GetNetworkDirectFilter `pulumi:"filter"`
	Href                     string                  `pulumi:"href"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Deprecated: Use metadata_entry instead
	Metadata        map[string]interface{}          `pulumi:"metadata"`
	MetadataEntries []GetNetworkDirectMetadataEntry `pulumi:"metadataEntries"`
	Name            *string                         `pulumi:"name"`
	Org             *string                         `pulumi:"org"`
	Shared          bool                            `pulumi:"shared"`
	Vdc             *string                         `pulumi:"vdc"`
}

func LookupNetworkDirectOutput(ctx *pulumi.Context, args LookupNetworkDirectOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkDirectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkDirectResult, error) {
			args := v.(LookupNetworkDirectArgs)
			r, err := LookupNetworkDirect(ctx, &args, opts...)
			var s LookupNetworkDirectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkDirectResultOutput)
}

// A collection of arguments for invoking getNetworkDirect.
type LookupNetworkDirectOutputArgs struct {
	Filter GetNetworkDirectFilterPtrInput `pulumi:"filter"`
	Name   pulumi.StringPtrInput          `pulumi:"name"`
	Org    pulumi.StringPtrInput          `pulumi:"org"`
	Vdc    pulumi.StringPtrInput          `pulumi:"vdc"`
}

func (LookupNetworkDirectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkDirectArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkDirect.
type LookupNetworkDirectResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkDirectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkDirectResult)(nil)).Elem()
}

func (o LookupNetworkDirectResultOutput) ToLookupNetworkDirectResultOutput() LookupNetworkDirectResultOutput {
	return o
}

func (o LookupNetworkDirectResultOutput) ToLookupNetworkDirectResultOutputWithContext(ctx context.Context) LookupNetworkDirectResultOutput {
	return o
}

func (o LookupNetworkDirectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) ExternalNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.ExternalNetwork }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) ExternalNetworkDns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.ExternalNetworkDns1 }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) ExternalNetworkDns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.ExternalNetworkDns2 }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) ExternalNetworkDnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.ExternalNetworkDnsSuffix }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) ExternalNetworkGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.ExternalNetworkGateway }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) ExternalNetworkNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.ExternalNetworkNetmask }).(pulumi.StringOutput)
}

func (o LookupNetworkDirectResultOutput) Filter() GetNetworkDirectFilterPtrOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) *GetNetworkDirectFilter { return v.Filter }).(GetNetworkDirectFilterPtrOutput)
}

func (o LookupNetworkDirectResultOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.Href }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkDirectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupNetworkDirectResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupNetworkDirectResultOutput) MetadataEntries() GetNetworkDirectMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) []GetNetworkDirectMetadataEntry { return v.MetadataEntries }).(GetNetworkDirectMetadataEntryArrayOutput)
}

func (o LookupNetworkDirectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkDirectResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkDirectResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) bool { return v.Shared }).(pulumi.BoolOutput)
}

func (o LookupNetworkDirectResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkDirectResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkDirectResultOutput{})
}