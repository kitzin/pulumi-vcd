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
    'GetVdcGroupResult',
    'AwaitableGetVdcGroupResult',
    'get_vdc_group',
    'get_vdc_group_output',
]

@pulumi.output_type
class GetVdcGroupResult:
    """
    A collection of values returned by getVdcGroup.
    """
    def __init__(__self__, default_policy_status=None, description=None, dfw_enabled=None, error_message=None, id=None, local_egress=None, name=None, network_pool_id=None, network_pool_universal_id=None, network_provider_type=None, org=None, participating_org_vdcs=None, status=None, type=None, universal_networking_enabled=None):
        if default_policy_status and not isinstance(default_policy_status, bool):
            raise TypeError("Expected argument 'default_policy_status' to be a bool")
        pulumi.set(__self__, "default_policy_status", default_policy_status)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dfw_enabled and not isinstance(dfw_enabled, bool):
            raise TypeError("Expected argument 'dfw_enabled' to be a bool")
        pulumi.set(__self__, "dfw_enabled", dfw_enabled)
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if local_egress and not isinstance(local_egress, bool):
            raise TypeError("Expected argument 'local_egress' to be a bool")
        pulumi.set(__self__, "local_egress", local_egress)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_pool_id and not isinstance(network_pool_id, str):
            raise TypeError("Expected argument 'network_pool_id' to be a str")
        pulumi.set(__self__, "network_pool_id", network_pool_id)
        if network_pool_universal_id and not isinstance(network_pool_universal_id, str):
            raise TypeError("Expected argument 'network_pool_universal_id' to be a str")
        pulumi.set(__self__, "network_pool_universal_id", network_pool_universal_id)
        if network_provider_type and not isinstance(network_provider_type, str):
            raise TypeError("Expected argument 'network_provider_type' to be a str")
        pulumi.set(__self__, "network_provider_type", network_provider_type)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if participating_org_vdcs and not isinstance(participating_org_vdcs, list):
            raise TypeError("Expected argument 'participating_org_vdcs' to be a list")
        pulumi.set(__self__, "participating_org_vdcs", participating_org_vdcs)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if universal_networking_enabled and not isinstance(universal_networking_enabled, bool):
            raise TypeError("Expected argument 'universal_networking_enabled' to be a bool")
        pulumi.set(__self__, "universal_networking_enabled", universal_networking_enabled)

    @property
    @pulumi.getter(name="defaultPolicyStatus")
    def default_policy_status(self) -> bool:
        return pulumi.get(self, "default_policy_status")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dfwEnabled")
    def dfw_enabled(self) -> bool:
        return pulumi.get(self, "dfw_enabled")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> str:
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="localEgress")
    def local_egress(self) -> bool:
        return pulumi.get(self, "local_egress")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkPoolId")
    def network_pool_id(self) -> str:
        return pulumi.get(self, "network_pool_id")

    @property
    @pulumi.getter(name="networkPoolUniversalId")
    def network_pool_universal_id(self) -> str:
        return pulumi.get(self, "network_pool_universal_id")

    @property
    @pulumi.getter(name="networkProviderType")
    def network_provider_type(self) -> str:
        return pulumi.get(self, "network_provider_type")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="participatingOrgVdcs")
    def participating_org_vdcs(self) -> Sequence['outputs.GetVdcGroupParticipatingOrgVdcResult']:
        return pulumi.get(self, "participating_org_vdcs")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="universalNetworkingEnabled")
    def universal_networking_enabled(self) -> bool:
        return pulumi.get(self, "universal_networking_enabled")


class AwaitableGetVdcGroupResult(GetVdcGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVdcGroupResult(
            default_policy_status=self.default_policy_status,
            description=self.description,
            dfw_enabled=self.dfw_enabled,
            error_message=self.error_message,
            id=self.id,
            local_egress=self.local_egress,
            name=self.name,
            network_pool_id=self.network_pool_id,
            network_pool_universal_id=self.network_pool_universal_id,
            network_provider_type=self.network_provider_type,
            org=self.org,
            participating_org_vdcs=self.participating_org_vdcs,
            status=self.status,
            type=self.type,
            universal_networking_enabled=self.universal_networking_enabled)


def get_vdc_group(default_policy_status: Optional[bool] = None,
                  description: Optional[str] = None,
                  error_message: Optional[str] = None,
                  id: Optional[str] = None,
                  local_egress: Optional[bool] = None,
                  name: Optional[str] = None,
                  network_pool_id: Optional[str] = None,
                  network_pool_universal_id: Optional[str] = None,
                  network_provider_type: Optional[str] = None,
                  org: Optional[str] = None,
                  status: Optional[str] = None,
                  type: Optional[str] = None,
                  universal_networking_enabled: Optional[bool] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVdcGroupResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['defaultPolicyStatus'] = default_policy_status
    __args__['description'] = description
    __args__['errorMessage'] = error_message
    __args__['id'] = id
    __args__['localEgress'] = local_egress
    __args__['name'] = name
    __args__['networkPoolId'] = network_pool_id
    __args__['networkPoolUniversalId'] = network_pool_universal_id
    __args__['networkProviderType'] = network_provider_type
    __args__['org'] = org
    __args__['status'] = status
    __args__['type'] = type
    __args__['universalNetworkingEnabled'] = universal_networking_enabled
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getVdcGroup:getVdcGroup', __args__, opts=opts, typ=GetVdcGroupResult).value

    return AwaitableGetVdcGroupResult(
        default_policy_status=__ret__.default_policy_status,
        description=__ret__.description,
        dfw_enabled=__ret__.dfw_enabled,
        error_message=__ret__.error_message,
        id=__ret__.id,
        local_egress=__ret__.local_egress,
        name=__ret__.name,
        network_pool_id=__ret__.network_pool_id,
        network_pool_universal_id=__ret__.network_pool_universal_id,
        network_provider_type=__ret__.network_provider_type,
        org=__ret__.org,
        participating_org_vdcs=__ret__.participating_org_vdcs,
        status=__ret__.status,
        type=__ret__.type,
        universal_networking_enabled=__ret__.universal_networking_enabled)


@_utilities.lift_output_func(get_vdc_group)
def get_vdc_group_output(default_policy_status: Optional[pulumi.Input[Optional[bool]]] = None,
                         description: Optional[pulumi.Input[Optional[str]]] = None,
                         error_message: Optional[pulumi.Input[Optional[str]]] = None,
                         id: Optional[pulumi.Input[Optional[str]]] = None,
                         local_egress: Optional[pulumi.Input[Optional[bool]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         network_pool_id: Optional[pulumi.Input[Optional[str]]] = None,
                         network_pool_universal_id: Optional[pulumi.Input[Optional[str]]] = None,
                         network_provider_type: Optional[pulumi.Input[Optional[str]]] = None,
                         org: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         type: Optional[pulumi.Input[Optional[str]]] = None,
                         universal_networking_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVdcGroupResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
