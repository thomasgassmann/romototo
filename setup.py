import json
from typing import List

from setuptools import find_packages, setup

from romototo import __version__

install_requires: List[str] = []

with open('README.md', 'r') as fh:
    long_description = fh.read()

with open('Pipfile.lock') as fd:
    lock_data = json.load(fd)
    install_requires = [
        package_name + package_data['version']
        for package_name, package_data in lock_data['default'].items()
    ]


setup(
    name='romototo',
    version=__version__,
    license='MIT',
    author='Thomas Gassmann',
    long_description=long_description,
    description='romototo housing',
    long_description_content_type='text/markdown',
    author_email='thomas@gassmann.dev',
    url='https://github.com/thomasgassmann/romototo',
    install_requires=install_requires,
    packages=find_packages(),
    include_package_data=True,
    entry_points={
        'console_scripts': ['romototo=romototo.__main__:main']
    },
    keywords='romototo housing',
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License"
    ]
)
