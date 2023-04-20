// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtIpsecVpnTunnelSecurityProfileCustomizationArgs : global::Pulumi.ResourceArgs
    {
        [Input("dpdProbeInternal")]
        public Input<int>? DpdProbeInternal { get; set; }

        [Input("ikeDhGroups", required: true)]
        private InputList<string>? _ikeDhGroups;
        public InputList<string> IkeDhGroups
        {
            get => _ikeDhGroups ?? (_ikeDhGroups = new InputList<string>());
            set => _ikeDhGroups = value;
        }

        [Input("ikeDigestAlgorithms")]
        private InputList<string>? _ikeDigestAlgorithms;
        public InputList<string> IkeDigestAlgorithms
        {
            get => _ikeDigestAlgorithms ?? (_ikeDigestAlgorithms = new InputList<string>());
            set => _ikeDigestAlgorithms = value;
        }

        [Input("ikeEncryptionAlgorithms", required: true)]
        private InputList<string>? _ikeEncryptionAlgorithms;
        public InputList<string> IkeEncryptionAlgorithms
        {
            get => _ikeEncryptionAlgorithms ?? (_ikeEncryptionAlgorithms = new InputList<string>());
            set => _ikeEncryptionAlgorithms = value;
        }

        [Input("ikeSaLifetime")]
        public Input<int>? IkeSaLifetime { get; set; }

        [Input("ikeVersion", required: true)]
        public Input<string> IkeVersion { get; set; } = null!;

        [Input("tunnelDfPolicy")]
        public Input<string>? TunnelDfPolicy { get; set; }

        [Input("tunnelDhGroups", required: true)]
        private InputList<string>? _tunnelDhGroups;
        public InputList<string> TunnelDhGroups
        {
            get => _tunnelDhGroups ?? (_tunnelDhGroups = new InputList<string>());
            set => _tunnelDhGroups = value;
        }

        [Input("tunnelDigestAlgorithms")]
        private InputList<string>? _tunnelDigestAlgorithms;
        public InputList<string> TunnelDigestAlgorithms
        {
            get => _tunnelDigestAlgorithms ?? (_tunnelDigestAlgorithms = new InputList<string>());
            set => _tunnelDigestAlgorithms = value;
        }

        [Input("tunnelEncryptionAlgorithms", required: true)]
        private InputList<string>? _tunnelEncryptionAlgorithms;
        public InputList<string> TunnelEncryptionAlgorithms
        {
            get => _tunnelEncryptionAlgorithms ?? (_tunnelEncryptionAlgorithms = new InputList<string>());
            set => _tunnelEncryptionAlgorithms = value;
        }

        [Input("tunnelPfsEnabled")]
        public Input<bool>? TunnelPfsEnabled { get; set; }

        [Input("tunnelSaLifetime")]
        public Input<int>? TunnelSaLifetime { get; set; }

        public NsxtIpsecVpnTunnelSecurityProfileCustomizationArgs()
        {
        }
        public static new NsxtIpsecVpnTunnelSecurityProfileCustomizationArgs Empty => new NsxtIpsecVpnTunnelSecurityProfileCustomizationArgs();
    }
}