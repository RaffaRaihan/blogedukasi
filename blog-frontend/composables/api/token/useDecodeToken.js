import { useCookie } from 'nuxt/app'
import { ref } from 'vue'

const userData = ref(null)

export default function useDecodeToken() {
  const decodeToken = async () => {
    try {
      const jwt_decode = (await import('jwt-decode')).default
      const token = useCookie('token')?.value
      if (token) {
        const decoded = jwt_decode(token)
        userData.value = decoded
      } else {
        userData.value = null
      }
    } catch (error) {
      console.error('Gagal decode token:', error)
      userData.value = null
    }
  }

  return {
    decodeToken,
    userData
  }
}
