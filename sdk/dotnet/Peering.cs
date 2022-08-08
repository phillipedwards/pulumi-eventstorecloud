// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EventStoreCloud
{
    /// <summary>
    /// Manages peering connections between Event Store Cloud VPCs and customer own VPCs
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using EventStoreCloud = Pulumi.EventStoreCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Example for AWS
    ///         var exampleProject = new EventStoreCloud.Project("exampleProject", new EventStoreCloud.ProjectArgs
    ///         {
    ///         });
    ///         var exampleNetwork = new EventStoreCloud.Network("exampleNetwork", new EventStoreCloud.NetworkArgs
    ///         {
    ///             ProjectId = exampleProject.Id,
    ///             ResourceProvider = "aws",
    ///             Region = "us-west-2",
    ///             CidrBlock = "172.21.0.0/16",
    ///         });
    ///         var examplePeering = new EventStoreCloud.Peering("examplePeering", new EventStoreCloud.PeeringArgs
    ///         {
    ///             ProjectId = exampleNetwork.ProjectId,
    ///             NetworkId = exampleNetwork.Id,
    ///             PeerResourceProvider = exampleNetwork.ResourceProvider,
    ///             PeerNetworkRegion = exampleNetwork.Region,
    ///             PeerAccountId = "&lt;Customer AWS Account ID&gt;",
    ///             PeerNetworkId = "&lt;Customer VPC ID&gt;",
    ///             Routes = 
    ///             {
    ///                 "&lt;Address space of the customer VPC&gt;",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import eventstorecloud:index/peering:Peering example project_id:peering_id
    /// ```
    /// 
    ///  ~&gt; Keep in mind that additional operations might be required to activate the peering link. Check our [provisioning guidelines](https://developers.eventstore.com/cloud/provision/) for each of the supported cloud providers to know more.
    /// </summary>
    [EventStoreCloudResourceType("eventstorecloud:index/peering:Peering")]
    public partial class Peering : Pulumi.CustomResource
    {
        /// <summary>
        /// Human-friendly name for the network
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network ID
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Account identifier in which to the peer network exists
        /// </summary>
        [Output("peerAccountId")]
        public Output<string> PeerAccountId { get; private set; } = null!;

        /// <summary>
        /// Network identifier of the peer network exists
        /// </summary>
        [Output("peerNetworkId")]
        public Output<string> PeerNetworkId { get; private set; } = null!;

        /// <summary>
        /// Provider region in which to the peer network exists
        /// </summary>
        [Output("peerNetworkRegion")]
        public Output<string> PeerNetworkRegion { get; private set; } = null!;

        /// <summary>
        /// Cloud Provider in which the target network exists
        /// </summary>
        [Output("peerResourceProvider")]
        public Output<string> PeerResourceProvider { get; private set; } = null!;

        /// <summary>
        /// Project ID
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Metadata about the remote end of the peering connection
        /// </summary>
        [Output("providerMetadata")]
        public Output<ImmutableDictionary<string, string>> ProviderMetadata { get; private set; } = null!;

        /// <summary>
        /// Routes to create from the Event Store network to the peer network
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<string>> Routes { get; private set; } = null!;


        /// <summary>
        /// Create a Peering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peering(string name, PeeringArgs args, CustomResourceOptions? options = null)
            : base("eventstorecloud:index/peering:Peering", name, args ?? new PeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peering(string name, Input<string> id, PeeringState? state = null, CustomResourceOptions? options = null)
            : base("eventstorecloud:index/peering:Peering", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/EventStore/pulumi-eventstorecloud/releases/download/0.2.6",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Peering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peering Get(string name, Input<string> id, PeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new Peering(name, id, state, options);
        }
    }

    public sealed class PeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-friendly name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network ID
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Account identifier in which to the peer network exists
        /// </summary>
        [Input("peerAccountId", required: true)]
        public Input<string> PeerAccountId { get; set; } = null!;

        /// <summary>
        /// Network identifier of the peer network exists
        /// </summary>
        [Input("peerNetworkId", required: true)]
        public Input<string> PeerNetworkId { get; set; } = null!;

        /// <summary>
        /// Provider region in which to the peer network exists
        /// </summary>
        [Input("peerNetworkRegion", required: true)]
        public Input<string> PeerNetworkRegion { get; set; } = null!;

        /// <summary>
        /// Cloud Provider in which the target network exists
        /// </summary>
        [Input("peerResourceProvider", required: true)]
        public Input<string> PeerResourceProvider { get; set; } = null!;

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("routes", required: true)]
        private InputList<string>? _routes;

        /// <summary>
        /// Routes to create from the Event Store network to the peer network
        /// </summary>
        public InputList<string> Routes
        {
            get => _routes ?? (_routes = new InputList<string>());
            set => _routes = value;
        }

        public PeeringArgs()
        {
        }
    }

    public sealed class PeeringState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-friendly name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network ID
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Account identifier in which to the peer network exists
        /// </summary>
        [Input("peerAccountId")]
        public Input<string>? PeerAccountId { get; set; }

        /// <summary>
        /// Network identifier of the peer network exists
        /// </summary>
        [Input("peerNetworkId")]
        public Input<string>? PeerNetworkId { get; set; }

        /// <summary>
        /// Provider region in which to the peer network exists
        /// </summary>
        [Input("peerNetworkRegion")]
        public Input<string>? PeerNetworkRegion { get; set; }

        /// <summary>
        /// Cloud Provider in which the target network exists
        /// </summary>
        [Input("peerResourceProvider")]
        public Input<string>? PeerResourceProvider { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("providerMetadata")]
        private InputMap<string>? _providerMetadata;

        /// <summary>
        /// Metadata about the remote end of the peering connection
        /// </summary>
        public InputMap<string> ProviderMetadata
        {
            get => _providerMetadata ?? (_providerMetadata = new InputMap<string>());
            set => _providerMetadata = value;
        }

        [Input("routes")]
        private InputList<string>? _routes;

        /// <summary>
        /// Routes to create from the Event Store network to the peer network
        /// </summary>
        public InputList<string> Routes
        {
            get => _routes ?? (_routes = new InputList<string>());
            set => _routes = value;
        }

        public PeeringState()
        {
        }
    }
}
