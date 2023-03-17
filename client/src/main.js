import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import Button from "primevue/button"
//main
import '@/assets/main.css';
//theme
import "primevue/resources/themes/lara-light-indigo/theme.css";
//core
import "primevue/resources/primevue.min.css";
//icons
import "primeicons/primeicons.css";



const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(PrimeVue, { ripple: true })
app.component('Button', Button);

app.mount('#app')
