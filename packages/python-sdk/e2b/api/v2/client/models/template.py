# coding: utf-8

"""
    Devbook

    Devbook API

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import json
import pprint
import re  # noqa: F401

from aenum import Enum, no_arg


class Template(str, Enum):
    """
    Template
    """

    """
    allowed enum values
    """
    NODEJS = "Nodejs"
    GO = "Go"
    BASH = "Bash"
    RUST = "Rust"
    PYTHON3 = "Python3"
    PHP = "PHP"
    JAVA = "Java"
    PERL = "Perl"
    DOTNET = "DotNET"

    @classmethod
    def from_json(cls, json_str: str) -> Template:
        """Create an instance of Template from a JSON string"""
        return Template(json.loads(json_str))
