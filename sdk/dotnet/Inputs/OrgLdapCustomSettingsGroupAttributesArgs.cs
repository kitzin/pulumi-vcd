// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgLdapCustomSettingsGroupAttributesArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupBackLinkIdentifier")]
        public Input<string>? GroupBackLinkIdentifier { get; set; }

        [Input("groupMembershipIdentifier", required: true)]
        public Input<string> GroupMembershipIdentifier { get; set; } = null!;

        [Input("membership", required: true)]
        public Input<string> Membership { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("objectClass", required: true)]
        public Input<string> ObjectClass { get; set; } = null!;

        [Input("uniqueIdentifier", required: true)]
        public Input<string> UniqueIdentifier { get; set; } = null!;

        public OrgLdapCustomSettingsGroupAttributesArgs()
        {
        }
        public static new OrgLdapCustomSettingsGroupAttributesArgs Empty => new OrgLdapCustomSettingsGroupAttributesArgs();
    }
}