#! /usr/bin/python2.7
#
# How to run this python server
#   $ sudo apt-get install python-virtualenv
#   $ cd ~/
#   $ virtualenv wsgi-demo
#   $ cp wsgi_app.py ~/wsgi-demo
#   $ python wsgi_app.py
#
# Now you can access this server through curl:
#   $ curl http://localhost:8031 -v
#

from __future__ import print_function
from wsgiref.simple_server import make_server
from urlparse import parse_qs

def myapp(env, start_response):
    msg = 'No Message!'
    response_headers = [('content-type', 'text/plain')]

    start_response('200 OK', response_headers)
    qs_params = parse_qs(env.get('QUERY_STRING'))

    if 'msg' in qs_params:
            msg = qs_params.get('msg')[0]

    return ['Check the headers']

class Middleware:
    def __init__(self, app):
        self.wrapped_app = app

    def __call__(self, env, start_response):
        def custom_start_response(status, headers, exec_info=None):
            headers.append(('X-A-SIMPLE-TOKEN', "1234567890"))
            return start_response(status, headers, exec_info)
        return self.wrapped_app(env, custom_start_response)

app = Middleware(myapp)

httpd = make_server('', 8031, app)
print("Starting the server on port 8031")

httpd.serve_forever()
