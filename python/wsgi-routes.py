#!/usr/bin/python
# -*- coding: UTF-8 -*-

import webob.dec
import eventlet
from eventlet import wsgi, listen
from routes import Mapper, middleware
 
class controller(object):
    def index(self):
        return "do index()"
 
    def add(self):
        return "do show()"
 
class App(object):
    def __init__(self):
        self.controller = controller()
        m = Mapper()
        m.connect('blog', '/blog/{action}/{id}', controller=controller,
                  conditions={'method': ['GET']})
        self.router = middleware.RoutesMiddleware(self.dispatch, m)
 
    @webob.dec.wsgify
    def dispatch(self, req):
        # RoutesMiddleware 会根据接收到的 url, 自动调用 map.match(),
        # 做路由匹配, 然后调用第一个参数, 这里即是 self.dispatch()
        print "kyang ----> req.environ: ", req.environ
        match = req.environ['wsgiorg.routing_args'][1]
        if not match:
            return 'error url: %s' % req.environ['PATH_INFO']

        print "kyang ----> match: ", match
 
        # 根据用户的请求动作, 调用相应的 action 处理函数(index/add)
        action = match['action']
        if hasattr(self.controller, action):
            func = getattr(self.controller, action)
            ret = func()
            return ret
        else:
            return "has no action:%s" % action
 
    @webob.dec.wsgify
    def __call__(self, req):
        return self.router
 
if __name__ == '__main__':
    socket = listen(('0.0.0.0', 8000))
    server = eventlet.spawn(wsgi.server, socket, App())
    server.wait()
