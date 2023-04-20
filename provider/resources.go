// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vcd

import (
	"fmt"
	"path/filepath"

	"github.com/kitzin/pulumi-vcd/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/vmware/terraform-provider-vcd/v3/vcd"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vcd"
	// modules:
	mainMod = "index" // the vcd module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(vcd.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "vcd",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "kitzin",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing vmware cloud directory resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "vcd", "vmware", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/kitzin/pulumi-vcd",
		Repository: "https://github.com/kitzin/pulumi-vcd",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "kitzin",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"vcd_catalog":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Catalog")},
			"vcd_catalog_access_control": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogAccessControl")},
			"vcd_catalog_item":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogItem")},
			"vcd_catalog_media":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogMedia")},
			"vcd_catalog_vapp_template":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogVappTemplate")},
			"vcd_edgegateway":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EdgeGateway")},
			"vcd_edgegateway_settings":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EdgeGatewaySettings")},
			"vcd_edgegateway_vpn":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EdgeGatewayVpn")},
			"vcd_external_network":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ExternalNetwork")},
			"vcd_external_network_v2":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ExternalNetworkV2")},
			"vcd_global_role":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "GlobalRole")},
			"vcd_independent_disk":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IndependentDisk")},
			"vcd_inserted_media":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InsertedMedia")},
			"vcd_lb_app_profile":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbAppProfile")},
			"vcd_lb_app_rule":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbAppRule")},
			"vcd_lb_server_pool":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbServerPool")},
			"vcd_lb_service_monitor":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbServiceMonitor")},
			"vcd_lb_virtual_server":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbVirtualServer")},
			"vcd_library_certificate":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LibraryCertificate")},
			"vcd_network_direct":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkDirect")},
			"vcd_network_isolated":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkIsolated")},
			"vcd_network_isolated_v2":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkIsolatedV2")},
			"vcd_network_routed":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkRouted")},
			"vcd_network_routed_v2":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkRoutedV2")},
			"vcd_nsxt_alb_cloud":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbCloud")},
			"vcd_nsxt_alb_controller":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbController")},
			"vcd_nsxt_alb_edgegateway_service_engine_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbEdgegatewayServiceEngineGroup"),
			},
			"vcd_nsxt_alb_pool": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbPool")},
			"vcd_nsxt_alb_service_engine_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbServiceEngineGroup"),
			},
			"vcd_nsxt_alb_settings":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbSettings")},
			"vcd_nsxt_alb_virtual_service":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAlbVirtualService")},
			"vcd_nsxt_app_port_profile":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtAppPortProfile")},
			"vcd_nsxt_distributed_firewall": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtDistributedFirewall")},
			"vcd_nsxt_dynamic_security_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtDynamicSecurityGroup"),
			},
			"vcd_nsxt_edgegateway": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtEdgegateway")},
			"vcd_nsxt_edgegateway_bgp_configuration": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtEdgegatewayBgpConfiguration"),
			},
			"vcd_nsxt_edgegateway_bgp_ip_prefix_list": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtEdgegatewayBgpIpPrefixList"),
			},
			"vcd_nsxt_edgegateway_bgp_neighbor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtEdgegatewayBgpNeighbor"),
			},
			"vcd_nsxt_firewall":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtFirewall")},
			"vcd_nsxt_ip_set":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtIpSet")},
			"vcd_nsxt_ipsec_vpn_tunnel":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtIpsecVpnTunnel")},
			"vcd_nsxt_nat_rule":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtNatRule")},
			"vcd_nsxt_network_dhcp":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtNetworkDhcp")},
			"vcd_nsxt_network_imported":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtNetworkImported")},
			"vcd_nsxt_route_advertisement": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtRouteAdvertisement")},
			"vcd_nsxt_security_group":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxtSecurityGroup")},
			"vcd_nsxv_dhcp_relay":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxvDhcpRelay")},
			"vcd_nsxv_dnat":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxvDnat")},
			"vcd_nsxv_firewall_rule":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxvFirewallRule")},
			"vcd_nsxv_ip_set":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxvIpSet")},
			"vcd_nsxv_snat":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NsxvSnat")},
			"vcd_org":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Org")},
			"vcd_org_group":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgGroup")},
			"vcd_org_ldap":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgLdap")},
			"vcd_org_user":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgUser")},
			"vcd_org_vdc":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgVdc")},
			"vcd_org_vdc_access_control":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "OrgVdcAccessControl")},
			"vcd_rights_bundle":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RightsBundle")},
			"vcd_role":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Role")},
			"vcd_security_tag":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecurityTag")},
			"vcd_subscribed_catalog":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SubscribedCatalog")},
			"vcd_vapp":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vapp")},
			"vcd_vapp_access_control":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappAccessControl")},
			"vcd_vapp_firewall_rules":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappFirewallRules")},
			"vcd_vapp_nat_rules":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappNatRules")},
			"vcd_vapp_network":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappNetwork")},
			"vcd_vapp_org_network":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappOrgNetwork")},
			"vcd_vapp_static_routing":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappStaticRouting")},
			"vcd_vapp_vm":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VappVm")},
			"vcd_vdc_group":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VdcGroup")},
			"vcd_vm":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vm")},
			"vcd_vm_affinity_rule":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VmAffinityRule")},
			"vcd_vm_internal_disk":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VmInternalDisk")},
			"vcd_vm_placement_policy":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VmPlacementPolicy")},
			"vcd_vm_sizing_policy":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VmSizingPolicy")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vcd_catalog":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalog")},
			"vcd_catalog_item":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogItem")},
			"vcd_catalog_media":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogMedia")},
			"vcd_catalog_vapp_template": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogVappTemplate")},
			"vcd_edgegateway":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEdgegateway")},
			"vcd_external_network":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getExternalNetwork")},
			"vcd_external_network_v2":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getExternalNetworkV2")},
			"vcd_global_role":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGlobalRole")},
			"vcd_independent_disk":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIndependentDisk")},
			"vcd_lb_app_profile":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbAppProfile")},
			"vcd_lb_app_rule":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbAppRule")},
			"vcd_lb_server_pool":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbServerPool")},
			"vcd_lb_service_monitor":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbServiceMonitor")},
			"vcd_lb_virtual_server":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbVirtualServer")},
			"vcd_library_certificate":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLibraryCertificate")},
			"vcd_network_direct":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkDirect")},
			"vcd_network_isolated":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkIsolated")},
			"vcd_network_isolated_v2":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkIsolatedV2")},
			"vcd_network_routed":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkRouted")},
			"vcd_network_routed_v2":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkRoutedV2")},
			"vcd_nsxt_alb_cloud":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbCloud")},
			"vcd_nsxt_alb_controller":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbController")},
			"vcd_nsxt_alb_edgegateway_service_engine_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbEdgegatewayServiceEngineGroup"),
			},
			"vcd_nsxt_alb_importable_cloud": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbImportableCloud"),
			},
			"vcd_nsxt_alb_pool": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbPool")},
			"vcd_nsxt_alb_service_engine_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbServiceEngineGroup"),
			},
			"vcd_nsxt_alb_settings": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbSettings")},
			"vcd_nsxt_alb_virtual_service": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAlbVirtualService"),
			},
			"vcd_nsxt_app_port_profile": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtAppPortProfile")},
			"vcd_nsxt_distributed_firewall": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtDistributedFirewall"),
			},
			"vcd_nsxt_dynamic_security_group": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtDynamicSecurityGroup"),
			},
			"vcd_nsxt_edge_cluster": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtEdgeCluster"),
			},
			"vcd_nsxt_edgegateway": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtEdgegateway")},
			"vcd_nsxt_edgegateway_bgp_configuration": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtEdgegatewayBgpConfiguration"),
			},
			"vcd_nsxt_edgegateway_bgp_ip_prefix_list": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtEdgegatewayBgpIpPrefixList"),
			},
			"vcd_nsxt_edgegateway_bgp_neighbor": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtEdgegatewayBgpNeighbor"),
			},
			"vcd_nsxt_firewall":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtFirewall")},
			"vcd_nsxt_ip_set":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtIpSet")},
			"vcd_nsxt_ipsec_vpn_tunnel": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtIpsecVpnTunnel")},
			"vcd_nsxt_manager":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtManager")},
			"vcd_nsxt_nat_rule":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtNatRule")},
			"vcd_nsxt_network_context_profile": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtNetworkContextProfile"),
			},
			"vcd_nsxt_network_dhcp":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtNetworkDhcp")},
			"vcd_nsxt_network_imported": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtNetworkImported")},
			"vcd_nsxt_route_advertisement": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtRouteAdvertisement"),
			},
			"vcd_nsxt_security_group": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtSecurityGroup")},
			"vcd_nsxt_tier0_router":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxtTier0Router")},
			"vcd_nsxv_dhcp_relay":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxvDhcpRelay")},
			"vcd_nsxv_dnat":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxvDnat")},
			"vcd_nsxv_firewall_rule":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxvFirewallRule")},
			"vcd_nsxv_ip_set":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxvIpSet")},
			"vcd_nsxv_snat":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNsxvSnat")},
			"vcd_org":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrg")},
			"vcd_org_group":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgGroup")},
			"vcd_org_ldap":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgLdap")},
			"vcd_org_user":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgUser")},
			"vcd_org_vdc":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOrgVdc")},
			"vcd_portgroup":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPortgroup")},
			"vcd_provider_vdc":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProviderVdc")},
			"vcd_resource_list":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getResourceList")},
			"vcd_resource_schema":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getResourceSchema")},
			"vcd_right":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRight")},
			"vcd_rights_bundle":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRightsBundle")},
			"vcd_role":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRole")},
			"vcd_storage_profile":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageProfile")},
			"vcd_subscribed_catalog":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSubscribedCatalog")},
			"vcd_task":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTask")},
			"vcd_vapp":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVapp")},
			"vcd_vapp_network":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVappNetwork")},
			"vcd_vapp_org_network":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVappOrgNetwork")},
			"vcd_vapp_vm":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVappVm")},
			"vcd_vcenter":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVcenter")},
			"vcd_vdc_group":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVdcGroup")},
			"vcd_vm":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVm")},
			"vcd_vm_affinity_rule":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVmAffinityRule")},
			"vcd_vm_group":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVmGroup")},
			"vcd_vm_placement_policy": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVmPlacementPolicy")},
			"vcd_vm_sizing_policy":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVmSizingPolicy")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
