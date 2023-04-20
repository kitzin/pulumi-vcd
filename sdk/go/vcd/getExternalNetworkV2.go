// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExternalNetworkV2(ctx *pulumi.Context, args *LookupExternalNetworkV2Args, opts ...pulumi.InvokeOption) (*LookupExternalNetworkV2Result, error) {
	var rv LookupExternalNetworkV2Result
	err := ctx.Invoke("vcd:index/getExternalNetworkV2:getExternalNetworkV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternalNetworkV2.
type LookupExternalNetworkV2Args struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getExternalNetworkV2.
type LookupExternalNetworkV2Result struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                               `pulumi:"id"`
	IpScopes        []GetExternalNetworkV2IpScope        `pulumi:"ipScopes"`
	Name            string                               `pulumi:"name"`
	NsxtNetworks    []GetExternalNetworkV2NsxtNetwork    `pulumi:"nsxtNetworks"`
	VsphereNetworks []GetExternalNetworkV2VsphereNetwork `pulumi:"vsphereNetworks"`
}

func LookupExternalNetworkV2Output(ctx *pulumi.Context, args LookupExternalNetworkV2OutputArgs, opts ...pulumi.InvokeOption) LookupExternalNetworkV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExternalNetworkV2Result, error) {
			args := v.(LookupExternalNetworkV2Args)
			r, err := LookupExternalNetworkV2(ctx, &args, opts...)
			var s LookupExternalNetworkV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExternalNetworkV2ResultOutput)
}

// A collection of arguments for invoking getExternalNetworkV2.
type LookupExternalNetworkV2OutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupExternalNetworkV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalNetworkV2Args)(nil)).Elem()
}

// A collection of values returned by getExternalNetworkV2.
type LookupExternalNetworkV2ResultOutput struct{ *pulumi.OutputState }

func (LookupExternalNetworkV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalNetworkV2Result)(nil)).Elem()
}

func (o LookupExternalNetworkV2ResultOutput) ToLookupExternalNetworkV2ResultOutput() LookupExternalNetworkV2ResultOutput {
	return o
}

func (o LookupExternalNetworkV2ResultOutput) ToLookupExternalNetworkV2ResultOutputWithContext(ctx context.Context) LookupExternalNetworkV2ResultOutput {
	return o
}

func (o LookupExternalNetworkV2ResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkV2Result) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupExternalNetworkV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExternalNetworkV2ResultOutput) IpScopes() GetExternalNetworkV2IpScopeArrayOutput {
	return o.ApplyT(func(v LookupExternalNetworkV2Result) []GetExternalNetworkV2IpScope { return v.IpScopes }).(GetExternalNetworkV2IpScopeArrayOutput)
}

func (o LookupExternalNetworkV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalNetworkV2Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExternalNetworkV2ResultOutput) NsxtNetworks() GetExternalNetworkV2NsxtNetworkArrayOutput {
	return o.ApplyT(func(v LookupExternalNetworkV2Result) []GetExternalNetworkV2NsxtNetwork { return v.NsxtNetworks }).(GetExternalNetworkV2NsxtNetworkArrayOutput)
}

func (o LookupExternalNetworkV2ResultOutput) VsphereNetworks() GetExternalNetworkV2VsphereNetworkArrayOutput {
	return o.ApplyT(func(v LookupExternalNetworkV2Result) []GetExternalNetworkV2VsphereNetwork { return v.VsphereNetworks }).(GetExternalNetworkV2VsphereNetworkArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExternalNetworkV2ResultOutput{})
}
