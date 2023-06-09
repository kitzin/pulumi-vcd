// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/vappNetwork:VappNetwork")]
    public partial class VappNetwork : global::Pulumi.CustomResource
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
        public Output<ImmutableArray<Outputs.VappNetworkDhcpPool>> DhcpPools { get; private set; } = null!;

        /// <summary>
        /// Primary DNS server
        /// </summary>
        [Output("dns1")]
        public Output<string?> Dns1 { get; private set; } = null!;

        /// <summary>
        /// Secondary DNS server
        /// </summary>
        [Output("dns2")]
        public Output<string?> Dns2 { get; private set; } = null!;

        /// <summary>
        /// DNS suffix
        /// </summary>
        [Output("dnsSuffix")]
        public Output<string?> DnsSuffix { get; private set; } = null!;

        /// <summary>
        /// Gateway of the network
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// True if Network allows guest VLAN tagging
        /// </summary>
        [Output("guestVlanAllowed")]
        public Output<bool?> GuestVlanAllowed { get; private set; } = null!;

        /// <summary>
        /// vApp network name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Netmask address for a subnet. Default is 255.255.255.0
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
        /// org network name to which vapp network is connected
        /// </summary>
        [Output("orgNetworkName")]
        public Output<string?> OrgNetworkName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
        /// </summary>
        [Output("retainIpMacEnabled")]
        public Output<bool?> RetainIpMacEnabled { get; private set; } = null!;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for virtual machines
        /// </summary>
        [Output("staticIpPools")]
        public Output<ImmutableArray<Outputs.VappNetworkStaticIpPool>> StaticIpPools { get; private set; } = null!;

        /// <summary>
        /// vApp to use
        /// </summary>
        [Output("vappName")]
        public Output<string> VappName { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a VappNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VappNetwork(string name, VappNetworkArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/vappNetwork:VappNetwork", name, args ?? new VappNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VappNetwork(string name, Input<string> id, VappNetworkState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vappNetwork:VappNetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VappNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VappNetwork Get(string name, Input<string> id, VappNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new VappNetwork(name, id, state, options);
        }
    }

    public sealed class VappNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description for the network
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dhcpPools")]
        private InputList<Inputs.VappNetworkDhcpPoolArgs>? _dhcpPools;

        /// <summary>
        /// A range of IPs to issue to virtual machines that don't have a static IP
        /// </summary>
        public InputList<Inputs.VappNetworkDhcpPoolArgs> DhcpPools
        {
            get => _dhcpPools ?? (_dhcpPools = new InputList<Inputs.VappNetworkDhcpPoolArgs>());
            set => _dhcpPools = value;
        }

        /// <summary>
        /// Primary DNS server
        /// </summary>
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        /// <summary>
        /// Secondary DNS server
        /// </summary>
        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        /// <summary>
        /// DNS suffix
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// Gateway of the network
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        /// <summary>
        /// True if Network allows guest VLAN tagging
        /// </summary>
        [Input("guestVlanAllowed")]
        public Input<bool>? GuestVlanAllowed { get; set; }

        /// <summary>
        /// vApp network name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Netmask address for a subnet. Default is 255.255.255.0
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
        /// org network name to which vapp network is connected
        /// </summary>
        [Input("orgNetworkName")]
        public Input<string>? OrgNetworkName { get; set; }

        /// <summary>
        /// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
        /// </summary>
        [Input("retainIpMacEnabled")]
        public Input<bool>? RetainIpMacEnabled { get; set; }

        [Input("staticIpPools")]
        private InputList<Inputs.VappNetworkStaticIpPoolArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for virtual machines
        /// </summary>
        public InputList<Inputs.VappNetworkStaticIpPoolArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.VappNetworkStaticIpPoolArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// vApp to use
        /// </summary>
        [Input("vappName", required: true)]
        public Input<string> VappName { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappNetworkArgs()
        {
        }
        public static new VappNetworkArgs Empty => new VappNetworkArgs();
    }

    public sealed class VappNetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional description for the network
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dhcpPools")]
        private InputList<Inputs.VappNetworkDhcpPoolGetArgs>? _dhcpPools;

        /// <summary>
        /// A range of IPs to issue to virtual machines that don't have a static IP
        /// </summary>
        public InputList<Inputs.VappNetworkDhcpPoolGetArgs> DhcpPools
        {
            get => _dhcpPools ?? (_dhcpPools = new InputList<Inputs.VappNetworkDhcpPoolGetArgs>());
            set => _dhcpPools = value;
        }

        /// <summary>
        /// Primary DNS server
        /// </summary>
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        /// <summary>
        /// Secondary DNS server
        /// </summary>
        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        /// <summary>
        /// DNS suffix
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// Gateway of the network
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// True if Network allows guest VLAN tagging
        /// </summary>
        [Input("guestVlanAllowed")]
        public Input<bool>? GuestVlanAllowed { get; set; }

        /// <summary>
        /// vApp network name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Netmask address for a subnet. Default is 255.255.255.0
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
        /// org network name to which vapp network is connected
        /// </summary>
        [Input("orgNetworkName")]
        public Input<string>? OrgNetworkName { get; set; }

        /// <summary>
        /// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
        /// </summary>
        [Input("retainIpMacEnabled")]
        public Input<bool>? RetainIpMacEnabled { get; set; }

        [Input("staticIpPools")]
        private InputList<Inputs.VappNetworkStaticIpPoolGetArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for virtual machines
        /// </summary>
        public InputList<Inputs.VappNetworkStaticIpPoolGetArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.VappNetworkStaticIpPoolGetArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// vApp to use
        /// </summary>
        [Input("vappName")]
        public Input<string>? VappName { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappNetworkState()
        {
        }
        public static new VappNetworkState Empty => new VappNetworkState();
    }
}
