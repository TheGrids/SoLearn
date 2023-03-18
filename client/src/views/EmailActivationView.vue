<template>
  <div class="flex center" style="height: 100vh; align-items: center">
    <ProgressSpinner style="width: 100px; height: 100px" strokeWidth="8" fill="var(--surface-ground)"
                     animationDuration=".5s" aria-label="Custom ProgressSpinner"/>
  </div>
</template>

<script>
import ProgressSpinner from "primevue/progressspinner";
import {api} from "@/http";

export default {
  name: "EmailActivation",
  components: {
    ProgressSpinner,
  },
  mounted() {
    this.activateEmail();
  },
  methods: {
    async activateEmail() {
      try {
        const id = this.$route.params.id;
        const response = await api.get(`activate/${id}`);
        this.$router.replace({name: 'login'});
        this.showSuccess(response.data.msg);
      } catch (error) {
        this.$router.replace({name: 'login'});
        const message = error.response.data.msg;
        this.showError(message);
      }
    },
    showError(message) {
      this.$toast.add({severity: 'error', summary: message, life: 10000});
    },
    showSuccess(message) {
      this.$toast.add({severity: 'success', summary: message, life: 10000});
    },
  }
}
</script>

<style scoped>

</style>