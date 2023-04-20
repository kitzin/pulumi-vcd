// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNsxtSecurityGroup(args: GetNsxtSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtSecurityGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxtSecurityGroup:getNsxtSecurityGroup", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtSecurityGroup.
 */
export interface GetNsxtSecurityGroupArgs {
    edgeGatewayId: string;
    name: string;
    org?: string;
    /**
     * @deprecated Deprecated in favor of `edge_gateway_id`. Security Group will inherit VDC from parent Edge Gateway.
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtSecurityGroup.
 */
export interface GetNsxtSecurityGroupResult {
    readonly description: string;
    readonly edgeGatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly memberOrgNetworkIds: string[];
    readonly memberVms: outputs.GetNsxtSecurityGroupMemberVm[];
    readonly name: string;
    readonly org?: string;
    readonly ownerId: string;
    /**
     * @deprecated Deprecated in favor of `edge_gateway_id`. Security Group will inherit VDC from parent Edge Gateway.
     */
    readonly vdc?: string;
}

export function getNsxtSecurityGroupOutput(args: GetNsxtSecurityGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxtSecurityGroupResult> {
    return pulumi.output(args).apply(a => getNsxtSecurityGroup(a, opts))
}

/**
 * A collection of arguments for invoking getNsxtSecurityGroup.
 */
export interface GetNsxtSecurityGroupOutputArgs {
    edgeGatewayId: pulumi.Input<string>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favor of `edge_gateway_id`. Security Group will inherit VDC from parent Edge Gateway.
     */
    vdc?: pulumi.Input<string>;
}