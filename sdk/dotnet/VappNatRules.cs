// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/vappNatRules:VappNatRules")]
    public partial class VappNatRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
        /// </summary>
        [Output("enableIpMasquerade")]
        public Output<bool?> EnableIpMasquerade { get; private set; } = null!;

        /// <summary>
        /// Enable or disable NAT service. Default is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
        /// </summary>
        [Output("natType")]
        public Output<string> NatType { get; private set; } = null!;

        /// <summary>
        /// vApp network identifier
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        [Output("rules")]
        public Output<ImmutableArray<Outputs.VappNatRulesRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// vApp identifier
        /// </summary>
        [Output("vappId")]
        public Output<string> VappId { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a VappNatRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VappNatRules(string name, VappNatRulesArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/vappNatRules:VappNatRules", name, args ?? new VappNatRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VappNatRules(string name, Input<string> id, VappNatRulesState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vappNatRules:VappNatRules", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VappNatRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VappNatRules Get(string name, Input<string> id, VappNatRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new VappNatRules(name, id, state, options);
        }
    }

    public sealed class VappNatRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
        /// </summary>
        [Input("enableIpMasquerade")]
        public Input<bool>? EnableIpMasquerade { get; set; }

        /// <summary>
        /// Enable or disable NAT service. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
        /// </summary>
        [Input("natType", required: true)]
        public Input<string> NatType { get; set; } = null!;

        /// <summary>
        /// vApp network identifier
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("rules")]
        private InputList<Inputs.VappNatRulesRuleArgs>? _rules;
        public InputList<Inputs.VappNatRulesRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.VappNatRulesRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// vApp identifier
        /// </summary>
        [Input("vappId", required: true)]
        public Input<string> VappId { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappNatRulesArgs()
        {
        }
        public static new VappNatRulesArgs Empty => new VappNatRulesArgs();
    }

    public sealed class VappNatRulesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
        /// </summary>
        [Input("enableIpMasquerade")]
        public Input<bool>? EnableIpMasquerade { get; set; }

        /// <summary>
        /// Enable or disable NAT service. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
        /// </summary>
        [Input("natType")]
        public Input<string>? NatType { get; set; }

        /// <summary>
        /// vApp network identifier
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("rules")]
        private InputList<Inputs.VappNatRulesRuleGetArgs>? _rules;
        public InputList<Inputs.VappNatRulesRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.VappNatRulesRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// vApp identifier
        /// </summary>
        [Input("vappId")]
        public Input<string>? VappId { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappNatRulesState()
        {
        }
        public static new VappNatRulesState Empty => new VappNatRulesState();
    }
}