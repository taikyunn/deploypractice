import Vue from 'vue'
import Router from 'vue-router'
import Todo from './components/Todo.vue'
import Edit from './components/Edit.vue'

Vue.use(Router)

export default new Router ({
  mode: "history",
  routes: [
    {path: "/", component: Todo},
    {path: "/edit/:id", component: Edit, name: 'edit', props: true},
    {path: "*",redirect: {path: "/"}}
  ]
})
