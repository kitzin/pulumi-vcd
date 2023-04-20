// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNsxtEdgegateway(args: GetNsxtEdgegatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtEdgegatewayResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxtEdgegateway:getNsxtEdgegateway", {
        "edgeClusterId": args.edgeClusterId,
        "name": args.name,
        "org": args.org,
        "ownerId": args.ownerId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtEdgegateway.
 */
export interface GetNsxtEdgegatewayArgs {
    edgeClusterId?: string;
    name: string;
    org?: string;
    ownerId?: string;
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtEdgegateway.
 */
export interface GetNsxtEdgegatewayResult {
    readonly dedicateExternalNetwork: boolean;
    readonly description: string;
    readonly edgeClusterId?: string;
    readonly externalNetworkId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly org?: string;
    readonly ownerId: string;
    readonly primaryIp: string;
    readonly subnets: outputs.GetNsxtEdgegatewaySubnet[];
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    readonly vdc: string;
}

export function getNsxtEdgegatewayOutput(args: GetNsxtEdgegatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxtEdgegatewayResult> {
    return pulumi.output(args).apply(a => getNsxtEdgegateway(a, opts))
}

/**
 * A collection of arguments for invoking getNsxtEdgegateway.
 */
export interface GetNsxtEdgegatewayOutputArgs {
    edgeClusterId?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    ownerId?: pulumi.Input<string>;
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}