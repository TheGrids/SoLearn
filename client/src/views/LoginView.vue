<template>
    <div class="form">
        <div class="form-name">Авторизация</div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-at" />
            <InputText placeholder="Email" style="width: 100%"/>
        </div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-lock" />
            <InputText placeholder="Пароль"  style="width: 100%" />
        </div>
        <Button label="Авторизоваться" style="width: 100%" severity="secondary" rounded />
        <span class="dop-text">Нет аккаунта?<router-link class="dop-link" to="/registration"> Создате</router-link></span>
    </div>
</template>

<script>
import InputText from "primevue/inputtext"

export default {
    name: 'login',
    components: {
        InputText
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