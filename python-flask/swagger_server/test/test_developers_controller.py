# coding: utf-8

from __future__ import absolute_import

from flask import json
from six import BytesIO

from swagger_server.models.item import Item  # noqa: E501
from swagger_server.test import BaseTestCase


class TestDevelopersController(BaseTestCase):
    """DevelopersController integration test stubs"""

    def test_create_user(self):
        """Test case for create_user

        createUser
        """
        query_string = [('extID', 'extID_example')]
        response = self.client.open(
            '/Nuno19/Recomendation_Service/1.0.0/createUser',
            method='GET',
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_get_recommended(self):
        """Test case for get_recommended

        searches recommended
        """
        query_string = [('maxCount', 56),
                        ('movieId', 1)]
        response = self.client.open(
            '/Nuno19/Recomendation_Service/1.0.0/getRecommended',
            method='GET',
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    import unittest
    unittest.main()
