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

export function useAuth(){
  async function fetchUser(){
    isLoading.value = true;
    try {
      const response = await fetch("/api/auth/user", {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        user.value = null;
        return false;
      }

      user.value = await response.json();
      return true
    } catch (error) {
      console.error("Error fetching user:", error);
      user.value = null;
      return false
    } finally {
      isLoading.value = false;
    }
  }

  return {
    user,
    isAuthenticated,
    isLoading,
    fetchUser
  }
}
