// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Outputs
{

    [OutputType]
    public sealed class GetNetworkRoutedFilterMetadataResult
    {
        public readonly bool? IsSystem;
        public readonly string Key;
        public readonly string? Type;
        public readonly bool? UseApiSearch;
        public readonly string Value;

        [OutputConstructor]
        private GetNetworkRoutedFilterMetadataResult(
            bool? isSystem,

            string key,

            string? type,

            bool? useApiSearch,

            string value)
        {
            IsSystem = isSystem;
            Key = key;
            Type = type;
            UseApiSearch = useApiSearch;
            Value = value;
        }
    }
}
