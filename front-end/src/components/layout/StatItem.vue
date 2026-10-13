<script setup lang="ts">
const props = defineProps<{
  label: string
  text: string
  amount: number
}>()
</script>

<template>
  <div class="w-36 flex flex-col px-4 py-2 rounded-lg border shadow-md">
    <p class="self-center text-lg">
      {{ props.text }}
    </p>
    <hr />
    <p class="flex-1 self-center text-lg font-bold text-orange-400">
      {{ props.amount }}
    </p>
  </div>
</template>
