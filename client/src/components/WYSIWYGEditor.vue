<template>
  <editor
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
</template>

<script>
import Editor from '@tinymce/tinymce-vue'

export default {
  name: 'WYSIWYGEditor',
  components: {
    'editor': Editor
  },
  setup() {
    const apiKey = import.meta.env.VITE_WYSIWYG_REDACTOR_API_KEY;

    if (!apiKey) {
      throw new Error('VITE_WYSIWYG_REDACTOR_API_KEY is missing!');
    }

    return {
      apiKey,
    }
  }
}
</script>
