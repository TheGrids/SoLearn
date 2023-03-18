<template>
  <form @submit.prevent="signUp" class="form">
    <div class="form-name">Регистрация</div>
    <div class="p-input-icon-left" style="width: 100%">
      <i style="z-index: 111;" class="pi pi-at"/>
      <InputText type="email" required v-model="formFields.email" placeholder="Email" style="width: 100%"/>
    </div>
    <div style="width: 100%">
      <InputText
          v-model="formFields.firstName" @input.stop="validateFirstNameInput"
          required
          placeholder="Имя"
          style="width: 100%"/>
    </div>
    <div style="width: 100%">
      <InputText required v-model="formFields.lastName" @input.stop="validateLastNameInput" placeholder="Фамилия"
                 style="width: 100%"/>
    </div>

    <div class="p-input-icon-left" style="width: 100%">
      <i style="z-index: 111;" class="pi pi-lock"/>
      <InputText min="8" minlength="8" required v-model="formFields.password" placeholder="Пароль"
                 style="width: 100%"/>
    </div>
    <div class="p-input-icon-left" style="width: 100%">
      <i style="z-index: 111;" class="pi pi-lock"/>
      <InputText v-model="repeatPassword" min="8" minlength="8" :class="{'p-invalid': !repeatPasswordIsValid}"
                 required placeholder="Повтор пароль" style="width: 100%"
      />
    </div>
    <Button :loading="signUpIsLoading" type="submit" label="Зарегистрироваться" style="width: 100%" severity="secondary"
            rounded/>
    <span class="dop-text">Есть аккаунт?<router-link class="dop-link" to="/login"> Войдите</router-link></span>
  </form>
</template>

<script>
import InputText from "primevue/inputtext"
import AuthService from "@/services/auth.service";
import {computed, ref} from "vue";
import {useAuthStore} from "@/stores/auth";
import {mapStores} from "pinia";

export default {
  name: 'registration',
  components: {
    InputText,
  },
  setup() {
    const formFields = ref({
      email: '',
      password: '',
      firstName: '',
      lastName: '',
    });
    const signUpIsLoading = ref(false);
    const repeatPassword = ref('');
    const validateFirstNameInput = () => {
      formFields.value.firstName = formFields.value.firstName.replace(/[^a-zA-Z]/g, '');
    }

    const validateLastNameInput = () => {
      formFields.value.lastName = formFields.value.lastName.replace(/[^a-zA-Z]/g, '');
    }

    const repeatPasswordIsValid = computed(() => {
      if (formFields.value.password && repeatPassword.value) {
        return formFields.value.password === repeatPassword.value;
      }
      return true;
    });

    return {
      formFields,
      validateFirstNameInput,
      validateLastNameInput,
      repeatPassword,
      repeatPasswordIsValid,
      signUpIsLoading,
    }
  },
  methods: {
    showError(message) {
      this.$toast.add({severity: 'error', summary: message, life: 10000});
    },
    showSuccess(message) {
      this.$toast.add({severity: 'success', summary: message, life: 10000});
    },
    async signUp() {
      if (this.repeatPasswordIsValid) {
        this.signUpIsLoading = true;
        try {
          const response = await this.authStore.signUp(this.formFields);
          this.showSuccess(response.data.msg);
          this.$router.push({path: '/'})
        } catch (error) {
          this.showError(error.response.data.msg);
        } finally {
          this.signUpIsLoading = false;
        }
      }
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