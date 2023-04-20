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
    'GetVappVmResult',
    'AwaitableGetVappVmResult',
    'get_vapp_vm',
    'get_vapp_vm_output',
]

@pulumi.output_type
class GetVappVmResult:
    """
    A collection of values returned by getVappVm.
    """
    def __init__(__self__, computer_name=None, cpu_cores=None, cpu_hot_add_enabled=None, cpu_limit=None, cpu_priority=None, cpu_reservation=None, cpu_shares=None, cpus=None, customizations=None, description=None, disks=None, expose_hardware_virtualization=None, guest_properties=None, hardware_version=None, href=None, id=None, internal_disks=None, memory=None, memory_hot_add_enabled=None, memory_limit=None, memory_priority=None, memory_reservation=None, memory_shares=None, metadata=None, metadata_entries=None, name=None, network_dhcp_wait_seconds=None, networks=None, org=None, os_type=None, placement_policy_id=None, sizing_policy_id=None, status=None, status_text=None, storage_profile=None, vapp_name=None, vdc=None, vm_type=None):
        if computer_name and not isinstance(computer_name, str):
            raise TypeError("Expected argument 'computer_name' to be a str")
        pulumi.set(__self__, "computer_name", computer_name)
        if cpu_cores and not isinstance(cpu_cores, int):
            raise TypeError("Expected argument 'cpu_cores' to be a int")
        pulumi.set(__self__, "cpu_cores", cpu_cores)
        if cpu_hot_add_enabled and not isinstance(cpu_hot_add_enabled, bool):
            raise TypeError("Expected argument 'cpu_hot_add_enabled' to be a bool")
        pulumi.set(__self__, "cpu_hot_add_enabled", cpu_hot_add_enabled)
        if cpu_limit and not isinstance(cpu_limit, int):
            raise TypeError("Expected argument 'cpu_limit' to be a int")
        pulumi.set(__self__, "cpu_limit", cpu_limit)
        if cpu_priority and not isinstance(cpu_priority, str):
            raise TypeError("Expected argument 'cpu_priority' to be a str")
        pulumi.set(__self__, "cpu_priority", cpu_priority)
        if cpu_reservation and not isinstance(cpu_reservation, int):
            raise TypeError("Expected argument 'cpu_reservation' to be a int")
        pulumi.set(__self__, "cpu_reservation", cpu_reservation)
        if cpu_shares and not isinstance(cpu_shares, int):
            raise TypeError("Expected argument 'cpu_shares' to be a int")
        pulumi.set(__self__, "cpu_shares", cpu_shares)
        if cpus and not isinstance(cpus, int):
            raise TypeError("Expected argument 'cpus' to be a int")
        pulumi.set(__self__, "cpus", cpus)
        if customizations and not isinstance(customizations, list):
            raise TypeError("Expected argument 'customizations' to be a list")
        pulumi.set(__self__, "customizations", customizations)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disks and not isinstance(disks, list):
            raise TypeError("Expected argument 'disks' to be a list")
        pulumi.set(__self__, "disks", disks)
        if expose_hardware_virtualization and not isinstance(expose_hardware_virtualization, bool):
            raise TypeError("Expected argument 'expose_hardware_virtualization' to be a bool")
        pulumi.set(__self__, "expose_hardware_virtualization", expose_hardware_virtualization)
        if guest_properties and not isinstance(guest_properties, dict):
            raise TypeError("Expected argument 'guest_properties' to be a dict")
        pulumi.set(__self__, "guest_properties", guest_properties)
        if hardware_version and not isinstance(hardware_version, str):
            raise TypeError("Expected argument 'hardware_version' to be a str")
        pulumi.set(__self__, "hardware_version", hardware_version)
        if href and not isinstance(href, str):
            raise TypeError("Expected argument 'href' to be a str")
        pulumi.set(__self__, "href", href)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_disks and not isinstance(internal_disks, list):
            raise TypeError("Expected argument 'internal_disks' to be a list")
        pulumi.set(__self__, "internal_disks", internal_disks)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if memory_hot_add_enabled and not isinstance(memory_hot_add_enabled, bool):
            raise TypeError("Expected argument 'memory_hot_add_enabled' to be a bool")
        pulumi.set(__self__, "memory_hot_add_enabled", memory_hot_add_enabled)
        if memory_limit and not isinstance(memory_limit, int):
            raise TypeError("Expected argument 'memory_limit' to be a int")
        pulumi.set(__self__, "memory_limit", memory_limit)
        if memory_priority and not isinstance(memory_priority, str):
            raise TypeError("Expected argument 'memory_priority' to be a str")
        pulumi.set(__self__, "memory_priority", memory_priority)
        if memory_reservation and not isinstance(memory_reservation, int):
            raise TypeError("Expected argument 'memory_reservation' to be a int")
        pulumi.set(__self__, "memory_reservation", memory_reservation)
        if memory_shares and not isinstance(memory_shares, int):
            raise TypeError("Expected argument 'memory_shares' to be a int")
        pulumi.set(__self__, "memory_shares", memory_shares)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        if metadata is not None:
            warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
            pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")

        pulumi.set(__self__, "metadata", metadata)
        if metadata_entries and not isinstance(metadata_entries, list):
            raise TypeError("Expected argument 'metadata_entries' to be a list")
        pulumi.set(__self__, "metadata_entries", metadata_entries)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_dhcp_wait_seconds and not isinstance(network_dhcp_wait_seconds, int):
            raise TypeError("Expected argument 'network_dhcp_wait_seconds' to be a int")
        pulumi.set(__self__, "network_dhcp_wait_seconds", network_dhcp_wait_seconds)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if placement_policy_id and not isinstance(placement_policy_id, str):
            raise TypeError("Expected argument 'placement_policy_id' to be a str")
        pulumi.set(__self__, "placement_policy_id", placement_policy_id)
        if sizing_policy_id and not isinstance(sizing_policy_id, str):
            raise TypeError("Expected argument 'sizing_policy_id' to be a str")
        pulumi.set(__self__, "sizing_policy_id", sizing_policy_id)
        if status and not isinstance(status, int):
            raise TypeError("Expected argument 'status' to be a int")
        pulumi.set(__self__, "status", status)
        if status_text and not isinstance(status_text, str):
            raise TypeError("Expected argument 'status_text' to be a str")
        pulumi.set(__self__, "status_text", status_text)
        if storage_profile and not isinstance(storage_profile, str):
            raise TypeError("Expected argument 'storage_profile' to be a str")
        pulumi.set(__self__, "storage_profile", storage_profile)
        if vapp_name and not isinstance(vapp_name, str):
            raise TypeError("Expected argument 'vapp_name' to be a str")
        pulumi.set(__self__, "vapp_name", vapp_name)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)
        if vm_type and not isinstance(vm_type, str):
            raise TypeError("Expected argument 'vm_type' to be a str")
        pulumi.set(__self__, "vm_type", vm_type)

    @property
    @pulumi.getter(name="computerName")
    def computer_name(self) -> str:
        return pulumi.get(self, "computer_name")

    @property
    @pulumi.getter(name="cpuCores")
    def cpu_cores(self) -> int:
        return pulumi.get(self, "cpu_cores")

    @property
    @pulumi.getter(name="cpuHotAddEnabled")
    def cpu_hot_add_enabled(self) -> bool:
        return pulumi.get(self, "cpu_hot_add_enabled")

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> int:
        return pulumi.get(self, "cpu_limit")

    @property
    @pulumi.getter(name="cpuPriority")
    def cpu_priority(self) -> str:
        return pulumi.get(self, "cpu_priority")

    @property
    @pulumi.getter(name="cpuReservation")
    def cpu_reservation(self) -> int:
        return pulumi.get(self, "cpu_reservation")

    @property
    @pulumi.getter(name="cpuShares")
    def cpu_shares(self) -> int:
        return pulumi.get(self, "cpu_shares")

    @property
    @pulumi.getter
    def cpus(self) -> int:
        return pulumi.get(self, "cpus")

    @property
    @pulumi.getter
    def customizations(self) -> Sequence['outputs.GetVappVmCustomizationResult']:
        return pulumi.get(self, "customizations")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disks(self) -> Sequence['outputs.GetVappVmDiskResult']:
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter(name="exposeHardwareVirtualization")
    def expose_hardware_virtualization(self) -> bool:
        return pulumi.get(self, "expose_hardware_virtualization")

    @property
    @pulumi.getter(name="guestProperties")
    def guest_properties(self) -> Mapping[str, Any]:
        return pulumi.get(self, "guest_properties")

    @property
    @pulumi.getter(name="hardwareVersion")
    def hardware_version(self) -> str:
        return pulumi.get(self, "hardware_version")

    @property
    @pulumi.getter
    def href(self) -> str:
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalDisks")
    def internal_disks(self) -> Sequence['outputs.GetVappVmInternalDiskResult']:
        return pulumi.get(self, "internal_disks")

    @property
    @pulumi.getter
    def memory(self) -> int:
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="memoryHotAddEnabled")
    def memory_hot_add_enabled(self) -> bool:
        return pulumi.get(self, "memory_hot_add_enabled")

    @property
    @pulumi.getter(name="memoryLimit")
    def memory_limit(self) -> int:
        return pulumi.get(self, "memory_limit")

    @property
    @pulumi.getter(name="memoryPriority")
    def memory_priority(self) -> str:
        return pulumi.get(self, "memory_priority")

    @property
    @pulumi.getter(name="memoryReservation")
    def memory_reservation(self) -> int:
        return pulumi.get(self, "memory_reservation")

    @property
    @pulumi.getter(name="memoryShares")
    def memory_shares(self) -> int:
        return pulumi.get(self, "memory_shares")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, Any]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> Sequence['outputs.GetVappVmMetadataEntryResult']:
        return pulumi.get(self, "metadata_entries")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkDhcpWaitSeconds")
    def network_dhcp_wait_seconds(self) -> Optional[int]:
        return pulumi.get(self, "network_dhcp_wait_seconds")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetVappVmNetworkResult']:
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> str:
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="placementPolicyId")
    def placement_policy_id(self) -> str:
        return pulumi.get(self, "placement_policy_id")

    @property
    @pulumi.getter(name="sizingPolicyId")
    def sizing_policy_id(self) -> str:
        return pulumi.get(self, "sizing_policy_id")

    @property
    @pulumi.getter
    def status(self) -> int:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusText")
    def status_text(self) -> str:
        return pulumi.get(self, "status_text")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> str:
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter(name="vappName")
    def vapp_name(self) -> str:
        return pulumi.get(self, "vapp_name")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")

    @property
    @pulumi.getter(name="vmType")
    def vm_type(self) -> str:
        return pulumi.get(self, "vm_type")


class AwaitableGetVappVmResult(GetVappVmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVappVmResult(
            computer_name=self.computer_name,
            cpu_cores=self.cpu_cores,
            cpu_hot_add_enabled=self.cpu_hot_add_enabled,
            cpu_limit=self.cpu_limit,
            cpu_priority=self.cpu_priority,
            cpu_reservation=self.cpu_reservation,
            cpu_shares=self.cpu_shares,
            cpus=self.cpus,
            customizations=self.customizations,
            description=self.description,
            disks=self.disks,
            expose_hardware_virtualization=self.expose_hardware_virtualization,
            guest_properties=self.guest_properties,
            hardware_version=self.hardware_version,
            href=self.href,
            id=self.id,
            internal_disks=self.internal_disks,
            memory=self.memory,
            memory_hot_add_enabled=self.memory_hot_add_enabled,
            memory_limit=self.memory_limit,
            memory_priority=self.memory_priority,
            memory_reservation=self.memory_reservation,
            memory_shares=self.memory_shares,
            metadata=self.metadata,
            metadata_entries=self.metadata_entries,
            name=self.name,
            network_dhcp_wait_seconds=self.network_dhcp_wait_seconds,
            networks=self.networks,
            org=self.org,
            os_type=self.os_type,
            placement_policy_id=self.placement_policy_id,
            sizing_policy_id=self.sizing_policy_id,
            status=self.status,
            status_text=self.status_text,
            storage_profile=self.storage_profile,
            vapp_name=self.vapp_name,
            vdc=self.vdc,
            vm_type=self.vm_type)


def get_vapp_vm(name: Optional[str] = None,
                network_dhcp_wait_seconds: Optional[int] = None,
                org: Optional[str] = None,
                placement_policy_id: Optional[str] = None,
                sizing_policy_id: Optional[str] = None,
                vapp_name: Optional[str] = None,
                vdc: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVappVmResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['networkDhcpWaitSeconds'] = network_dhcp_wait_seconds
    __args__['org'] = org
    __args__['placementPolicyId'] = placement_policy_id
    __args__['sizingPolicyId'] = sizing_policy_id
    __args__['vappName'] = vapp_name
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getVappVm:getVappVm', __args__, opts=opts, typ=GetVappVmResult).value

    return AwaitableGetVappVmResult(
        computer_name=__ret__.computer_name,
        cpu_cores=__ret__.cpu_cores,
        cpu_hot_add_enabled=__ret__.cpu_hot_add_enabled,
        cpu_limit=__ret__.cpu_limit,
        cpu_priority=__ret__.cpu_priority,
        cpu_reservation=__ret__.cpu_reservation,
        cpu_shares=__ret__.cpu_shares,
        cpus=__ret__.cpus,
        customizations=__ret__.customizations,
        description=__ret__.description,
        disks=__ret__.disks,
        expose_hardware_virtualization=__ret__.expose_hardware_virtualization,
        guest_properties=__ret__.guest_properties,
        hardware_version=__ret__.hardware_version,
        href=__ret__.href,
        id=__ret__.id,
        internal_disks=__ret__.internal_disks,
        memory=__ret__.memory,
        memory_hot_add_enabled=__ret__.memory_hot_add_enabled,
        memory_limit=__ret__.memory_limit,
        memory_priority=__ret__.memory_priority,
        memory_reservation=__ret__.memory_reservation,
        memory_shares=__ret__.memory_shares,
        metadata=__ret__.metadata,
        metadata_entries=__ret__.metadata_entries,
        name=__ret__.name,
        network_dhcp_wait_seconds=__ret__.network_dhcp_wait_seconds,
        networks=__ret__.networks,
        org=__ret__.org,
        os_type=__ret__.os_type,
        placement_policy_id=__ret__.placement_policy_id,
        sizing_policy_id=__ret__.sizing_policy_id,
        status=__ret__.status,
        status_text=__ret__.status_text,
        storage_profile=__ret__.storage_profile,
        vapp_name=__ret__.vapp_name,
        vdc=__ret__.vdc,
        vm_type=__ret__.vm_type)


@_utilities.lift_output_func(get_vapp_vm)
def get_vapp_vm_output(name: Optional[pulumi.Input[str]] = None,
                       network_dhcp_wait_seconds: Optional[pulumi.Input[Optional[int]]] = None,
                       org: Optional[pulumi.Input[Optional[str]]] = None,
                       placement_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                       sizing_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                       vapp_name: Optional[pulumi.Input[str]] = None,
                       vdc: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVappVmResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
