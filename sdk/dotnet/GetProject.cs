// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Eventstorecloud
{
    public static class GetProject
    {
        /// <summary>
        /// Retrieves data for an existing `Project` resource
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Eventstorecloud = Pulumi.Eventstorecloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Eventstorecloud.GetProject.InvokeAsync(new Eventstorecloud.GetProjectArgs
        ///         {
        ///             Name = "Example Project",
        ///         }));
        ///         this.ProjectId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("projectId")]
        ///     public Output&lt;string&gt; ProjectId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("eventstorecloud:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithVersion());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetProjectArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetProjectResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
