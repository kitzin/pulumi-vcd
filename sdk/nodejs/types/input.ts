// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface CatalogAccessControlSharedWith {
    accessLevel: pulumi.Input<string>;
    groupId?: pulumi.Input<string>;
    orgId?: pulumi.Input<string>;
    subjectName?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
}

export interface CatalogItemMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface CatalogMediaMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface CatalogMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface CatalogVappTemplateMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface EdgeGatewayExternalNetwork {
    enableRateLimit?: pulumi.Input<boolean>;
    incomingRateLimit?: pulumi.Input<number>;
    name: pulumi.Input<string>;
    outgoingRateLimit?: pulumi.Input<number>;
    subnets?: pulumi.Input<pulumi.Input<inputs.EdgeGatewayExternalNetworkSubnet>[]>;
}

export interface EdgeGatewayExternalNetworkSubnet {
    gateway: pulumi.Input<string>;
    ipAddress?: pulumi.Input<string>;
    netmask: pulumi.Input<string>;
    suballocatePools?: pulumi.Input<pulumi.Input<inputs.EdgeGatewayExternalNetworkSubnetSuballocatePool>[]>;
    useForDefaultRoute?: pulumi.Input<boolean>;
}

export interface EdgeGatewayExternalNetworkSubnetSuballocatePool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface EdgeGatewayVpnLocalSubnet {
    localSubnetGateway: pulumi.Input<string>;
    localSubnetMask: pulumi.Input<string>;
    localSubnetName: pulumi.Input<string>;
}

export interface EdgeGatewayVpnPeerSubnet {
    peerSubnetGateway: pulumi.Input<string>;
    peerSubnetMask: pulumi.Input<string>;
    peerSubnetName: pulumi.Input<string>;
}

export interface ExternalNetworkIpScope {
    dns1?: pulumi.Input<string>;
    dns2?: pulumi.Input<string>;
    dnsSuffix?: pulumi.Input<string>;
    gateway: pulumi.Input<string>;
    netmask: pulumi.Input<string>;
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.ExternalNetworkIpScopeStaticIpPool>[]>;
}

export interface ExternalNetworkIpScopeStaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface ExternalNetworkV2IpScope {
    dns1?: pulumi.Input<string>;
    dns2?: pulumi.Input<string>;
    dnsSuffix?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    gateway: pulumi.Input<string>;
    prefixLength: pulumi.Input<number>;
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.ExternalNetworkV2IpScopeStaticIpPool>[]>;
}

export interface ExternalNetworkV2IpScopeStaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface ExternalNetworkV2NsxtNetwork {
    nsxtManagerId: pulumi.Input<string>;
    nsxtSegmentName?: pulumi.Input<string>;
    nsxtTier0RouterId?: pulumi.Input<string>;
}

export interface ExternalNetworkV2VsphereNetwork {
    portgroupId: pulumi.Input<string>;
    vcenterId: pulumi.Input<string>;
}

export interface ExternalNetworkVsphereNetwork {
    name: pulumi.Input<string>;
    type: pulumi.Input<string>;
    vcenter: pulumi.Input<string>;
}

export interface GetCatalogFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: inputs.GetCatalogFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogFilterArgs {
    date?: pulumi.Input<string>;
    earliest?: pulumi.Input<boolean>;
    latest?: pulumi.Input<boolean>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetCatalogFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetCatalogFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetCatalogItemFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: inputs.GetCatalogItemFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogItemFilterArgs {
    date?: pulumi.Input<string>;
    earliest?: pulumi.Input<boolean>;
    latest?: pulumi.Input<boolean>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetCatalogItemFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetCatalogItemFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogItemFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetCatalogMediaFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: inputs.GetCatalogMediaFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogMediaFilterArgs {
    date?: pulumi.Input<string>;
    earliest?: pulumi.Input<boolean>;
    latest?: pulumi.Input<boolean>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetCatalogMediaFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetCatalogMediaFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogMediaFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetCatalogVappTemplateFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: inputs.GetCatalogVappTemplateFilterMetadata[];
    nameRegex?: string;
}

export interface GetCatalogVappTemplateFilterArgs {
    date?: pulumi.Input<string>;
    earliest?: pulumi.Input<boolean>;
    latest?: pulumi.Input<boolean>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetCatalogVappTemplateFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetCatalogVappTemplateFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetCatalogVappTemplateFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetEdgegatewayFilter {
    nameRegex?: string;
}

export interface GetEdgegatewayFilterArgs {
    nameRegex?: pulumi.Input<string>;
}

export interface GetNetworkDirectFilter {
    ip?: string;
    metadatas?: inputs.GetNetworkDirectFilterMetadata[];
    nameRegex?: string;
}

export interface GetNetworkDirectFilterArgs {
    ip?: pulumi.Input<string>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetNetworkDirectFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetNetworkDirectFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetNetworkDirectFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetNetworkIsolatedFilterArgs {
    ip?: pulumi.Input<string>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetNetworkIsolatedFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetNetworkIsolatedFilter {
    ip?: string;
    metadatas?: inputs.GetNetworkIsolatedFilterMetadata[];
    nameRegex?: string;
}

export interface GetNetworkIsolatedFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetNetworkIsolatedFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetNetworkIsolatedV2Filter {
    ip?: string;
    nameRegex?: string;
}

export interface GetNetworkIsolatedV2FilterArgs {
    ip?: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetNetworkRoutedFilter {
    ip?: string;
    metadatas?: inputs.GetNetworkRoutedFilterMetadata[];
    nameRegex?: string;
}

export interface GetNetworkRoutedFilterArgs {
    ip?: pulumi.Input<string>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetNetworkRoutedFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetNetworkRoutedFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface GetNetworkRoutedFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetNetworkRoutedV2Filter {
    ip?: string;
    nameRegex?: string;
}

export interface GetNetworkRoutedV2FilterArgs {
    ip?: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetNsxtNetworkImportedFilter {
    ip?: string;
    nameRegex?: string;
}

export interface GetNsxtNetworkImportedFilterArgs {
    ip?: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetOrgVdcMetadataEntry {
    isSystem?: boolean;
    key?: string;
    type?: string;
    userAccess?: string;
    value?: string;
}

export interface GetOrgVdcMetadataEntryArgs {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface GetSubscribedCatalogFilterArgs {
    date?: pulumi.Input<string>;
    earliest?: pulumi.Input<boolean>;
    latest?: pulumi.Input<boolean>;
    metadatas?: pulumi.Input<pulumi.Input<inputs.GetSubscribedCatalogFilterMetadataArgs>[]>;
    nameRegex?: pulumi.Input<string>;
}

export interface GetSubscribedCatalogFilter {
    date?: string;
    earliest?: boolean;
    latest?: boolean;
    metadatas?: inputs.GetSubscribedCatalogFilterMetadata[];
    nameRegex?: string;
}

export interface GetSubscribedCatalogFilterMetadata {
    isSystem?: boolean;
    key: string;
    type?: string;
    useApiSearch?: boolean;
    value: string;
}

export interface GetSubscribedCatalogFilterMetadataArgs {
    isSystem?: pulumi.Input<boolean>;
    key: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    useApiSearch?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export interface IndependentDiskMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface LbServerPoolMember {
    condition: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    ipAddress: pulumi.Input<string>;
    maxConnections?: pulumi.Input<number>;
    minConnections?: pulumi.Input<number>;
    monitorPort: pulumi.Input<number>;
    name: pulumi.Input<string>;
    port: pulumi.Input<number>;
    weight: pulumi.Input<number>;
}

export interface NetworkDirectMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface NetworkIsolatedDhcpPool {
    defaultLeaseTime?: pulumi.Input<number>;
    endAddress: pulumi.Input<string>;
    maxLeaseTime?: pulumi.Input<number>;
    startAddress: pulumi.Input<string>;
}

export interface NetworkIsolatedMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface NetworkIsolatedStaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NetworkIsolatedV2MetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface NetworkIsolatedV2StaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NetworkRoutedDhcpPool {
    defaultLeaseTime?: pulumi.Input<number>;
    endAddress: pulumi.Input<string>;
    maxLeaseTime?: pulumi.Input<number>;
    startAddress: pulumi.Input<string>;
}

export interface NetworkRoutedMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface NetworkRoutedStaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NetworkRoutedV2MetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface NetworkRoutedV2StaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NsxtAlbPoolHealthMonitor {
    name?: pulumi.Input<string>;
    systemDefined?: pulumi.Input<boolean>;
    type: pulumi.Input<string>;
}

export interface NsxtAlbPoolMember {
    detailedHealthMessage?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    healthStatus?: pulumi.Input<string>;
    ipAddress: pulumi.Input<string>;
    markedDownBies?: pulumi.Input<pulumi.Input<string>[]>;
    port?: pulumi.Input<number>;
    ratio?: pulumi.Input<number>;
}

export interface NsxtAlbPoolPersistenceProfile {
    name?: pulumi.Input<string>;
    type: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface NsxtAlbVirtualServiceServicePort {
    endPort?: pulumi.Input<number>;
    sslEnabled?: pulumi.Input<boolean>;
    startPort: pulumi.Input<number>;
    type: pulumi.Input<string>;
}

export interface NsxtAppPortProfileAppPort {
    ports?: pulumi.Input<pulumi.Input<string>[]>;
    protocol: pulumi.Input<string>;
}

export interface NsxtDistributedFirewallRule {
    action: pulumi.Input<string>;
    appPortProfileIds?: pulumi.Input<pulumi.Input<string>[]>;
    comment?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    destinationGroupsExcluded?: pulumi.Input<boolean>;
    destinationIds?: pulumi.Input<pulumi.Input<string>[]>;
    direction?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    ipProtocol?: pulumi.Input<string>;
    logging?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    networkContextProfileIds?: pulumi.Input<pulumi.Input<string>[]>;
    sourceGroupsExcluded?: pulumi.Input<boolean>;
    sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface NsxtDynamicSecurityGroupCriteria {
    rules?: pulumi.Input<pulumi.Input<inputs.NsxtDynamicSecurityGroupCriteriaRule>[]>;
}

export interface NsxtDynamicSecurityGroupCriteriaRule {
    operator: pulumi.Input<string>;
    type: pulumi.Input<string>;
    value: pulumi.Input<string>;
}

export interface NsxtDynamicSecurityGroupMemberVm {
    vappId?: pulumi.Input<string>;
    vappName?: pulumi.Input<string>;
    vmId?: pulumi.Input<string>;
    vmName?: pulumi.Input<string>;
}

export interface NsxtEdgegatewayBgpIpPrefixListIpPrefix {
    action: pulumi.Input<string>;
    greaterThanOrEqualTo?: pulumi.Input<number>;
    lessThanOrEqualTo?: pulumi.Input<number>;
    network: pulumi.Input<string>;
}

export interface NsxtEdgegatewaySubnet {
    allocatedIps: pulumi.Input<pulumi.Input<inputs.NsxtEdgegatewaySubnetAllocatedIp>[]>;
    gateway: pulumi.Input<string>;
    prefixLength: pulumi.Input<number>;
    primaryIp?: pulumi.Input<string>;
}

export interface NsxtEdgegatewaySubnetAllocatedIp {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NsxtFirewallRule {
    action: pulumi.Input<string>;
    appPortProfileIds?: pulumi.Input<pulumi.Input<string>[]>;
    destinationIds?: pulumi.Input<pulumi.Input<string>[]>;
    direction: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    ipProtocol: pulumi.Input<string>;
    logging?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface NsxtIpsecVpnTunnelSecurityProfileCustomization {
    dpdProbeInternal?: pulumi.Input<number>;
    ikeDhGroups: pulumi.Input<pulumi.Input<string>[]>;
    ikeDigestAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    ikeEncryptionAlgorithms: pulumi.Input<pulumi.Input<string>[]>;
    ikeSaLifetime?: pulumi.Input<number>;
    ikeVersion: pulumi.Input<string>;
    tunnelDfPolicy?: pulumi.Input<string>;
    tunnelDhGroups: pulumi.Input<pulumi.Input<string>[]>;
    tunnelDigestAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    tunnelEncryptionAlgorithms: pulumi.Input<pulumi.Input<string>[]>;
    tunnelPfsEnabled?: pulumi.Input<boolean>;
    tunnelSaLifetime?: pulumi.Input<number>;
}

export interface NsxtNetworkDhcpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NsxtNetworkImportedStaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface NsxtSecurityGroupMemberVm {
    vappId?: pulumi.Input<string>;
    vappName?: pulumi.Input<string>;
    vmId?: pulumi.Input<string>;
    vmName?: pulumi.Input<string>;
}

export interface NsxvDhcpRelayRelayAgent {
    gatewayIpAddress?: pulumi.Input<string>;
    networkName: pulumi.Input<string>;
}

export interface NsxvFirewallRuleDestination {
    exclude?: pulumi.Input<boolean>;
    gatewayInterfaces?: pulumi.Input<pulumi.Input<string>[]>;
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    ipSets?: pulumi.Input<pulumi.Input<string>[]>;
    orgNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    vmIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface NsxvFirewallRuleService {
    port?: pulumi.Input<string>;
    protocol: pulumi.Input<string>;
    sourcePort?: pulumi.Input<string>;
}

export interface NsxvFirewallRuleSource {
    exclude?: pulumi.Input<boolean>;
    gatewayInterfaces?: pulumi.Input<pulumi.Input<string>[]>;
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    ipSets?: pulumi.Input<pulumi.Input<string>[]>;
    orgNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    vmIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface OrgLdapCustomSettings {
    authenticationMethod: pulumi.Input<string>;
    baseDistinguishedName?: pulumi.Input<string>;
    connectorType: pulumi.Input<string>;
    groupAttributes: pulumi.Input<inputs.OrgLdapCustomSettingsGroupAttributes>;
    isSsl?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    port: pulumi.Input<number>;
    server: pulumi.Input<string>;
    userAttributes: pulumi.Input<inputs.OrgLdapCustomSettingsUserAttributes>;
    username?: pulumi.Input<string>;
}

export interface OrgLdapCustomSettingsGroupAttributes {
    groupBackLinkIdentifier?: pulumi.Input<string>;
    groupMembershipIdentifier: pulumi.Input<string>;
    membership: pulumi.Input<string>;
    name: pulumi.Input<string>;
    objectClass: pulumi.Input<string>;
    uniqueIdentifier: pulumi.Input<string>;
}

export interface OrgLdapCustomSettingsUserAttributes {
    displayName: pulumi.Input<string>;
    email: pulumi.Input<string>;
    givenName: pulumi.Input<string>;
    groupBackLinkIdentifier?: pulumi.Input<string>;
    groupMembershipIdentifier: pulumi.Input<string>;
    objectClass: pulumi.Input<string>;
    surname: pulumi.Input<string>;
    telephone: pulumi.Input<string>;
    uniqueIdentifier: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface OrgMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface OrgVappLease {
    deleteOnStorageLeaseExpiration: pulumi.Input<boolean>;
    maximumRuntimeLeaseInSec: pulumi.Input<number>;
    maximumStorageLeaseInSec: pulumi.Input<number>;
    powerOffOnRuntimeLeaseExpiration: pulumi.Input<boolean>;
}

export interface OrgVappTemplateLease {
    deleteOnStorageLeaseExpiration: pulumi.Input<boolean>;
    maximumStorageLeaseInSec: pulumi.Input<number>;
}

export interface OrgVdcAccessControlSharedWith {
    accessLevel: pulumi.Input<string>;
    groupId?: pulumi.Input<string>;
    subjectName?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
}

export interface OrgVdcComputeCapacity {
    cpu: pulumi.Input<inputs.OrgVdcComputeCapacityCpu>;
    memory: pulumi.Input<inputs.OrgVdcComputeCapacityMemory>;
}

export interface OrgVdcComputeCapacityCpu {
    allocated?: pulumi.Input<number>;
    limit?: pulumi.Input<number>;
    reserved?: pulumi.Input<number>;
    used?: pulumi.Input<number>;
}

export interface OrgVdcComputeCapacityMemory {
    allocated?: pulumi.Input<number>;
    limit?: pulumi.Input<number>;
    reserved?: pulumi.Input<number>;
    used?: pulumi.Input<number>;
}

export interface OrgVdcMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface OrgVdcStorageProfile {
    default: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    limit: pulumi.Input<number>;
    name: pulumi.Input<string>;
    storageUsedInMb?: pulumi.Input<number>;
}

export interface VappAccessControlSharedWith {
    accessLevel: pulumi.Input<string>;
    groupId?: pulumi.Input<string>;
    subjectName?: pulumi.Input<string>;
    userId?: pulumi.Input<string>;
}

export interface VappFirewallRulesRule {
    destinationIp?: pulumi.Input<string>;
    destinationPort?: pulumi.Input<string>;
    destinationVmId?: pulumi.Input<string>;
    destinationVmIpType?: pulumi.Input<string>;
    destinationVmNicId?: pulumi.Input<number>;
    enableLogging?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    policy?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    sourceIp?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<string>;
    sourceVmId?: pulumi.Input<string>;
    sourceVmIpType?: pulumi.Input<string>;
    sourceVmNicId?: pulumi.Input<number>;
}

export interface VappLease {
    runtimeLeaseInSec: pulumi.Input<number>;
    storageLeaseInSec: pulumi.Input<number>;
}

export interface VappMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface VappNatRulesRule {
    externalIp?: pulumi.Input<string>;
    externalPort?: pulumi.Input<number>;
    forwardToPort?: pulumi.Input<number>;
    id?: pulumi.Input<string>;
    mappingMode?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    vmId: pulumi.Input<string>;
    vmNicId: pulumi.Input<number>;
}

export interface VappNetworkDhcpPool {
    defaultLeaseTime?: pulumi.Input<number>;
    enabled?: pulumi.Input<boolean>;
    endAddress?: pulumi.Input<string>;
    maxLeaseTime?: pulumi.Input<number>;
    startAddress: pulumi.Input<string>;
}

export interface VappNetworkStaticIpPool {
    endAddress: pulumi.Input<string>;
    startAddress: pulumi.Input<string>;
}

export interface VappStaticRoutingRule {
    name: pulumi.Input<string>;
    networkCidr: pulumi.Input<string>;
    nextHopIp: pulumi.Input<string>;
}

export interface VappVmCustomization {
    adminPassword?: pulumi.Input<string>;
    allowLocalAdminPassword?: pulumi.Input<boolean>;
    autoGeneratePassword?: pulumi.Input<boolean>;
    changeSid?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    force?: pulumi.Input<boolean>;
    initscript?: pulumi.Input<string>;
    joinDomain?: pulumi.Input<boolean>;
    joinDomainAccountOu?: pulumi.Input<string>;
    joinDomainName?: pulumi.Input<string>;
    joinDomainPassword?: pulumi.Input<string>;
    joinDomainUser?: pulumi.Input<string>;
    joinOrgDomain?: pulumi.Input<boolean>;
    mustChangePasswordOnFirstLogin?: pulumi.Input<boolean>;
    numberOfAutoLogons?: pulumi.Input<number>;
}

export interface VappVmDisk {
    busNumber: pulumi.Input<string>;
    name: pulumi.Input<string>;
    sizeInMb?: pulumi.Input<number>;
    unitNumber: pulumi.Input<string>;
}

export interface VappVmInternalDisk {
    busNumber?: pulumi.Input<number>;
    busType?: pulumi.Input<string>;
    diskId?: pulumi.Input<string>;
    iops?: pulumi.Input<number>;
    sizeInMb?: pulumi.Input<number>;
    storageProfile?: pulumi.Input<string>;
    thinProvisioned?: pulumi.Input<boolean>;
    unitNumber?: pulumi.Input<number>;
}

export interface VappVmMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface VappVmNetwork {
    adapterType?: pulumi.Input<string>;
    connected?: pulumi.Input<boolean>;
    ip?: pulumi.Input<string>;
    ipAllocationMode?: pulumi.Input<string>;
    isPrimary?: pulumi.Input<boolean>;
    mac?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface VappVmOverrideTemplateDisk {
    busNumber: pulumi.Input<number>;
    busType: pulumi.Input<string>;
    iops?: pulumi.Input<number>;
    sizeInMb: pulumi.Input<number>;
    storageProfile?: pulumi.Input<string>;
    unitNumber: pulumi.Input<number>;
}

export interface VdcGroupParticipatingOrgVdc {
    faultDomainTag?: pulumi.Input<string>;
    isRemoteOrg?: pulumi.Input<boolean>;
    networkProviderScope?: pulumi.Input<string>;
    orgId?: pulumi.Input<string>;
    orgName?: pulumi.Input<string>;
    siteId?: pulumi.Input<string>;
    siteName?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdcId?: pulumi.Input<string>;
    vdcName?: pulumi.Input<string>;
}

export interface VmCustomization {
    adminPassword?: pulumi.Input<string>;
    allowLocalAdminPassword?: pulumi.Input<boolean>;
    autoGeneratePassword?: pulumi.Input<boolean>;
    changeSid?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    force?: pulumi.Input<boolean>;
    initscript?: pulumi.Input<string>;
    joinDomain?: pulumi.Input<boolean>;
    joinDomainAccountOu?: pulumi.Input<string>;
    joinDomainName?: pulumi.Input<string>;
    joinDomainPassword?: pulumi.Input<string>;
    joinDomainUser?: pulumi.Input<string>;
    joinOrgDomain?: pulumi.Input<boolean>;
    mustChangePasswordOnFirstLogin?: pulumi.Input<boolean>;
    numberOfAutoLogons?: pulumi.Input<number>;
}

export interface VmDisk {
    busNumber: pulumi.Input<string>;
    name: pulumi.Input<string>;
    sizeInMb?: pulumi.Input<number>;
    unitNumber: pulumi.Input<string>;
}

export interface VmInternalDisk {
    busNumber?: pulumi.Input<number>;
    busType?: pulumi.Input<string>;
    diskId?: pulumi.Input<string>;
    iops?: pulumi.Input<number>;
    sizeInMb?: pulumi.Input<number>;
    storageProfile?: pulumi.Input<string>;
    thinProvisioned?: pulumi.Input<boolean>;
    unitNumber?: pulumi.Input<number>;
}

export interface VmMetadataEntry {
    isSystem?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userAccess?: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface VmNetwork {
    adapterType?: pulumi.Input<string>;
    connected?: pulumi.Input<boolean>;
    ip?: pulumi.Input<string>;
    ipAllocationMode?: pulumi.Input<string>;
    isPrimary?: pulumi.Input<boolean>;
    mac?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface VmOverrideTemplateDisk {
    busNumber: pulumi.Input<number>;
    busType: pulumi.Input<string>;
    iops?: pulumi.Input<number>;
    sizeInMb: pulumi.Input<number>;
    storageProfile?: pulumi.Input<string>;
    unitNumber: pulumi.Input<number>;
}

export interface VmSizingPolicyCpu {
    coresPerSocket?: pulumi.Input<string>;
    count?: pulumi.Input<string>;
    limitInMhz?: pulumi.Input<string>;
    reservationGuarantee?: pulumi.Input<string>;
    shares?: pulumi.Input<string>;
    speedInMhz?: pulumi.Input<string>;
}

export interface VmSizingPolicyMemory {
    limitInMb?: pulumi.Input<string>;
    reservationGuarantee?: pulumi.Input<string>;
    shares?: pulumi.Input<string>;
    sizeInMb?: pulumi.Input<string>;
}
