// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxvDnat:NsxvDnat")]
    public partial class NsxvDnat : global::Pulumi.CustomResource
    {
        /// <summary>
        /// NAT rule description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Edge gateway name in which NAT Rule is located
        /// </summary>
        [Output("edgeGateway")]
        public Output<string> EdgeGateway { get; private set; } = null!;

        /// <summary>
        /// Whether the rule should be enabled. Default 'true'
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
        /// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
        /// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
        /// </summary>
        [Output("icmpType")]
        public Output<string?> IcmpType { get; private set; } = null!;

        /// <summary>
        /// Whether logging should be enabled for this rule. Default 'false'
        /// </summary>
        [Output("loggingEnabled")]
        public Output<bool?> LoggingEnabled { get; private set; } = null!;

        /// <summary>
        /// Org or external network name
        /// </summary>
        [Output("networkName")]
        public Output<string> NetworkName { get; private set; } = null!;

        /// <summary>
        /// Network type. One of 'ext', 'org'
        /// </summary>
        [Output("networkType")]
        public Output<string> NetworkType { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Original address or address range. This is the the destination address for DNAT rules.
        /// </summary>
        [Output("originalAddress")]
        public Output<string> OriginalAddress { get; private set; } = null!;

        /// <summary>
        /// Original port. This is the destination port for DNAT rules
        /// </summary>
        [Output("originalPort")]
        public Output<string?> OriginalPort { get; private set; } = null!;

        /// <summary>
        /// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// Optional. Allows to set custom rule tag
        /// </summary>
        [Output("ruleTag")]
        public Output<int> RuleTag { get; private set; } = null!;

        /// <summary>
        /// Read only. Possible values 'user', 'internal_high'
        /// </summary>
        [Output("ruleType")]
        public Output<string> RuleType { get; private set; } = null!;

        /// <summary>
        /// Translated address or address range
        /// </summary>
        [Output("translatedAddress")]
        public Output<string?> TranslatedAddress { get; private set; } = null!;

        /// <summary>
        /// Translated port
        /// </summary>
        [Output("translatedPort")]
        public Output<string?> TranslatedPort { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxvDnat resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxvDnat(string name, NsxvDnatArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxvDnat:NsxvDnat", name, args ?? new NsxvDnatArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxvDnat(string name, Input<string> id, NsxvDnatState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxvDnat:NsxvDnat", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxvDnat resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxvDnat Get(string name, Input<string> id, NsxvDnatState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxvDnat(name, id, state, options);
        }
    }

    public sealed class NsxvDnatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NAT rule description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Edge gateway name in which NAT Rule is located
        /// </summary>
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        /// <summary>
        /// Whether the rule should be enabled. Default 'true'
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
        /// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
        /// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
        /// </summary>
        [Input("icmpType")]
        public Input<string>? IcmpType { get; set; }

        /// <summary>
        /// Whether logging should be enabled for this rule. Default 'false'
        /// </summary>
        [Input("loggingEnabled")]
        public Input<bool>? LoggingEnabled { get; set; }

        /// <summary>
        /// Org or external network name
        /// </summary>
        [Input("networkName", required: true)]
        public Input<string> NetworkName { get; set; } = null!;

        /// <summary>
        /// Network type. One of 'ext', 'org'
        /// </summary>
        [Input("networkType", required: true)]
        public Input<string> NetworkType { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Original address or address range. This is the the destination address for DNAT rules.
        /// </summary>
        [Input("originalAddress", required: true)]
        public Input<string> OriginalAddress { get; set; } = null!;

        /// <summary>
        /// Original port. This is the destination port for DNAT rules
        /// </summary>
        [Input("originalPort")]
        public Input<string>? OriginalPort { get; set; }

        /// <summary>
        /// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Optional. Allows to set custom rule tag
        /// </summary>
        [Input("ruleTag")]
        public Input<int>? RuleTag { get; set; }

        /// <summary>
        /// Read only. Possible values 'user', 'internal_high'
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        /// <summary>
        /// Translated address or address range
        /// </summary>
        [Input("translatedAddress")]
        public Input<string>? TranslatedAddress { get; set; }

        /// <summary>
        /// Translated port
        /// </summary>
        [Input("translatedPort")]
        public Input<string>? TranslatedPort { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxvDnatArgs()
        {
        }
        public static new NsxvDnatArgs Empty => new NsxvDnatArgs();
    }

    public sealed class NsxvDnatState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NAT rule description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Edge gateway name in which NAT Rule is located
        /// </summary>
        [Input("edgeGateway")]
        public Input<string>? EdgeGateway { get; set; }

        /// <summary>
        /// Whether the rule should be enabled. Default 'true'
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// ICMP type. Only supported when protocol is ICMP. One of `any`, `address-mask-request`, `address-mask-reply`,
        /// `destination-unreachable`, `echo-request`, `echo-reply`, `parameter-problem`, `redirect`, `router-advertisement`,
        /// `router-solicitation`, `source-quench`, `time-exceeded`, `timestamp-request`, `timestamp-reply`. Default `any`
        /// </summary>
        [Input("icmpType")]
        public Input<string>? IcmpType { get; set; }

        /// <summary>
        /// Whether logging should be enabled for this rule. Default 'false'
        /// </summary>
        [Input("loggingEnabled")]
        public Input<bool>? LoggingEnabled { get; set; }

        /// <summary>
        /// Org or external network name
        /// </summary>
        [Input("networkName")]
        public Input<string>? NetworkName { get; set; }

        /// <summary>
        /// Network type. One of 'ext', 'org'
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Original address or address range. This is the the destination address for DNAT rules.
        /// </summary>
        [Input("originalAddress")]
        public Input<string>? OriginalAddress { get; set; }

        /// <summary>
        /// Original port. This is the destination port for DNAT rules
        /// </summary>
        [Input("originalPort")]
        public Input<string>? OriginalPort { get; set; }

        /// <summary>
        /// Protocol. Such as 'tcp', 'udp', 'icmp', 'any'
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Optional. Allows to set custom rule tag
        /// </summary>
        [Input("ruleTag")]
        public Input<int>? RuleTag { get; set; }

        /// <summary>
        /// Read only. Possible values 'user', 'internal_high'
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        /// <summary>
        /// Translated address or address range
        /// </summary>
        [Input("translatedAddress")]
        public Input<string>? TranslatedAddress { get; set; }

        /// <summary>
        /// Translated port
        /// </summary>
        [Input("translatedPort")]
        public Input<string>? TranslatedPort { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxvDnatState()
        {
        }
        public static new NsxvDnatState Empty => new NsxvDnatState();
    }
}
