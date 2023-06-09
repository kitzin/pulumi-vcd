// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EdgeGatewayVpn struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput               `pulumi:"description"`
	EdgeGateway        pulumi.StringOutput                  `pulumi:"edgeGateway"`
	EncryptionProtocol pulumi.StringOutput                  `pulumi:"encryptionProtocol"`
	LocalId            pulumi.StringOutput                  `pulumi:"localId"`
	LocalIpAddress     pulumi.StringOutput                  `pulumi:"localIpAddress"`
	LocalSubnets       EdgeGatewayVpnLocalSubnetArrayOutput `pulumi:"localSubnets"`
	Mtu                pulumi.IntOutput                     `pulumi:"mtu"`
	Name               pulumi.StringOutput                  `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org           pulumi.StringPtrOutput              `pulumi:"org"`
	PeerId        pulumi.StringOutput                 `pulumi:"peerId"`
	PeerIpAddress pulumi.StringOutput                 `pulumi:"peerIpAddress"`
	PeerSubnets   EdgeGatewayVpnPeerSubnetArrayOutput `pulumi:"peerSubnets"`
	SharedSecret  pulumi.StringOutput                 `pulumi:"sharedSecret"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewEdgeGatewayVpn registers a new resource with the given unique name, arguments, and options.
func NewEdgeGatewayVpn(ctx *pulumi.Context,
	name string, args *EdgeGatewayVpnArgs, opts ...pulumi.ResourceOption) (*EdgeGatewayVpn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGateway == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGateway'")
	}
	if args.EncryptionProtocol == nil {
		return nil, errors.New("invalid value for required argument 'EncryptionProtocol'")
	}
	if args.LocalId == nil {
		return nil, errors.New("invalid value for required argument 'LocalId'")
	}
	if args.LocalIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'LocalIpAddress'")
	}
	if args.Mtu == nil {
		return nil, errors.New("invalid value for required argument 'Mtu'")
	}
	if args.PeerId == nil {
		return nil, errors.New("invalid value for required argument 'PeerId'")
	}
	if args.PeerIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'PeerIpAddress'")
	}
	if args.SharedSecret == nil {
		return nil, errors.New("invalid value for required argument 'SharedSecret'")
	}
	var resource EdgeGatewayVpn
	err := ctx.RegisterResource("vcd:index/edgeGatewayVpn:EdgeGatewayVpn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeGatewayVpn gets an existing EdgeGatewayVpn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeGatewayVpn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeGatewayVpnState, opts ...pulumi.ResourceOption) (*EdgeGatewayVpn, error) {
	var resource EdgeGatewayVpn
	err := ctx.ReadResource("vcd:index/edgeGatewayVpn:EdgeGatewayVpn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeGatewayVpn resources.
type edgeGatewayVpnState struct {
	Description        *string                     `pulumi:"description"`
	EdgeGateway        *string                     `pulumi:"edgeGateway"`
	EncryptionProtocol *string                     `pulumi:"encryptionProtocol"`
	LocalId            *string                     `pulumi:"localId"`
	LocalIpAddress     *string                     `pulumi:"localIpAddress"`
	LocalSubnets       []EdgeGatewayVpnLocalSubnet `pulumi:"localSubnets"`
	Mtu                *int                        `pulumi:"mtu"`
	Name               *string                     `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org           *string                    `pulumi:"org"`
	PeerId        *string                    `pulumi:"peerId"`
	PeerIpAddress *string                    `pulumi:"peerIpAddress"`
	PeerSubnets   []EdgeGatewayVpnPeerSubnet `pulumi:"peerSubnets"`
	SharedSecret  *string                    `pulumi:"sharedSecret"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type EdgeGatewayVpnState struct {
	Description        pulumi.StringPtrInput
	EdgeGateway        pulumi.StringPtrInput
	EncryptionProtocol pulumi.StringPtrInput
	LocalId            pulumi.StringPtrInput
	LocalIpAddress     pulumi.StringPtrInput
	LocalSubnets       EdgeGatewayVpnLocalSubnetArrayInput
	Mtu                pulumi.IntPtrInput
	Name               pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org           pulumi.StringPtrInput
	PeerId        pulumi.StringPtrInput
	PeerIpAddress pulumi.StringPtrInput
	PeerSubnets   EdgeGatewayVpnPeerSubnetArrayInput
	SharedSecret  pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (EdgeGatewayVpnState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeGatewayVpnState)(nil)).Elem()
}

type edgeGatewayVpnArgs struct {
	Description        *string                     `pulumi:"description"`
	EdgeGateway        string                      `pulumi:"edgeGateway"`
	EncryptionProtocol string                      `pulumi:"encryptionProtocol"`
	LocalId            string                      `pulumi:"localId"`
	LocalIpAddress     string                      `pulumi:"localIpAddress"`
	LocalSubnets       []EdgeGatewayVpnLocalSubnet `pulumi:"localSubnets"`
	Mtu                int                         `pulumi:"mtu"`
	Name               *string                     `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org           *string                    `pulumi:"org"`
	PeerId        string                     `pulumi:"peerId"`
	PeerIpAddress string                     `pulumi:"peerIpAddress"`
	PeerSubnets   []EdgeGatewayVpnPeerSubnet `pulumi:"peerSubnets"`
	SharedSecret  string                     `pulumi:"sharedSecret"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a EdgeGatewayVpn resource.
type EdgeGatewayVpnArgs struct {
	Description        pulumi.StringPtrInput
	EdgeGateway        pulumi.StringInput
	EncryptionProtocol pulumi.StringInput
	LocalId            pulumi.StringInput
	LocalIpAddress     pulumi.StringInput
	LocalSubnets       EdgeGatewayVpnLocalSubnetArrayInput
	Mtu                pulumi.IntInput
	Name               pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org           pulumi.StringPtrInput
	PeerId        pulumi.StringInput
	PeerIpAddress pulumi.StringInput
	PeerSubnets   EdgeGatewayVpnPeerSubnetArrayInput
	SharedSecret  pulumi.StringInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (EdgeGatewayVpnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeGatewayVpnArgs)(nil)).Elem()
}

type EdgeGatewayVpnInput interface {
	pulumi.Input

	ToEdgeGatewayVpnOutput() EdgeGatewayVpnOutput
	ToEdgeGatewayVpnOutputWithContext(ctx context.Context) EdgeGatewayVpnOutput
}

func (*EdgeGatewayVpn) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeGatewayVpn)(nil)).Elem()
}

func (i *EdgeGatewayVpn) ToEdgeGatewayVpnOutput() EdgeGatewayVpnOutput {
	return i.ToEdgeGatewayVpnOutputWithContext(context.Background())
}

func (i *EdgeGatewayVpn) ToEdgeGatewayVpnOutputWithContext(ctx context.Context) EdgeGatewayVpnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeGatewayVpnOutput)
}

// EdgeGatewayVpnArrayInput is an input type that accepts EdgeGatewayVpnArray and EdgeGatewayVpnArrayOutput values.
// You can construct a concrete instance of `EdgeGatewayVpnArrayInput` via:
//
//	EdgeGatewayVpnArray{ EdgeGatewayVpnArgs{...} }
type EdgeGatewayVpnArrayInput interface {
	pulumi.Input

	ToEdgeGatewayVpnArrayOutput() EdgeGatewayVpnArrayOutput
	ToEdgeGatewayVpnArrayOutputWithContext(context.Context) EdgeGatewayVpnArrayOutput
}

type EdgeGatewayVpnArray []EdgeGatewayVpnInput

func (EdgeGatewayVpnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeGatewayVpn)(nil)).Elem()
}

func (i EdgeGatewayVpnArray) ToEdgeGatewayVpnArrayOutput() EdgeGatewayVpnArrayOutput {
	return i.ToEdgeGatewayVpnArrayOutputWithContext(context.Background())
}

func (i EdgeGatewayVpnArray) ToEdgeGatewayVpnArrayOutputWithContext(ctx context.Context) EdgeGatewayVpnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeGatewayVpnArrayOutput)
}

// EdgeGatewayVpnMapInput is an input type that accepts EdgeGatewayVpnMap and EdgeGatewayVpnMapOutput values.
// You can construct a concrete instance of `EdgeGatewayVpnMapInput` via:
//
//	EdgeGatewayVpnMap{ "key": EdgeGatewayVpnArgs{...} }
type EdgeGatewayVpnMapInput interface {
	pulumi.Input

	ToEdgeGatewayVpnMapOutput() EdgeGatewayVpnMapOutput
	ToEdgeGatewayVpnMapOutputWithContext(context.Context) EdgeGatewayVpnMapOutput
}

type EdgeGatewayVpnMap map[string]EdgeGatewayVpnInput

func (EdgeGatewayVpnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeGatewayVpn)(nil)).Elem()
}

func (i EdgeGatewayVpnMap) ToEdgeGatewayVpnMapOutput() EdgeGatewayVpnMapOutput {
	return i.ToEdgeGatewayVpnMapOutputWithContext(context.Background())
}

func (i EdgeGatewayVpnMap) ToEdgeGatewayVpnMapOutputWithContext(ctx context.Context) EdgeGatewayVpnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeGatewayVpnMapOutput)
}

type EdgeGatewayVpnOutput struct{ *pulumi.OutputState }

func (EdgeGatewayVpnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeGatewayVpn)(nil)).Elem()
}

func (o EdgeGatewayVpnOutput) ToEdgeGatewayVpnOutput() EdgeGatewayVpnOutput {
	return o
}

func (o EdgeGatewayVpnOutput) ToEdgeGatewayVpnOutputWithContext(ctx context.Context) EdgeGatewayVpnOutput {
	return o
}

func (o EdgeGatewayVpnOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EdgeGatewayVpnOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.EdgeGateway }).(pulumi.StringOutput)
}

func (o EdgeGatewayVpnOutput) EncryptionProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.EncryptionProtocol }).(pulumi.StringOutput)
}

func (o EdgeGatewayVpnOutput) LocalId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.LocalId }).(pulumi.StringOutput)
}

func (o EdgeGatewayVpnOutput) LocalIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.LocalIpAddress }).(pulumi.StringOutput)
}

func (o EdgeGatewayVpnOutput) LocalSubnets() EdgeGatewayVpnLocalSubnetArrayOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) EdgeGatewayVpnLocalSubnetArrayOutput { return v.LocalSubnets }).(EdgeGatewayVpnLocalSubnetArrayOutput)
}

func (o EdgeGatewayVpnOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.IntOutput { return v.Mtu }).(pulumi.IntOutput)
}

func (o EdgeGatewayVpnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o EdgeGatewayVpnOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

func (o EdgeGatewayVpnOutput) PeerId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.PeerId }).(pulumi.StringOutput)
}

func (o EdgeGatewayVpnOutput) PeerIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.PeerIpAddress }).(pulumi.StringOutput)
}

func (o EdgeGatewayVpnOutput) PeerSubnets() EdgeGatewayVpnPeerSubnetArrayOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) EdgeGatewayVpnPeerSubnetArrayOutput { return v.PeerSubnets }).(EdgeGatewayVpnPeerSubnetArrayOutput)
}

func (o EdgeGatewayVpnOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringOutput { return v.SharedSecret }).(pulumi.StringOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o EdgeGatewayVpnOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeGatewayVpn) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type EdgeGatewayVpnArrayOutput struct{ *pulumi.OutputState }

func (EdgeGatewayVpnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeGatewayVpn)(nil)).Elem()
}

func (o EdgeGatewayVpnArrayOutput) ToEdgeGatewayVpnArrayOutput() EdgeGatewayVpnArrayOutput {
	return o
}

func (o EdgeGatewayVpnArrayOutput) ToEdgeGatewayVpnArrayOutputWithContext(ctx context.Context) EdgeGatewayVpnArrayOutput {
	return o
}

func (o EdgeGatewayVpnArrayOutput) Index(i pulumi.IntInput) EdgeGatewayVpnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgeGatewayVpn {
		return vs[0].([]*EdgeGatewayVpn)[vs[1].(int)]
	}).(EdgeGatewayVpnOutput)
}

type EdgeGatewayVpnMapOutput struct{ *pulumi.OutputState }

func (EdgeGatewayVpnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeGatewayVpn)(nil)).Elem()
}

func (o EdgeGatewayVpnMapOutput) ToEdgeGatewayVpnMapOutput() EdgeGatewayVpnMapOutput {
	return o
}

func (o EdgeGatewayVpnMapOutput) ToEdgeGatewayVpnMapOutputWithContext(ctx context.Context) EdgeGatewayVpnMapOutput {
	return o
}

func (o EdgeGatewayVpnMapOutput) MapIndex(k pulumi.StringInput) EdgeGatewayVpnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgeGatewayVpn {
		return vs[0].(map[string]*EdgeGatewayVpn)[vs[1].(string)]
	}).(EdgeGatewayVpnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeGatewayVpnInput)(nil)).Elem(), &EdgeGatewayVpn{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeGatewayVpnArrayInput)(nil)).Elem(), EdgeGatewayVpnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeGatewayVpnMapInput)(nil)).Elem(), EdgeGatewayVpnMap{})
	pulumi.RegisterOutputType(EdgeGatewayVpnOutput{})
	pulumi.RegisterOutputType(EdgeGatewayVpnArrayOutput{})
	pulumi.RegisterOutputType(EdgeGatewayVpnMapOutput{})
}
