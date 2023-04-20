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
from ._inputs import *

__all__ = ['VmSizingPolicyArgs', 'VmSizingPolicy']

@pulumi.input_type
class VmSizingPolicyArgs:
    def __init__(__self__, *,
                 cpu: Optional[pulumi.Input['VmSizingPolicyCpuArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input['VmSizingPolicyMemoryArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VmSizingPolicy resource.
        :param pulumi.Input[str] org: The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            warnings.warn("""Unneeded property, which was included by mistake""", DeprecationWarning)
            pulumi.log.warn("""org is deprecated: Unneeded property, which was included by mistake""")
        if org is not None:
            pulumi.set(__self__, "org", org)

    @property
    @pulumi.getter
    def cpu(self) -> Optional[pulumi.Input['VmSizingPolicyCpuArgs']]:
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: Optional[pulumi.Input['VmSizingPolicyCpuArgs']]):
        pulumi.set(self, "cpu", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input['VmSizingPolicyMemoryArgs']]:
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input['VmSizingPolicyMemoryArgs']]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)


@pulumi.input_type
class _VmSizingPolicyState:
    def __init__(__self__, *,
                 cpu: Optional[pulumi.Input['VmSizingPolicyCpuArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input['VmSizingPolicyMemoryArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VmSizingPolicy resources.
        :param pulumi.Input[str] org: The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            warnings.warn("""Unneeded property, which was included by mistake""", DeprecationWarning)
            pulumi.log.warn("""org is deprecated: Unneeded property, which was included by mistake""")
        if org is not None:
            pulumi.set(__self__, "org", org)

    @property
    @pulumi.getter
    def cpu(self) -> Optional[pulumi.Input['VmSizingPolicyCpuArgs']]:
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: Optional[pulumi.Input['VmSizingPolicyCpuArgs']]):
        pulumi.set(self, "cpu", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input['VmSizingPolicyMemoryArgs']]:
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input['VmSizingPolicyMemoryArgs']]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)


class VmSizingPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu: Optional[pulumi.Input[pulumi.InputType['VmSizingPolicyCpuArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[pulumi.InputType['VmSizingPolicyMemoryArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VmSizingPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org: The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VmSizingPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VmSizingPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VmSizingPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VmSizingPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu: Optional[pulumi.Input[pulumi.InputType['VmSizingPolicyCpuArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[pulumi.InputType['VmSizingPolicyMemoryArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VmSizingPolicyArgs.__new__(VmSizingPolicyArgs)

            __props__.__dict__["cpu"] = cpu
            __props__.__dict__["description"] = description
            __props__.__dict__["memory"] = memory
            __props__.__dict__["name"] = name
            if org is not None and not opts.urn:
                warnings.warn("""Unneeded property, which was included by mistake""", DeprecationWarning)
                pulumi.log.warn("""org is deprecated: Unneeded property, which was included by mistake""")
            __props__.__dict__["org"] = org
        super(VmSizingPolicy, __self__).__init__(
            'vcd:index/vmSizingPolicy:VmSizingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cpu: Optional[pulumi.Input[pulumi.InputType['VmSizingPolicyCpuArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            memory: Optional[pulumi.Input[pulumi.InputType['VmSizingPolicyMemoryArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None) -> 'VmSizingPolicy':
        """
        Get an existing VmSizingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org: The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VmSizingPolicyState.__new__(_VmSizingPolicyState)

        __props__.__dict__["cpu"] = cpu
        __props__.__dict__["description"] = description
        __props__.__dict__["memory"] = memory
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        return VmSizingPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cpu(self) -> pulumi.Output[Optional['outputs.VmSizingPolicyCpu']]:
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Output[Optional['outputs.VmSizingPolicyMemory']]:
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use - Deprecated and unneeded: will be ignored if used
        """
        return pulumi.get(self, "org")

