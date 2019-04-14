#! /usr/bin/python2.7
#
# How to run this python server
#   $ sudo apt-get install python-virtualenv
#   $ cd ~/
#   $ virtualenv pastedeploy-demo
#   $ cp pastedeploylab.py ~/pastedeploy-demo
#   $ touch pastedeploylab.ini
#   $ python pastedeploylab.py
#
# Now you can access this server through curl:
#   $ curl http://localhost:8080 -v
#   $ curl http://127.0.0.1:8080/calc?operator=plus&operand1=12&operand2=23 -v

'''ATTENTION: the INI configure file is at the bottom of source file'''

import os
import webob
from webob import Request
from webob import Response
from paste.deploy import loadapp
from wsgiref.simple_server import make_server

#Filter
class LogFilter():
    def __init__(self, app):
        self.app = app
        pass

    def __call__(self, environ, start_response):
        print "filter:LogFilter is called."
        return self.app(environ, start_response)

    @classmethod
    def factory(cls, global_conf, **kwargs):
        print "----------- LogFilter -------------"
        print "in LogFilter.factory", global_conf, kwargs
        return LogFilter

class ShowVersion():
    def __init__(self):
        pass

    def __call__(self, environ, start_response):
        start_response("200 OK", [("Content-type", "text/plain")])
        return ["Paste Deploy LAB: Version = 1.0.0",]

    @classmethod
    def factory(cls, global_conf, **kwargs):
        print "----------- ShowVersion -------------"
        print "in ShowVersion.factory", global_conf, kwargs
        return ShowVersion()

class Calculator():
    def __init__(self):
        pass

    def __call__(self, environ, start_response):
        req = Request(environ)
        res = Response()
        res.status = "200 OK"
        res.content_type = "text/plain"

        # get operands
        operator = req.GET.get("operator", None)
        operand1 = req.GET.get("operand1", None)
        operand2 = req.GET.get("operand2", None)

        #print req.GET
        opnd1 = int(operand1)
        opnd2 = int(operand2)

        if operator == u'plus':
            opnd1 = opnd1 + opnd2
        elif operator == u'minus':
            opnd1 = opnd1 - opnd2
        elif operator == u'star':
            opnd1 = opnd1 * opnd2
        elif operator == u'slash':
            opnd1 = opnd1 / opnd2

        res.body = "%s \nRESULT= %d" % (str(req.GET) , opnd1)
        return res(environ,start_response)

    @classmethod
    def factory(cls, global_conf, **kwargs):
        print "----------- Calculator -------------"
        print "in Calculator.factory", global_conf, kwargs
        return Calculator()

if __name__ == '__main__':
    configfile = "pastedeploylab.ini"
    appname = "pdl"
    wsgi_app = loadapp("config:%s" % os.path.abspath(configfile), appname)
    server = make_server('localhost', 8080, wsgi_app)
    server.serve_forever()
    pass


# Here is the INI configure file 'pastedeploylab.py'
'''
[DEFAULT]
global-key1=yang
global-key2=kk

[composite:pdl]
use=egg:Paste#urlmap
/:root
/calc:calc

[pipeline:root]
pipeline = logrequest showversion

[pipeline:calc]
pipeline = logrequest calculator

[filter:logrequest]
username = root
password = root123
paste.filter_factory = pastedeploylab:LogFilter.factory

[app:showversion]
version = 1.0.0
paste.app_factory = pastedeploylab:ShowVersion.factory

[app:calculator]
description = This is an "+-*/" Calculator
paste.app_factory = pastedeploylab:Calculator.factory
'''
