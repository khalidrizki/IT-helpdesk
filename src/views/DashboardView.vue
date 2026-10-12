<script setup lang="ts">
import Stats from '@/components/layout/Stats.vue'
import TicketPanel from '@/components/layout/TicketPanel.vue'
import { ref } from 'vue'

const props = defineProps<{
  activeUser: string
}>()
</script>

<template>
  <div class="h-screen px-1 sm:px-2 md:px-4 lg:px-8 py-4 flex flex-col gap-4">
    <Stats :active-user="activeUser" />
    <TicketPanel :active-user="activeUser" class="flex-1 min-h-0" />
  </div>
</template>
