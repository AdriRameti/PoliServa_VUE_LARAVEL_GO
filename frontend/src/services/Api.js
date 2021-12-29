import axios from 'axios'

export default (URL) => {
    
    const InstanceAxios = axios.create({
        baseURL: URL
    })

    return InstanceAxios
}