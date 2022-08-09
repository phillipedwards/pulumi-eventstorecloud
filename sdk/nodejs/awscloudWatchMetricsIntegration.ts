// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AWSCloudWatchMetricsIntegration extends pulumi.CustomResource {
    /**
     * Get an existing AWSCloudWatchMetricsIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AWSCloudWatchMetricsIntegrationState, opts?: pulumi.CustomResourceOptions): AWSCloudWatchMetricsIntegration {
        return new AWSCloudWatchMetricsIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'eventstorecloud:index/aWSCloudWatchMetricsIntegration:AWSCloudWatchMetricsIntegration';

    /**
     * Returns true if the given object is an instance of AWSCloudWatchMetricsIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AWSCloudWatchMetricsIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AWSCloudWatchMetricsIntegration.__pulumiType;
    }

    /**
     * AWS IAM access key
     */
    public readonly accessKeyId!: pulumi.Output<string | undefined>;
    /**
     * Clusters to be used with this integration
     */
    public readonly clusterIds!: pulumi.Output<string[]>;
    /**
     * Human readable description of the integration
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Name of the CloudWatch namespace
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * ID of the project to which the integration applies
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * AWS region for group
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * AWS IAM secret access key
     */
    public readonly secretAccessKey!: pulumi.Output<string | undefined>;

    /**
     * Create a AWSCloudWatchMetricsIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AWSCloudWatchMetricsIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AWSCloudWatchMetricsIntegrationArgs | AWSCloudWatchMetricsIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AWSCloudWatchMetricsIntegrationState | undefined;
            resourceInputs["accessKeyId"] = state ? state.accessKeyId : undefined;
            resourceInputs["clusterIds"] = state ? state.clusterIds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretAccessKey"] = state ? state.secretAccessKey : undefined;
        } else {
            const args = argsOrState as AWSCloudWatchMetricsIntegrationArgs | undefined;
            if ((!args || args.clusterIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterIds'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["accessKeyId"] = args ? args.accessKeyId : undefined;
            resourceInputs["clusterIds"] = args ? args.clusterIds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretAccessKey"] = args ? args.secretAccessKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AWSCloudWatchMetricsIntegration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AWSCloudWatchMetricsIntegration resources.
 */
export interface AWSCloudWatchMetricsIntegrationState {
    /**
     * AWS IAM access key
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * Clusters to be used with this integration
     */
    clusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human readable description of the integration
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the CloudWatch namespace
     */
    namespace?: pulumi.Input<string>;
    /**
     * ID of the project to which the integration applies
     */
    projectId?: pulumi.Input<string>;
    /**
     * AWS region for group
     */
    region?: pulumi.Input<string>;
    /**
     * AWS IAM secret access key
     */
    secretAccessKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AWSCloudWatchMetricsIntegration resource.
 */
export interface AWSCloudWatchMetricsIntegrationArgs {
    /**
     * AWS IAM access key
     */
    accessKeyId?: pulumi.Input<string>;
    /**
     * Clusters to be used with this integration
     */
    clusterIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human readable description of the integration
     */
    description: pulumi.Input<string>;
    /**
     * Name of the CloudWatch namespace
     */
    namespace: pulumi.Input<string>;
    /**
     * ID of the project to which the integration applies
     */
    projectId: pulumi.Input<string>;
    /**
     * AWS region for group
     */
    region: pulumi.Input<string>;
    /**
     * AWS IAM secret access key
     */
    secretAccessKey?: pulumi.Input<string>;
}