<template>
  <div>
    <h1>編集ページ</h1>
    <label for="content">今日のContent：</label>
    <!-- v-model="content"のみで出力できる -->
    <input
      type="text"
      v-model="content"
      >
    <button @click="updateContent">編集する</button>
    <button @click="back">戻る</button>
  </div>
</template>

<script>
import axios from 'axios'

export default{
  data() {
    return {
      content: "",
      todos: [],
    }
  },
  props:["id"],
  created() {
    this.fetchTodo()
  },
  methods: {
    fetchTodo() {
      const params = new URLSearchParams()
      params.append('id', this.id)
      axios.post('getOneTodo', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultTodo = response.data
          // Contentの値だけ取得する
          this.content = resultTodo[0].Content
        }
      })
    },
    updateContent() {
      const params = new URLSearchParams()
      params.append('content', this.content)
      params.append('id', this.id)
      axios.post("updateTodo", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          // アラートを表示してから/へ戻る感じ。
          alert("編集しました。")
          this.$router.push('/')
        }
      })
    },
    back() {
      this.$router.push('/')
    }
  },
}

</script>
