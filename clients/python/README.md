# Python Client

O-stor is the Python client used to talk to Zero-OS 0-Stor REST API

## Install
```shell
pip install 0-directory
```

## Example:
### List resource pools:
```python
from zeroos.directory import Client

cl = Client('http://localhost:8080')    
resp = cl.api.providers.ListProviders()
for item in resp.json():
    provid = item['name']
    print("ResourcePool from provider {}".format(provid))
    pools = cl.api.providers.ListResourcePools(query_params={'providerid': provid}).json()
    for pool in pools:
        print(pool)
```

## To update the client from the RAML file
```shell
go-raml client -l python --ramlfile ../../specs/directory.raml --dir zeroos/directory/
```
