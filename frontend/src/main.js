import 'core-js/stable';
import 'regenerator-runtime/runtime';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import 'bootstrap/dist/css/bootstrap.css';
import Vue from 'vue';
import App from './App.vue';
import BootstrapVue from 'bootstrap-vue'
import router from "./router";

Vue.use(BootstrapVue);
Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
import vuetify from './plugins/vuetify';

Wails.Init(() => {
	new Vue({
        router,
        vuetify,
        render: h => h(App)
    }).$mount('#app');
});
