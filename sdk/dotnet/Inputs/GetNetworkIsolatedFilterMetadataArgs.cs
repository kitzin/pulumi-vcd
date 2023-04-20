// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetNetworkIsolatedFilterMetadataInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("isSystem")]
        public Input<bool>? IsSystem { get; set; }

        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("useApiSearch")]
        public Input<bool>? UseApiSearch { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GetNetworkIsolatedFilterMetadataInputArgs()
        {
        }
        public static new GetNetworkIsolatedFilterMetadataInputArgs Empty => new GetNetworkIsolatedFilterMetadataInputArgs();
    }
}