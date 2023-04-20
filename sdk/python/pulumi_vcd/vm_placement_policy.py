# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VmPlacementPolicyArgs', 'VmPlacementPolicy']

@pulumi.input_type
class VmPlacementPolicyArgs:
    def __init__(__self__, *,
                 provider_vdc_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 logical_vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VmPlacementPolicy resource.
        :param pulumi.Input[str] provider_vdc_id: ID of the Provider VDC to which the VM Placement Policy belongs
        :param pulumi.Input[str] description: Description of the VM Placement Policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] logical_vm_group_ids: IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
               entries set in this attribute
        :param pulumi.Input[str] name: Name of the VM Placement Policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vm_group_ids: IDs of the collection of VMs with similar host requirements
        """
        pulumi.set(__self__, "provider_vdc_id", provider_vdc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if logical_vm_group_ids is not None:
            pulumi.set(__self__, "logical_vm_group_ids", logical_vm_group_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vm_group_ids is not None:
            pulumi.set(__self__, "vm_group_ids", vm_group_ids)

    @property
    @pulumi.getter(name="providerVdcId")
    def provider_vdc_id(self) -> pulumi.Input[str]:
        """
        ID of the Provider VDC to which the VM Placement Policy belongs
        """
        return pulumi.get(self, "provider_vdc_id")

    @provider_vdc_id.setter
    def provider_vdc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_vdc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the VM Placement Policy
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="logicalVmGroupIds")
    def logical_vm_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
        entries set in this attribute
        """
        return pulumi.get(self, "logical_vm_group_ids")

    @logical_vm_group_ids.setter
    def logical_vm_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "logical_vm_group_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the VM Placement Policy
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="vmGroupIds")
    def vm_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of the collection of VMs with similar host requirements
        """
        return pulumi.get(self, "vm_group_ids")

    @vm_group_ids.setter
    def vm_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vm_group_ids", value)


@pulumi.input_type
class _VmPlacementPolicyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 logical_vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_vdc_id: Optional[pulumi.Input[str]] = None,
                 vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering VmPlacementPolicy resources.
        :param pulumi.Input[str] description: Description of the VM Placement Policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] logical_vm_group_ids: IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
               entries set in this attribute
        :param pulumi.Input[str] name: Name of the VM Placement Policy
        :param pulumi.Input[str] provider_vdc_id: ID of the Provider VDC to which the VM Placement Policy belongs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vm_group_ids: IDs of the collection of VMs with similar host requirements
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if logical_vm_group_ids is not None:
            pulumi.set(__self__, "logical_vm_group_ids", logical_vm_group_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if provider_vdc_id is not None:
            pulumi.set(__self__, "provider_vdc_id", provider_vdc_id)
        if vm_group_ids is not None:
            pulumi.set(__self__, "vm_group_ids", vm_group_ids)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the VM Placement Policy
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="logicalVmGroupIds")
    def logical_vm_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
        entries set in this attribute
        """
        return pulumi.get(self, "logical_vm_group_ids")

    @logical_vm_group_ids.setter
    def logical_vm_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "logical_vm_group_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the VM Placement Policy
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="providerVdcId")
    def provider_vdc_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Provider VDC to which the VM Placement Policy belongs
        """
        return pulumi.get(self, "provider_vdc_id")

    @provider_vdc_id.setter
    def provider_vdc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_vdc_id", value)

    @property
    @pulumi.getter(name="vmGroupIds")
    def vm_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of the collection of VMs with similar host requirements
        """
        return pulumi.get(self, "vm_group_ids")

    @vm_group_ids.setter
    def vm_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vm_group_ids", value)


class VmPlacementPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 logical_vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_vdc_id: Optional[pulumi.Input[str]] = None,
                 vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a VmPlacementPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the VM Placement Policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] logical_vm_group_ids: IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
               entries set in this attribute
        :param pulumi.Input[str] name: Name of the VM Placement Policy
        :param pulumi.Input[str] provider_vdc_id: ID of the Provider VDC to which the VM Placement Policy belongs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vm_group_ids: IDs of the collection of VMs with similar host requirements
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VmPlacementPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VmPlacementPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VmPlacementPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VmPlacementPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 logical_vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_vdc_id: Optional[pulumi.Input[str]] = None,
                 vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VmPlacementPolicyArgs.__new__(VmPlacementPolicyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["logical_vm_group_ids"] = logical_vm_group_ids
            __props__.__dict__["name"] = name
            if provider_vdc_id is None and not opts.urn:
                raise TypeError("Missing required property 'provider_vdc_id'")
            __props__.__dict__["provider_vdc_id"] = provider_vdc_id
            __props__.__dict__["vm_group_ids"] = vm_group_ids
        super(VmPlacementPolicy, __self__).__init__(
            'vcd:index/vmPlacementPolicy:VmPlacementPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            logical_vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            provider_vdc_id: Optional[pulumi.Input[str]] = None,
            vm_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'VmPlacementPolicy':
        """
        Get an existing VmPlacementPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the VM Placement Policy
        :param pulumi.Input[Sequence[pulumi.Input[str]]] logical_vm_group_ids: IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
               entries set in this attribute
        :param pulumi.Input[str] name: Name of the VM Placement Policy
        :param pulumi.Input[str] provider_vdc_id: ID of the Provider VDC to which the VM Placement Policy belongs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vm_group_ids: IDs of the collection of VMs with similar host requirements
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VmPlacementPolicyState.__new__(_VmPlacementPolicyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["logical_vm_group_ids"] = logical_vm_group_ids
        __props__.__dict__["name"] = name
        __props__.__dict__["provider_vdc_id"] = provider_vdc_id
        __props__.__dict__["vm_group_ids"] = vm_group_ids
        return VmPlacementPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the VM Placement Policy
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="logicalVmGroupIds")
    def logical_vm_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
        entries set in this attribute
        """
        return pulumi.get(self, "logical_vm_group_ids")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the VM Placement Policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="providerVdcId")
    def provider_vdc_id(self) -> pulumi.Output[str]:
        """
        ID of the Provider VDC to which the VM Placement Policy belongs
        """
        return pulumi.get(self, "provider_vdc_id")

    @property
    @pulumi.getter(name="vmGroupIds")
    def vm_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        IDs of the collection of VMs with similar host requirements
        """
        return pulumi.get(self, "vm_group_ids")
