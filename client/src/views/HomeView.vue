<template>
  <main>
    <WYSIWYGEditor />
    <Button label="OK"/>
  </main>
</template>

<script>
import WYSIWYGEditor from "@/components/WYSIWYGEditor.vue";
import Button from "primevue/button"
export default {
  name: 'app',
  components: {
    WYSIWYGEditor,
  },
}
</script>