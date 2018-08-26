import Vue from 'vue';
import App from '@/App';
import router from '@/router';

export default new Vue({
  el: '#app',
  router,
  render: (λ) => λ(App)
});
