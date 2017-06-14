import requests


from .DownPeriod import DownPeriod
from .EnumResourcePoolCreateDatacenter_tier import EnumResourcePoolCreateDatacenter_tier
from .EnumResourcePoolCreateDatacenter_tier import EnumResourcePoolCreateDatacenter_tier
from .EnumResourcePoolCreatePricing_currency import EnumResourcePoolCreatePricing_currency
from .EnumResourcePoolCreatePricing_currency import EnumResourcePoolCreatePricing_currency
from .Location import Location
from .Node import Node
from .PoolQuality import PoolQuality
from .Pricing import Pricing
from .Provider import Provider
from .ResourcePool import ResourcePool
from .ResourcePoolCreate import ResourcePoolCreate

from .client import Client as APIClient

from .oauth2_client_itsyouonline import Oauth2ClientItsyouonline

class Client:
    def __init__(self, base_uri="https://discovery.gig.tech"):
        self.api = APIClient(base_uri)
        
        self.oauth2_client_itsyouonline = Oauth2ClientItsyouonline()