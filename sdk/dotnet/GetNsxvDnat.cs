// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxvDnat
    {
        public static Task<GetNsxvDnatResult> InvokeAsync(GetNsxvDnatArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxvDnatResult>("vcd:index/getNsxvDnat:getNsxvDnat", args ?? new GetNsxvDnatArgs(), options.WithDefaults());

        public static Output<GetNsxvDnatResult> Invoke(GetNsxvDnatInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxvDnatResult>("vcd:index/getNsxvDnat:getNsxvDnat", args ?? new GetNsxvDnatInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxvDnatArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public string EdgeGateway { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("ruleId", required: true)]
        public string RuleId { get; set; } = null!;

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxvDnatArgs()
        {
        }
        public static new GetNsxvDnatArgs Empty => new GetNsxvDnatArgs();
    }

    public sealed class GetNsxvDnatInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxvDnatInvokeArgs()
        {
        }
        public static new GetNsxvDnatInvokeArgs Empty => new GetNsxvDnatInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxvDnatResult
    {
        public readonly string Description;
        public readonly string EdgeGateway;
        public readonly bool Enabled;
        public readonly string IcmpType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool LoggingEnabled;
        public readonly string NetworkName;
        public readonly string NetworkType;
        public readonly string? Org;
        public readonly string OriginalAddress;
        public readonly string OriginalPort;
        public readonly string Protocol;
        public readonly string RuleId;
        public readonly int RuleTag;
        public readonly string RuleType;
        public readonly string TranslatedAddress;
        public readonly string TranslatedPort;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxvDnatResult(
            string description,

            string edgeGateway,

            bool enabled,

            string icmpType,

            string id,

            bool loggingEnabled,

            string networkName,

            string networkType,

            string? org,

            string originalAddress,

            string originalPort,

            string protocol,

            string ruleId,

            int ruleTag,

            string ruleType,

            string translatedAddress,

            string translatedPort,

            string? vdc)
        {
            Description = description;
            EdgeGateway = edgeGateway;
            Enabled = enabled;
            IcmpType = icmpType;
            Id = id;
            LoggingEnabled = loggingEnabled;
            NetworkName = networkName;
            NetworkType = networkType;
            Org = org;
            OriginalAddress = originalAddress;
            OriginalPort = originalPort;
            Protocol = protocol;
            RuleId = ruleId;
            RuleTag = ruleTag;
            RuleType = ruleType;
            TranslatedAddress = translatedAddress;
            TranslatedPort = translatedPort;
            Vdc = vdc;
        }
    }
}
