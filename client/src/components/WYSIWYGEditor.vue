<template>
  <div style="height: 100%; z-index: 9999 !important">
    <editor
        v-model="content"
        :api-key="apiKey"
        :init="{
         height: 500,
         menubar: true,
         plugins: [
           'media',
           'advlist autolink lists link image charmap print preview anchor',
           'searchreplace visualblocks code fullscreen',
           'insertdatetime media table paste code help wordcount'
         ],
         toolbar:
           'undo redo | formatselect | bold italic backcolor | \
           alignleft aligncenter alignright alignjustify | \
           bullist numlist outdent indent | removeformat | help'
       }"
    />
    <div class="mt-5 mb-5" style="display: flex; flex-direction: row; justify-content: flex-end">
      <Button>Сохранить</Button>
    </div>
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import {computed} from "vue";

export default {
  name: 'WYSIWYGEditor',
  props: {
    modelValue: {
      type: String,
      required: true,
    }
  },
  emits: [
    'update:modelValue'
  ],
  components: {
    'editor': Editor
  },
  setup(props, {emit}) {
    const apiKey = import.meta.env.VITE_WYSIWYG_REDACTOR_API_KEY;
    const content = computed({
      get: () => props.modelValue,
      set: (value) => emit('update:modelValue', value),
    });

    if (!apiKey) {
      throw new Error('VITE_WYSIWYG_REDACTOR_API_KEY is missing!');
    }

    return {
      content,
      apiKey,
    }
  },
}
</script>
