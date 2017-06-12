"""
Auto-generated class for Pricing
"""

from . import client_support


class Pricing(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(cu, min_nr_months, su, tu):
        """
        :type cu: int
        :type min_nr_months: int
        :type su: int
        :type tu: int
        :rtype: Pricing
        """

        return Pricing(
            cu=cu,
            min_nr_months=min_nr_months,
            su=su,
            tu=tu,
        )

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'Pricing'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'cu'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.cu = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'min_nr_months'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.min_nr_months = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'su'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.su = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'tu'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.tu = client_support.val_factory(val, datatypes)
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
