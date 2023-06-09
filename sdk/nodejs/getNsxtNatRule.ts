// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getNsxtNatRule(args: GetNsxtNatRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtNatRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxtNatRule:getNsxtNatRule", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtNatRule.
 */
export interface GetNsxtNatRuleArgs {
    edgeGatewayId: string;
    name: string;
    org?: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtNatRule.
 */
export interface GetNsxtNatRuleResult {
    readonly appPortProfileId: string;
    readonly description: string;
    readonly dnatExternalPort: string;
    readonly edgeGatewayId: string;
    readonly enabled: boolean;
    readonly externalAddress: string;
    readonly firewallMatch: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internalAddress: string;
    readonly logging: boolean;
    readonly name: string;
    readonly org?: string;
    readonly priority: number;
    readonly ruleType: string;
    readonly snatDestinationAddress: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    readonly vdc?: string;
}

export function getNsxtNatRuleOutput(args: GetNsxtNatRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxtNatRuleResult> {
    return pulumi.output(args).apply(a => getNsxtNatRule(a, opts))
}

/**
 * A collection of arguments for invoking getNsxtNatRule.
 */
export interface GetNsxtNatRuleOutputArgs {
    edgeGatewayId: pulumi.Input<string>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}
