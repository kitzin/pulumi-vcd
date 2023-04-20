// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetVapp
    {
        public static Task<GetVappResult> InvokeAsync(GetVappArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVappResult>("vcd:index/getVapp:getVapp", args ?? new GetVappArgs(), options.WithDefaults());

        public static Output<GetVappResult> Invoke(GetVappInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVappResult>("vcd:index/getVapp:getVapp", args ?? new GetVappInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVappArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetVappArgs()
        {
        }
        public static new GetVappArgs Empty => new GetVappArgs();
    }

    public sealed class GetVappInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetVappInvokeArgs()
        {
        }
        public static new GetVappInvokeArgs Empty => new GetVappInvokeArgs();
    }


    [OutputType]
    public sealed class GetVappResult
    {
        public readonly string Description;
        public readonly ImmutableDictionary<string, object> GuestProperties;
        public readonly string Href;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetVappLeaseResult> Leases;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetVappMetadataEntryResult> MetadataEntries;
        public readonly string Name;
        public readonly string? Org;
        public readonly int Status;
        public readonly string StatusText;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetVappResult(
            string description,

            ImmutableDictionary<string, object> guestProperties,

            string href,

            string id,

            ImmutableArray<Outputs.GetVappLeaseResult> leases,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetVappMetadataEntryResult> metadataEntries,

            string name,

            string? org,

            int status,

            string statusText,

            string? vdc)
        {
            Description = description;
            GuestProperties = guestProperties;
            Href = href;
            Id = id;
            Leases = leases;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            Status = status;
            StatusText = statusText;
            Vdc = vdc;
        }
    }
}