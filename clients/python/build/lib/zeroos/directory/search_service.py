class SearchService:
    def __init__(self, client):
        self.client = client



    def search_get(self, headers=None, query_params=None, content_type="application/json"):
        """
        Search for resource pools.
        It is method for GET /search
        """
        uri = self.client.base_url + "/search"
        return self.client.get(uri, None, headers, query_params, content_type)
