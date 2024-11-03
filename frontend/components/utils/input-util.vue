<template>
  <input
      @input="onInput"
      class="transition-all border-2 border-input hover:opacity-80 focus:opacity-70 focus:border-emerald-400 duration-150 w-full bg-input drop-shadow-2xl px-5 py-2.5 rounded-xl disabled:cursor-not-allowed"
  />
</template>

<script setup lang="ts">
const emit = defineEmits(['update:modelValue'])

const onInput = (e: Event) => {
  emit('update:modelValue', (e.target as HTMLInputElement).value)
}
</script>