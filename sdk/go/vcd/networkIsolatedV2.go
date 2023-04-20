// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkIsolatedV2 struct {
	pulumi.CustomResourceState

	// Network description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// DNS server 1
	Dns1 pulumi.StringPtrOutput `pulumi:"dns1"`
	// DNS server 1
	Dns2 pulumi.StringPtrOutput `pulumi:"dns2"`
	// DNS suffix
	DnsSuffix pulumi.StringPtrOutput `pulumi:"dnsSuffix"`
	// Gateway IP address
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// NSX-V only - share this network with other VDCs in this organization. Default - false
	IsShared pulumi.BoolOutput `pulumi:"isShared"`
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Metadata entries for the given Network
	MetadataEntries NetworkIsolatedV2MetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// Network name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// ID of VDC or VDC Group
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Network prefix
	PrefixLength pulumi.IntOutput `pulumi:"prefixLength"`
	// IP ranges used for static pool allocation in the network
	StaticIpPools NetworkIsolatedV2StaticIpPoolArrayOutput `pulumi:"staticIpPools"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringOutput `pulumi:"vdc"`
}

// NewNetworkIsolatedV2 registers a new resource with the given unique name, arguments, and options.
func NewNetworkIsolatedV2(ctx *pulumi.Context,
	name string, args *NetworkIsolatedV2Args, opts ...pulumi.ResourceOption) (*NetworkIsolatedV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.PrefixLength == nil {
		return nil, errors.New("invalid value for required argument 'PrefixLength'")
	}
	var resource NetworkIsolatedV2
	err := ctx.RegisterResource("vcd:index/networkIsolatedV2:NetworkIsolatedV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkIsolatedV2 gets an existing NetworkIsolatedV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkIsolatedV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkIsolatedV2State, opts ...pulumi.ResourceOption) (*NetworkIsolatedV2, error) {
	var resource NetworkIsolatedV2
	err := ctx.ReadResource("vcd:index/networkIsolatedV2:NetworkIsolatedV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkIsolatedV2 resources.
type networkIsolatedV2State struct {
	// Network description
	Description *string `pulumi:"description"`
	// DNS server 1
	Dns1 *string `pulumi:"dns1"`
	// DNS server 1
	Dns2 *string `pulumi:"dns2"`
	// DNS suffix
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Gateway IP address
	Gateway *string `pulumi:"gateway"`
	// NSX-V only - share this network with other VDCs in this organization. Default - false
	IsShared *bool `pulumi:"isShared"`
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata entries for the given Network
	MetadataEntries []NetworkIsolatedV2MetadataEntry `pulumi:"metadataEntries"`
	// Network name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// ID of VDC or VDC Group
	OwnerId *string `pulumi:"ownerId"`
	// Network prefix
	PrefixLength *int `pulumi:"prefixLength"`
	// IP ranges used for static pool allocation in the network
	StaticIpPools []NetworkIsolatedV2StaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

type NetworkIsolatedV2State struct {
	// Network description
	Description pulumi.StringPtrInput
	// DNS server 1
	Dns1 pulumi.StringPtrInput
	// DNS server 1
	Dns2 pulumi.StringPtrInput
	// DNS suffix
	DnsSuffix pulumi.StringPtrInput
	// Gateway IP address
	Gateway pulumi.StringPtrInput
	// NSX-V only - share this network with other VDCs in this organization. Default - false
	IsShared pulumi.BoolPtrInput
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata pulumi.MapInput
	// Metadata entries for the given Network
	MetadataEntries NetworkIsolatedV2MetadataEntryArrayInput
	// Network name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// ID of VDC or VDC Group
	OwnerId pulumi.StringPtrInput
	// Network prefix
	PrefixLength pulumi.IntPtrInput
	// IP ranges used for static pool allocation in the network
	StaticIpPools NetworkIsolatedV2StaticIpPoolArrayInput
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput
}

func (NetworkIsolatedV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*networkIsolatedV2State)(nil)).Elem()
}

type networkIsolatedV2Args struct {
	// Network description
	Description *string `pulumi:"description"`
	// DNS server 1
	Dns1 *string `pulumi:"dns1"`
	// DNS server 1
	Dns2 *string `pulumi:"dns2"`
	// DNS suffix
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Gateway IP address
	Gateway string `pulumi:"gateway"`
	// NSX-V only - share this network with other VDCs in this organization. Default - false
	IsShared *bool `pulumi:"isShared"`
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata entries for the given Network
	MetadataEntries []NetworkIsolatedV2MetadataEntry `pulumi:"metadataEntries"`
	// Network name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// ID of VDC or VDC Group
	OwnerId *string `pulumi:"ownerId"`
	// Network prefix
	PrefixLength int `pulumi:"prefixLength"`
	// IP ranges used for static pool allocation in the network
	StaticIpPools []NetworkIsolatedV2StaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NetworkIsolatedV2 resource.
type NetworkIsolatedV2Args struct {
	// Network description
	Description pulumi.StringPtrInput
	// DNS server 1
	Dns1 pulumi.StringPtrInput
	// DNS server 1
	Dns2 pulumi.StringPtrInput
	// DNS suffix
	DnsSuffix pulumi.StringPtrInput
	// Gateway IP address
	Gateway pulumi.StringInput
	// NSX-V only - share this network with other VDCs in this organization. Default - false
	IsShared pulumi.BoolPtrInput
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata pulumi.MapInput
	// Metadata entries for the given Network
	MetadataEntries NetworkIsolatedV2MetadataEntryArrayInput
	// Network name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// ID of VDC or VDC Group
	OwnerId pulumi.StringPtrInput
	// Network prefix
	PrefixLength pulumi.IntInput
	// IP ranges used for static pool allocation in the network
	StaticIpPools NetworkIsolatedV2StaticIpPoolArrayInput
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput
}

func (NetworkIsolatedV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*networkIsolatedV2Args)(nil)).Elem()
}

type NetworkIsolatedV2Input interface {
	pulumi.Input

	ToNetworkIsolatedV2Output() NetworkIsolatedV2Output
	ToNetworkIsolatedV2OutputWithContext(ctx context.Context) NetworkIsolatedV2Output
}

func (*NetworkIsolatedV2) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkIsolatedV2)(nil)).Elem()
}

func (i *NetworkIsolatedV2) ToNetworkIsolatedV2Output() NetworkIsolatedV2Output {
	return i.ToNetworkIsolatedV2OutputWithContext(context.Background())
}

func (i *NetworkIsolatedV2) ToNetworkIsolatedV2OutputWithContext(ctx context.Context) NetworkIsolatedV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkIsolatedV2Output)
}

// NetworkIsolatedV2ArrayInput is an input type that accepts NetworkIsolatedV2Array and NetworkIsolatedV2ArrayOutput values.
// You can construct a concrete instance of `NetworkIsolatedV2ArrayInput` via:
//
//	NetworkIsolatedV2Array{ NetworkIsolatedV2Args{...} }
type NetworkIsolatedV2ArrayInput interface {
	pulumi.Input

	ToNetworkIsolatedV2ArrayOutput() NetworkIsolatedV2ArrayOutput
	ToNetworkIsolatedV2ArrayOutputWithContext(context.Context) NetworkIsolatedV2ArrayOutput
}

type NetworkIsolatedV2Array []NetworkIsolatedV2Input

func (NetworkIsolatedV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkIsolatedV2)(nil)).Elem()
}

func (i NetworkIsolatedV2Array) ToNetworkIsolatedV2ArrayOutput() NetworkIsolatedV2ArrayOutput {
	return i.ToNetworkIsolatedV2ArrayOutputWithContext(context.Background())
}

func (i NetworkIsolatedV2Array) ToNetworkIsolatedV2ArrayOutputWithContext(ctx context.Context) NetworkIsolatedV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkIsolatedV2ArrayOutput)
}

// NetworkIsolatedV2MapInput is an input type that accepts NetworkIsolatedV2Map and NetworkIsolatedV2MapOutput values.
// You can construct a concrete instance of `NetworkIsolatedV2MapInput` via:
//
//	NetworkIsolatedV2Map{ "key": NetworkIsolatedV2Args{...} }
type NetworkIsolatedV2MapInput interface {
	pulumi.Input

	ToNetworkIsolatedV2MapOutput() NetworkIsolatedV2MapOutput
	ToNetworkIsolatedV2MapOutputWithContext(context.Context) NetworkIsolatedV2MapOutput
}

type NetworkIsolatedV2Map map[string]NetworkIsolatedV2Input

func (NetworkIsolatedV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkIsolatedV2)(nil)).Elem()
}

func (i NetworkIsolatedV2Map) ToNetworkIsolatedV2MapOutput() NetworkIsolatedV2MapOutput {
	return i.ToNetworkIsolatedV2MapOutputWithContext(context.Background())
}

func (i NetworkIsolatedV2Map) ToNetworkIsolatedV2MapOutputWithContext(ctx context.Context) NetworkIsolatedV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkIsolatedV2MapOutput)
}

type NetworkIsolatedV2Output struct{ *pulumi.OutputState }

func (NetworkIsolatedV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkIsolatedV2)(nil)).Elem()
}

func (o NetworkIsolatedV2Output) ToNetworkIsolatedV2Output() NetworkIsolatedV2Output {
	return o
}

func (o NetworkIsolatedV2Output) ToNetworkIsolatedV2OutputWithContext(ctx context.Context) NetworkIsolatedV2Output {
	return o
}

// Network description
func (o NetworkIsolatedV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// DNS server 1
func (o NetworkIsolatedV2Output) Dns1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Dns1 }).(pulumi.StringPtrOutput)
}

// DNS server 1
func (o NetworkIsolatedV2Output) Dns2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Dns2 }).(pulumi.StringPtrOutput)
}

// DNS suffix
func (o NetworkIsolatedV2Output) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

// Gateway IP address
func (o NetworkIsolatedV2Output) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// NSX-V only - share this network with other VDCs in this organization. Default - false
func (o NetworkIsolatedV2Output) IsShared() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.BoolOutput { return v.IsShared }).(pulumi.BoolOutput)
}

// Key value map of metadata to assign to this network. Key and value can be any string
//
// Deprecated: Use metadata_entry instead
func (o NetworkIsolatedV2Output) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// Metadata entries for the given Network
func (o NetworkIsolatedV2Output) MetadataEntries() NetworkIsolatedV2MetadataEntryArrayOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) NetworkIsolatedV2MetadataEntryArrayOutput { return v.MetadataEntries }).(NetworkIsolatedV2MetadataEntryArrayOutput)
}

// Network name
func (o NetworkIsolatedV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NetworkIsolatedV2Output) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// ID of VDC or VDC Group
func (o NetworkIsolatedV2Output) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Network prefix
func (o NetworkIsolatedV2Output) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.IntOutput { return v.PrefixLength }).(pulumi.IntOutput)
}

// IP ranges used for static pool allocation in the network
func (o NetworkIsolatedV2Output) StaticIpPools() NetworkIsolatedV2StaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) NetworkIsolatedV2StaticIpPoolArrayOutput { return v.StaticIpPools }).(NetworkIsolatedV2StaticIpPoolArrayOutput)
}

// The name of VDC to use, optional if defined at provider level
//
// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
func (o NetworkIsolatedV2Output) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.Vdc }).(pulumi.StringOutput)
}

type NetworkIsolatedV2ArrayOutput struct{ *pulumi.OutputState }

func (NetworkIsolatedV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkIsolatedV2)(nil)).Elem()
}

func (o NetworkIsolatedV2ArrayOutput) ToNetworkIsolatedV2ArrayOutput() NetworkIsolatedV2ArrayOutput {
	return o
}

func (o NetworkIsolatedV2ArrayOutput) ToNetworkIsolatedV2ArrayOutputWithContext(ctx context.Context) NetworkIsolatedV2ArrayOutput {
	return o
}

func (o NetworkIsolatedV2ArrayOutput) Index(i pulumi.IntInput) NetworkIsolatedV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkIsolatedV2 {
		return vs[0].([]*NetworkIsolatedV2)[vs[1].(int)]
	}).(NetworkIsolatedV2Output)
}

type NetworkIsolatedV2MapOutput struct{ *pulumi.OutputState }

func (NetworkIsolatedV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkIsolatedV2)(nil)).Elem()
}

func (o NetworkIsolatedV2MapOutput) ToNetworkIsolatedV2MapOutput() NetworkIsolatedV2MapOutput {
	return o
}

func (o NetworkIsolatedV2MapOutput) ToNetworkIsolatedV2MapOutputWithContext(ctx context.Context) NetworkIsolatedV2MapOutput {
	return o
}

func (o NetworkIsolatedV2MapOutput) MapIndex(k pulumi.StringInput) NetworkIsolatedV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkIsolatedV2 {
		return vs[0].(map[string]*NetworkIsolatedV2)[vs[1].(string)]
	}).(NetworkIsolatedV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkIsolatedV2Input)(nil)).Elem(), &NetworkIsolatedV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkIsolatedV2ArrayInput)(nil)).Elem(), NetworkIsolatedV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkIsolatedV2MapInput)(nil)).Elem(), NetworkIsolatedV2Map{})
	pulumi.RegisterOutputType(NetworkIsolatedV2Output{})
	pulumi.RegisterOutputType(NetworkIsolatedV2ArrayOutput{})
	pulumi.RegisterOutputType(NetworkIsolatedV2MapOutput{})
}