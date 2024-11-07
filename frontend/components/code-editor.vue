<template>
  <codemirror
      class="jetbrains"
      v-model="code"
      placeholder="Code goes here..."
      :autofocus="true"
      :indent-with-tab="true"
      :tab-size="2"
      :extensions="extensions"
      @ready="handleReady"
      @change="log('change', $event)"
      @focus="log('focus', $event)"
      @blur="log('blur', $event)"
  />
</template>

<script>
import { defineComponent, ref, shallowRef } from 'vue'
import { Codemirror } from 'vue-codemirror'
import { java } from '@codemirror/lang-java'
import { oneDark } from '@codemirror/theme-one-dark'

export default defineComponent({
  props: {
    language: String,
    content: String,
  },
  components: {
    Codemirror
  },
  setup() {
    const code = ref('')
    const extensions = [java(), oneDark]

    // Codemirror EditorView instance ref
    const view = shallowRef()
    const handleReady = (payload) => {
      view.value = payload.view
    }

    return {
      code,
      extensions,
      handleReady,
      log: console.log
    }
  },
  mounted() {
    this.code = this.content
  }
})
</script>