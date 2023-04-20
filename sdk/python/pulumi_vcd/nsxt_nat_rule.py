# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NsxtNatRuleArgs', 'NsxtNatRule']

@pulumi.input_type
class NsxtNatRuleArgs:
    def __init__(__self__, *,
                 edge_gateway_id: pulumi.Input[str],
                 rule_type: pulumi.Input[str],
                 app_port_profile_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dnat_external_port: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 external_address: Optional[pulumi.Input[str]] = None,
                 firewall_match: Optional[pulumi.Input[str]] = None,
                 internal_address: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 snat_destination_address: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NsxtNatRule resource.
        :param pulumi.Input[str] edge_gateway_id: Edge gateway name in which NAT Rule is located
        :param pulumi.Input[str] rule_type: Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        :param pulumi.Input[str] app_port_profile_id: Application Port Profile to apply for this rule
        :param pulumi.Input[str] description: Description of NAT rule
        :param pulumi.Input[str] dnat_external_port: For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        :param pulumi.Input[bool] enabled: Enables or disables this rule
        :param pulumi.Input[str] external_address: IP address or CIDR of external network
        :param pulumi.Input[str] firewall_match: VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
               'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        :param pulumi.Input[str] internal_address: IP address or CIDR of the virtual machines for which you are configuring NAT
        :param pulumi.Input[bool] logging: Enable logging when this rule is applied
        :param pulumi.Input[str] name: Name of NAT rule
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[int] priority: VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
               higher precedence for this rule.
        :param pulumi.Input[str] snat_destination_address: For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
               or an IP address range in CIDR format.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        pulumi.set(__self__, "rule_type", rule_type)
        if app_port_profile_id is not None:
            pulumi.set(__self__, "app_port_profile_id", app_port_profile_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dnat_external_port is not None:
            pulumi.set(__self__, "dnat_external_port", dnat_external_port)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if external_address is not None:
            pulumi.set(__self__, "external_address", external_address)
        if firewall_match is not None:
            pulumi.set(__self__, "firewall_match", firewall_match)
        if internal_address is not None:
            pulumi.set(__self__, "internal_address", internal_address)
        if logging is not None:
            pulumi.set(__self__, "logging", logging)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if snat_destination_address is not None:
            pulumi.set(__self__, "snat_destination_address", snat_destination_address)
        if vdc is not None:
            warnings.warn("""Edge Gateway will be looked up based on 'edge_gateway_id' field""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field""")
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Input[str]:
        """
        Edge gateway name in which NAT Rule is located
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[str]:
        """
        Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_type", value)

    @property
    @pulumi.getter(name="appPortProfileId")
    def app_port_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application Port Profile to apply for this rule
        """
        return pulumi.get(self, "app_port_profile_id")

    @app_port_profile_id.setter
    def app_port_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_port_profile_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of NAT rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnatExternalPort")
    def dnat_external_port(self) -> Optional[pulumi.Input[str]]:
        """
        For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        """
        return pulumi.get(self, "dnat_external_port")

    @dnat_external_port.setter
    def dnat_external_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dnat_external_port", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables or disables this rule
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="externalAddress")
    def external_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or CIDR of external network
        """
        return pulumi.get(self, "external_address")

    @external_address.setter
    def external_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_address", value)

    @property
    @pulumi.getter(name="firewallMatch")
    def firewall_match(self) -> Optional[pulumi.Input[str]]:
        """
        VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
        'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        """
        return pulumi.get(self, "firewall_match")

    @firewall_match.setter
    def firewall_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_match", value)

    @property
    @pulumi.getter(name="internalAddress")
    def internal_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or CIDR of the virtual machines for which you are configuring NAT
        """
        return pulumi.get(self, "internal_address")

    @internal_address.setter
    def internal_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_address", value)

    @property
    @pulumi.getter
    def logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable logging when this rule is applied
        """
        return pulumi.get(self, "logging")

    @logging.setter
    def logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "logging", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of NAT rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
        higher precedence for this rule.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="snatDestinationAddress")
    def snat_destination_address(self) -> Optional[pulumi.Input[str]]:
        """
        For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
        or an IP address range in CIDR format.
        """
        return pulumi.get(self, "snat_destination_address")

    @snat_destination_address.setter
    def snat_destination_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_destination_address", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


@pulumi.input_type
class _NsxtNatRuleState:
    def __init__(__self__, *,
                 app_port_profile_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dnat_external_port: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 external_address: Optional[pulumi.Input[str]] = None,
                 firewall_match: Optional[pulumi.Input[str]] = None,
                 internal_address: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 rule_type: Optional[pulumi.Input[str]] = None,
                 snat_destination_address: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NsxtNatRule resources.
        :param pulumi.Input[str] app_port_profile_id: Application Port Profile to apply for this rule
        :param pulumi.Input[str] description: Description of NAT rule
        :param pulumi.Input[str] dnat_external_port: For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        :param pulumi.Input[str] edge_gateway_id: Edge gateway name in which NAT Rule is located
        :param pulumi.Input[bool] enabled: Enables or disables this rule
        :param pulumi.Input[str] external_address: IP address or CIDR of external network
        :param pulumi.Input[str] firewall_match: VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
               'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        :param pulumi.Input[str] internal_address: IP address or CIDR of the virtual machines for which you are configuring NAT
        :param pulumi.Input[bool] logging: Enable logging when this rule is applied
        :param pulumi.Input[str] name: Name of NAT rule
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[int] priority: VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
               higher precedence for this rule.
        :param pulumi.Input[str] rule_type: Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        :param pulumi.Input[str] snat_destination_address: For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
               or an IP address range in CIDR format.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if app_port_profile_id is not None:
            pulumi.set(__self__, "app_port_profile_id", app_port_profile_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dnat_external_port is not None:
            pulumi.set(__self__, "dnat_external_port", dnat_external_port)
        if edge_gateway_id is not None:
            pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if external_address is not None:
            pulumi.set(__self__, "external_address", external_address)
        if firewall_match is not None:
            pulumi.set(__self__, "firewall_match", firewall_match)
        if internal_address is not None:
            pulumi.set(__self__, "internal_address", internal_address)
        if logging is not None:
            pulumi.set(__self__, "logging", logging)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if rule_type is not None:
            pulumi.set(__self__, "rule_type", rule_type)
        if snat_destination_address is not None:
            pulumi.set(__self__, "snat_destination_address", snat_destination_address)
        if vdc is not None:
            warnings.warn("""Edge Gateway will be looked up based on 'edge_gateway_id' field""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field""")
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="appPortProfileId")
    def app_port_profile_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application Port Profile to apply for this rule
        """
        return pulumi.get(self, "app_port_profile_id")

    @app_port_profile_id.setter
    def app_port_profile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_port_profile_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of NAT rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnatExternalPort")
    def dnat_external_port(self) -> Optional[pulumi.Input[str]]:
        """
        For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        """
        return pulumi.get(self, "dnat_external_port")

    @dnat_external_port.setter
    def dnat_external_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dnat_external_port", value)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Edge gateway name in which NAT Rule is located
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables or disables this rule
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="externalAddress")
    def external_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or CIDR of external network
        """
        return pulumi.get(self, "external_address")

    @external_address.setter
    def external_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_address", value)

    @property
    @pulumi.getter(name="firewallMatch")
    def firewall_match(self) -> Optional[pulumi.Input[str]]:
        """
        VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
        'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        """
        return pulumi.get(self, "firewall_match")

    @firewall_match.setter
    def firewall_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_match", value)

    @property
    @pulumi.getter(name="internalAddress")
    def internal_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or CIDR of the virtual machines for which you are configuring NAT
        """
        return pulumi.get(self, "internal_address")

    @internal_address.setter
    def internal_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_address", value)

    @property
    @pulumi.getter
    def logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable logging when this rule is applied
        """
        return pulumi.get(self, "logging")

    @logging.setter
    def logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "logging", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of NAT rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
        higher precedence for this rule.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> Optional[pulumi.Input[str]]:
        """
        Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_type", value)

    @property
    @pulumi.getter(name="snatDestinationAddress")
    def snat_destination_address(self) -> Optional[pulumi.Input[str]]:
        """
        For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
        or an IP address range in CIDR format.
        """
        return pulumi.get(self, "snat_destination_address")

    @snat_destination_address.setter
    def snat_destination_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_destination_address", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


class NsxtNatRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_port_profile_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dnat_external_port: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 external_address: Optional[pulumi.Input[str]] = None,
                 firewall_match: Optional[pulumi.Input[str]] = None,
                 internal_address: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 rule_type: Optional[pulumi.Input[str]] = None,
                 snat_destination_address: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NsxtNatRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_port_profile_id: Application Port Profile to apply for this rule
        :param pulumi.Input[str] description: Description of NAT rule
        :param pulumi.Input[str] dnat_external_port: For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        :param pulumi.Input[str] edge_gateway_id: Edge gateway name in which NAT Rule is located
        :param pulumi.Input[bool] enabled: Enables or disables this rule
        :param pulumi.Input[str] external_address: IP address or CIDR of external network
        :param pulumi.Input[str] firewall_match: VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
               'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        :param pulumi.Input[str] internal_address: IP address or CIDR of the virtual machines for which you are configuring NAT
        :param pulumi.Input[bool] logging: Enable logging when this rule is applied
        :param pulumi.Input[str] name: Name of NAT rule
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[int] priority: VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
               higher precedence for this rule.
        :param pulumi.Input[str] rule_type: Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        :param pulumi.Input[str] snat_destination_address: For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
               or an IP address range in CIDR format.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NsxtNatRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NsxtNatRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NsxtNatRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxtNatRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_port_profile_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dnat_external_port: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 external_address: Optional[pulumi.Input[str]] = None,
                 firewall_match: Optional[pulumi.Input[str]] = None,
                 internal_address: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 rule_type: Optional[pulumi.Input[str]] = None,
                 snat_destination_address: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NsxtNatRuleArgs.__new__(NsxtNatRuleArgs)

            __props__.__dict__["app_port_profile_id"] = app_port_profile_id
            __props__.__dict__["description"] = description
            __props__.__dict__["dnat_external_port"] = dnat_external_port
            if edge_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'edge_gateway_id'")
            __props__.__dict__["edge_gateway_id"] = edge_gateway_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["external_address"] = external_address
            __props__.__dict__["firewall_match"] = firewall_match
            __props__.__dict__["internal_address"] = internal_address
            __props__.__dict__["logging"] = logging
            __props__.__dict__["name"] = name
            __props__.__dict__["org"] = org
            __props__.__dict__["priority"] = priority
            if rule_type is None and not opts.urn:
                raise TypeError("Missing required property 'rule_type'")
            __props__.__dict__["rule_type"] = rule_type
            __props__.__dict__["snat_destination_address"] = snat_destination_address
            if vdc is not None and not opts.urn:
                warnings.warn("""Edge Gateway will be looked up based on 'edge_gateway_id' field""", DeprecationWarning)
                pulumi.log.warn("""vdc is deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field""")
            __props__.__dict__["vdc"] = vdc
        super(NsxtNatRule, __self__).__init__(
            'vcd:index/nsxtNatRule:NsxtNatRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_port_profile_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dnat_external_port: Optional[pulumi.Input[str]] = None,
            edge_gateway_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            external_address: Optional[pulumi.Input[str]] = None,
            firewall_match: Optional[pulumi.Input[str]] = None,
            internal_address: Optional[pulumi.Input[str]] = None,
            logging: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            rule_type: Optional[pulumi.Input[str]] = None,
            snat_destination_address: Optional[pulumi.Input[str]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'NsxtNatRule':
        """
        Get an existing NsxtNatRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_port_profile_id: Application Port Profile to apply for this rule
        :param pulumi.Input[str] description: Description of NAT rule
        :param pulumi.Input[str] dnat_external_port: For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        :param pulumi.Input[str] edge_gateway_id: Edge gateway name in which NAT Rule is located
        :param pulumi.Input[bool] enabled: Enables or disables this rule
        :param pulumi.Input[str] external_address: IP address or CIDR of external network
        :param pulumi.Input[str] firewall_match: VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
               'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        :param pulumi.Input[str] internal_address: IP address or CIDR of the virtual machines for which you are configuring NAT
        :param pulumi.Input[bool] logging: Enable logging when this rule is applied
        :param pulumi.Input[str] name: Name of NAT rule
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[int] priority: VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
               higher precedence for this rule.
        :param pulumi.Input[str] rule_type: Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        :param pulumi.Input[str] snat_destination_address: For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
               or an IP address range in CIDR format.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxtNatRuleState.__new__(_NsxtNatRuleState)

        __props__.__dict__["app_port_profile_id"] = app_port_profile_id
        __props__.__dict__["description"] = description
        __props__.__dict__["dnat_external_port"] = dnat_external_port
        __props__.__dict__["edge_gateway_id"] = edge_gateway_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["external_address"] = external_address
        __props__.__dict__["firewall_match"] = firewall_match
        __props__.__dict__["internal_address"] = internal_address
        __props__.__dict__["logging"] = logging
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        __props__.__dict__["priority"] = priority
        __props__.__dict__["rule_type"] = rule_type
        __props__.__dict__["snat_destination_address"] = snat_destination_address
        __props__.__dict__["vdc"] = vdc
        return NsxtNatRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appPortProfileId")
    def app_port_profile_id(self) -> pulumi.Output[Optional[str]]:
        """
        Application Port Profile to apply for this rule
        """
        return pulumi.get(self, "app_port_profile_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of NAT rule
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnatExternalPort")
    def dnat_external_port(self) -> pulumi.Output[Optional[str]]:
        """
        For DNAT only. Enter a port into which the DNAT rule is translating for the packets inbound to the virtual machines.
        """
        return pulumi.get(self, "dnat_external_port")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Output[str]:
        """
        Edge gateway name in which NAT Rule is located
        """
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables or disables this rule
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="externalAddress")
    def external_address(self) -> pulumi.Output[Optional[str]]:
        """
        IP address or CIDR of external network
        """
        return pulumi.get(self, "external_address")

    @property
    @pulumi.getter(name="firewallMatch")
    def firewall_match(self) -> pulumi.Output[str]:
        """
        VCD 10.2.2+ Determines how the firewall matches the address during NATing if firewall stage is not skipped. One of
        'MATCH_INTERNAL_ADDRESS', 'MATCH_EXTERNAL_ADDRESS', 'BYPASS'
        """
        return pulumi.get(self, "firewall_match")

    @property
    @pulumi.getter(name="internalAddress")
    def internal_address(self) -> pulumi.Output[Optional[str]]:
        """
        IP address or CIDR of the virtual machines for which you are configuring NAT
        """
        return pulumi.get(self, "internal_address")

    @property
    @pulumi.getter
    def logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable logging when this rule is applied
        """
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of NAT rule
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        VCD 10.2.2+ If an address has multiple NAT rules, the rule with the highest priority is applied. A lower value means a
        higher precedence for this rule.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Output[str]:
        """
        Rule type - one of 'DNAT', 'NO_DNAT', 'SNAT', 'NO_SNAT', 'REFLEXIVE'
        """
        return pulumi.get(self, "rule_type")

    @property
    @pulumi.getter(name="snatDestinationAddress")
    def snat_destination_address(self) -> pulumi.Output[Optional[str]]:
        """
        For SNAT only. If you want the rule to apply only for traffic to a specific domain, enter an IP address for this domain
        or an IP address range in CIDR format.
        """
        return pulumi.get(self, "snat_destination_address")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[str]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")
