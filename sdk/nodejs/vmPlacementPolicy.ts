// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VmPlacementPolicy extends pulumi.CustomResource {
    /**
     * Get an existing VmPlacementPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VmPlacementPolicyState, opts?: pulumi.CustomResourceOptions): VmPlacementPolicy {
        return new VmPlacementPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/vmPlacementPolicy:VmPlacementPolicy';

    /**
     * Returns true if the given object is an instance of VmPlacementPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VmPlacementPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VmPlacementPolicy.__pulumiType;
    }

    /**
     * Description of the VM Placement Policy
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
     * entries set in this attribute
     */
    public readonly logicalVmGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the VM Placement Policy
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the Provider VDC to which the VM Placement Policy belongs
     */
    public readonly providerVdcId!: pulumi.Output<string>;
    /**
     * IDs of the collection of VMs with similar host requirements
     */
    public readonly vmGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VmPlacementPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VmPlacementPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VmPlacementPolicyArgs | VmPlacementPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VmPlacementPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["logicalVmGroupIds"] = state ? state.logicalVmGroupIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["providerVdcId"] = state ? state.providerVdcId : undefined;
            resourceInputs["vmGroupIds"] = state ? state.vmGroupIds : undefined;
        } else {
            const args = argsOrState as VmPlacementPolicyArgs | undefined;
            if ((!args || args.providerVdcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerVdcId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["logicalVmGroupIds"] = args ? args.logicalVmGroupIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["providerVdcId"] = args ? args.providerVdcId : undefined;
            resourceInputs["vmGroupIds"] = args ? args.vmGroupIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VmPlacementPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VmPlacementPolicy resources.
 */
export interface VmPlacementPolicyState {
    /**
     * Description of the VM Placement Policy
     */
    description?: pulumi.Input<string>;
    /**
     * IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
     * entries set in this attribute
     */
    logicalVmGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the VM Placement Policy
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the Provider VDC to which the VM Placement Policy belongs
     */
    providerVdcId?: pulumi.Input<string>;
    /**
     * IDs of the collection of VMs with similar host requirements
     */
    vmGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a VmPlacementPolicy resource.
 */
export interface VmPlacementPolicyArgs {
    /**
     * Description of the VM Placement Policy
     */
    description?: pulumi.Input<string>;
    /**
     * IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
     * entries set in this attribute
     */
    logicalVmGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the VM Placement Policy
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the Provider VDC to which the VM Placement Policy belongs
     */
    providerVdcId: pulumi.Input<string>;
    /**
     * IDs of the collection of VMs with similar host requirements
     */
    vmGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
