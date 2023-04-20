// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkIsolated(ctx *pulumi.Context, args *LookupNetworkIsolatedArgs, opts ...pulumi.InvokeOption) (*LookupNetworkIsolatedResult, error) {
	var rv LookupNetworkIsolatedResult
	err := ctx.Invoke("vcd:index/getNetworkIsolated:getNetworkIsolated", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkIsolated.
type LookupNetworkIsolatedArgs struct {
	Filter *GetNetworkIsolatedFilter `pulumi:"filter"`
	Name   *string                   `pulumi:"name"`
	Org    *string                   `pulumi:"org"`
	Vdc    *string                   `pulumi:"vdc"`
}

// A collection of values returned by getNetworkIsolated.
type LookupNetworkIsolatedResult struct {
	Description string                       `pulumi:"description"`
	DhcpPools   []GetNetworkIsolatedDhcpPool `pulumi:"dhcpPools"`
	Dns1        string                       `pulumi:"dns1"`
	Dns2        string                       `pulumi:"dns2"`
	DnsSuffix   string                       `pulumi:"dnsSuffix"`
	Filter      *GetNetworkIsolatedFilter    `pulumi:"filter"`
	Gateway     string                       `pulumi:"gateway"`
	Href        string                       `pulumi:"href"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Deprecated: Use metadata_entry instead
	Metadata        map[string]interface{}            `pulumi:"metadata"`
	MetadataEntries []GetNetworkIsolatedMetadataEntry `pulumi:"metadataEntries"`
	Name            *string                           `pulumi:"name"`
	Netmask         string                            `pulumi:"netmask"`
	Org             *string                           `pulumi:"org"`
	Shared          bool                              `pulumi:"shared"`
	StaticIpPools   []GetNetworkIsolatedStaticIpPool  `pulumi:"staticIpPools"`
	Vdc             *string                           `pulumi:"vdc"`
}

func LookupNetworkIsolatedOutput(ctx *pulumi.Context, args LookupNetworkIsolatedOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkIsolatedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkIsolatedResult, error) {
			args := v.(LookupNetworkIsolatedArgs)
			r, err := LookupNetworkIsolated(ctx, &args, opts...)
			var s LookupNetworkIsolatedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkIsolatedResultOutput)
}

// A collection of arguments for invoking getNetworkIsolated.
type LookupNetworkIsolatedOutputArgs struct {
	Filter GetNetworkIsolatedFilterPtrInput `pulumi:"filter"`
	Name   pulumi.StringPtrInput            `pulumi:"name"`
	Org    pulumi.StringPtrInput            `pulumi:"org"`
	Vdc    pulumi.StringPtrInput            `pulumi:"vdc"`
}

func (LookupNetworkIsolatedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkIsolatedArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkIsolated.
type LookupNetworkIsolatedResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkIsolatedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkIsolatedResult)(nil)).Elem()
}

func (o LookupNetworkIsolatedResultOutput) ToLookupNetworkIsolatedResultOutput() LookupNetworkIsolatedResultOutput {
	return o
}

func (o LookupNetworkIsolatedResultOutput) ToLookupNetworkIsolatedResultOutputWithContext(ctx context.Context) LookupNetworkIsolatedResultOutput {
	return o
}

func (o LookupNetworkIsolatedResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedResultOutput) DhcpPools() GetNetworkIsolatedDhcpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) []GetNetworkIsolatedDhcpPool { return v.DhcpPools }).(GetNetworkIsolatedDhcpPoolArrayOutput)
}

func (o LookupNetworkIsolatedResultOutput) Dns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Dns1 }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedResultOutput) Dns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Dns2 }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedResultOutput) DnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.DnsSuffix }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedResultOutput) Filter() GetNetworkIsolatedFilterPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) *GetNetworkIsolatedFilter { return v.Filter }).(GetNetworkIsolatedFilterPtrOutput)
}

func (o LookupNetworkIsolatedResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedResultOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Href }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkIsolatedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupNetworkIsolatedResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupNetworkIsolatedResultOutput) MetadataEntries() GetNetworkIsolatedMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) []GetNetworkIsolatedMetadataEntry { return v.MetadataEntries }).(GetNetworkIsolatedMetadataEntryArrayOutput)
}

func (o LookupNetworkIsolatedResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkIsolatedResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) string { return v.Netmask }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkIsolatedResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) bool { return v.Shared }).(pulumi.BoolOutput)
}

func (o LookupNetworkIsolatedResultOutput) StaticIpPools() GetNetworkIsolatedStaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) []GetNetworkIsolatedStaticIpPool { return v.StaticIpPools }).(GetNetworkIsolatedStaticIpPoolArrayOutput)
}

func (o LookupNetworkIsolatedResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkIsolatedResultOutput{})
}
