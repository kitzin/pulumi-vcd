// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVappNetwork(ctx *pulumi.Context, args *LookupVappNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVappNetworkResult, error) {
	var rv LookupVappNetworkResult
	err := ctx.Invoke("vcd:index/getVappNetwork:getVappNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVappNetwork.
type LookupVappNetworkArgs struct {
	Name     string  `pulumi:"name"`
	Org      *string `pulumi:"org"`
	VappName string  `pulumi:"vappName"`
	Vdc      *string `pulumi:"vdc"`
}

// A collection of values returned by getVappNetwork.
type LookupVappNetworkResult struct {
	Description      string                   `pulumi:"description"`
	DhcpPools        []GetVappNetworkDhcpPool `pulumi:"dhcpPools"`
	Dns1             string                   `pulumi:"dns1"`
	Dns2             string                   `pulumi:"dns2"`
	DnsSuffix        string                   `pulumi:"dnsSuffix"`
	Gateway          string                   `pulumi:"gateway"`
	GuestVlanAllowed bool                     `pulumi:"guestVlanAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                       `pulumi:"id"`
	Name               string                       `pulumi:"name"`
	Netmask            string                       `pulumi:"netmask"`
	Org                *string                      `pulumi:"org"`
	OrgNetworkName     string                       `pulumi:"orgNetworkName"`
	RetainIpMacEnabled bool                         `pulumi:"retainIpMacEnabled"`
	StaticIpPools      []GetVappNetworkStaticIpPool `pulumi:"staticIpPools"`
	VappName           string                       `pulumi:"vappName"`
	Vdc                *string                      `pulumi:"vdc"`
}

func LookupVappNetworkOutput(ctx *pulumi.Context, args LookupVappNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVappNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVappNetworkResult, error) {
			args := v.(LookupVappNetworkArgs)
			r, err := LookupVappNetwork(ctx, &args, opts...)
			var s LookupVappNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVappNetworkResultOutput)
}

// A collection of arguments for invoking getVappNetwork.
type LookupVappNetworkOutputArgs struct {
	Name     pulumi.StringInput    `pulumi:"name"`
	Org      pulumi.StringPtrInput `pulumi:"org"`
	VappName pulumi.StringInput    `pulumi:"vappName"`
	Vdc      pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupVappNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVappNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getVappNetwork.
type LookupVappNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVappNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVappNetworkResult)(nil)).Elem()
}

func (o LookupVappNetworkResultOutput) ToLookupVappNetworkResultOutput() LookupVappNetworkResultOutput {
	return o
}

func (o LookupVappNetworkResultOutput) ToLookupVappNetworkResultOutputWithContext(ctx context.Context) LookupVappNetworkResultOutput {
	return o
}

func (o LookupVappNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) DhcpPools() GetVappNetworkDhcpPoolArrayOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) []GetVappNetworkDhcpPool { return v.DhcpPools }).(GetVappNetworkDhcpPoolArrayOutput)
}

func (o LookupVappNetworkResultOutput) Dns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Dns1 }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) Dns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Dns2 }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) DnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.DnsSuffix }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) GuestVlanAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) bool { return v.GuestVlanAllowed }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVappNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.Netmask }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupVappNetworkResultOutput) OrgNetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.OrgNetworkName }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) RetainIpMacEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) bool { return v.RetainIpMacEnabled }).(pulumi.BoolOutput)
}

func (o LookupVappNetworkResultOutput) StaticIpPools() GetVappNetworkStaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) []GetVappNetworkStaticIpPool { return v.StaticIpPools }).(GetVappNetworkStaticIpPoolArrayOutput)
}

func (o LookupVappNetworkResultOutput) VappName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) string { return v.VappName }).(pulumi.StringOutput)
}

func (o LookupVappNetworkResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVappNetworkResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVappNetworkResultOutput{})
}
