export default defineNuxtRouteMiddleware((to, from) => {
    const token = useCookie('token').value
  
    if (!token) {
      // Jika tidak ada token, redirect ke login
      return navigateTo('/login')
    }
  })
  