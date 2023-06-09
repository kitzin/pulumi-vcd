// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkDirect struct {
	pulumi.CustomResourceState

	// Optional description for the network
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the external network
	ExternalNetwork pulumi.StringOutput `pulumi:"externalNetwork"`
	// Main DNS of the external network
	ExternalNetworkDns1 pulumi.StringOutput `pulumi:"externalNetworkDns1"`
	// Secondary DNS of the external network
	ExternalNetworkDns2 pulumi.StringOutput `pulumi:"externalNetworkDns2"`
	// DNS suffix of the external network
	ExternalNetworkDnsSuffix pulumi.StringOutput `pulumi:"externalNetworkDnsSuffix"`
	// Gateway of the external network
	ExternalNetworkGateway pulumi.StringOutput `pulumi:"externalNetworkGateway"`
	// Net mask of the external network
	ExternalNetworkNetmask pulumi.StringOutput `pulumi:"externalNetworkNetmask"`
	// Network Hypertext Reference
	Href pulumi.StringOutput `pulumi:"href"`
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Metadata entries for the given Network
	MetadataEntries NetworkDirectMetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// A unique name for this network
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Defines if this network is shared between multiple VDCs in the Org
	Shared pulumi.BoolPtrOutput `pulumi:"shared"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewNetworkDirect registers a new resource with the given unique name, arguments, and options.
func NewNetworkDirect(ctx *pulumi.Context,
	name string, args *NetworkDirectArgs, opts ...pulumi.ResourceOption) (*NetworkDirect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalNetwork == nil {
		return nil, errors.New("invalid value for required argument 'ExternalNetwork'")
	}
	var resource NetworkDirect
	err := ctx.RegisterResource("vcd:index/networkDirect:NetworkDirect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkDirect gets an existing NetworkDirect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkDirect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkDirectState, opts ...pulumi.ResourceOption) (*NetworkDirect, error) {
	var resource NetworkDirect
	err := ctx.ReadResource("vcd:index/networkDirect:NetworkDirect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkDirect resources.
type networkDirectState struct {
	// Optional description for the network
	Description *string `pulumi:"description"`
	// The name of the external network
	ExternalNetwork *string `pulumi:"externalNetwork"`
	// Main DNS of the external network
	ExternalNetworkDns1 *string `pulumi:"externalNetworkDns1"`
	// Secondary DNS of the external network
	ExternalNetworkDns2 *string `pulumi:"externalNetworkDns2"`
	// DNS suffix of the external network
	ExternalNetworkDnsSuffix *string `pulumi:"externalNetworkDnsSuffix"`
	// Gateway of the external network
	ExternalNetworkGateway *string `pulumi:"externalNetworkGateway"`
	// Net mask of the external network
	ExternalNetworkNetmask *string `pulumi:"externalNetworkNetmask"`
	// Network Hypertext Reference
	Href *string `pulumi:"href"`
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata entries for the given Network
	MetadataEntries []NetworkDirectMetadataEntry `pulumi:"metadataEntries"`
	// A unique name for this network
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Defines if this network is shared between multiple VDCs in the Org
	Shared *bool `pulumi:"shared"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type NetworkDirectState struct {
	// Optional description for the network
	Description pulumi.StringPtrInput
	// The name of the external network
	ExternalNetwork pulumi.StringPtrInput
	// Main DNS of the external network
	ExternalNetworkDns1 pulumi.StringPtrInput
	// Secondary DNS of the external network
	ExternalNetworkDns2 pulumi.StringPtrInput
	// DNS suffix of the external network
	ExternalNetworkDnsSuffix pulumi.StringPtrInput
	// Gateway of the external network
	ExternalNetworkGateway pulumi.StringPtrInput
	// Net mask of the external network
	ExternalNetworkNetmask pulumi.StringPtrInput
	// Network Hypertext Reference
	Href pulumi.StringPtrInput
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata pulumi.MapInput
	// Metadata entries for the given Network
	MetadataEntries NetworkDirectMetadataEntryArrayInput
	// A unique name for this network
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Defines if this network is shared between multiple VDCs in the Org
	Shared pulumi.BoolPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NetworkDirectState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkDirectState)(nil)).Elem()
}

type networkDirectArgs struct {
	// Optional description for the network
	Description *string `pulumi:"description"`
	// The name of the external network
	ExternalNetwork string `pulumi:"externalNetwork"`
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata entries for the given Network
	MetadataEntries []NetworkDirectMetadataEntry `pulumi:"metadataEntries"`
	// A unique name for this network
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Defines if this network is shared between multiple VDCs in the Org
	Shared *bool `pulumi:"shared"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NetworkDirect resource.
type NetworkDirectArgs struct {
	// Optional description for the network
	Description pulumi.StringPtrInput
	// The name of the external network
	ExternalNetwork pulumi.StringInput
	// Key value map of metadata to assign to this network. Key and value can be any string
	//
	// Deprecated: Use metadata_entry instead
	Metadata pulumi.MapInput
	// Metadata entries for the given Network
	MetadataEntries NetworkDirectMetadataEntryArrayInput
	// A unique name for this network
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Defines if this network is shared between multiple VDCs in the Org
	Shared pulumi.BoolPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (NetworkDirectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkDirectArgs)(nil)).Elem()
}

type NetworkDirectInput interface {
	pulumi.Input

	ToNetworkDirectOutput() NetworkDirectOutput
	ToNetworkDirectOutputWithContext(ctx context.Context) NetworkDirectOutput
}

func (*NetworkDirect) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkDirect)(nil)).Elem()
}

func (i *NetworkDirect) ToNetworkDirectOutput() NetworkDirectOutput {
	return i.ToNetworkDirectOutputWithContext(context.Background())
}

func (i *NetworkDirect) ToNetworkDirectOutputWithContext(ctx context.Context) NetworkDirectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkDirectOutput)
}

// NetworkDirectArrayInput is an input type that accepts NetworkDirectArray and NetworkDirectArrayOutput values.
// You can construct a concrete instance of `NetworkDirectArrayInput` via:
//
//	NetworkDirectArray{ NetworkDirectArgs{...} }
type NetworkDirectArrayInput interface {
	pulumi.Input

	ToNetworkDirectArrayOutput() NetworkDirectArrayOutput
	ToNetworkDirectArrayOutputWithContext(context.Context) NetworkDirectArrayOutput
}

type NetworkDirectArray []NetworkDirectInput

func (NetworkDirectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkDirect)(nil)).Elem()
}

func (i NetworkDirectArray) ToNetworkDirectArrayOutput() NetworkDirectArrayOutput {
	return i.ToNetworkDirectArrayOutputWithContext(context.Background())
}

func (i NetworkDirectArray) ToNetworkDirectArrayOutputWithContext(ctx context.Context) NetworkDirectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkDirectArrayOutput)
}

// NetworkDirectMapInput is an input type that accepts NetworkDirectMap and NetworkDirectMapOutput values.
// You can construct a concrete instance of `NetworkDirectMapInput` via:
//
//	NetworkDirectMap{ "key": NetworkDirectArgs{...} }
type NetworkDirectMapInput interface {
	pulumi.Input

	ToNetworkDirectMapOutput() NetworkDirectMapOutput
	ToNetworkDirectMapOutputWithContext(context.Context) NetworkDirectMapOutput
}

type NetworkDirectMap map[string]NetworkDirectInput

func (NetworkDirectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkDirect)(nil)).Elem()
}

func (i NetworkDirectMap) ToNetworkDirectMapOutput() NetworkDirectMapOutput {
	return i.ToNetworkDirectMapOutputWithContext(context.Background())
}

func (i NetworkDirectMap) ToNetworkDirectMapOutputWithContext(ctx context.Context) NetworkDirectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkDirectMapOutput)
}

type NetworkDirectOutput struct{ *pulumi.OutputState }

func (NetworkDirectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkDirect)(nil)).Elem()
}

func (o NetworkDirectOutput) ToNetworkDirectOutput() NetworkDirectOutput {
	return o
}

func (o NetworkDirectOutput) ToNetworkDirectOutputWithContext(ctx context.Context) NetworkDirectOutput {
	return o
}

// Optional description for the network
func (o NetworkDirectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the external network
func (o NetworkDirectOutput) ExternalNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.ExternalNetwork }).(pulumi.StringOutput)
}

// Main DNS of the external network
func (o NetworkDirectOutput) ExternalNetworkDns1() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.ExternalNetworkDns1 }).(pulumi.StringOutput)
}

// Secondary DNS of the external network
func (o NetworkDirectOutput) ExternalNetworkDns2() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.ExternalNetworkDns2 }).(pulumi.StringOutput)
}

// DNS suffix of the external network
func (o NetworkDirectOutput) ExternalNetworkDnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.ExternalNetworkDnsSuffix }).(pulumi.StringOutput)
}

// Gateway of the external network
func (o NetworkDirectOutput) ExternalNetworkGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.ExternalNetworkGateway }).(pulumi.StringOutput)
}

// Net mask of the external network
func (o NetworkDirectOutput) ExternalNetworkNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.ExternalNetworkNetmask }).(pulumi.StringOutput)
}

// Network Hypertext Reference
func (o NetworkDirectOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.Href }).(pulumi.StringOutput)
}

// Key value map of metadata to assign to this network. Key and value can be any string
//
// Deprecated: Use metadata_entry instead
func (o NetworkDirectOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// Metadata entries for the given Network
func (o NetworkDirectOutput) MetadataEntries() NetworkDirectMetadataEntryArrayOutput {
	return o.ApplyT(func(v *NetworkDirect) NetworkDirectMetadataEntryArrayOutput { return v.MetadataEntries }).(NetworkDirectMetadataEntryArrayOutput)
}

// A unique name for this network
func (o NetworkDirectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NetworkDirectOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Defines if this network is shared between multiple VDCs in the Org
func (o NetworkDirectOutput) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.BoolPtrOutput { return v.Shared }).(pulumi.BoolPtrOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o NetworkDirectOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkDirect) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type NetworkDirectArrayOutput struct{ *pulumi.OutputState }

func (NetworkDirectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkDirect)(nil)).Elem()
}

func (o NetworkDirectArrayOutput) ToNetworkDirectArrayOutput() NetworkDirectArrayOutput {
	return o
}

func (o NetworkDirectArrayOutput) ToNetworkDirectArrayOutputWithContext(ctx context.Context) NetworkDirectArrayOutput {
	return o
}

func (o NetworkDirectArrayOutput) Index(i pulumi.IntInput) NetworkDirectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkDirect {
		return vs[0].([]*NetworkDirect)[vs[1].(int)]
	}).(NetworkDirectOutput)
}

type NetworkDirectMapOutput struct{ *pulumi.OutputState }

func (NetworkDirectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkDirect)(nil)).Elem()
}

func (o NetworkDirectMapOutput) ToNetworkDirectMapOutput() NetworkDirectMapOutput {
	return o
}

func (o NetworkDirectMapOutput) ToNetworkDirectMapOutputWithContext(ctx context.Context) NetworkDirectMapOutput {
	return o
}

func (o NetworkDirectMapOutput) MapIndex(k pulumi.StringInput) NetworkDirectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkDirect {
		return vs[0].(map[string]*NetworkDirect)[vs[1].(string)]
	}).(NetworkDirectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkDirectInput)(nil)).Elem(), &NetworkDirect{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkDirectArrayInput)(nil)).Elem(), NetworkDirectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkDirectMapInput)(nil)).Elem(), NetworkDirectMap{})
	pulumi.RegisterOutputType(NetworkDirectOutput{})
	pulumi.RegisterOutputType(NetworkDirectArrayOutput{})
	pulumi.RegisterOutputType(NetworkDirectMapOutput{})
}
