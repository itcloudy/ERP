import axiosAjax from 'utils/axiosAjax';

export const axiosAjaxLogin = (type, url, data, callback, errback, others) => { return axiosAjax(type, url, data, callback, errback, others) };