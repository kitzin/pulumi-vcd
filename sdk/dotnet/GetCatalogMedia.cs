// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetCatalogMedia
    {
        public static Task<GetCatalogMediaResult> InvokeAsync(GetCatalogMediaArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCatalogMediaResult>("vcd:index/getCatalogMedia:getCatalogMedia", args ?? new GetCatalogMediaArgs(), options.WithDefaults());

        public static Output<GetCatalogMediaResult> Invoke(GetCatalogMediaInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCatalogMediaResult>("vcd:index/getCatalogMedia:getCatalogMedia", args ?? new GetCatalogMediaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCatalogMediaArgs : global::Pulumi.InvokeArgs
    {
        [Input("catalog")]
        public string? Catalog { get; set; }

        [Input("catalogId")]
        public string? CatalogId { get; set; }

        [Input("filter")]
        public Inputs.GetCatalogMediaFilterArgs? Filter { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        public GetCatalogMediaArgs()
        {
        }
        public static new GetCatalogMediaArgs Empty => new GetCatalogMediaArgs();
    }

    public sealed class GetCatalogMediaInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        [Input("filter")]
        public Input<Inputs.GetCatalogMediaFilterInputArgs>? Filter { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        public GetCatalogMediaInvokeArgs()
        {
        }
        public static new GetCatalogMediaInvokeArgs Empty => new GetCatalogMediaInvokeArgs();
    }


    [OutputType]
    public sealed class GetCatalogMediaResult
    {
        public readonly string Catalog;
        public readonly string? CatalogId;
        public readonly string CreationDate;
        public readonly string Description;
        public readonly Outputs.GetCatalogMediaFilterResult? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsIso;
        public readonly bool IsPublished;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetCatalogMediaMetadataEntryResult> MetadataEntries;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string OwnerName;
        public readonly int Size;
        public readonly string Status;
        public readonly string StorageProfileName;

        [OutputConstructor]
        private GetCatalogMediaResult(
            string catalog,

            string? catalogId,

            string creationDate,

            string description,

            Outputs.GetCatalogMediaFilterResult? filter,

            string id,

            bool isIso,

            bool isPublished,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetCatalogMediaMetadataEntryResult> metadataEntries,

            string? name,

            string? org,

            string ownerName,

            int size,

            string status,

            string storageProfileName)
        {
            Catalog = catalog;
            CatalogId = catalogId;
            CreationDate = creationDate;
            Description = description;
            Filter = filter;
            Id = id;
            IsIso = isIso;
            IsPublished = isPublished;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            OwnerName = ownerName;
            Size = size;
            Status = status;
            StorageProfileName = storageProfileName;
        }
    }
}
