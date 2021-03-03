<template>
  <div class="box">
    <h1>Hello World</h1>
    <p>学习笔记，工作笔记，其他感想，也许还会上传其他的。</p>
    <p>不一定什么时候更新</p>
    <p class="name">by @duqy@sitech</p>
    <img src="@/assets/image/decoration.png" alt="没了">
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
    }
  },
  created() {
    document.title = "DooQY的涂鸦墙"
  }
}
</script>

<style scoped>
.box {
  padding-top: 5vh;
  text-align: center;
  background-color: rgba(256,256,256,.85);
  min-height: calc(100vh - 48px);
}

h1 {
  font-size: 50px;
}

p {
  font-size: 32px;
}

.name {
  font-size: 26px;
  margin-right: 10vw;
  text-align: right;
}

img {
  width: 80%;
}
</style>