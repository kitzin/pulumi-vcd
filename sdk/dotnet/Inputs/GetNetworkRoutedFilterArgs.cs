// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetNetworkRoutedFilterInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.GetNetworkRoutedFilterMetadataInputArgs>? _metadatas;
        public InputList<Inputs.GetNetworkRoutedFilterMetadataInputArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.GetNetworkRoutedFilterMetadataInputArgs>());
            set => _metadatas = value;
        }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        public GetNetworkRoutedFilterInputArgs()
        {
        }
        public static new GetNetworkRoutedFilterInputArgs Empty => new GetNetworkRoutedFilterInputArgs();
    }
}
