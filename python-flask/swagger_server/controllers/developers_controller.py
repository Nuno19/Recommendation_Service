import connexion
import six

from swagger_server.models.item import Item  # noqa: E501
from swagger_server import util


def create_user(extID):  # noqa: E501
    """createUser

    By passing in the appropriate options, you can search for available cinemas  # noqa: E501

    :param extID: 
    :type extID: str

    :rtype: None
    """
    return 'do some magic!'


def get_recommended(maxCount=None, movieId=None):  # noqa: E501
    """searches recommended

    By passing in the appropriate options, you can search for available cinemas  # noqa: E501

    :param maxCount: 
    :type maxCount: int
    :param movieId: 
    :type movieId: int

    :rtype: List[Item]
    """
    return 'do some magic!'
