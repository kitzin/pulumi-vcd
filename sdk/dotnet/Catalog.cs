// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/catalog:Catalog")]
    public partial class Catalog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// True enables early catalog export to optimize synchronization
        /// </summary>
        [Output("cacheEnabled")]
        public Output<bool?> CacheEnabled { get; private set; } = null!;

        /// <summary>
        /// Catalog version number.
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
        public Output<bool> DeleteForce { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Output("deleteRecursive")]
        public Output<bool> DeleteRecursive { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

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
        /// True if this catalog is published.
        /// </summary>
        [Output("isPublished")]
        public Output<bool> IsPublished { get; private set; } = null!;

        /// <summary>
        /// True if this catalog is shared.
        /// </summary>
        [Output("isShared")]
        public Output<bool> IsShared { get; private set; } = null!;

        /// <summary>
        /// List of Media items in this catalog
        /// </summary>
        [Output("mediaItemLists")]
        public Output<ImmutableArray<string>> MediaItemLists { get; private set; } = null!;

        /// <summary>
        /// Key and value pairs for catalog metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Metadata entries for the given Catalog
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.CatalogMetadataEntry>> MetadataEntries { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of Medias this catalog contains.
        /// </summary>
        [Output("numberOfMedia")]
        public Output<int> NumberOfMedia { get; private set; } = null!;

        /// <summary>
        /// Number of vApps templates this catalog contains.
        /// </summary>
        [Output("numberOfVappTemplates")]
        public Output<int> NumberOfVappTemplates { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Owner name from the catalog.
        /// </summary>
        [Output("ownerName")]
        public Output<string> OwnerName { get; private set; } = null!;

        /// <summary>
        /// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Include BIOS UUIDs and MAC addresses in the downloaded OVF package. Preserving the identity information limits the
        /// portability of the package and you should use it only when necessary.
        /// </summary>
        [Output("preserveIdentityInformation")]
        public Output<bool?> PreserveIdentityInformation { get; private set; } = null!;

        /// <summary>
        /// True allows to publish a catalog externally to make its vApp templates and media files available for subscription by
        /// organizations outside the Cloud Director installation.
        /// </summary>
        [Output("publishEnabled")]
        public Output<bool?> PublishEnabled { get; private set; } = null!;

        /// <summary>
        /// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise.
        /// </summary>
        [Output("publishSubscriptionType")]
        public Output<string> PublishSubscriptionType { get; private set; } = null!;

        /// <summary>
        /// URL to which other catalogs can subscribe
        /// </summary>
        [Output("publishSubscriptionUrl")]
        public Output<string> PublishSubscriptionUrl { get; private set; } = null!;

        /// <summary>
        /// Optional storage profile ID
        /// </summary>
        [Output("storageProfileId")]
        public Output<string?> StorageProfileId { get; private set; } = null!;

        /// <summary>
        /// List of catalog items in this catalog
        /// </summary>
        [Output("vappTemplateLists")]
        public Output<ImmutableArray<string>> VappTemplateLists { get; private set; } = null!;


        /// <summary>
        /// Create a Catalog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Catalog(string name, CatalogArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/catalog:Catalog", name, args ?? new CatalogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Catalog(string name, Input<string> id, CatalogState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/catalog:Catalog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Catalog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Catalog Get(string name, Input<string> id, CatalogState? state = null, CustomResourceOptions? options = null)
        {
            return new Catalog(name, id, state, options);
        }
    }

    public sealed class CatalogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True enables early catalog export to optimize synchronization
        /// </summary>
        [Input("cacheEnabled")]
        public Input<bool>? CacheEnabled { get; set; }

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove a catalog and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Input("deleteForce", required: true)]
        public Input<bool> DeleteForce { get; set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive", required: true)]
        public Input<bool> DeleteRecursive { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key and value pairs for catalog metadata.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.CatalogMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given Catalog
        /// </summary>
        public InputList<Inputs.CatalogMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.CatalogMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Include BIOS UUIDs and MAC addresses in the downloaded OVF package. Preserving the identity information limits the
        /// portability of the package and you should use it only when necessary.
        /// </summary>
        [Input("preserveIdentityInformation")]
        public Input<bool>? PreserveIdentityInformation { get; set; }

        /// <summary>
        /// True allows to publish a catalog externally to make its vApp templates and media files available for subscription by
        /// organizations outside the Cloud Director installation.
        /// </summary>
        [Input("publishEnabled")]
        public Input<bool>? PublishEnabled { get; set; }

        /// <summary>
        /// Optional storage profile ID
        /// </summary>
        [Input("storageProfileId")]
        public Input<string>? StorageProfileId { get; set; }

        public CatalogArgs()
        {
        }
        public static new CatalogArgs Empty => new CatalogArgs();
    }

    public sealed class CatalogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True enables early catalog export to optimize synchronization
        /// </summary>
        [Input("cacheEnabled")]
        public Input<bool>? CacheEnabled { get; set; }

        /// <summary>
        /// Catalog version number.
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

        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// True if this catalog is published.
        /// </summary>
        [Input("isPublished")]
        public Input<bool>? IsPublished { get; set; }

        /// <summary>
        /// True if this catalog is shared.
        /// </summary>
        [Input("isShared")]
        public Input<bool>? IsShared { get; set; }

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
        /// Key and value pairs for catalog metadata.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.CatalogMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given Catalog
        /// </summary>
        public InputList<Inputs.CatalogMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.CatalogMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of Medias this catalog contains.
        /// </summary>
        [Input("numberOfMedia")]
        public Input<int>? NumberOfMedia { get; set; }

        /// <summary>
        /// Number of vApps templates this catalog contains.
        /// </summary>
        [Input("numberOfVappTemplates")]
        public Input<int>? NumberOfVappTemplates { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Owner name from the catalog.
        /// </summary>
        [Input("ownerName")]
        public Input<string>? OwnerName { get; set; }

        /// <summary>
        /// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Include BIOS UUIDs and MAC addresses in the downloaded OVF package. Preserving the identity information limits the
        /// portability of the package and you should use it only when necessary.
        /// </summary>
        [Input("preserveIdentityInformation")]
        public Input<bool>? PreserveIdentityInformation { get; set; }

        /// <summary>
        /// True allows to publish a catalog externally to make its vApp templates and media files available for subscription by
        /// organizations outside the Cloud Director installation.
        /// </summary>
        [Input("publishEnabled")]
        public Input<bool>? PublishEnabled { get; set; }

        /// <summary>
        /// PUBLISHED if published externally, SUBSCRIBED if subscribed to an external catalog, UNPUBLISHED otherwise.
        /// </summary>
        [Input("publishSubscriptionType")]
        public Input<string>? PublishSubscriptionType { get; set; }

        /// <summary>
        /// URL to which other catalogs can subscribe
        /// </summary>
        [Input("publishSubscriptionUrl")]
        public Input<string>? PublishSubscriptionUrl { get; set; }

        /// <summary>
        /// Optional storage profile ID
        /// </summary>
        [Input("storageProfileId")]
        public Input<string>? StorageProfileId { get; set; }

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

        public CatalogState()
        {
        }
        public static new CatalogState Empty => new CatalogState();
    }
}
