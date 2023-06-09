// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtAlbImportableCloud
    {
        public static Task<GetNsxtAlbImportableCloudResult> InvokeAsync(GetNsxtAlbImportableCloudArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtAlbImportableCloudResult>("vcd:index/getNsxtAlbImportableCloud:getNsxtAlbImportableCloud", args ?? new GetNsxtAlbImportableCloudArgs(), options.WithDefaults());

        public static Output<GetNsxtAlbImportableCloudResult> Invoke(GetNsxtAlbImportableCloudInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtAlbImportableCloudResult>("vcd:index/getNsxtAlbImportableCloud:getNsxtAlbImportableCloud", args ?? new GetNsxtAlbImportableCloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtAlbImportableCloudArgs : global::Pulumi.InvokeArgs
    {
        [Input("controllerId", required: true)]
        public string ControllerId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNsxtAlbImportableCloudArgs()
        {
        }
        public static new GetNsxtAlbImportableCloudArgs Empty => new GetNsxtAlbImportableCloudArgs();
    }

    public sealed class GetNsxtAlbImportableCloudInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controllerId", required: true)]
        public Input<string> ControllerId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNsxtAlbImportableCloudInvokeArgs()
        {
        }
        public static new GetNsxtAlbImportableCloudInvokeArgs Empty => new GetNsxtAlbImportableCloudInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtAlbImportableCloudResult
    {
        public readonly bool AlreadyImported;
        public readonly string ControllerId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string NetworkPoolId;
        public readonly string NetworkPoolName;
        public readonly string TransportZoneName;

        [OutputConstructor]
        private GetNsxtAlbImportableCloudResult(
            bool alreadyImported,

            string controllerId,

            string id,

            string name,

            string networkPoolId,

            string networkPoolName,

            string transportZoneName)
        {
            AlreadyImported = alreadyImported;
            ControllerId = controllerId;
            Id = id;
            Name = name;
            NetworkPoolId = networkPoolId;
            NetworkPoolName = networkPoolName;
            TransportZoneName = transportZoneName;
        }
    }
}
