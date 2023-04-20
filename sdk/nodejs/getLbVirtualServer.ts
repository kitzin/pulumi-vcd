// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getLbVirtualServer(args: GetLbVirtualServerArgs, opts?: pulumi.InvokeOptions): Promise<GetLbVirtualServerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getLbVirtualServer:getLbVirtualServer", {
        "edgeGateway": args.edgeGateway,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbVirtualServer.
 */
export interface GetLbVirtualServerArgs {
    edgeGateway: string;
    name: string;
    org?: string;
    vdc?: string;
}

/**
 * A collection of values returned by getLbVirtualServer.
 */
export interface GetLbVirtualServerResult {
    readonly appProfileId: string;
    readonly appRuleIds: string[];
    readonly connectionLimit: number;
    readonly connectionRateLimit: number;
    readonly description: string;
    readonly edgeGateway: string;
    readonly enableAcceleration: boolean;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddress: string;
    readonly name: string;
    readonly org?: string;
    readonly port: number;
    readonly protocol: string;
    readonly serverPoolId: string;
    readonly vdc?: string;
}

export function getLbVirtualServerOutput(args: GetLbVirtualServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbVirtualServerResult> {
    return pulumi.output(args).apply(a => getLbVirtualServer(a, opts))
}

/**
 * A collection of arguments for invoking getLbVirtualServer.
 */
export interface GetLbVirtualServerOutputArgs {
    edgeGateway: pulumi.Input<string>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}
