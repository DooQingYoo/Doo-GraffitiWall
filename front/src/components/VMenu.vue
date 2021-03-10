<template>
  <a-menu mode="vertical" @click="getNote">
    <a-sub-menu v-for="subMenu in subDir" :key="subMenu.name" :title="subMenu.name">
      <a-menu-item v-for="item in subMenu.files" :key="item.name">{{ item.name }}</a-menu-item>
    </a-sub-menu>
    <a-menu-item v-for="item in files" :key="item.name">{{ item.name }}</a-menu-item>
  </a-menu>
</template>

<script>
export default {
  name: "VMenu",
  data() {
    return {
      subDir: [],
      files: []
    }
  },
  props: {
    directory: '',
  },
  methods: {
    async updateMenu(dir) {
      this.subDir = [];
      const list = await this.axiosGet("/" + dir);
      // 竖向菜单的一级菜单
      for (const element of list) {
        // 拥有二级菜单
        if (element.dir) {
          this.subDir.push({
            name: element.name,
            files: []
          })
          // 直接就是文件
        } else {
          this.files.push({
            // 去掉 .md 后缀
            name: element.name.substring(0, element.name.length - 3)
          })
        }
      }

      // 逐个查询二级菜单的内容
      for (const element of this.subDir) {
        const key = element.name;
        const list = await this.axiosGet(`/${dir}/${key}`);
        for (const listEle of list) {
          element.files.push({
            name: listEle.name.substring(0, listEle.name.length - 3)
          })
        }
      }
    },
    getNote(param) {
      let url = '/' + this.directory
      if (param.keyPath.length > 1) {
        url += '/' + param.keyPath[1];
      }
      url += '/' + param.keyPath[0] + '.md';
      document.title = param.keyPath[0];
      this.$emit('changeNote', url);
    },
  },
  watch: {
    directory(newVal) {
      this.updateMenu(newVal);
    }
  },
  mounted() {
    this.updateMenu(this.directory);
  }
}
</script>

<style scoped>
.ant-menu {
  font-size: 16px;
  background-color: transparent;
}
</style>