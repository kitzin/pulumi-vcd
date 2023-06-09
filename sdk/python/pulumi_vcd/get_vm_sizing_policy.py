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
    'GetVmSizingPolicyResult',
    'AwaitableGetVmSizingPolicyResult',
    'get_vm_sizing_policy',
    'get_vm_sizing_policy_output',
]

@pulumi.output_type
class GetVmSizingPolicyResult:
    """
    A collection of values returned by getVmSizingPolicy.
    """
    def __init__(__self__, cpus=None, description=None, id=None, memories=None, name=None, org=None):
        if cpus and not isinstance(cpus, list):
            raise TypeError("Expected argument 'cpus' to be a list")
        pulumi.set(__self__, "cpus", cpus)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if memories and not isinstance(memories, list):
            raise TypeError("Expected argument 'memories' to be a list")
        pulumi.set(__self__, "memories", memories)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        if org is not None:
            warnings.warn("""Unneeded property, which was included by mistake""", DeprecationWarning)
            pulumi.log.warn("""org is deprecated: Unneeded property, which was included by mistake""")

        pulumi.set(__self__, "org", org)

    @property
    @pulumi.getter
    def cpus(self) -> Sequence['outputs.GetVmSizingPolicyCpusResult']:
        return pulumi.get(self, "cpus")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def memories(self) -> Sequence['outputs.GetVmSizingPolicyMemoryResult']:
        return pulumi.get(self, "memories")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")


class AwaitableGetVmSizingPolicyResult(GetVmSizingPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVmSizingPolicyResult(
            cpus=self.cpus,
            description=self.description,
            id=self.id,
            memories=self.memories,
            name=self.name,
            org=self.org)


def get_vm_sizing_policy(name: Optional[str] = None,
                         org: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVmSizingPolicyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['org'] = org
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getVmSizingPolicy:getVmSizingPolicy', __args__, opts=opts, typ=GetVmSizingPolicyResult).value

    return AwaitableGetVmSizingPolicyResult(
        cpus=__ret__.cpus,
        description=__ret__.description,
        id=__ret__.id,
        memories=__ret__.memories,
        name=__ret__.name,
        org=__ret__.org)


@_utilities.lift_output_func(get_vm_sizing_policy)
def get_vm_sizing_policy_output(name: Optional[pulumi.Input[str]] = None,
                                org: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVmSizingPolicyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
