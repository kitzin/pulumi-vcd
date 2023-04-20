// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SecurityTag extends pulumi.CustomResource {
    /**
     * Get an existing SecurityTag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityTagState, opts?: pulumi.CustomResourceOptions): SecurityTag {
        return new SecurityTag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/securityTag:SecurityTag';

    /**
     * Returns true if the given object is an instance of SecurityTag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityTag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityTag.__pulumiType;
    }

    /**
     * Security tag name to be created
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * List of VM IDs that the security tags is going to be tied to
     */
    public readonly vmIds!: pulumi.Output<string[]>;

    /**
     * Create a SecurityTag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityTagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityTagArgs | SecurityTagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityTagState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["vmIds"] = state ? state.vmIds : undefined;
        } else {
            const args = argsOrState as SecurityTagArgs | undefined;
            if ((!args || args.vmIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmIds'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["vmIds"] = args ? args.vmIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityTag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityTag resources.
 */
export interface SecurityTagState {
    /**
     * Security tag name to be created
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * List of VM IDs that the security tags is going to be tied to
     */
    vmIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecurityTag resource.
 */
export interface SecurityTagArgs {
    /**
     * Security tag name to be created
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * List of VM IDs that the security tags is going to be tied to
     */
    vmIds: pulumi.Input<pulumi.Input<string>[]>;
}
