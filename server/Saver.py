import redis
import google.protobuf.json_format as js

conn = redis.Redis('localhost', port=6379, db=0)


def SaveUser(user):
    messageJson = (js.MessageToJson(user))
    conn.set(user.ip_address, messageJson)
    response = user.user_name + " kayıt edildi."
    return response
