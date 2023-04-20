// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtIpSet:NsxtIpSet")]
    public partial class NsxtIpSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP Set description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Edge Gateway name in which IP Set is located
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// A set of IP address, CIDR, IP range objects
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// IP Set name
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
        /// ID of VDC or VDC Group
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtIpSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtIpSet(string name, NsxtIpSetArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtIpSet:NsxtIpSet", name, args ?? new NsxtIpSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtIpSet(string name, Input<string> id, NsxtIpSetState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtIpSet:NsxtIpSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtIpSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtIpSet Get(string name, Input<string> id, NsxtIpSetState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtIpSet(name, id, state, options);
        }
    }

    public sealed class NsxtIpSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Set description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Edge Gateway name in which IP Set is located
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// A set of IP address, CIDR, IP range objects
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// IP Set name
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
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtIpSetArgs()
        {
        }
        public static new NsxtIpSetArgs Empty => new NsxtIpSetArgs();
    }

    public sealed class NsxtIpSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Set description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Edge Gateway name in which IP Set is located
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// A set of IP address, CIDR, IP range objects
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// IP Set name
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
        /// ID of VDC or VDC Group
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtIpSetState()
        {
        }
        public static new NsxtIpSetState Empty => new NsxtIpSetState();
    }
}
