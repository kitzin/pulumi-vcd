// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtAlbController
    {
        public static Task<GetNsxtAlbControllerResult> InvokeAsync(GetNsxtAlbControllerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtAlbControllerResult>("vcd:index/getNsxtAlbController:getNsxtAlbController", args ?? new GetNsxtAlbControllerArgs(), options.WithDefaults());

        public static Output<GetNsxtAlbControllerResult> Invoke(GetNsxtAlbControllerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtAlbControllerResult>("vcd:index/getNsxtAlbController:getNsxtAlbController", args ?? new GetNsxtAlbControllerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtAlbControllerArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNsxtAlbControllerArgs()
        {
        }
        public static new GetNsxtAlbControllerArgs Empty => new GetNsxtAlbControllerArgs();
    }

    public sealed class GetNsxtAlbControllerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNsxtAlbControllerInvokeArgs()
        {
        }
        public static new GetNsxtAlbControllerInvokeArgs Empty => new GetNsxtAlbControllerInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtAlbControllerResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LicenseType;
        public readonly string Name;
        public readonly string Url;
        public readonly string Username;
        public readonly string Version;

        [OutputConstructor]
        private GetNsxtAlbControllerResult(
            string description,

            string id,

            string licenseType,

            string name,

            string url,

            string username,

            string version)
        {
            Description = description;
            Id = id;
            LicenseType = licenseType;
            Name = name;
            Url = url;
            Username = username;
            Version = version;
        }
    }
}