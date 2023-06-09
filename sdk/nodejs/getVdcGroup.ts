// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getVdcGroup(args?: GetVdcGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetVdcGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getVdcGroup:getVdcGroup", {
        "defaultPolicyStatus": args.defaultPolicyStatus,
        "description": args.description,
        "errorMessage": args.errorMessage,
        "id": args.id,
        "localEgress": args.localEgress,
        "name": args.name,
        "networkPoolId": args.networkPoolId,
        "networkPoolUniversalId": args.networkPoolUniversalId,
        "networkProviderType": args.networkProviderType,
        "org": args.org,
        "status": args.status,
        "type": args.type,
        "universalNetworkingEnabled": args.universalNetworkingEnabled,
    }, opts);
}

/**
 * A collection of arguments for invoking getVdcGroup.
 */
export interface GetVdcGroupArgs {
    defaultPolicyStatus?: boolean;
    description?: string;
    errorMessage?: string;
    id?: string;
    localEgress?: boolean;
    name?: string;
    networkPoolId?: string;
    networkPoolUniversalId?: string;
    networkProviderType?: string;
    org?: string;
    status?: string;
    type?: string;
    universalNetworkingEnabled?: boolean;
}

/**
 * A collection of values returned by getVdcGroup.
 */
export interface GetVdcGroupResult {
    readonly defaultPolicyStatus: boolean;
    readonly description: string;
    readonly dfwEnabled: boolean;
    readonly errorMessage: string;
    readonly id: string;
    readonly localEgress: boolean;
    readonly name: string;
    readonly networkPoolId: string;
    readonly networkPoolUniversalId: string;
    readonly networkProviderType: string;
    readonly org?: string;
    readonly participatingOrgVdcs: outputs.GetVdcGroupParticipatingOrgVdc[];
    readonly status: string;
    readonly type: string;
    readonly universalNetworkingEnabled: boolean;
}

export function getVdcGroupOutput(args?: GetVdcGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVdcGroupResult> {
    return pulumi.output(args).apply(a => getVdcGroup(a, opts))
}

/**
 * A collection of arguments for invoking getVdcGroup.
 */
export interface GetVdcGroupOutputArgs {
    defaultPolicyStatus?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    errorMessage?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    localEgress?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    networkPoolId?: pulumi.Input<string>;
    networkPoolUniversalId?: pulumi.Input<string>;
    networkProviderType?: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    universalNetworkingEnabled?: pulumi.Input<boolean>;
}
