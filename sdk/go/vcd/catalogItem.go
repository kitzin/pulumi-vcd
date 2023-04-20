// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CatalogItem struct {
	pulumi.CustomResourceState

	// catalog name where upload the OVA file
	Catalog pulumi.StringOutput `pulumi:"catalog"`
	// Key and value pairs for catalog item metadata
	//
	// Deprecated: Use metadata_entry instead
	CatalogItemMetadata pulumi.MapOutput `pulumi:"catalogItemMetadata"`
	// Time stamp of when the item was created
	Created     pulumi.StringOutput    `pulumi:"created"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Key and value pairs for the metadata of the vApp template associated to this catalog item
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Metadata entries for the given Catalog Item
	MetadataEntries CatalogItemMetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// catalog item name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Absolute or relative path to OVA
	OvaPath pulumi.StringPtrOutput `pulumi:"ovaPath"`
	// URL of OVF file
	OvfUrl pulumi.StringPtrOutput `pulumi:"ovfUrl"`
	// shows upload progress in stdout
	ShowUploadProgress pulumi.BoolPtrOutput `pulumi:"showUploadProgress"`
	// size of upload file piece size in mega bytes
	UploadPieceSize pulumi.IntPtrOutput `pulumi:"uploadPieceSize"`
}

// NewCatalogItem registers a new resource with the given unique name, arguments, and options.
func NewCatalogItem(ctx *pulumi.Context,
	name string, args *CatalogItemArgs, opts ...pulumi.ResourceOption) (*CatalogItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Catalog == nil {
		return nil, errors.New("invalid value for required argument 'Catalog'")
	}
	var resource CatalogItem
	err := ctx.RegisterResource("vcd:index/catalogItem:CatalogItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCatalogItem gets an existing CatalogItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalogItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogItemState, opts ...pulumi.ResourceOption) (*CatalogItem, error) {
	var resource CatalogItem
	err := ctx.ReadResource("vcd:index/catalogItem:CatalogItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CatalogItem resources.
type catalogItemState struct {
	// catalog name where upload the OVA file
	Catalog *string `pulumi:"catalog"`
	// Key and value pairs for catalog item metadata
	//
	// Deprecated: Use metadata_entry instead
	CatalogItemMetadata map[string]interface{} `pulumi:"catalogItemMetadata"`
	// Time stamp of when the item was created
	Created     *string `pulumi:"created"`
	Description *string `pulumi:"description"`
	// Key and value pairs for the metadata of the vApp template associated to this catalog item
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata entries for the given Catalog Item
	MetadataEntries []CatalogItemMetadataEntry `pulumi:"metadataEntries"`
	// catalog item name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Absolute or relative path to OVA
	OvaPath *string `pulumi:"ovaPath"`
	// URL of OVF file
	OvfUrl *string `pulumi:"ovfUrl"`
	// shows upload progress in stdout
	ShowUploadProgress *bool `pulumi:"showUploadProgress"`
	// size of upload file piece size in mega bytes
	UploadPieceSize *int `pulumi:"uploadPieceSize"`
}

type CatalogItemState struct {
	// catalog name where upload the OVA file
	Catalog pulumi.StringPtrInput
	// Key and value pairs for catalog item metadata
	//
	// Deprecated: Use metadata_entry instead
	CatalogItemMetadata pulumi.MapInput
	// Time stamp of when the item was created
	Created     pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	// Key and value pairs for the metadata of the vApp template associated to this catalog item
	Metadata pulumi.MapInput
	// Metadata entries for the given Catalog Item
	MetadataEntries CatalogItemMetadataEntryArrayInput
	// catalog item name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Absolute or relative path to OVA
	OvaPath pulumi.StringPtrInput
	// URL of OVF file
	OvfUrl pulumi.StringPtrInput
	// shows upload progress in stdout
	ShowUploadProgress pulumi.BoolPtrInput
	// size of upload file piece size in mega bytes
	UploadPieceSize pulumi.IntPtrInput
}

func (CatalogItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogItemState)(nil)).Elem()
}

type catalogItemArgs struct {
	// catalog name where upload the OVA file
	Catalog string `pulumi:"catalog"`
	// Key and value pairs for catalog item metadata
	//
	// Deprecated: Use metadata_entry instead
	CatalogItemMetadata map[string]interface{} `pulumi:"catalogItemMetadata"`
	Description         *string                `pulumi:"description"`
	// Key and value pairs for the metadata of the vApp template associated to this catalog item
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata entries for the given Catalog Item
	MetadataEntries []CatalogItemMetadataEntry `pulumi:"metadataEntries"`
	// catalog item name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
	// Absolute or relative path to OVA
	OvaPath *string `pulumi:"ovaPath"`
	// URL of OVF file
	OvfUrl *string `pulumi:"ovfUrl"`
	// shows upload progress in stdout
	ShowUploadProgress *bool `pulumi:"showUploadProgress"`
	// size of upload file piece size in mega bytes
	UploadPieceSize *int `pulumi:"uploadPieceSize"`
}

// The set of arguments for constructing a CatalogItem resource.
type CatalogItemArgs struct {
	// catalog name where upload the OVA file
	Catalog pulumi.StringInput
	// Key and value pairs for catalog item metadata
	//
	// Deprecated: Use metadata_entry instead
	CatalogItemMetadata pulumi.MapInput
	Description         pulumi.StringPtrInput
	// Key and value pairs for the metadata of the vApp template associated to this catalog item
	Metadata pulumi.MapInput
	// Metadata entries for the given Catalog Item
	MetadataEntries CatalogItemMetadataEntryArrayInput
	// catalog item name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
	// Absolute or relative path to OVA
	OvaPath pulumi.StringPtrInput
	// URL of OVF file
	OvfUrl pulumi.StringPtrInput
	// shows upload progress in stdout
	ShowUploadProgress pulumi.BoolPtrInput
	// size of upload file piece size in mega bytes
	UploadPieceSize pulumi.IntPtrInput
}

func (CatalogItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogItemArgs)(nil)).Elem()
}

type CatalogItemInput interface {
	pulumi.Input

	ToCatalogItemOutput() CatalogItemOutput
	ToCatalogItemOutputWithContext(ctx context.Context) CatalogItemOutput
}

func (*CatalogItem) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogItem)(nil)).Elem()
}

func (i *CatalogItem) ToCatalogItemOutput() CatalogItemOutput {
	return i.ToCatalogItemOutputWithContext(context.Background())
}

func (i *CatalogItem) ToCatalogItemOutputWithContext(ctx context.Context) CatalogItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemOutput)
}

// CatalogItemArrayInput is an input type that accepts CatalogItemArray and CatalogItemArrayOutput values.
// You can construct a concrete instance of `CatalogItemArrayInput` via:
//
//	CatalogItemArray{ CatalogItemArgs{...} }
type CatalogItemArrayInput interface {
	pulumi.Input

	ToCatalogItemArrayOutput() CatalogItemArrayOutput
	ToCatalogItemArrayOutputWithContext(context.Context) CatalogItemArrayOutput
}

type CatalogItemArray []CatalogItemInput

func (CatalogItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CatalogItem)(nil)).Elem()
}

func (i CatalogItemArray) ToCatalogItemArrayOutput() CatalogItemArrayOutput {
	return i.ToCatalogItemArrayOutputWithContext(context.Background())
}

func (i CatalogItemArray) ToCatalogItemArrayOutputWithContext(ctx context.Context) CatalogItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemArrayOutput)
}

// CatalogItemMapInput is an input type that accepts CatalogItemMap and CatalogItemMapOutput values.
// You can construct a concrete instance of `CatalogItemMapInput` via:
//
//	CatalogItemMap{ "key": CatalogItemArgs{...} }
type CatalogItemMapInput interface {
	pulumi.Input

	ToCatalogItemMapOutput() CatalogItemMapOutput
	ToCatalogItemMapOutputWithContext(context.Context) CatalogItemMapOutput
}

type CatalogItemMap map[string]CatalogItemInput

func (CatalogItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CatalogItem)(nil)).Elem()
}

func (i CatalogItemMap) ToCatalogItemMapOutput() CatalogItemMapOutput {
	return i.ToCatalogItemMapOutputWithContext(context.Background())
}

func (i CatalogItemMap) ToCatalogItemMapOutputWithContext(ctx context.Context) CatalogItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemMapOutput)
}

type CatalogItemOutput struct{ *pulumi.OutputState }

func (CatalogItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogItem)(nil)).Elem()
}

func (o CatalogItemOutput) ToCatalogItemOutput() CatalogItemOutput {
	return o
}

func (o CatalogItemOutput) ToCatalogItemOutputWithContext(ctx context.Context) CatalogItemOutput {
	return o
}

// catalog name where upload the OVA file
func (o CatalogItemOutput) Catalog() pulumi.StringOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringOutput { return v.Catalog }).(pulumi.StringOutput)
}

// Key and value pairs for catalog item metadata
//
// Deprecated: Use metadata_entry instead
func (o CatalogItemOutput) CatalogItemMetadata() pulumi.MapOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.MapOutput { return v.CatalogItemMetadata }).(pulumi.MapOutput)
}

// Time stamp of when the item was created
func (o CatalogItemOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o CatalogItemOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Key and value pairs for the metadata of the vApp template associated to this catalog item
func (o CatalogItemOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// Metadata entries for the given Catalog Item
func (o CatalogItemOutput) MetadataEntries() CatalogItemMetadataEntryArrayOutput {
	return o.ApplyT(func(v *CatalogItem) CatalogItemMetadataEntryArrayOutput { return v.MetadataEntries }).(CatalogItemMetadataEntryArrayOutput)
}

// catalog item name
func (o CatalogItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o CatalogItemOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Absolute or relative path to OVA
func (o CatalogItemOutput) OvaPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringPtrOutput { return v.OvaPath }).(pulumi.StringPtrOutput)
}

// URL of OVF file
func (o CatalogItemOutput) OvfUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.StringPtrOutput { return v.OvfUrl }).(pulumi.StringPtrOutput)
}

// shows upload progress in stdout
func (o CatalogItemOutput) ShowUploadProgress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.BoolPtrOutput { return v.ShowUploadProgress }).(pulumi.BoolPtrOutput)
}

// size of upload file piece size in mega bytes
func (o CatalogItemOutput) UploadPieceSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CatalogItem) pulumi.IntPtrOutput { return v.UploadPieceSize }).(pulumi.IntPtrOutput)
}

type CatalogItemArrayOutput struct{ *pulumi.OutputState }

func (CatalogItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CatalogItem)(nil)).Elem()
}

func (o CatalogItemArrayOutput) ToCatalogItemArrayOutput() CatalogItemArrayOutput {
	return o
}

func (o CatalogItemArrayOutput) ToCatalogItemArrayOutputWithContext(ctx context.Context) CatalogItemArrayOutput {
	return o
}

func (o CatalogItemArrayOutput) Index(i pulumi.IntInput) CatalogItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CatalogItem {
		return vs[0].([]*CatalogItem)[vs[1].(int)]
	}).(CatalogItemOutput)
}

type CatalogItemMapOutput struct{ *pulumi.OutputState }

func (CatalogItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CatalogItem)(nil)).Elem()
}

func (o CatalogItemMapOutput) ToCatalogItemMapOutput() CatalogItemMapOutput {
	return o
}

func (o CatalogItemMapOutput) ToCatalogItemMapOutputWithContext(ctx context.Context) CatalogItemMapOutput {
	return o
}

func (o CatalogItemMapOutput) MapIndex(k pulumi.StringInput) CatalogItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CatalogItem {
		return vs[0].(map[string]*CatalogItem)[vs[1].(string)]
	}).(CatalogItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CatalogItemInput)(nil)).Elem(), &CatalogItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*CatalogItemArrayInput)(nil)).Elem(), CatalogItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CatalogItemMapInput)(nil)).Elem(), CatalogItemMap{})
	pulumi.RegisterOutputType(CatalogItemOutput{})
	pulumi.RegisterOutputType(CatalogItemArrayOutput{})
	pulumi.RegisterOutputType(CatalogItemMapOutput{})
}