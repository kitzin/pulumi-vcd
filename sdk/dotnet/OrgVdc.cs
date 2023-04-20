// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/orgVdc:OrgVdc")]
    public partial class OrgVdc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The allocation model used by this VDC; must be one of {AllocationVApp, AllocationPool, ReservationPool, Flex}
        /// </summary>
        [Output("allocationModel")]
        public Output<string> AllocationModel { get; private set; } = null!;

        /// <summary>
        /// Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the
        /// ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
        /// </summary>
        [Output("allowOverCommit")]
        public Output<bool> AllowOverCommit { get; private set; } = null!;

        /// <summary>
        /// The compute capacity allocated to this VDC.
        /// </summary>
        [Output("computeCapacity")]
        public Output<Outputs.OrgVdcComputeCapacity> ComputeCapacity { get; private set; } = null!;

        /// <summary>
        /// Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then
        /// 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If the
        /// element is empty, vCD sets a value
        /// </summary>
        [Output("cpuGuaranteed")]
        public Output<double> CpuGuaranteed { get; private set; } = null!;

        /// <summary>
        /// Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will
        /// consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or
        /// AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if the element is empty or missing.
        /// </summary>
        [Output("cpuSpeed")]
        public Output<int> CpuSpeed { get; private set; } = null!;

        /// <summary>
        /// ID of default Compute policy for this VDC, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
        /// </summary>
        [Output("defaultComputePolicyId")]
        public Output<string> DefaultComputePolicyId { get; private set; } = null!;

        /// <summary>
        /// ID of default VM Compute policy, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
        /// </summary>
        [Output("defaultVmSizingPolicyId")]
        public Output<string> DefaultVmSizingPolicyId { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_force=True to remove a VDC and any objects it contains, regardless of their state.
        /// </summary>
        [Output("deleteForce")]
        public Output<bool> DeleteForce { get; private set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the VDC and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Output("deleteRecursive")]
        public Output<bool> DeleteRecursive { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of NSX-T Edge Cluster (provider vApp networking services and DHCP capability for Isolated networks)
        /// </summary>
        [Output("edgeClusterId")]
        public Output<string?> EdgeClusterId { get; private set; } = null!;

        /// <summary>
        /// Set to true to indicate if the Flex VDC is to be elastic.
        /// </summary>
        [Output("elasticity")]
        public Output<bool> Elasticity { get; private set; } = null!;

        /// <summary>
        /// Request for fast provisioning. Request will be honored only if the underlying datas tore supports it. Fast provisioning
        /// can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast
        /// provisioning, all provisioning operations will result in full clones.
        /// </summary>
        [Output("enableFastProvisioning")]
        public Output<bool?> EnableFastProvisioning { get; private set; } = null!;

        /// <summary>
        /// Boolean to request thin provisioning. Request will be honored only if the underlying datastore supports it. Thin
        /// provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
        /// </summary>
        [Output("enableThinProvisioning")]
        public Output<bool?> EnableThinProvisioning { get; private set; } = null!;

        /// <summary>
        /// True if discovery of vCenter VMs is enabled for resource pools backing this VDC. If left unspecified, the actual
        /// behaviour depends on enablement at the organization level and at the system level.
        /// </summary>
        [Output("enableVmDiscovery")]
        public Output<bool?> EnableVmDiscovery { get; private set; } = null!;

        /// <summary>
        /// True if this VDC is enabled for use by the organization VDCs. Default is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Set to true to indicate if the Flex VDC is to include memory overhead into its accounting for admission control.
        /// </summary>
        [Output("includeVmMemoryOverhead")]
        public Output<bool> IncludeVmMemoryOverhead { get; private set; } = null!;

        /// <summary>
        /// Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75,
        /// then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. When
        /// Allocation model is AllocationPool minimum value is 0.2. If the element is empty, vCD sets a value.
        /// </summary>
        [Output("memoryGuaranteed")]
        public Output<double> MemoryGuaranteed { get; private set; } = null!;

        /// <summary>
        /// Key and value pairs for Org VDC metadata
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Metadata entries for the given VDC
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.OrgVdcMetadataEntry>> MetadataEntries { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
        /// </summary>
        [Output("networkPoolName")]
        public Output<string?> NetworkPoolName { get; private set; } = null!;

        /// <summary>
        /// Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be
        /// deployed.
        /// </summary>
        [Output("networkQuota")]
        public Output<int?> NetworkQuota { get; private set; } = null!;

        /// <summary>
        /// Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
        /// </summary>
        [Output("nicQuota")]
        public Output<int?> NicQuota { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// A reference to the Provider VDC from which this organization VDC is provisioned.
        /// </summary>
        [Output("providerVdcName")]
        public Output<string> ProviderVdcName { get; private set; } = null!;

        /// <summary>
        /// Storage profiles supported by this VDC.
        /// </summary>
        [Output("storageProfiles")]
        public Output<ImmutableArray<Outputs.OrgVdcStorageProfile>> StorageProfiles { get; private set; } = null!;

        /// <summary>
        /// Set of VM Placement Policy IDs
        /// </summary>
        [Output("vmPlacementPolicyIds")]
        public Output<ImmutableArray<string>> VmPlacementPolicyIds { get; private set; } = null!;

        /// <summary>
        /// The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp
        /// templates. Defaults to 0, which specifies an unlimited number.
        /// </summary>
        [Output("vmQuota")]
        public Output<int?> VmQuota { get; private set; } = null!;

        /// <summary>
        /// Set of VM Sizing Policy IDs
        /// </summary>
        [Output("vmSizingPolicyIds")]
        public Output<ImmutableArray<string>> VmSizingPolicyIds { get; private set; } = null!;


        /// <summary>
        /// Create a OrgVdc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgVdc(string name, OrgVdcArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/orgVdc:OrgVdc", name, args ?? new OrgVdcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgVdc(string name, Input<string> id, OrgVdcState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/orgVdc:OrgVdc", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrgVdc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgVdc Get(string name, Input<string> id, OrgVdcState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgVdc(name, id, state, options);
        }
    }

    public sealed class OrgVdcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation model used by this VDC; must be one of {AllocationVApp, AllocationPool, ReservationPool, Flex}
        /// </summary>
        [Input("allocationModel", required: true)]
        public Input<string> AllocationModel { get; set; } = null!;

        /// <summary>
        /// Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the
        /// ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
        /// </summary>
        [Input("allowOverCommit")]
        public Input<bool>? AllowOverCommit { get; set; }

        /// <summary>
        /// The compute capacity allocated to this VDC.
        /// </summary>
        [Input("computeCapacity", required: true)]
        public Input<Inputs.OrgVdcComputeCapacityArgs> ComputeCapacity { get; set; } = null!;

        /// <summary>
        /// Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then
        /// 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If the
        /// element is empty, vCD sets a value
        /// </summary>
        [Input("cpuGuaranteed")]
        public Input<double>? CpuGuaranteed { get; set; }

        /// <summary>
        /// Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will
        /// consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or
        /// AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if the element is empty or missing.
        /// </summary>
        [Input("cpuSpeed")]
        public Input<int>? CpuSpeed { get; set; }

        /// <summary>
        /// ID of default Compute policy for this VDC, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
        /// </summary>
        [Input("defaultComputePolicyId")]
        public Input<string>? DefaultComputePolicyId { get; set; }

        /// <summary>
        /// ID of default VM Compute policy, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
        /// </summary>
        [Input("defaultVmSizingPolicyId")]
        public Input<string>? DefaultVmSizingPolicyId { get; set; }

        /// <summary>
        /// When destroying use delete_force=True to remove a VDC and any objects it contains, regardless of their state.
        /// </summary>
        [Input("deleteForce", required: true)]
        public Input<bool> DeleteForce { get; set; } = null!;

        /// <summary>
        /// When destroying use delete_recursive=True to remove the VDC and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive", required: true)]
        public Input<bool> DeleteRecursive { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of NSX-T Edge Cluster (provider vApp networking services and DHCP capability for Isolated networks)
        /// </summary>
        [Input("edgeClusterId")]
        public Input<string>? EdgeClusterId { get; set; }

        /// <summary>
        /// Set to true to indicate if the Flex VDC is to be elastic.
        /// </summary>
        [Input("elasticity")]
        public Input<bool>? Elasticity { get; set; }

        /// <summary>
        /// Request for fast provisioning. Request will be honored only if the underlying datas tore supports it. Fast provisioning
        /// can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast
        /// provisioning, all provisioning operations will result in full clones.
        /// </summary>
        [Input("enableFastProvisioning")]
        public Input<bool>? EnableFastProvisioning { get; set; }

        /// <summary>
        /// Boolean to request thin provisioning. Request will be honored only if the underlying datastore supports it. Thin
        /// provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
        /// </summary>
        [Input("enableThinProvisioning")]
        public Input<bool>? EnableThinProvisioning { get; set; }

        /// <summary>
        /// True if discovery of vCenter VMs is enabled for resource pools backing this VDC. If left unspecified, the actual
        /// behaviour depends on enablement at the organization level and at the system level.
        /// </summary>
        [Input("enableVmDiscovery")]
        public Input<bool>? EnableVmDiscovery { get; set; }

        /// <summary>
        /// True if this VDC is enabled for use by the organization VDCs. Default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Set to true to indicate if the Flex VDC is to include memory overhead into its accounting for admission control.
        /// </summary>
        [Input("includeVmMemoryOverhead")]
        public Input<bool>? IncludeVmMemoryOverhead { get; set; }

        /// <summary>
        /// Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75,
        /// then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. When
        /// Allocation model is AllocationPool minimum value is 0.2. If the element is empty, vCD sets a value.
        /// </summary>
        [Input("memoryGuaranteed")]
        public Input<double>? MemoryGuaranteed { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key and value pairs for Org VDC metadata
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.OrgVdcMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given VDC
        /// </summary>
        public InputList<Inputs.OrgVdcMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.OrgVdcMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
        /// </summary>
        [Input("networkPoolName")]
        public Input<string>? NetworkPoolName { get; set; }

        /// <summary>
        /// Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be
        /// deployed.
        /// </summary>
        [Input("networkQuota")]
        public Input<int>? NetworkQuota { get; set; }

        /// <summary>
        /// Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
        /// </summary>
        [Input("nicQuota")]
        public Input<int>? NicQuota { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A reference to the Provider VDC from which this organization VDC is provisioned.
        /// </summary>
        [Input("providerVdcName", required: true)]
        public Input<string> ProviderVdcName { get; set; } = null!;

        [Input("storageProfiles", required: true)]
        private InputList<Inputs.OrgVdcStorageProfileArgs>? _storageProfiles;

        /// <summary>
        /// Storage profiles supported by this VDC.
        /// </summary>
        public InputList<Inputs.OrgVdcStorageProfileArgs> StorageProfiles
        {
            get => _storageProfiles ?? (_storageProfiles = new InputList<Inputs.OrgVdcStorageProfileArgs>());
            set => _storageProfiles = value;
        }

        [Input("vmPlacementPolicyIds")]
        private InputList<string>? _vmPlacementPolicyIds;

        /// <summary>
        /// Set of VM Placement Policy IDs
        /// </summary>
        public InputList<string> VmPlacementPolicyIds
        {
            get => _vmPlacementPolicyIds ?? (_vmPlacementPolicyIds = new InputList<string>());
            set => _vmPlacementPolicyIds = value;
        }

        /// <summary>
        /// The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp
        /// templates. Defaults to 0, which specifies an unlimited number.
        /// </summary>
        [Input("vmQuota")]
        public Input<int>? VmQuota { get; set; }

        [Input("vmSizingPolicyIds")]
        private InputList<string>? _vmSizingPolicyIds;

        /// <summary>
        /// Set of VM Sizing Policy IDs
        /// </summary>
        public InputList<string> VmSizingPolicyIds
        {
            get => _vmSizingPolicyIds ?? (_vmSizingPolicyIds = new InputList<string>());
            set => _vmSizingPolicyIds = value;
        }

        public OrgVdcArgs()
        {
        }
        public static new OrgVdcArgs Empty => new OrgVdcArgs();
    }

    public sealed class OrgVdcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation model used by this VDC; must be one of {AllocationVApp, AllocationPool, ReservationPool, Flex}
        /// </summary>
        [Input("allocationModel")]
        public Input<string>? AllocationModel { get; set; }

        /// <summary>
        /// Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the
        /// ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
        /// </summary>
        [Input("allowOverCommit")]
        public Input<bool>? AllowOverCommit { get; set; }

        /// <summary>
        /// The compute capacity allocated to this VDC.
        /// </summary>
        [Input("computeCapacity")]
        public Input<Inputs.OrgVdcComputeCapacityGetArgs>? ComputeCapacity { get; set; }

        /// <summary>
        /// Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then
        /// 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If the
        /// element is empty, vCD sets a value
        /// </summary>
        [Input("cpuGuaranteed")]
        public Input<double>? CpuGuaranteed { get; set; }

        /// <summary>
        /// Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will
        /// consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or
        /// AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if the element is empty or missing.
        /// </summary>
        [Input("cpuSpeed")]
        public Input<int>? CpuSpeed { get; set; }

        /// <summary>
        /// ID of default Compute policy for this VDC, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
        /// </summary>
        [Input("defaultComputePolicyId")]
        public Input<string>? DefaultComputePolicyId { get; set; }

        /// <summary>
        /// ID of default VM Compute policy, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
        /// </summary>
        [Input("defaultVmSizingPolicyId")]
        public Input<string>? DefaultVmSizingPolicyId { get; set; }

        /// <summary>
        /// When destroying use delete_force=True to remove a VDC and any objects it contains, regardless of their state.
        /// </summary>
        [Input("deleteForce")]
        public Input<bool>? DeleteForce { get; set; }

        /// <summary>
        /// When destroying use delete_recursive=True to remove the VDC and any objects it contains that are in a state that
        /// normally allows removal.
        /// </summary>
        [Input("deleteRecursive")]
        public Input<bool>? DeleteRecursive { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of NSX-T Edge Cluster (provider vApp networking services and DHCP capability for Isolated networks)
        /// </summary>
        [Input("edgeClusterId")]
        public Input<string>? EdgeClusterId { get; set; }

        /// <summary>
        /// Set to true to indicate if the Flex VDC is to be elastic.
        /// </summary>
        [Input("elasticity")]
        public Input<bool>? Elasticity { get; set; }

        /// <summary>
        /// Request for fast provisioning. Request will be honored only if the underlying datas tore supports it. Fast provisioning
        /// can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast
        /// provisioning, all provisioning operations will result in full clones.
        /// </summary>
        [Input("enableFastProvisioning")]
        public Input<bool>? EnableFastProvisioning { get; set; }

        /// <summary>
        /// Boolean to request thin provisioning. Request will be honored only if the underlying datastore supports it. Thin
        /// provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
        /// </summary>
        [Input("enableThinProvisioning")]
        public Input<bool>? EnableThinProvisioning { get; set; }

        /// <summary>
        /// True if discovery of vCenter VMs is enabled for resource pools backing this VDC. If left unspecified, the actual
        /// behaviour depends on enablement at the organization level and at the system level.
        /// </summary>
        [Input("enableVmDiscovery")]
        public Input<bool>? EnableVmDiscovery { get; set; }

        /// <summary>
        /// True if this VDC is enabled for use by the organization VDCs. Default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Set to true to indicate if the Flex VDC is to include memory overhead into its accounting for admission control.
        /// </summary>
        [Input("includeVmMemoryOverhead")]
        public Input<bool>? IncludeVmMemoryOverhead { get; set; }

        /// <summary>
        /// Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75,
        /// then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. When
        /// Allocation model is AllocationPool minimum value is 0.2. If the element is empty, vCD sets a value.
        /// </summary>
        [Input("memoryGuaranteed")]
        public Input<double>? MemoryGuaranteed { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Key and value pairs for Org VDC metadata
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.OrgVdcMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// Metadata entries for the given VDC
        /// </summary>
        public InputList<Inputs.OrgVdcMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.OrgVdcMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
        /// </summary>
        [Input("networkPoolName")]
        public Input<string>? NetworkPoolName { get; set; }

        /// <summary>
        /// Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be
        /// deployed.
        /// </summary>
        [Input("networkQuota")]
        public Input<int>? NetworkQuota { get; set; }

        /// <summary>
        /// Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
        /// </summary>
        [Input("nicQuota")]
        public Input<int>? NicQuota { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A reference to the Provider VDC from which this organization VDC is provisioned.
        /// </summary>
        [Input("providerVdcName")]
        public Input<string>? ProviderVdcName { get; set; }

        [Input("storageProfiles")]
        private InputList<Inputs.OrgVdcStorageProfileGetArgs>? _storageProfiles;

        /// <summary>
        /// Storage profiles supported by this VDC.
        /// </summary>
        public InputList<Inputs.OrgVdcStorageProfileGetArgs> StorageProfiles
        {
            get => _storageProfiles ?? (_storageProfiles = new InputList<Inputs.OrgVdcStorageProfileGetArgs>());
            set => _storageProfiles = value;
        }

        [Input("vmPlacementPolicyIds")]
        private InputList<string>? _vmPlacementPolicyIds;

        /// <summary>
        /// Set of VM Placement Policy IDs
        /// </summary>
        public InputList<string> VmPlacementPolicyIds
        {
            get => _vmPlacementPolicyIds ?? (_vmPlacementPolicyIds = new InputList<string>());
            set => _vmPlacementPolicyIds = value;
        }

        /// <summary>
        /// The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp
        /// templates. Defaults to 0, which specifies an unlimited number.
        /// </summary>
        [Input("vmQuota")]
        public Input<int>? VmQuota { get; set; }

        [Input("vmSizingPolicyIds")]
        private InputList<string>? _vmSizingPolicyIds;

        /// <summary>
        /// Set of VM Sizing Policy IDs
        /// </summary>
        public InputList<string> VmSizingPolicyIds
        {
            get => _vmSizingPolicyIds ?? (_vmSizingPolicyIds = new InputList<string>());
            set => _vmSizingPolicyIds = value;
        }

        public OrgVdcState()
        {
        }
        public static new OrgVdcState Empty => new OrgVdcState();
    }
}
