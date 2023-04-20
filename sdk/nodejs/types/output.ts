// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface CatalogAccessControlSharedWith {
    accessLevel: string;
    groupId?: string;
    orgId?: string;
    subjectName: string;
    userId?: string;
}

export interface CatalogItemMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface CatalogMediaMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface CatalogMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface CatalogVappTemplateMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface EdgeGatewayExternalNetwork {
    enableRateLimit?: boolean;
    incomingRateLimit?: number;
    name: string;
    outgoingRateLimit?: number;
    subnets: outputs.EdgeGatewayExternalNetworkSubnet[];
}

export interface EdgeGatewayExternalNetworkSubnet {
    gateway: string;
    ipAddress?: string;
    netmask: string;
    suballocatePools?: outputs.EdgeGatewayExternalNetworkSubnetSuballocatePool[];
    useForDefaultRoute?: boolean;
}

export interface EdgeGatewayExternalNetworkSubnetSuballocatePool {
    endAddress: string;
    startAddress: string;
}

export interface EdgeGatewayVpnLocalSubnet {
    localSubnetGateway: string;
    localSubnetMask: string;
    localSubnetName: string;
}

export interface EdgeGatewayVpnPeerSubnet {
    peerSubnetGateway: string;
    peerSubnetMask: string;
    peerSubnetName: string;
}

export interface ExternalNetworkIpScope {
    dns1?: string;
    dns2?: string;
    dnsSuffix?: string;
    gateway: string;
    netmask: string;
    staticIpPools?: outputs.ExternalNetworkIpScopeStaticIpPool[];
}

export interface ExternalNetworkIpScopeStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface ExternalNetworkV2IpScope {
    dns1?: string;
    dns2?: string;
    dnsSuffix?: string;
    enabled?: boolean;
    gateway: string;
    prefixLength: number;
    staticIpPools?: outputs.ExternalNetworkV2IpScopeStaticIpPool[];
}

export interface ExternalNetworkV2IpScopeStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface ExternalNetworkV2NsxtNetwork {
    nsxtManagerId: string;
    nsxtSegmentName?: string;
    nsxtTier0RouterId?: string;
}

export interface ExternalNetworkV2VsphereNetwork {
    portgroupId: string;
    vcenterId: string;
}

export interface ExternalNetworkVsphereNetwork {
    name: string;
    type: string;
    vcenter: string;
}

export interface GetCatalogFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: outputs.GetCatalogFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogItemFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: outputs.GetCatalogItemFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogItemFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogItemMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetCatalogMediaFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: outputs.GetCatalogMediaFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogMediaFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogMediaMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetCatalogMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetCatalogVappTemplateFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: outputs.GetCatalogVappTemplateFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogVappTemplateFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogVappTemplateMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetEdgegatewayExternalNetwork {
    enableRateLimit: boolean;
    incomingRateLimit: number;
    name: string;
    outgoingRateLimit: number;
    subnets: outputs.GetEdgegatewayExternalNetworkSubnet[];
}

export interface GetEdgegatewayExternalNetworkSubnet {
    gateway: string;
    ipAddress: string;
    netmask: string;
    suballocatePools: outputs.GetEdgegatewayExternalNetworkSubnetSuballocatePool[];
    useForDefaultRoute: boolean;
}

export interface GetEdgegatewayExternalNetworkSubnetSuballocatePool {
    endAddress: string;
    startAddress: string;
}

export interface GetEdgegatewayFilter {
    nameRegex?: string;
}

export interface GetExternalNetworkIpScope {
    dns1: string;
    dns2: string;
    dnsSuffix: string;
    gateway: string;
    netmask: string;
    staticIpPools: outputs.GetExternalNetworkIpScopeStaticIpPool[];
}

export interface GetExternalNetworkIpScopeStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetExternalNetworkV2IpScope {
    dns1?: string;
    dns2?: string;
    dnsSuffix?: string;
    enabled?: boolean;
    gateway: string;
    prefixLength: number;
    staticIpPools?: outputs.GetExternalNetworkV2IpScopeStaticIpPool[];
}

export interface GetExternalNetworkV2IpScopeStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetExternalNetworkV2NsxtNetwork {
    nsxtManagerId: string;
    nsxtSegmentName: string;
    nsxtTier0RouterId: string;
}

export interface GetExternalNetworkV2VsphereNetwork {
    portgroupId: string;
    vcenterId: string;
}

export interface GetExternalNetworkVsphereNetwork {
    name: string;
    type: string;
    vcenter: string;
}

export interface GetIndependentDiskMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetLbServerPoolMember {
    condition: string;
    id: string;
    ipAddress: string;
    maxConnections: number;
    minConnections: number;
    monitorPort: number;
    name: string;
    port: number;
    weight: number;
}

export interface GetNetworkDirectFilter {
    ip?: string;
    metadatas?: outputs.GetNetworkDirectFilterMetadata[];
    nameRegex?: string;
}

export interface GetNetworkDirectFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetNetworkDirectMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetNetworkIsolatedDhcpPool {
    defaultLeaseTime: number;
    endAddress: string;
    maxLeaseTime: number;
    startAddress: string;
}

export interface GetNetworkIsolatedFilter {
    ip?: string;
    metadatas?: outputs.GetNetworkIsolatedFilterMetadata[];
    nameRegex?: string;
}

export interface GetNetworkIsolatedFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetNetworkIsolatedMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetNetworkIsolatedStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetNetworkIsolatedV2Filter {
    ip?: string;
    nameRegex?: string;
}

export interface GetNetworkIsolatedV2MetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetNetworkIsolatedV2StaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetNetworkRoutedDhcpPool {
    defaultLeaseTime: number;
    endAddress: string;
    maxLeaseTime: number;
    startAddress: string;
}

export interface GetNetworkRoutedFilter {
    ip?: string;
    metadatas?: outputs.GetNetworkRoutedFilterMetadata[];
    nameRegex?: string;
}

export interface GetNetworkRoutedFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetNetworkRoutedMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetNetworkRoutedStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetNetworkRoutedV2Filter {
    ip?: string;
    nameRegex?: string;
}

export interface GetNetworkRoutedV2MetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetNetworkRoutedV2StaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetNsxtAlbPoolHealthMonitor {
    name: string;
    systemDefined: boolean;
    type: string;
}

export interface GetNsxtAlbPoolMember {
    detailedHealthMessage: string;
    enabled: boolean;
    healthStatus: string;
    ipAddress: string;
    markedDownBies: string[];
    port: number;
    ratio: number;
}

export interface GetNsxtAlbPoolPersistenceProfile {
    name: string;
    type: string;
    value: string;
}

export interface GetNsxtAlbVirtualServiceServicePort {
    endPort: number;
    sslEnabled: boolean;
    startPort: number;
    type: string;
}

export interface GetNsxtAppPortProfileAppPort {
    ports: string[];
    protocol: string;
}

export interface GetNsxtDistributedFirewallRule {
    action: string;
    appPortProfileIds: string[];
    comment: string;
    description: string;
    destinationGroupsExcluded: boolean;
    destinationIds: string[];
    direction: string;
    enabled: boolean;
    id: string;
    ipProtocol: string;
    logging: boolean;
    name: string;
    networkContextProfileIds: string[];
    sourceGroupsExcluded: boolean;
    sourceIds: string[];
}

export interface GetNsxtDynamicSecurityGroupCriteria {
    rules: outputs.GetNsxtDynamicSecurityGroupCriteriaRule[];
}

export interface GetNsxtDynamicSecurityGroupCriteriaRule {
    operator: string;
    type: string;
    value: string;
}

export interface GetNsxtDynamicSecurityGroupMemberVm {
    vappId: string;
    vappName: string;
    vmId: string;
    vmName: string;
}

export interface GetNsxtEdgegatewayBgpIpPrefixListIpPrefix {
    action: string;
    greaterThanOrEqualTo: number;
    lessThanOrEqualTo: number;
    network: string;
}

export interface GetNsxtEdgegatewaySubnet {
    allocatedIps: outputs.GetNsxtEdgegatewaySubnetAllocatedIp[];
    gateway: string;
    prefixLength: number;
    primaryIp: string;
}

export interface GetNsxtEdgegatewaySubnetAllocatedIp {
    endAddress: string;
    startAddress: string;
}

export interface GetNsxtFirewallRule {
    action: string;
    appPortProfileIds: string[];
    destinationIds: string[];
    direction: string;
    enabled: boolean;
    id: string;
    ipProtocol: string;
    logging: boolean;
    name: string;
    sourceIds: string[];
}

export interface GetNsxtIpsecVpnTunnelSecurityProfileCustomization {
    dpdProbeInternal: number;
    ikeDhGroups: string[];
    ikeDigestAlgorithms: string[];
    ikeEncryptionAlgorithms: string[];
    ikeSaLifetime: number;
    ikeVersion: string;
    tunnelDfPolicy: string;
    tunnelDhGroups: string[];
    tunnelDigestAlgorithms: string[];
    tunnelEncryptionAlgorithms: string[];
    tunnelPfsEnabled: boolean;
    tunnelSaLifetime: number;
}

export interface GetNsxtNetworkDhcpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetNsxtNetworkImportedFilter {
    ip?: string;
    nameRegex?: string;
}

export interface GetNsxtNetworkImportedStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetNsxtSecurityGroupMemberVm {
    vappId: string;
    vappName: string;
    vmId: string;
    vmName: string;
}

export interface GetNsxvDhcpRelayRelayAgent {
    gatewayIpAddress: string;
    networkName: string;
}

export interface GetNsxvFirewallRuleDestination {
    exclude: boolean;
    gatewayInterfaces: string[];
    ipAddresses: string[];
    ipSets: string[];
    orgNetworks: string[];
    vmIds: string[];
}

export interface GetNsxvFirewallRuleService {
    port: string;
    protocol: string;
    sourcePort: string;
}

export interface GetNsxvFirewallRuleSource {
    exclude: boolean;
    gatewayInterfaces: string[];
    ipAddresses: string[];
    ipSets: string[];
    orgNetworks: string[];
    vmIds: string[];
}

export interface GetOrgLdapCustomSetting {
    authenticationMethod: string;
    baseDistinguishedName: string;
    connectorType: string;
    groupAttributes: outputs.GetOrgLdapCustomSettingGroupAttribute[];
    isSsl: boolean;
    password: string;
    port: number;
    server: string;
    userAttributes: outputs.GetOrgLdapCustomSettingUserAttribute[];
    username: string;
}

export interface GetOrgLdapCustomSettingGroupAttribute {
    groupBackLinkIdentifier: string;
    groupMembershipIdentifier: string;
    membership: string;
    name: string;
    objectClass: string;
    uniqueIdentifier: string;
}

export interface GetOrgLdapCustomSettingUserAttribute {
    displayName: string;
    email: string;
    givenName: string;
    groupBackLinkIdentifier: string;
    groupMembershipIdentifier: string;
    objectClass: string;
    surname: string;
    telephone: string;
    uniqueIdentifier: string;
    username: string;
}

export interface GetOrgMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetOrgVappLease {
    deleteOnStorageLeaseExpiration: boolean;
    maximumRuntimeLeaseInSec: number;
    maximumStorageLeaseInSec: number;
    powerOffOnRuntimeLeaseExpiration: boolean;
}

export interface GetOrgVappTemplateLease {
    deleteOnStorageLeaseExpiration: boolean;
    maximumStorageLeaseInSec: number;
}

export interface GetOrgVdcComputeCapacity {
    cpus: outputs.GetOrgVdcComputeCapacityCpus[];
    memories: outputs.GetOrgVdcComputeCapacityMemory[];
}

export interface GetOrgVdcComputeCapacityCpus {
    allocated: number;
    limit: number;
    reserved: number;
    used: number;
}

export interface GetOrgVdcComputeCapacityMemory {
    allocated: number;
    limit: number;
    reserved: number;
    used: number;
}

export interface GetOrgVdcMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface GetOrgVdcStorageProfile {
    default: boolean;
    enabled: boolean;
    limit: number;
    name: string;
    storageUsedInMb: number;
}

export interface GetProviderVdcComputeCapacity {
    cpus: outputs.GetProviderVdcComputeCapacityCpus[];
    isElastic: boolean;
    isHa: boolean;
    memories: outputs.GetProviderVdcComputeCapacityMemory[];
}

export interface GetProviderVdcComputeCapacityCpus {
    allocation: number;
    overhead: number;
    reserved: number;
    total: number;
    units: string;
    used: number;
}

export interface GetProviderVdcComputeCapacityMemory {
    allocation: number;
    overhead: number;
    reserved: number;
    total: number;
    units: string;
    used: number;
}

export interface GetProviderVdcMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetResourceSchemaAttribute {
    computed: boolean;
    description: string;
    name: string;
    optional: boolean;
    required: boolean;
    sensitive: boolean;
    type: string;
}

export interface GetResourceSchemaBlockAttribute {
    attributes: outputs.GetResourceSchemaBlockAttributeAttribute[];
    name: string;
    nestingMode: string;
}

export interface GetResourceSchemaBlockAttributeAttribute {
    computed: boolean;
    description: string;
    name: string;
    optional: boolean;
    required: boolean;
    sensitive: boolean;
    type: string;
}

export interface GetRightImpliedRight {
    id: string;
    name: string;
}

export interface GetStorageProfileIopsSetting {
    defaultDiskIops: number;
    diskIopsPerGbMax: number;
    iopsLimit: number;
    iopsLimitingEnabled: boolean;
    maximumDiskIops: number;
}

export interface GetStorageProfileMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetSubscribedCatalogFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: outputs.GetSubscribedCatalogFilterMetadata[];
    nameRegex?: string;
}

export interface GetSubscribedCatalogFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetVappLease {
    runtimeLeaseInSec: number;
    storageLeaseInSec: number;
}

export interface GetVappMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetVappNetworkDhcpPool {
    defaultLeaseTime: number;
    enabled: boolean;
    endAddress: string;
    maxLeaseTime: number;
    startAddress: string;
}

export interface GetVappNetworkStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface GetVappVmCustomization {
    adminPassword: string;
    allowLocalAdminPassword: boolean;
    autoGeneratePassword: boolean;
    changeSid: boolean;
    enabled: boolean;
    force: boolean;
    initscript: string;
    joinDomain: boolean;
    joinDomainAccountOu: string;
    joinDomainName: string;
    joinDomainPassword: string;
    joinDomainUser: string;
    joinOrgDomain: boolean;
    mustChangePasswordOnFirstLogin: boolean;
    numberOfAutoLogons: number;
}

export interface GetVappVmDisk {
    busNumber: string;
    name: string;
    sizeInMb: number;
    unitNumber: string;
}

export interface GetVappVmInternalDisk {
    busNumber: number;
    busType: string;
    diskId: string;
    iops: number;
    sizeInMb: number;
    storageProfile: string;
    thinProvisioned: boolean;
    unitNumber: number;
}

export interface GetVappVmMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetVappVmNetwork {
    adapterType: string;
    connected: boolean;
    ip: string;
    ipAllocationMode: string;
    isPrimary: boolean;
    mac: string;
    name: string;
    type: string;
}

export interface GetVdcGroupParticipatingOrgVdc {
    faultDomainTag: string;
    isRemoteOrg: boolean;
    networkProviderScope: string;
    orgId: string;
    orgName: string;
    siteId: string;
    siteName: string;
    status: string;
    vdcId: string;
    vdcName: string;
}

export interface GetVmCustomization {
    adminPassword: string;
    allowLocalAdminPassword: boolean;
    autoGeneratePassword: boolean;
    changeSid: boolean;
    enabled: boolean;
    force: boolean;
    initscript: string;
    joinDomain: boolean;
    joinDomainAccountOu: string;
    joinDomainName: string;
    joinDomainPassword: string;
    joinDomainUser: string;
    joinOrgDomain: boolean;
    mustChangePasswordOnFirstLogin: boolean;
    numberOfAutoLogons: number;
}

export interface GetVmDisk {
    busNumber: string;
    name: string;
    sizeInMb: number;
    unitNumber: string;
}

export interface GetVmInternalDisk {
    busNumber: number;
    busType: string;
    diskId: string;
    iops: number;
    sizeInMb: number;
    storageProfile: string;
    thinProvisioned: boolean;
    unitNumber: number;
}

export interface GetVmMetadataEntry {
    isSystem: boolean;
    key: string;
    type: string;
    userAccess: string;
    value: string;
}

export interface GetVmNetwork {
    adapterType: string;
    connected: boolean;
    ip: string;
    ipAllocationMode: string;
    isPrimary: boolean;
    mac: string;
    name: string;
    type: string;
}

export interface GetVmSizingPolicyCpus {
    coresPerSocket: string;
    count: string;
    limitInMhz: string;
    reservationGuarantee: string;
    shares: string;
    speedInMhz: string;
}

export interface GetVmSizingPolicyMemory {
    limitInMb: string;
    reservationGuarantee: string;
    shares: string;
    sizeInMb: string;
}

export interface IndependentDiskMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface LbServerPoolMember {
    condition: string;
    id: string;
    ipAddress: string;
    maxConnections?: number;
    minConnections?: number;
    monitorPort: number;
    name: string;
    port: number;
    weight: number;
}

export interface NetworkDirectMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface NetworkIsolatedDhcpPool {
    defaultLeaseTime?: number;
    endAddress: string;
    maxLeaseTime?: number;
    startAddress: string;
}

export interface NetworkIsolatedMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface NetworkIsolatedStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface NetworkIsolatedV2MetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface NetworkIsolatedV2StaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface NetworkRoutedDhcpPool {
    defaultLeaseTime: number;
    endAddress: string;
    maxLeaseTime?: number;
    startAddress: string;
}

export interface NetworkRoutedMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface NetworkRoutedStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface NetworkRoutedV2MetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface NetworkRoutedV2StaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface NsxtAlbPoolHealthMonitor {
    name: string;
    systemDefined: boolean;
    type: string;
}

export interface NsxtAlbPoolMember {
    detailedHealthMessage: string;
    enabled?: boolean;
    healthStatus: string;
    ipAddress: string;
    markedDownBies: string[];
    port?: number;
    ratio?: number;
}

export interface NsxtAlbPoolPersistenceProfile {
    name: string;
    type: string;
    value?: string;
}

export interface NsxtAlbVirtualServiceServicePort {
    endPort?: number;
    sslEnabled?: boolean;
    startPort: number;
    type: string;
}

export interface NsxtAppPortProfileAppPort {
    ports?: string[];
    protocol: string;
}

export interface NsxtDistributedFirewallRule {
    action: string;
    appPortProfileIds?: string[];
    comment?: string;
    description?: string;
    destinationGroupsExcluded?: boolean;
    destinationIds?: string[];
    direction?: string;
    enabled?: boolean;
    id: string;
    ipProtocol?: string;
    logging?: boolean;
    name: string;
    networkContextProfileIds?: string[];
    sourceGroupsExcluded?: boolean;
    sourceIds?: string[];
}

export interface NsxtDynamicSecurityGroupCriteria {
    rules?: outputs.NsxtDynamicSecurityGroupCriteriaRule[];
}

export interface NsxtDynamicSecurityGroupCriteriaRule {
    operator: string;
    type: string;
    value: string;
}

export interface NsxtDynamicSecurityGroupMemberVm {
    vappId: string;
    vappName: string;
    vmId: string;
    vmName: string;
}

export interface NsxtEdgegatewayBgpIpPrefixListIpPrefix {
    action: string;
    greaterThanOrEqualTo?: number;
    lessThanOrEqualTo?: number;
    network: string;
}

export interface NsxtEdgegatewaySubnet {
    allocatedIps: outputs.NsxtEdgegatewaySubnetAllocatedIp[];
    gateway: string;
    prefixLength: number;
    primaryIp?: string;
}

export interface NsxtEdgegatewaySubnetAllocatedIp {
    endAddress: string;
    startAddress: string;
}

export interface NsxtFirewallRule {
    action: string;
    appPortProfileIds?: string[];
    destinationIds?: string[];
    direction: string;
    enabled?: boolean;
    id: string;
    ipProtocol: string;
    logging?: boolean;
    name: string;
    sourceIds?: string[];
}

export interface NsxtIpsecVpnTunnelSecurityProfileCustomization {
    dpdProbeInternal?: number;
    ikeDhGroups: string[];
    ikeDigestAlgorithms?: string[];
    ikeEncryptionAlgorithms: string[];
    ikeSaLifetime?: number;
    ikeVersion: string;
    tunnelDfPolicy?: string;
    tunnelDhGroups: string[];
    tunnelDigestAlgorithms?: string[];
    tunnelEncryptionAlgorithms: string[];
    tunnelPfsEnabled?: boolean;
    tunnelSaLifetime?: number;
}

export interface NsxtNetworkDhcpPool {
    endAddress: string;
    startAddress: string;
}

export interface NsxtNetworkImportedStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface NsxtSecurityGroupMemberVm {
    vappId: string;
    vappName: string;
    vmId: string;
    vmName: string;
}

export interface NsxvDhcpRelayRelayAgent {
    gatewayIpAddress: string;
    networkName: string;
}

export interface NsxvFirewallRuleDestination {
    exclude?: boolean;
    gatewayInterfaces?: string[];
    ipAddresses?: string[];
    ipSets?: string[];
    orgNetworks?: string[];
    vmIds?: string[];
}

export interface NsxvFirewallRuleService {
    port: string;
    protocol: string;
    sourcePort: string;
}

export interface NsxvFirewallRuleSource {
    exclude?: boolean;
    gatewayInterfaces?: string[];
    ipAddresses?: string[];
    ipSets?: string[];
    orgNetworks?: string[];
    vmIds?: string[];
}

export interface OrgLdapCustomSettings {
    authenticationMethod: string;
    baseDistinguishedName?: string;
    connectorType: string;
    groupAttributes: outputs.OrgLdapCustomSettingsGroupAttributes;
    isSsl?: boolean;
    password?: string;
    port: number;
    server: string;
    userAttributes: outputs.OrgLdapCustomSettingsUserAttributes;
    username?: string;
}

export interface OrgLdapCustomSettingsGroupAttributes {
    groupBackLinkIdentifier?: string;
    groupMembershipIdentifier: string;
    membership: string;
    name: string;
    objectClass: string;
    uniqueIdentifier: string;
}

export interface OrgLdapCustomSettingsUserAttributes {
    displayName: string;
    email: string;
    givenName: string;
    groupBackLinkIdentifier?: string;
    groupMembershipIdentifier: string;
    objectClass: string;
    surname: string;
    telephone: string;
    uniqueIdentifier: string;
    username: string;
}

export interface OrgMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface OrgVappLease {
    deleteOnStorageLeaseExpiration: boolean;
    maximumRuntimeLeaseInSec: number;
    maximumStorageLeaseInSec: number;
    powerOffOnRuntimeLeaseExpiration: boolean;
}

export interface OrgVappTemplateLease {
    deleteOnStorageLeaseExpiration: boolean;
    maximumStorageLeaseInSec: number;
}

export interface OrgVdcAccessControlSharedWith {
    accessLevel: string;
    groupId?: string;
    subjectName: string;
    userId?: string;
}

export interface OrgVdcComputeCapacity {
    cpu: outputs.OrgVdcComputeCapacityCpu;
    memory: outputs.OrgVdcComputeCapacityMemory;
}

export interface OrgVdcComputeCapacityCpu {
    allocated: number;
    limit: number;
    reserved: number;
    used: number;
}

export interface OrgVdcComputeCapacityMemory {
    allocated: number;
    limit: number;
    reserved: number;
    used: number;
}

export interface OrgVdcMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface OrgVdcStorageProfile {
    default: boolean;
    enabled?: boolean;
    limit: number;
    name: string;
    storageUsedInMb: number;
}

export interface VappAccessControlSharedWith {
    accessLevel: string;
    groupId?: string;
    subjectName: string;
    userId?: string;
}

export interface VappFirewallRulesRule {
    destinationIp?: string;
    destinationPort?: string;
    destinationVmId?: string;
    destinationVmIpType?: string;
    destinationVmNicId?: number;
    enableLogging?: boolean;
    enabled?: boolean;
    name?: string;
    policy?: string;
    protocol?: string;
    sourceIp?: string;
    sourcePort?: string;
    sourceVmId?: string;
    sourceVmIpType?: string;
    sourceVmNicId?: number;
}

export interface VappLease {
    runtimeLeaseInSec: number;
    storageLeaseInSec: number;
}

export interface VappMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface VappNatRulesRule {
    externalIp: string;
    externalPort?: number;
    forwardToPort?: number;
    id: string;
    mappingMode?: string;
    protocol?: string;
    vmId: string;
    vmNicId: number;
}

export interface VappNetworkDhcpPool {
    defaultLeaseTime?: number;
    enabled?: boolean;
    endAddress?: string;
    maxLeaseTime?: number;
    startAddress: string;
}

export interface VappNetworkStaticIpPool {
    endAddress: string;
    startAddress: string;
}

export interface VappStaticRoutingRule {
    name: string;
    networkCidr: string;
    nextHopIp: string;
}

export interface VappVmCustomization {
    adminPassword: string;
    allowLocalAdminPassword: boolean;
    autoGeneratePassword: boolean;
    changeSid: boolean;
    enabled: boolean;
    force?: boolean;
    initscript: string;
    joinDomain: boolean;
    joinDomainAccountOu: string;
    joinDomainName: string;
    joinDomainPassword: string;
    joinDomainUser: string;
    joinOrgDomain: boolean;
    mustChangePasswordOnFirstLogin: boolean;
    numberOfAutoLogons: number;
}

export interface VappVmDisk {
    busNumber: string;
    name: string;
    sizeInMb: number;
    unitNumber: string;
}

export interface VappVmInternalDisk {
    busNumber: number;
    busType: string;
    diskId: string;
    iops: number;
    sizeInMb: number;
    storageProfile: string;
    thinProvisioned: boolean;
    unitNumber: number;
}

export interface VappVmMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface VappVmNetwork {
    adapterType: string;
    connected?: boolean;
    ip: string;
    ipAllocationMode?: string;
    isPrimary: boolean;
    mac: string;
    name?: string;
    type: string;
}

export interface VappVmOverrideTemplateDisk {
    busNumber: number;
    busType: string;
    iops?: number;
    sizeInMb: number;
    storageProfile?: string;
    unitNumber: number;
}

export interface VdcGroupParticipatingOrgVdc {
    faultDomainTag: string;
    isRemoteOrg: boolean;
    networkProviderScope: string;
    orgId: string;
    orgName: string;
    siteId: string;
    siteName: string;
    status: string;
    vdcId: string;
    vdcName: string;
}

export interface VmCustomization {
    adminPassword: string;
    allowLocalAdminPassword: boolean;
    autoGeneratePassword: boolean;
    changeSid: boolean;
    enabled: boolean;
    force?: boolean;
    initscript: string;
    joinDomain: boolean;
    joinDomainAccountOu: string;
    joinDomainName: string;
    joinDomainPassword: string;
    joinDomainUser: string;
    joinOrgDomain: boolean;
    mustChangePasswordOnFirstLogin: boolean;
    numberOfAutoLogons: number;
}

export interface VmDisk {
    busNumber: string;
    name: string;
    sizeInMb: number;
    unitNumber: string;
}

export interface VmInternalDisk {
    busNumber: number;
    busType: string;
    diskId: string;
    iops: number;
    sizeInMb: number;
    storageProfile: string;
    thinProvisioned: boolean;
    unitNumber: number;
}

export interface VmMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface VmNetwork {
    adapterType: string;
    connected?: boolean;
    ip: string;
    ipAllocationMode?: string;
    isPrimary: boolean;
    mac: string;
    name?: string;
    type: string;
}

export interface VmOverrideTemplateDisk {
    busNumber: number;
    busType: string;
    iops?: number;
    sizeInMb: number;
    storageProfile?: string;
    unitNumber: number;
}

export interface VmSizingPolicyCpu {
    coresPerSocket?: string;
    count?: string;
    limitInMhz?: string;
    reservationGuarantee?: string;
    shares?: string;
    speedInMhz?: string;
}

export interface VmSizingPolicyMemory {
    limitInMb?: string;
    reservationGuarantee?: string;
    shares?: string;
    sizeInMb?: string;
}

