// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceServicePortGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        [Input("sslEnabled")]
        public Input<bool>? SslEnabled { get; set; }

        [Input("startPort", required: true)]
        public Input<int> StartPort { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public NsxtAlbVirtualServiceServicePortGetArgs()
        {
        }
        public static new NsxtAlbVirtualServiceServicePortGetArgs Empty => new NsxtAlbVirtualServiceServicePortGetArgs();
    }
}
