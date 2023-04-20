// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/orgLdap:OrgLdap")]
    public partial class OrgLdap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Custom settings when `ldap_mode` is CUSTOM
        /// </summary>
        [Output("customSettings")]
        public Output<Outputs.OrgLdapCustomSettings?> CustomSettings { get; private set; } = null!;

        /// <summary>
        /// Type of LDAP settings (one of NONE, SYSTEM, CUSTOM)
        /// </summary>
        [Output("ldapMode")]
        public Output<string> LdapMode { get; private set; } = null!;

        /// <summary>
        /// Organization ID
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;


        /// <summary>
        /// Create a OrgLdap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgLdap(string name, OrgLdapArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/orgLdap:OrgLdap", name, args ?? new OrgLdapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgLdap(string name, Input<string> id, OrgLdapState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/orgLdap:OrgLdap", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrgLdap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgLdap Get(string name, Input<string> id, OrgLdapState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgLdap(name, id, state, options);
        }
    }

    public sealed class OrgLdapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom settings when `ldap_mode` is CUSTOM
        /// </summary>
        [Input("customSettings")]
        public Input<Inputs.OrgLdapCustomSettingsArgs>? CustomSettings { get; set; }

        /// <summary>
        /// Type of LDAP settings (one of NONE, SYSTEM, CUSTOM)
        /// </summary>
        [Input("ldapMode", required: true)]
        public Input<string> LdapMode { get; set; } = null!;

        /// <summary>
        /// Organization ID
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        public OrgLdapArgs()
        {
        }
        public static new OrgLdapArgs Empty => new OrgLdapArgs();
    }

    public sealed class OrgLdapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom settings when `ldap_mode` is CUSTOM
        /// </summary>
        [Input("customSettings")]
        public Input<Inputs.OrgLdapCustomSettingsGetArgs>? CustomSettings { get; set; }

        /// <summary>
        /// Type of LDAP settings (one of NONE, SYSTEM, CUSTOM)
        /// </summary>
        [Input("ldapMode")]
        public Input<string>? LdapMode { get; set; }

        /// <summary>
        /// Organization ID
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public OrgLdapState()
        {
        }
        public static new OrgLdapState Empty => new OrgLdapState();
    }
}
