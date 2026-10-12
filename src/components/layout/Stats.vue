<script setup lang="ts">
import { computed } from 'vue'
import StatItem from './StatItem.vue'
import { ticketData } from '@/data/tickets.ts'

const props = defineProps<{
  activeUser: string
}>()
// const statItems = [
//   {
//     label: 'open',
//     text: 'Open Tickets',
//     amount: 10,
//   },
//   {
//     label: 'unassigned',
//     text: 'Unassigned',
//     amount: 3,
//   },
//   {
//     label: 'mine',
//     text: 'Mine',
//     amount: 1,
//   },
// ]

const statItems = computed(() => {
  return [
    {
      label: 'all-open',
      text: 'Open Tickets',
      amount: ticketData.filter((ticket) => ticket.status === 'Open').length,
    },
    {
      label: 'unassigned',
      text: 'Unassigned',
      amount: ticketData.filter((ticket) => ticket.status === 'Open' && ticket.assignee === null)
        .length,
    },
    {
      label: 'mine',
      text: 'Mine',
      amount: ticketData.filter(
        (ticket) => ticket.status === 'Open' && ticket.assignee === props.activeUser,
      ).length,
    },
  ]
})
</script>

<template>
  <div class="w-full h-20 flex flex-row flex-start gap-3">
    <StatItem
      v-for="item in statItems"
      :key="item.label"
      :label="item.label"
      :text="item.text"
      :amount="item.amount"
    />
  </div>
</template>
