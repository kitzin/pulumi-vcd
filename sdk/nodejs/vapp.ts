// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Vapp extends pulumi.CustomResource {
    /**
     * Get an existing Vapp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VappState, opts?: pulumi.CustomResourceOptions): Vapp {
        return new Vapp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/vapp:Vapp';

    /**
     * Returns true if the given object is an instance of Vapp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vapp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vapp.__pulumiType;
    }

    /**
     * Optional description of the vApp
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Key/value settings for guest properties. Will be picked up by new VMs when created.
     */
    public readonly guestProperties!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * vApp Hyper Reference
     */
    public /*out*/ readonly href!: pulumi.Output<string>;
    /**
     * Defines lease parameters for this vApp
     */
    public readonly lease!: pulumi.Output<outputs.VappLease>;
    /**
     * Key value map of metadata to assign to this vApp. Key and value can be any string.
     *
     * @deprecated Use metadata_entry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * Metadata entries for the given vApp
     */
    public readonly metadataEntries!: pulumi.Output<outputs.VappMetadataEntry[]>;
    /**
     * A name for the vApp, unique withing the VDC
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * A boolean value stating if this vApp should be powered on
     */
    public readonly powerOn!: pulumi.Output<boolean | undefined>;
    /**
     * Shows the status code of the vApp
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Shows the status of the vApp
     */
    public /*out*/ readonly statusText!: pulumi.Output<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a Vapp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VappArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VappArgs | VappState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VappState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["guestProperties"] = state ? state.guestProperties : undefined;
            resourceInputs["href"] = state ? state.href : undefined;
            resourceInputs["lease"] = state ? state.lease : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["powerOn"] = state ? state.powerOn : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusText"] = state ? state.statusText : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as VappArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["guestProperties"] = args ? args.guestProperties : undefined;
            resourceInputs["lease"] = args ? args.lease : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["powerOn"] = args ? args.powerOn : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["href"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusText"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vapp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vapp resources.
 */
export interface VappState {
    /**
     * Optional description of the vApp
     */
    description?: pulumi.Input<string>;
    /**
     * Key/value settings for guest properties. Will be picked up by new VMs when created.
     */
    guestProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * vApp Hyper Reference
     */
    href?: pulumi.Input<string>;
    /**
     * Defines lease parameters for this vApp
     */
    lease?: pulumi.Input<inputs.VappLease>;
    /**
     * Key value map of metadata to assign to this vApp. Key and value can be any string.
     *
     * @deprecated Use metadata_entry instead
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given vApp
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.VappMetadataEntry>[]>;
    /**
     * A name for the vApp, unique withing the VDC
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * A boolean value stating if this vApp should be powered on
     */
    powerOn?: pulumi.Input<boolean>;
    /**
     * Shows the status code of the vApp
     */
    status?: pulumi.Input<number>;
    /**
     * Shows the status of the vApp
     */
    statusText?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vapp resource.
 */
export interface VappArgs {
    /**
     * Optional description of the vApp
     */
    description?: pulumi.Input<string>;
    /**
     * Key/value settings for guest properties. Will be picked up by new VMs when created.
     */
    guestProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Defines lease parameters for this vApp
     */
    lease?: pulumi.Input<inputs.VappLease>;
    /**
     * Key value map of metadata to assign to this vApp. Key and value can be any string.
     *
     * @deprecated Use metadata_entry instead
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given vApp
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.VappMetadataEntry>[]>;
    /**
     * A name for the vApp, unique withing the VDC
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * A boolean value stating if this vApp should be powered on
     */
    powerOn?: pulumi.Input<boolean>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
