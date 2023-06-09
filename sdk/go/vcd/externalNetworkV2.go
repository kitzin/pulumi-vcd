// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExternalNetworkV2 struct {
	pulumi.CustomResourceState

	// Network description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A set of IP scopes for the network
	IpScopes ExternalNetworkV2IpScopeArrayOutput `pulumi:"ipScopes"`
	// Network name
	Name pulumi.StringOutput `pulumi:"name"`
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork ExternalNetworkV2NsxtNetworkPtrOutput `pulumi:"nsxtNetwork"`
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
	// registered with the system.
	VsphereNetworks ExternalNetworkV2VsphereNetworkArrayOutput `pulumi:"vsphereNetworks"`
}

// NewExternalNetworkV2 registers a new resource with the given unique name, arguments, and options.
func NewExternalNetworkV2(ctx *pulumi.Context,
	name string, args *ExternalNetworkV2Args, opts ...pulumi.ResourceOption) (*ExternalNetworkV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpScopes == nil {
		return nil, errors.New("invalid value for required argument 'IpScopes'")
	}
	var resource ExternalNetworkV2
	err := ctx.RegisterResource("vcd:index/externalNetworkV2:ExternalNetworkV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalNetworkV2 gets an existing ExternalNetworkV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalNetworkV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalNetworkV2State, opts ...pulumi.ResourceOption) (*ExternalNetworkV2, error) {
	var resource ExternalNetworkV2
	err := ctx.ReadResource("vcd:index/externalNetworkV2:ExternalNetworkV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalNetworkV2 resources.
type externalNetworkV2State struct {
	// Network description
	Description *string `pulumi:"description"`
	// A set of IP scopes for the network
	IpScopes []ExternalNetworkV2IpScope `pulumi:"ipScopes"`
	// Network name
	Name *string `pulumi:"name"`
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork *ExternalNetworkV2NsxtNetwork `pulumi:"nsxtNetwork"`
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
	// registered with the system.
	VsphereNetworks []ExternalNetworkV2VsphereNetwork `pulumi:"vsphereNetworks"`
}

type ExternalNetworkV2State struct {
	// Network description
	Description pulumi.StringPtrInput
	// A set of IP scopes for the network
	IpScopes ExternalNetworkV2IpScopeArrayInput
	// Network name
	Name pulumi.StringPtrInput
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork ExternalNetworkV2NsxtNetworkPtrInput
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
	// registered with the system.
	VsphereNetworks ExternalNetworkV2VsphereNetworkArrayInput
}

func (ExternalNetworkV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*externalNetworkV2State)(nil)).Elem()
}

type externalNetworkV2Args struct {
	// Network description
	Description *string `pulumi:"description"`
	// A set of IP scopes for the network
	IpScopes []ExternalNetworkV2IpScope `pulumi:"ipScopes"`
	// Network name
	Name *string `pulumi:"name"`
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork *ExternalNetworkV2NsxtNetwork `pulumi:"nsxtNetwork"`
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
	// registered with the system.
	VsphereNetworks []ExternalNetworkV2VsphereNetwork `pulumi:"vsphereNetworks"`
}

// The set of arguments for constructing a ExternalNetworkV2 resource.
type ExternalNetworkV2Args struct {
	// Network description
	Description pulumi.StringPtrInput
	// A set of IP scopes for the network
	IpScopes ExternalNetworkV2IpScopeArrayInput
	// Network name
	Name pulumi.StringPtrInput
	// Reference to NSX-T Tier-0 router or segment and manager
	NsxtNetwork ExternalNetworkV2NsxtNetworkPtrInput
	// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
	// registered with the system.
	VsphereNetworks ExternalNetworkV2VsphereNetworkArrayInput
}

func (ExternalNetworkV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*externalNetworkV2Args)(nil)).Elem()
}

type ExternalNetworkV2Input interface {
	pulumi.Input

	ToExternalNetworkV2Output() ExternalNetworkV2Output
	ToExternalNetworkV2OutputWithContext(ctx context.Context) ExternalNetworkV2Output
}

func (*ExternalNetworkV2) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalNetworkV2)(nil)).Elem()
}

func (i *ExternalNetworkV2) ToExternalNetworkV2Output() ExternalNetworkV2Output {
	return i.ToExternalNetworkV2OutputWithContext(context.Background())
}

func (i *ExternalNetworkV2) ToExternalNetworkV2OutputWithContext(ctx context.Context) ExternalNetworkV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNetworkV2Output)
}

// ExternalNetworkV2ArrayInput is an input type that accepts ExternalNetworkV2Array and ExternalNetworkV2ArrayOutput values.
// You can construct a concrete instance of `ExternalNetworkV2ArrayInput` via:
//
//	ExternalNetworkV2Array{ ExternalNetworkV2Args{...} }
type ExternalNetworkV2ArrayInput interface {
	pulumi.Input

	ToExternalNetworkV2ArrayOutput() ExternalNetworkV2ArrayOutput
	ToExternalNetworkV2ArrayOutputWithContext(context.Context) ExternalNetworkV2ArrayOutput
}

type ExternalNetworkV2Array []ExternalNetworkV2Input

func (ExternalNetworkV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalNetworkV2)(nil)).Elem()
}

func (i ExternalNetworkV2Array) ToExternalNetworkV2ArrayOutput() ExternalNetworkV2ArrayOutput {
	return i.ToExternalNetworkV2ArrayOutputWithContext(context.Background())
}

func (i ExternalNetworkV2Array) ToExternalNetworkV2ArrayOutputWithContext(ctx context.Context) ExternalNetworkV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNetworkV2ArrayOutput)
}

// ExternalNetworkV2MapInput is an input type that accepts ExternalNetworkV2Map and ExternalNetworkV2MapOutput values.
// You can construct a concrete instance of `ExternalNetworkV2MapInput` via:
//
//	ExternalNetworkV2Map{ "key": ExternalNetworkV2Args{...} }
type ExternalNetworkV2MapInput interface {
	pulumi.Input

	ToExternalNetworkV2MapOutput() ExternalNetworkV2MapOutput
	ToExternalNetworkV2MapOutputWithContext(context.Context) ExternalNetworkV2MapOutput
}

type ExternalNetworkV2Map map[string]ExternalNetworkV2Input

func (ExternalNetworkV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalNetworkV2)(nil)).Elem()
}

func (i ExternalNetworkV2Map) ToExternalNetworkV2MapOutput() ExternalNetworkV2MapOutput {
	return i.ToExternalNetworkV2MapOutputWithContext(context.Background())
}

func (i ExternalNetworkV2Map) ToExternalNetworkV2MapOutputWithContext(ctx context.Context) ExternalNetworkV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNetworkV2MapOutput)
}

type ExternalNetworkV2Output struct{ *pulumi.OutputState }

func (ExternalNetworkV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalNetworkV2)(nil)).Elem()
}

func (o ExternalNetworkV2Output) ToExternalNetworkV2Output() ExternalNetworkV2Output {
	return o
}

func (o ExternalNetworkV2Output) ToExternalNetworkV2OutputWithContext(ctx context.Context) ExternalNetworkV2Output {
	return o
}

// Network description
func (o ExternalNetworkV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalNetworkV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A set of IP scopes for the network
func (o ExternalNetworkV2Output) IpScopes() ExternalNetworkV2IpScopeArrayOutput {
	return o.ApplyT(func(v *ExternalNetworkV2) ExternalNetworkV2IpScopeArrayOutput { return v.IpScopes }).(ExternalNetworkV2IpScopeArrayOutput)
}

// Network name
func (o ExternalNetworkV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetworkV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Reference to NSX-T Tier-0 router or segment and manager
func (o ExternalNetworkV2Output) NsxtNetwork() ExternalNetworkV2NsxtNetworkPtrOutput {
	return o.ApplyT(func(v *ExternalNetworkV2) ExternalNetworkV2NsxtNetworkPtrOutput { return v.NsxtNetwork }).(ExternalNetworkV2NsxtNetworkPtrOutput)
}

// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
// registered with the system.
func (o ExternalNetworkV2Output) VsphereNetworks() ExternalNetworkV2VsphereNetworkArrayOutput {
	return o.ApplyT(func(v *ExternalNetworkV2) ExternalNetworkV2VsphereNetworkArrayOutput { return v.VsphereNetworks }).(ExternalNetworkV2VsphereNetworkArrayOutput)
}

type ExternalNetworkV2ArrayOutput struct{ *pulumi.OutputState }

func (ExternalNetworkV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalNetworkV2)(nil)).Elem()
}

func (o ExternalNetworkV2ArrayOutput) ToExternalNetworkV2ArrayOutput() ExternalNetworkV2ArrayOutput {
	return o
}

func (o ExternalNetworkV2ArrayOutput) ToExternalNetworkV2ArrayOutputWithContext(ctx context.Context) ExternalNetworkV2ArrayOutput {
	return o
}

func (o ExternalNetworkV2ArrayOutput) Index(i pulumi.IntInput) ExternalNetworkV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalNetworkV2 {
		return vs[0].([]*ExternalNetworkV2)[vs[1].(int)]
	}).(ExternalNetworkV2Output)
}

type ExternalNetworkV2MapOutput struct{ *pulumi.OutputState }

func (ExternalNetworkV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalNetworkV2)(nil)).Elem()
}

func (o ExternalNetworkV2MapOutput) ToExternalNetworkV2MapOutput() ExternalNetworkV2MapOutput {
	return o
}

func (o ExternalNetworkV2MapOutput) ToExternalNetworkV2MapOutputWithContext(ctx context.Context) ExternalNetworkV2MapOutput {
	return o
}

func (o ExternalNetworkV2MapOutput) MapIndex(k pulumi.StringInput) ExternalNetworkV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalNetworkV2 {
		return vs[0].(map[string]*ExternalNetworkV2)[vs[1].(string)]
	}).(ExternalNetworkV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalNetworkV2Input)(nil)).Elem(), &ExternalNetworkV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalNetworkV2ArrayInput)(nil)).Elem(), ExternalNetworkV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalNetworkV2MapInput)(nil)).Elem(), ExternalNetworkV2Map{})
	pulumi.RegisterOutputType(ExternalNetworkV2Output{})
	pulumi.RegisterOutputType(ExternalNetworkV2ArrayOutput{})
	pulumi.RegisterOutputType(ExternalNetworkV2MapOutput{})
}
