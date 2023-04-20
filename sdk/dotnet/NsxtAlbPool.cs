// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtAlbPool:NsxtAlbPool")]
    public partial class NsxtAlbPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`,
        /// `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`
        /// </summary>
        [Output("algorithm")]
        public Output<string?> Algorithm { get; private set; } = null!;

        /// <summary>
        /// IDs of associated virtual services
        /// </summary>
        [Output("associatedVirtualServiceIds")]
        public Output<ImmutableArray<string>> AssociatedVirtualServiceIds { get; private set; } = null!;

        /// <summary>
        /// Names of associated virtual services
        /// </summary>
        [Output("associatedVirtualServices")]
        public Output<ImmutableArray<string>> AssociatedVirtualServices { get; private set; } = null!;

        /// <summary>
        /// A set of root certificate IDs to use when validating certificates presented by pool members
        /// </summary>
        [Output("caCertificateIds")]
        public Output<ImmutableArray<string>> CaCertificateIds { get; private set; } = null!;

        /// <summary>
        /// Boolean flag if common name check of the certificate should be enabled
        /// </summary>
        [Output("cnCheckEnabled")]
        public Output<bool?> CnCheckEnabled { get; private set; } = null!;

        /// <summary>
        /// Default Port defines destination server port used by the traffic sent to the member (default 80)
        /// </summary>
        [Output("defaultPort")]
        public Output<int?> DefaultPort { get; private set; } = null!;

        /// <summary>
        /// Description of ALB Pool
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of domain names which will be used to verify common names
        /// </summary>
        [Output("domainNames")]
        public Output<ImmutableArray<string>> DomainNames { get; private set; } = null!;

        /// <summary>
        /// Edge gateway ID in which ALB Pool should be created
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// Boolean value if ALB Pool is enabled or not (default true)
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Number of enabled members in the pool
        /// </summary>
        [Output("enabledMemberCount")]
        public Output<int> EnabledMemberCount { get; private set; } = null!;

        /// <summary>
        /// Maximum time in minutes to gracefully disable pool member (default 1)
        /// </summary>
        [Output("gracefulTimeoutPeriod")]
        public Output<int?> GracefulTimeoutPeriod { get; private set; } = null!;

        /// <summary>
        /// Health message
        /// </summary>
        [Output("healthMessage")]
        public Output<string> HealthMessage { get; private set; } = null!;

        [Output("healthMonitors")]
        public Output<ImmutableArray<Outputs.NsxtAlbPoolHealthMonitor>> HealthMonitors { get; private set; } = null!;

        /// <summary>
        /// Number of members in the pool
        /// </summary>
        [Output("memberCount")]
        public Output<int> MemberCount { get; private set; } = null!;

        /// <summary>
        /// ALB Pool Members
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.NsxtAlbPoolMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Name of ALB Pool
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
        /// Monitors if the traffic is accepted by node (default true)
        /// </summary>
        [Output("passiveMonitoringEnabled")]
        public Output<bool?> PassiveMonitoringEnabled { get; private set; } = null!;

        [Output("persistenceProfile")]
        public Output<Outputs.NsxtAlbPoolPersistenceProfile?> PersistenceProfile { get; private set; } = null!;

        /// <summary>
        /// Number of members in the pool serving the traffic
        /// </summary>
        [Output("upMemberCount")]
        public Output<int> UpMemberCount { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtAlbPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtAlbPool(string name, NsxtAlbPoolArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbPool:NsxtAlbPool", name, args ?? new NsxtAlbPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtAlbPool(string name, Input<string> id, NsxtAlbPoolState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbPool:NsxtAlbPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtAlbPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtAlbPool Get(string name, Input<string> id, NsxtAlbPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtAlbPool(name, id, state, options);
        }
    }

    public sealed class NsxtAlbPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`,
        /// `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("caCertificateIds")]
        private InputList<string>? _caCertificateIds;

        /// <summary>
        /// A set of root certificate IDs to use when validating certificates presented by pool members
        /// </summary>
        public InputList<string> CaCertificateIds
        {
            get => _caCertificateIds ?? (_caCertificateIds = new InputList<string>());
            set => _caCertificateIds = value;
        }

        /// <summary>
        /// Boolean flag if common name check of the certificate should be enabled
        /// </summary>
        [Input("cnCheckEnabled")]
        public Input<bool>? CnCheckEnabled { get; set; }

        /// <summary>
        /// Default Port defines destination server port used by the traffic sent to the member (default 80)
        /// </summary>
        [Input("defaultPort")]
        public Input<int>? DefaultPort { get; set; }

        /// <summary>
        /// Description of ALB Pool
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domainNames")]
        private InputList<string>? _domainNames;

        /// <summary>
        /// List of domain names which will be used to verify common names
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// Edge gateway ID in which ALB Pool should be created
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Boolean value if ALB Pool is enabled or not (default true)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Maximum time in minutes to gracefully disable pool member (default 1)
        /// </summary>
        [Input("gracefulTimeoutPeriod")]
        public Input<int>? GracefulTimeoutPeriod { get; set; }

        [Input("healthMonitors")]
        private InputList<Inputs.NsxtAlbPoolHealthMonitorArgs>? _healthMonitors;
        public InputList<Inputs.NsxtAlbPoolHealthMonitorArgs> HealthMonitors
        {
            get => _healthMonitors ?? (_healthMonitors = new InputList<Inputs.NsxtAlbPoolHealthMonitorArgs>());
            set => _healthMonitors = value;
        }

        [Input("members")]
        private InputList<Inputs.NsxtAlbPoolMemberArgs>? _members;

        /// <summary>
        /// ALB Pool Members
        /// </summary>
        public InputList<Inputs.NsxtAlbPoolMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.NsxtAlbPoolMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Name of ALB Pool
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
        /// Monitors if the traffic is accepted by node (default true)
        /// </summary>
        [Input("passiveMonitoringEnabled")]
        public Input<bool>? PassiveMonitoringEnabled { get; set; }

        [Input("persistenceProfile")]
        public Input<Inputs.NsxtAlbPoolPersistenceProfileArgs>? PersistenceProfile { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtAlbPoolArgs()
        {
        }
        public static new NsxtAlbPoolArgs Empty => new NsxtAlbPoolArgs();
    }

    public sealed class NsxtAlbPoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Algorithm for choosing pool members (default LEAST_CONNECTIONS). Other `ROUND_ROBIN`,`CONSISTENT_HASH`,
        /// `FASTEST_RESPONSE`, `LEAST_LOAD`, `FEWEST_SERVERS`, `RANDOM`, `FEWEST_TASKS`,`CORE_AFFINITY`
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("associatedVirtualServiceIds")]
        private InputList<string>? _associatedVirtualServiceIds;

        /// <summary>
        /// IDs of associated virtual services
        /// </summary>
        public InputList<string> AssociatedVirtualServiceIds
        {
            get => _associatedVirtualServiceIds ?? (_associatedVirtualServiceIds = new InputList<string>());
            set => _associatedVirtualServiceIds = value;
        }

        [Input("associatedVirtualServices")]
        private InputList<string>? _associatedVirtualServices;

        /// <summary>
        /// Names of associated virtual services
        /// </summary>
        public InputList<string> AssociatedVirtualServices
        {
            get => _associatedVirtualServices ?? (_associatedVirtualServices = new InputList<string>());
            set => _associatedVirtualServices = value;
        }

        [Input("caCertificateIds")]
        private InputList<string>? _caCertificateIds;

        /// <summary>
        /// A set of root certificate IDs to use when validating certificates presented by pool members
        /// </summary>
        public InputList<string> CaCertificateIds
        {
            get => _caCertificateIds ?? (_caCertificateIds = new InputList<string>());
            set => _caCertificateIds = value;
        }

        /// <summary>
        /// Boolean flag if common name check of the certificate should be enabled
        /// </summary>
        [Input("cnCheckEnabled")]
        public Input<bool>? CnCheckEnabled { get; set; }

        /// <summary>
        /// Default Port defines destination server port used by the traffic sent to the member (default 80)
        /// </summary>
        [Input("defaultPort")]
        public Input<int>? DefaultPort { get; set; }

        /// <summary>
        /// Description of ALB Pool
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domainNames")]
        private InputList<string>? _domainNames;

        /// <summary>
        /// List of domain names which will be used to verify common names
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// Edge gateway ID in which ALB Pool should be created
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// Boolean value if ALB Pool is enabled or not (default true)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Number of enabled members in the pool
        /// </summary>
        [Input("enabledMemberCount")]
        public Input<int>? EnabledMemberCount { get; set; }

        /// <summary>
        /// Maximum time in minutes to gracefully disable pool member (default 1)
        /// </summary>
        [Input("gracefulTimeoutPeriod")]
        public Input<int>? GracefulTimeoutPeriod { get; set; }

        /// <summary>
        /// Health message
        /// </summary>
        [Input("healthMessage")]
        public Input<string>? HealthMessage { get; set; }

        [Input("healthMonitors")]
        private InputList<Inputs.NsxtAlbPoolHealthMonitorGetArgs>? _healthMonitors;
        public InputList<Inputs.NsxtAlbPoolHealthMonitorGetArgs> HealthMonitors
        {
            get => _healthMonitors ?? (_healthMonitors = new InputList<Inputs.NsxtAlbPoolHealthMonitorGetArgs>());
            set => _healthMonitors = value;
        }

        /// <summary>
        /// Number of members in the pool
        /// </summary>
        [Input("memberCount")]
        public Input<int>? MemberCount { get; set; }

        [Input("members")]
        private InputList<Inputs.NsxtAlbPoolMemberGetArgs>? _members;

        /// <summary>
        /// ALB Pool Members
        /// </summary>
        public InputList<Inputs.NsxtAlbPoolMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.NsxtAlbPoolMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Name of ALB Pool
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
        /// Monitors if the traffic is accepted by node (default true)
        /// </summary>
        [Input("passiveMonitoringEnabled")]
        public Input<bool>? PassiveMonitoringEnabled { get; set; }

        [Input("persistenceProfile")]
        public Input<Inputs.NsxtAlbPoolPersistenceProfileGetArgs>? PersistenceProfile { get; set; }

        /// <summary>
        /// Number of members in the pool serving the traffic
        /// </summary>
        [Input("upMemberCount")]
        public Input<int>? UpMemberCount { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtAlbPoolState()
        {
        }
        public static new NsxtAlbPoolState Empty => new NsxtAlbPoolState();
    }
}
