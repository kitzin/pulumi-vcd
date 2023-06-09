// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getLbServerPool(args: GetLbServerPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetLbServerPoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getLbServerPool:getLbServerPool", {
        "edgeGateway": args.edgeGateway,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbServerPool.
 */
export interface GetLbServerPoolArgs {
    edgeGateway: string;
    name: string;
    org?: string;
    vdc?: string;
}

/**
 * A collection of values returned by getLbServerPool.
 */
export interface GetLbServerPoolResult {
    readonly algorithm: string;
    readonly algorithmParameters: string;
    readonly description: string;
    readonly edgeGateway: string;
    readonly enableTransparency: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly members: outputs.GetLbServerPoolMember[];
    readonly monitorId: string;
    readonly name: string;
    readonly org?: string;
    readonly vdc?: string;
}

export function getLbServerPoolOutput(args: GetLbServerPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbServerPoolResult> {
    return pulumi.output(args).apply(a => getLbServerPool(a, opts))
}

/**
 * A collection of arguments for invoking getLbServerPool.
 */
export interface GetLbServerPoolOutputArgs {
    edgeGateway: pulumi.Input<string>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}
