// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NsxtDistributedFirewall extends pulumi.CustomResource {
    /**
     * Get an existing NsxtDistributedFirewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtDistributedFirewallState, opts?: pulumi.CustomResourceOptions): NsxtDistributedFirewall {
        return new NsxtDistributedFirewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtDistributedFirewall:NsxtDistributedFirewall';

    /**
     * Returns true if the given object is an instance of NsxtDistributedFirewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtDistributedFirewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtDistributedFirewall.__pulumiType;
    }

    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Ordered list of firewall rules
     */
    public readonly rules!: pulumi.Output<outputs.NsxtDistributedFirewallRule[]>;
    /**
     * ID of VDC Group for Distributed Firewall
     */
    public readonly vdcGroupId!: pulumi.Output<string>;

    /**
     * Create a NsxtDistributedFirewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtDistributedFirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtDistributedFirewallArgs | NsxtDistributedFirewallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtDistributedFirewallState | undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vdcGroupId"] = state ? state.vdcGroupId : undefined;
        } else {
            const args = argsOrState as NsxtDistributedFirewallArgs | undefined;
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            if ((!args || args.vdcGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vdcGroupId'");
            }
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vdcGroupId"] = args ? args.vdcGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtDistributedFirewall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtDistributedFirewall resources.
 */
export interface NsxtDistributedFirewallState {
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Ordered list of firewall rules
     */
    rules?: pulumi.Input<pulumi.Input<inputs.NsxtDistributedFirewallRule>[]>;
    /**
     * ID of VDC Group for Distributed Firewall
     */
    vdcGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtDistributedFirewall resource.
 */
export interface NsxtDistributedFirewallArgs {
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Ordered list of firewall rules
     */
    rules: pulumi.Input<pulumi.Input<inputs.NsxtDistributedFirewallRule>[]>;
    /**
     * ID of VDC Group for Distributed Firewall
     */
    vdcGroupId: pulumi.Input<string>;
}