// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/externalNetworkV2:ExternalNetworkV2")]
    public partial class ExternalNetworkV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Network description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A set of IP scopes for the network
        /// </summary>
        [Output("ipScopes")]
        public Output<ImmutableArray<Outputs.ExternalNetworkV2IpScope>> IpScopes { get; private set; } = null!;

        /// <summary>
        /// Network name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Reference to NSX-T Tier-0 router or segment and manager
        /// </summary>
        [Output("nsxtNetwork")]
        public Output<Outputs.ExternalNetworkV2NsxtNetwork?> NsxtNetwork { get; private set; } = null!;

        /// <summary>
        /// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
        /// registered with the system.
        /// </summary>
        [Output("vsphereNetworks")]
        public Output<ImmutableArray<Outputs.ExternalNetworkV2VsphereNetwork>> VsphereNetworks { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalNetworkV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalNetworkV2(string name, ExternalNetworkV2Args args, CustomResourceOptions? options = null)
            : base("vcd:index/externalNetworkV2:ExternalNetworkV2", name, args ?? new ExternalNetworkV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalNetworkV2(string name, Input<string> id, ExternalNetworkV2State? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/externalNetworkV2:ExternalNetworkV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExternalNetworkV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalNetworkV2 Get(string name, Input<string> id, ExternalNetworkV2State? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalNetworkV2(name, id, state, options);
        }
    }

    public sealed class ExternalNetworkV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ipScopes", required: true)]
        private InputList<Inputs.ExternalNetworkV2IpScopeArgs>? _ipScopes;

        /// <summary>
        /// A set of IP scopes for the network
        /// </summary>
        public InputList<Inputs.ExternalNetworkV2IpScopeArgs> IpScopes
        {
            get => _ipScopes ?? (_ipScopes = new InputList<Inputs.ExternalNetworkV2IpScopeArgs>());
            set => _ipScopes = value;
        }

        /// <summary>
        /// Network name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Reference to NSX-T Tier-0 router or segment and manager
        /// </summary>
        [Input("nsxtNetwork")]
        public Input<Inputs.ExternalNetworkV2NsxtNetworkArgs>? NsxtNetwork { get; set; }

        [Input("vsphereNetworks")]
        private InputList<Inputs.ExternalNetworkV2VsphereNetworkArgs>? _vsphereNetworks;

        /// <summary>
        /// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
        /// registered with the system.
        /// </summary>
        public InputList<Inputs.ExternalNetworkV2VsphereNetworkArgs> VsphereNetworks
        {
            get => _vsphereNetworks ?? (_vsphereNetworks = new InputList<Inputs.ExternalNetworkV2VsphereNetworkArgs>());
            set => _vsphereNetworks = value;
        }

        public ExternalNetworkV2Args()
        {
        }
        public static new ExternalNetworkV2Args Empty => new ExternalNetworkV2Args();
    }

    public sealed class ExternalNetworkV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ipScopes")]
        private InputList<Inputs.ExternalNetworkV2IpScopeGetArgs>? _ipScopes;

        /// <summary>
        /// A set of IP scopes for the network
        /// </summary>
        public InputList<Inputs.ExternalNetworkV2IpScopeGetArgs> IpScopes
        {
            get => _ipScopes ?? (_ipScopes = new InputList<Inputs.ExternalNetworkV2IpScopeGetArgs>());
            set => _ipScopes = value;
        }

        /// <summary>
        /// Network name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Reference to NSX-T Tier-0 router or segment and manager
        /// </summary>
        [Input("nsxtNetwork")]
        public Input<Inputs.ExternalNetworkV2NsxtNetworkGetArgs>? NsxtNetwork { get; set; }

        [Input("vsphereNetworks")]
        private InputList<Inputs.ExternalNetworkV2VsphereNetworkGetArgs>? _vsphereNetworks;

        /// <summary>
        /// A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
        /// registered with the system.
        /// </summary>
        public InputList<Inputs.ExternalNetworkV2VsphereNetworkGetArgs> VsphereNetworks
        {
            get => _vsphereNetworks ?? (_vsphereNetworks = new InputList<Inputs.ExternalNetworkV2VsphereNetworkGetArgs>());
            set => _vsphereNetworks = value;
        }

        public ExternalNetworkV2State()
        {
        }
        public static new ExternalNetworkV2State Empty => new ExternalNetworkV2State();
    }
}
