# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetNsxvFirewallRuleResult',
    'AwaitableGetNsxvFirewallRuleResult',
    'get_nsxv_firewall_rule',
    'get_nsxv_firewall_rule_output',
]

@pulumi.output_type
class GetNsxvFirewallRuleResult:
    """
    A collection of values returned by getNsxvFirewallRule.
    """
    def __init__(__self__, action=None, destinations=None, edge_gateway=None, enabled=None, id=None, logging_enabled=None, name=None, org=None, rule_id=None, rule_tag=None, rule_type=None, services=None, sources=None, vdc=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if destinations and not isinstance(destinations, list):
            raise TypeError("Expected argument 'destinations' to be a list")
        pulumi.set(__self__, "destinations", destinations)
        if edge_gateway and not isinstance(edge_gateway, str):
            raise TypeError("Expected argument 'edge_gateway' to be a str")
        pulumi.set(__self__, "edge_gateway", edge_gateway)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logging_enabled and not isinstance(logging_enabled, bool):
            raise TypeError("Expected argument 'logging_enabled' to be a bool")
        pulumi.set(__self__, "logging_enabled", logging_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if rule_id and not isinstance(rule_id, str):
            raise TypeError("Expected argument 'rule_id' to be a str")
        pulumi.set(__self__, "rule_id", rule_id)
        if rule_tag and not isinstance(rule_tag, int):
            raise TypeError("Expected argument 'rule_tag' to be a int")
        pulumi.set(__self__, "rule_tag", rule_tag)
        if rule_type and not isinstance(rule_type, str):
            raise TypeError("Expected argument 'rule_type' to be a str")
        pulumi.set(__self__, "rule_type", rule_type)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if sources and not isinstance(sources, list):
            raise TypeError("Expected argument 'sources' to be a list")
        pulumi.set(__self__, "sources", sources)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def destinations(self) -> Sequence['outputs.GetNsxvFirewallRuleDestinationResult']:
        return pulumi.get(self, "destinations")

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> str:
        return pulumi.get(self, "edge_gateway")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loggingEnabled")
    def logging_enabled(self) -> bool:
        return pulumi.get(self, "logging_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> str:
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="ruleTag")
    def rule_tag(self) -> int:
        return pulumi.get(self, "rule_tag")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> str:
        return pulumi.get(self, "rule_type")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetNsxvFirewallRuleServiceResult']:
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def sources(self) -> Sequence['outputs.GetNsxvFirewallRuleSourceResult']:
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxvFirewallRuleResult(GetNsxvFirewallRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxvFirewallRuleResult(
            action=self.action,
            destinations=self.destinations,
            edge_gateway=self.edge_gateway,
            enabled=self.enabled,
            id=self.id,
            logging_enabled=self.logging_enabled,
            name=self.name,
            org=self.org,
            rule_id=self.rule_id,
            rule_tag=self.rule_tag,
            rule_type=self.rule_type,
            services=self.services,
            sources=self.sources,
            vdc=self.vdc)


def get_nsxv_firewall_rule(edge_gateway: Optional[str] = None,
                           org: Optional[str] = None,
                           rule_id: Optional[str] = None,
                           vdc: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxvFirewallRuleResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['edgeGateway'] = edge_gateway
    __args__['org'] = org
    __args__['ruleId'] = rule_id
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxvFirewallRule:getNsxvFirewallRule', __args__, opts=opts, typ=GetNsxvFirewallRuleResult).value

    return AwaitableGetNsxvFirewallRuleResult(
        action=__ret__.action,
        destinations=__ret__.destinations,
        edge_gateway=__ret__.edge_gateway,
        enabled=__ret__.enabled,
        id=__ret__.id,
        logging_enabled=__ret__.logging_enabled,
        name=__ret__.name,
        org=__ret__.org,
        rule_id=__ret__.rule_id,
        rule_tag=__ret__.rule_tag,
        rule_type=__ret__.rule_type,
        services=__ret__.services,
        sources=__ret__.sources,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_nsxv_firewall_rule)
def get_nsxv_firewall_rule_output(edge_gateway: Optional[pulumi.Input[str]] = None,
                                  org: Optional[pulumi.Input[Optional[str]]] = None,
                                  rule_id: Optional[pulumi.Input[str]] = None,
                                  vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNsxvFirewallRuleResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...