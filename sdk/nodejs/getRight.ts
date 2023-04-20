// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getRight(args: GetRightArgs, opts?: pulumi.InvokeOptions): Promise<GetRightResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getRight:getRight", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRight.
 */
export interface GetRightArgs {
    name: string;
}

/**
 * A collection of values returned by getRight.
 */
export interface GetRightResult {
    readonly bundleKey: string;
    readonly categoryId: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly impliedRights: outputs.GetRightImpliedRight[];
    readonly name: string;
    readonly rightType: string;
}

export function getRightOutput(args: GetRightOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRightResult> {
    return pulumi.output(args).apply(a => getRight(a, opts))
}

/**
 * A collection of arguments for invoking getRight.
 */
export interface GetRightOutputArgs {
    name: pulumi.Input<string>;
}
