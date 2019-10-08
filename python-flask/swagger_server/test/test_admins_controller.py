# coding: utf-8

from __future__ import absolute_import

from flask import json
from six import BytesIO

from swagger_server.test import BaseTestCase


class TestAdminsController(BaseTestCase):
    """AdminsController integration test stubs"""

    def test_load_item(self):
        """Test case for load_item

        load item(csv) to the collection
        """
        query_string = [('item', 'item_example')]
        response = self.client.open(
            '/Nuno19/Recomendation_Service/1.0.0/loadItem',
            method='POST',
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_load_item_list(self):
        """Test case for load_item_list

        load list of items(csv) to the collection
        """
        query_string = [('itemList', 'itemList_example')]
        response = self.client.open(
            '/Nuno19/Recomendation_Service/1.0.0/loadItemList',
            method='POST',
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_set_cluster_count(self):
        """Test case for set_cluster_count

        set number of clusters
        """
        query_string = [('itemList', 56)]
        response = self.client.open(
            '/Nuno19/Recomendation_Service/1.0.0/setClusterNumber',
            method='POST',
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    import unittest
    unittest.main()
