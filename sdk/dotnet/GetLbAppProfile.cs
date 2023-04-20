// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetLbAppProfile
    {
        public static Task<GetLbAppProfileResult> InvokeAsync(GetLbAppProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLbAppProfileResult>("vcd:index/getLbAppProfile:getLbAppProfile", args ?? new GetLbAppProfileArgs(), options.WithDefaults());

        public static Output<GetLbAppProfileResult> Invoke(GetLbAppProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLbAppProfileResult>("vcd:index/getLbAppProfile:getLbAppProfile", args ?? new GetLbAppProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbAppProfileArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public string EdgeGateway { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetLbAppProfileArgs()
        {
        }
        public static new GetLbAppProfileArgs Empty => new GetLbAppProfileArgs();
    }

    public sealed class GetLbAppProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetLbAppProfileInvokeArgs()
        {
        }
        public static new GetLbAppProfileInvokeArgs Empty => new GetLbAppProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbAppProfileResult
    {
        public readonly string CookieMode;
        public readonly string CookieName;
        public readonly string EdgeGateway;
        public readonly bool EnablePoolSideSsl;
        public readonly bool EnableSslPassthrough;
        public readonly int Expiration;
        public readonly string HttpRedirectUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool InsertXForwardedHttpHeader;
        public readonly string Name;
        public readonly string? Org;
        public readonly string PersistenceMechanism;
        public readonly string Type;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetLbAppProfileResult(
            string cookieMode,

            string cookieName,

            string edgeGateway,

            bool enablePoolSideSsl,

            bool enableSslPassthrough,

            int expiration,

            string httpRedirectUrl,

            string id,

            bool insertXForwardedHttpHeader,

            string name,

            string? org,

            string persistenceMechanism,

            string type,

            string? vdc)
        {
            CookieMode = cookieMode;
            CookieName = cookieName;
            EdgeGateway = edgeGateway;
            EnablePoolSideSsl = enablePoolSideSsl;
            EnableSslPassthrough = enableSslPassthrough;
            Expiration = expiration;
            HttpRedirectUrl = httpRedirectUrl;
            Id = id;
            InsertXForwardedHttpHeader = insertXForwardedHttpHeader;
            Name = name;
            Org = org;
            PersistenceMechanism = persistenceMechanism;
            Type = type;
            Vdc = vdc;
        }
    }
}
