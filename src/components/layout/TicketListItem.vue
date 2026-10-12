<script setup lang="ts">
import type { Ticket } from '@/types/ticket'
import { UserRoundMinus } from '@lucide/vue'
const props = defineProps<{
  ticket: Ticket
  active: boolean
}>()

const emit = defineEmits<{
  select: []
}>()
</script>

<template>
  <div
    @click="emit('select')"
    class="px-4 py-1 grid grid-cols-[minmax(0,2fr)_repeat(2,minmax(0,1fr))] sm:grid-cols-[minmax(0,2fr)_repeat(4,minmax(0,1fr))] xl:grid-cols-[minmax(0,2fr)_repeat(5,minmax(0,1fr))] items-center gap-4 border-b border-gray-500 transition-colors duration-150 cursor-pointer"
    :class="[props.active ? 'bg-yellow-100 hover:bg-yellow-200' : 'bg-white hover:bg-gray-100']"
  >
    <div class="flex-1 font-normal">{{ props.ticket.title }}</div>
    <div class="hidden sm:block text-center font-normal">{{ props.ticket.requester }}</div>
    <div class="hidden sm:block text-center font-normal">
      <template v-if="props.ticket.assignee === null">
        <span
          class="inline-flex items-center justify-center rounded-sm border bg-yellow-200 px-1 lg:hidden"
        >
          <UserRoundMinus class="size-4" />
        </span>
        <span
          class="hidden lg:inline-block rounded-sm border bg-yellow-200 px-1 xl:text-sm 2xl:text-base"
        >
          Unassigned
        </span>
      </template>
      <span v-else>{{ props.ticket.assignee }}</span>
    </div>
    <div class="hidden xl:block text-center font-normal">{{ props.ticket.category }}</div>
    <div class="text-center font-normal">{{ props.ticket.status }}</div>
    <div class="text-center font-normal">{{ props.ticket.priority }}</div>
  </div>
</template>
