// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./getProject";
export * from "./integration";
export * from "./managedCluster";
export * from "./network";
export * from "./peering";
export * from "./project";
export * from "./provider";
export * from "./scheduledBackup";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Integration } from "./integration";
import { ManagedCluster } from "./managedCluster";
import { Network } from "./network";
import { Peering } from "./peering";
import { Project } from "./project";
import { ScheduledBackup } from "./scheduledBackup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "eventstorecloud:index/integration:Integration":
                return new Integration(name, <any>undefined, { urn })
            case "eventstorecloud:index/managedCluster:ManagedCluster":
                return new ManagedCluster(name, <any>undefined, { urn })
            case "eventstorecloud:index/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "eventstorecloud:index/peering:Peering":
                return new Peering(name, <any>undefined, { urn })
            case "eventstorecloud:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "eventstorecloud:index/scheduledBackup:ScheduledBackup":
                return new ScheduledBackup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("eventstorecloud", "index/integration", _module)
pulumi.runtime.registerResourceModule("eventstorecloud", "index/managedCluster", _module)
pulumi.runtime.registerResourceModule("eventstorecloud", "index/network", _module)
pulumi.runtime.registerResourceModule("eventstorecloud", "index/peering", _module)
pulumi.runtime.registerResourceModule("eventstorecloud", "index/project", _module)
pulumi.runtime.registerResourceModule("eventstorecloud", "index/scheduledBackup", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("eventstorecloud", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:eventstorecloud") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});