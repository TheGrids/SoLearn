<template>
    <div class="form">
        <div class="form-name">Регистрация</div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-at" />
            <InputText placeholder="Email" style="width: 100%"/>
        </div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-lock" />
            <InputText placeholder="Пароль"  style="width: 100%" />
        </div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-lock" />
            <InputText placeholder="Повтор пароль"  style="width: 100%" />
        </div>
        <Button label="Зарегистрироваться" style="width: 100%" severity="secondary" rounded />
        <span class="dop-text">Есть аккаунт?<router-link class="dop-link" to="/login"> Войдите</router-link></span>
    </div>
</template>

<script>
import InputText from "primevue/inputtext"

export default {
    name: 'registration',
    components: {
        InputText,
    }
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