// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtAlbCloud
    {
        public static Task<GetNsxtAlbCloudResult> InvokeAsync(GetNsxtAlbCloudArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtAlbCloudResult>("vcd:index/getNsxtAlbCloud:getNsxtAlbCloud", args ?? new GetNsxtAlbCloudArgs(), options.WithDefaults());

        public static Output<GetNsxtAlbCloudResult> Invoke(GetNsxtAlbCloudInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtAlbCloudResult>("vcd:index/getNsxtAlbCloud:getNsxtAlbCloud", args ?? new GetNsxtAlbCloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtAlbCloudArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNsxtAlbCloudArgs()
        {
        }
        public static new GetNsxtAlbCloudArgs Empty => new GetNsxtAlbCloudArgs();
    }

    public sealed class GetNsxtAlbCloudInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNsxtAlbCloudInvokeArgs()
        {
        }
        public static new GetNsxtAlbCloudInvokeArgs Empty => new GetNsxtAlbCloudInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtAlbCloudResult
    {
        public readonly string ControllerId;
        public readonly string Description;
        public readonly string HealthMessage;
        public readonly string HealthStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImportableCloudId;
        public readonly string Name;
        public readonly string NetworkPoolId;
        public readonly string NetworkPoolName;

        [OutputConstructor]
        private GetNsxtAlbCloudResult(
            string controllerId,

            string description,

            string healthMessage,

            string healthStatus,

            string id,

            string importableCloudId,

            string name,

            string networkPoolId,

            string networkPoolName)
        {
            ControllerId = controllerId;
            Description = description;
            HealthMessage = healthMessage;
            HealthStatus = healthStatus;
            Id = id;
            ImportableCloudId = importableCloudId;
            Name = name;
            NetworkPoolId = networkPoolId;
            NetworkPoolName = networkPoolName;
        }
    }
}
