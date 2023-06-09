// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getSubscribedCatalog(args: GetSubscribedCatalogArgs, opts?: pulumi.InvokeOptions): Promise<GetSubscribedCatalogResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getSubscribedCatalog:getSubscribedCatalog", {
        "filter": args.filter,
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubscribedCatalog.
 */
export interface GetSubscribedCatalogArgs {
    filter?: inputs.GetSubscribedCatalogFilter;
    name: string;
    org?: string;
}

/**
 * A collection of values returned by getSubscribedCatalog.
 */
export interface GetSubscribedCatalogResult {
    readonly catalogVersion: number;
    readonly created: string;
    readonly description: string;
    readonly failedTasks: string[];
    readonly filter?: outputs.GetSubscribedCatalogFilter;
    readonly href: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isLocal: boolean;
    readonly isPublished: boolean;
    readonly isShared: boolean;
    readonly makeLocalCopy: boolean;
    readonly mediaItemLists: string[];
    readonly name: string;
    readonly numberOfMedia: number;
    readonly numberOfVappTemplates: number;
    readonly org?: string;
    readonly ownerName: string;
    readonly publishSubscriptionType: string;
    readonly runningTasks: string[];
    readonly storageProfileId: string;
    readonly subscriptionUrl: string;
    readonly vappTemplateLists: string[];
}

export function getSubscribedCatalogOutput(args: GetSubscribedCatalogOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubscribedCatalogResult> {
    return pulumi.output(args).apply(a => getSubscribedCatalog(a, opts))
}

/**
 * A collection of arguments for invoking getSubscribedCatalog.
 */
export interface GetSubscribedCatalogOutputArgs {
    filter?: pulumi.Input<inputs.GetSubscribedCatalogFilterArgs>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
}
