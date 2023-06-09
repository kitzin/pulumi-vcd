// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNsxtFirewall(args: GetNsxtFirewallArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtFirewallResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxtFirewall:getNsxtFirewall", {
        "edgeGatewayId": args.edgeGatewayId,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtFirewall.
 */
export interface GetNsxtFirewallArgs {
    edgeGatewayId: string;
    org?: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtFirewall.
 */
export interface GetNsxtFirewallResult {
    readonly edgeGatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly org?: string;
    readonly rules: outputs.GetNsxtFirewallRule[];
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    readonly vdc?: string;
}

export function getNsxtFirewallOutput(args: GetNsxtFirewallOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxtFirewallResult> {
    return pulumi.output(args).apply(a => getNsxtFirewall(a, opts))
}

/**
 * A collection of arguments for invoking getNsxtFirewall.
 */
export interface GetNsxtFirewallOutputArgs {
    edgeGatewayId: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}
