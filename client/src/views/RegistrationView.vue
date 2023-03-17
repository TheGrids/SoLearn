<template>
    <div class="form">
        <div class="form-name">Регистрация</div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-at" />
            <InputText v-model="value1" placeholder="Email" style="width: 100%"/>
        </div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-lock" />
            <InputText v-model="value1" placeholder="Пароль"  style="width: 100%" />
        </div>
        <div class="p-input-icon-left" style="width: 100%">
            <i style="z-index: 111;" class="pi pi-lock" />
            <InputText v-model="value1" placeholder="Повтор пароль"  style="width: 100%" />
        </div>
        <Button label="Зарегистрироваться" style="width: 100%" severity="secondary" rounded />
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
}
</style>