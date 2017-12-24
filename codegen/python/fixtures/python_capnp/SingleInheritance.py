"""
Auto-generated class for SingleInheritance
"""
import capnp
import os
from .EnumCity import EnumCity
from six import string_types

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class SingleInheritance(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :type single: bool
        :rtype: SingleInheritance
        """

        return SingleInheritance(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'SingleInheritance'
        data = json or kwargs

        # set attributes
        data_types = [EnumCity]
        self.cities = client_support.set_property('cities', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.colours = client_support.set_property('colours', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.name = client_support.set_property('name', data, data_types, False, [], False, True, class_name)
        data_types = [bool]
        self.single = client_support.set_property('single', data, data_types, False, [], False, True, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)

    def to_capnp(self):
        """
        Load the class in capnp schema SingleInheritance.capnp
        :rtype bytes
        """
        template = capnp.load('%s/SingleInheritance.capnp' % dir)
        return template.SingleInheritance.new_message(**self.as_dict()).to_bytes()


class SingleInheritanceCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(binary=None):
        """
        Load the binary of SingleInheritance.capnp into class SingleInheritance
        :type binary: bytes. If none creates an empty capnp object.
        rtype: SingleInheritance
        """
        template = capnp.load('%s/SingleInheritance.capnp' % dir)
        struct = template.SingleInheritance.from_bytes(binary) if binary else template.SingleInheritance.new_message()
        return SingleInheritance(**struct.to_dict(verbose=True))