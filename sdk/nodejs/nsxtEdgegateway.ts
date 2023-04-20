// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NsxtEdgegateway extends pulumi.CustomResource {
    /**
     * Get an existing NsxtEdgegateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtEdgegatewayState, opts?: pulumi.CustomResourceOptions): NsxtEdgegateway {
        return new NsxtEdgegateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtEdgegateway:NsxtEdgegateway';

    /**
     * Returns true if the given object is an instance of NsxtEdgegateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtEdgegateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtEdgegateway.__pulumiType;
    }

    /**
     * Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
     */
    public readonly dedicateExternalNetwork!: pulumi.Output<boolean | undefined>;
    /**
     * Edge Gateway description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified
     */
    public readonly edgeClusterId!: pulumi.Output<string>;
    /**
     * External network ID
     */
    public readonly externalNetworkId!: pulumi.Output<string>;
    /**
     * Edge Gateway name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * ID of VDC or VDC Group
     */
    public readonly ownerId!: pulumi.Output<string>;
    /**
     * Primary IP address of edge gateway. Read-only (can be specified in specific subnet)
     */
    public /*out*/ readonly primaryIp!: pulumi.Output<string>;
    /**
     * Optional ID of starting VDC if the 'owner_id' is a VDC Group
     */
    public readonly startingVdcId!: pulumi.Output<string | undefined>;
    /**
     * One or more blocks with external network information to be attached to this gateway's interface
     */
    public readonly subnets!: pulumi.Output<outputs.NsxtEdgegatewaySubnet[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    public readonly vdc!: pulumi.Output<string>;

    /**
     * Create a NsxtEdgegateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtEdgegatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtEdgegatewayArgs | NsxtEdgegatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtEdgegatewayState | undefined;
            resourceInputs["dedicateExternalNetwork"] = state ? state.dedicateExternalNetwork : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeClusterId"] = state ? state.edgeClusterId : undefined;
            resourceInputs["externalNetworkId"] = state ? state.externalNetworkId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["primaryIp"] = state ? state.primaryIp : undefined;
            resourceInputs["startingVdcId"] = state ? state.startingVdcId : undefined;
            resourceInputs["subnets"] = state ? state.subnets : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NsxtEdgegatewayArgs | undefined;
            if ((!args || args.externalNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalNetworkId'");
            }
            if ((!args || args.subnets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnets'");
            }
            resourceInputs["dedicateExternalNetwork"] = args ? args.dedicateExternalNetwork : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeClusterId"] = args ? args.edgeClusterId : undefined;
            resourceInputs["externalNetworkId"] = args ? args.externalNetworkId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["startingVdcId"] = args ? args.startingVdcId : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["primaryIp"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtEdgegateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtEdgegateway resources.
 */
export interface NsxtEdgegatewayState {
    /**
     * Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
     */
    dedicateExternalNetwork?: pulumi.Input<boolean>;
    /**
     * Edge Gateway description
     */
    description?: pulumi.Input<string>;
    /**
     * Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified
     */
    edgeClusterId?: pulumi.Input<string>;
    /**
     * External network ID
     */
    externalNetworkId?: pulumi.Input<string>;
    /**
     * Edge Gateway name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * ID of VDC or VDC Group
     */
    ownerId?: pulumi.Input<string>;
    /**
     * Primary IP address of edge gateway. Read-only (can be specified in specific subnet)
     */
    primaryIp?: pulumi.Input<string>;
    /**
     * Optional ID of starting VDC if the 'owner_id' is a VDC Group
     */
    startingVdcId?: pulumi.Input<string>;
    /**
     * One or more blocks with external network information to be attached to this gateway's interface
     */
    subnets?: pulumi.Input<pulumi.Input<inputs.NsxtEdgegatewaySubnet>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtEdgegateway resource.
 */
export interface NsxtEdgegatewayArgs {
    /**
     * Dedicating the External Network will enable Route Advertisement for this Edge Gateway.
     */
    dedicateExternalNetwork?: pulumi.Input<boolean>;
    /**
     * Edge Gateway description
     */
    description?: pulumi.Input<string>;
    /**
     * Select specific NSX-T Edge Cluster. Will be inherited from external network if not specified
     */
    edgeClusterId?: pulumi.Input<string>;
    /**
     * External network ID
     */
    externalNetworkId: pulumi.Input<string>;
    /**
     * Edge Gateway name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * ID of VDC or VDC Group
     */
    ownerId?: pulumi.Input<string>;
    /**
     * Optional ID of starting VDC if the 'owner_id' is a VDC Group
     */
    startingVdcId?: pulumi.Input<string>;
    /**
     * One or more blocks with external network information to be attached to this gateway's interface
     */
    subnets: pulumi.Input<pulumi.Input<inputs.NsxtEdgegatewaySubnet>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}