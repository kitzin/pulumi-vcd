// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/vapp:Vapp")]
    public partial class Vapp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional description of the vApp
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Key/value settings for guest properties. Will be picked up by new VMs when created.
        /// </summary>
        [Output("guestProperties")]
        public Output<ImmutableDictionary<string, object>?> GuestProperties { get; private set; } = null!;

        /// <summary>
        /// vApp Hyper Reference
        /// </summary>
        [Output("href")]
        public Output<string> Href { get; private set; } = null!;

        /// <summary>
        /// Defines lease parameters for this vApp
        /// </summary>
        [Output("lease")]
        public Output<Outputs.VappLease> Lease { get; private set; } = null!;

        /// <summary>
        /// Key value map of metadata to assign to this vApp. Key and value can be any string.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Metadata entries for the given vApp
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.VappMetadataEntry>> MetadataEntries { get; private set; } = null!;

        /// <summary>
        /// A name for the vApp, unique withing the VDC
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// A boolean value stating if this vApp should be powered on
        /// </summary>
        [Output("powerOn")]
        public Output<bool?> PowerOn { get; private set; } = null!;

        /// <summary>
        /// Shows the status code of the vApp
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Shows the status of the vApp
        /// </summary>
        [Output("statusText")]
        public Output<string> StatusText { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a Vapp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vapp(string name, VappArgs? args = null, CustomResourceOptions? options = null)
            : base("vcd:index/vapp:Vapp", name, args ?? new VappArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vapp(string name, Input<string> id, VappState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vapp:Vapp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vapp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vapp Get(string name, Input<string> id, VappState? state = null, CustomResourceOptions? options = null)
        {
            return new Vapp(name, id, state, options);
        }
    }

    public sealed class VappArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description of the vApp
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("guestProperties")]
        private InputMap<object>? _guestProperties;

        /// <summary>
        /// Key/value settings for guest properties. Will be picked up by new VMs when created.
        /// </summary>
        public InputMap<object> GuestProperties
        {
            get => _guestProperties ?? (_guestProperties = new InputMap<object>());
            set => _guestProperties = value;
        }

        /// <summary>
        /// Defines lease parameters for this vApp
        /// </summary>
        [Input("lease")]
        public Input<Inputs.VappLeaseArgs>? Lease { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key value map of metadata to assign to this vApp. Key and value can be any string.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.VappMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given vApp
        /// </summary>
        public InputList<Inputs.VappMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.VappMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// A name for the vApp, unique withing the VDC
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A boolean value stating if this vApp should be powered on
        /// </summary>
        [Input("powerOn")]
        public Input<bool>? PowerOn { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappArgs()
        {
        }
        public static new VappArgs Empty => new VappArgs();
    }

    public sealed class VappState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description of the vApp
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("guestProperties")]
        private InputMap<object>? _guestProperties;

        /// <summary>
        /// Key/value settings for guest properties. Will be picked up by new VMs when created.
        /// </summary>
        public InputMap<object> GuestProperties
        {
            get => _guestProperties ?? (_guestProperties = new InputMap<object>());
            set => _guestProperties = value;
        }

        /// <summary>
        /// vApp Hyper Reference
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// Defines lease parameters for this vApp
        /// </summary>
        [Input("lease")]
        public Input<Inputs.VappLeaseGetArgs>? Lease { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key value map of metadata to assign to this vApp. Key and value can be any string.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.VappMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given vApp
        /// </summary>
        public InputList<Inputs.VappMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.VappMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// A name for the vApp, unique withing the VDC
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A boolean value stating if this vApp should be powered on
        /// </summary>
        [Input("powerOn")]
        public Input<bool>? PowerOn { get; set; }

        /// <summary>
        /// Shows the status code of the vApp
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Shows the status of the vApp
        /// </summary>
        [Input("statusText")]
        public Input<string>? StatusText { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappState()
        {
        }
        public static new VappState Empty => new VappState();
    }
}
