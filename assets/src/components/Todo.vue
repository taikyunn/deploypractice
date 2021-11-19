<template>
  <div>
    <h1>Todoリスト</h1>
    <label for="content">今日のContent:</label>
    <input type="text" name="content" v-model="content">
    <p v-if="isInValidContent" class="error">Contentを入力してください</p>
    <button @click="createContent">送信する</button>
    <h2>登録データ一覧</h2>
    <thead v-pre>
        <tr>
            <th class="id">ID</th>
            <th class="content">Content</th>
            <th class="create_date">作成日時</th>
            <th class="update_date">更新日時</th>
            <th class="delete">削除</th>
            <th class="delete">編集</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in todos" :key="todo.content">
          <td>{{todo.ID}}</td>
          <td>{{todo.Content}}</td>
          <td>{{todo.CreatedAt}}</td>
          <td>{{todo.UpdatedAt}}</td>
          <td class="delete">
            <button @click="deleteTodo(todo)">
              削除
            </button>
          </td>
          <td>
          <router-link
            :to="{name: 'edit', params: {id:(Number(todo.ID))}}"
          >編集</router-link>
          </td>
        </tr>
      </tbody>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      content: "",
      todos: [],
      errors: [],
    }
  },
  created() {
    axios.get('getAllTodos')
    .then(response => {
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        var resultTodos = response.data
        this.todos = resultTodos
      }
    })
  },
  computed: {
    isInValidContent() {
      return this.content < 1;
    }
  },
  methods: {
    getAllTodos() {
      axios.get('/getAllTodos')
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultTodos = response.data
          this.todos = resultTodos
        }
      })
    },
    createContent() {
      const params = new URLSearchParams()
      params.append('content', this.content)
      axios.post('/new', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          // 全件取得する
          this.getAllTodos()
          this.content = ''
        }
      })
    },
    deleteTodo(todo) {
      const params = new URLSearchParams()
      params.append('todoId', todo.ID)
      console.log(params)
      axios.post('deleteTodo', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          this.getAllTodos()
        }
      })
    }
  }
}
</script>
