// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VappNetwork struct {
	pulumi.CustomResourceState

	// Optional description for the network
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A range of IPs to issue to virtual machines that don't have a static IP
	DhcpPools VappNetworkDhcpPoolArrayOutput `pulumi:"dhcpPools"`
	// Primary DNS server
	Dns1 pulumi.StringPtrOutput `pulumi:"dns1"`
	// Secondary DNS server
	Dns2 pulumi.StringPtrOutput `pulumi:"dns2"`
	// DNS suffix
	DnsSuffix pulumi.StringPtrOutput `pulumi:"dnsSuffix"`
	// Gateway of the network
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed pulumi.BoolPtrOutput `pulumi:"guestVlanAllowed"`
	// vApp network name
	Name pulumi.StringOutput `pulumi:"name"`
	// Netmask address for a subnet. Default is 255.255.255.0
	Netmask pulumi.StringPtrOutput `pulumi:"netmask"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// org network name to which vapp network is connected
	OrgNetworkName pulumi.StringPtrOutput `pulumi:"orgNetworkName"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled pulumi.BoolPtrOutput `pulumi:"retainIpMacEnabled"`
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIpPools VappNetworkStaticIpPoolArrayOutput `pulumi:"staticIpPools"`
	// vApp to use
	VappName pulumi.StringOutput `pulumi:"vappName"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewVappNetwork registers a new resource with the given unique name, arguments, and options.
func NewVappNetwork(ctx *pulumi.Context,
	name string, args *VappNetworkArgs, opts ...pulumi.ResourceOption) (*VappNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.VappName == nil {
		return nil, errors.New("invalid value for required argument 'VappName'")
	}
	var resource VappNetwork
	err := ctx.RegisterResource("vcd:index/vappNetwork:VappNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVappNetwork gets an existing VappNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VappNetworkState, opts ...pulumi.ResourceOption) (*VappNetwork, error) {
	var resource VappNetwork
	err := ctx.ReadResource("vcd:index/vappNetwork:VappNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VappNetwork resources.
type vappNetworkState struct {
	// Optional description for the network
	Description *string `pulumi:"description"`
	// A range of IPs to issue to virtual machines that don't have a static IP
	DhcpPools []VappNetworkDhcpPool `pulumi:"dhcpPools"`
	// Primary DNS server
	Dns1 *string `pulumi:"dns1"`
	// Secondary DNS server
	Dns2 *string `pulumi:"dns2"`
	// DNS suffix
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Gateway of the network
	Gateway *string `pulumi:"gateway"`
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed *bool `pulumi:"guestVlanAllowed"`
	// vApp network name
	Name *string `pulumi:"name"`
	// Netmask address for a subnet. Default is 255.255.255.0
	Netmask *string `pulumi:"netmask"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// org network name to which vapp network is connected
	OrgNetworkName *string `pulumi:"orgNetworkName"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled *bool `pulumi:"retainIpMacEnabled"`
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIpPools []VappNetworkStaticIpPool `pulumi:"staticIpPools"`
	// vApp to use
	VappName *string `pulumi:"vappName"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type VappNetworkState struct {
	// Optional description for the network
	Description pulumi.StringPtrInput
	// A range of IPs to issue to virtual machines that don't have a static IP
	DhcpPools VappNetworkDhcpPoolArrayInput
	// Primary DNS server
	Dns1 pulumi.StringPtrInput
	// Secondary DNS server
	Dns2 pulumi.StringPtrInput
	// DNS suffix
	DnsSuffix pulumi.StringPtrInput
	// Gateway of the network
	Gateway pulumi.StringPtrInput
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed pulumi.BoolPtrInput
	// vApp network name
	Name pulumi.StringPtrInput
	// Netmask address for a subnet. Default is 255.255.255.0
	Netmask pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// org network name to which vapp network is connected
	OrgNetworkName pulumi.StringPtrInput
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled pulumi.BoolPtrInput
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIpPools VappNetworkStaticIpPoolArrayInput
	// vApp to use
	VappName pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (VappNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vappNetworkState)(nil)).Elem()
}

type vappNetworkArgs struct {
	// Optional description for the network
	Description *string `pulumi:"description"`
	// A range of IPs to issue to virtual machines that don't have a static IP
	DhcpPools []VappNetworkDhcpPool `pulumi:"dhcpPools"`
	// Primary DNS server
	Dns1 *string `pulumi:"dns1"`
	// Secondary DNS server
	Dns2 *string `pulumi:"dns2"`
	// DNS suffix
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Gateway of the network
	Gateway string `pulumi:"gateway"`
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed *bool `pulumi:"guestVlanAllowed"`
	// vApp network name
	Name *string `pulumi:"name"`
	// Netmask address for a subnet. Default is 255.255.255.0
	Netmask *string `pulumi:"netmask"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// org network name to which vapp network is connected
	OrgNetworkName *string `pulumi:"orgNetworkName"`
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled *bool `pulumi:"retainIpMacEnabled"`
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIpPools []VappNetworkStaticIpPool `pulumi:"staticIpPools"`
	// vApp to use
	VappName string `pulumi:"vappName"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a VappNetwork resource.
type VappNetworkArgs struct {
	// Optional description for the network
	Description pulumi.StringPtrInput
	// A range of IPs to issue to virtual machines that don't have a static IP
	DhcpPools VappNetworkDhcpPoolArrayInput
	// Primary DNS server
	Dns1 pulumi.StringPtrInput
	// Secondary DNS server
	Dns2 pulumi.StringPtrInput
	// DNS suffix
	DnsSuffix pulumi.StringPtrInput
	// Gateway of the network
	Gateway pulumi.StringInput
	// True if Network allows guest VLAN tagging
	GuestVlanAllowed pulumi.BoolPtrInput
	// vApp network name
	Name pulumi.StringPtrInput
	// Netmask address for a subnet. Default is 255.255.255.0
	Netmask pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// org network name to which vapp network is connected
	OrgNetworkName pulumi.StringPtrInput
	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIpMacEnabled pulumi.BoolPtrInput
	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIpPools VappNetworkStaticIpPoolArrayInput
	// vApp to use
	VappName pulumi.StringInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (VappNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vappNetworkArgs)(nil)).Elem()
}

type VappNetworkInput interface {
	pulumi.Input

	ToVappNetworkOutput() VappNetworkOutput
	ToVappNetworkOutputWithContext(ctx context.Context) VappNetworkOutput
}

func (*VappNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VappNetwork)(nil)).Elem()
}

func (i *VappNetwork) ToVappNetworkOutput() VappNetworkOutput {
	return i.ToVappNetworkOutputWithContext(context.Background())
}

func (i *VappNetwork) ToVappNetworkOutputWithContext(ctx context.Context) VappNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappNetworkOutput)
}

// VappNetworkArrayInput is an input type that accepts VappNetworkArray and VappNetworkArrayOutput values.
// You can construct a concrete instance of `VappNetworkArrayInput` via:
//
//	VappNetworkArray{ VappNetworkArgs{...} }
type VappNetworkArrayInput interface {
	pulumi.Input

	ToVappNetworkArrayOutput() VappNetworkArrayOutput
	ToVappNetworkArrayOutputWithContext(context.Context) VappNetworkArrayOutput
}

type VappNetworkArray []VappNetworkInput

func (VappNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappNetwork)(nil)).Elem()
}

func (i VappNetworkArray) ToVappNetworkArrayOutput() VappNetworkArrayOutput {
	return i.ToVappNetworkArrayOutputWithContext(context.Background())
}

func (i VappNetworkArray) ToVappNetworkArrayOutputWithContext(ctx context.Context) VappNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappNetworkArrayOutput)
}

// VappNetworkMapInput is an input type that accepts VappNetworkMap and VappNetworkMapOutput values.
// You can construct a concrete instance of `VappNetworkMapInput` via:
//
//	VappNetworkMap{ "key": VappNetworkArgs{...} }
type VappNetworkMapInput interface {
	pulumi.Input

	ToVappNetworkMapOutput() VappNetworkMapOutput
	ToVappNetworkMapOutputWithContext(context.Context) VappNetworkMapOutput
}

type VappNetworkMap map[string]VappNetworkInput

func (VappNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappNetwork)(nil)).Elem()
}

func (i VappNetworkMap) ToVappNetworkMapOutput() VappNetworkMapOutput {
	return i.ToVappNetworkMapOutputWithContext(context.Background())
}

func (i VappNetworkMap) ToVappNetworkMapOutputWithContext(ctx context.Context) VappNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappNetworkMapOutput)
}

type VappNetworkOutput struct{ *pulumi.OutputState }

func (VappNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VappNetwork)(nil)).Elem()
}

func (o VappNetworkOutput) ToVappNetworkOutput() VappNetworkOutput {
	return o
}

func (o VappNetworkOutput) ToVappNetworkOutputWithContext(ctx context.Context) VappNetworkOutput {
	return o
}

// Optional description for the network
func (o VappNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A range of IPs to issue to virtual machines that don't have a static IP
func (o VappNetworkOutput) DhcpPools() VappNetworkDhcpPoolArrayOutput {
	return o.ApplyT(func(v *VappNetwork) VappNetworkDhcpPoolArrayOutput { return v.DhcpPools }).(VappNetworkDhcpPoolArrayOutput)
}

// Primary DNS server
func (o VappNetworkOutput) Dns1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.Dns1 }).(pulumi.StringPtrOutput)
}

// Secondary DNS server
func (o VappNetworkOutput) Dns2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.Dns2 }).(pulumi.StringPtrOutput)
}

// DNS suffix
func (o VappNetworkOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

// Gateway of the network
func (o VappNetworkOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// True if Network allows guest VLAN tagging
func (o VappNetworkOutput) GuestVlanAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.BoolPtrOutput { return v.GuestVlanAllowed }).(pulumi.BoolPtrOutput)
}

// vApp network name
func (o VappNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Netmask address for a subnet. Default is 255.255.255.0
func (o VappNetworkOutput) Netmask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.Netmask }).(pulumi.StringPtrOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o VappNetworkOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// org network name to which vapp network is connected
func (o VappNetworkOutput) OrgNetworkName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.OrgNetworkName }).(pulumi.StringPtrOutput)
}

// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
func (o VappNetworkOutput) RetainIpMacEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.BoolPtrOutput { return v.RetainIpMacEnabled }).(pulumi.BoolPtrOutput)
}

// A range of IPs permitted to be used as static IPs for virtual machines
func (o VappNetworkOutput) StaticIpPools() VappNetworkStaticIpPoolArrayOutput {
	return o.ApplyT(func(v *VappNetwork) VappNetworkStaticIpPoolArrayOutput { return v.StaticIpPools }).(VappNetworkStaticIpPoolArrayOutput)
}

// vApp to use
func (o VappNetworkOutput) VappName() pulumi.StringOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringOutput { return v.VappName }).(pulumi.StringOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o VappNetworkOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNetwork) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type VappNetworkArrayOutput struct{ *pulumi.OutputState }

func (VappNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappNetwork)(nil)).Elem()
}

func (o VappNetworkArrayOutput) ToVappNetworkArrayOutput() VappNetworkArrayOutput {
	return o
}

func (o VappNetworkArrayOutput) ToVappNetworkArrayOutputWithContext(ctx context.Context) VappNetworkArrayOutput {
	return o
}

func (o VappNetworkArrayOutput) Index(i pulumi.IntInput) VappNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VappNetwork {
		return vs[0].([]*VappNetwork)[vs[1].(int)]
	}).(VappNetworkOutput)
}

type VappNetworkMapOutput struct{ *pulumi.OutputState }

func (VappNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappNetwork)(nil)).Elem()
}

func (o VappNetworkMapOutput) ToVappNetworkMapOutput() VappNetworkMapOutput {
	return o
}

func (o VappNetworkMapOutput) ToVappNetworkMapOutputWithContext(ctx context.Context) VappNetworkMapOutput {
	return o
}

func (o VappNetworkMapOutput) MapIndex(k pulumi.StringInput) VappNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VappNetwork {
		return vs[0].(map[string]*VappNetwork)[vs[1].(string)]
	}).(VappNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VappNetworkInput)(nil)).Elem(), &VappNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappNetworkArrayInput)(nil)).Elem(), VappNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappNetworkMapInput)(nil)).Elem(), VappNetworkMap{})
	pulumi.RegisterOutputType(VappNetworkOutput{})
	pulumi.RegisterOutputType(VappNetworkArrayOutput{})
	pulumi.RegisterOutputType(VappNetworkMapOutput{})
}