<template>
  <div class="card">
    <div class=" row" style="align-items: center; justify-content: space-between">
      <div class="row">
        <div style="font-size: 20px; margin-right: 10px">{{ stepNumber }}.</div>
        <InputText placeholder="Название темы" v-model="themeValue.title"/>
      </div>
      <div>
        <Button @click="addSubTheme" v-tooltip.bottom="'Добавить пункт'" link placeholder="Название"
                v-model="themeValue.title"
                icon="pi pi-plus"/>
        <Button v-tooltip.bottom="'Убрать тему'" @click="$emit('remove', themeValue.id)" link placeholder="Название"
                v-model="themeValue.title"
                icon="pi pi-trash"/>
      </div>
    </div>
    <div class="p-4">
      <div v-for="subTheme in themeValue.subThemes">
        <div class="mt-5 mb-5" style="display: flex; flex-direction: row; justify-content: space-between">
          <p class="text-lg">Основная часть</p>
          <Button @click="editSubTheme(subTheme)" v-tooltip.bottom="'Изменить пункт'" link placeholder="Название"
                  v-model="themeValue.title"
                  icon="pi pi-pencil"/>
        </div>
        <p style="word-wrap: anywhere">
          {{ subTheme.name }}
        </p>
        <p style="word-wrap: anywhere" v-html="subTheme.content"/>
      </div>
    </div>
    <Dialog v-model:visible="dialogVisible" maximizable header="Содержание темы"
            :style="{ width: '80vw' }">
      <div class="mb-5 mt-1" style="z-index: 999999">
        <InputText v-model="currentSubTheme.name" placeholder="Название"/>
      </div>
      <WYSIWYGEditor v-model="currentSubTheme.content"/>
    </Dialog>
  </div>
</template>

<script>
import {computed, ref} from "vue";
import Dialog from "primevue/dialog";
import WYSIWYGEditor from "@/components/WYSIWYGEditor.vue";

export default {
  name: "CourseStep",
  components: {
    WYSIWYGEditor,
    Dialog,
  },
  emits: [
    'update:theme',
    'remove'
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
      subThemes: {
        type: Array,
        default: [],
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
    const dialogVisible = ref(false);
    const currentSubTheme = ref()

    const editSubTheme = (subTheme) => {
      currentSubTheme.value = subTheme;
      dialogVisible.value = true;
    }

    const addSubTheme = () => {
      dialogVisible.value = true;
      const subtheme = {
        name: '',
        content: '',
      }
      themeValue.value.subThemes.push(subtheme)
      currentSubTheme.value = subtheme;
    }

    return {
      themeValue,
      addSubTheme,
      dialogVisible,
      currentSubTheme,
      editSubTheme,
    }
  },
}
</script>

<style>

.card {
  width: 100%;
  background: #ffffff;
  border-radius: 10px;
  padding: 15px;
  margin-bottom: 10px;
}

.p-dialog-mask {
  z-index: 0 !important;
}
</style>