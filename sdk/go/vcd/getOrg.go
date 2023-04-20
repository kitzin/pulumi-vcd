// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrg(ctx *pulumi.Context, args *LookupOrgArgs, opts ...pulumi.InvokeOption) (*LookupOrgResult, error) {
	var rv LookupOrgResult
	err := ctx.Invoke("vcd:index/getOrg:getOrg", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrg.
type LookupOrgArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getOrg.
type LookupOrgResult struct {
	CanPublishCatalogs           bool   `pulumi:"canPublishCatalogs"`
	CanPublishExternalCatalogs   bool   `pulumi:"canPublishExternalCatalogs"`
	CanSubscribeExternalCatalogs bool   `pulumi:"canSubscribeExternalCatalogs"`
	DelayAfterPowerOnSeconds     int    `pulumi:"delayAfterPowerOnSeconds"`
	DeployedVmQuota              int    `pulumi:"deployedVmQuota"`
	Description                  string `pulumi:"description"`
	FullName                     string `pulumi:"fullName"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	IsEnabled bool   `pulumi:"isEnabled"`
	// Deprecated: Use metadata_entry instead
	Metadata           map[string]interface{}    `pulumi:"metadata"`
	MetadataEntries    []GetOrgMetadataEntry     `pulumi:"metadataEntries"`
	Name               string                    `pulumi:"name"`
	StoredVmQuota      int                       `pulumi:"storedVmQuota"`
	VappLeases         []GetOrgVappLease         `pulumi:"vappLeases"`
	VappTemplateLeases []GetOrgVappTemplateLease `pulumi:"vappTemplateLeases"`
}

func LookupOrgOutput(ctx *pulumi.Context, args LookupOrgOutputArgs, opts ...pulumi.InvokeOption) LookupOrgResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrgResult, error) {
			args := v.(LookupOrgArgs)
			r, err := LookupOrg(ctx, &args, opts...)
			var s LookupOrgResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrgResultOutput)
}

// A collection of arguments for invoking getOrg.
type LookupOrgOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupOrgOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgArgs)(nil)).Elem()
}

// A collection of values returned by getOrg.
type LookupOrgResultOutput struct{ *pulumi.OutputState }

func (LookupOrgResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgResult)(nil)).Elem()
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutput() LookupOrgResultOutput {
	return o
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutputWithContext(ctx context.Context) LookupOrgResultOutput {
	return o
}

func (o LookupOrgResultOutput) CanPublishCatalogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.CanPublishCatalogs }).(pulumi.BoolOutput)
}

func (o LookupOrgResultOutput) CanPublishExternalCatalogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.CanPublishExternalCatalogs }).(pulumi.BoolOutput)
}

func (o LookupOrgResultOutput) CanSubscribeExternalCatalogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.CanSubscribeExternalCatalogs }).(pulumi.BoolOutput)
}

func (o LookupOrgResultOutput) DelayAfterPowerOnSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.DelayAfterPowerOnSeconds }).(pulumi.IntOutput)
}

func (o LookupOrgResultOutput) DeployedVmQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.DeployedVmQuota }).(pulumi.IntOutput)
}

func (o LookupOrgResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupOrgResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrgResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupOrgResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupOrgResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupOrgResultOutput) MetadataEntries() GetOrgMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgMetadataEntry { return v.MetadataEntries }).(GetOrgMetadataEntryArrayOutput)
}

func (o LookupOrgResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrgResultOutput) StoredVmQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.StoredVmQuota }).(pulumi.IntOutput)
}

func (o LookupOrgResultOutput) VappLeases() GetOrgVappLeaseArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgVappLease { return v.VappLeases }).(GetOrgVappLeaseArrayOutput)
}

func (o LookupOrgResultOutput) VappTemplateLeases() GetOrgVappTemplateLeaseArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgVappTemplateLease { return v.VappTemplateLeases }).(GetOrgVappTemplateLeaseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgResultOutput{})
}
