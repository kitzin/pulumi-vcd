// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetPortgroup
    {
        public static Task<GetPortgroupResult> InvokeAsync(GetPortgroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPortgroupResult>("vcd:index/getPortgroup:getPortgroup", args ?? new GetPortgroupArgs(), options.WithDefaults());

        public static Output<GetPortgroupResult> Invoke(GetPortgroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPortgroupResult>("vcd:index/getPortgroup:getPortgroup", args ?? new GetPortgroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortgroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetPortgroupArgs()
        {
        }
        public static new GetPortgroupArgs Empty => new GetPortgroupArgs();
    }

    public sealed class GetPortgroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetPortgroupInvokeArgs()
        {
        }
        public static new GetPortgroupInvokeArgs Empty => new GetPortgroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortgroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetPortgroupResult(
            string id,

            string name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}