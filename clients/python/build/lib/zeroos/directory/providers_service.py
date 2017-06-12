class ProvidersService:
    def __init__(self, client):
        self.client = client



    def ListNodes(self, poolid, headers=None, query_params=None, content_type="application/json"):
        """
        List all nodes
        It is method for GET /providers/resource_pools/{poolid}/nodes
        """
        uri = self.client.base_url + "/providers/resource_pools/"+poolid+"/nodes"
        return self.client.get(uri, None, headers, query_params, content_type)


    def DeleteResourcePool(self, poolid, headers=None, query_params=None, content_type="application/json"):
        """
        Delete resource_pool
        It is method for DELETE /providers/resource_pools/{poolid}
        """
        uri = self.client.base_url + "/providers/resource_pools/"+poolid
        return self.client.delete(uri, None, headers, query_params, content_type)


    def GetResourcePool(self, poolid, headers=None, query_params=None, content_type="application/json"):
        """
        Get detail view about resource_pool
        It is method for GET /providers/resource_pools/{poolid}
        """
        uri = self.client.base_url + "/providers/resource_pools/"+poolid
        return self.client.get(uri, None, headers, query_params, content_type)


    def UpdateResourcePool(self, poolid, headers=None, query_params=None, content_type="application/json"):
        """
        Update resource_pool
        It is method for PUT /providers/resource_pools/{poolid}
        """
        uri = self.client.base_url + "/providers/resource_pools/"+poolid
        return self.client.put(uri, None, headers, query_params, content_type)


    def ListResourcePools(self, headers=None, query_params=None, content_type="application/json"):
        """
        List all resource_pools
        It is method for GET /providers/resource_pools
        """
        uri = self.client.base_url + "/providers/resource_pools"
        return self.client.get(uri, None, headers, query_params, content_type)


    def CreateResourcePool(self, data, headers=None, query_params=None, content_type="application/json"):
        """
        Create a new resource_pool
        It is method for POST /providers/resource_pools
        """
        uri = self.client.base_url + "/providers/resource_pools"
        return self.client.post(uri, data, headers, query_params, content_type)


    def ListProviders(self, headers=None, query_params=None, content_type="application/json"):
        """
        List all providers
        It is method for GET /providers
        """
        uri = self.client.base_url + "/providers"
        return self.client.get(uri, None, headers, query_params, content_type)
