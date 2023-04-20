// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NsxtEdgegatewayBgpConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing NsxtEdgegatewayBgpConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtEdgegatewayBgpConfigurationState, opts?: pulumi.CustomResourceOptions): NsxtEdgegatewayBgpConfiguration {
        return new NsxtEdgegatewayBgpConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtEdgegatewayBgpConfiguration:NsxtEdgegatewayBgpConfiguration';

    /**
     * Returns true if the given object is an instance of NsxtEdgegatewayBgpConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtEdgegatewayBgpConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtEdgegatewayBgpConfiguration.__pulumiType;
    }

    /**
     * Defines if ECMP (Equal-cost multi-path routing) is enabled
     */
    public readonly ecmpEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Edge gateway ID for BGP Configuration
     */
    public readonly edgeGatewayId!: pulumi.Output<string>;
    /**
     * Defines if BGP service is enabled
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Graceful restart configuration on Edge Gateway. One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
     */
    public readonly gracefulRestartMode!: pulumi.Output<string>;
    /**
     * Maximum time taken (in seconds) for a BGP session to be established after a restart
     */
    public readonly gracefulRestartTimer!: pulumi.Output<number>;
    /**
     * Autonomous system number
     */
    public readonly localAsNumber!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Maximum time (in seconds) before stale routes are removed when BGP restarts
     */
    public readonly staleRouteTimer!: pulumi.Output<number>;

    /**
     * Create a NsxtEdgegatewayBgpConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtEdgegatewayBgpConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtEdgegatewayBgpConfigurationArgs | NsxtEdgegatewayBgpConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtEdgegatewayBgpConfigurationState | undefined;
            resourceInputs["ecmpEnabled"] = state ? state.ecmpEnabled : undefined;
            resourceInputs["edgeGatewayId"] = state ? state.edgeGatewayId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["gracefulRestartMode"] = state ? state.gracefulRestartMode : undefined;
            resourceInputs["gracefulRestartTimer"] = state ? state.gracefulRestartTimer : undefined;
            resourceInputs["localAsNumber"] = state ? state.localAsNumber : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["staleRouteTimer"] = state ? state.staleRouteTimer : undefined;
        } else {
            const args = argsOrState as NsxtEdgegatewayBgpConfigurationArgs | undefined;
            if ((!args || args.edgeGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGatewayId'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            resourceInputs["ecmpEnabled"] = args ? args.ecmpEnabled : undefined;
            resourceInputs["edgeGatewayId"] = args ? args.edgeGatewayId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["gracefulRestartMode"] = args ? args.gracefulRestartMode : undefined;
            resourceInputs["gracefulRestartTimer"] = args ? args.gracefulRestartTimer : undefined;
            resourceInputs["localAsNumber"] = args ? args.localAsNumber : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["staleRouteTimer"] = args ? args.staleRouteTimer : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtEdgegatewayBgpConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtEdgegatewayBgpConfiguration resources.
 */
export interface NsxtEdgegatewayBgpConfigurationState {
    /**
     * Defines if ECMP (Equal-cost multi-path routing) is enabled
     */
    ecmpEnabled?: pulumi.Input<boolean>;
    /**
     * Edge gateway ID for BGP Configuration
     */
    edgeGatewayId?: pulumi.Input<string>;
    /**
     * Defines if BGP service is enabled
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Graceful restart configuration on Edge Gateway. One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
     */
    gracefulRestartMode?: pulumi.Input<string>;
    /**
     * Maximum time taken (in seconds) for a BGP session to be established after a restart
     */
    gracefulRestartTimer?: pulumi.Input<number>;
    /**
     * Autonomous system number
     */
    localAsNumber?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Maximum time (in seconds) before stale routes are removed when BGP restarts
     */
    staleRouteTimer?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NsxtEdgegatewayBgpConfiguration resource.
 */
export interface NsxtEdgegatewayBgpConfigurationArgs {
    /**
     * Defines if ECMP (Equal-cost multi-path routing) is enabled
     */
    ecmpEnabled?: pulumi.Input<boolean>;
    /**
     * Edge gateway ID for BGP Configuration
     */
    edgeGatewayId: pulumi.Input<string>;
    /**
     * Defines if BGP service is enabled
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Graceful restart configuration on Edge Gateway. One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
     */
    gracefulRestartMode?: pulumi.Input<string>;
    /**
     * Maximum time taken (in seconds) for a BGP session to be established after a restart
     */
    gracefulRestartTimer?: pulumi.Input<number>;
    /**
     * Autonomous system number
     */
    localAsNumber?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Maximum time (in seconds) before stale routes are removed when BGP restarts
     */
    staleRouteTimer?: pulumi.Input<number>;
}