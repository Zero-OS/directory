"""
Auto-generated class for ResourcePool
"""
from .EnumResourcePoolCreateDatacenter_tier import EnumResourcePoolCreateDatacenter_tier
from .EnumResourcePoolCreatePricing_currency import EnumResourcePoolCreatePricing_currency
from .Location import Location
from .Node import Node
from .PoolQuality import PoolQuality
from .Pricing import Pricing

from . import client_support


class ResourcePool(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(Location, Nodes, Pricings, UID, cu_max, cu_planned, datacenter_tier, datacenter_uptime_sla_min, ipv4_enabled, ipv4_nr_pub, ipv6_enabled, network_redundant, network_speed, organization, pricing_currency, provider, quality, su_max, su_planned, tu_max, tu_planned):
        """
        :type Location: Location
        :type Nodes: list[Node]
        :type Pricings: list[Pricing]
        :type UID: str
        :type cu_max: int
        :type cu_planned: int
        :type datacenter_tier: EnumResourcePoolCreateDatacenter_tier
        :type datacenter_uptime_sla_min: float
        :type ipv4_enabled: bool
        :type ipv4_nr_pub: int
        :type ipv6_enabled: bool
        :type network_redundant: bool
        :type network_speed: int
        :type organization: str
        :type pricing_currency: EnumResourcePoolCreatePricing_currency
        :type provider: str
        :type quality: PoolQuality
        :type su_max: int
        :type su_planned: int
        :type tu_max: int
        :type tu_planned: int
        :rtype: ResourcePool
        """

        return ResourcePool(
            Location=Location,
            Nodes=Nodes,
            Pricings=Pricings,
            UID=UID,
            cu_max=cu_max,
            cu_planned=cu_planned,
            datacenter_tier=datacenter_tier,
            datacenter_uptime_sla_min=datacenter_uptime_sla_min,
            ipv4_enabled=ipv4_enabled,
            ipv4_nr_pub=ipv4_nr_pub,
            ipv6_enabled=ipv6_enabled,
            network_redundant=network_redundant,
            network_speed=network_speed,
            organization=organization,
            pricing_currency=pricing_currency,
            provider=provider,
            quality=quality,
            su_max=su_max,
            su_planned=su_planned,
            tu_max=tu_max,
            tu_planned=tu_planned,
        )

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'ResourcePool'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'Location'
        val = data.get(property_name)
        if val is not None:
            datatypes = [Location]
            try:
                self.Location = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'Nodes'
        val = data.get(property_name)
        if val is not None:
            datatypes = [Node]
            try:
                self.Nodes = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'Pricings'
        val = data.get(property_name)
        if val is not None:
            datatypes = [Pricing]
            try:
                self.Pricings = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'UID'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.UID = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'cu_max'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.cu_max = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'cu_planned'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.cu_planned = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'datacenter_tier'
        val = data.get(property_name)
        if val is not None:
            datatypes = [EnumResourcePoolCreateDatacenter_tier]
            try:
                self.datacenter_tier = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'datacenter_uptime_sla_min'
        val = data.get(property_name)
        if val is not None:
            datatypes = [float]
            try:
                self.datacenter_uptime_sla_min = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'ipv4_enabled'
        val = data.get(property_name)
        if val is not None:
            datatypes = [bool]
            try:
                self.ipv4_enabled = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'ipv4_nr_pub'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.ipv4_nr_pub = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'ipv6_enabled'
        val = data.get(property_name)
        if val is not None:
            datatypes = [bool]
            try:
                self.ipv6_enabled = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'network_redundant'
        val = data.get(property_name)
        if val is not None:
            datatypes = [bool]
            try:
                self.network_redundant = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'network_speed'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.network_speed = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'organization'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.organization = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'pricing_currency'
        val = data.get(property_name)
        if val is not None:
            datatypes = [EnumResourcePoolCreatePricing_currency]
            try:
                self.pricing_currency = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'provider'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.provider = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'quality'
        val = data.get(property_name)
        if val is not None:
            datatypes = [PoolQuality]
            try:
                self.quality = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'su_max'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.su_max = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'su_planned'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.su_planned = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'tu_max'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.tu_max = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'tu_planned'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.tu_planned = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
