// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtNetworkContextProfile
    {
        public static Task<GetNsxtNetworkContextProfileResult> InvokeAsync(GetNsxtNetworkContextProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtNetworkContextProfileResult>("vcd:index/getNsxtNetworkContextProfile:getNsxtNetworkContextProfile", args ?? new GetNsxtNetworkContextProfileArgs(), options.WithDefaults());

        public static Output<GetNsxtNetworkContextProfileResult> Invoke(GetNsxtNetworkContextProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtNetworkContextProfileResult>("vcd:index/getNsxtNetworkContextProfile:getNsxtNetworkContextProfile", args ?? new GetNsxtNetworkContextProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtNetworkContextProfileArgs : global::Pulumi.InvokeArgs
    {
        [Input("contextId", required: true)]
        public string ContextId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("scope")]
        public string? Scope { get; set; }

        public GetNsxtNetworkContextProfileArgs()
        {
        }
        public static new GetNsxtNetworkContextProfileArgs Empty => new GetNsxtNetworkContextProfileArgs();
    }

    public sealed class GetNsxtNetworkContextProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("contextId", required: true)]
        public Input<string> ContextId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public GetNsxtNetworkContextProfileInvokeArgs()
        {
        }
        public static new GetNsxtNetworkContextProfileInvokeArgs Empty => new GetNsxtNetworkContextProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtNetworkContextProfileResult
    {
        public readonly string ContextId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Scope;

        [OutputConstructor]
        private GetNsxtNetworkContextProfileResult(
            string contextId,

            string id,

            string name,

            string? scope)
        {
            ContextId = contextId;
            Id = id;
            Name = name;
            Scope = scope;
        }
    }
}
