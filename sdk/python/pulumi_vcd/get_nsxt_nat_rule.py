# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetNsxtNatRuleResult',
    'AwaitableGetNsxtNatRuleResult',
    'get_nsxt_nat_rule',
    'get_nsxt_nat_rule_output',
]

@pulumi.output_type
class GetNsxtNatRuleResult:
    """
    A collection of values returned by getNsxtNatRule.
    """
    def __init__(__self__, app_port_profile_id=None, description=None, dnat_external_port=None, edge_gateway_id=None, enabled=None, external_address=None, firewall_match=None, id=None, internal_address=None, logging=None, name=None, org=None, priority=None, rule_type=None, snat_destination_address=None, vdc=None):
        if app_port_profile_id and not isinstance(app_port_profile_id, str):
            raise TypeError("Expected argument 'app_port_profile_id' to be a str")
        pulumi.set(__self__, "app_port_profile_id", app_port_profile_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dnat_external_port and not isinstance(dnat_external_port, str):
            raise TypeError("Expected argument 'dnat_external_port' to be a str")
        pulumi.set(__self__, "dnat_external_port", dnat_external_port)
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if external_address and not isinstance(external_address, str):
            raise TypeError("Expected argument 'external_address' to be a str")
        pulumi.set(__self__, "external_address", external_address)
        if firewall_match and not isinstance(firewall_match, str):
            raise TypeError("Expected argument 'firewall_match' to be a str")
        pulumi.set(__self__, "firewall_match", firewall_match)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_address and not isinstance(internal_address, str):
            raise TypeError("Expected argument 'internal_address' to be a str")
        pulumi.set(__self__, "internal_address", internal_address)
        if logging and not isinstance(logging, bool):
            raise TypeError("Expected argument 'logging' to be a bool")
        pulumi.set(__self__, "logging", logging)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if rule_type and not isinstance(rule_type, str):
            raise TypeError("Expected argument 'rule_type' to be a str")
        pulumi.set(__self__, "rule_type", rule_type)
        if snat_destination_address and not isinstance(snat_destination_address, str):
            raise TypeError("Expected argument 'snat_destination_address' to be a str")
        pulumi.set(__self__, "snat_destination_address", snat_destination_address)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        if vdc is not None:
            warnings.warn("""Edge Gateway will be looked up based on 'edge_gateway_id' field""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field""")

        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="appPortProfileId")
    def app_port_profile_id(self) -> str:
        return pulumi.get(self, "app_port_profile_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnatExternalPort")
    def dnat_external_port(self) -> str:
        return pulumi.get(self, "dnat_external_port")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> str:
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="externalAddress")
    def external_address(self) -> str:
        return pulumi.get(self, "external_address")

    @property
    @pulumi.getter(name="firewallMatch")
    def firewall_match(self) -> str:
        return pulumi.get(self, "firewall_match")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalAddress")
    def internal_address(self) -> str:
        return pulumi.get(self, "internal_address")

    @property
    @pulumi.getter
    def logging(self) -> bool:
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def priority(self) -> int:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> str:
        return pulumi.get(self, "rule_type")

    @property
    @pulumi.getter(name="snatDestinationAddress")
    def snat_destination_address(self) -> str:
        return pulumi.get(self, "snat_destination_address")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtNatRuleResult(GetNsxtNatRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtNatRuleResult(
            app_port_profile_id=self.app_port_profile_id,
            description=self.description,
            dnat_external_port=self.dnat_external_port,
            edge_gateway_id=self.edge_gateway_id,
            enabled=self.enabled,
            external_address=self.external_address,
            firewall_match=self.firewall_match,
            id=self.id,
            internal_address=self.internal_address,
            logging=self.logging,
            name=self.name,
            org=self.org,
            priority=self.priority,
            rule_type=self.rule_type,
            snat_destination_address=self.snat_destination_address,
            vdc=self.vdc)


def get_nsxt_nat_rule(edge_gateway_id: Optional[str] = None,
                      name: Optional[str] = None,
                      org: Optional[str] = None,
                      vdc: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtNatRuleResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtNatRule:getNsxtNatRule', __args__, opts=opts, typ=GetNsxtNatRuleResult).value

    return AwaitableGetNsxtNatRuleResult(
        app_port_profile_id=__ret__.app_port_profile_id,
        description=__ret__.description,
        dnat_external_port=__ret__.dnat_external_port,
        edge_gateway_id=__ret__.edge_gateway_id,
        enabled=__ret__.enabled,
        external_address=__ret__.external_address,
        firewall_match=__ret__.firewall_match,
        id=__ret__.id,
        internal_address=__ret__.internal_address,
        logging=__ret__.logging,
        name=__ret__.name,
        org=__ret__.org,
        priority=__ret__.priority,
        rule_type=__ret__.rule_type,
        snat_destination_address=__ret__.snat_destination_address,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_nsxt_nat_rule)
def get_nsxt_nat_rule_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                             name: Optional[pulumi.Input[str]] = None,
                             org: Optional[pulumi.Input[Optional[str]]] = None,
                             vdc: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNsxtNatRuleResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...