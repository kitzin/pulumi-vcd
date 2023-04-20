// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/securityTag:SecurityTag")]
    public partial class SecurityTag : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Security tag name to be created
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// List of VM IDs that the security tags is going to be tied to
        /// </summary>
        [Output("vmIds")]
        public Output<ImmutableArray<string>> VmIds { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityTag(string name, SecurityTagArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/securityTag:SecurityTag", name, args ?? new SecurityTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityTag(string name, Input<string> id, SecurityTagState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/securityTag:SecurityTag", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityTag Get(string name, Input<string> id, SecurityTagState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityTag(name, id, state, options);
        }
    }

    public sealed class SecurityTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Security tag name to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vmIds", required: true)]
        private InputList<string>? _vmIds;

        /// <summary>
        /// List of VM IDs that the security tags is going to be tied to
        /// </summary>
        public InputList<string> VmIds
        {
            get => _vmIds ?? (_vmIds = new InputList<string>());
            set => _vmIds = value;
        }

        public SecurityTagArgs()
        {
        }
        public static new SecurityTagArgs Empty => new SecurityTagArgs();
    }

    public sealed class SecurityTagState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Security tag name to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vmIds")]
        private InputList<string>? _vmIds;

        /// <summary>
        /// List of VM IDs that the security tags is going to be tied to
        /// </summary>
        public InputList<string> VmIds
        {
            get => _vmIds ?? (_vmIds = new InputList<string>());
            set => _vmIds = value;
        }

        public SecurityTagState()
        {
        }
        public static new SecurityTagState Empty => new SecurityTagState();
    }
}
