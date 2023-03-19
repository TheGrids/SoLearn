import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import Button from "primevue/button"
import InputText from "primevue/inputtext";
import InputNumber from "primevue/inputnumber";
//main
import '@/assets/main.css';
//theme
import "primevue/resources/themes/lara-light-indigo/theme.css";
//core
import "primevue/resources/primevue.min.css";
import Tooltip from 'primevue/tooltip';

//icons
import "primeicons/primeicons.css";

import ToastService from 'primevue/toastservice';
import {useAuthStore} from "@/stores/auth";


const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(PrimeVue, {ripple: true})
app.use(ToastService)
app.component('Button', Button)
app.component('InputText', InputText)
app.component('InputNumber', InputNumber)
app.directive('tooltip', Tooltip);

const authStore = useAuthStore();
authStore.checkAuth();

app.mount('#app')
