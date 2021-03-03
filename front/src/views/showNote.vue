<template>
  <div class="wrapper">
    <VMenu v-model:directory="directory" class="dir" @changeNote="getNote"></VMenu>
    <div id="content" v-html="noteContent"></div>
  </div>

</template>

<script>
import VMenu from "@/components/VMenu";
import Markdown from "markdown-it";
import hljs from "highlight.js";

export default {
  name: "showNote",
  components: {VMenu},
  data() {
    return {
      directory: '',
      md: {},
      noteContent:
          `<h1>关于网站</h1>
          <p>笔记是平时学习和工作的总结或者思考</p>
          <p>原本都是 <code>markdown</code> 格式的，感谢 <code>markdown-it</code> 和 <code>highlight.js</code> 的帮助，可以把 <code>markdown</code> 转化为 html 格式。</p>
          <p>以及 <code>VNote</code> 提供的 markdown 笔记管理功能。</p>
          <p>另外还有，前端的Vue框架，在学习笔记里有记载学习的过程。<b>尤大牛逼ヾ(^Д^*)/</b> </p>
          <p>为了节约服务器资源，后端用的 <code>golang</code> 原生 http 包</p>
          <p>最后，网站还在建设中，未来会继续完善的。</p>
            <br/><br/>
            <h1 style="text-align: center;color: red">最后祝所有人圣诞快乐</h1>`
    }
  },
  methods: {
    initData() {
      this.directory = this.$route.params.path;
    },
    initMarkdown() {
      this.md = new Markdown({
        highlight: (str, lang) => {
          if (lang && hljs.getLanguage(lang)) {
            try {
              return hljs.highlight(lang, str).value;
            } catch (e) {
              console.log(e);
            }
          }
        }
      });
    },
    async getNote(url) {
      let response = await this.axios.get(url);
      document.body.scrollTop = 0;
      document.documentElement.scrollTop = 0;
      if (response) {
        this.noteContent = this.md.render(response.data);
      }
      let index = url.lastIndexOf('/');
      this.handleImage(url.substring(0, index + 1));
    },
    handleImage(path) {
      this.noteContent = this.noteContent.replace(/(<img.+src=")(.+?)"/g, '$1/api' + path + '$2"');
    }
  },
  created() {
    this.initData();
    this.initMarkdown();
  },
  watch: {
    '$route'() {
      this.initData();
    }
  }
}
</script>

<style scoped>
.wrapper {
  width: 100%;
}

.wrapper > #content {
  width: calc(100% - 200px);
  min-height: calc(100vh - 48px);
  display: inline-block;
  padding: 4vh 4vw;
  font-size: 16px;
  background-color: rgba(256, 256, 256, .85);
  position: relative;
  left: 200px;
}

.wrapper > .dir {
  width: 200px;
  height: calc(100vh - 48px);
  background-color: rgba(256, 256, 256, .85);
  position: fixed;
  top: 48px;
}
</style>