<script setup lang="ts">
import type { Ticket } from '@/types/ticket'
import { CircleUserRound } from '@lucide/vue'
const props = defineProps<{
  ticket: Ticket
}>()
</script>

<template>
  <div class="flex flex-col gap-2 w-full py-1 rounded-lg border shadow-md">
    <div class="flex flex-col gap-1 px-2">
      <p class="my-4 text-xl text-black font-bold">{{ props.ticket.title }}</p>
      <div class="flex flex-row gap-2">
        <div class="px-2 rounded-xl border text-xs">{{ props.ticket.status }}</div>
        <div class="px-2 rounded-xl border text-xs">{{ props.ticket.priority }}</div>
      </div>
      <p class="text-black text-sm">{{ props.ticket.description }}</p>
    </div>
    <hr />
    <div class="flex flex-row items-center">
      <CircleUserRound class="hidden md:block size-10 md:pl-2" />
      <div class="flex flex-col md:px-2">
        <p class="text-sm">{{ props.ticket.requester }}</p>
        <p class="hidden lg:block text-xs">{{ props.ticket.requesterEmail }}</p>
      </div>
    </div>
    <hr />
    <div class="flex flex-col px-2 gap-2">
      <p class="text-sm">Properties</p>
      <div class="px-2">
        <div
          class="grid md:grid-cols-[minmax(0,1fr)_minmax(0,1fr)] lg:grid-cols-[minmax(0,1fr)_minmax(0,2fr)]"
        >
          <p class="text-xs text-gray-700">Assignee</p>
          <p class="text-xs text-black">{{ props.ticket.assignee ?? '< Unassigned >' }}</p>
        </div>
        <div
          class="grid md:grid-cols-[minmax(0,1fr)_minmax(0,1fr)] lg:grid-cols-[minmax(0,1fr)_minmax(0,2fr)]"
        >
          <p class="text-xs text-gray-700">Category</p>
          <p class="text-xs text-black">{{ props.ticket.category }}</p>
        </div>
        <div
          v-if="props.ticket.affectedAsset"
          class="grid md:grid-cols-[minmax(0,1fr)_minmax(0,1fr)] lg:grid-cols-[minmax(0,1fr)_minmax(0,2fr)]"
        >
          <p class="text-xs text-gray-700">Affected Asset</p>
          <p class="text-xs text-black self-center">{{ props.ticket.affectedAsset }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
