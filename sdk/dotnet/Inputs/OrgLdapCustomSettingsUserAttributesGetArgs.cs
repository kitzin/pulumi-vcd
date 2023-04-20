// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgLdapCustomSettingsUserAttributesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("givenName", required: true)]
        public Input<string> GivenName { get; set; } = null!;

        [Input("groupBackLinkIdentifier")]
        public Input<string>? GroupBackLinkIdentifier { get; set; }

        [Input("groupMembershipIdentifier", required: true)]
        public Input<string> GroupMembershipIdentifier { get; set; } = null!;

        [Input("objectClass", required: true)]
        public Input<string> ObjectClass { get; set; } = null!;

        [Input("surname", required: true)]
        public Input<string> Surname { get; set; } = null!;

        [Input("telephone", required: true)]
        public Input<string> Telephone { get; set; } = null!;

        [Input("uniqueIdentifier", required: true)]
        public Input<string> UniqueIdentifier { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public OrgLdapCustomSettingsUserAttributesGetArgs()
        {
        }
        public static new OrgLdapCustomSettingsUserAttributesGetArgs Empty => new OrgLdapCustomSettingsUserAttributesGetArgs();
    }
}
