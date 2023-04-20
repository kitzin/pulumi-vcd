// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NsxtEdgegatewayBgpNeighbor extends pulumi.CustomResource {
    /**
     * Get an existing NsxtEdgegatewayBgpNeighbor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtEdgegatewayBgpNeighborState, opts?: pulumi.CustomResourceOptions): NsxtEdgegatewayBgpNeighbor {
        return new NsxtEdgegatewayBgpNeighbor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtEdgegatewayBgpNeighbor:NsxtEdgegatewayBgpNeighbor';

    /**
     * Returns true if the given object is an instance of NsxtEdgegatewayBgpNeighbor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtEdgegatewayBgpNeighbor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtEdgegatewayBgpNeighbor.__pulumiType;
    }

    /**
     * A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false')
     */
    public readonly allowAsIn!: pulumi.Output<boolean>;
    /**
     * Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
     */
    public readonly bfdDeadMultiple!: pulumi.Output<number>;
    /**
     * BFD configuration for failure detection
     */
    public readonly bfdEnabled!: pulumi.Output<boolean>;
    /**
     * Time interval (in milliseconds) between heartbeat packets
     */
    public readonly bfdInterval!: pulumi.Output<number>;
    /**
     * Edge gateway ID for BGP Neighbor Configuration
     */
    public readonly edgeGatewayId!: pulumi.Output<string>;
    /**
     * One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
     */
    public readonly gracefulRestartMode!: pulumi.Output<string>;
    /**
     * Time interval (in seconds) before declaring a peer dead
     */
    public readonly holdDownTimer!: pulumi.Output<number>;
    /**
     * An optional IP Prefix List ID for filtering 'IN' direction.
     */
    public readonly inFilterIpPrefixListId!: pulumi.Output<string>;
    /**
     * BGP Neighbor IP address (IPv4 or IPv6)
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * Time interval (in seconds) between sending keep alive messages to a peer
     */
    public readonly keepAliveTimer!: pulumi.Output<number>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * An optional IP Prefix List ID for filtering 'OUT' direction.
     */
    public readonly outFilterIpPrefixListId!: pulumi.Output<string>;
    /**
     * Neighbor password
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Remote Autonomous System (AS) number
     */
    public readonly remoteAsNumber!: pulumi.Output<string>;
    /**
     * One of 'DISABLED', 'IPV4', 'IPV6'
     */
    public readonly routeFiltering!: pulumi.Output<string>;

    /**
     * Create a NsxtEdgegatewayBgpNeighbor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtEdgegatewayBgpNeighborArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtEdgegatewayBgpNeighborArgs | NsxtEdgegatewayBgpNeighborState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtEdgegatewayBgpNeighborState | undefined;
            resourceInputs["allowAsIn"] = state ? state.allowAsIn : undefined;
            resourceInputs["bfdDeadMultiple"] = state ? state.bfdDeadMultiple : undefined;
            resourceInputs["bfdEnabled"] = state ? state.bfdEnabled : undefined;
            resourceInputs["bfdInterval"] = state ? state.bfdInterval : undefined;
            resourceInputs["edgeGatewayId"] = state ? state.edgeGatewayId : undefined;
            resourceInputs["gracefulRestartMode"] = state ? state.gracefulRestartMode : undefined;
            resourceInputs["holdDownTimer"] = state ? state.holdDownTimer : undefined;
            resourceInputs["inFilterIpPrefixListId"] = state ? state.inFilterIpPrefixListId : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["keepAliveTimer"] = state ? state.keepAliveTimer : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["outFilterIpPrefixListId"] = state ? state.outFilterIpPrefixListId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["remoteAsNumber"] = state ? state.remoteAsNumber : undefined;
            resourceInputs["routeFiltering"] = state ? state.routeFiltering : undefined;
        } else {
            const args = argsOrState as NsxtEdgegatewayBgpNeighborArgs | undefined;
            if ((!args || args.edgeGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGatewayId'");
            }
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            if ((!args || args.remoteAsNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteAsNumber'");
            }
            resourceInputs["allowAsIn"] = args ? args.allowAsIn : undefined;
            resourceInputs["bfdDeadMultiple"] = args ? args.bfdDeadMultiple : undefined;
            resourceInputs["bfdEnabled"] = args ? args.bfdEnabled : undefined;
            resourceInputs["bfdInterval"] = args ? args.bfdInterval : undefined;
            resourceInputs["edgeGatewayId"] = args ? args.edgeGatewayId : undefined;
            resourceInputs["gracefulRestartMode"] = args ? args.gracefulRestartMode : undefined;
            resourceInputs["holdDownTimer"] = args ? args.holdDownTimer : undefined;
            resourceInputs["inFilterIpPrefixListId"] = args ? args.inFilterIpPrefixListId : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["keepAliveTimer"] = args ? args.keepAliveTimer : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["outFilterIpPrefixListId"] = args ? args.outFilterIpPrefixListId : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["remoteAsNumber"] = args ? args.remoteAsNumber : undefined;
            resourceInputs["routeFiltering"] = args ? args.routeFiltering : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtEdgegatewayBgpNeighbor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtEdgegatewayBgpNeighbor resources.
 */
export interface NsxtEdgegatewayBgpNeighborState {
    /**
     * A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false')
     */
    allowAsIn?: pulumi.Input<boolean>;
    /**
     * Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
     */
    bfdDeadMultiple?: pulumi.Input<number>;
    /**
     * BFD configuration for failure detection
     */
    bfdEnabled?: pulumi.Input<boolean>;
    /**
     * Time interval (in milliseconds) between heartbeat packets
     */
    bfdInterval?: pulumi.Input<number>;
    /**
     * Edge gateway ID for BGP Neighbor Configuration
     */
    edgeGatewayId?: pulumi.Input<string>;
    /**
     * One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
     */
    gracefulRestartMode?: pulumi.Input<string>;
    /**
     * Time interval (in seconds) before declaring a peer dead
     */
    holdDownTimer?: pulumi.Input<number>;
    /**
     * An optional IP Prefix List ID for filtering 'IN' direction.
     */
    inFilterIpPrefixListId?: pulumi.Input<string>;
    /**
     * BGP Neighbor IP address (IPv4 or IPv6)
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Time interval (in seconds) between sending keep alive messages to a peer
     */
    keepAliveTimer?: pulumi.Input<number>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * An optional IP Prefix List ID for filtering 'OUT' direction.
     */
    outFilterIpPrefixListId?: pulumi.Input<string>;
    /**
     * Neighbor password
     */
    password?: pulumi.Input<string>;
    /**
     * Remote Autonomous System (AS) number
     */
    remoteAsNumber?: pulumi.Input<string>;
    /**
     * One of 'DISABLED', 'IPV4', 'IPV6'
     */
    routeFiltering?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtEdgegatewayBgpNeighbor resource.
 */
export interface NsxtEdgegatewayBgpNeighborArgs {
    /**
     * A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false')
     */
    allowAsIn?: pulumi.Input<boolean>;
    /**
     * Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
     */
    bfdDeadMultiple?: pulumi.Input<number>;
    /**
     * BFD configuration for failure detection
     */
    bfdEnabled?: pulumi.Input<boolean>;
    /**
     * Time interval (in milliseconds) between heartbeat packets
     */
    bfdInterval?: pulumi.Input<number>;
    /**
     * Edge gateway ID for BGP Neighbor Configuration
     */
    edgeGatewayId: pulumi.Input<string>;
    /**
     * One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
     */
    gracefulRestartMode?: pulumi.Input<string>;
    /**
     * Time interval (in seconds) before declaring a peer dead
     */
    holdDownTimer?: pulumi.Input<number>;
    /**
     * An optional IP Prefix List ID for filtering 'IN' direction.
     */
    inFilterIpPrefixListId?: pulumi.Input<string>;
    /**
     * BGP Neighbor IP address (IPv4 or IPv6)
     */
    ipAddress: pulumi.Input<string>;
    /**
     * Time interval (in seconds) between sending keep alive messages to a peer
     */
    keepAliveTimer?: pulumi.Input<number>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * An optional IP Prefix List ID for filtering 'OUT' direction.
     */
    outFilterIpPrefixListId?: pulumi.Input<string>;
    /**
     * Neighbor password
     */
    password?: pulumi.Input<string>;
    /**
     * Remote Autonomous System (AS) number
     */
    remoteAsNumber: pulumi.Input<string>;
    /**
     * One of 'DISABLED', 'IPV4', 'IPV6'
     */
    routeFiltering?: pulumi.Input<string>;
}