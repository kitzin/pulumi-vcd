// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxvFirewallRule
    {
        public static Task<GetNsxvFirewallRuleResult> InvokeAsync(GetNsxvFirewallRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxvFirewallRuleResult>("vcd:index/getNsxvFirewallRule:getNsxvFirewallRule", args ?? new GetNsxvFirewallRuleArgs(), options.WithDefaults());

        public static Output<GetNsxvFirewallRuleResult> Invoke(GetNsxvFirewallRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxvFirewallRuleResult>("vcd:index/getNsxvFirewallRule:getNsxvFirewallRule", args ?? new GetNsxvFirewallRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxvFirewallRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public string EdgeGateway { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("ruleId", required: true)]
        public string RuleId { get; set; } = null!;

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxvFirewallRuleArgs()
        {
        }
        public static new GetNsxvFirewallRuleArgs Empty => new GetNsxvFirewallRuleArgs();
    }

    public sealed class GetNsxvFirewallRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxvFirewallRuleInvokeArgs()
        {
        }
        public static new GetNsxvFirewallRuleInvokeArgs Empty => new GetNsxvFirewallRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxvFirewallRuleResult
    {
        public readonly string Action;
        public readonly ImmutableArray<Outputs.GetNsxvFirewallRuleDestinationResult> Destinations;
        public readonly string EdgeGateway;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool LoggingEnabled;
        public readonly string Name;
        public readonly string? Org;
        public readonly string RuleId;
        public readonly int RuleTag;
        public readonly string RuleType;
        public readonly ImmutableArray<Outputs.GetNsxvFirewallRuleServiceResult> Services;
        public readonly ImmutableArray<Outputs.GetNsxvFirewallRuleSourceResult> Sources;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxvFirewallRuleResult(
            string action,

            ImmutableArray<Outputs.GetNsxvFirewallRuleDestinationResult> destinations,

            string edgeGateway,

            bool enabled,

            string id,

            bool loggingEnabled,

            string name,

            string? org,

            string ruleId,

            int ruleTag,

            string ruleType,

            ImmutableArray<Outputs.GetNsxvFirewallRuleServiceResult> services,

            ImmutableArray<Outputs.GetNsxvFirewallRuleSourceResult> sources,

            string? vdc)
        {
            Action = action;
            Destinations = destinations;
            EdgeGateway = edgeGateway;
            Enabled = enabled;
            Id = id;
            LoggingEnabled = loggingEnabled;
            Name = name;
            Org = org;
            RuleId = ruleId;
            RuleTag = ruleTag;
            RuleType = ruleType;
            Services = services;
            Sources = sources;
            Vdc = vdc;
        }
    }
}
