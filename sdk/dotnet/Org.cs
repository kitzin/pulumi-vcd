// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/org:Org")]
    public partial class Org : global::Pulumi.CustomResource
    {
        /// <summary>
        /// True if this organization is allowed to share catalogs.
        /// </summary>
        [Output("canPublishCatalogs")]
        public Output<bool?> CanPublishCatalogs { get; private set; } = null!;

        /// <summary>
        /// True if this organization is allowed to publish external catalogs.
        /// </summary>
        [Output("canPublishExternalCatalogs")]
        public Output<bool?> CanPublishExternalCatalogs { get; private set; } = null!;

        /// <summary>
        /// True if this organization is allowed to subscribe to external catalogs.
        /// </summary>
        [Output("canSubscribeExternalCatalogs")]
        public Output<bool?> CanSubscribeExternalCatalogs { get; private set; } = null!;

        /// <summary>
        /// Specifies this organization's default for virtual machine boot delay after power on.
        /// </summary>
        [Output("delayAfterPowerOnSeconds")]
        public Output<int?> DelayAfterPowerOnSeconds { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Output("deleteForce")]
        public Output<bool> DeleteForce { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Output("deleteRecursive")]
        public Output<bool> DeleteRecursive { get; private set; } = null!;

        /// <summary>
        /// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited)
        /// </summary>
        [Output("deployedVmQuota")]
        public Output<int?> DeployedVmQuota { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("fullName")]
        public Output<string> FullName { get; private set; } = null!;

        /// <summary>
        /// True if this organization is enabled (allows login and all other operations).
        /// </summary>
        [Output("isEnabled")]
        public Output<bool?> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// Key value map of metadata to assign to this organization. Key and value can be any string.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Metadata entries for the given Organization
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.OrgMetadataEntry>> MetadataEntries { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of
        /// this organization. (0 = unlimited)
        /// </summary>
        [Output("storedVmQuota")]
        public Output<int?> StoredVmQuota { get; private set; } = null!;

        /// <summary>
        /// Defines lease parameters for vApps created in this organization
        /// </summary>
        [Output("vappLease")]
        public Output<Outputs.OrgVappLease> VappLease { get; private set; } = null!;

        /// <summary>
        /// Defines lease parameters for vApp templates created in this organization
        /// </summary>
        [Output("vappTemplateLease")]
        public Output<Outputs.OrgVappTemplateLease> VappTemplateLease { get; private set; } = null!;


        /// <summary>
        /// Create a Org resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Org(string name, OrgArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/org:Org", name, args ?? new OrgArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Org(string name, Input<string> id, OrgState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/org:Org", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Org resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Org Get(string name, Input<string> id, OrgState? state = null, CustomResourceOptions? options = null)
        {
            return new Org(name, id, state, options);
        }
    }

    public sealed class OrgArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if this organization is allowed to share catalogs.
        /// </summary>
        [Input("canPublishCatalogs")]
        public Input<bool>? CanPublishCatalogs { get; set; }

        /// <summary>
        /// True if this organization is allowed to publish external catalogs.
        /// </summary>
        [Input("canPublishExternalCatalogs")]
        public Input<bool>? CanPublishExternalCatalogs { get; set; }

        /// <summary>
        /// True if this organization is allowed to subscribe to external catalogs.
        /// </summary>
        [Input("canSubscribeExternalCatalogs")]
        public Input<bool>? CanSubscribeExternalCatalogs { get; set; }

        /// <summary>
        /// Specifies this organization's default for virtual machine boot delay after power on.
        /// </summary>
        [Input("delayAfterPowerOnSeconds")]
        public Input<int>? DelayAfterPowerOnSeconds { get; set; }

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Input("deleteForce", required: true)]
        public Input<bool> DeleteForce { get; set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive", required: true)]
        public Input<bool> DeleteRecursive { get; set; } = null!;

        /// <summary>
        /// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited)
        /// </summary>
        [Input("deployedVmQuota")]
        public Input<int>? DeployedVmQuota { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fullName", required: true)]
        public Input<string> FullName { get; set; } = null!;

        /// <summary>
        /// True if this organization is enabled (allows login and all other operations).
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key value map of metadata to assign to this organization. Key and value can be any string.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.OrgMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given Organization
        /// </summary>
        public InputList<Inputs.OrgMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.OrgMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of
        /// this organization. (0 = unlimited)
        /// </summary>
        [Input("storedVmQuota")]
        public Input<int>? StoredVmQuota { get; set; }

        /// <summary>
        /// Defines lease parameters for vApps created in this organization
        /// </summary>
        [Input("vappLease")]
        public Input<Inputs.OrgVappLeaseArgs>? VappLease { get; set; }

        /// <summary>
        /// Defines lease parameters for vApp templates created in this organization
        /// </summary>
        [Input("vappTemplateLease")]
        public Input<Inputs.OrgVappTemplateLeaseArgs>? VappTemplateLease { get; set; }

        public OrgArgs()
        {
        }
        public static new OrgArgs Empty => new OrgArgs();
    }

    public sealed class OrgState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if this organization is allowed to share catalogs.
        /// </summary>
        [Input("canPublishCatalogs")]
        public Input<bool>? CanPublishCatalogs { get; set; }

        /// <summary>
        /// True if this organization is allowed to publish external catalogs.
        /// </summary>
        [Input("canPublishExternalCatalogs")]
        public Input<bool>? CanPublishExternalCatalogs { get; set; }

        /// <summary>
        /// True if this organization is allowed to subscribe to external catalogs.
        /// </summary>
        [Input("canSubscribeExternalCatalogs")]
        public Input<bool>? CanSubscribeExternalCatalogs { get; set; }

        /// <summary>
        /// Specifies this organization's default for virtual machine boot delay after power on.
        /// </summary>
        [Input("delayAfterPowerOnSeconds")]
        public Input<int>? DelayAfterPowerOnSeconds { get; set; }

        /// <summary>
        /// When destroying use delete_force=True with delete_recursive=True to remove an org and any objects it contains,
        /// regardless of their state.
        /// </summary>
        [Input("deleteForce")]
        public Input<bool>? DeleteForce { get; set; }

        /// <summary>
        /// When destroying use delete_recursive=True to remove the org and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive")]
        public Input<bool>? DeleteRecursive { get; set; }

        /// <summary>
        /// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. (0 = unlimited)
        /// </summary>
        [Input("deployedVmQuota")]
        public Input<int>? DeployedVmQuota { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// True if this organization is enabled (allows login and all other operations).
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key value map of metadata to assign to this organization. Key and value can be any string.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.OrgMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given Organization
        /// </summary>
        public InputList<Inputs.OrgMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.OrgMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of
        /// this organization. (0 = unlimited)
        /// </summary>
        [Input("storedVmQuota")]
        public Input<int>? StoredVmQuota { get; set; }

        /// <summary>
        /// Defines lease parameters for vApps created in this organization
        /// </summary>
        [Input("vappLease")]
        public Input<Inputs.OrgVappLeaseGetArgs>? VappLease { get; set; }

        /// <summary>
        /// Defines lease parameters for vApp templates created in this organization
        /// </summary>
        [Input("vappTemplateLease")]
        public Input<Inputs.OrgVappTemplateLeaseGetArgs>? VappTemplateLease { get; set; }

        public OrgState()
        {
        }
        public static new OrgState Empty => new OrgState();
    }
}