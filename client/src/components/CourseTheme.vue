<template>
  <div class="card row" style="align-items: center;">
    <div style="font-size: 20px; margin-right: 10px">{{ stepNumber }}.</div>
    <InputText placeholder="Название" v-model="themeValue.title"/>
    <Button placeholder="Название" v-model="themeValue.title">
    </Button>
  </div>
</template>

<script>
import { computed } from "vue";

export default {
  name: "CourseStep",
  emits: [
    'update:theme'
  ],
  props: {
    theme: {
      type: Object,
      required: true,
      id: {
        type: String,
        required: true,
      },
      title: {
        type: String,
        required: true,
      },
    },
    stepNumber: {
      type: Number,
      required: true,
    }
  },
  setup(props, {emit}) {
    const themeValue = computed({
      get: () => props.theme,
      set: () => emit('update:theme')
    });

    return {
      themeValue,
    }
  }
}
</script>

<style scoped>

.card {
  width: 100%;
  background: #ffffff;
  border-radius: 10px;
  padding: 15px;
  margin-bottom: 10px;
}
</style>