// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/subscribedCatalog:SubscribedCatalog")]
    public partial class SubscribedCatalog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When true, the subscribed catalog will attempt canceling failed tasks
        /// </summary>
        [Output("cancelFailedTasks")]
        public Output<bool?> CancelFailedTasks { get; private set; } = null!;

        /// <summary>
        /// Catalog version number. Inherited from publishing catalog and updated on sync.
        /// </summary>
        [Output("catalogVersion")]
        public Output<int> CatalogVersion { get; private set; } = null!;

        /// <summary>
        /// Time stamp of when the catalog was created
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Output("deleteForce")]
        public Output<bool?> DeleteForce { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Output("deleteRecursive")]
        public Output<bool?> DeleteRecursive { get; private set; } = null!;

        /// <summary>
        /// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// List of failed synchronization tasks
        /// </summary>
        [Output("failedTasks")]
        public Output<ImmutableArray<string>> FailedTasks { get; private set; } = null!;

        /// <summary>
        /// Catalog HREF
        /// </summary>
        [Output("href")]
        public Output<string> Href { get; private set; } = null!;

        /// <summary>
        /// True if this catalog belongs to the current organization.
        /// </summary>
        [Output("isLocal")]
        public Output<bool> IsLocal { get; private set; } = null!;

        /// <summary>
        /// True if this catalog is published. (Always false)
        /// </summary>
        [Output("isPublished")]
        public Output<bool> IsPublished { get; private set; } = null!;

        /// <summary>
        /// True if this catalog is shared.
        /// </summary>
        [Output("isShared")]
        public Output<bool> IsShared { get; private set; } = null!;

        /// <summary>
        /// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
        /// copy of catalog items unless a sync operation is performed.
        /// </summary>
        [Output("makeLocalCopy")]
        public Output<bool?> MakeLocalCopy { get; private set; } = null!;

        /// <summary>
        /// List of Media items in this catalog
        /// </summary>
        [Output("mediaItemLists")]
        public Output<ImmutableArray<string>> MediaItemLists { get; private set; } = null!;

        /// <summary>
        /// Key and value pairs for catalog metadata. Inherited from publishing catalog
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the catalog
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of Media items this catalog contains.
        /// </summary>
        [Output("numberOfMedia")]
        public Output<int> NumberOfMedia { get; private set; } = null!;

        /// <summary>
        /// Number of vApp templates this catalog contains.
        /// </summary>
        [Output("numberOfVappTemplates")]
        public Output<int> NumberOfVappTemplates { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Owner name from the catalog.
        /// </summary>
        [Output("ownerName")]
        public Output<string> OwnerName { get; private set; } = null!;

        /// <summary>
        /// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always
        /// SUBSCRIBED)
        /// </summary>
        [Output("publishSubscriptionType")]
        public Output<string> PublishSubscriptionType { get; private set; } = null!;

        /// <summary>
        /// List of running synchronization tasks
        /// </summary>
        [Output("runningTasks")]
        public Output<ImmutableArray<string>> RunningTasks { get; private set; } = null!;

        /// <summary>
        /// Optional storage profile ID
        /// </summary>
        [Output("storageProfileId")]
        public Output<string?> StorageProfileId { get; private set; } = null!;

        /// <summary>
        /// If true, saves list of tasks to file for later update
        /// </summary>
        [Output("storeTasks")]
        public Output<bool?> StoreTasks { get; private set; } = null!;

        /// <summary>
        /// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
        /// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
        /// </summary>
        [Output("subscriptionPassword")]
        public Output<string> SubscriptionPassword { get; private set; } = null!;

        /// <summary>
        /// The URL to subscribe to the external catalog.
        /// </summary>
        [Output("subscriptionUrl")]
        public Output<string> SubscriptionUrl { get; private set; } = null!;

        /// <summary>
        /// If true, synchronise this catalog and all items
        /// </summary>
        [Output("syncAll")]
        public Output<bool?> SyncAll { get; private set; } = null!;

        /// <summary>
        /// If true, synchronises all media items
        /// </summary>
        [Output("syncAllMediaItems")]
        public Output<bool?> SyncAllMediaItems { get; private set; } = null!;

        /// <summary>
        /// If true, synchronises all vApp templates
        /// </summary>
        [Output("syncAllVappTemplates")]
        public Output<bool?> SyncAllVappTemplates { get; private set; } = null!;

        /// <summary>
        /// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
        /// fetches the items data.
        /// </summary>
        [Output("syncCatalog")]
        public Output<bool?> SyncCatalog { get; private set; } = null!;

        /// <summary>
        /// Synchronises media items from this list of names.
        /// </summary>
        [Output("syncMediaItems")]
        public Output<ImmutableArray<string>> SyncMediaItems { get; private set; } = null!;

        /// <summary>
        /// Boolean value that shows if sync should be performed on every refresh.
        /// </summary>
        [Output("syncOnRefresh")]
        public Output<bool?> SyncOnRefresh { get; private set; } = null!;

        /// <summary>
        /// Synchronises vApp templates from this list of names.
        /// </summary>
        [Output("syncVappTemplates")]
        public Output<ImmutableArray<string>> SyncVappTemplates { get; private set; } = null!;

        /// <summary>
        /// Where the running tasks IDs have been stored. Only if `store_tasks` is set
        /// </summary>
        [Output("tasksFileName")]
        public Output<string> TasksFileName { get; private set; } = null!;

        /// <summary>
        /// List of catalog items in this catalog
        /// </summary>
        [Output("vappTemplateLists")]
        public Output<ImmutableArray<string>> VappTemplateLists { get; private set; } = null!;


        /// <summary>
        /// Create a SubscribedCatalog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubscribedCatalog(string name, SubscribedCatalogArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/subscribedCatalog:SubscribedCatalog", name, args ?? new SubscribedCatalogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubscribedCatalog(string name, Input<string> id, SubscribedCatalogState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/subscribedCatalog:SubscribedCatalog", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SubscribedCatalog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubscribedCatalog Get(string name, Input<string> id, SubscribedCatalogState? state = null, CustomResourceOptions? options = null)
        {
            return new SubscribedCatalog(name, id, state, options);
        }
    }

    public sealed class SubscribedCatalogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, the subscribed catalog will attempt canceling failed tasks
        /// </summary>
        [Input("cancelFailedTasks")]
        public Input<bool>? CancelFailedTasks { get; set; }

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Input("deleteForce")]
        public Input<bool>? DeleteForce { get; set; }

        /// <summary>
        /// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive")]
        public Input<bool>? DeleteRecursive { get; set; }

        /// <summary>
        /// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
        /// copy of catalog items unless a sync operation is performed.
        /// </summary>
        [Input("makeLocalCopy")]
        public Input<bool>? MakeLocalCopy { get; set; }

        /// <summary>
        /// The name of the catalog
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Optional storage profile ID
        /// </summary>
        [Input("storageProfileId")]
        public Input<string>? StorageProfileId { get; set; }

        /// <summary>
        /// If true, saves list of tasks to file for later update
        /// </summary>
        [Input("storeTasks")]
        public Input<bool>? StoreTasks { get; set; }

        /// <summary>
        /// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
        /// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
        /// </summary>
        [Input("subscriptionPassword")]
        public Input<string>? SubscriptionPassword { get; set; }

        /// <summary>
        /// The URL to subscribe to the external catalog.
        /// </summary>
        [Input("subscriptionUrl", required: true)]
        public Input<string> SubscriptionUrl { get; set; } = null!;

        /// <summary>
        /// If true, synchronise this catalog and all items
        /// </summary>
        [Input("syncAll")]
        public Input<bool>? SyncAll { get; set; }

        /// <summary>
        /// If true, synchronises all media items
        /// </summary>
        [Input("syncAllMediaItems")]
        public Input<bool>? SyncAllMediaItems { get; set; }

        /// <summary>
        /// If true, synchronises all vApp templates
        /// </summary>
        [Input("syncAllVappTemplates")]
        public Input<bool>? SyncAllVappTemplates { get; set; }

        /// <summary>
        /// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
        /// fetches the items data.
        /// </summary>
        [Input("syncCatalog")]
        public Input<bool>? SyncCatalog { get; set; }

        [Input("syncMediaItems")]
        private InputList<string>? _syncMediaItems;

        /// <summary>
        /// Synchronises media items from this list of names.
        /// </summary>
        public InputList<string> SyncMediaItems
        {
            get => _syncMediaItems ?? (_syncMediaItems = new InputList<string>());
            set => _syncMediaItems = value;
        }

        /// <summary>
        /// Boolean value that shows if sync should be performed on every refresh.
        /// </summary>
        [Input("syncOnRefresh")]
        public Input<bool>? SyncOnRefresh { get; set; }

        [Input("syncVappTemplates")]
        private InputList<string>? _syncVappTemplates;

        /// <summary>
        /// Synchronises vApp templates from this list of names.
        /// </summary>
        public InputList<string> SyncVappTemplates
        {
            get => _syncVappTemplates ?? (_syncVappTemplates = new InputList<string>());
            set => _syncVappTemplates = value;
        }

        public SubscribedCatalogArgs()
        {
        }
        public static new SubscribedCatalogArgs Empty => new SubscribedCatalogArgs();
    }

    public sealed class SubscribedCatalogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, the subscribed catalog will attempt canceling failed tasks
        /// </summary>
        [Input("cancelFailedTasks")]
        public Input<bool>? CancelFailedTasks { get; set; }

        /// <summary>
        /// Catalog version number. Inherited from publishing catalog and updated on sync.
        /// </summary>
        [Input("catalogVersion")]
        public Input<int>? CatalogVersion { get; set; }

        /// <summary>
        /// Time stamp of when the catalog was created
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Input("deleteForce")]
        public Input<bool>? DeleteForce { get; set; }

        /// <summary>
        /// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive")]
        public Input<bool>? DeleteRecursive { get; set; }

        /// <summary>
        /// A subscribed catalog description is inherited from the publisher catalog and cannot be changed. It is updated on sync
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("failedTasks")]
        private InputList<string>? _failedTasks;

        /// <summary>
        /// List of failed synchronization tasks
        /// </summary>
        public InputList<string> FailedTasks
        {
            get => _failedTasks ?? (_failedTasks = new InputList<string>());
            set => _failedTasks = value;
        }

        /// <summary>
        /// Catalog HREF
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// True if this catalog belongs to the current organization.
        /// </summary>
        [Input("isLocal")]
        public Input<bool>? IsLocal { get; set; }

        /// <summary>
        /// True if this catalog is published. (Always false)
        /// </summary>
        [Input("isPublished")]
        public Input<bool>? IsPublished { get; set; }

        /// <summary>
        /// True if this catalog is shared.
        /// </summary>
        [Input("isShared")]
        public Input<bool>? IsShared { get; set; }

        /// <summary>
        /// If true, subscription to a catalog creates a local copy of all items. Defaults to false, which does not create a local
        /// copy of catalog items unless a sync operation is performed.
        /// </summary>
        [Input("makeLocalCopy")]
        public Input<bool>? MakeLocalCopy { get; set; }

        [Input("mediaItemLists")]
        private InputList<string>? _mediaItemLists;

        /// <summary>
        /// List of Media items in this catalog
        /// </summary>
        public InputList<string> MediaItemLists
        {
            get => _mediaItemLists ?? (_mediaItemLists = new InputList<string>());
            set => _mediaItemLists = value;
        }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key and value pairs for catalog metadata. Inherited from publishing catalog
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the catalog
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of Media items this catalog contains.
        /// </summary>
        [Input("numberOfMedia")]
        public Input<int>? NumberOfMedia { get; set; }

        /// <summary>
        /// Number of vApp templates this catalog contains.
        /// </summary>
        [Input("numberOfVappTemplates")]
        public Input<int>? NumberOfVappTemplates { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Owner name from the catalog.
        /// </summary>
        [Input("ownerName")]
        public Input<string>? OwnerName { get; set; }

        /// <summary>
        /// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise. (Always
        /// SUBSCRIBED)
        /// </summary>
        [Input("publishSubscriptionType")]
        public Input<string>? PublishSubscriptionType { get; set; }

        [Input("runningTasks")]
        private InputList<string>? _runningTasks;

        /// <summary>
        /// List of running synchronization tasks
        /// </summary>
        public InputList<string> RunningTasks
        {
            get => _runningTasks ?? (_runningTasks = new InputList<string>());
            set => _runningTasks = value;
        }

        /// <summary>
        /// Optional storage profile ID
        /// </summary>
        [Input("storageProfileId")]
        public Input<string>? StorageProfileId { get; set; }

        /// <summary>
        /// If true, saves list of tasks to file for later update
        /// </summary>
        [Input("storeTasks")]
        public Input<bool>? StoreTasks { get; set; }

        /// <summary>
        /// An optional password to access the catalog. Only ASCII characters are allowed in a valid password. Passing in six
        /// asterisks '******' indicates to keep current password. Passing in null or empty string indicates to remove password.
        /// </summary>
        [Input("subscriptionPassword")]
        public Input<string>? SubscriptionPassword { get; set; }

        /// <summary>
        /// The URL to subscribe to the external catalog.
        /// </summary>
        [Input("subscriptionUrl")]
        public Input<string>? SubscriptionUrl { get; set; }

        /// <summary>
        /// If true, synchronise this catalog and all items
        /// </summary>
        [Input("syncAll")]
        public Input<bool>? SyncAll { get; set; }

        /// <summary>
        /// If true, synchronises all media items
        /// </summary>
        [Input("syncAllMediaItems")]
        public Input<bool>? SyncAllMediaItems { get; set; }

        /// <summary>
        /// If true, synchronises all vApp templates
        /// </summary>
        [Input("syncAllVappTemplates")]
        public Input<bool>? SyncAllVappTemplates { get; set; }

        /// <summary>
        /// If true, synchronise this catalog. This operation fetches the list of items. If `make_local_copy` is set, it also
        /// fetches the items data.
        /// </summary>
        [Input("syncCatalog")]
        public Input<bool>? SyncCatalog { get; set; }

        [Input("syncMediaItems")]
        private InputList<string>? _syncMediaItems;

        /// <summary>
        /// Synchronises media items from this list of names.
        /// </summary>
        public InputList<string> SyncMediaItems
        {
            get => _syncMediaItems ?? (_syncMediaItems = new InputList<string>());
            set => _syncMediaItems = value;
        }

        /// <summary>
        /// Boolean value that shows if sync should be performed on every refresh.
        /// </summary>
        [Input("syncOnRefresh")]
        public Input<bool>? SyncOnRefresh { get; set; }

        [Input("syncVappTemplates")]
        private InputList<string>? _syncVappTemplates;

        /// <summary>
        /// Synchronises vApp templates from this list of names.
        /// </summary>
        public InputList<string> SyncVappTemplates
        {
            get => _syncVappTemplates ?? (_syncVappTemplates = new InputList<string>());
            set => _syncVappTemplates = value;
        }

        /// <summary>
        /// Where the running tasks IDs have been stored. Only if `store_tasks` is set
        /// </summary>
        [Input("tasksFileName")]
        public Input<string>? TasksFileName { get; set; }

        [Input("vappTemplateLists")]
        private InputList<string>? _vappTemplateLists;

        /// <summary>
        /// List of catalog items in this catalog
        /// </summary>
        public InputList<string> VappTemplateLists
        {
            get => _vappTemplateLists ?? (_vappTemplateLists = new InputList<string>());
            set => _vappTemplateLists = value;
        }

        public SubscribedCatalogState()
        {
        }
        public static new SubscribedCatalogState Empty => new SubscribedCatalogState();
    }
}
