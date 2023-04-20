// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNsxtAlbImportableCloud(ctx *pulumi.Context, args *GetNsxtAlbImportableCloudArgs, opts ...pulumi.InvokeOption) (*GetNsxtAlbImportableCloudResult, error) {
	var rv GetNsxtAlbImportableCloudResult
	err := ctx.Invoke("vcd:index/getNsxtAlbImportableCloud:getNsxtAlbImportableCloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbImportableCloud.
type GetNsxtAlbImportableCloudArgs struct {
	ControllerId string `pulumi:"controllerId"`
	Name         string `pulumi:"name"`
}

// A collection of values returned by getNsxtAlbImportableCloud.
type GetNsxtAlbImportableCloudResult struct {
	AlreadyImported bool   `pulumi:"alreadyImported"`
	ControllerId    string `pulumi:"controllerId"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	NetworkPoolId     string `pulumi:"networkPoolId"`
	NetworkPoolName   string `pulumi:"networkPoolName"`
	TransportZoneName string `pulumi:"transportZoneName"`
}

func GetNsxtAlbImportableCloudOutput(ctx *pulumi.Context, args GetNsxtAlbImportableCloudOutputArgs, opts ...pulumi.InvokeOption) GetNsxtAlbImportableCloudResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNsxtAlbImportableCloudResult, error) {
			args := v.(GetNsxtAlbImportableCloudArgs)
			r, err := GetNsxtAlbImportableCloud(ctx, &args, opts...)
			var s GetNsxtAlbImportableCloudResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNsxtAlbImportableCloudResultOutput)
}

// A collection of arguments for invoking getNsxtAlbImportableCloud.
type GetNsxtAlbImportableCloudOutputArgs struct {
	ControllerId pulumi.StringInput `pulumi:"controllerId"`
	Name         pulumi.StringInput `pulumi:"name"`
}

func (GetNsxtAlbImportableCloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtAlbImportableCloudArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbImportableCloud.
type GetNsxtAlbImportableCloudResultOutput struct{ *pulumi.OutputState }

func (GetNsxtAlbImportableCloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtAlbImportableCloudResult)(nil)).Elem()
}

func (o GetNsxtAlbImportableCloudResultOutput) ToGetNsxtAlbImportableCloudResultOutput() GetNsxtAlbImportableCloudResultOutput {
	return o
}

func (o GetNsxtAlbImportableCloudResultOutput) ToGetNsxtAlbImportableCloudResultOutputWithContext(ctx context.Context) GetNsxtAlbImportableCloudResultOutput {
	return o
}

func (o GetNsxtAlbImportableCloudResultOutput) AlreadyImported() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) bool { return v.AlreadyImported }).(pulumi.BoolOutput)
}

func (o GetNsxtAlbImportableCloudResultOutput) ControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) string { return v.ControllerId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNsxtAlbImportableCloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNsxtAlbImportableCloudResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNsxtAlbImportableCloudResultOutput) NetworkPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) string { return v.NetworkPoolId }).(pulumi.StringOutput)
}

func (o GetNsxtAlbImportableCloudResultOutput) NetworkPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) string { return v.NetworkPoolName }).(pulumi.StringOutput)
}

func (o GetNsxtAlbImportableCloudResultOutput) TransportZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtAlbImportableCloudResult) string { return v.TransportZoneName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNsxtAlbImportableCloudResultOutput{})
}