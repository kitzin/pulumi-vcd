// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkRouted(ctx *pulumi.Context, args *LookupNetworkRoutedArgs, opts ...pulumi.InvokeOption) (*LookupNetworkRoutedResult, error) {
	var rv LookupNetworkRoutedResult
	err := ctx.Invoke("vcd:index/getNetworkRouted:getNetworkRouted", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkRouted.
type LookupNetworkRoutedArgs struct {
	Filter *GetNetworkRoutedFilter `pulumi:"filter"`
	Name   *string                 `pulumi:"name"`
	Org    *string                 `pulumi:"org"`
	Vdc    *string                 `pulumi:"vdc"`
}

// A collection of values returned by getNetworkRouted.
type LookupNetworkRoutedResult struct {
	Description string                     `pulumi:"description"`
	DhcpPools   []GetNetworkRoutedDhcpPool `pulumi:"dhcpPools"`
	Dns1        string                     `pulumi:"dns1"`
	Dns2        string                     `pulumi:"dns2"`
	DnsSuffix   string                     `pulumi:"dnsSuffix"`
	EdgeGateway string                     `pulumi:"edgeGateway"`
	Filter      *GetNetworkRoutedFilter    `pulumi:"filter"`
	Gateway     string                     `pulumi:"gateway"`
	Href        string                     `pulumi:"href"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	InterfaceType string `pulumi:"interfaceType"`
	// Deprecated: Use metadata_entry instead
	Metadata        map[string]interface{}          `pulumi:"metadata"`
	MetadataEntries []GetNetworkRoutedMetadataEntry `pulumi:"metadataEntries"`
	Name            *string                         `pulumi:"name"`
	Netmask         string                          `pulumi:"netmask"`
	Org             *string                         `pulumi:"org"`
	Shared          bool                            `pulumi:"shared"`
	StaticIpPools   []GetNetworkRoutedStaticIpPool  `pulumi:"staticIpPools"`
	Vdc             *string                         `pulumi:"vdc"`
}

func LookupNetworkRoutedOutput(ctx *pulumi.Context, args LookupNetworkRoutedOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkRoutedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkRoutedResult, error) {
			args := v.(LookupNetworkRoutedArgs)
			r, err := LookupNetworkRouted(ctx, &args, opts...)
			var s LookupNetworkRoutedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkRoutedResultOutput)
}

// A collection of arguments for invoking getNetworkRouted.
type LookupNetworkRoutedOutputArgs struct {
	Filter GetNetworkRoutedFilterPtrInput `pulumi:"filter"`
	Name   pulumi.StringPtrInput          `pulumi:"name"`
	Org    pulumi.StringPtrInput          `pulumi:"org"`
	Vdc    pulumi.StringPtrInput          `pulumi:"vdc"`
}

func (LookupNetworkRoutedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkRoutedArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkRouted.
type LookupNetworkRoutedResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkRoutedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkRoutedResult)(nil)).Elem()
}

func (o LookupNetworkRoutedResultOutput) ToLookupNetworkRoutedResultOutput() LookupNetworkRoutedResultOutput {
	return o
}

func (o LookupNetworkRoutedResultOutput) ToLookupNetworkRoutedResultOutputWithContext(ctx context.Context) LookupNetworkRoutedResultOutput {
	return o
}

func (o LookupNetworkRoutedResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) DhcpPools() GetNetworkRoutedDhcpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) []GetNetworkRoutedDhcpPool { return v.DhcpPools }).(GetNetworkRoutedDhcpPoolArrayOutput)
}

func (o LookupNetworkRoutedResultOutput) Dns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Dns1 }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) Dns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Dns2 }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) DnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.DnsSuffix }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) Filter() GetNetworkRoutedFilterPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) *GetNetworkRoutedFilter { return v.Filter }).(GetNetworkRoutedFilterPtrOutput)
}

func (o LookupNetworkRoutedResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Href }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkRoutedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) InterfaceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.InterfaceType }).(pulumi.StringOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupNetworkRoutedResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupNetworkRoutedResultOutput) MetadataEntries() GetNetworkRoutedMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) []GetNetworkRoutedMetadataEntry { return v.MetadataEntries }).(GetNetworkRoutedMetadataEntryArrayOutput)
}

func (o LookupNetworkRoutedResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkRoutedResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) string { return v.Netmask }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkRoutedResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) bool { return v.Shared }).(pulumi.BoolOutput)
}

func (o LookupNetworkRoutedResultOutput) StaticIpPools() GetNetworkRoutedStaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) []GetNetworkRoutedStaticIpPool { return v.StaticIpPools }).(GetNetworkRoutedStaticIpPoolArrayOutput)
}

func (o LookupNetworkRoutedResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkRoutedResultOutput{})
}
