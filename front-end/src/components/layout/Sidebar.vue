<script setup lang="ts">
import { ref } from 'vue'
import SidebarNavItem from './SidebarNavItem.vue'
import { CircleUserRound } from '@lucide/vue'
import router from '@/router/index.ts'
const props = defineProps<{
  activeUser: string
}>()
const navigationItems = [
  {
    label: 'Dashboard',
    path: '/dashboard',
  },
  {
    label: 'Tickets',
    path: '/tickets',
  },
  {
    label: 'Accounts',
    path: '/accounts',
  },
]

function logout() {
  router.push('/login')
}

const showLogOut = ref(false)
</script>

<template>
  <div
    class="hidden md:flex md:w-40 lg:w-50 xl:w-60 min-h-screen flex-col justify-between p-4 bg-blue-900 gap-8"
  >
    <h1 class="text-2xl text-white">IT Helpdesk</h1>
    <div class="flex-1">
      <SidebarNavItem
        v-for="item in navigationItems"
        :key="item.path"
        :label="item.label"
        :path="item.path"
      />
    </div>
    <div class="relative">
      <div
        class="flex flex-col gap-2 border border-white rounded-lg cursor-pointer hover:bg-blue-800"
      >
        <div class="flex flex-row items-center px-4 py-2 gap-2" @click="showLogOut = !showLogOut">
          <CircleUserRound class="size-10 text-white" />
          <div class="flex flex-col">
            <div class="text-white">{{ activeUser }}</div>
            <div class="text-sm text-white/70">IT Support</div>
          </div>
        </div>
        <hr class="w-full border-white" v-if="showLogOut" />
        <button
          class="bg-blue-500 mb-2 text-black text-white font-bold text-center rounded-md mx-2 px-4 py-1 cursor-pointer transition-colors duration-150 hover:bg-blue-600"
          v-if="showLogOut"
          @click="logout"
        >
          Log Out
        </button>
      </div>
    </div>
  </div>
</template>
