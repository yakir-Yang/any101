"""Spawn multiple workers and collect their results.
 
Demonstrates how to use the eventlet.green.socket module.
"""
import os
import eventlet
from eventlet.green import socket
 
 
def geturl(url):
    c = socket.socket()
    ip = socket.gethostbyname(url)
    c.connect((ip, 80))
    print('%s connected' % url)
    c.sendall('GET /\r\n\r\n')
    return c.recv(1024)
 
 
urls = ['www.csdn.com', 'www.qq.com', 'www.kyang.cc']
pool = eventlet.GreenPool()
pile = eventlet.GreenPile(pool)

print "Start to spwan..."
for x in urls:
    pile.spawn(geturl, x)
print "End to spwan..."

pool.waitall()

# note that the pile acts as a collection of return values from the functions
# if any exceptions are raised by the function they'll get raised here
for url, result in zip(urls, pile):
    print('%s: %s' % (url, repr(result)[:80]))
