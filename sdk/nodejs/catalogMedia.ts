// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class CatalogMedia extends pulumi.CustomResource {
    /**
     * Get an existing CatalogMedia resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogMediaState, opts?: pulumi.CustomResourceOptions): CatalogMedia {
        return new CatalogMedia(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/catalogMedia:CatalogMedia';

    /**
     * Returns true if the given object is an instance of CatalogMedia.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogMedia {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogMedia.__pulumiType;
    }

    /**
     * catalog name where to upload the Media file
     *
     * @deprecated Use catalog_id instead
     */
    public readonly catalog!: pulumi.Output<string>;
    /**
     * ID of the catalog where to upload the Media file
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * Creation date
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * True if this media file is ISO
     */
    public /*out*/ readonly isIso!: pulumi.Output<boolean>;
    /**
     * True if this media file is in a published catalog
     */
    public /*out*/ readonly isPublished!: pulumi.Output<boolean>;
    /**
     * absolute or relative path to Media file
     */
    public readonly mediaPath!: pulumi.Output<string>;
    /**
     * Key and value pairs for catalog item metadata
     *
     * @deprecated Use metadata_entry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * Metadata entries for the given Catalog Media
     */
    public readonly metadataEntries!: pulumi.Output<outputs.CatalogMediaMetadataEntry[]>;
    /**
     * media name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string>;
    /**
     * Owner name
     */
    public /*out*/ readonly ownerName!: pulumi.Output<string>;
    /**
     * shows upload progress in stdout
     */
    public readonly showUploadProgress!: pulumi.Output<boolean | undefined>;
    /**
     * Media storage in Bytes
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Media status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Storage profile name
     */
    public /*out*/ readonly storageProfileName!: pulumi.Output<string>;
    /**
     * size of upload file piece size in mega bytes
     */
    public readonly uploadPieceSize!: pulumi.Output<number | undefined>;

    /**
     * Create a CatalogMedia resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogMediaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogMediaArgs | CatalogMediaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogMediaState | undefined;
            resourceInputs["catalog"] = state ? state.catalog : undefined;
            resourceInputs["catalogId"] = state ? state.catalogId : undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isIso"] = state ? state.isIso : undefined;
            resourceInputs["isPublished"] = state ? state.isPublished : undefined;
            resourceInputs["mediaPath"] = state ? state.mediaPath : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ownerName"] = state ? state.ownerName : undefined;
            resourceInputs["showUploadProgress"] = state ? state.showUploadProgress : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageProfileName"] = state ? state.storageProfileName : undefined;
            resourceInputs["uploadPieceSize"] = state ? state.uploadPieceSize : undefined;
        } else {
            const args = argsOrState as CatalogMediaArgs | undefined;
            if ((!args || args.mediaPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mediaPath'");
            }
            resourceInputs["catalog"] = args ? args.catalog : undefined;
            resourceInputs["catalogId"] = args ? args.catalogId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["mediaPath"] = args ? args.mediaPath : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["showUploadProgress"] = args ? args.showUploadProgress : undefined;
            resourceInputs["uploadPieceSize"] = args ? args.uploadPieceSize : undefined;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["isIso"] = undefined /*out*/;
            resourceInputs["isPublished"] = undefined /*out*/;
            resourceInputs["ownerName"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageProfileName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CatalogMedia.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogMedia resources.
 */
export interface CatalogMediaState {
    /**
     * catalog name where to upload the Media file
     *
     * @deprecated Use catalog_id instead
     */
    catalog?: pulumi.Input<string>;
    /**
     * ID of the catalog where to upload the Media file
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Creation date
     */
    creationDate?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * True if this media file is ISO
     */
    isIso?: pulumi.Input<boolean>;
    /**
     * True if this media file is in a published catalog
     */
    isPublished?: pulumi.Input<boolean>;
    /**
     * absolute or relative path to Media file
     */
    mediaPath?: pulumi.Input<string>;
    /**
     * Key and value pairs for catalog item metadata
     *
     * @deprecated Use metadata_entry instead
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given Catalog Media
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogMediaMetadataEntry>[]>;
    /**
     * media name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Owner name
     */
    ownerName?: pulumi.Input<string>;
    /**
     * shows upload progress in stdout
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * Media storage in Bytes
     */
    size?: pulumi.Input<number>;
    /**
     * Media status
     */
    status?: pulumi.Input<string>;
    /**
     * Storage profile name
     */
    storageProfileName?: pulumi.Input<string>;
    /**
     * size of upload file piece size in mega bytes
     */
    uploadPieceSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CatalogMedia resource.
 */
export interface CatalogMediaArgs {
    /**
     * catalog name where to upload the Media file
     *
     * @deprecated Use catalog_id instead
     */
    catalog?: pulumi.Input<string>;
    /**
     * ID of the catalog where to upload the Media file
     */
    catalogId?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * absolute or relative path to Media file
     */
    mediaPath: pulumi.Input<string>;
    /**
     * Key and value pairs for catalog item metadata
     *
     * @deprecated Use metadata_entry instead
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given Catalog Media
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogMediaMetadataEntry>[]>;
    /**
     * media name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * shows upload progress in stdout
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * size of upload file piece size in mega bytes
     */
    uploadPieceSize?: pulumi.Input<number>;
}
