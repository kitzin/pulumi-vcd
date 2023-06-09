// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtManager
    {
        public static Task<GetNsxtManagerResult> InvokeAsync(GetNsxtManagerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtManagerResult>("vcd:index/getNsxtManager:getNsxtManager", args ?? new GetNsxtManagerArgs(), options.WithDefaults());

        public static Output<GetNsxtManagerResult> Invoke(GetNsxtManagerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtManagerResult>("vcd:index/getNsxtManager:getNsxtManager", args ?? new GetNsxtManagerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtManagerArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNsxtManagerArgs()
        {
        }
        public static new GetNsxtManagerArgs Empty => new GetNsxtManagerArgs();
    }

    public sealed class GetNsxtManagerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNsxtManagerInvokeArgs()
        {
        }
        public static new GetNsxtManagerInvokeArgs Empty => new GetNsxtManagerInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtManagerResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetNsxtManagerResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
