// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscribedCatalog struct {
	pulumi.CustomResourceState

	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks pulumi.BoolPtrOutput `pulumi:"cancelFailedTasks"`
	// Catalog version number. Inherited from publishing catalog and updated on sync.
	CatalogVersion pulumi.IntOutput `pulumi:"catalogVersion"`
	// Time stamp of when the catalog was created
	Created pulumi.StringOutput `pulumi:"created"`
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
	// regardless of their state.
	DeleteForce pulumi.BoolPtrOutput `pulumi:"deleteForce"`
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
	// normally allows removal.
	DeleteRecursive pulumi.BoolPtrOutput `pulumi:"deleteRecursive"`
	// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
	Description pulumi.StringOutput `pulumi:"description"`
	// List of failed synchronization tasks
	FailedTasks pulumi.StringArrayOutput `pulumi:"failedTasks"`
	// Catalog HREF
	Href pulumi.StringOutput `pulumi:"href"`
	// True if this catalog belongs to the current organization.
	IsLocal pulumi.BoolOutput `pulumi:"isLocal"`
	// True if this catalog is published. (Always false)
	IsPublished pulumi.BoolOutput `pulumi:"isPublished"`
	// True if this catalog is shared.
	IsShared pulumi.BoolOutput `pulumi:"isShared"`
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
	// copy of catalog items unless a sync operation is performed.
	MakeLocalCopy pulumi.BoolPtrOutput `pulumi:"makeLocalCopy"`
	// List of Media items in this catalog
	MediaItemLists pulumi.StringArrayOutput `pulumi:"mediaItemLists"`
	// Key and value pairs for catalog metadata. Inherited from publishing catalog
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The name of the catalog
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of Media items this catalog contains.
	NumberOfMedia pulumi.IntOutput `pulumi:"numberOfMedia"`
	// Number of vApp templates this catalog contains.
	NumberOfVappTemplates pulumi.IntOutput `pulumi:"numberOfVappTemplates"`
	// The name of organization to use, optional if defined at provider level.
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Owner name from the catalog.
	OwnerName pulumi.StringOutput `pulumi:"ownerName"`
	// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always
	// SUBSCRIBED)
	PublishSubscriptionType pulumi.StringOutput `pulumi:"publishSubscriptionType"`
	// List of running synchronization tasks
	RunningTasks pulumi.StringArrayOutput `pulumi:"runningTasks"`
	// Optional storage profile ID
	StorageProfileId pulumi.StringPtrOutput `pulumi:"storageProfileId"`
	// If true, saves list of tasks to file for later update
	StoreTasks pulumi.BoolPtrOutput `pulumi:"storeTasks"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
	// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	SubscriptionPassword pulumi.StringOutput `pulumi:"subscriptionPassword"`
	// The URL to subscribe to the external catalog.
	SubscriptionUrl pulumi.StringOutput `pulumi:"subscriptionUrl"`
	// If true, synchronise this catalog and all items
	SyncAll pulumi.BoolPtrOutput `pulumi:"syncAll"`
	// If true, synchronises all media items
	SyncAllMediaItems pulumi.BoolPtrOutput `pulumi:"syncAllMediaItems"`
	// If true, synchronises all vApp templates
	SyncAllVappTemplates pulumi.BoolPtrOutput `pulumi:"syncAllVappTemplates"`
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
	// fetches the items data.
	SyncCatalog pulumi.BoolPtrOutput `pulumi:"syncCatalog"`
	// Synchronises media items from this list of names.
	SyncMediaItems pulumi.StringArrayOutput `pulumi:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh pulumi.BoolPtrOutput `pulumi:"syncOnRefresh"`
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates pulumi.StringArrayOutput `pulumi:"syncVappTemplates"`
	// Where the running tasks IDs have been stored. Only if `store_tasks` is set
	TasksFileName pulumi.StringOutput `pulumi:"tasksFileName"`
	// List of catalog items in this catalog
	VappTemplateLists pulumi.StringArrayOutput `pulumi:"vappTemplateLists"`
}

// NewSubscribedCatalog registers a new resource with the given unique name, arguments, and options.
func NewSubscribedCatalog(ctx *pulumi.Context,
	name string, args *SubscribedCatalogArgs, opts ...pulumi.ResourceOption) (*SubscribedCatalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubscriptionUrl == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionUrl'")
	}
	var resource SubscribedCatalog
	err := ctx.RegisterResource("vcd:index/subscribedCatalog:SubscribedCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscribedCatalog gets an existing SubscribedCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscribedCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscribedCatalogState, opts ...pulumi.ResourceOption) (*SubscribedCatalog, error) {
	var resource SubscribedCatalog
	err := ctx.ReadResource("vcd:index/subscribedCatalog:SubscribedCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscribedCatalog resources.
type subscribedCatalogState struct {
	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks *bool `pulumi:"cancelFailedTasks"`
	// Catalog version number. Inherited from publishing catalog and updated on sync.
	CatalogVersion *int `pulumi:"catalogVersion"`
	// Time stamp of when the catalog was created
	Created *string `pulumi:"created"`
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
	// regardless of their state.
	DeleteForce *bool `pulumi:"deleteForce"`
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
	// normally allows removal.
	DeleteRecursive *bool `pulumi:"deleteRecursive"`
	// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
	Description *string `pulumi:"description"`
	// List of failed synchronization tasks
	FailedTasks []string `pulumi:"failedTasks"`
	// Catalog HREF
	Href *string `pulumi:"href"`
	// True if this catalog belongs to the current organization.
	IsLocal *bool `pulumi:"isLocal"`
	// True if this catalog is published. (Always false)
	IsPublished *bool `pulumi:"isPublished"`
	// True if this catalog is shared.
	IsShared *bool `pulumi:"isShared"`
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
	// copy of catalog items unless a sync operation is performed.
	MakeLocalCopy *bool `pulumi:"makeLocalCopy"`
	// List of Media items in this catalog
	MediaItemLists []string `pulumi:"mediaItemLists"`
	// Key and value pairs for catalog metadata. Inherited from publishing catalog
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the catalog
	Name *string `pulumi:"name"`
	// Number of Media items this catalog contains.
	NumberOfMedia *int `pulumi:"numberOfMedia"`
	// Number of vApp templates this catalog contains.
	NumberOfVappTemplates *int `pulumi:"numberOfVappTemplates"`
	// The name of organization to use, optional if defined at provider level.
	Org *string `pulumi:"org"`
	// Owner name from the catalog.
	OwnerName *string `pulumi:"ownerName"`
	// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always
	// SUBSCRIBED)
	PublishSubscriptionType *string `pulumi:"publishSubscriptionType"`
	// List of running synchronization tasks
	RunningTasks []string `pulumi:"runningTasks"`
	// Optional storage profile ID
	StorageProfileId *string `pulumi:"storageProfileId"`
	// If true, saves list of tasks to file for later update
	StoreTasks *bool `pulumi:"storeTasks"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
	// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	SubscriptionPassword *string `pulumi:"subscriptionPassword"`
	// The URL to subscribe to the external catalog.
	SubscriptionUrl *string `pulumi:"subscriptionUrl"`
	// If true, synchronise this catalog and all items
	SyncAll *bool `pulumi:"syncAll"`
	// If true, synchronises all media items
	SyncAllMediaItems *bool `pulumi:"syncAllMediaItems"`
	// If true, synchronises all vApp templates
	SyncAllVappTemplates *bool `pulumi:"syncAllVappTemplates"`
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
	// fetches the items data.
	SyncCatalog *bool `pulumi:"syncCatalog"`
	// Synchronises media items from this list of names.
	SyncMediaItems []string `pulumi:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh *bool `pulumi:"syncOnRefresh"`
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates []string `pulumi:"syncVappTemplates"`
	// Where the running tasks IDs have been stored. Only if `store_tasks` is set
	TasksFileName *string `pulumi:"tasksFileName"`
	// List of catalog items in this catalog
	VappTemplateLists []string `pulumi:"vappTemplateLists"`
}

type SubscribedCatalogState struct {
	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks pulumi.BoolPtrInput
	// Catalog version number. Inherited from publishing catalog and updated on sync.
	CatalogVersion pulumi.IntPtrInput
	// Time stamp of when the catalog was created
	Created pulumi.StringPtrInput
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
	// regardless of their state.
	DeleteForce pulumi.BoolPtrInput
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
	// normally allows removal.
	DeleteRecursive pulumi.BoolPtrInput
	// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
	Description pulumi.StringPtrInput
	// List of failed synchronization tasks
	FailedTasks pulumi.StringArrayInput
	// Catalog HREF
	Href pulumi.StringPtrInput
	// True if this catalog belongs to the current organization.
	IsLocal pulumi.BoolPtrInput
	// True if this catalog is published. (Always false)
	IsPublished pulumi.BoolPtrInput
	// True if this catalog is shared.
	IsShared pulumi.BoolPtrInput
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
	// copy of catalog items unless a sync operation is performed.
	MakeLocalCopy pulumi.BoolPtrInput
	// List of Media items in this catalog
	MediaItemLists pulumi.StringArrayInput
	// Key and value pairs for catalog metadata. Inherited from publishing catalog
	Metadata pulumi.MapInput
	// The name of the catalog
	Name pulumi.StringPtrInput
	// Number of Media items this catalog contains.
	NumberOfMedia pulumi.IntPtrInput
	// Number of vApp templates this catalog contains.
	NumberOfVappTemplates pulumi.IntPtrInput
	// The name of organization to use, optional if defined at provider level.
	Org pulumi.StringPtrInput
	// Owner name from the catalog.
	OwnerName pulumi.StringPtrInput
	// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always
	// SUBSCRIBED)
	PublishSubscriptionType pulumi.StringPtrInput
	// List of running synchronization tasks
	RunningTasks pulumi.StringArrayInput
	// Optional storage profile ID
	StorageProfileId pulumi.StringPtrInput
	// If true, saves list of tasks to file for later update
	StoreTasks pulumi.BoolPtrInput
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
	// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	SubscriptionPassword pulumi.StringPtrInput
	// The URL to subscribe to the external catalog.
	SubscriptionUrl pulumi.StringPtrInput
	// If true, synchronise this catalog and all items
	SyncAll pulumi.BoolPtrInput
	// If true, synchronises all media items
	SyncAllMediaItems pulumi.BoolPtrInput
	// If true, synchronises all vApp templates
	SyncAllVappTemplates pulumi.BoolPtrInput
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
	// fetches the items data.
	SyncCatalog pulumi.BoolPtrInput
	// Synchronises media items from this list of names.
	SyncMediaItems pulumi.StringArrayInput
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh pulumi.BoolPtrInput
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates pulumi.StringArrayInput
	// Where the running tasks IDs have been stored. Only if `store_tasks` is set
	TasksFileName pulumi.StringPtrInput
	// List of catalog items in this catalog
	VappTemplateLists pulumi.StringArrayInput
}

func (SubscribedCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscribedCatalogState)(nil)).Elem()
}

type subscribedCatalogArgs struct {
	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks *bool `pulumi:"cancelFailedTasks"`
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
	// regardless of their state.
	DeleteForce *bool `pulumi:"deleteForce"`
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
	// normally allows removal.
	DeleteRecursive *bool `pulumi:"deleteRecursive"`
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
	// copy of catalog items unless a sync operation is performed.
	MakeLocalCopy *bool `pulumi:"makeLocalCopy"`
	// The name of the catalog
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level.
	Org *string `pulumi:"org"`
	// Optional storage profile ID
	StorageProfileId *string `pulumi:"storageProfileId"`
	// If true, saves list of tasks to file for later update
	StoreTasks *bool `pulumi:"storeTasks"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
	// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	SubscriptionPassword *string `pulumi:"subscriptionPassword"`
	// The URL to subscribe to the external catalog.
	SubscriptionUrl string `pulumi:"subscriptionUrl"`
	// If true, synchronise this catalog and all items
	SyncAll *bool `pulumi:"syncAll"`
	// If true, synchronises all media items
	SyncAllMediaItems *bool `pulumi:"syncAllMediaItems"`
	// If true, synchronises all vApp templates
	SyncAllVappTemplates *bool `pulumi:"syncAllVappTemplates"`
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
	// fetches the items data.
	SyncCatalog *bool `pulumi:"syncCatalog"`
	// Synchronises media items from this list of names.
	SyncMediaItems []string `pulumi:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh *bool `pulumi:"syncOnRefresh"`
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates []string `pulumi:"syncVappTemplates"`
}

// The set of arguments for constructing a SubscribedCatalog resource.
type SubscribedCatalogArgs struct {
	// When true, the subscribed catalog will attempt canceling failed tasks
	CancelFailedTasks pulumi.BoolPtrInput
	// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
	// regardless of their state.
	DeleteForce pulumi.BoolPtrInput
	// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
	// normally allows removal.
	DeleteRecursive pulumi.BoolPtrInput
	// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
	// copy of catalog items unless a sync operation is performed.
	MakeLocalCopy pulumi.BoolPtrInput
	// The name of the catalog
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level.
	Org pulumi.StringPtrInput
	// Optional storage profile ID
	StorageProfileId pulumi.StringPtrInput
	// If true, saves list of tasks to file for later update
	StoreTasks pulumi.BoolPtrInput
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
	// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
	SubscriptionPassword pulumi.StringPtrInput
	// The URL to subscribe to the external catalog.
	SubscriptionUrl pulumi.StringInput
	// If true, synchronise this catalog and all items
	SyncAll pulumi.BoolPtrInput
	// If true, synchronises all media items
	SyncAllMediaItems pulumi.BoolPtrInput
	// If true, synchronises all vApp templates
	SyncAllVappTemplates pulumi.BoolPtrInput
	// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
	// fetches the items data.
	SyncCatalog pulumi.BoolPtrInput
	// Synchronises media items from this list of names.
	SyncMediaItems pulumi.StringArrayInput
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh pulumi.BoolPtrInput
	// Synchronises vApp templates from this list of names.
	SyncVappTemplates pulumi.StringArrayInput
}

func (SubscribedCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscribedCatalogArgs)(nil)).Elem()
}

type SubscribedCatalogInput interface {
	pulumi.Input

	ToSubscribedCatalogOutput() SubscribedCatalogOutput
	ToSubscribedCatalogOutputWithContext(ctx context.Context) SubscribedCatalogOutput
}

func (*SubscribedCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscribedCatalog)(nil)).Elem()
}

func (i *SubscribedCatalog) ToSubscribedCatalogOutput() SubscribedCatalogOutput {
	return i.ToSubscribedCatalogOutputWithContext(context.Background())
}

func (i *SubscribedCatalog) ToSubscribedCatalogOutputWithContext(ctx context.Context) SubscribedCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribedCatalogOutput)
}

// SubscribedCatalogArrayInput is an input type that accepts SubscribedCatalogArray and SubscribedCatalogArrayOutput values.
// You can construct a concrete instance of `SubscribedCatalogArrayInput` via:
//
//	SubscribedCatalogArray{ SubscribedCatalogArgs{...} }
type SubscribedCatalogArrayInput interface {
	pulumi.Input

	ToSubscribedCatalogArrayOutput() SubscribedCatalogArrayOutput
	ToSubscribedCatalogArrayOutputWithContext(context.Context) SubscribedCatalogArrayOutput
}

type SubscribedCatalogArray []SubscribedCatalogInput

func (SubscribedCatalogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscribedCatalog)(nil)).Elem()
}

func (i SubscribedCatalogArray) ToSubscribedCatalogArrayOutput() SubscribedCatalogArrayOutput {
	return i.ToSubscribedCatalogArrayOutputWithContext(context.Background())
}

func (i SubscribedCatalogArray) ToSubscribedCatalogArrayOutputWithContext(ctx context.Context) SubscribedCatalogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribedCatalogArrayOutput)
}

// SubscribedCatalogMapInput is an input type that accepts SubscribedCatalogMap and SubscribedCatalogMapOutput values.
// You can construct a concrete instance of `SubscribedCatalogMapInput` via:
//
//	SubscribedCatalogMap{ "key": SubscribedCatalogArgs{...} }
type SubscribedCatalogMapInput interface {
	pulumi.Input

	ToSubscribedCatalogMapOutput() SubscribedCatalogMapOutput
	ToSubscribedCatalogMapOutputWithContext(context.Context) SubscribedCatalogMapOutput
}

type SubscribedCatalogMap map[string]SubscribedCatalogInput

func (SubscribedCatalogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscribedCatalog)(nil)).Elem()
}

func (i SubscribedCatalogMap) ToSubscribedCatalogMapOutput() SubscribedCatalogMapOutput {
	return i.ToSubscribedCatalogMapOutputWithContext(context.Background())
}

func (i SubscribedCatalogMap) ToSubscribedCatalogMapOutputWithContext(ctx context.Context) SubscribedCatalogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribedCatalogMapOutput)
}

type SubscribedCatalogOutput struct{ *pulumi.OutputState }

func (SubscribedCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscribedCatalog)(nil)).Elem()
}

func (o SubscribedCatalogOutput) ToSubscribedCatalogOutput() SubscribedCatalogOutput {
	return o
}

func (o SubscribedCatalogOutput) ToSubscribedCatalogOutputWithContext(ctx context.Context) SubscribedCatalogOutput {
	return o
}

// When true, the subscribed catalog will attempt canceling failed tasks
func (o SubscribedCatalogOutput) CancelFailedTasks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.CancelFailedTasks }).(pulumi.BoolPtrOutput)
}

// Catalog version number. Inherited from publishing catalog and updated on sync.
func (o SubscribedCatalogOutput) CatalogVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.IntOutput { return v.CatalogVersion }).(pulumi.IntOutput)
}

// Time stamp of when the catalog was created
func (o SubscribedCatalogOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
// regardless of their state.
func (o SubscribedCatalogOutput) DeleteForce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.DeleteForce }).(pulumi.BoolPtrOutput)
}

// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
// normally allows removal.
func (o SubscribedCatalogOutput) DeleteRecursive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.DeleteRecursive }).(pulumi.BoolPtrOutput)
}

// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
func (o SubscribedCatalogOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// List of failed synchronization tasks
func (o SubscribedCatalogOutput) FailedTasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.FailedTasks }).(pulumi.StringArrayOutput)
}

// Catalog HREF
func (o SubscribedCatalogOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Href }).(pulumi.StringOutput)
}

// True if this catalog belongs to the current organization.
func (o SubscribedCatalogOutput) IsLocal() pulumi.BoolOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolOutput { return v.IsLocal }).(pulumi.BoolOutput)
}

// True if this catalog is published. (Always false)
func (o SubscribedCatalogOutput) IsPublished() pulumi.BoolOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolOutput { return v.IsPublished }).(pulumi.BoolOutput)
}

// True if this catalog is shared.
func (o SubscribedCatalogOutput) IsShared() pulumi.BoolOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolOutput { return v.IsShared }).(pulumi.BoolOutput)
}

// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
// copy of catalog items unless a sync operation is performed.
func (o SubscribedCatalogOutput) MakeLocalCopy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.MakeLocalCopy }).(pulumi.BoolPtrOutput)
}

// List of Media items in this catalog
func (o SubscribedCatalogOutput) MediaItemLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.MediaItemLists }).(pulumi.StringArrayOutput)
}

// Key and value pairs for catalog metadata. Inherited from publishing catalog
func (o SubscribedCatalogOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// The name of the catalog
func (o SubscribedCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of Media items this catalog contains.
func (o SubscribedCatalogOutput) NumberOfMedia() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.IntOutput { return v.NumberOfMedia }).(pulumi.IntOutput)
}

// Number of vApp templates this catalog contains.
func (o SubscribedCatalogOutput) NumberOfVappTemplates() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.IntOutput { return v.NumberOfVappTemplates }).(pulumi.IntOutput)
}

// The name of organization to use, optional if defined at provider level.
func (o SubscribedCatalogOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Owner name from the catalog.
func (o SubscribedCatalogOutput) OwnerName() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.OwnerName }).(pulumi.StringOutput)
}

// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always
// SUBSCRIBED)
func (o SubscribedCatalogOutput) PublishSubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.PublishSubscriptionType }).(pulumi.StringOutput)
}

// List of running synchronization tasks
func (o SubscribedCatalogOutput) RunningTasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.RunningTasks }).(pulumi.StringArrayOutput)
}

// Optional storage profile ID
func (o SubscribedCatalogOutput) StorageProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringPtrOutput { return v.StorageProfileId }).(pulumi.StringPtrOutput)
}

// If true, saves list of tasks to file for later update
func (o SubscribedCatalogOutput) StoreTasks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.StoreTasks }).(pulumi.BoolPtrOutput)
}

// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
func (o SubscribedCatalogOutput) SubscriptionPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.SubscriptionPassword }).(pulumi.StringOutput)
}

// The URL to subscribe to the external catalog.
func (o SubscribedCatalogOutput) SubscriptionUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.SubscriptionUrl }).(pulumi.StringOutput)
}

// If true, synchronise this catalog and all items
func (o SubscribedCatalogOutput) SyncAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncAll }).(pulumi.BoolPtrOutput)
}

// If true, synchronises all media items
func (o SubscribedCatalogOutput) SyncAllMediaItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncAllMediaItems }).(pulumi.BoolPtrOutput)
}

// If true, synchronises all vApp templates
func (o SubscribedCatalogOutput) SyncAllVappTemplates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncAllVappTemplates }).(pulumi.BoolPtrOutput)
}

// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
// fetches the items data.
func (o SubscribedCatalogOutput) SyncCatalog() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncCatalog }).(pulumi.BoolPtrOutput)
}

// Synchronises media items from this list of names.
func (o SubscribedCatalogOutput) SyncMediaItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.SyncMediaItems }).(pulumi.StringArrayOutput)
}

// Boolean value that shows if sync should be performed on every refresh.
func (o SubscribedCatalogOutput) SyncOnRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncOnRefresh }).(pulumi.BoolPtrOutput)
}

// Synchronises vApp templates from this list of names.
func (o SubscribedCatalogOutput) SyncVappTemplates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.SyncVappTemplates }).(pulumi.StringArrayOutput)
}

// Where the running tasks IDs have been stored. Only if `store_tasks` is set
func (o SubscribedCatalogOutput) TasksFileName() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.TasksFileName }).(pulumi.StringOutput)
}

// List of catalog items in this catalog
func (o SubscribedCatalogOutput) VappTemplateLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.VappTemplateLists }).(pulumi.StringArrayOutput)
}

type SubscribedCatalogArrayOutput struct{ *pulumi.OutputState }

func (SubscribedCatalogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscribedCatalog)(nil)).Elem()
}

func (o SubscribedCatalogArrayOutput) ToSubscribedCatalogArrayOutput() SubscribedCatalogArrayOutput {
	return o
}

func (o SubscribedCatalogArrayOutput) ToSubscribedCatalogArrayOutputWithContext(ctx context.Context) SubscribedCatalogArrayOutput {
	return o
}

func (o SubscribedCatalogArrayOutput) Index(i pulumi.IntInput) SubscribedCatalogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubscribedCatalog {
		return vs[0].([]*SubscribedCatalog)[vs[1].(int)]
	}).(SubscribedCatalogOutput)
}

type SubscribedCatalogMapOutput struct{ *pulumi.OutputState }

func (SubscribedCatalogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscribedCatalog)(nil)).Elem()
}

func (o SubscribedCatalogMapOutput) ToSubscribedCatalogMapOutput() SubscribedCatalogMapOutput {
	return o
}

func (o SubscribedCatalogMapOutput) ToSubscribedCatalogMapOutputWithContext(ctx context.Context) SubscribedCatalogMapOutput {
	return o
}

func (o SubscribedCatalogMapOutput) MapIndex(k pulumi.StringInput) SubscribedCatalogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubscribedCatalog {
		return vs[0].(map[string]*SubscribedCatalog)[vs[1].(string)]
	}).(SubscribedCatalogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribedCatalogInput)(nil)).Elem(), &SubscribedCatalog{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribedCatalogArrayInput)(nil)).Elem(), SubscribedCatalogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribedCatalogMapInput)(nil)).Elem(), SubscribedCatalogMap{})
	pulumi.RegisterOutputType(SubscribedCatalogOutput{})
	pulumi.RegisterOutputType(SubscribedCatalogArrayOutput{})
	pulumi.RegisterOutputType(SubscribedCatalogMapOutput{})
}