<template>
  <form @submit.prevent="signIn" class="form">
    <div class="form-name">Авторизация</div>
    <div class="p-input-icon-left" style="width: 100%">
      <i style="z-index: 111;" class="pi pi-at"/>
      <InputText v-model="formFields.email" type="email" autocomplete="on" required placeholder="Email"
                 style="width: 100%"/>
    </div>
    <div class="p-input-icon-left" style="width: 100%">
      <i style="z-index: 111;" class="pi pi-lock"/>
      <InputText v-model="formFields.password" required placeholder="Пароль" style="width: 100%"/>
    </div>
    <Button type="submit" label="Авторизоваться" style="width: 100%" severity="secondary" rounded/>
    <span class="dop-text">Нет аккаунта?<router-link class="dop-link" to="/registration"> Создайте</router-link></span>
  </form>
</template>

<script>
import InputText from "primevue/inputtext"
import {useAuthStore} from "@/stores/auth";
import {mapStores} from "pinia";

export default {
  name: 'login',
  components: {
    InputText
  },
  data() {
    return {
      formFields: {
        email: '',
        password: '',
      },
    }
  },
  methods: {
    async signIn() {
      try {
        await this.authStore.signIn(this.formFields);
        this.$router.push({path: '/'});
      } catch (error) {
        this.showError(error.response.data.msg);
      }
    },
    showError(message) {
      this.$toast.add({severity: 'error', summary: message, life: 10000});
    },
    showSuccess(message) {
      this.$toast.add({severity: 'success', summary: message, life: 10000});
    },
  },
  computed: {
    ...mapStores(useAuthStore)
  },
}
</script>

<style lang="scss" scoped>
.form {
  color: #383B41;

  max-width: 400px;
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  gap: 20px;

  .form-name {
    font-weight: 700;
    font-size: 28px;
    margin: 0 auto;
  }

  .dop-text {
    text-align: center;
    color: #8E8E8E;
    font-size: 20px;

    .dop-link {
      text-decoration: none;
      color: #FF7676;
    }
  }
}
</style>