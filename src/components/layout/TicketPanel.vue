<script setup lang="ts">
import type { Ticket } from '@/types/ticket.ts'
import { ticketData } from '@/data/tickets.ts'
import FilterItem from './FilterItem.vue'
import TicketListItem from './TicketListItem.vue'
import TicketDetail from './TicketDetail.vue'
import { computed, ref, watch } from 'vue'
import { matchesSearch } from '@/utils/search.ts'
import { Ghost, Search } from '@lucide/vue'
import router from '@/router/index.ts'

const props = defineProps<{
  activeUser: string
}>()

const tickets = ref(ticketData)

const filters = [
  {
    label: 'all-open',
    text: 'All Open',
  },
  {
    label: 'unassigned',
    text: 'Unassigned',
  },
  {
    label: 'mine',
    text: 'Mine',
  },
]
const activeFilter = ref('mine')
const searchQuery = ref('')
const selectedTicket = ref<Ticket | null>(null)

const filteredTickets = computed(() => {
  let result = tickets.value

  if (activeFilter.value === 'all-open') {
    result = result.filter((ticket) => ticket.status === 'Open')
  }

  if (activeFilter.value === 'unassigned') {
    result = result.filter((ticket) => ticket.status === 'Open' && ticket.assignee === null)
  }

  if (activeFilter.value === 'mine') {
    result = result.filter(
      (ticket) => ticket.status === 'Open' && ticket.assignee === props.activeUser,
    )
  }

  if (searchQuery.value.trim() !== '') {
    result = result.filter((ticket) => matchesSearch(ticket.title, searchQuery.value))
  }

  return result
})

watch(filteredTickets, (newTickets) => {
  if (selectedTicket.value !== null && !newTickets.includes(selectedTicket.value)) {
    selectedTicket.value = null
  }
})

function moveToCreate() {
  router.push('/create')
}
</script>

<template>
  <div class="h-full flex flex-row gap-4">
    <div class="flex-1 flex flex-col rounded-lg border shadow-md">
      <!-- Judul dan tombol new ticket -->
      <div class="h-12 px-2 flex flex-row py-2 justify-between bg-white rounded-lg">
        <p class="text-lg font-bold">Active queue</p>
        <button
          class="bg-blue-500 px-3 py-1 rounded-md text-white/80 text-center transition-colors duration-150 hover:text-white hover:bg-blue-700 cursor-pointer"
          @click="moveToCreate"
        >
          + New Ticket
        </button>
      </div>
      <!-- Filter: all open, unassigned, find by keyword, dan filter -->
      <div class="px-4 py-2 w-full flex flex-row gap-4 bg-white">
        <div class="flex flex-row items-center bg-gray-200 px-1 rounded-md border">
          <FilterItem
            v-for="item in filters"
            :key="item.label"
            :text="item.text"
            :active="activeFilter === item.label"
            @select="activeFilter = item.label"
          />
        </div>
        <div class="flex-1 flex flex-row items-center gap-2 rounded-sm border px-2">
          <Search class="w-4 h-4 text-gray-500" />
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search..."
            class="w-full outline-none"
          />
        </div>
        <button>Filters</button>
      </div>
      <!-- Tabel: judul, requester, assignee, category, status, priority -->
      <hr />
      <div
        class="px-4 py-1 grid grid-cols-[minmax(0,2fr)_repeat(2,minmax(0,1fr))] sm:grid-cols-[minmax(0,2fr)_repeat(4,minmax(0,1fr))] xl:grid-cols-[minmax(0,2fr)_repeat(5,minmax(0,1fr))] gap-4 bg-gray-100 border-b border-gray-500"
      >
        <div class="flex-1 font-bold">Ticket</div>
        <div class="hidden md:block lg:hidden text-center font-bold">Req.</div>
        <div class="hidden sm:block md:hidden lg:block text-center font-bold">Requester</div>
        <div class="hidden sm:block text-center font-bold">Assignee</div>
        <!-- <div class="hidden xl:block xl:block text-center font-bold">Cat.</div> -->
        <div class="hidden xl:block text-center font-bold">Category</div>
        <div class="text-center font-bold">Status</div>
        <div class="text-center font-bold">Priority</div>
      </div>
      <TicketListItem
        v-if="filteredTickets.length > 0"
        v-for="item in filteredTickets"
        :key="item.id"
        :ticket="item"
        :active="selectedTicket?.id === item.id"
        @select="selectedTicket = item"
      />
      <div v-else class="w-full h-full flex flex-col items-center justify-center gap-2 transition">
        <Ghost class="size-36 text-gray-500" />
        <p>No tickets found...</p>
      </div>
    </div>
    <div class="hidden md:block md:w-39 lg:w-68 xl:w-92" v-if="selectedTicket !== null">
      <TicketDetail :ticket="selectedTicket" />
    </div>
  </div>
</template>
