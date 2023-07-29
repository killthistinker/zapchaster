import api from "../requests"

export const getParts  = async () => await api.get('/carpart')

export const getProductData = async (product_data) => await api.get(`/detail?partId=${product_data}`)

export const counter = async (payload) => await api.post('counter/update',payload)

export default getParts