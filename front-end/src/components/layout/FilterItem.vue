<script setup lang="ts">
const props = defineProps<{
  text: string
  active: boolean
}>()

const emit = defineEmits<{
  select: []
}>()
</script>

<template>
  <button
    @click="emit('select')"
    class="h-5 w-20 rounded-sm text-xs transition-colors duration:150 hover:bg-gray-100 cursor-pointer"
    :class="props.active ? 'bg-white text-black' : 'bg-gray-200 text-gray-800'"
  >
    {{ props.text }}
  </button>
</template>
