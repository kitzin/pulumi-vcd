// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtTier0Router
    {
        public static Task<GetNsxtTier0RouterResult> InvokeAsync(GetNsxtTier0RouterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtTier0RouterResult>("vcd:index/getNsxtTier0Router:getNsxtTier0Router", args ?? new GetNsxtTier0RouterArgs(), options.WithDefaults());

        public static Output<GetNsxtTier0RouterResult> Invoke(GetNsxtTier0RouterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtTier0RouterResult>("vcd:index/getNsxtTier0Router:getNsxtTier0Router", args ?? new GetNsxtTier0RouterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtTier0RouterArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("nsxtManagerId", required: true)]
        public string NsxtManagerId { get; set; } = null!;

        public GetNsxtTier0RouterArgs()
        {
        }
        public static new GetNsxtTier0RouterArgs Empty => new GetNsxtTier0RouterArgs();
    }

    public sealed class GetNsxtTier0RouterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("nsxtManagerId", required: true)]
        public Input<string> NsxtManagerId { get; set; } = null!;

        public GetNsxtTier0RouterInvokeArgs()
        {
        }
        public static new GetNsxtTier0RouterInvokeArgs Empty => new GetNsxtTier0RouterInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtTier0RouterResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsAssigned;
        public readonly string Name;
        public readonly string NsxtManagerId;

        [OutputConstructor]
        private GetNsxtTier0RouterResult(
            string id,

            bool isAssigned,

            string name,

            string nsxtManagerId)
        {
            Id = id;
            IsAssigned = isAssigned;
            Name = name;
            NsxtManagerId = nsxtManagerId;
        }
    }
}
