"""
Auto-generated class for PoolQuality
"""
from .DownPeriod import DownPeriod

from . import client_support


class PoolQuality(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(average_restauration_time, down_periods, uptime):
        """
        :type average_restauration_time: int
        :type down_periods: list[DownPeriod]
        :type uptime: int
        :rtype: PoolQuality
        """

        return PoolQuality(
            average_restauration_time=average_restauration_time,
            down_periods=down_periods,
            uptime=uptime,
        )

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'PoolQuality'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'average_restauration_time'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.average_restauration_time = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'down_periods'
        val = data.get(property_name)
        if val is not None:
            datatypes = [DownPeriod]
            try:
                self.down_periods = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'uptime'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.uptime = client_support.val_factory(val, datatypes)
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
