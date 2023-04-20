// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkRoutedV2(ctx *pulumi.Context, args *LookupNetworkRoutedV2Args, opts ...pulumi.InvokeOption) (*LookupNetworkRoutedV2Result, error) {
	var rv LookupNetworkRoutedV2Result
	err := ctx.Invoke("vcd:index/getNetworkRoutedV2:getNetworkRoutedV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkRoutedV2.
type LookupNetworkRoutedV2Args struct {
	EdgeGatewayId *string                   `pulumi:"edgeGatewayId"`
	Filter        *GetNetworkRoutedV2Filter `pulumi:"filter"`
	Name          *string                   `pulumi:"name"`
	Org           *string                   `pulumi:"org"`
	// Deprecated: Deprecated in favor of `edge_gateway_id`. Routed networks will inherit VDC from parent Edge Gateway.
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNetworkRoutedV2.
type LookupNetworkRoutedV2Result struct {
	Description   string                    `pulumi:"description"`
	Dns1          string                    `pulumi:"dns1"`
	Dns2          string                    `pulumi:"dns2"`
	DnsSuffix     string                    `pulumi:"dnsSuffix"`
	EdgeGatewayId string                    `pulumi:"edgeGatewayId"`
	Filter        *GetNetworkRoutedV2Filter `pulumi:"filter"`
	Gateway       string                    `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	InterfaceType string `pulumi:"interfaceType"`
	// Deprecated: Use metadata_entry instead
	Metadata        map[string]interface{}            `pulumi:"metadata"`
	MetadataEntries []GetNetworkRoutedV2MetadataEntry `pulumi:"metadataEntries"`
	Name            *string                           `pulumi:"name"`
	Org             *string                           `pulumi:"org"`
	OwnerId         string                            `pulumi:"ownerId"`
	PrefixLength    int                               `pulumi:"prefixLength"`
	StaticIpPools   []GetNetworkRoutedV2StaticIpPool  `pulumi:"staticIpPools"`
	// Deprecated: Deprecated in favor of `edge_gateway_id`. Routed networks will inherit VDC from parent Edge Gateway.
	Vdc *string `pulumi:"vdc"`
}

func LookupNetworkRoutedV2Output(ctx *pulumi.Context, args LookupNetworkRoutedV2OutputArgs, opts ...pulumi.InvokeOption) LookupNetworkRoutedV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkRoutedV2Result, error) {
			args := v.(LookupNetworkRoutedV2Args)
			r, err := LookupNetworkRoutedV2(ctx, &args, opts...)
			var s LookupNetworkRoutedV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkRoutedV2ResultOutput)
}

// A collection of arguments for invoking getNetworkRoutedV2.
type LookupNetworkRoutedV2OutputArgs struct {
	EdgeGatewayId pulumi.StringPtrInput            `pulumi:"edgeGatewayId"`
	Filter        GetNetworkRoutedV2FilterPtrInput `pulumi:"filter"`
	Name          pulumi.StringPtrInput            `pulumi:"name"`
	Org           pulumi.StringPtrInput            `pulumi:"org"`
	// Deprecated: Deprecated in favor of `edge_gateway_id`. Routed networks will inherit VDC from parent Edge Gateway.
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNetworkRoutedV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkRoutedV2Args)(nil)).Elem()
}

// A collection of values returned by getNetworkRoutedV2.
type LookupNetworkRoutedV2ResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkRoutedV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkRoutedV2Result)(nil)).Elem()
}

func (o LookupNetworkRoutedV2ResultOutput) ToLookupNetworkRoutedV2ResultOutput() LookupNetworkRoutedV2ResultOutput {
	return o
}

func (o LookupNetworkRoutedV2ResultOutput) ToLookupNetworkRoutedV2ResultOutputWithContext(ctx context.Context) LookupNetworkRoutedV2ResultOutput {
	return o
}

func (o LookupNetworkRoutedV2ResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) Dns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.Dns1 }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) Dns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.Dns2 }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) DnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.DnsSuffix }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) Filter() GetNetworkRoutedV2FilterPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) *GetNetworkRoutedV2Filter { return v.Filter }).(GetNetworkRoutedV2FilterPtrOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkRoutedV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) InterfaceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.InterfaceType }).(pulumi.StringOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupNetworkRoutedV2ResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) MetadataEntries() GetNetworkRoutedV2MetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) []GetNetworkRoutedV2MetadataEntry { return v.MetadataEntries }).(GetNetworkRoutedV2MetadataEntryArrayOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) int { return v.PrefixLength }).(pulumi.IntOutput)
}

func (o LookupNetworkRoutedV2ResultOutput) StaticIpPools() GetNetworkRoutedV2StaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) []GetNetworkRoutedV2StaticIpPool { return v.StaticIpPools }).(GetNetworkRoutedV2StaticIpPoolArrayOutput)
}

// Deprecated: Deprecated in favor of `edge_gateway_id`. Routed networks will inherit VDC from parent Edge Gateway.
func (o LookupNetworkRoutedV2ResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkRoutedV2Result) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkRoutedV2ResultOutput{})
}