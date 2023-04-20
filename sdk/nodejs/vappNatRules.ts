// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class VappNatRules extends pulumi.CustomResource {
    /**
     * Get an existing VappNatRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VappNatRulesState, opts?: pulumi.CustomResourceOptions): VappNatRules {
        return new VappNatRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/vappNatRules:VappNatRules';

    /**
     * Returns true if the given object is an instance of VappNatRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VappNatRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VappNatRules.__pulumiType;
    }

    /**
     * When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
     */
    public readonly enableIpMasquerade!: pulumi.Output<boolean | undefined>;
    /**
     * Enable or disable NAT service. Default is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
     */
    public readonly natType!: pulumi.Output<string>;
    /**
     * vApp network identifier
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    public readonly rules!: pulumi.Output<outputs.VappNatRulesRule[] | undefined>;
    /**
     * vApp identifier
     */
    public readonly vappId!: pulumi.Output<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a VappNatRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VappNatRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VappNatRulesArgs | VappNatRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VappNatRulesState | undefined;
            resourceInputs["enableIpMasquerade"] = state ? state.enableIpMasquerade : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["natType"] = state ? state.natType : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vappId"] = state ? state.vappId : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as VappNatRulesArgs | undefined;
            if ((!args || args.natType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natType'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.vappId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vappId'");
            }
            resourceInputs["enableIpMasquerade"] = args ? args.enableIpMasquerade : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["natType"] = args ? args.natType : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vappId"] = args ? args.vappId : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VappNatRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VappNatRules resources.
 */
export interface VappNatRulesState {
    /**
     * When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
     */
    enableIpMasquerade?: pulumi.Input<boolean>;
    /**
     * Enable or disable NAT service. Default is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
     */
    natType?: pulumi.Input<string>;
    /**
     * vApp network identifier
     */
    networkId?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.VappNatRulesRule>[]>;
    /**
     * vApp identifier
     */
    vappId?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VappNatRules resource.
 */
export interface VappNatRulesArgs {
    /**
     * When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
     */
    enableIpMasquerade?: pulumi.Input<boolean>;
    /**
     * Enable or disable NAT service. Default is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
     */
    natType: pulumi.Input<string>;
    /**
     * vApp network identifier
     */
    networkId: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.VappNatRulesRule>[]>;
    /**
     * vApp identifier
     */
    vappId: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
