// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtNatRule
    {
        public static Task<GetNsxtNatRuleResult> InvokeAsync(GetNsxtNatRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtNatRuleResult>("vcd:index/getNsxtNatRule:getNsxtNatRule", args ?? new GetNsxtNatRuleArgs(), options.WithDefaults());

        public static Output<GetNsxtNatRuleResult> Invoke(GetNsxtNatRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtNatRuleResult>("vcd:index/getNsxtNatRule:getNsxtNatRule", args ?? new GetNsxtNatRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtNatRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGatewayId", required: true)]
        public string EdgeGatewayId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxtNatRuleArgs()
        {
        }
        public static new GetNsxtNatRuleArgs Empty => new GetNsxtNatRuleArgs();
    }

    public sealed class GetNsxtNatRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxtNatRuleInvokeArgs()
        {
        }
        public static new GetNsxtNatRuleInvokeArgs Empty => new GetNsxtNatRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtNatRuleResult
    {
        public readonly string AppPortProfileId;
        public readonly string Description;
        public readonly string DnatExternalPort;
        public readonly string EdgeGatewayId;
        public readonly bool Enabled;
        public readonly string ExternalAddress;
        public readonly string FirewallMatch;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InternalAddress;
        public readonly bool Logging;
        public readonly string Name;
        public readonly string? Org;
        public readonly int Priority;
        public readonly string RuleType;
        public readonly string SnatDestinationAddress;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxtNatRuleResult(
            string appPortProfileId,

            string description,

            string dnatExternalPort,

            string edgeGatewayId,

            bool enabled,

            string externalAddress,

            string firewallMatch,

            string id,

            string internalAddress,

            bool logging,

            string name,

            string? org,

            int priority,

            string ruleType,

            string snatDestinationAddress,

            string? vdc)
        {
            AppPortProfileId = appPortProfileId;
            Description = description;
            DnatExternalPort = dnatExternalPort;
            EdgeGatewayId = edgeGatewayId;
            Enabled = enabled;
            ExternalAddress = externalAddress;
            FirewallMatch = firewallMatch;
            Id = id;
            InternalAddress = internalAddress;
            Logging = logging;
            Name = name;
            Org = org;
            Priority = priority;
            RuleType = ruleType;
            SnatDestinationAddress = snatDestinationAddress;
            Vdc = vdc;
        }
    }
}
