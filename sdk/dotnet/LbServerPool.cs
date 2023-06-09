// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/lbServerPool:LbServerPool")]
    public partial class LbServerPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Balancing method for the service. One of 'ip-hash', 'round-robin', 'uri', 'leastconn', 'url', or 'httpheader'
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Additional options for load balancing algorithm for httpheader or url algorithms
        /// </summary>
        [Output("algorithmParameters")]
        public Output<string?> AlgorithmParameters { get; private set; } = null!;

        /// <summary>
        /// Server pool description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Edge gateway name in which the LB Server Pool is located
        /// </summary>
        [Output("edgeGateway")]
        public Output<string> EdgeGateway { get; private set; } = null!;

        /// <summary>
        /// Makes client IP addresses visible to the backend servers
        /// </summary>
        [Output("enableTransparency")]
        public Output<bool?> EnableTransparency { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<Outputs.LbServerPoolMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Load Balancer Service Monitor ID
        /// </summary>
        [Output("monitorId")]
        public Output<string?> MonitorId { get; private set; } = null!;

        /// <summary>
        /// Unique LB Server Pool name
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
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a LbServerPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbServerPool(string name, LbServerPoolArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/lbServerPool:LbServerPool", name, args ?? new LbServerPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbServerPool(string name, Input<string> id, LbServerPoolState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/lbServerPool:LbServerPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LbServerPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbServerPool Get(string name, Input<string> id, LbServerPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new LbServerPool(name, id, state, options);
        }
    }

    public sealed class LbServerPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Balancing method for the service. One of 'ip-hash', 'round-robin', 'uri', 'leastconn', 'url', or 'httpheader'
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// Additional options for load balancing algorithm for httpheader or url algorithms
        /// </summary>
        [Input("algorithmParameters")]
        public Input<string>? AlgorithmParameters { get; set; }

        /// <summary>
        /// Server pool description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Edge gateway name in which the LB Server Pool is located
        /// </summary>
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        /// <summary>
        /// Makes client IP addresses visible to the backend servers
        /// </summary>
        [Input("enableTransparency")]
        public Input<bool>? EnableTransparency { get; set; }

        [Input("members")]
        private InputList<Inputs.LbServerPoolMemberArgs>? _members;
        public InputList<Inputs.LbServerPoolMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.LbServerPoolMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Load Balancer Service Monitor ID
        /// </summary>
        [Input("monitorId")]
        public Input<string>? MonitorId { get; set; }

        /// <summary>
        /// Unique LB Server Pool name
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

        public LbServerPoolArgs()
        {
        }
        public static new LbServerPoolArgs Empty => new LbServerPoolArgs();
    }

    public sealed class LbServerPoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Balancing method for the service. One of 'ip-hash', 'round-robin', 'uri', 'leastconn', 'url', or 'httpheader'
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Additional options for load balancing algorithm for httpheader or url algorithms
        /// </summary>
        [Input("algorithmParameters")]
        public Input<string>? AlgorithmParameters { get; set; }

        /// <summary>
        /// Server pool description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Edge gateway name in which the LB Server Pool is located
        /// </summary>
        [Input("edgeGateway")]
        public Input<string>? EdgeGateway { get; set; }

        /// <summary>
        /// Makes client IP addresses visible to the backend servers
        /// </summary>
        [Input("enableTransparency")]
        public Input<bool>? EnableTransparency { get; set; }

        [Input("members")]
        private InputList<Inputs.LbServerPoolMemberGetArgs>? _members;
        public InputList<Inputs.LbServerPoolMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.LbServerPoolMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Load Balancer Service Monitor ID
        /// </summary>
        [Input("monitorId")]
        public Input<string>? MonitorId { get; set; }

        /// <summary>
        /// Unique LB Server Pool name
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

        public LbServerPoolState()
        {
        }
        public static new LbServerPoolState Empty => new LbServerPoolState();
    }
}
