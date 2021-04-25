import {useState, useEffect} from 'react'
import axios from "axios";

const useAxios = url => {
    const [name, setName] = useState('')
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        const instance = axios.create({
            withCredentials: true,
        })

        const axiosResource = async () => {
            try {
                await instance.get(url)
                    .then(res => {
                        setName(res.data.username)
                    })

                setLoading(false)
            } catch (error) {
                setLoading(false)
                setError(error)
            }
        }
        axiosResource()
    }, [url])

    return { name, setName, loading, error }
}

export default useAxios