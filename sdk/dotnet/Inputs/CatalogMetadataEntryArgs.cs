// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class CatalogMetadataEntryArgs : global::Pulumi.ResourceArgs
    {
        [Input("isSystem")]
        public Input<bool>? IsSystem { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userAccess")]
        public Input<string>? UserAccess { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public CatalogMetadataEntryArgs()
        {
        }
        public static new CatalogMetadataEntryArgs Empty => new CatalogMetadataEntryArgs();
    }
}
