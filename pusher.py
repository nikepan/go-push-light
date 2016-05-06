import requests
import json
from django.conf import settings
from requests.exceptions import ConnectionError


def push_intent(intent, obj=None):
    if intent:
        data = {'intent': intent}
        if obj is not None:
            if type(obj) == str:
                data['obj'] = obj
            else:
                data['obj'] = json.dumps(obj)

        pusher_url = settings.PUSHER_URL

        try:
            resp = requests.post(pusher_url, data)
        except ConnectionError:
            return False
        if resp.status_code == 200:
            if resp.json()['status'] == 'ok':
                return True
        return False
    else:
        return False