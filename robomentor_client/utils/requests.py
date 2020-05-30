# -*- coding: utf-8 -*-
"""
RoboMentor_Client: Python library and framework for RoboMentor_Client.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
:copyright: (c) 2020 by RoboMentor.
:license: MIT, see LICENSE for more details.
"""

import requests
from .log import Log


class HttpRequest:

    @staticmethod
    def do(self, url, params, http_method):
        res = ''
        if http_method.upper() == 'POST':
            try:
                res = requests.post(url, params)
                Log.info("正在进行post请求")
            except Exception as e:
                Log.error("post请求出现了异常：{0}".format(e))
        elif http_method.upper() == 'GET':
            try:
                res = requests.post(url, params)
                Log.info("正在进行get请求")
            except Exception as e:
                Log.error("get请求出现了异常：{0}".format(e))
        return res
