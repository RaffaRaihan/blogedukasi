export default defineNuxtRouteMiddleware((to, from) => {
  // Ambil token dari cookies atau localStorage
  const token = useCookie('token') || useLocalStorage('token', null);

  if (!token.value) {
    // Jika token tidak ada, redirect ke halaman login
    return navigateTo('/login');
  }

  // Decode token jika perlu memverifikasi peran pengguna
  try {
    const base64Url = token.value.split('.')[1];
    const decoded = JSON.parse(atob(base64Url));

    // Jika route membutuhkan admin, pastikan peran adalah admin
    if (to.meta.requiresAdmin && decoded.role !== 'admin') {
      return navigateTo('/forbidden');
    }
  } catch (error) {
    // Jika token tidak valid, redirect ke login
    return navigateTo('/login');
  }
});
