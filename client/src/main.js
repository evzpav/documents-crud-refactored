import Vue from 'vue'
import App from './App.vue'
import store from './store'

import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
Vue.use(Buefy);

import VueTheMask from 'vue-the-mask'
Vue.use(VueTheMask);

import VueGoodTablePlugin from 'vue-good-table';
import 'vue-good-table/dist/vue-good-table.css'
Vue.use(VueGoodTablePlugin);

import VModal from 'vue-js-modal'
Vue.use(VModal);

Vue.config.productionTip = false;
import "./assets/scss/styles.scss"

import VueRouter from 'vue-router'
Vue.use(VueRouter);
import routes from './routes';

const router = new VueRouter({
  routes,
  mode: 'history'
});

new Vue({
  render: h => h(App),
  store,
  router
}).$mount('#app');
