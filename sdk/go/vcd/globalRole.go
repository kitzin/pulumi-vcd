// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalRole struct {
	pulumi.CustomResourceState

	// Key used for internationalization
	BundleKey pulumi.StringOutput `pulumi:"bundleKey"`
	// Global role description
	Description pulumi.StringOutput `pulumi:"description"`
	// Name of global role.
	Name pulumi.StringOutput `pulumi:"name"`
	// When true, publishes the global role to all tenants
	PublishToAllTenants pulumi.BoolOutput `pulumi:"publishToAllTenants"`
	// Whether this global role is read-only
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// list of rights assigned to this global role
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// list of tenants to which this global role is published
	Tenants pulumi.StringArrayOutput `pulumi:"tenants"`
}

// NewGlobalRole registers a new resource with the given unique name, arguments, and options.
func NewGlobalRole(ctx *pulumi.Context,
	name string, args *GlobalRoleArgs, opts ...pulumi.ResourceOption) (*GlobalRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.PublishToAllTenants == nil {
		return nil, errors.New("invalid value for required argument 'PublishToAllTenants'")
	}
	var resource GlobalRole
	err := ctx.RegisterResource("vcd:index/globalRole:GlobalRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalRole gets an existing GlobalRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalRoleState, opts ...pulumi.ResourceOption) (*GlobalRole, error) {
	var resource GlobalRole
	err := ctx.ReadResource("vcd:index/globalRole:GlobalRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalRole resources.
type globalRoleState struct {
	// Key used for internationalization
	BundleKey *string `pulumi:"bundleKey"`
	// Global role description
	Description *string `pulumi:"description"`
	// Name of global role.
	Name *string `pulumi:"name"`
	// When true, publishes the global role to all tenants
	PublishToAllTenants *bool `pulumi:"publishToAllTenants"`
	// Whether this global role is read-only
	ReadOnly *bool `pulumi:"readOnly"`
	// list of rights assigned to this global role
	Rights []string `pulumi:"rights"`
	// list of tenants to which this global role is published
	Tenants []string `pulumi:"tenants"`
}

type GlobalRoleState struct {
	// Key used for internationalization
	BundleKey pulumi.StringPtrInput
	// Global role description
	Description pulumi.StringPtrInput
	// Name of global role.
	Name pulumi.StringPtrInput
	// When true, publishes the global role to all tenants
	PublishToAllTenants pulumi.BoolPtrInput
	// Whether this global role is read-only
	ReadOnly pulumi.BoolPtrInput
	// list of rights assigned to this global role
	Rights pulumi.StringArrayInput
	// list of tenants to which this global role is published
	Tenants pulumi.StringArrayInput
}

func (GlobalRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalRoleState)(nil)).Elem()
}

type globalRoleArgs struct {
	// Global role description
	Description string `pulumi:"description"`
	// Name of global role.
	Name *string `pulumi:"name"`
	// When true, publishes the global role to all tenants
	PublishToAllTenants bool `pulumi:"publishToAllTenants"`
	// list of rights assigned to this global role
	Rights []string `pulumi:"rights"`
	// list of tenants to which this global role is published
	Tenants []string `pulumi:"tenants"`
}

// The set of arguments for constructing a GlobalRole resource.
type GlobalRoleArgs struct {
	// Global role description
	Description pulumi.StringInput
	// Name of global role.
	Name pulumi.StringPtrInput
	// When true, publishes the global role to all tenants
	PublishToAllTenants pulumi.BoolInput
	// list of rights assigned to this global role
	Rights pulumi.StringArrayInput
	// list of tenants to which this global role is published
	Tenants pulumi.StringArrayInput
}

func (GlobalRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalRoleArgs)(nil)).Elem()
}

type GlobalRoleInput interface {
	pulumi.Input

	ToGlobalRoleOutput() GlobalRoleOutput
	ToGlobalRoleOutputWithContext(ctx context.Context) GlobalRoleOutput
}

func (*GlobalRole) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalRole)(nil)).Elem()
}

func (i *GlobalRole) ToGlobalRoleOutput() GlobalRoleOutput {
	return i.ToGlobalRoleOutputWithContext(context.Background())
}

func (i *GlobalRole) ToGlobalRoleOutputWithContext(ctx context.Context) GlobalRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalRoleOutput)
}

// GlobalRoleArrayInput is an input type that accepts GlobalRoleArray and GlobalRoleArrayOutput values.
// You can construct a concrete instance of `GlobalRoleArrayInput` via:
//
//	GlobalRoleArray{ GlobalRoleArgs{...} }
type GlobalRoleArrayInput interface {
	pulumi.Input

	ToGlobalRoleArrayOutput() GlobalRoleArrayOutput
	ToGlobalRoleArrayOutputWithContext(context.Context) GlobalRoleArrayOutput
}

type GlobalRoleArray []GlobalRoleInput

func (GlobalRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalRole)(nil)).Elem()
}

func (i GlobalRoleArray) ToGlobalRoleArrayOutput() GlobalRoleArrayOutput {
	return i.ToGlobalRoleArrayOutputWithContext(context.Background())
}

func (i GlobalRoleArray) ToGlobalRoleArrayOutputWithContext(ctx context.Context) GlobalRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalRoleArrayOutput)
}

// GlobalRoleMapInput is an input type that accepts GlobalRoleMap and GlobalRoleMapOutput values.
// You can construct a concrete instance of `GlobalRoleMapInput` via:
//
//	GlobalRoleMap{ "key": GlobalRoleArgs{...} }
type GlobalRoleMapInput interface {
	pulumi.Input

	ToGlobalRoleMapOutput() GlobalRoleMapOutput
	ToGlobalRoleMapOutputWithContext(context.Context) GlobalRoleMapOutput
}

type GlobalRoleMap map[string]GlobalRoleInput

func (GlobalRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalRole)(nil)).Elem()
}

func (i GlobalRoleMap) ToGlobalRoleMapOutput() GlobalRoleMapOutput {
	return i.ToGlobalRoleMapOutputWithContext(context.Background())
}

func (i GlobalRoleMap) ToGlobalRoleMapOutputWithContext(ctx context.Context) GlobalRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalRoleMapOutput)
}

type GlobalRoleOutput struct{ *pulumi.OutputState }

func (GlobalRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalRole)(nil)).Elem()
}

func (o GlobalRoleOutput) ToGlobalRoleOutput() GlobalRoleOutput {
	return o
}

func (o GlobalRoleOutput) ToGlobalRoleOutputWithContext(ctx context.Context) GlobalRoleOutput {
	return o
}

// Key used for internationalization
func (o GlobalRoleOutput) BundleKey() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.StringOutput { return v.BundleKey }).(pulumi.StringOutput)
}

// Global role description
func (o GlobalRoleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Name of global role.
func (o GlobalRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// When true, publishes the global role to all tenants
func (o GlobalRoleOutput) PublishToAllTenants() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.BoolOutput { return v.PublishToAllTenants }).(pulumi.BoolOutput)
}

// Whether this global role is read-only
func (o GlobalRoleOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// list of rights assigned to this global role
func (o GlobalRoleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

// list of tenants to which this global role is published
func (o GlobalRoleOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GlobalRole) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

type GlobalRoleArrayOutput struct{ *pulumi.OutputState }

func (GlobalRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalRole)(nil)).Elem()
}

func (o GlobalRoleArrayOutput) ToGlobalRoleArrayOutput() GlobalRoleArrayOutput {
	return o
}

func (o GlobalRoleArrayOutput) ToGlobalRoleArrayOutputWithContext(ctx context.Context) GlobalRoleArrayOutput {
	return o
}

func (o GlobalRoleArrayOutput) Index(i pulumi.IntInput) GlobalRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GlobalRole {
		return vs[0].([]*GlobalRole)[vs[1].(int)]
	}).(GlobalRoleOutput)
}

type GlobalRoleMapOutput struct{ *pulumi.OutputState }

func (GlobalRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalRole)(nil)).Elem()
}

func (o GlobalRoleMapOutput) ToGlobalRoleMapOutput() GlobalRoleMapOutput {
	return o
}

func (o GlobalRoleMapOutput) ToGlobalRoleMapOutputWithContext(ctx context.Context) GlobalRoleMapOutput {
	return o
}

func (o GlobalRoleMapOutput) MapIndex(k pulumi.StringInput) GlobalRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GlobalRole {
		return vs[0].(map[string]*GlobalRole)[vs[1].(string)]
	}).(GlobalRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalRoleInput)(nil)).Elem(), &GlobalRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalRoleArrayInput)(nil)).Elem(), GlobalRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalRoleMapInput)(nil)).Elem(), GlobalRoleMap{})
	pulumi.RegisterOutputType(GlobalRoleOutput{})
	pulumi.RegisterOutputType(GlobalRoleArrayOutput{})
	pulumi.RegisterOutputType(GlobalRoleMapOutput{})
}
