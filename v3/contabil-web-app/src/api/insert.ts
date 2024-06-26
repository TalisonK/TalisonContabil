import axios from 'axios'

const url = `${process.env.REACT_APP_BACKEND_API}`

export const categoryList = async () => {
    const response = await axios.get(url + '/category')
    return response.data
}

export const submitActivity = async (data: any) => {
    let result = {}

    let type = data.type

    delete data.type

    if (type === 'Income') {
        result = await axios.post(url + '/income/create', data)
    } else if (data.list) {
        result = await axios.post(url + '/expense/list', data)
    } else {
        result = await axios.post(url + '/expense/create', data)
    }
    return result
}
