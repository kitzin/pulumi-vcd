// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtDistributedFirewall
    {
        public static Task<GetNsxtDistributedFirewallResult> InvokeAsync(GetNsxtDistributedFirewallArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtDistributedFirewallResult>("vcd:index/getNsxtDistributedFirewall:getNsxtDistributedFirewall", args ?? new GetNsxtDistributedFirewallArgs(), options.WithDefaults());

        public static Output<GetNsxtDistributedFirewallResult> Invoke(GetNsxtDistributedFirewallInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtDistributedFirewallResult>("vcd:index/getNsxtDistributedFirewall:getNsxtDistributedFirewall", args ?? new GetNsxtDistributedFirewallInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtDistributedFirewallArgs : global::Pulumi.InvokeArgs
    {
        [Input("org")]
        public string? Org { get; set; }

        [Input("vdcGroupId", required: true)]
        public string VdcGroupId { get; set; } = null!;

        public GetNsxtDistributedFirewallArgs()
        {
        }
        public static new GetNsxtDistributedFirewallArgs Empty => new GetNsxtDistributedFirewallArgs();
    }

    public sealed class GetNsxtDistributedFirewallInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdcGroupId", required: true)]
        public Input<string> VdcGroupId { get; set; } = null!;

        public GetNsxtDistributedFirewallInvokeArgs()
        {
        }
        public static new GetNsxtDistributedFirewallInvokeArgs Empty => new GetNsxtDistributedFirewallInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtDistributedFirewallResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Org;
        public readonly ImmutableArray<Outputs.GetNsxtDistributedFirewallRuleResult> Rules;
        public readonly string VdcGroupId;

        [OutputConstructor]
        private GetNsxtDistributedFirewallResult(
            string id,

            string? org,

            ImmutableArray<Outputs.GetNsxtDistributedFirewallRuleResult> rules,

            string vdcGroupId)
        {
            Id = id;
            Org = org;
            Rules = rules;
            VdcGroupId = vdcGroupId;
        }
    }
}
