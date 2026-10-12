<script setup lang="ts">
import router from '@/router'

function goHome() {
  router.push('/dashboard')
}
</script>
<template>
  <div class="h-screen flex flex-col justify-center items-center gap-3 -translate-y-6">
    <p class="text-9xl">404</p>
    <p class="text-4xl">Page not found</p>
    <button
      class="text-2xl font-bold mt-4 border-b-2 transition-colors duration-150 cursor-pointer hover:text-blue-400"
      @click="goHome"
    >
      Go back home
    </button>
  </div>
</template>
