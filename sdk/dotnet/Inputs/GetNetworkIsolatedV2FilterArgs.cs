// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetNetworkIsolatedV2FilterInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        public GetNetworkIsolatedV2FilterInputArgs()
        {
        }
        public static new GetNetworkIsolatedV2FilterInputArgs Empty => new GetNetworkIsolatedV2FilterInputArgs();
    }
}
