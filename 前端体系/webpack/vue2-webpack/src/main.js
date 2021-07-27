import App from "./app.vue";
import Vue from "vue";

new Vue({
  render: function(h) {
    return h(App);
  },
}).$mount("#app");
