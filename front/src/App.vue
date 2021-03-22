<template>
  <div class="big-daddy">
    <Menu></Menu>
    <router-view/>
  </div>
</template>
<script>
import Menu from "@/components/Menu";
export default {
  components: {Menu},
  data() {
    return {
      backgrounds: {
        spring: require("@/assets/image/Spring/bg.png"),
        autumn: require("@/assets/image/Autumn/bg.png"),
        christmas: require("@/assets/image/Christmas/bg.png"),
        children: require("@/assets/image/Children/bg.png"),
        wall: require("@/assets/image/bg.png")
      },
      backgroundsArray: ["spring", "autumn", "christmas", "children", "wall"],
      current: -1
    }
  },
  methods: {
    setBackGround(img) {
      const body = document.querySelector("body");
      body.setAttribute("style", "background-image: url('" + img + "')")
    }
  },
  created() {
    const date = new Date();
    switch (date.getMonth()) {
      case 0:
      case 1:
        this.setBackGround(this.backgrounds.spring);
        this.current = 0;
        break;
      case 8:
      case 9:
      case 10:
        this.setBackGround(this.backgrounds.autumn);
        this.current = 1;
        break;
      case 11:
        this.setBackGround(this.backgrounds.christmas);
        this.current = 2;
        break;
      case 5:
        this.setBackGround(this.backgrounds.children);
        this.current = 3;
        break;
      default:
        this.setBackGround(this.backgrounds.wall);
        this.current = 4;
    }
    document.oncontextmenu = () => {
      this.current = (this.current + 1) % this.backgroundsArray.length;
      const img = this.backgrounds[this.backgroundsArray[this.current]];
      this.setBackGround(img);
      return false;
    };
  },
}
</script>
<style>
.big-daddy {
  max-width: 1200px;
  margin: 0 auto;
}
</style>
