// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getEdgegateway(args?: GetEdgegatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetEdgegatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getEdgegateway:getEdgegateway", {
        "filter": args.filter,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getEdgegateway.
 */
export interface GetEdgegatewayArgs {
    filter?: inputs.GetEdgegatewayFilter;
    name?: string;
    org?: string;
    vdc?: string;
}

/**
 * A collection of values returned by getEdgegateway.
 */
export interface GetEdgegatewayResult {
    readonly configuration: string;
    readonly defaultExternalNetworkIp: string;
    readonly description: string;
    readonly distributedRouting: boolean;
    readonly externalNetworkIps: string[];
    readonly externalNetworks: outputs.GetEdgegatewayExternalNetwork[];
    readonly filter?: outputs.GetEdgegatewayFilter;
    readonly fipsModeEnabled: boolean;
    readonly fwDefaultRuleAction: string;
    readonly fwDefaultRuleLoggingEnabled: boolean;
    readonly fwEnabled: boolean;
    readonly haEnabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lbAccelerationEnabled: boolean;
    readonly lbEnabled: boolean;
    readonly lbLoggingEnabled: boolean;
    readonly lbLoglevel: string;
    readonly name?: string;
    readonly org?: string;
    readonly useDefaultRouteForDnsRelay: boolean;
    readonly vdc?: string;
}

export function getEdgegatewayOutput(args?: GetEdgegatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEdgegatewayResult> {
    return pulumi.output(args).apply(a => getEdgegateway(a, opts))
}

/**
 * A collection of arguments for invoking getEdgegateway.
 */
export interface GetEdgegatewayOutputArgs {
    filter?: pulumi.Input<inputs.GetEdgegatewayFilterArgs>;
    name?: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}
