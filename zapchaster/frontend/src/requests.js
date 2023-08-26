import axios from 'axios'

const api = axios.create({
  baseURL: 'https://zapchastik.kz/backend/'
})


export default api
