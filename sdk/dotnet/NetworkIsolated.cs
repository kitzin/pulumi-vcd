// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/networkIsolated:NetworkIsolated")]
    public partial class NetworkIsolated : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional description for the network
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A range of IPs to issue to virtual machines that don't have a static IP
        /// </summary>
        [Output("dhcpPools")]
        public Output<ImmutableArray<Outputs.NetworkIsolatedDhcpPool>> DhcpPools { get; private set; } = null!;

        /// <summary>
        /// First DNS server to use
        /// </summary>
        [Output("dns1")]
        public Output<string?> Dns1 { get; private set; } = null!;

        /// <summary>
        /// Second DNS server to use
        /// </summary>
        [Output("dns2")]
        public Output<string?> Dns2 { get; private set; } = null!;

        /// <summary>
        /// A FQDN for the virtual machines on this network
        /// </summary>
        [Output("dnsSuffix")]
        public Output<string?> DnsSuffix { get; private set; } = null!;

        /// <summary>
        /// The gateway for this network
        /// </summary>
        [Output("gateway")]
        public Output<string?> Gateway { get; private set; } = null!;

        /// <summary>
        /// Network Hyper Reference
        /// </summary>
        [Output("href")]
        public Output<string> Href { get; private set; } = null!;

        /// <summary>
        /// Key value map of metadata to assign to this network. Key and value can be any string
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Metadata entries for the given Network
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.NetworkIsolatedMetadataEntry>> MetadataEntries { get; private set; } = null!;

        /// <summary>
        /// A unique name for this network
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The netmask for the new network
        /// </summary>
        [Output("netmask")]
        public Output<string?> Netmask { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Defines if this network is shared between multiple VDCs in the Org
        /// </summary>
        [Output("shared")]
        public Output<bool?> Shared { get; private set; } = null!;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for virtual machines
        /// </summary>
        [Output("staticIpPools")]
        public Output<ImmutableArray<Outputs.NetworkIsolatedStaticIpPool>> StaticIpPools { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkIsolated resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkIsolated(string name, NetworkIsolatedArgs? args = null, CustomResourceOptions? options = null)
            : base("vcd:index/networkIsolated:NetworkIsolated", name, args ?? new NetworkIsolatedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkIsolated(string name, Input<string> id, NetworkIsolatedState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/networkIsolated:NetworkIsolated", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkIsolated resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkIsolated Get(string name, Input<string> id, NetworkIsolatedState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkIsolated(name, id, state, options);
        }
    }

    public sealed class NetworkIsolatedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description for the network
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dhcpPools")]
        private InputList<Inputs.NetworkIsolatedDhcpPoolArgs>? _dhcpPools;

        /// <summary>
        /// A range of IPs to issue to virtual machines that don't have a static IP
        /// </summary>
        public InputList<Inputs.NetworkIsolatedDhcpPoolArgs> DhcpPools
        {
            get => _dhcpPools ?? (_dhcpPools = new InputList<Inputs.NetworkIsolatedDhcpPoolArgs>());
            set => _dhcpPools = value;
        }

        /// <summary>
        /// First DNS server to use
        /// </summary>
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        /// <summary>
        /// Second DNS server to use
        /// </summary>
        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        /// <summary>
        /// A FQDN for the virtual machines on this network
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// The gateway for this network
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key value map of metadata to assign to this network. Key and value can be any string
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.NetworkIsolatedMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given Network
        /// </summary>
        public InputList<Inputs.NetworkIsolatedMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.NetworkIsolatedMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// A unique name for this network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The netmask for the new network
        /// </summary>
        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Defines if this network is shared between multiple VDCs in the Org
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("staticIpPools")]
        private InputList<Inputs.NetworkIsolatedStaticIpPoolArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for virtual machines
        /// </summary>
        public InputList<Inputs.NetworkIsolatedStaticIpPoolArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.NetworkIsolatedStaticIpPoolArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NetworkIsolatedArgs()
        {
        }
        public static new NetworkIsolatedArgs Empty => new NetworkIsolatedArgs();
    }

    public sealed class NetworkIsolatedState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description for the network
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dhcpPools")]
        private InputList<Inputs.NetworkIsolatedDhcpPoolGetArgs>? _dhcpPools;

        /// <summary>
        /// A range of IPs to issue to virtual machines that don't have a static IP
        /// </summary>
        public InputList<Inputs.NetworkIsolatedDhcpPoolGetArgs> DhcpPools
        {
            get => _dhcpPools ?? (_dhcpPools = new InputList<Inputs.NetworkIsolatedDhcpPoolGetArgs>());
            set => _dhcpPools = value;
        }

        /// <summary>
        /// First DNS server to use
        /// </summary>
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        /// <summary>
        /// Second DNS server to use
        /// </summary>
        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        /// <summary>
        /// A FQDN for the virtual machines on this network
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// The gateway for this network
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Network Hyper Reference
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key value map of metadata to assign to this network. Key and value can be any string
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.NetworkIsolatedMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given Network
        /// </summary>
        public InputList<Inputs.NetworkIsolatedMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.NetworkIsolatedMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// A unique name for this network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The netmask for the new network
        /// </summary>
        [Input("netmask")]
        public Input<string>? Netmask { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Defines if this network is shared between multiple VDCs in the Org
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("staticIpPools")]
        private InputList<Inputs.NetworkIsolatedStaticIpPoolGetArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for virtual machines
        /// </summary>
        public InputList<Inputs.NetworkIsolatedStaticIpPoolGetArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.NetworkIsolatedStaticIpPoolGetArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NetworkIsolatedState()
        {
        }
        public static new NetworkIsolatedState Empty => new NetworkIsolatedState();
    }
}