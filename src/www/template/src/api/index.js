import axios from 'axios'
import { Message } from 'element-ui'

let message

const HTTP = axios.create({
    baseURL: '/'
});

HTTP.interceptors.request.use(config => {
    config.headers = {
      'Content-Type':'application/json'
    }
    return config
}, (error) => {
    return Promise.reject(error)
})

// response拦截器
HTTP.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        if (error && error.response) {
            switch (error.response.status) {
                case 400:
                    break
                case 401:
                    break
                case 403:
                    break
                case 404:
                    break
                case 408:
                    break
                case 500:
                    break
                case 504:
                    break
                default:
                    error.message = '未知错误'
            }
            message = Message.error(error.message)
        } else {
            message = Message.error(error.message)
        }
        return {error: -1, data: ''}
    }
)

export function Check() {
    return HTTP({
        url: 'common/ping',
        method: 'GET',
        data: '',
    });
}

export function GetHomeIndex() {
    return HTTP({
        url: 'common/home/index',
        method: 'GET',
        data: '',
    });
}

export function GetHomeTools() {
    return HTTP({
        url: 'common/home/tools',
        method: 'GET',
        data: '',
    });
}

export function SetHomeToolsSerial(data) {
    return HTTP({
        url: 'common/home/tools/serial',
        method: 'POST',
        data: data,
    });
}
