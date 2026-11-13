import { ref, computed } from "vue";

interface User {
  id: string;
  name: string;
  email: string;
  department: string;
  role: string;
}

const user = ref<User | null>(null);
const isAuthenticated = computed(() => user.value !== null);
const isLoading = ref(false);
const isInitialized = ref(false);

export function useAuth(){
  async function fetchUser(force = false){
    if (isInitialized.value && !force){
      return isAuthenticated.value;
    }

    isLoading.value = true;

    try {
      const response = await fetch("/auth/me", {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        user.value = null;
        isInitialized.value = true
        return false;
      }

      user.value = await response.json();
      isInitialized.value = true;

      return true
    } catch (error) {
      console.error("Error fetching user:", error);
      return false
    } finally {
      isLoading.value = false;
    }
  }

  async function logout() {
  const response = await fetch('/auth/logout', {
    method: 'POST',
    credentials: 'include',
  });

  if (!response.ok) {
    throw new Error('Failed to log out');
  }

  user.value = null;
  isInitialized.value = true;
}

  return {
    user,
    isAuthenticated,
    isLoading,
    isInitialized,
    fetchUser,
    logout
  }
}
