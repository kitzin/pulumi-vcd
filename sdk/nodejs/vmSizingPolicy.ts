// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class VmSizingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing VmSizingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VmSizingPolicyState, opts?: pulumi.CustomResourceOptions): VmSizingPolicy {
        return new VmSizingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/vmSizingPolicy:VmSizingPolicy';

    /**
     * Returns true if the given object is an instance of VmSizingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VmSizingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VmSizingPolicy.__pulumiType;
    }

    public readonly cpu!: pulumi.Output<outputs.VmSizingPolicyCpu | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly memory!: pulumi.Output<outputs.VmSizingPolicyMemory | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use - Deprecated and unneeded: will be ignored if used
     *
     * @deprecated Unneeded property, which was included by mistake
     */
    public readonly org!: pulumi.Output<string | undefined>;

    /**
     * Create a VmSizingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VmSizingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VmSizingPolicyArgs | VmSizingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VmSizingPolicyState | undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
        } else {
            const args = argsOrState as VmSizingPolicyArgs | undefined;
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VmSizingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VmSizingPolicy resources.
 */
export interface VmSizingPolicyState {
    cpu?: pulumi.Input<inputs.VmSizingPolicyCpu>;
    description?: pulumi.Input<string>;
    memory?: pulumi.Input<inputs.VmSizingPolicyMemory>;
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use - Deprecated and unneeded: will be ignored if used
     *
     * @deprecated Unneeded property, which was included by mistake
     */
    org?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VmSizingPolicy resource.
 */
export interface VmSizingPolicyArgs {
    cpu?: pulumi.Input<inputs.VmSizingPolicyCpu>;
    description?: pulumi.Input<string>;
    memory?: pulumi.Input<inputs.VmSizingPolicyMemory>;
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use - Deprecated and unneeded: will be ignored if used
     *
     * @deprecated Unneeded property, which was included by mistake
     */
    org?: pulumi.Input<string>;
}
