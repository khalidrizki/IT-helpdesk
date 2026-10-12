<script setup lang="ts">
import { RouterView } from 'vue-router'
import Sidebar from '@/components/layout/Sidebar.vue'
import { ref } from 'vue'
const activeUser = ref('Will')
</script>

<template>
  <div class="min-h-screen flex flex-row">
    <Sidebar :active-user="activeUser" />
    <main class="flex-1 min-h-screen">
      <RouterView v-slot="{ Component }"
        ><component :is="Component" :active-user="activeUser"
      /></RouterView>
    </main>
  </div>
</template>
