<template>
  <div class="box">
    <h1>Hello World</h1>
    <p>学习笔记，工作笔记，其他感想，也许还会上传其他的。</p>
    <p><del>不一定什么时候更新</del></p>
    <p>实时更新哦</p>
    <p class="name">by @duqy@sitech</p>
    <img class="decoration" src="" alt="没了">
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      decorations: {
        spring: require("@/assets/image/Spring/decoration.png"),
        autumn: require("@/assets/image/Autumn/decoration.png"),
        christmas: require("@/assets/image/Christmas/decoration.png"),
        children: require("@/assets/image/Children/decoration.png"),
        wall: require("@/assets/image/decoration.png")
      },
      decorationsArray: ["spring", "autumn", "christmas", "children", "wall"],
      current: -1
    }
  },
  mounted() {
    document.title = "DooQY的涂鸦墙"
    document.querySelector(".decoration").addEventListener("contextmenu", this.changeImg);
    const date = new Date();
    switch (date.getMonth()) {
      case 0:
      case 1:
        this.setDecoration(this.decorations.spring);
        this.current = 0;
        break;
      case 8:
      case 9:
      case 10:
        this.setDecoration(this.decorations.autumn);
        this.current = 1;
        break;
      case 11:
        this.setDecoration(this.decorations.christmas);
        this.current = 2;
        break;
      case 5:
        this.setDecoration(this.decorations.children);
        this.current = 3;
        break;
      default:
        this.setDecoration(this.decorations.wall);
        this.current = 4;
    }
  },
  methods: {
    changeImg(e) {
      this.current = (this.current + 1) % this.decorationsArray.length;
      const decoration = this.decorations[this.decorationsArray[this.current]];
      this.setDecoration(decoration);
      // 阻止事件冒泡
      e.stopPropagation();
      // 阻止默认事件
      e.preventDefault();
    },
    setDecoration(img) {
      const image = document.querySelector(".decoration");
      image.setAttribute("src", img);
    }
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

.decoration {
  width: 80%;
}
</style>