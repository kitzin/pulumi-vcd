// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetNetworkRoutedFilterMetadataArgs : global::Pulumi.InvokeArgs
    {
        [Input("isSystem")]
        public bool? IsSystem { get; set; }

        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("type")]
        public string? Type { get; set; }

        [Input("useApiSearch")]
        public bool? UseApiSearch { get; set; }

        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetNetworkRoutedFilterMetadataArgs()
        {
        }
        public static new GetNetworkRoutedFilterMetadataArgs Empty => new GetNetworkRoutedFilterMetadataArgs();
    }
}
