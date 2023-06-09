// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./catalog";
export * from "./catalogAccessControl";
export * from "./catalogItem";
export * from "./catalogMedia";
export * from "./catalogVappTemplate";
export * from "./edgeGateway";
export * from "./edgeGatewaySettings";
export * from "./edgeGatewayVpn";
export * from "./externalNetwork";
export * from "./externalNetworkV2";
export * from "./getCatalog";
export * from "./getCatalogItem";
export * from "./getCatalogMedia";
export * from "./getCatalogVappTemplate";
export * from "./getEdgegateway";
export * from "./getExternalNetwork";
export * from "./getExternalNetworkV2";
export * from "./getGlobalRole";
export * from "./getIndependentDisk";
export * from "./getLbAppProfile";
export * from "./getLbAppRule";
export * from "./getLbServerPool";
export * from "./getLbServiceMonitor";
export * from "./getLbVirtualServer";
export * from "./getLibraryCertificate";
export * from "./getNetworkDirect";
export * from "./getNetworkIsolated";
export * from "./getNetworkIsolatedV2";
export * from "./getNetworkRouted";
export * from "./getNetworkRoutedV2";
export * from "./getNsxtAlbCloud";
export * from "./getNsxtAlbController";
export * from "./getNsxtAlbEdgegatewayServiceEngineGroup";
export * from "./getNsxtAlbImportableCloud";
export * from "./getNsxtAlbPool";
export * from "./getNsxtAlbServiceEngineGroup";
export * from "./getNsxtAlbSettings";
export * from "./getNsxtAlbVirtualService";
export * from "./getNsxtAppPortProfile";
export * from "./getNsxtDistributedFirewall";
export * from "./getNsxtDynamicSecurityGroup";
export * from "./getNsxtEdgeCluster";
export * from "./getNsxtEdgegateway";
export * from "./getNsxtEdgegatewayBgpConfiguration";
export * from "./getNsxtEdgegatewayBgpIpPrefixList";
export * from "./getNsxtEdgegatewayBgpNeighbor";
export * from "./getNsxtFirewall";
export * from "./getNsxtIpSet";
export * from "./getNsxtIpsecVpnTunnel";
export * from "./getNsxtManager";
export * from "./getNsxtNatRule";
export * from "./getNsxtNetworkContextProfile";
export * from "./getNsxtNetworkDhcp";
export * from "./getNsxtNetworkImported";
export * from "./getNsxtRouteAdvertisement";
export * from "./getNsxtSecurityGroup";
export * from "./getNsxtTier0Router";
export * from "./getNsxvDhcpRelay";
export * from "./getNsxvDnat";
export * from "./getNsxvFirewallRule";
export * from "./getNsxvIpSet";
export * from "./getNsxvSnat";
export * from "./getOrg";
export * from "./getOrgGroup";
export * from "./getOrgLdap";
export * from "./getOrgUser";
export * from "./getOrgVdc";
export * from "./getPortgroup";
export * from "./getProviderVdc";
export * from "./getResourceList";
export * from "./getResourceSchema";
export * from "./getRight";
export * from "./getRightsBundle";
export * from "./getRole";
export * from "./getStorageProfile";
export * from "./getSubscribedCatalog";
export * from "./getTask";
export * from "./getVapp";
export * from "./getVappNetwork";
export * from "./getVappOrgNetwork";
export * from "./getVappVm";
export * from "./getVcenter";
export * from "./getVdcGroup";
export * from "./getVm";
export * from "./getVmAffinityRule";
export * from "./getVmGroup";
export * from "./getVmPlacementPolicy";
export * from "./getVmSizingPolicy";
export * from "./globalRole";
export * from "./independentDisk";
export * from "./insertedMedia";
export * from "./lbAppProfile";
export * from "./lbAppRule";
export * from "./lbServerPool";
export * from "./lbServiceMonitor";
export * from "./lbVirtualServer";
export * from "./libraryCertificate";
export * from "./networkDirect";
export * from "./networkIsolated";
export * from "./networkIsolatedV2";
export * from "./networkRouted";
export * from "./networkRoutedV2";
export * from "./nsxtAlbCloud";
export * from "./nsxtAlbController";
export * from "./nsxtAlbEdgegatewayServiceEngineGroup";
export * from "./nsxtAlbPool";
export * from "./nsxtAlbServiceEngineGroup";
export * from "./nsxtAlbSettings";
export * from "./nsxtAlbVirtualService";
export * from "./nsxtAppPortProfile";
export * from "./nsxtDistributedFirewall";
export * from "./nsxtDynamicSecurityGroup";
export * from "./nsxtEdgegateway";
export * from "./nsxtEdgegatewayBgpConfiguration";
export * from "./nsxtEdgegatewayBgpIpPrefixList";
export * from "./nsxtEdgegatewayBgpNeighbor";
export * from "./nsxtFirewall";
export * from "./nsxtIpSet";
export * from "./nsxtIpsecVpnTunnel";
export * from "./nsxtNatRule";
export * from "./nsxtNetworkDhcp";
export * from "./nsxtNetworkImported";
export * from "./nsxtRouteAdvertisement";
export * from "./nsxtSecurityGroup";
export * from "./nsxvDhcpRelay";
export * from "./nsxvDnat";
export * from "./nsxvFirewallRule";
export * from "./nsxvIpSet";
export * from "./nsxvSnat";
export * from "./org";
export * from "./orgGroup";
export * from "./orgLdap";
export * from "./orgUser";
export * from "./orgVdc";
export * from "./orgVdcAccessControl";
export * from "./provider";
export * from "./rightsBundle";
export * from "./role";
export * from "./securityTag";
export * from "./subscribedCatalog";
export * from "./vapp";
export * from "./vappAccessControl";
export * from "./vappFirewallRules";
export * from "./vappNatRules";
export * from "./vappNetwork";
export * from "./vappOrgNetwork";
export * from "./vappStaticRouting";
export * from "./vappVm";
export * from "./vdcGroup";
export * from "./vm";
export * from "./vmAffinityRule";
export * from "./vmInternalDisk";
export * from "./vmPlacementPolicy";
export * from "./vmSizingPolicy";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Catalog } from "./catalog";
import { CatalogAccessControl } from "./catalogAccessControl";
import { CatalogItem } from "./catalogItem";
import { CatalogMedia } from "./catalogMedia";
import { CatalogVappTemplate } from "./catalogVappTemplate";
import { EdgeGateway } from "./edgeGateway";
import { EdgeGatewaySettings } from "./edgeGatewaySettings";
import { EdgeGatewayVpn } from "./edgeGatewayVpn";
import { ExternalNetwork } from "./externalNetwork";
import { ExternalNetworkV2 } from "./externalNetworkV2";
import { GlobalRole } from "./globalRole";
import { IndependentDisk } from "./independentDisk";
import { InsertedMedia } from "./insertedMedia";
import { LbAppProfile } from "./lbAppProfile";
import { LbAppRule } from "./lbAppRule";
import { LbServerPool } from "./lbServerPool";
import { LbServiceMonitor } from "./lbServiceMonitor";
import { LbVirtualServer } from "./lbVirtualServer";
import { LibraryCertificate } from "./libraryCertificate";
import { NetworkDirect } from "./networkDirect";
import { NetworkIsolated } from "./networkIsolated";
import { NetworkIsolatedV2 } from "./networkIsolatedV2";
import { NetworkRouted } from "./networkRouted";
import { NetworkRoutedV2 } from "./networkRoutedV2";
import { NsxtAlbCloud } from "./nsxtAlbCloud";
import { NsxtAlbController } from "./nsxtAlbController";
import { NsxtAlbEdgegatewayServiceEngineGroup } from "./nsxtAlbEdgegatewayServiceEngineGroup";
import { NsxtAlbPool } from "./nsxtAlbPool";
import { NsxtAlbServiceEngineGroup } from "./nsxtAlbServiceEngineGroup";
import { NsxtAlbSettings } from "./nsxtAlbSettings";
import { NsxtAlbVirtualService } from "./nsxtAlbVirtualService";
import { NsxtAppPortProfile } from "./nsxtAppPortProfile";
import { NsxtDistributedFirewall } from "./nsxtDistributedFirewall";
import { NsxtDynamicSecurityGroup } from "./nsxtDynamicSecurityGroup";
import { NsxtEdgegateway } from "./nsxtEdgegateway";
import { NsxtEdgegatewayBgpConfiguration } from "./nsxtEdgegatewayBgpConfiguration";
import { NsxtEdgegatewayBgpIpPrefixList } from "./nsxtEdgegatewayBgpIpPrefixList";
import { NsxtEdgegatewayBgpNeighbor } from "./nsxtEdgegatewayBgpNeighbor";
import { NsxtFirewall } from "./nsxtFirewall";
import { NsxtIpSet } from "./nsxtIpSet";
import { NsxtIpsecVpnTunnel } from "./nsxtIpsecVpnTunnel";
import { NsxtNatRule } from "./nsxtNatRule";
import { NsxtNetworkDhcp } from "./nsxtNetworkDhcp";
import { NsxtNetworkImported } from "./nsxtNetworkImported";
import { NsxtRouteAdvertisement } from "./nsxtRouteAdvertisement";
import { NsxtSecurityGroup } from "./nsxtSecurityGroup";
import { NsxvDhcpRelay } from "./nsxvDhcpRelay";
import { NsxvDnat } from "./nsxvDnat";
import { NsxvFirewallRule } from "./nsxvFirewallRule";
import { NsxvIpSet } from "./nsxvIpSet";
import { NsxvSnat } from "./nsxvSnat";
import { Org } from "./org";
import { OrgGroup } from "./orgGroup";
import { OrgLdap } from "./orgLdap";
import { OrgUser } from "./orgUser";
import { OrgVdc } from "./orgVdc";
import { OrgVdcAccessControl } from "./orgVdcAccessControl";
import { RightsBundle } from "./rightsBundle";
import { Role } from "./role";
import { SecurityTag } from "./securityTag";
import { SubscribedCatalog } from "./subscribedCatalog";
import { Vapp } from "./vapp";
import { VappAccessControl } from "./vappAccessControl";
import { VappFirewallRules } from "./vappFirewallRules";
import { VappNatRules } from "./vappNatRules";
import { VappNetwork } from "./vappNetwork";
import { VappOrgNetwork } from "./vappOrgNetwork";
import { VappStaticRouting } from "./vappStaticRouting";
import { VappVm } from "./vappVm";
import { VdcGroup } from "./vdcGroup";
import { Vm } from "./vm";
import { VmAffinityRule } from "./vmAffinityRule";
import { VmInternalDisk } from "./vmInternalDisk";
import { VmPlacementPolicy } from "./vmPlacementPolicy";
import { VmSizingPolicy } from "./vmSizingPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vcd:index/catalog:Catalog":
                return new Catalog(name, <any>undefined, { urn })
            case "vcd:index/catalogAccessControl:CatalogAccessControl":
                return new CatalogAccessControl(name, <any>undefined, { urn })
            case "vcd:index/catalogItem:CatalogItem":
                return new CatalogItem(name, <any>undefined, { urn })
            case "vcd:index/catalogMedia:CatalogMedia":
                return new CatalogMedia(name, <any>undefined, { urn })
            case "vcd:index/catalogVappTemplate:CatalogVappTemplate":
                return new CatalogVappTemplate(name, <any>undefined, { urn })
            case "vcd:index/edgeGateway:EdgeGateway":
                return new EdgeGateway(name, <any>undefined, { urn })
            case "vcd:index/edgeGatewaySettings:EdgeGatewaySettings":
                return new EdgeGatewaySettings(name, <any>undefined, { urn })
            case "vcd:index/edgeGatewayVpn:EdgeGatewayVpn":
                return new EdgeGatewayVpn(name, <any>undefined, { urn })
            case "vcd:index/externalNetwork:ExternalNetwork":
                return new ExternalNetwork(name, <any>undefined, { urn })
            case "vcd:index/externalNetworkV2:ExternalNetworkV2":
                return new ExternalNetworkV2(name, <any>undefined, { urn })
            case "vcd:index/globalRole:GlobalRole":
                return new GlobalRole(name, <any>undefined, { urn })
            case "vcd:index/independentDisk:IndependentDisk":
                return new IndependentDisk(name, <any>undefined, { urn })
            case "vcd:index/insertedMedia:InsertedMedia":
                return new InsertedMedia(name, <any>undefined, { urn })
            case "vcd:index/lbAppProfile:LbAppProfile":
                return new LbAppProfile(name, <any>undefined, { urn })
            case "vcd:index/lbAppRule:LbAppRule":
                return new LbAppRule(name, <any>undefined, { urn })
            case "vcd:index/lbServerPool:LbServerPool":
                return new LbServerPool(name, <any>undefined, { urn })
            case "vcd:index/lbServiceMonitor:LbServiceMonitor":
                return new LbServiceMonitor(name, <any>undefined, { urn })
            case "vcd:index/lbVirtualServer:LbVirtualServer":
                return new LbVirtualServer(name, <any>undefined, { urn })
            case "vcd:index/libraryCertificate:LibraryCertificate":
                return new LibraryCertificate(name, <any>undefined, { urn })
            case "vcd:index/networkDirect:NetworkDirect":
                return new NetworkDirect(name, <any>undefined, { urn })
            case "vcd:index/networkIsolated:NetworkIsolated":
                return new NetworkIsolated(name, <any>undefined, { urn })
            case "vcd:index/networkIsolatedV2:NetworkIsolatedV2":
                return new NetworkIsolatedV2(name, <any>undefined, { urn })
            case "vcd:index/networkRouted:NetworkRouted":
                return new NetworkRouted(name, <any>undefined, { urn })
            case "vcd:index/networkRoutedV2:NetworkRoutedV2":
                return new NetworkRoutedV2(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbCloud:NsxtAlbCloud":
                return new NsxtAlbCloud(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbController:NsxtAlbController":
                return new NsxtAlbController(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbEdgegatewayServiceEngineGroup:NsxtAlbEdgegatewayServiceEngineGroup":
                return new NsxtAlbEdgegatewayServiceEngineGroup(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbPool:NsxtAlbPool":
                return new NsxtAlbPool(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbServiceEngineGroup:NsxtAlbServiceEngineGroup":
                return new NsxtAlbServiceEngineGroup(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbSettings:NsxtAlbSettings":
                return new NsxtAlbSettings(name, <any>undefined, { urn })
            case "vcd:index/nsxtAlbVirtualService:NsxtAlbVirtualService":
                return new NsxtAlbVirtualService(name, <any>undefined, { urn })
            case "vcd:index/nsxtAppPortProfile:NsxtAppPortProfile":
                return new NsxtAppPortProfile(name, <any>undefined, { urn })
            case "vcd:index/nsxtDistributedFirewall:NsxtDistributedFirewall":
                return new NsxtDistributedFirewall(name, <any>undefined, { urn })
            case "vcd:index/nsxtDynamicSecurityGroup:NsxtDynamicSecurityGroup":
                return new NsxtDynamicSecurityGroup(name, <any>undefined, { urn })
            case "vcd:index/nsxtEdgegateway:NsxtEdgegateway":
                return new NsxtEdgegateway(name, <any>undefined, { urn })
            case "vcd:index/nsxtEdgegatewayBgpConfiguration:NsxtEdgegatewayBgpConfiguration":
                return new NsxtEdgegatewayBgpConfiguration(name, <any>undefined, { urn })
            case "vcd:index/nsxtEdgegatewayBgpIpPrefixList:NsxtEdgegatewayBgpIpPrefixList":
                return new NsxtEdgegatewayBgpIpPrefixList(name, <any>undefined, { urn })
            case "vcd:index/nsxtEdgegatewayBgpNeighbor:NsxtEdgegatewayBgpNeighbor":
                return new NsxtEdgegatewayBgpNeighbor(name, <any>undefined, { urn })
            case "vcd:index/nsxtFirewall:NsxtFirewall":
                return new NsxtFirewall(name, <any>undefined, { urn })
            case "vcd:index/nsxtIpSet:NsxtIpSet":
                return new NsxtIpSet(name, <any>undefined, { urn })
            case "vcd:index/nsxtIpsecVpnTunnel:NsxtIpsecVpnTunnel":
                return new NsxtIpsecVpnTunnel(name, <any>undefined, { urn })
            case "vcd:index/nsxtNatRule:NsxtNatRule":
                return new NsxtNatRule(name, <any>undefined, { urn })
            case "vcd:index/nsxtNetworkDhcp:NsxtNetworkDhcp":
                return new NsxtNetworkDhcp(name, <any>undefined, { urn })
            case "vcd:index/nsxtNetworkImported:NsxtNetworkImported":
                return new NsxtNetworkImported(name, <any>undefined, { urn })
            case "vcd:index/nsxtRouteAdvertisement:NsxtRouteAdvertisement":
                return new NsxtRouteAdvertisement(name, <any>undefined, { urn })
            case "vcd:index/nsxtSecurityGroup:NsxtSecurityGroup":
                return new NsxtSecurityGroup(name, <any>undefined, { urn })
            case "vcd:index/nsxvDhcpRelay:NsxvDhcpRelay":
                return new NsxvDhcpRelay(name, <any>undefined, { urn })
            case "vcd:index/nsxvDnat:NsxvDnat":
                return new NsxvDnat(name, <any>undefined, { urn })
            case "vcd:index/nsxvFirewallRule:NsxvFirewallRule":
                return new NsxvFirewallRule(name, <any>undefined, { urn })
            case "vcd:index/nsxvIpSet:NsxvIpSet":
                return new NsxvIpSet(name, <any>undefined, { urn })
            case "vcd:index/nsxvSnat:NsxvSnat":
                return new NsxvSnat(name, <any>undefined, { urn })
            case "vcd:index/org:Org":
                return new Org(name, <any>undefined, { urn })
            case "vcd:index/orgGroup:OrgGroup":
                return new OrgGroup(name, <any>undefined, { urn })
            case "vcd:index/orgLdap:OrgLdap":
                return new OrgLdap(name, <any>undefined, { urn })
            case "vcd:index/orgUser:OrgUser":
                return new OrgUser(name, <any>undefined, { urn })
            case "vcd:index/orgVdc:OrgVdc":
                return new OrgVdc(name, <any>undefined, { urn })
            case "vcd:index/orgVdcAccessControl:OrgVdcAccessControl":
                return new OrgVdcAccessControl(name, <any>undefined, { urn })
            case "vcd:index/rightsBundle:RightsBundle":
                return new RightsBundle(name, <any>undefined, { urn })
            case "vcd:index/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "vcd:index/securityTag:SecurityTag":
                return new SecurityTag(name, <any>undefined, { urn })
            case "vcd:index/subscribedCatalog:SubscribedCatalog":
                return new SubscribedCatalog(name, <any>undefined, { urn })
            case "vcd:index/vapp:Vapp":
                return new Vapp(name, <any>undefined, { urn })
            case "vcd:index/vappAccessControl:VappAccessControl":
                return new VappAccessControl(name, <any>undefined, { urn })
            case "vcd:index/vappFirewallRules:VappFirewallRules":
                return new VappFirewallRules(name, <any>undefined, { urn })
            case "vcd:index/vappNatRules:VappNatRules":
                return new VappNatRules(name, <any>undefined, { urn })
            case "vcd:index/vappNetwork:VappNetwork":
                return new VappNetwork(name, <any>undefined, { urn })
            case "vcd:index/vappOrgNetwork:VappOrgNetwork":
                return new VappOrgNetwork(name, <any>undefined, { urn })
            case "vcd:index/vappStaticRouting:VappStaticRouting":
                return new VappStaticRouting(name, <any>undefined, { urn })
            case "vcd:index/vappVm:VappVm":
                return new VappVm(name, <any>undefined, { urn })
            case "vcd:index/vdcGroup:VdcGroup":
                return new VdcGroup(name, <any>undefined, { urn })
            case "vcd:index/vm:Vm":
                return new Vm(name, <any>undefined, { urn })
            case "vcd:index/vmAffinityRule:VmAffinityRule":
                return new VmAffinityRule(name, <any>undefined, { urn })
            case "vcd:index/vmInternalDisk:VmInternalDisk":
                return new VmInternalDisk(name, <any>undefined, { urn })
            case "vcd:index/vmPlacementPolicy:VmPlacementPolicy":
                return new VmPlacementPolicy(name, <any>undefined, { urn })
            case "vcd:index/vmSizingPolicy:VmSizingPolicy":
                return new VmSizingPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vcd", "index/catalog", _module)
pulumi.runtime.registerResourceModule("vcd", "index/catalogAccessControl", _module)
pulumi.runtime.registerResourceModule("vcd", "index/catalogItem", _module)
pulumi.runtime.registerResourceModule("vcd", "index/catalogMedia", _module)
pulumi.runtime.registerResourceModule("vcd", "index/catalogVappTemplate", _module)
pulumi.runtime.registerResourceModule("vcd", "index/edgeGateway", _module)
pulumi.runtime.registerResourceModule("vcd", "index/edgeGatewaySettings", _module)
pulumi.runtime.registerResourceModule("vcd", "index/edgeGatewayVpn", _module)
pulumi.runtime.registerResourceModule("vcd", "index/externalNetwork", _module)
pulumi.runtime.registerResourceModule("vcd", "index/externalNetworkV2", _module)
pulumi.runtime.registerResourceModule("vcd", "index/globalRole", _module)
pulumi.runtime.registerResourceModule("vcd", "index/independentDisk", _module)
pulumi.runtime.registerResourceModule("vcd", "index/insertedMedia", _module)
pulumi.runtime.registerResourceModule("vcd", "index/lbAppProfile", _module)
pulumi.runtime.registerResourceModule("vcd", "index/lbAppRule", _module)
pulumi.runtime.registerResourceModule("vcd", "index/lbServerPool", _module)
pulumi.runtime.registerResourceModule("vcd", "index/lbServiceMonitor", _module)
pulumi.runtime.registerResourceModule("vcd", "index/lbVirtualServer", _module)
pulumi.runtime.registerResourceModule("vcd", "index/libraryCertificate", _module)
pulumi.runtime.registerResourceModule("vcd", "index/networkDirect", _module)
pulumi.runtime.registerResourceModule("vcd", "index/networkIsolated", _module)
pulumi.runtime.registerResourceModule("vcd", "index/networkIsolatedV2", _module)
pulumi.runtime.registerResourceModule("vcd", "index/networkRouted", _module)
pulumi.runtime.registerResourceModule("vcd", "index/networkRoutedV2", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbCloud", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbController", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbEdgegatewayServiceEngineGroup", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbPool", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbServiceEngineGroup", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbSettings", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAlbVirtualService", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtAppPortProfile", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtDistributedFirewall", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtDynamicSecurityGroup", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtEdgegateway", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtEdgegatewayBgpConfiguration", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtEdgegatewayBgpIpPrefixList", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtEdgegatewayBgpNeighbor", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtFirewall", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtIpSet", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtIpsecVpnTunnel", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtNatRule", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtNetworkDhcp", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtNetworkImported", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtRouteAdvertisement", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxtSecurityGroup", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxvDhcpRelay", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxvDnat", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxvFirewallRule", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxvIpSet", _module)
pulumi.runtime.registerResourceModule("vcd", "index/nsxvSnat", _module)
pulumi.runtime.registerResourceModule("vcd", "index/org", _module)
pulumi.runtime.registerResourceModule("vcd", "index/orgGroup", _module)
pulumi.runtime.registerResourceModule("vcd", "index/orgLdap", _module)
pulumi.runtime.registerResourceModule("vcd", "index/orgUser", _module)
pulumi.runtime.registerResourceModule("vcd", "index/orgVdc", _module)
pulumi.runtime.registerResourceModule("vcd", "index/orgVdcAccessControl", _module)
pulumi.runtime.registerResourceModule("vcd", "index/rightsBundle", _module)
pulumi.runtime.registerResourceModule("vcd", "index/role", _module)
pulumi.runtime.registerResourceModule("vcd", "index/securityTag", _module)
pulumi.runtime.registerResourceModule("vcd", "index/subscribedCatalog", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vapp", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappAccessControl", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappFirewallRules", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappNatRules", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappNetwork", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappOrgNetwork", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappStaticRouting", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vappVm", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vdcGroup", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vm", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vmAffinityRule", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vmInternalDisk", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vmPlacementPolicy", _module)
pulumi.runtime.registerResourceModule("vcd", "index/vmSizingPolicy", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("vcd", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:vcd") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
