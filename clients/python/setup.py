from setuptools import setup, find_packages
# To use a consistent encoding
from codecs import open
from os import path

here = path.abspath(path.dirname(__file__))

# Get the long description from the README file
with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='0-directory',
    version='0.0.1',
    description='Zero-OS Directory',
    long_description=long_description,
    url='https://github.com/zero-os/0-directory',
    author='Christophe de Carvalho',
    author_email='christophe@gig.tech',
    license='Apache 2.0',
    packages=find_packages(),
    namespace_packages=['zeroos'],
    install_requires=['requests'],
)
