import requests
import json
from django.conf import settings


def push_intent(intent, obj=None):
    if intent:
        data = {'intent': intent}
        if obj is not None:
            if type(obj) == str:
                data['obj'] = obj
            else:
                data['obj'] = json.dumps(obj)

        pusher_url = settings.PUSHER_URL

        resp = requests.post(pusher_url, data)
        if resp.status_code == 200:
            if resp.json()['status'] == 'ok':
                return True
        return False
    else:
        return False