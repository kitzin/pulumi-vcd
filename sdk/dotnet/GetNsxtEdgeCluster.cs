// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtEdgeCluster
    {
        public static Task<GetNsxtEdgeClusterResult> InvokeAsync(GetNsxtEdgeClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtEdgeClusterResult>("vcd:index/getNsxtEdgeCluster:getNsxtEdgeCluster", args ?? new GetNsxtEdgeClusterArgs(), options.WithDefaults());

        public static Output<GetNsxtEdgeClusterResult> Invoke(GetNsxtEdgeClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtEdgeClusterResult>("vcd:index/getNsxtEdgeCluster:getNsxtEdgeCluster", args ?? new GetNsxtEdgeClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtEdgeClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("providerVdcId")]
        public string? ProviderVdcId { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        [Input("vdcGroupId")]
        public string? VdcGroupId { get; set; }

        [Input("vdcId")]
        public string? VdcId { get; set; }

        public GetNsxtEdgeClusterArgs()
        {
        }
        public static new GetNsxtEdgeClusterArgs Empty => new GetNsxtEdgeClusterArgs();
    }

    public sealed class GetNsxtEdgeClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("providerVdcId")]
        public Input<string>? ProviderVdcId { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        [Input("vdcGroupId")]
        public Input<string>? VdcGroupId { get; set; }

        [Input("vdcId")]
        public Input<string>? VdcId { get; set; }

        public GetNsxtEdgeClusterInvokeArgs()
        {
        }
        public static new GetNsxtEdgeClusterInvokeArgs Empty => new GetNsxtEdgeClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtEdgeClusterResult
    {
        public readonly string DeploymentType;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly int NodeCount;
        public readonly string NodeType;
        public readonly string? Org;
        public readonly string? ProviderVdcId;
        public readonly string? Vdc;
        public readonly string? VdcGroupId;
        public readonly string? VdcId;

        [OutputConstructor]
        private GetNsxtEdgeClusterResult(
            string deploymentType,

            string description,

            string id,

            string name,

            int nodeCount,

            string nodeType,

            string? org,

            string? providerVdcId,

            string? vdc,

            string? vdcGroupId,

            string? vdcId)
        {
            DeploymentType = deploymentType;
            Description = description;
            Id = id;
            Name = name;
            NodeCount = nodeCount;
            NodeType = nodeType;
            Org = org;
            ProviderVdcId = providerVdcId;
            Vdc = vdc;
            VdcGroupId = vdcGroupId;
            VdcId = vdcId;
        }
    }
}
