// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetVmAffinityRule
    {
        public static Task<GetVmAffinityRuleResult> InvokeAsync(GetVmAffinityRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVmAffinityRuleResult>("vcd:index/getVmAffinityRule:getVmAffinityRule", args ?? new GetVmAffinityRuleArgs(), options.WithDefaults());

        public static Output<GetVmAffinityRuleResult> Invoke(GetVmAffinityRuleInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVmAffinityRuleResult>("vcd:index/getVmAffinityRule:getVmAffinityRule", args ?? new GetVmAffinityRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVmAffinityRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        [Input("ruleId")]
        public string? RuleId { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetVmAffinityRuleArgs()
        {
        }
        public static new GetVmAffinityRuleArgs Empty => new GetVmAffinityRuleArgs();
    }

    public sealed class GetVmAffinityRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetVmAffinityRuleInvokeArgs()
        {
        }
        public static new GetVmAffinityRuleInvokeArgs Empty => new GetVmAffinityRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetVmAffinityRuleResult
    {
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string Polarity;
        public readonly bool Required;
        public readonly string? RuleId;
        public readonly string? Vdc;
        public readonly ImmutableArray<string> VmIds;

        [OutputConstructor]
        private GetVmAffinityRuleResult(
            bool enabled,

            string id,

            string? name,

            string? org,

            string polarity,

            bool required,

            string? ruleId,

            string? vdc,

            ImmutableArray<string> vmIds)
        {
            Enabled = enabled;
            Id = id;
            Name = name;
            Org = org;
            Polarity = polarity;
            Required = required;
            RuleId = ruleId;
            Vdc = vdc;
            VmIds = vmIds;
        }
    }
}
