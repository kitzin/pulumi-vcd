// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetOrgVdcMetadataEntryArgs : global::Pulumi.InvokeArgs
    {
        [Input("isSystem")]
        public bool? IsSystem { get; set; }

        [Input("key")]
        public string? Key { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        [Input("userAccess")]
        public string? UserAccess { get; set; }

        [Input("value")]
        public string? Value { get; set; }

        public GetOrgVdcMetadataEntryArgs()
        {
        }
        public static new GetOrgVdcMetadataEntryArgs Empty => new GetOrgVdcMetadataEntryArgs();
    }
}
