// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LbServiceMonitor extends pulumi.CustomResource {
    /**
     * Get an existing LbServiceMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbServiceMonitorState, opts?: pulumi.CustomResourceOptions): LbServiceMonitor {
        return new LbServiceMonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/lbServiceMonitor:LbServiceMonitor';

    /**
     * Returns true if the given object is an instance of LbServiceMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbServiceMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbServiceMonitor.__pulumiType;
    }

    /**
     * Edge gateway name in which the LB Service Monitor is located
     */
    public readonly edgeGateway!: pulumi.Output<string>;
    /**
     * String that the monitor expects to match in the status line of the http or https response (for example, HTTP/1.1)
     */
    public readonly expected!: pulumi.Output<string | undefined>;
    /**
     * Advanced monitor parameters as key=value pairs
     */
    public readonly extension!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Interval in seconds at which a server is to be monitored (defaults to 10)
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * Number of times the specified monitoring Method must fail sequentially before the server is declared down (defaults to
     * 3)
     */
    public readonly maxRetries!: pulumi.Output<number | undefined>;
    /**
     * Method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
     */
    public readonly method!: pulumi.Output<string | undefined>;
    /**
     * Unique LB Service Monitor name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * String to be matched in the response content
     */
    public readonly receive!: pulumi.Output<string | undefined>;
    /**
     * Data to be sent
     */
    public readonly send!: pulumi.Output<string | undefined>;
    /**
     * Maximum time in seconds within which a response from the server must be received (defaults to 15)
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Way in which you want to send the health check request to the server. One of http, https, tcp, icmp, or udp
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * URL to be used in the server status request
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a LbServiceMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbServiceMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbServiceMonitorArgs | LbServiceMonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbServiceMonitorState | undefined;
            resourceInputs["edgeGateway"] = state ? state.edgeGateway : undefined;
            resourceInputs["expected"] = state ? state.expected : undefined;
            resourceInputs["extension"] = state ? state.extension : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["maxRetries"] = state ? state.maxRetries : undefined;
            resourceInputs["method"] = state ? state.method : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["receive"] = state ? state.receive : undefined;
            resourceInputs["send"] = state ? state.send : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as LbServiceMonitorArgs | undefined;
            if ((!args || args.edgeGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGateway'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["edgeGateway"] = args ? args.edgeGateway : undefined;
            resourceInputs["expected"] = args ? args.expected : undefined;
            resourceInputs["extension"] = args ? args.extension : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["maxRetries"] = args ? args.maxRetries : undefined;
            resourceInputs["method"] = args ? args.method : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["receive"] = args ? args.receive : undefined;
            resourceInputs["send"] = args ? args.send : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbServiceMonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbServiceMonitor resources.
 */
export interface LbServiceMonitorState {
    /**
     * Edge gateway name in which the LB Service Monitor is located
     */
    edgeGateway?: pulumi.Input<string>;
    /**
     * String that the monitor expects to match in the status line of the http or https response (for example, HTTP/1.1)
     */
    expected?: pulumi.Input<string>;
    /**
     * Advanced monitor parameters as key=value pairs
     */
    extension?: pulumi.Input<{[key: string]: any}>;
    /**
     * Interval in seconds at which a server is to be monitored (defaults to 10)
     */
    interval?: pulumi.Input<number>;
    /**
     * Number of times the specified monitoring Method must fail sequentially before the server is declared down (defaults to
     * 3)
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
     */
    method?: pulumi.Input<string>;
    /**
     * Unique LB Service Monitor name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * String to be matched in the response content
     */
    receive?: pulumi.Input<string>;
    /**
     * Data to be sent
     */
    send?: pulumi.Input<string>;
    /**
     * Maximum time in seconds within which a response from the server must be received (defaults to 15)
     */
    timeout?: pulumi.Input<number>;
    /**
     * Way in which you want to send the health check request to the server. One of http, https, tcp, icmp, or udp
     */
    type?: pulumi.Input<string>;
    /**
     * URL to be used in the server status request
     */
    url?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbServiceMonitor resource.
 */
export interface LbServiceMonitorArgs {
    /**
     * Edge gateway name in which the LB Service Monitor is located
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * String that the monitor expects to match in the status line of the http or https response (for example, HTTP/1.1)
     */
    expected?: pulumi.Input<string>;
    /**
     * Advanced monitor parameters as key=value pairs
     */
    extension?: pulumi.Input<{[key: string]: any}>;
    /**
     * Interval in seconds at which a server is to be monitored (defaults to 10)
     */
    interval?: pulumi.Input<number>;
    /**
     * Number of times the specified monitoring Method must fail sequentially before the server is declared down (defaults to
     * 3)
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
     */
    method?: pulumi.Input<string>;
    /**
     * Unique LB Service Monitor name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * String to be matched in the response content
     */
    receive?: pulumi.Input<string>;
    /**
     * Data to be sent
     */
    send?: pulumi.Input<string>;
    /**
     * Maximum time in seconds within which a response from the server must be received (defaults to 15)
     */
    timeout?: pulumi.Input<number>;
    /**
     * Way in which you want to send the health check request to the server. One of http, https, tcp, icmp, or udp
     */
    type: pulumi.Input<string>;
    /**
     * URL to be used in the server status request
     */
    url?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}